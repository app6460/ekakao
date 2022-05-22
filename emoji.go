package emoji

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

var (
	csrfReg, _   = regexp.Compile("csrf-token\" content=\"(.+)\"")
	cryptoReg, _ = regexp.Compile("name=\"p\" value=\"(.+)\"")
)

type (
	Emoji struct {
		Email       string
		Password    string
		Cookies     []*http.Cookie
		Csrf        string
		CryptoToken string
	}

	AuthRes struct {
		Status      int    `json:"status"`
		Message     string `json:"message"`
		ContinueURL string `json:"continue_url"`
	}
)

func (e *Emoji) GetloginRes() {
	res, err := http.Get("https://accounts.kakao.com/login?continue=https%3A%2F%2Fe.kakao.com%2F")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	body := string(bodyBytes)
	csrf := csrfReg.FindStringSubmatch(body)
	crypto := cryptoReg.FindStringSubmatch(body)

	e.Csrf = csrf[1]
	e.CryptoToken = crypto[1]
	e.Cookies = append(e.Cookies, res.Cookies()...)
}

func (e *Emoji) GetTiara() {
	res, err := http.Get(GetTiaraUrl())

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	e.Cookies = append(e.Cookies, res.Cookies()...)
}

func (e *Emoji) Pad(data []byte) []byte {
	length := 16 - len(data)%16
	var b bytes.Buffer
	b.Write(data)
	b.Write(bytes.Repeat([]byte{byte(length)}, length))
	return b.Bytes()
}

func (e *Emoji) BytesToKey(data, salt []byte, output int) ([]byte, []byte) {
	key := make([]byte, 0)
	finalKey := make([]byte, 0)
	for len(finalKey) < output {
		var b bytes.Buffer
		b.Write(key)
		b.Write(data)
		b.Write(salt)
		sum := md5.Sum(b.Bytes())
		key = sum[:]
		finalKey = append(finalKey, key...)
	}
	return finalKey[:32], finalKey[32:output]
}

func (e *Emoji) AESEncrypt(message, passphrase string) string {
	salt := make([]byte, 8)
	rand.Read(salt)

	key, iv := e.BytesToKey([]byte(passphrase), salt, 48)
	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}
	msg := e.Pad([]byte(message))
	res := make([]byte, len(msg))

	cipher.NewCBCEncrypter(block, iv).CryptBlocks(res, msg)

	var b bytes.Buffer
	b.Write([]byte("Salted__"))
	b.Write(salt)
	b.Write(res)
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func (e *Emoji) GetAuth() AuthRes {
	email := e.AESEncrypt(e.Email, e.CryptoToken)
	pass := e.AESEncrypt(e.Password, e.CryptoToken)

	data := url.Values{}
	data.Add("k", "true")
	data.Add("os", "web")
	data.Add("lang", "ko")
	data.Add("email", email)
	data.Add("password", pass)
	data.Add("webview_v", "2")
	data.Add("third", "false")
	data.Add("authenticity_token", e.Csrf)
	data.Add("continue", "https://e.kakao.com/")

	req, _ := http.NewRequest("POST", "https://accounts.kakao.com/weblogin/authenticate.json", bytes.NewBuffer([]byte(data.Encode())))

	for _, v := range e.Cookies {
		req.AddCookie(v)
	}

	req.Header.Add("Referer", "https://accounts.kakao.com/login?continue=https%3A%2F%2Fe.kakao.com%2F")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	e.Cookies = append(e.Cookies, res.Cookies()...)

	body, _ := ioutil.ReadAll(res.Body)
	auth := AuthRes{}
	json.Unmarshal(body, &auth)

	return auth
}

func New(email, pass string) *Emoji {
	emoji := Emoji{}
	emoji.Email = email
	emoji.Password = pass
	return &emoji
}

func (e *Emoji) Login() {
	e.GetloginRes()
	e.GetTiara()
	e.GetAuth()
}

func (e *Emoji) SendEmoji(name string, id int) {
	data, _ := json.Marshal(map[string]interface{}{
		"agree":    "Y",
		"idx":      id,
		"titleUrl": name,
	})

	req, _ := http.NewRequest("POST", "https://e.kakao.com/api/v1/previews/send-preview-message", bytes.NewBuffer(data))

	for _, v := range e.Cookies {
		req.AddCookie(v)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}

package ekakao

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/app6460/webkakao"
)

type (
	emoji struct {
		email    string
		password string
		cookies  []*http.Cookie
	}
)

func New(email, pass string) *emoji {
	emoji := emoji{}
	emoji.email = email
	emoji.password = pass
	return &emoji
}

func (e *emoji) Login() {
	client := webkakao.New(e.email, e.password, "https://e.kakao.com/")
	client.Login()
	e.cookies = client.Cookies()
}

func (e *emoji) SendEmoji(name string, id int) {
	data, _ := json.Marshal(map[string]interface{}{
		"agree":    "Y",
		"idx":      id,
		"titleUrl": name,
	})

	req, _ := http.NewRequest("POST", "https://e.kakao.com/api/v1/previews/send-preview-message", bytes.NewBuffer(data))

	for _, v := range e.cookies {
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

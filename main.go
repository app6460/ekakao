package emotpreview

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	csrfReg, _ = regexp.Compile("csrf-token\" content=\"(.+)\"")
)

type (
	Emoji struct {
		Email    string
		Password string
		Cookies  []*http.Cookie
		Csrf     string
	}
)

func (e *Emoji) getloginRes() {
	res, err := http.Get("https://accounts.kakao.com/login?continue=https%3A%2F%2Fe.kakao.com%2F")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	regex := csrfReg.FindStringSubmatch(string(body))

	e.Csrf = regex[1]
	e.Cookies = append(e.Cookies, res.Cookies()...)
}

func (e *Emoji) getTiara() {
	res, err := http.Get(getTiaraUrl())

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	e.Cookies = append(e.Cookies, res.Cookies()...)
}

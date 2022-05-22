package emoji

import (
	"encoding/json"
	"net/url"
)

func GetTiaraUrl() string {
	dt := GenerateRandomUUIDWithDateTime()
	dn := GenerateRandomUUIDWithDateNumber()
	data := map[string]interface{}{
		"sdk": map[string]interface{}{
			"type":    "WEB",
			"version": "1.1.20",
		},
		"env": map[string]interface{}{
			"screen": "1920X1080",
			"tz":     "+9",
			"cke":    "Y",
		},
		"common": map[string]interface{}{
			"svcdomain":  "accounts.kakao.com",
			"deployment": "production",
			"url":        "https://accounts.kakao.com/login",
			"referrer":   "https://m.search.daum.net/",
			"title":      "카카오계정",
			"section":    "login",
			"page":       "pageLogin",
		},
		"etc": map[string]interface{}{
			"client_info": map[string]interface{}{
				"tuid": dt,
				"tsid": dt,
				"uuid": dn,
				"suid": dn,
			},
		},
		"action": map[string]interface{}{
			"type": "Pageview",
			"name": "pageLogin",
			"kind": "",
		},
	}
	res, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	param := url.Values{}
	param.Add("d", string(res))

	return "https://stat.tiara.kakao.com/track?" + param.Encode()
}

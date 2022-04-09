package lib

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type RequestOption struct {
	Method string
	Url    string
	Header map[string]string
	Cookie string
	Body   *string
}

func Request(op *RequestOption) (string, error) {
	// fmt.Printf("(%%#v) %#v\n", op)
	req, _ := http.NewRequest(op.Method, op.Url, nil)

	// header処理
	for k,v := range op.Header {
		req.Header.Set(k, v)
	}

	for _, c := range cookieParse(op.Cookie) {
		req.AddCookie(&c)
	}

	client := new(http.Client)
	resp, e := client.Do(req); if e != nil {
		return "", e
	}

	defer resp.Body.Close()

	byteArray, e := ioutil.ReadAll(resp.Body); if e != nil {
		return  "", e
	}

	return string(byteArray), nil
}

func cookieParse(cookie string) []http.Cookie {
	var cookies []http.Cookie

	for _, s := range strings.Split(cookie, ";") {
		kv := strings.Split(s, "=")

		if len(kv) != 2 {
			continue
		}

		c := &http.Cookie{
			Name: kv[0],
			Value: kv[1],
		}

		cookies = append(cookies, *c)
	}
	return cookies
}

package lib

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type RequestOption struct {
	Method string
	Url    string
	Header map[string]string
	Cookie string
	Param  string
}

func Request(op *RequestOption) (string, error) {
	// fmt.Printf("(%%#v) %#v\n", op)
	var req *http.Request

	switch strings.ToUpper(op.Method) {
	case "GET":
		req, _ = http.NewRequest(op.Method, op.Url+"?"+op.Param, nil)
		break
	case "POST":
		req, _ = http.NewRequest(op.Method, op.Url, bytes.NewBuffer([]byte(op.Param)))
		break
	default:
		return "", errors.New("method type is require GET or POST")
	}

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

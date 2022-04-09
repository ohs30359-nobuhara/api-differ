package lib

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
)

func xmlSort(s string) (string, error) {
	var o interface{}

	if err := xml.Unmarshal([]byte(s), &o); err != nil {
		return "",err
	}

	out, err := xml.Marshal(o); if err != nil {
		return "", err
	}

	return string(out), nil
}

func jsonSort(s string) (string, error) {
	var o interface{}

	// 一度json obに変換することでkeyでsortする
	if err := json.Unmarshal([]byte(s), &o); err != nil {
		return "", err
	}
	out, err := json.Marshal(o); if err != nil {
		return "", err
	}

	var b bytes.Buffer
	json.Indent(&b, out, "", "  ")

	return b.String(), nil
}

func Sort(s string) string {
	if out, e := jsonSort(s); e == nil {
		return out
	}

	if out, e := xmlSort(s); e == nil {
		return out
	}

	return s
}

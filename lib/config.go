package lib

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Expect Target `json:expect`
	Actual Target `json:actual`
	Scenario Scenario `json:scenario`
}

type Target struct {
	Url string `json:url`
	Header map[string]string `json:header`
	Cookie string `json:cookie`
}

type Scenario struct {
	Method string `json:method`
	Type string `json:type`
	Params []string `json:params`
}

func LoadConfig(file string) Config {
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var conf Config
	if err := json.Unmarshal(bytes, &conf); err != nil {
		panic(err)
	}

	return conf
}

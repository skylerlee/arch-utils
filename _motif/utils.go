package main

import (
	"encoding/json"
	"io/ioutil"
)

type Conf struct {
	GistId string `json:"gistId"`
	Token  string `json:"token"`
}

func LoadConf(filename string) Conf {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	cfg := Conf{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

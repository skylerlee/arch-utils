package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Conf struct {
	GistID string `json:"gistId"`
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

type Filter func(*http.Request) *http.Request

type Client struct {
	client http.Client
	filter Filter
}

func (c *Client) preprocess(req *http.Request) *http.Request {
	if c.filter != nil {
		req = c.filter(req)
	}
	return req
}

func (c *Client) Get(url string, v interface{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req = c.preprocess(req)
	resp, err := c.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		panic(err)
	}
}

func (c *Client) Patch(url string, v interface{}) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PATCH", url, buf)
	if err != nil {
		panic(err)
	}
	req = c.preprocess(req)
	resp, err := c.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(resp.Status)
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Conf represents the gist config
type Conf struct {
	GistID string `json:"gistId"`
	Token  string `json:"token"`
}

// LoadConf loads the Conf from a given filename
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

// RequestFilter represents a preprocessor of the http.Request
type RequestFilter func(*http.Request) *http.Request

// Client wraps the http.Client for better JSON processing
type Client struct {
	client http.Client
	Filter RequestFilter
}

func (c *Client) preprocess(req *http.Request) *http.Request {
	if c.Filter != nil {
		req = c.Filter(req)
	}
	return req
}

func (c *Client) dumpData(data interface{}) (rd io.Reader, err error) {
	if data == nil {
		return
	}
	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(data)
	if err != nil {
		return
	}
	rd = buf
	return
}

func (c *Client) loadResult(result interface{}, rd io.Reader) (err error) {
	err = json.NewDecoder(rd).Decode(result)
	return
}

// Request sends and receives JSON data
func (c *Client) Request(method, url string, data interface{}, result interface{}) (err error) {
	body, err := c.dumpData(data)
	if err != nil {
		return
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req = c.preprocess(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(resp.Status)
	}
	err = c.loadResult(result, resp.Body)
	return
}

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Conf represents the gist config
type Conf struct {
	GistID string `json:"gistId"`
	Token  string `json:"token"`
}

// LoadConf loads the Conf from a given filename
func LoadConf(filename string) (cfg Conf, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	tmp := Conf{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return
	}
	cfg = tmp
	return
}

// RequestFilter represents a preprocessor of the http.Request
type RequestFilter func(*http.Request) *http.Request

// Client wraps the http.Client for better JSON processing
type Client struct {
	client http.Client
	Filter RequestFilter
	Base   *url.URL
}

// NewClient creates a Client using a base URL
func NewClient(basePath string) *Client {
	base, err := url.Parse(basePath)
	if err != nil {
		panic(err)
	}
	return &Client{Base: base}
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
	reqURL, err := c.Base.Parse(url)
	if err != nil {
		return
	}
	req, err := http.NewRequest(method, reqURL.String(), body)
	if err != nil {
		return
	}
	req = c.preprocess(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}
	err = c.loadResult(result, resp.Body)
	return
}

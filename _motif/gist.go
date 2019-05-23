package main

import (
	"net/http"
	"path"
)

// GistFile represents a file in gist
type GistFile struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

// Gist represents a gist
type Gist struct {
	Description string              `json:"description"`
	Files       map[string]GistFile `json:"files"`
}

// NewGist creates an empty Gist
func NewGist() *Gist {
	return &Gist{"", make(map[string]GistFile)}
}

// GetGist reads the gist from a given gistID
func (c *Client) GetGist(gistID string) (ret *Gist, err error) {
	url := path.Join("/gists", gistID)
	tmp := &Gist{}
	err = c.Request(http.MethodGet, url, nil, tmp)
	if err != nil {
		return
	}
	ret = tmp
	return
}

// PatchGist updates the associated gist by gistID
func (c *Client) PatchGist(gistID string, gist *Gist) (ret *Gist, err error) {
	url := path.Join("/gists", gistID)
	tmp := &Gist{}
	err = c.Request(http.MethodPatch, url, gist, tmp)
	if err != nil {
		return
	}
	ret = tmp
	return
}

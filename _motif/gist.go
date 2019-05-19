package main

const (
	baseURL = "https://api.github.com"
)

type GistFile struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

type Gist struct {
	Description string              `json:"description"`
	Files       map[string]GistFile `json:"files"`
}

func NewGist() *Gist {
	return &Gist{"", make(map[string]GistFile)}
}

func (c *Client) GetGist(gistId string) *Gist {
	url := baseURL + "/gists/" + gistId
	gist := &Gist{}
	c.Get(url, gist)
	return gist
}

func (c *Client) PatchGist(gistId string, gist *Gist) {
	url := baseURL + "/gists/" + gistId
	c.Patch(url, gist)
}

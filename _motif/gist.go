package main

const (
	baseURL = "https://api.github.com"
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

func (c *Client) GetGist(gistId string) (ret *Gist, err error) {
	url := baseURL + "/gists/" + gistId
	tmp := &Gist{}
	err = c.Request("GET", url, nil, tmp)
	if err != nil {
		return
	}
	ret = tmp
	return
}

func (c *Client) PatchGist(gistId string, gist *Gist) (ret *Gist, err error) {
	url := baseURL + "/gists/" + gistId
	tmp := &Gist{}
	err = c.Request("PATCH", url, gist, tmp)
	if err != nil {
		return
	}
	ret = tmp
	return
}

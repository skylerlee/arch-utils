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

func GetGist(gistId string) *Gist {
	url := baseURL + "/gists/" + gistId
	gist := &Gist{}
	client := Client{}
	client.Get(url, gist)
	return gist
}

func PatchGist(gistId string, gist *Gist) bool {
	url := baseURL + "/gists/" + gistId
	client := Client{}
	return client.Patch(url, gist)
}

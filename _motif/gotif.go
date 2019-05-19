package main

import (
	"fmt"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	conf := LoadConf("conf/gist.json")
	client := Client{}
	// client.Filter = func(req *http.Request) *http.Request {
	// 	req.Header.Set("Authorization", "token "+conf.Token)
	// 	return req
	// }
	gist := NewGist()
	gist.Files["savebox.zenc.txt"] = GistFile{"savebox.zenc.txt", ""}
	client.PatchGist(conf.GistID, gist)
}

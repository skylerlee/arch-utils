package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help    bool
	listing bool
	target  string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: gotif [OPTION...]")
		flag.PrintDefaults()
	}
	flag.BoolVar(&help, "h", false, "print help message")
	flag.BoolVar(&listing, "l", false, "list gist files")
	flag.StringVar(&target, "o", "", "open a gist file")
}

// UsageError means there is a problem with command usage
type UsageError struct {
	message string
}

// NewUsageError creates an UsageError with a message
func NewUsageError(message string) *UsageError {
	return &UsageError{message}
}

func (e *UsageError) Error() string {
	return e.message
}

func handleError() {
	if err := recover(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		switch err.(type) {
		case *UsageError:
			flag.Usage()
		}
		os.Exit(1)
	}
}

func main() {
	defer handleError()
	flag.Parse()
	if flag.NFlag() < 1 {
		panic(NewUsageError("not enough arguments"))
	}
	if help {
		flag.Usage()
		return
	}
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

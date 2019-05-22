package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
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

func process() {
	conf, err := LoadConf("conf/gist.json")
	if err != nil {
		panic(err)
	}
	client := Client{}
	if conf.Token != "" {
		client.Filter = func(req *http.Request) *http.Request {
			req.Header.Set("Authorization", "token "+conf.Token)
			return req
		}
	}
	switch {
	case listing:
		gist, err := client.GetGist(conf.GistID)
		if err != nil {
			panic(err)
		}
		for name := range gist.Files {
			fmt.Println("*", name)
		}
	case target != "":
		gist, err := client.GetGist(conf.GistID)
		if err != nil {
			panic(err)
		}
		targetFile, ok := gist.Files[target]
		if !ok {
			panic(errors.New("file not found"))
		}
		fmt.Println(targetFile.Content)
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
	process()
}

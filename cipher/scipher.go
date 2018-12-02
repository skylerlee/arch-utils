package main

import (
	"flag"
)

var (
	password string
	output string
)

func init() {
	flag.StringVar(&password, "p", "", "password to be used")
	flag.StringVar(&output, "o", "", "output file path")
}

func main() {
	flag.Parse()
}

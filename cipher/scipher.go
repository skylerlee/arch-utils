package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	password string
	output   string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: cipher FILE -p <password> -o <path>")
		flag.PrintDefaults()
	}
	flag.StringVar(&password, "p", "", "password to be applied")
	flag.StringVar(&output, "o", "", "output file path")
}

func process(input string) {
	fmt.Println(input, password, output)
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
	input := flag.Arg(0)
	process(input)
}

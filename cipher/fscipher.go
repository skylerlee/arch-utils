package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	passwd string
	output string
)

func init() {
	flag.Usage = func() {
		printErr("Usage: cipher [OPTION...] FILE")
		flag.PrintDefaults()
	}
	flag.StringVar(&passwd, "p", "", "password to be applied")
	flag.StringVar(&output, "o", "", "file to write output\nUse - to write to standard output")
}

func printErr(message string) {
	fmt.Fprintln(os.Stderr, message)
}

func process(input string) {
	var file *os.File
	if input == "-" {
		// read from standard input
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(input)
		if err != nil {
			printErr(fmt.Sprintf("No such file: %s", input))
		}
	}
	source(file)
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		printErr("error: missing input file")
		flag.Usage()
		os.Exit(1)
	}
	input := flag.Arg(0)
	process(input)
}

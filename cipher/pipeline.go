package main

import (
	"bufio"
	"fmt"
	"io"
)

func source(reader io.Reader) {
	reader = bufio.NewReader(reader)
	buffer := make([]byte, 16)
	size := 0
	for {
		n, err := reader.Read(buffer)
		size += n
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("size: %d bytes\n", size)
}

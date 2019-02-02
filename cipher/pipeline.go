package main

import (
	"bufio"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
)

// Pipeline represents a cipher data flow
type Pipeline struct {
	s *crypto.Stream
}

// NewPipeline creates a cipher pipeline
func NewPipeline() Pipeline {
	return Pipeline{}
}

func (p *Pipeline) InitEncrypter(key string, iv []byte) {
	block, err = aes.NewCipher()
	p.s = cipher.NewCFBEncrypter(block, iv)
}

func (p *Pipeline) InitDecrypter(key string, iv []byte) {
	block, err = aes.NewCipher()
	p.s = cipher.NewCFBDecrypter(block, iv)
}

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

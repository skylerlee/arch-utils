package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
)

var (
	host string = "127.0.0.1"
	port string = "8080"
)

func init() {
}

func main() {
	path := os.Args[1]
	addr := fmt.Sprintf("%s:%s", host, port)
	server := http.FileServer(http.Dir(path))
	err := http.ListenAndServe(addr, server)
	log.Fatal(err)
}

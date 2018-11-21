package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "a", "127.0.0.1", "address to use")
	flag.IntVar(&port, "p", 8080, "port to bind to")
}

func main() {
	flag.Parse()
	path := os.Args[1]
	addr := fmt.Sprintf("%s:%d", host, port)
	handler := http.FileServer(http.Dir(path))
	http.Handle("/", handler)
	log.Printf("start service at: %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

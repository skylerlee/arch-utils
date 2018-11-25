package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	host string
	port int
)

type middleware func(http.Handler) http.Handler

func init() {
	flag.StringVar(&host, "a", "0.0.0.0", "address to use")
	flag.IntVar(&port, "p", 8080, "port to bind to")
}

func compose(h http.Handler, mws []middleware) http.Handler {
	for i := 0; i < len(mws); i++ {
		h = mws[i](h)
	}
	return h
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	addr := fmt.Sprintf("%s:%d", host, port)
	handler := compose(http.FileServer(http.Dir(root)), []middleware{logRequest})
	http.Handle("/", handler)
	log.Printf("start service at: %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

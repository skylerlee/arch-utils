package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var (
	host string
	port string
)

type middleware func(http.Handler) http.Handler

func init() {
	flag.StringVar(&host, "a", "0.0.0.0", "address to use")
	flag.StringVar(&port, "p", "8080", "port to bind to")
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

func LocalIP() (string, error) {
	addr, err := net.ResolveUDPAddr("udp", "1.2.3.4:1")
	if err != nil {
		return "", err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	ip, _, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		return "", err
	}
	return ip, nil
}

func printServerInfo(root string) {
	fmt.Println("Serving at:", root)
	ip, err := LocalIP()
	if host == "0.0.0.0" && err == nil {
		fmt.Println("Available on:", net.JoinHostPort(ip, port))
	}
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	addr := net.JoinHostPort(host, port)
	handler := compose(http.FileServer(http.Dir(root)), []middleware{logRequest})
	http.Handle("/", handler)
	printServerInfo(root)
	log.Fatal(http.ListenAndServe(addr, nil))
}

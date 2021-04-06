package main

import (
	"fmt"
	"log"

	"flag"

	"time"

	"net/http"
)

const (
	timeout    = time.Duration(1 * time.Second)
	timeoutMsg = "your request has timed out"
)

var (
	port int
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this request will always timeout!
	time.Sleep(timeout)
}

func init() {
	flag.IntVar(&port, "port", 8080, "HTTP Server Port")

	flag.Parse()
}

func main() {
	httpAddr := fmt.Sprintf(":%v", port)

	log.Printf("Listening to %v", httpAddr)

	rootHandler := &MyHandler{}

	// http://golang.org/pkg/net/http/#TimeoutHandler
	http.Handle("/", http.TimeoutHandler(rootHandler, timeout, timeoutMsg))
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}

package main

import (
	"fmt"
	"net/http"
)

// type ServeMux struct {
// 	mu sync.RWMutex        // because of concurrency, we have to use a mutex here
// 	m  map[string]muxEntry // router rules, every string mapping to a handler
// }

// type muxEntry struct {
// 	explicit bool // exact match or not
// 	h        Handler
// }

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request) // routing implementer
// }

// type HandlerFunc func(ResponseWriter, *Request)

// // ServeHTTP calls f(w, r).
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
// 	f(w, r)
// }

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}

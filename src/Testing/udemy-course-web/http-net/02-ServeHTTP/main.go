package main

import (
	"fmt"
	"net/http"
)

type nitu int

func (n nitu) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My apple of eye is Riddham")
}

func main() {

	var ni nitu

	http.ListenAndServe(":8080", ni)

}

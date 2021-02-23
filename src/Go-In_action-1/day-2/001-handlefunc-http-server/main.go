package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", SimpleServer)
	http.ListenAndServe(":8080", nil)

}

//SimpleServer is exported
func SimpleServer(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello", req.URL.Path[1:])
}

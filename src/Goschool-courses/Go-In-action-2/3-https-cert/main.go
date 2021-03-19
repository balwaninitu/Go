package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8082", "C:/Users/Lenovo/cert.pem", "C:/Users/Lenovo/key.pem", nil)

}

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", ExampleHandler)

	log.Println("** Service Started on Port 8080 **")

	// Use ListenAndServeTLS() instead of ListenAndServe() which accepts two extra parameters.
	// We need to specify both the certificate file and the key file (which we've named
	// https-server.crt and https-server.key).
	err := http.ListenAndServeTLS(":8080", "C:/Users/Lenovo/cert.pem", "C:/Users/Lenovo/key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There!")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":8081", "C:/Users/Lenovo/cert.pem", "C:/Users/Lenovo/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}

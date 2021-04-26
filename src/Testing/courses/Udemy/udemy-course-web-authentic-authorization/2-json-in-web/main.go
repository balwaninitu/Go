package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	/*encode and decode provide memory benefit compared to using json on fly
	because json is very big to transfer it is better to transfer single struct than big json*/
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}
	//json encoder needs a writer and we have response writer
	//after you got encoder encode actual data which is p1
	//encode return error
	err := json.NewEncoder(w).Encode(p1)
	if err != nil { //we dont do panic here coz dont want server to go down
		log.Println("Encoded bad data", err) //log is stdout, log by default is timestamp

	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	//for decode we need person/client to enter data in empty struct
	//var p1 person
	people := []person{}
	//json decoder takes request so we put r with body and decode to pointer p1, returns error
	//err := json.NewDecoder(r.Body).Decode(&p1)
	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil { //we dont do panic here coz dont want server to go down
		log.Println("Decoded bad data", err)
	}
	//log.Println("Person:", p1)
	log.Println("People:", people)

}

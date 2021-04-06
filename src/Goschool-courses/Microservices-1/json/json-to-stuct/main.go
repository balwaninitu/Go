package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	FirstName string
	Lastname  string
}

func main() {
	var person People
	jsonString := `{"firstname":"Mary", "lastname":"Daisy"}`
	err := json.Unmarshal([]byte(jsonString), &person)
	fmt.Println(person)
	fmt.Println(err)
}

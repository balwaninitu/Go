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
	var persons []People
	jsonString :=
		`[

		{
			"firstname":"Mary",
			 "lastname":"Daisy"
			},
		{
			"firstname":"Janson", 
			"lastname":"Wong"
		}
]`

	err := json.Unmarshal([]byte(jsonString), &persons)
	fmt.Println(persons)
	fmt.Println(err)
}

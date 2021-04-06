package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	FirstName string
	LastName  string
	Details   struct {
		Height int
		Weight float32
	}
}

type Rates struct {
	Base   string `json:"base currency"`
	Symbol string `json:"destination currency"`
}

func main() {
	var persons []People
	jsonString :=
		`[

		{
			"firstname":"Mary",
			 "lastname":"Daisy",
			 "Details": {
				 "height":154,
				 "weight":55.3
			 }
			},
		{
			"firstname":"Janson", 
			"lastname":"Wong",
			"Details": {
				"height":154,
				"weight":55.3
			}
		}
]`

	err := json.Unmarshal([]byte(jsonString), &persons)
	for _, v := range persons {
		fmt.Println(v.FirstName)
		fmt.Println(v.LastName)
		fmt.Println(v.Details.Height)
		fmt.Println(v.Details.Weight)
	}
	fmt.Println(err)
	fmt.Println(persons)

	jsonString2 :=
		`{
		"base currency" : "EUR",
		"destination currency": "USD"
	}`

	var rates Rates
	json.Unmarshal([]byte(jsonString2), &rates)
	fmt.Println(rates.Base)
	fmt.Println(rates.Symbol)

}

package main

import (
	"fmt"
)

type details struct {
	id           int
	name         string
	availableDay string
}
type doctor struct {
	details
	specialisation string
}

func main() {
	doctors := []doctor{
		{details: details{id: 1, name: "dr1", availableDay: "Mon"}, specialisation: "Periodontist"},
		{details: details{id: 2, name: "dr2", availableDay: "Tue"}, specialisation: "Endodontist"},
		{details: details{id: 3, name: "dr3", availableDay: "Wed"}, specialisation: "Orthodontist"},
		{details: details{id: 4, name: "dr4", availableDay: "Thu"}, specialisation: "General Dentist"},
		{details: details{id: 5, name: "dr5", availableDay: "Thu"}, specialisation: "Prosthodontist"},
	}

	var slice []int
	for i := range doctors {

		slice = append(slice, i)
	}

	searchDrSlice := appendSlice(doctors)
	fmt.Println(searchDrSlice)

	search(searchDrSlice, len(searchDrSlice), 6)

}

func appendSlice(doctors []doctor) []int {
	var s []int
	var i int
	for i = range doctors {
		s = append(s, i)

	}
	return s
}

func search(s []int, n int, target int) error {
	var i int
	for i = range s {
		if i == target {
			fmt.Println("id found")

		}
	}
	return nil

}

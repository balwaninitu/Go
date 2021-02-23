package main

import "fmt"

type celebrity struct {
	name     string
	age      int
	coActors []string
	movies   []string
}

func (c celebrity) speak() {
	fmt.Println(c.name, `says, "Hello, Nitu"`)
}

func main() {

	aCelebrity := celebrity{

		name:     "vin diesel",
		age:      55,
		coActors: []string{"paul walker", "Gal Gadok"},
		movies:   []string{"f&f 1", "f&f2"},
	}

	fmt.Println(aCelebrity)
	fmt.Println(aCelebrity.movies[0], aCelebrity.name, aCelebrity.coActors[1])

	aCelebrity.speak()

}

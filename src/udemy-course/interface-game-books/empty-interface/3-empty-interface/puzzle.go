package main

import "fmt"

type puzzle struct {
	title string
	price money
}

func (p puzzle) print() {

	fmt.Printf("%s %s\n", p.title, p.price.string())
}

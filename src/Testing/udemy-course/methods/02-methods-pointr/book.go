package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b book) print() {

	fmt.Printf("%s %.2f\n", b.title, b.price)
}

func (b book) discount(ratio float64) {

	b.price *= (1 - ratio)
}

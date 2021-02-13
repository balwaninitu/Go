package main

import "fmt"

type book struct {
	title string
	price float64
}

// func printBook(b book) {

// 	fmt.Printf("%s: %.2f\n", b.title, b.price)
// }

func (b book) print() {

	fmt.Printf("%s: %.2f\n", b.title, b.price)

}

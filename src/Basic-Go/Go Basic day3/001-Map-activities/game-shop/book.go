package main

import "fmt"

type book struct {
	title string
	price float32
}

func (b *book) print() {
	fmt.Println(b.title, b.price)
}

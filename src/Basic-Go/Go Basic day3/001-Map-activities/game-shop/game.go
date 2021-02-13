package main

import "fmt"

type game struct {
	title string
	price float32
}

func (g *game) print() {

	fmt.Println(g.title, g.price)
}

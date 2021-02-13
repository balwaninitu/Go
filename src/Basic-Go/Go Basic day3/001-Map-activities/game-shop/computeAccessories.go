package main

import "fmt"

type computerAccessories struct {
	title string
	price float32
}

func (c *computerAccessories) print() {
	fmt.Println(c.title, c.price)

}

package main

import (
	"fmt"
)

type printer interface {
	print()
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Error")
		return
	}
	for _, it := range l {
		fmt.Printf("(%-10T) --> ", it)
		it.print()
	}
}

func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64)
	}
	for _, it := range l {
		fmt.Printf("%T\n", it)
		if it, ok := it.(discounter); ok {
			it.discount(ratio)

		}

	}
}

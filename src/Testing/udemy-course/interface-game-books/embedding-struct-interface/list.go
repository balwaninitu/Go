package main

import (
	"fmt"
)

type item interface {
	print()
	discount(ratio float64)
}

type list []item

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

}

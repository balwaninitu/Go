package main

import (
	"fmt"
)

type list []*product

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Error")
		return
	}
	for _, p := range l {
		fmt.Printf("(%-10T) --> ", p)
		p.print()
	}
}

func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64)
	}

}

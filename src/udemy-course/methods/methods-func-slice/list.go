package main

import "fmt"

type list []*game

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Error")
		return
	}
	for _, it := range l {
		it.print()
	}
}

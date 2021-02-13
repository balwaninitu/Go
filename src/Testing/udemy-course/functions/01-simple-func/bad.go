package main

import "fmt"

func showN() {

	if n == 0 {
		return
	}
	fmt.Printf("ShowN       : N is %d\n", n)
}

func incrN() {

	n++
}

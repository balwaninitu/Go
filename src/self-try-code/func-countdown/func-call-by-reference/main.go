package main

import "fmt"

func swap(a, b *int) int {
	var o int
	o = *a
	*a = *b
	*b = o

	return 0
}

func main() {

	var p int = 6
	var q int = 10

	fmt.Println(p, q)
	swap(&p, &q)

	fmt.Println(p, q)
}

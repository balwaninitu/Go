package main

import (
	"fmt"
)

// sort through function
// func smallestNum(s []int) int {

// 	sort.Ints(s)
// 	return s[0]

// }

//find smallest number in slice

func main() {

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	// smallestNum(x)
	// fmt.Println(x[0])

	// sort through algorithm
	n := x[0]
	for _, v := range x[1:] {

		if v < n {
			n = v

		}

	}

	fmt.Println(n)

}

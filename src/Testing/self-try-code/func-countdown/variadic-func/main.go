package main

import "fmt"

func add(slice ...int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum

}

func main() {

	slice1 := []int{4, 7, 2, 1}

	fmt.Println(add(slice1...))
	fmt.Println(add(1, 2, 3))

}

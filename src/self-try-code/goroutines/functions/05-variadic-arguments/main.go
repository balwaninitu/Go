package main

import "fmt"

func find(num int, nums ...int) {
	found := false
	index := 0
	for i, v := range nums {
		if num == v {
			found = true
			index = i
		}

	}
	if found {
		fmt.Printf("%d  found at index %d\n", num, index)
	} else {
		fmt.Printf("%d is not present\n", num)
	}
}

func main() {

	slice := []int{5, 7, 8, 3, 2, 6}

	find(5, slice...)
	find(10, slice...)

}

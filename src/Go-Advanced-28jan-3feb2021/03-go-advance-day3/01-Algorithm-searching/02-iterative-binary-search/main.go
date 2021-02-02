package main

import "fmt"

func search(slice []int, n int, target int) int {
	first := 0
	last := n

	for i := 0; i < n; i++ {
		if slice[i] == target {
			return i
		}else {
			return -1
		}

		if slice[i]

	}

	return -1
}

func main() {

	slice1 := []int{45, 78, 23, 12, 78, 90, 34, 56, 23, 67, 45}

	s := search(slice1, len(slice1), 78)

	fmt.Println(s)

}
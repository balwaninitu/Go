package main

import "fmt"

func search(arr []int, n int, start int, target int) int {
	if start > n-1 {
		return -1
	} else {
		if arr[start] == target {
			return start
		} else {
			if arr[start] > target {
				return -1
			} else {
				return search(arr, n, start+1, target)
			}
		}
	}
}

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := search(array, 10, 1, 15)

	fmt.Println(result)

}

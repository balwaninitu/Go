package main

import "fmt"

func findSum(vals ...int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}

func main() {

	nums := []int{1, 2, 3, 4}
	nums1 := []int{3, 5, 7, 9, 11, 12}

	fmt.Println(findSum())
	fmt.Println(findSum(5, 6, 7))
	fmt.Println(findSum(1, 2, 3))
	fmt.Println(findSum(0, 0, 0))
	fmt.Println(findSum(nums...))
	fmt.Println(findSum(nums1...))

}

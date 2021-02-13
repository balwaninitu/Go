package main

import "fmt"

func main() {

	slice()

}

func slice() {

	nums := []int{1, 2, 3}

	incrByPtr(nums)
	fmt.Println(nums)

}

func incrByPtr(nums []int) {
	for i := range nums {
		nums[i]++
	}
}

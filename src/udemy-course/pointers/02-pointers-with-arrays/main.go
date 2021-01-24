package main

import "fmt"

func main() {

	fmt.Println("Arrays")
	arrays()

}

func arrays() {

	nums := [...]int{1, 2, 3}

	incr(nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println("pointer", nums)

}

func incr(nums [3]int) {
	for i := range nums {
		nums[i]++
	}
}

func incrByPtr(nums *[3]int) {
	for i := range nums {
		nums[i]++
	}
}

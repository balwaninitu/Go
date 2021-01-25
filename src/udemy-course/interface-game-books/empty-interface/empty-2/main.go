package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	var any interface{}
	any = nums
	//_ = len(any)
	_ = len(any.([]int))

	var many []interface{}

	//many = nums --> wont work
	// var words[]string = nums --> is like type mismatch
	// solution is to loop over

	for _, n := range nums {
		many = append(many, n)
	}
	fmt.Println(many)

	//many is a slice of interface{}values.it can not store any type. only interface{}store any value.

}

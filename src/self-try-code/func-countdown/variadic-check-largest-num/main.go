package main

import "fmt"

//func to find greatest number in slice

func findLarge(s ...int) int {
	largest := s[0]
	for _, v := range s[1:] {

		if v > largest {
			largest = v
		}
	}

	return largest
}

//func to find greatest number in list

func findLargelist(args ...int) int {
	largest := int(0)
	for _, v := range args {

		if v > largest {
			largest = v
		}
	}

	return largest
}
func main() {

	// slice := []int{

	// 	45, 56, 78, 23,
	// 	67, 89, 12, 90, 23,
	// 	100, 56, 890, 234, 1234,
	// }

	// fmt.Println(findLarge(slice...))

	fmt.Println(findLargelist(23, 56, 23, 45, 12, 0, 7))

}

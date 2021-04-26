package main

import (
	"fmt"
)

/*Len is the number of elements in the collection.
  Len() int
  // Less reports whether the element with
  // index i should sort before the element with index j.
  Less(i, j int) bool
  // Swap swaps the elements with indexes i and j.
  Swap(i, j int)
*/
func main() {
	nums := []int{6, 2, 3, 4, 7, 11, 13, 20}
	i := 0
	j := len(nums) - 1 //len
	for i < j {        //less
		nums[i], nums[j] = nums[j], nums[i] //swap
		i++
		j--
	}
	fmt.Println(nums)
}

//fmt.Println(nums)
//reverse(nums)

// func reverse(nums []int) {
// 	i := 0
// 	j := len(nums) - 1 //len
// 	for i < j {        //less
// 		nums[i], nums[j] = nums[j], nums[i] //swap
// 		i++
// 		j--
// 	}
// }

//2nd approach
// package main

// import "fmt"

// //reverse array - Len, less , swap
// func main() {
// 	s := []int{3, 5, 6, 2, 1}

// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// 	fmt.Println(s)
// }

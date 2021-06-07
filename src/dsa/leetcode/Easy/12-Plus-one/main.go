/*

Given a non-empty array of decimal digits representing a non-negative integer, increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.
Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Example 3:

Input: digits = [0]
Output: [1]
*/
package main

import "fmt"

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		//if last element of slice is 9 it become 0
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			//otherwise 1 will get added to last element
			digits[i] += 1
			return digits
		}
	}
	//when slice is empty 1 will get append to it and slice get return
	//variadict func has been use to append slice, in this func it is allowed to pass 0 or more arguments
	digits = append([]int{1}, digits...)
	return digits
}

func main() {

	slice := []int{}

	ans := plusOne(slice)

	fmt.Println(ans)
}

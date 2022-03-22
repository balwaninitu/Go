package main

import "fmt"

/*
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.
For example, 121 is palindrome while 123 is not.



Example 1:

Input: x = 121
Output: true
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
Example 4:

Input: x = -101
Output: false


*/

//Now the question is, how do we know that we've reached the half of the number?

//Since we divided the number by 10, and multiplied the reversed number by 10,
//when the original number is less than the reversed number,
//it means we've processed half of the number digits.
//if first half is eqaul to second half its pallindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x <= 9 {
		return true
	} else if x%10 == 0 {
		return false
	}

	var y int
	for x > y {
		r := x % 10
		x = x / 10
		y = y*10 + r

		if x == y || x/10 == y {
			return true
		}
	}
	return false
}

func main() {

	ans := isPalindrome(1231)

	fmt.Println(ans)

}

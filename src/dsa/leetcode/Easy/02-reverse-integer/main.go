/*Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
Example 1:

Input: x = 123
Output: 321

Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21
Example 4:

Input: x = 0
Output: 0*/

package main

import (
	"fmt"
	"math"
)

// to pass signed test cases
//math.MaxInt32 = 1<<31 - 1
//math.MinInt32 = -1 << 31

func reverseInt(x int) int {
	n := 0

	//for x > 0{
	for x != 0 {
		if n > math.MaxInt32/10 || n < math.MinInt32/10 {
			return 0
		}

		//isolate last digit num by modulus 10
		//it return last digit as remainder
		//e.g incase of 1234, 1000%10, 200%10, 30%10 == no remainder only 4(unit place) will be remainder
		/*1st iteration:
		x = -1234
		n = 0
		remainder= -4,
		n has to move to next place from unit to ten place by multiplying 10
		n become when remainder added = -4,
		after x/10, x become -123


		2nd Iteration
		x = -123
		n = -4
		remainder = -3
		n= -4 will move to ten place by multiplication with 10
		remainder get added it become -4*10+(-3) = -43
		n = -43
		x = - 12*/

		remainder := x % 10
		//adding remainder to temp num and multiplying with 10 n(if n=0)*10+4=4
		n *= 10
		n += remainder
		//removing last number 1234/10 = 123.4 it gives nearest integer
		x /= 10
	}
	return n
}
func reverse(x int) int {
	out := int(0)
	for x != 0 {
		// check bound condition
		if out > math.MaxInt32/10 || out < math.MinInt32/10 {
			return 0
		}
		out = out*10 + int(x%10)
		x = int(x / 10)
	}
	return out
}

func main() {
	num := reverseInt(-1234)
	//num1 := reverse(-1234)
	fmt.Println(num)
	//fmt.Println(num1)
}

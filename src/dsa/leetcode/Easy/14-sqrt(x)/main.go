/*
Given a non-negative integer x, compute and return the square root of x.

Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.

Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.

Example 1:

Input: x = 4
Output: 2
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since the decimal part is truncated, 2 is returned.


Constraints:

0 <= x <= 231 - 1
In square root need to find nearest possible num e.g if need to find sqrt of 9 it is 3
sqrt of 10 will be nearest perfect sq 4*4=16 but 16 > 10 so 3 will be sqrt of 10
binary search method : It will start from min= 1 & max = x, mid = 1+x/2
if x = 5

for array :

Set first to start of array
Set last to end of array
While (first <= last)
set mid = (first + last) / 2
if item at mid is equal to target
item found (return mid)
else
if target is smaller than item at mid
set last = mid
1 // search first half
else
set first = mid + 1 // search second half
End of search, item not found (return
1)

func binarySearch arr []int, n int, target int) int
first := 0
last := n
for first <= last {
mid := (first + last) / 2
if
arr [mid] == target { // found
return mid
} else {
if target <
arr [mid]
last = mid
1 // search first half
} else {
first = mid + 1 // search second half
}
}
} //end of for
return
1 // not found
}

*/
package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	start, end := 1, x //suppose x = 10

	for {
		mid := start + end/2 //1+10/2 = 5 (approx)

		square := mid * mid //5*5 = 25

		if x < square { //if condition true, search will be from start to mid, 1 - 5
			end = mid

		} else if x > square { //if condition true, search will be from mid to end
			start = mid
		} else {
			return mid
		}

		if start == end {
			return start
		} else if start+1 == end { //1+1 = 2,
			if end*end < x { //2*2 = 4, return 2
				return end
			}
			return start
		}

	}
}

func main() {

	x := 5

	ans := mySqrt(x)

	fmt.Println(ans)

}

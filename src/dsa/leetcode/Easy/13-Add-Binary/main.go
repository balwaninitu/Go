/*Given two binary strings a and b, return their sum as a binary string.
Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"

The idea is to start from last characters of two strings and compute digit sum one by one.
 If sum becomes more than 1, then store carry for next digits.

 binary addition: 1+0 = 1, 0+0 = 0, 1+1 = 10(here 1+1 become 0 and need to carry over 1 hence 10)
 1+1+1 = 11(here 1+1 become 0 then carry 1 and next 0+1 = 1 so ans 11)

 carry1   1                carry 1    1
 +        1                      +    1
          1
 -------------                    ----------
       1  1                         1 0
*/
package main

import (
	"fmt"
	"strconv"
)

func addBinary1(a string, b string) string {

	if a == "" {
		return b
	} else if b == "" {
		return a
	}
	return a + b

}

func addBinary(a string, b string) string {
	carry, ans := 0, ""
	/*when a = 11, b = 1
	1st iteration:
	i = 1(len(a)-1) - last element
	j = 0 (len(b)-1) - last element

	*/
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		if sum >= 2 {
			carry = 1
			sum -= 2
		} else {
			carry = 0
		}
		ans = strconv.Itoa(sum) + ans
	}
	return ans
}

func main() {

	a := "11"

	b := "1"

	ans := addBinary(a, b)
	fmt.Println(ans)

}

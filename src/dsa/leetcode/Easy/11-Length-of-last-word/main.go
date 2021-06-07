/*
Given a string s consists of some words separated by spaces, return the length of the last word in the string. If the last word does not exist, return 0.

A word is a maximal substring consisting of non-space characters only.

 Example 1:

Input: s = "Hello World"
Output: 5
Example 2:

Input: s = " "
Output: 0

*/

package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	//convert string to array
	array := strings.Fields(s)

	lengthOfArray := len(array)

	if lengthOfArray == 0 {
		return 0
	}
	//find out last element in array
	lastWord := array[lengthOfArray-1]
	//return length of last element
	lenOfLastWord := len(lastWord)

	return lenOfLastWord

}

//without strings function
func lengthOfLastWord1(s string) int {
	var length int
	//start loop from last word in the string
	for i := len(s) - 1; i >= 0; i-- {
		//when s is empty space and length less than 1 loop will continue
		if s[i] == ' ' && length < 1 {
			continue
		}
		//once s is empty space and length more than 0 then loop will break means we found out length of last word
		if s[i] == ' ' && length > 0 {
			break
		}
		//loop will continue and length will be added after each iteration untill no empty space
		if s[i] != ' ' {
			length++
		}
	}
	return length
}

func main() {

	s := "Hello World Singapore"

	ans := lengthOfLastWord1(s)

	fmt.Println(ans)

}

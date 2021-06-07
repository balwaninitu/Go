/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().



Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Example 3:

Input: haystack = "", needle = ""
Output: 0

*/

package main

import "fmt"

func strStr(haystack string, needle string) int {

	lengthH := len(haystack)
	lengthN := len(needle)

	// temporary storage of index
	indexHay := 0
	indexNeedle := 0

	for i := 0; i < lengthH && indexNeedle < lengthN; i++ {

		if haystack[i] == needle[indexNeedle] {

			indexNeedle++

		} else {
			//increment of indexHay to iterate over second element if no match

			indexNeedle, indexHay, i = 0, indexHay+1, indexHay
		}
	}

	if indexNeedle == lengthN {

		return indexHay
	}

	return -1

}

func strStr1(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)

	// Corner cases
	if haystack == needle {
		return 0
	}

	if haystackLen <= 0 {
		return -1
	}

	if haystackLen < needleLen {
		return -1
	}

	// Doesn't make sense to search on latest M chars (needle)
	for i := 0; i < haystackLen-needleLen+1; i++ {
		found := true
		for j := 0; j < needleLen; j++ {
			if haystack[i+j] != needle[j] {
				// Save CPU, lookup till first mismatch!
				found = false
				break
			}
		}
		if found {
			// As per problem looking only for first match
			return i
		}
	}
	return -1
}

func main() {

	Haystack := "heello"
	Needdle := "ll"

	fmt.Println("len", len(Haystack))

	ans := strStr1(Haystack, Needdle)

	fmt.Println(ans)

}

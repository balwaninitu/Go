package main

import "fmt"

//Input: s = "xyzzaz"
//Output: 1
//Explanation: There are 4 substrings of size 3: "xyz", "yzz", "zza", and "zaz".
//The only good substring of length 3 is "xyz".

func main() {

	s := "aababcabc"
	//fmt.Println(s)

	//ans := countGoodSubstrings(s)

	//fmt.Println(ans)

	//intSlice := []int{1,5,3,6,9,9,4,2,3,1,5}
	//fmt.Println(s)
	uniqueSlice := unique(s)
	fmt.Println("unique", uniqueSlice)
}

func unique(s string) int {
	keys := make(map[rune]bool)
	list := 0
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			// list = append(list, entry)
			list += 1
		}
	}
	return list
}

func countGoodSubstrings(s string) int {
	//previous := s[0]
	sum := 0
	previous := s[0]
	next := s[1]
	for i := 2; i < len(s); {
		if s[i] == previous || s[i] == next {
			fmt.Println("s[i]", s[i])
			fmt.Println("previous", previous)
			i++
		} else {
			sum += 1
			fmt.Println("sum", sum)

		}
	}
	return sum
}

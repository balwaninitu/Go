package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	list1 := [4]string{"ans", "wer", "s", "of"}
	list2 := [4]string{"-", "+", "*", "/"}
	list3 := [4]string{"5", "10", "4", "0"}

	title := strings.Title(list1[0]) // make title case of list[0] i.e. ans to Ans

	concate := title + list1[1] // join Ans + wer to get Answer

	str1, _ := strconv.Atoi(list3[0]) // convert string to int
	str2, _ := strconv.Atoi(list3[2])

	result := str1 + str2

	s := "=" // missing in array

	fmt.Printf("%s %s %s %s %s %s %d\n", concate, list1[3], list3[2], list2[1], list3[0], s, result)

}

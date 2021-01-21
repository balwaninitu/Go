package main

import (
	"fmt"
	"strconv"
)

func main() {

	var res1, res2 string

	for {

		var n string

		fmt.Println("Enter number")
		fmt.Scanln(&n)

		num, _ := strconv.ParseInt(n, 10, 0)

		if num%2 == 0 && num >= 10 {

			res1 = ("Even number")
			res2 = ("more than one digit")

		} else if num%2 == 0 && num < 10 {

			res1 = ("Even number")
			res2 = ("less than one digit")

		} else if num%2 != 0 && num >= 10 {

			res1 = ("Odd num")
			res2 = ("more than one digit")

		} else if num%2 != 0 && num < 10 {

			res1 = ("Odd num")
			res2 = ("less than one digit")

		}
		fmt.Println(res1)
		fmt.Println(res2)
	}

}

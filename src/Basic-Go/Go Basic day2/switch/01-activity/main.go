package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday() {

	case time.Monday, time.Wednesday, time.Friday, time.Sunday:

		fmt.Println("Today is Odd")

	case time.Tuesday, time.Thursday, time.Saturday:

		fmt.Println("Today is even")

	default:

		fmt.Println("Today is a weekday")

	}

}

package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now().Weekday()

	var weekday int

	switch today {

	case weekday == 0:

		fmt.Println("Day is even")

	default:

		fmt.Println("Day is odd")

	}

}

package main

import (
	"fmt"
	"math"
)

func main() {

	var lengthOne, lengthTwo, angle float64 //user input

	fmt.Println("Please input lengthOne of triangle")

	fmt.Scanln(&lengthOne)

	fmt.Println("Please input lengthTwo of triangle")

	fmt.Scanln(&lengthTwo)

	fmt.Println("Please input angle of triangle")

	fmt.Scanln(&angle)

	resultantLen := math.Sqrt(lengthOne*lengthOne) + (lengthTwo * lengthTwo) - 2*lengthOne*lengthTwo*math.Cos(angle)

	fmt.Printf("Resultant length of triangle is %.2f", resultantLen)

}

package main

import (
	"fmt"
	"math"
)

func main() {

	var inputLength1, inputLength2, inputAngle float64

	fmt.Println("Enter length1")
	fmt.Scanln(&inputLength1)

	fmt.Println("Enter length2")
	fmt.Scanln(&inputLength2)

	fmt.Println("Enter angle")
	fmt.Scanln(&inputAngle)

	fmt.Printf("Resultant length of triangle is %.2f\n", math.Sqrt(inputLength1)+math.Sqrt(inputLength2)-2*inputLength1*inputLength2*math.Cos(inputAngle))

}

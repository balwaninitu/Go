package main

import (
	"fmt"
	"math"
)

func main() {

	var inputHeight, inputWeight float64

	fmt.Println("Enter height in meter")
	fmt.Scanln(&inputHeight)

	fmt.Println("Enter weight in kg")
	fmt.Scanln(&inputWeight)

	resultBMI := inputWeight / math.Sqrt(inputHeight)

	var res string

	switch {

	case resultBMI < 18.5:

		res = "Underweight"

	case resultBMI >= 18.5 || resultBMI <= 24.9:

		res = "Healthy Weight"

	case resultBMI >= 25 || resultBMI <= 29.9:

		res = "Overweight"

	case resultBMI >= 30 || resultBMI <= 34.9:

		res = "Obese"

	case resultBMI >= 35 || resultBMI <= 39.9:
		res = "Severely Obese"
	case resultBMI >= 40:
		{

			res = "Morbidly Obese"

		}

	}

	fmt.Printf("Your BMI is %.2f & you are %s", resultBMI, res)

}

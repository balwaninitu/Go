package main

import (
	"fmt"
)

func main() {

	var inputHeight, inputWeight float64

	fmt.Println("Enter height in meter")
	fmt.Scanln(&inputHeight)

	fmt.Println("Enter weight in kg")
	fmt.Scanln(&inputWeight)

	switch resultBMI := inputWeight / (inputHeight * inputHeight); {

	case resultBMI < 18.5:

		fmt.Println("Underweight")

	case resultBMI >= 18.5 || resultBMI < 24.9:

		fmt.Println("Healthy Weight")

	case resultBMI >= 25 || resultBMI < 29.9:

		fmt.Println("Overweight")

	case resultBMI >= 30 || resultBMI < 34.9:

		fmt.Println("Obese")

	case resultBMI >= 35 || resultBMI < 39.9:
		fmt.Println("Severely Obese")
	case resultBMI >= 40:
		{

			fmt.Println("Morbidly Obese")

		}

	}

}

package main

import "fmt"

func main() {

	var input float64

	var temp float64

	fmt.Printf("Please enter format of temperature\n(1 for Kelvin, 2 for Celsius and 3 for Fahrenheit)\n")

	fmt.Scanln(&input)

	fmt.Println("Please enter current temperature")

	fmt.Scanln(&temp)

	if input == 3 {

		temp1 := (temp - 32) * 5 / 9

		temp2 := (temp + 459.67) * 5 / 9

		fmt.Printf("%.2f Farenheit = %.2f celcius\n", temp, temp1)

		fmt.Printf("%.2f Farenheit = %.2f Kelvin\n", temp, temp2)

	} else if input == 1 {

		temp3 := (temp - 273.15)

		temp4 := (temp - 459.67) * 5 / 9

		fmt.Printf("%.2f Kelvin = %.2f celcius\n", temp, temp3)

		fmt.Printf("%.2f Kelvin = %.2f farenheit\n", temp, temp4)

	} else if input == 2 {

		temp5 := (temp + 273.15)

		temp6 := (temp*9/5 + 32)

		fmt.Printf("%.2f celcius = %.2f kelvin\n", temp, temp5)

		fmt.Printf("%.2f celcius = %.2f farenheit\n", temp, temp6)

	} else {

		fmt.Println("Invalid Input")
	}

}

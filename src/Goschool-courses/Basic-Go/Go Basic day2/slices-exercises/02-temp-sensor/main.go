package main

import "fmt"

func main() {

	temperature := [][]float64{

		{20, 21, 23, 25, 22},
		{27, 23, 25, 20, 30, 24},
		{22, 23, 24, 22},
	}

	for i := 0; i < len(temperature); i++ {
		sum := 0.0
		for j := 0; j < len(temperature[i]); j++ {
			sum += temperature[i][j]
		}

		fmt.Printf("Average of room  is %.2f\n", sum/float64(len(temperature[i])))

	}

}

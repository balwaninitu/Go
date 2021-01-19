package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		oneDollar      float64
		fiftyCentCoin  float64
		twentyCentCoin float64
		tenCentCoin    float64
		fiveCentcoin   float64
	)

	fmt.Println("Enter number of 1 dollar coins")

	fmt.Scanln(&oneDollar)

	fmt.Println("Enter number of 50 cent coins")

	fmt.Scanln(&fiftyCentCoin)

	fmt.Println("Enter number of 20 cent coins")

	fmt.Scanln(&twentyCentCoin)

	fmt.Println("Enter number of 10 cent coins")

	fmt.Scanln(&tenCentCoin)

	fmt.Println("Enter number of 5 cent coins")

	fmt.Scanln(&fiveCentcoin)

	totalAmt := (oneDollar + fiftyCentCoin*0.5 + twentyCentCoin*0.2 + tenCentCoin*0.1 + fiveCentcoin*0.05)

	fmt.Println("Number of two dollar notes: ", math.Floor(totalAmt)/2)

	fmt.Println("Remaining change: ", math.Mod(float64(totalAmt), float64(2)))

}

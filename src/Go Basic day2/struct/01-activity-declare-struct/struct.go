package main

import "fmt"

type customer struct {
	fname            string
	lname            string
	age              int
	subscriber       bool
	homeAdd          string
	phone            int
	creditAvailable  float32
	currentCartCost  float32
	currentOrderCost float32
}

func main() {

	customer1 := customer{

		fname:            "Lee",
		lname:            "Ng",
		age:              25,
		subscriber:       true,
		homeAdd:          "sg",
		phone:            345678,
		creditAvailable:  34.7,
		currentCartCost:  45.3,
		currentOrderCost: 23.89,
	}

	fmt.Println(customer1)

}

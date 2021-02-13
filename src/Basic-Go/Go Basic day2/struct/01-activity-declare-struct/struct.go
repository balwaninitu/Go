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

	customer2 := customer{

		fname:            "Annakin",
		lname:            "Skywalker",
		age:              45,
		subscriber:       true,
		homeAdd:          "death star",
		phone:            123456,
		creditAvailable:  10000.0,
		currentCartCost:  34.8,
		currentOrderCost: 23.80,
	}

	fmt.Println(customer2)
}

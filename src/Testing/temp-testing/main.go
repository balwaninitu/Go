package main

import (
	"fmt"
	"log"
)

func main() {

	// category := []string{"Household", "Food", "Drink", "Electonics"}
	// var tempCategoryNameInput int

	// for i, v := range category {
	// 	if v == "Drink" {
	// 		tempCategoryNameInput = i
	// 	}
	// }

	//fmt.Println("Index value for  is", tempCategoryNameInput)

	// func errors() error{

	// 	return errors.New("string")
	// }
	var res1, res2 string

	var name string

	var num int
	fmt.Println("Enter name")

	_, err := fmt.Scanln(&name)

	fmt.Println("Name:", name)

	if err != nil {
		res1 = "name"

	}

	fmt.Println("Enter num")

	_, err = fmt.Scanln(&num)

	fmt.Println("Num:", num)

	if err != nil {

		res2 = "num"
	}

	log.Printf(" %s not found", res1)
	log.Printf(" %s not found", res2)

	// var input string

	// fmt.Println("Enter name")
	// n, err := fmt.Scanln(&input)
	// if err != nil {
	// 	log.Println("Error:")
	// }
	// fmt.Println(n, input)
}

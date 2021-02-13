package main

import "fmt"

func appendIfMissing(category []string, newCatName string) []string {

	for i, ele := range category {
		if ele == newCatName {
			fmt.Printf("Category : %s already exist at index %d", newCatName, i)
			return category
		}
	}

	findIndex := len(category)
	category = append(category, newCatName)
	fmt.Printf("New Category : %s added at index %d", newCatName, findIndex)

	return category
}

func main() {

	category := []string{"Household", "Food", "Drink"}

	// add new name to category slice

	var newCatName string
	fmt.Println("Add New Category Name")
	fmt.Println("What is the New Category Name to add?")
	fmt.Scanln(&newCatName)
	category = appendIfMissing(category, newCatName)
	fmt.Println(category)
}

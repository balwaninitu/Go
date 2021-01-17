package main

import "fmt"

type itemInfo struct {
	itemCategory int
	quantity     int
	unitCost     float64
}

func main() {

	item1 := itemInfo{itemCategory: 0, quantity: 5, unitCost: 3}
	item2 := itemInfo{itemCategory: 0, quantity: 4, unitCost: 3}
	item3 := itemInfo{itemCategory: 0, quantity: 4, unitCost: 3}
	item4 := itemInfo{itemCategory: 1, quantity: 3, unitCost: 1}
	item5 := itemInfo{itemCategory: 1, quantity: 2, unitCost: 2}
	item6 := itemInfo{itemCategory: 2, quantity: 5, unitCost: 2}
	item7 := itemInfo{itemCategory: 2, quantity: 5, unitCost: 2}

	category := []string{"Household", "Food", "Drink"}

	fmt.Println(category)

	//fmt.Println(itemInfo)

	itemsName := map[string]itemInfo{

		"Cups":   item1,
		"Fork":   item2,
		"Plates": item3,
		"Cake":   item4,
		"Bread":  item5,
		"Coke":   item6,
		"Sprite": item7,
	}

	// Print current data fields

	var printCurrentData string

	fmt.Println("Print Current Data.")

	fmt.Scanln(&printCurrentData)

	for item, info := range itemsName {

		if printCurrentData == item {

			fmt.Printf("%s - %v\n", item, info)

			return

		}

	}

	fmt.Println("No data found!")

	// Add New Category Name

	var addNewCategory string

	var newCategory []string

	fmt.Println("Add New Category Name")
	fmt.Println("What is the New category Name to add?")

	fmt.Scanln(&addNewCategory)

	//var v string

	for i, v := range category {

		addNewCategory = v

		if addNewCategory == v {

			fmt.Printf("Category: %s already exixst at index %d", addNewCategory, i)

			return
		} else if addNewCategory != v {

			fmt.Printf("New category: %s added at index %d", addNewCategory, i)

			return
		} else {

			fmt.Println("No Input Found!")

		}

		{

			fmt.Println("No Input Found!")

			return

		}
	}

	// newCategory = append(category, addNewCategory)

	fmt.Println(newCategory)

	// fmt.Printf("New category: %s added at index %d", addNewCategory, i)
	// // } else if addNewCategory == v {

	// 	fmt.Printf("Category: %s already exixst at index %d", addNewCategory, i)
	// } else {

	// 	fmt.Println("No Input Found!")
}

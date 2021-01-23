package main

import (
	"fmt"
	"strings"
)

/*
1) Delcare constants for
			1) Shoppin List Menu
			2) Generating Report
			3) Each Category Cost

			Const was used as values for  display menus will not change
*/

const (

	// Shopping List title
	title = "Shopping List Application"

	// Shopping List Menu
	shoppingListMenu = `  
1. View entire shopping list.
2. Generate Shopping List Report
3. Add Items information.
4. Modify Existing Items.
5. Delete Items from shopping List.
6. Print Current Data Fields.
7. Add New Category Name.
Select your choice(pick desired number):`

	// Report Generation
	report = `  
Generate Report
1. Total Cost of each category.
2. List of item by category.
3. Main Menu.
Choose your choice (pick desired number):`
)

type itemInfo struct {
	itemCategory int
	quantity     int
	unitCost     float64
}

type itemCostSlice []itemInfo

func main() {

	item1 := itemInfo{itemCategory: 0, quantity: 5, unitCost: 3}
	item2 := itemInfo{itemCategory: 0, quantity: 4, unitCost: 3}
	item3 := itemInfo{itemCategory: 0, quantity: 4, unitCost: 3}
	item4 := itemInfo{itemCategory: 1, quantity: 3, unitCost: 1}
	item5 := itemInfo{itemCategory: 1, quantity: 2, unitCost: 2}
	item6 := itemInfo{itemCategory: 2, quantity: 5, unitCost: 2}
	item7 := itemInfo{itemCategory: 2, quantity: 5, unitCost: 2}

	eachCost := itemCostSlice{

		itemInfo{0, 5, 3},
		itemInfo{0, 4, 3},
		itemInfo{0, 4, 3},
		itemInfo{1, 3, 1},
		itemInfo{1, 2, 2},
		itemInfo{2, 5, 2},
		itemInfo{2, 5, 2},
	}

	category := []string{"Household", "Food", "Drink"}

	itemsName := map[string]itemInfo{

		"Cups":   item1,
		"Fork":   item2,
		"Plates": item3,
		"Cake":   item4,
		"Bread":  item5,
		"Coke":   item6,
		"Sprite": item7,
	}
	// List menu in infinite loop

	for {
		userShoppingListMenuInput := displayShoppingListMenu()

		if userShoppingListMenuInput < 1 || userShoppingListMenuInput > 7 {
			fmt.Println("\n*******Enter valid choice between 1 to 7 !!!")
		} else {
			fmt.Print("User action function called")
		}
	}

	fmt.Print(category, itemsName, eachCost)
}

// function to display List menu to user
func displayShoppingListMenu() int {
	fmt.Println()
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", 25))
	fmt.Print(strings.TrimSpace(shoppingListMenu))
	var userShoppingListMenuInput int
	if _, err := fmt.Scan(&userShoppingListMenuInput); err != nil {
		fmt.Println("Error:", err)
		userShoppingListMenuInput = 0
		return userShoppingListMenuInput
	}

	return userShoppingListMenuInput
}

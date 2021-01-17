package main

import (
	"fmt"
	"strings"
)

const (
	title = "Shopping List Application"

	shoppingListMenu = `  
1. View entire shopping list.
2. Generate Shopping List Report
3. Add Items.
4. Modify Items.
5. Delete Items.
Select your choice : `

	report = `  
Generate Report
1. Total Cost of each category.bufio
2. List of item by category.
3. Main Menu.

Choose your report : `

	eachCatCost = `  
Total cost by Category.
Household cost : 39
Food cost : 7
Drink cost : 20`
)

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

	itemsName := map[string]itemInfo{

		"Cups":   item1,
		"Fork":   item2,
		"Plates": item3,
		"Cake":   item4,
		"Bread":  item5,
		"Coke":   item6,
		"Sprite": item7,
	}

	for {
		userShoppingListMenuInput := displayShoppingListMenu()
		userChoiceAction(userShoppingListMenuInput, itemsName, category)
	}

}

func displayShoppingListMenu() int {
	fmt.Println()
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", 25))
	fmt.Print(strings.TrimSpace(shoppingListMenu))
	var userShoppingListMenuInput int
	fmt.Scanln(&userShoppingListMenuInput)
	return userShoppingListMenuInput
}

func userChoiceAction(userShoppingListMenuInput int, itemsName map[string]itemInfo, category []string) {
	switch userShoppingListMenuInput {
	case 1:
		var item string
		var info itemInfo
		fmt.Println("\nShopping List Contents:")
		for item, info = range itemsName {
			fmt.Printf("Category: %s - Item: %s  Quantity : %d Unit Cost: %.0f\n", category[info.itemCategory], item, info.quantity, info.unitCost)
		}
	}
	return
}

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
		userChoiceAction(userShoppingListMenuInput, itemsName, category, eachCost)
	}

}

// function to display List menu to user
func displayShoppingListMenu() int {
	fmt.Println()
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", 25))
	fmt.Print(strings.TrimSpace(shoppingListMenu))
	var userShoppingListMenuInput int
	fmt.Scanln(&userShoppingListMenuInput)

	return userShoppingListMenuInput
}

// displya to user based on choice selected
func userChoiceAction(userShoppingListMenuInput int, itemsName map[string]itemInfo, category []string, eachCost itemCostSlice) {
	switch userShoppingListMenuInput {
	case 1: // Disply Entire shopping List
		var item string
		var info itemInfo
		fmt.Println("\nShopping List Contents:")

		for item, info = range itemsName {
			fmt.Printf("Category: %s - Item: %s  Quantity : %d Unit Cost: %.0f\n", category[info.itemCategory], item, info.quantity, info.unitCost)
		}

	case 2: //Display shopping list report
		for {
			fmt.Println()
			fmt.Print(strings.TrimSpace(report))
			var reportInput int
			fmt.Scanln(&reportInput)

			switch reportInput {
			case 1: //Display total cost of each category
				fmt.Println()
				fmt.Println(strings.TrimSpace(eachCatCost))

				fmt.Println(eachCost)
				m := make(map[int]int)
				var eachCosts itemInfo
				for _, eachCosts = range eachCost {

					m[eachCosts.itemCategory] += (eachCosts.quantity * int(eachCosts.unitCost))

				}

				fmt.Print(m)
				fmt.Println(category)

				for key, element := range m {
					fmt.Println("Category:", category[key], "=>", "Element:", element)
				}

			case 2: // Display list of item by category
				fmt.Println("\nList By category.")
				for item, info := range itemsName {
					fmt.Printf("Category: %s - Item: %s  Quantity : %d Unit Cost: %.0f\n", category[info.itemCategory], item, info.quantity, info.unitCost)

				}
				fmt.Println()
			case 3:
				return
			}
		}

	case 3: // Add items to list
		var itemNameInput string
		var categoryNameInput string
		var unitInput int
		var costInput float64
		var tempCategoryNameInput int
		fmt.Println("What is the name of your item?")
		fmt.Scanln(&itemNameInput)
		fmt.Println("What category does it belong to?")
		fmt.Scanln(&categoryNameInput)

		fmt.Println("How many units are there?")
		fmt.Scanln(&unitInput)
		fmt.Println("How much does it cost per unit")
		fmt.Scanln(&costInput)

		//Find Index of Category Value entered by user (String -> Int value). Int value is required as struct category information in int
		for i, v := range category {
			if v == categoryNameInput {
				tempCategoryNameInput = i
			}
		}

		tempItem := itemInfo{itemCategory: tempCategoryNameInput, quantity: unitInput, unitCost: costInput}

		itemsName[itemNameInput] = tempItem

		//fmt.Println("New item added in the list")
		//fmt.Println(itemsName)

	case 4: //Modify items in the list

		var modifyItem, modifyName, modifyCategory string
		var modifyQuantity, tempModifyCategory int
		var modifyUnitCost float64

		fmt.Println("Modify Items.")
		fmt.Println("Which item would you wish to modify?")
		fmt.Scanln(&modifyItem)

		var info itemInfo = itemsName[modifyItem]
		fmt.Printf("Current item is %s - Category is %s  - Quantity is %d - Unit Cost %0.f\n", modifyItem, category[info.itemCategory], info.quantity, info.unitCost)

		fmt.Println("Enter new name.Enter for no change.")
		if _, err := fmt.Scanln(&modifyName); err != nil {
			defer fmt.Println("No changes to item name made")
		}

		fmt.Println("Enter new Category.Enter for no change.")
		if _, err := fmt.Scanln(&modifyCategory); err != nil {
			defer fmt.Println("No changes to category name made")
		}
		fmt.Println("Enter new Quantity.Enter for no change.")
		if _, err := fmt.Scanln(&modifyQuantity); err != nil {
			defer fmt.Println("No changes to quantity name made")
		}

		fmt.Println("Enter new Unit cost.Enter for no change.")
		if _, err := fmt.Scanln(&modifyUnitCost); err != nil {
			defer fmt.Println("No changes to unit cost name made")
		}

		for i, v := range category {
			if v == modifyCategory {
				tempModifyCategory = i
			}
		}

		//add new key to map
		itemsName[modifyName] = itemInfo{itemCategory: tempModifyCategory, quantity: modifyQuantity, unitCost: modifyUnitCost}
		var infoAfterUpdate itemInfo = itemsName[modifyName]
		fmt.Printf("Category: %s - Item: %s  Quantity : %d Unit Cost: %.0f\n", category[infoAfterUpdate.itemCategory], modifyName, infoAfterUpdate.quantity, infoAfterUpdate.unitCost)
		delete(itemsName, modifyItem)
		fmt.Println(itemsName)

	case 5: // Delete items from list
		var deleteItemName string
		fmt.Println("Delete Item.")
		fmt.Println("Enter item name to delete:")
		fmt.Scanln((&deleteItemName))

		var found bool
		_, found = itemsName[strings.Title(deleteItemName)]

		if found == true {

			fmt.Printf("Deleted %s\n", deleteItemName)

		} else {

			fmt.Println("Item not found. Nothing to delete!")
		}
	}
}

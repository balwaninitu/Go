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
Choose your (pick desired number):`
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
		"Coke":   item6,
		"Plates": item3,
		"Cake":   item4,
		"Fork":   item2,
		"Bread":  item5,
		"Sprite": item7,
	}
	// List menu in infinite loop

	for {
		userShoppingListMenuInput := displayShoppingListMenu()

		if userShoppingListMenuInput < 1 || userShoppingListMenuInput > 7 {
			fmt.Println("\n*******Enter valid choice between 1 to 7 !!!")
		} else {
			userChoiceAction(userShoppingListMenuInput, itemsName, category, eachCost)
		}
	}
}

func appendIfMissing(category []string, newCatName string) []string {

	for i, ele := range category {
		if ele == newCatName {
			fmt.Println("Add New Category Name")
			fmt.Printf("Category : %s already exist at index %d", newCatName, i)
			return category
		}
	}
	category = append(category, newCatName)
	findIndex := len(category) - 1
	fmt.Println("Add New Category Name")
	fmt.Printf("New Category : %s added at index %d", newCatName, findIndex)
	return category
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

// switch case to display to user based on choice selected
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
			case 1:

				/*Display total cost of each category: Created slice of struct to group as per itemcategory
				ans then range over it to get total cost by category. */

				fmt.Println()
				m := make(map[int]int)
				var eachCosts itemInfo
				for _, eachCosts = range eachCost {

					m[eachCosts.itemCategory] += (eachCosts.quantity * int(eachCosts.unitCost))

				}

				for key, element := range m {
					fmt.Printf("%s cost : %d\n", category[key], element)
				}

			case 2: // Display list of item by category
				fmt.Println()
				fmt.Println("List by Category")
				for item, info := range itemsName {
					fmt.Printf("Category: %s - Item: %s  Quantity : %d Unit Cost: %.0f\n", category[info.itemCategory], item, info.quantity, info.unitCost)

				}

				fmt.Println()
			case 3:

				return

			}
		}

	case 3: // Add items to list

		var newItemNameInput string
		var newCategoryNameInput string
		var newUnitInput int
		var newCostInput float64
		var tempCategoryNameInput int
		fmt.Println("What is the name of your item?")
		fmt.Scanln(&newItemNameInput)

		fmt.Println("What category does it belong to?")
		fmt.Scanln(&newCategoryNameInput)

		fmt.Println("How many units are there?")
		fmt.Scanln(&newUnitInput)

		fmt.Println("How much does it cost per unit")
		fmt.Scanln(&newCostInput)

		//Find Index of Category Value entered by user (String -> Int value). Int value is required a struct category information in int
		tempCategoryNameInput = -1
		for i, v := range category {
			if strings.Title(v) == strings.Title(newCategoryNameInput) {
				tempCategoryNameInput = i
			}
		}

		if tempCategoryNameInput >= 0 {
			tempItem := itemInfo{itemCategory: tempCategoryNameInput, quantity: newUnitInput, unitCost: newCostInput}
			itemsName[strings.Title(newItemNameInput)] = tempItem
			fmt.Println("New item added in the list")
			fmt.Println(itemsName)
		} else {
			fmt.Printf("Category %s does not exist. Pleae add from given category %s!", newCategoryNameInput, category)
		}

	case 4: //Modify items in the list

		var modifyItem, modifyName, modifyCategory string
		var modifyQuantity, tempModifyCategory int
		var modifyUnitCost float64

		fmt.Println("Modify Items.")
		fmt.Println("Which item would you wish to modify?")
		fmt.Scanln(&modifyItem)

		var info itemInfo
		var found bool
		info, found = itemsName[strings.Title(modifyItem)]

		if found == true {

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
				fmt.Println("No changes to unit cost name made")
			}

			for i, v := range category {
				if v == modifyCategory {
					tempModifyCategory = i
					return
				}
			}

			//add new key to map
			itemsName[strings.Title(modifyName)] = itemInfo{itemCategory: tempModifyCategory, quantity: modifyQuantity, unitCost: modifyUnitCost}
			var infoAfterUpdate itemInfo = itemsName[modifyName]

			fmt.Printf("Category: %s - Item: %s  Quantity : %d Unit Cost: %.0f\n", category[infoAfterUpdate.itemCategory], modifyName, infoAfterUpdate.quantity, infoAfterUpdate.unitCost)
			delete(itemsName, strings.Title(modifyItem))
			fmt.Println(itemsName)

		} else {
			fmt.Printf("Item not found. Nothing to modify!!!")
		}

	case 5: // Delete items from list
		var deleteItemName string
		fmt.Println("Delete Item.")
		fmt.Println("Enter item name to delete:")
		fmt.Scanln(&deleteItemName)

		for item := range itemsName {
			switch strings.Title(deleteItemName) == item {
			case true:
				delete(itemsName, strings.Title(deleteItemName))
				fmt.Printf("Deleted %s\n", deleteItemName)
				fmt.Println(itemsName)
				return
			}
		}
		fmt.Println("Item not found. Nothing to delete!")

	case 6:

		var printData string
		fmt.Println("Print current data")
		fmt.Println("Which item data you wish to see?")
		fmt.Scanln(&printData)

		var infoprint itemInfo = itemsName[printData]
		fmt.Printf("%s - {%d %d %0.f}\n", printData, infoprint.itemCategory, infoprint.quantity, infoprint.unitCost)

		for item := range itemsName {

			switch strings.Title(printData) == item {

			case true:

				var infoprint itemInfo = itemsName[strings.Title(printData)]
				fmt.Printf("%s - {%d %d %0.f}\n", strings.Title(printData), infoprint.itemCategory, infoprint.quantity, infoprint.unitCost)
				return
			}
		}

		fmt.Println("Print current data")
		fmt.Println("No data found!")

	case 7:
		var newCatName string
		fmt.Println("Add New Category Name")
		fmt.Println("What is the New Category Name to add?")
		fmt.Scanln(&newCatName)
		category = appendIfMissing(category, strings.Title(newCatName))
		fmt.Println(category)

	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
Select your choice: `

	report = `  
Generate Report
1. Total Cost of each category.bufio
2. List of item by category.
3. Main Menu.

Choose your report: `

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

	// The shopping list menu on runtime prompt and allow the user to make a selection.

	fmt.Println(title)
	fmt.Println(strings.Repeat("=", 25))
	fmt.Println(strings.TrimSpace(shoppingListMenu))

	var listScan = bufio.NewScanner(os.Stdin)

	var item string

	var info itemInfo

	//  View Entire Shopping List

loop1:
	for listScan.Scan() {

		scanInputList, _ := strconv.Atoi(listScan.Text())

		if scanInputList == 1 {

			for item, info = range itemsName {

				fmt.Printf("Category: %s - Item: %s  Quantity : %d Unit Cost: %.0f\n", category[info.itemCategory], item, info.quantity, info.unitCost)

			}

			goto loop1

		} else if scanInputList == 3 { // Add Items to the List

			var word string
			var num1 int
			var num2 float64

			fmt.Println("What is the name of your item?")
			fmt.Scanln(&word)

			fmt.Println("What category does it belong to?")
			fmt.Scanln((&word))

			fmt.Println("How many units are there?")
			fmt.Scanln((&num1))

			fmt.Println("How much does it cost per unit?")
			fmt.Scanln((&num2))

		} else if scanInputList == 4 { // Modify items in the List

			var inputItem, newName, newCategory string

			var newQuantity int
			var newUnitCost float64

			fmt.Println("Modify Items.")
			fmt.Println("Which item would you wish to modify?")

			fmt.Scanln(&inputItem)

			for item, info = range itemsName {

				if inputItem == item {

					fmt.Printf("Current item is %s - Category is %s  - Quantity is %d - Unit Cost %0.f\n", inputItem, category[info.itemCategory], info.quantity, info.unitCost)

					fmt.Println("Enter new name.Enter for no change.")

					fmt.Scanln(&newName)

					fmt.Println("Enter new Category.Enter for no change.")

					fmt.Scanln(&newCategory)

					fmt.Println("Enter new Quantity.Enter for no change.")

					fmt.Scanln(&newQuantity)

					fmt.Println("Enter new Unit cost.Enter for no change.")

					fmt.Scanln(&newUnitCost)

					fmt.Println("No changes to category made.")
					fmt.Println("No changes to quantity made.")
					fmt.Println("No changes to unit cost made.")
					fmt.Println("No changes to item name made.")

				}

			}
		} else if scanInputList == 5 { // Delete Items from the List

			var deleteItemName string

			fmt.Println("Delete Item.")
			fmt.Println("Enter item name to delete:")
			fmt.Scanln(&deleteItemName)

			for item, info = range itemsName {

				if deleteItemName == item {

					fmt.Printf("Deleted %s\n", deleteItemName)

					return

				}

			}

			fmt.Println("Item not found. Nothing to delete!")

		} else if scanInputList == 2 { // Generate Shopping List Report

			fmt.Println(strings.TrimSpace(report))

			reportScan := bufio.NewScanner(os.Stdin)

			goto loop2

		loop2:
			for reportScan.Scan() {

				scanInputReport, _ := strconv.Atoi(reportScan.Text())

				if scanInputReport == 1 { // Report of Total cost of each category

					fmt.Println(strings.TrimSpace(eachCatCost))

				} else if scanInputReport == 2 { // List of item by category

					for item, info := range itemsName {

						fmt.Printf("Category: %s - Item: %s  Quantity : %d Unit Cost: %.0f\n", category[info.itemCategory], item, info.quantity, info.unitCost)

					}
				} else if scanInputReport == 3 { // Main Menu

					fmt.Println(title)
					fmt.Println(strings.Repeat("=", 25))
					fmt.Println(strings.TrimSpace(shoppingListMenu))

					listScan = bufio.NewScanner(os.Stdin)

					goto loop1

				}

			}

		}
	}

}

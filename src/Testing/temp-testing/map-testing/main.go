package main

import (
	"fmt"
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

	var modifyItem, modifyName, modifyCategory string
	var modifyQuantity, tempModifyCategory int
	var modifyUnitCost float64

	fmt.Println("Modify Items.")
	fmt.Println("Which item would you wish to modify?")
	fmt.Scanln(&modifyItem)

	var info itemInfo = itemsName[modifyItem]
	fmt.Printf("Category: %s - Item: %s  Quantity : %d Unit Cost: %.0f\n", category[info.itemCategory], modifyItem, info.quantity, info.unitCost)

	fmt.Println("Enter new name.Enter for no change.")
	fmt.Scanln(&modifyName)

	fmt.Println("Enter new Category.Enter for no change.")
	fmt.Scanln(&modifyCategory)

	fmt.Println("Enter new Quantity.Enter for no change.")
	fmt.Scanln(&modifyQuantity)

	fmt.Println("Enter new Unit cost.Enter for no change.")
	fmt.Scanln(&modifyUnitCost)

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

	// Print Current Data Files :

	var printData string

	fmt.Println("Print current data")
	fmt.Println("Which item data you wish to see?")
	fmt.Scanln(&printData)

	var infoprint itemInfo = itemsName[printData]

	fmt.Printf("%s - {%d %d %0.f}\n", printData, infoprint.itemCategory, infoprint.quantity, infoprint.unitCost)

}

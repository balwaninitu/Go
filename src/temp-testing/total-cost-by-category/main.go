package main

import "fmt"

type itemInfo struct {
	itemCategory int
	quantity     int
	unitCost     float64
}

type itemCostSlice []itemInfo

func main() {

	category := []string{"Household", "Food", "Drink"}

	// item1 := itemInfo{itemCategory: 0, quantity: 5, unitCost: 3}
	// item2 := itemInfo{itemCategory: 0, quantity: 4, unitCost: 3}
	// item3 := itemInfo{itemCategory: 0, quantity: 4, unitCost: 3}
	// item4 := itemInfo{itemCategory: 1, quantity: 3, unitCost: 1}
	// item5 := itemInfo{itemCategory: 1, quantity: 2, unitCost: 2}
	// item6 := itemInfo{itemCategory: 2, quantity: 5, unitCost: 2}
	// item7 := itemInfo{itemCategory: 2, quantity: 5, unitCost: 2}

	eachCost := itemCostSlice{

		itemInfo{0, 5, 3},
		itemInfo{0, 4, 3},
		itemInfo{0, 4, 3},
		itemInfo{1, 3, 1},
		itemInfo{1, 2, 2},
		itemInfo{2, 5, 2},
		itemInfo{2, 5, 2},
	}

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
}

// item1.itemCategory := itemInfo.quantity*iteInfo.unitCost
// item2.itemCategory := itemInfo.quantity*iteInfo.unitCost
// item3.itemCategory := itemInfo.quantity*iteInfo.unitCost
// item4.itemCategory := itemInfo.quantity*iteInfo.unitCost
// item5.itemCategory := itemInfo.quantity*iteInfo.unitCost
// item6.itemCategory := itemInfo.quantity*iteInfo.unitCost
// item7.itemCategory := itemInfo.quantity*iteInfo.unitCost

// totoatCostof1 = item1.itemCategory + item2.itemCategory + item3.itemCategory

// totalCostof2 = item4.itemCategory + item5.itemCategory
// totalCostof3 = item6.itemCategory + item7.itemCategory

//category[itemInfo.itemCategory]= [eachCosts.itemCategory]

//fmt.Println("Cat", eachCosts.itemCategory)

//fmt.Printf("Category %s, %d\n", (category[itemInfo.itemCategory]), (eachCosts.quantity * (eachCosts.unitCost)))

// var k int
// var v int

// for k, v = range m {

// 	fmt.Println("Map key & value", k, v)
// }

// fmt.Println("Map key 1", m[0])

// var i int

// for i = range category {

// 	i = m[k]

// //	fmt.Println("Category i", i)
// }

// 	fmt.Println(m)

// 	fmt.Println()

// 	itemsName := map[string]itemInfo{

// 		"Cups":   item1,
// 		"Fork":   item2,
// 		"Plates": item3,
// 		"Cake":   item4,
// 		"Bread":  item5,
// 		"Coke":   item6,
// 		"Sprite": item7,
// 	}

// 	fmt.Println()

// 	fmt.Println(itemsName)

// 	fmt.Println(category)

// }

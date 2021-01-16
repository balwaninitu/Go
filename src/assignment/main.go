package main

import "fmt"

const (
	title = "Shopping List Application"

	shoppingListMenu = `  
1. View entire shopping list.
2. Generate Shopping List Report
3. Add Items.
4. Modify Items.
5. Delete Items.
Select your choice:`
)

type itemInfo struct {
	itemCategory int
	quantity     int
	unitCost     float64
}

//var m map[string]itemInfo

func main() {

	category := []string{"Household", "Food", "Drinks"}

	fmt.Println(category[0])

	var i int

	for i = range category {

		//fmt.Println(itemCategory)
		fmt.Println(i)

	}

	// i :=itemInfo{}

	// i.itemCategory = category[i]

	var m = map[string]itemInfo{

		"Cups":   {itemCategory: 0, quantity: 5, unitCost: 3},
		"Fork":   {itemCategory: 0, quantity: 4, unitCost: 3},
		"Plates": {itemCategory: 0, quantity: 4, unitCost: 3},
		"Cake":   {itemCategory: 1, quantity: 3, unitCost: 1},
		"Bread":  {itemCategory: 1, quantity: 2, unitCost: 2},
		"Coke":   {itemCategory: 2, quantity: 5, unitCost: 2},
		"Sprite": {itemCategory: 2, quantity: 5, unitCost: 2},
	}

	fmt.Println(m)
	fmt.Println(category[i.itemCategory])

}

// }

// categorys := []category{
// 	{

// 		categoryName: "Household",

// 		item: item{itemName: "Cups", quantity: 5, unitCost: 3},
// 		{itemName: "Fork", quantity: 4, unitCost: 3},
// 		//item{itemName:"Plates", quantity: 4, unitCost: 3},

// 	},
// }

// {

// 	item: item{"Cake", quantity: 3, unitCost: 1},{"Bread", quantity: 2, unitCost: 2},
// 	categoryName: "Food",
// },
// {

// 	item: item{"Coke", quantity: 5, unitCost: 2},{"Sprite", quantity: 5, unitCost: 2},
// 	categoryName:"Drinks",
// },

// }

//in := bufio.NewScanner(os.Stdin)

//  for {

//  	fmt.Println(title)
//  	fmt.Println(strings.Repeat("=", 25))
//  	fmt.Println(strings.TrimSpace(shoppingListMenu))

//  	if !in.Scan() {
//  		break
//  	}
//  }

//  cmd := strings.Fields(in.Text())

//  if len(cmd) == 1 {

//  	continue
//  }

// fmt.Println(categorys)

//}

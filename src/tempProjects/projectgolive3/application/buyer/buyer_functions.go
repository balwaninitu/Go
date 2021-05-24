package buyer

/*
import "fmt"

//check if sellername available
func checkSellerName() string {
	var sellername string
	sellerSlice := MakeSellerNameSlice()
	for _, v := range sellerSlice {
		if v == sellername {
			fmt.Println("seller added to list")
			return sellername
		}
	}
	fmt.Println("seller not avialable")
	return ""
}

func checkItemName() string {
	var itemname string
	var sellername string
	mapItemName := MapItemNameDB()

	if mapItemName[sellername] == itemname {
		fmt.Println("Item added to cart")
		return itemname

	}
	fmt.Println("Item not available with selle")
	return ""
}

func checkQty() int {
	var quantity int
	var adqty int
	mapQty := MapItemQtyDB()

	for _, value := range mapQty {
		if value == quantity {
			fmt.Println("quantity added")
			return quantity
		}
		if value < quantity {
			fmt.Println("Not enough quantity")
			return 0
		}
		if value > quantity {
			adqty = value - quantity
			return adqty
		}
	}
	return quantity
}

func checkCost() float64 {
	var itemname string
	var cost float64
	mapCost := MapItemCostDB()

	if mapCost[itemname] == cost {
		fmt.Printf("Cost of Item is %0.2f", cost)
		return cost

	}
	fmt.Println("Item not available with selle")
	return 0
}
*/

package buyer

/*
type ItemDetailsToBuy struct {
	name     string
	quantity int
	//sellername string
	next *ItemDetailsToBuy
}

type DetailsList struct {
	start  *ItemDetailsToBuy
	length int
}

func (d *DetailsList) addItemDetails(name string, quantity int) {
	newDetails := &ItemDetailsToBuy{
		name:     name,
		quantity: quantity,
		//sellername: sellername,
		next: nil,
	}
	if d.start == nil {
		d.start = newDetails
	} else {
		currentDetails := d.start
		if currentDetails.name == name {
			d.start.quantity = d.start.quantity + quantity
			//d.start.sellername = d.start.sellername + sellername
			return
		}
		for currentDetails.next != nil {
			currentDetails = currentDetails.next
			if currentDetails.name == name {
				currentDetails.quantity = currentDetails.quantity + quantity
				//currentDetails.sellername = currentDetails.sellername + sellername
				return
			}
		}
		currentDetails.next = newDetails
	}
	d.length++
	return
}

func (d *DetailsList) AddedItemQty() (map[string]int, error) {
	AddedItemQty := map[string]int{}
	currentDetails := d.start
	if currentDetails == nil {
		return nil, nil
	}
	//item := r.FormValue
	if currentDetails.name == "apple" {
		AddedItemQty["apple"] = currentDetails.quantity
	} else if currentDetails.name == "orange" {
		AddedItemQty["orange"] = currentDetails.quantity
	} else if currentDetails.name == "banana" {
		AddedItemQty["banana"] = currentDetails.quantity
	}
	for currentDetails.next != nil {
		currentDetails = currentDetails.next
		if currentDetails.name == "apple" {
			AddedItemQty["apple"] = currentDetails.quantity
		} else if currentDetails.name == "orange" {
			AddedItemQty["orange"] = currentDetails.quantity
		} else if currentDetails.name == "banana" {
			AddedItemQty["banana"] = currentDetails.quantity
		}
	}
	return AddedItemQty, nil
}
func (d *DetailsList) DeleteAllItems() {
	d.start = nil
	d.length = 0
}

func (d *DetailsList) AddAllItems() map[string]int {
	CartItems := map[string]int{}
	currentDetails := d.start
	if currentDetails == nil {
		CartItems["emptycart"] = 1
		return CartItems
	}
	CartItems[currentDetails.name] = currentDetails.quantity
	fmt.Printf("Name: %s, Quantity: %v\n", currentDetails.name, currentDetails.quantity)
	for currentDetails.next != nil {
		currentDetails = currentDetails.next
		fmt.Printf("Name: %s, Quantity: %v\n", currentDetails.name, currentDetails.quantity)
		CartItems[currentDetails.name] = currentDetails.quantity
	}
	return CartItems
}

func (d *DetailsList) ItemsCost() map[string]int {
	currentDetails := d.start
	var appleCost int
	var orangeCost int
	var bananaCost int
	ItemCost := map[string]int{}
	if currentDetails == nil {
		return nil
	}
	if currentDetails.name == "apple" {
		appleCost = currentDetails.quantity * applePrice
		ItemCost["apple"] = appleCost
	} else if currentDetails.name == "orange" {
		orangeCost = currentDetails.quantity * orangePrice
		ItemCost["orange"] = orangeCost
	} else if currentDetails.name == "banana" {
		bananaCost = currentDetails.quantity * bananaPrice
		ItemCost["banana"] = bananaCost
	}

	for currentDetails.next != nil {
		currentDetails = currentDetails.next
		if currentDetails.name == "apple" {
			appleCost = currentDetails.quantity * applePrice
			ItemCost["apple"] = appleCost
		} else if currentDetails.name == "orange" {
			orangeCost = currentDetails.quantity * orangePrice
			ItemCost["orange"] = orangeCost
		} else if currentDetails.name == "banana" {
			bananaCost = currentDetails.quantity * bananaPrice
			ItemCost["banana"] = bananaCost
		}
	}
	ItemCost["totalcost"] = appleCost + orangeCost + bananaCost
	fmt.Println(ItemCost)
	return ItemCost
}*/

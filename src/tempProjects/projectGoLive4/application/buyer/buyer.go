//Handles buyer functions - talks to templates and apiclient
//gets login info from server.go
package buyer

import (
	"fmt"
	"net/http"
	"projectGoLive/application/apiclient"
	"projectGoLive/application/config"
	"projectGoLive/application/email"
	"projectGoLive/application/server"
	"strconv"
	"strings"
)

type buyerStruct struct {
	Buyername   string
	Operation   string
	Mainmessage []string
	Selleritems []apiclient.ItemsDetails
}

type buyerCartStruct struct {
	Buyername   string
	Operation   string
	Mainmessage []string
	Cartitems   []apiclient.ItemsDetails
}

var buyerCartll CartLinkedList

func init() {
	buyerCartll = CartLinkedList{Head: nil, Size: 0}
}

func BuyerHandler(w http.ResponseWriter, req *http.Request) {
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	user := server.GetUser(w, req)

	buyerMessage := buyerStruct{
		Buyername:   user.Username,
		Operation:   "view",
		Mainmessage: nil,
		Selleritems: nil,
	}

	buyerMessage.Mainmessage = append(buyerMessage.Mainmessage, "Welcome to Peel Rescue! ")

	allSellerItems, ok := apiclient.GetItem("", "", user.IsBuyer)
	if !ok {
		config.Error.Println("Unable to connect to Database!")
		return
	}

	allSellerItems = removeCartItems(allSellerItems)

	buyerMessage.Selleritems = allSellerItems
	buyerMessage.Mainmessage = append(buyerMessage.Mainmessage, "List of all items: ")

	if req.Method == http.MethodPost {
		addthisitem := req.FormValue("product_id")
		newquantity := req.FormValue("newquantity")
		fmt.Println("Quantity from form :", newquantity)

		if addthisitem != "" {
			// Add this item to linked list
			item := ConvStringtoSlice(addthisitem)
			fmt.Println(item)
			item.Quantity = ConvertToInt(newquantity)
			fmt.Println(item)
			buyerCartll.AddNode(item)
			// redirect to main index
			http.Redirect(w, req, "/buyer/buyercart", http.StatusSeeOther)
			return
		}
	}
	//display available items on browser
	config.TPL.ExecuteTemplate(w, "buyertemplate.gohtml", buyerMessage)
}

func LookForItemHandler(w http.ResponseWriter, req *http.Request) {
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	user := server.GetUser(w, req)
	var oneItemAllSellers []apiclient.ItemsDetails

	buyerMessage := buyerStruct{
		Buyername:   user.Username,
		Operation:   "finditem",
		Mainmessage: nil,
		Selleritems: nil,
	}
	buyerMessage.Mainmessage = append(buyerMessage.Mainmessage, "Please choose an item to search: ")

	// process form submission , when buyer clicks submit
	if req.Method == http.MethodPost {
		chosenitem := req.FormValue("fruit")
		//checkaddtocart(w, req)
		addthisitem := req.FormValue("product_id")

		//fmt.Println("added item", addthisitem)
		//fmt.Println("search item", chosenitem)

		if chosenitem != "" {
			allSellerItems, ok := apiclient.GetItem("", "", user.IsBuyer)
			if !ok {
				config.Error.Println("Unable to connect to Database!")
				return
			}
			allSellerItems = removeCartItems(allSellerItems)

			for _, item := range allSellerItems {
				if item.Item == chosenitem {
					oneItemAllSellers = append(oneItemAllSellers, item)
				}
			}
			buyerMessage.Selleritems = oneItemAllSellers
		} else if addthisitem != "" {
			// Add this item to linked list
			item := ConvStringtoSlice(addthisitem)
			fmt.Println(item)
			newquantity := req.FormValue("newquantity")
			fmt.Println("Quantity from form :", newquantity)
			item.Quantity = ConvertToInt(newquantity)
			fmt.Println(item)
			buyerCartll.AddNode(item)
			// redirect to main index
			http.Redirect(w, req, "/buyer/buyercart", http.StatusSeeOther)
			return
		}
	}
	config.TPL.ExecuteTemplate(w, "buyertemplate.gohtml", buyerMessage)
}

func ConvStringtoSlice(iteminput string) apiclient.ItemsDetails {
	iteminput = strings.Replace(iteminput, "{", "", -1)
	iteminput = strings.Replace(iteminput, "}", "", -1)
	//itemstr := []string{addthisitem}
	itemslice := strings.Split(iteminput, " ")
	item := apiclient.ItemsDetails{}
	item.Item = itemslice[0]
	item.Quantity = ConvertToInt(itemslice[1])
	item.Cost = ConvertToFloat(itemslice[2])
	item.Username = itemslice[3]
	return item
}

func CartHandler(w http.ResponseWriter, req *http.Request) {
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	user := server.GetUser(w, req)

	buyerCart := buyerCartStruct{
		Buyername:   user.Username,
		Operation:   "view",
		Mainmessage: nil,
		Cartitems:   nil,
	}

	buyerCart.Mainmessage = append(buyerCart.Mainmessage, "Shopping Cart: ")

	_, allitems := buyerCartll.GetAllItems()

	buyerCart.Cartitems = allitems

	// process form submission , when buyer clicks submit
	if req.Method == http.MethodPost {
		addmore := req.FormValue("add_more")
		reset := req.FormValue("reset")
		checkout := req.FormValue("checkout")

		if addmore != "" {
			http.Redirect(w, req, "/buyer", http.StatusSeeOther)
			return
		} else if reset != "" {
			// reset shopping cart linked list
			buyerCartll = CartLinkedList{Head: nil, Size: 0}
			http.Redirect(w, req, "/buyer", http.StatusSeeOther)
			return
		} else if checkout != "" {
			// perform checkout
			allSellerItems, ok := apiclient.GetItem("", "", user.IsBuyer)
			if !ok {
				config.Error.Println("Unable to connect to Database!")
				return
			}
			_, allitems := buyerCartll.GetAllItems()

			allok := true

			for _, item := range allitems {
				ok := updateDB(user.IsBuyer, allSellerItems, item)
				allok = allok && ok
				if !ok {
					buyerCart.Mainmessage = append(buyerCart.Mainmessage, "Error while performing check out!")
					buyerCart.Mainmessage = append(buyerCart.Mainmessage, "Try again")
				} else {
					// If ok, remove that item from the linked list
					_, index, err := buyerCartll.SearchItemandSellerName(item.Item, item.Username)
					if err != nil {
						config.Error.Println("Not able to find item in linked list")
						config.Error.Println(err)
					} else {
						_, err := buyerCartll.Remove(index)
						if err != nil {
							config.Error.Println("Not able to remove item from linked list")
							config.Error.Println(err)
						}
					}
				}
			}
			if allok {
				// Send invoice email to buyer and sellers
				email.Sendemail(user.Username, allitems)

				http.Redirect(w, req, "/buyer/checkoutsuccess", http.StatusSeeOther)
				return
			}
		}
	}

	config.TPL.ExecuteTemplate(w, "buyercart.gohtml", buyerCart)
}

func CheckoutSuccessHandler(w http.ResponseWriter, req *http.Request) {
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	user := server.GetUser(w, req)

	buyerCart := buyerCartStruct{
		Buyername:   user.Username,
		Operation:   "checkoutsuccess",
		Mainmessage: nil,
		Cartitems:   nil,
	}

	buyerCart.Mainmessage = append(buyerCart.Mainmessage, "Checkout successful!")
	buyerCart.Mainmessage = append(buyerCart.Mainmessage, "Items have been purchased!")
	buyerCart.Mainmessage = append(buyerCart.Mainmessage, "Please make payment during collection!")
	buyerCartll = CartLinkedList{Head: nil, Size: 0}

	config.TPL.ExecuteTemplate(w, "buyercart.gohtml", buyerCart)
}

func ConvertToInt(stringInput string) int {
	number, _ := strconv.ParseInt(stringInput, 10, 0)
	return int(number)
}

func ConvertToFloat(stringInput string) float64 {
	number, _ := strconv.ParseFloat(stringInput, 64)
	return float64(number)
}

func ConvertToBool(stringInput string) bool {
	done, _ := strconv.ParseBool(stringInput)
	return done
}

func updateDB(isBuyer bool, allSellerItems []apiclient.ItemsDetails, oneCartItem apiclient.ItemsDetails) bool {
	tempItem := apiclient.ItemsDetails{}
	for _, item := range allSellerItems {
		if item.Item == oneCartItem.Item && item.Username == oneCartItem.Username {
			if item.Quantity > oneCartItem.Quantity {
				tempItem.Quantity = item.Quantity - oneCartItem.Quantity
				tempItem.Cost = item.Cost
				tempItem.Item = item.Item
				tempItem.Username = item.Username
				ok := apiclient.UpdateItem(item.Item, item.Username, isBuyer, tempItem)
				return ok
			} else {
				ok := apiclient.DeleteItem(item.Item, item.Username, isBuyer)
				return ok

			}
		}
	}
	return false
}

func removeCartItems(allSellerItems []apiclient.ItemsDetails) []apiclient.ItemsDetails {
	tempItem := apiclient.ItemsDetails{}
	_, allcartitems := buyerCartll.GetAllItems()

	for _, cartitem := range allcartitems {
		for index, item := range allSellerItems {
			if item.Item == cartitem.Item && item.Username == cartitem.Username {
				if item.Quantity > cartitem.Quantity {
					tempItem.Quantity = item.Quantity - cartitem.Quantity
					tempItem.Cost = item.Cost
					tempItem.Item = item.Item
					tempItem.Username = item.Username
					//allSellerItems = append(allSellerItems[0:index], tempItem)
					//allSellerItems = append(allSellerItems, allSellerItems[index+1:]...)
					allSellerItems[index] = tempItem
				} else {
					allSellerItems = append(allSellerItems[0:index], allSellerItems[index+1:]...)
				}
				continue
			}
		}
	}
	return allSellerItems
}

/*

func AddItemToCart(w http.ResponseWriter, req *http.Request) {
	config.TPL.ExecuteTemplate(w, "addItemtocart.gohtml", nil)
}

func ItemsAddedTocart(w http.ResponseWriter, req *http.Request) {
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	user := server.GetUser(w, req)

	checkItemQty := CheckItemsQtyDB(user.IsBuyer)

	if req.Method == http.MethodPost {
		if server.AlreadyLoggedIn(req) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
		myCookie, err := req.Cookie("myCookie")
		if err != nil {
			panic(err.Error())
		}
		myList := &DetailsList{nil, 0}
		if buyername, ok := mapSessions[myCookie.Value]; ok {
			*myList = mapShoppingCart[buyername]
			apple := req.FormValue("apple")
			orange := req.FormValue("orange")
			banana := req.FormValue("banana")
			seller := req.FormValue(apple)
			appleQty, _ := strconv.Atoi(apple)
			orangeQty, _ := strconv.Atoi(orange)
			bananaQty, _ := strconv.Atoi(banana)
			errorMessage := map[string]string{}
			fmt.Println("apple", apple)
			fmt.Println("orange", orange)
			fmt.Println("banana", banana)
			fmt.Println("appleQty", appleQty)
			fmt.Println("orangeQty", orangeQty)
			fmt.Println("bananaQty", bananaQty)
			fmt.Println(seller)
			for k, v := range checkItemQty {
				if k == "apple" {
					if appleQty > v {
						errorMessage["exceedapplequantity"] = "Unable to add to cart as number of items requested exceeds quantity available"
						config.TPL.ExecuteTemplate(w, "itemlist.html", errorMessage)
						delete(errorMessage, "exceedapplequantity")
						return
					}
				} else if k == "orange" {
					if orangeQty > v {
						errorMessage["exceedorangequantity"] = "Unable to add to cart as number of items requested exceeds quantity available"
						config.TPL.ExecuteTemplate(w, "itemlist.html", errorMessage)
						delete(errorMessage, "exceedorangequantity")
						return
					}
				} else if k == "banana" {
					if bananaQty > v {
						errorMessage["exceedbananaquantity"] = "Unable to add to cart as number of items requested exceeds quantity available"
						config.TPL.ExecuteTemplate(w, "itemlist.html", errorMessage)
						delete(errorMessage, "exceedbananaquantity")
						return
					}
				}
			}
			myList.addItemDetails("apple", appleQty)
			myList.addItemDetails("orange", orangeQty)
			myList.addItemDetails("banana", bananaQty)
			mapShoppingCart[buyername] = *myList
			http.Redirect(w, req, "/checkout", http.StatusSeeOther)
		}
	}
	config.TPL.ExecuteTemplate(w, "itemlist.gohtml", checkItemQty)
	for k := range checkItemQty {
		if k == "apple" || k == "orange" || k == "banana" {
			delete(checkItemQty, k)
		}
	}
}

func ViewCart(w http.ResponseWriter, req *http.Request) {
	if !AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		panic(err.Error())
	}
	if buyername, ok := mapSessions[myCookie.Value]; ok {
		myList := mapShoppingCart[buyername]
		mapItemsAddedToCart := myList.AddAllItems()
		fmt.Println(mapItemsAddedToCart)
		if mapItemsAddedToCart["emptycart"] == 1 {
			config.TPL.ExecuteTemplate(w, "emptycart.gohtml", mapItemsAddedToCart)
		} else {
			config.TPL.ExecuteTemplate(w, "shoppingcart.gohtml", mapItemsAddedToCart)
		}
	}
}

func Checkout(w http.ResponseWriter, req *http.Request) {
	if !AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		panic(err.Error())
	}
	if buyername, ok := mapSessions[myCookie.Value]; ok {
		myList := mapShoppingCart[buyername]
		fmt.Println(myList)
		fmt.Println(mapShoppingCart)
		fmt.Println(buyername)
		mapCheckout := myList.ItemsCost()
		fmt.Println(mapCheckout)
		if mapCheckout["totalcost"] == 0 {
			config.TPL.ExecuteTemplate(w, "emptycart.gohtml", mapCheckout)
			return
		}
		if req.Method == http.MethodPost {
			myList := mapShoppingCart[buyername]
			checkoutQuantities, _ := myList.AddedItemQty()
			lock := make(chan bool, 1)
			channel := make(chan bool)
			go DeductSellerItemQuantity(checkoutQuantities, lock, channel)
			itemsDeducted := <-channel
			if itemsDeducted == true {
				http.Redirect(w, req, "/successfulcheckout", http.StatusSeeOther)
			} else if itemsDeducted == false {
				errorMessage := map[string]string{}
				errorMessage["notenoughavailableitems"] = "Sorry. There are not enough available items."
				config.TPL.ExecuteTemplate(w, "checkoutresult.gohtml", errorMessage)
				return
			}
		}
		config.TPL.ExecuteTemplate(w, "cartcheckout.gohtml", mapCheckout)
	}
}

func SuccessfulCheckout(w http.ResponseWriter, req *http.Request) {
	if !AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		panic(err.Error())
	}
	if buyername, ok := mapSessions[myCookie.Value]; ok {
		myList := mapShoppingCart[buyername]
		myList.DeleteAllItems()
		delete(mapShoppingCart, buyername)
		mapCheckoutResult := map[string]string{}
		mapCheckoutResult["successfulcheckout"] = "You have successfully bought your items! You may now click on any link below to be redirected."
		config.TPL.ExecuteTemplate(w, "checkoutresult.gohtml", mapCheckoutResult)
		delete(mapCheckoutResult, "successfulcheckout")
	}
	return
}

func CancelCheckout(w http.ResponseWriter, req *http.Request) {
	if !AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		panic(err.Error())
	}
	if buyername, ok := mapSessions[myCookie.Value]; ok {
		myList := mapShoppingCart[buyername]
		checkoutQuantities, _ := myList.AddedItemQty()
		for k := range checkoutQuantities {
			delete(checkoutQuantities, k)
		}
		//addItemQty(checkoutQuantities)
		delete(mapShoppingCart, buyername)
		myList.DeleteAllItems()
		mapCheckoutResult := map[string]string{}
		mapCheckoutResult["failedcheckout"] = "You have cancelled your order and removed all items from your shopping cart."
		config.TPL.ExecuteTemplate(w, "checkoutresult.html", mapCheckoutResult)
		delete(mapCheckoutResult, "failedcheckout")
	}

}
*/

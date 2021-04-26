package main

import "fmt"

//no descript data type
type Item struct {
	title string
	body  string
}

//creating slice of item type in memory
var database []Item

//read func read from database nad get item by name
func GetByName(title string) Item {
	//create variable which point to item in database
	var getItem Item

	//for loop to iterate through items in database
	for _, val := range database {
		/*if value is same to title string we are passing then take that value
		  and assign it to item */
		if val.title == title {
			getItem = val
		}
	}
	return getItem //return getitem variable
}

/*pass in item and pt in database and return item
so that we know what we put in database for that append item in database slice
and return item*/
func CreatItem(item Item) Item {
	database = append(database, item)
	return item

}

//below way we know which item we added to database
func AddItem(item Item) Item {
	database = append(database, item)
	return item

}

/*to edit we need to pass in title string which we get from database
and pass in edit item (new) which we want to replace old item */
func EditItem(title string, edit Item) Item {
	var changed Item
	//we grabbed edit item by its index and pass it to variable changed
	for idx, val := range database {
		if val.title == edit.title {
			database[idx] = edit
			changed = edit
		}
	}
	return changed

}

func DeleteItem(item Item) Item {
	var del Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			/*to remove item from slice we use append spread operator , take database all the way upto the index
			  of item to delete and database from after the item(+1 because idx starts from 0) we need to remove all the way untill end */
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}

	}
	return del

}

func main() {
	fmt.Println("intial database:", database)
	a := Item{"first", "a first item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("database after add item:", database)

	DeleteItem(c)
	fmt.Println("database after delete item b:", database)

	CreatItem(Item{"fourth", "a new item"})
	fmt.Println("create item", database)

	EditItem("first", Item{"fifth", "new"})
	fmt.Println("edit item", database)

	x := GetByName("fourth")
	y := GetByName("second")
	fmt.Println(x, y)
	/*for rpc function need to satisfy some criteria,
	  1.function need to be a methods
	  2.function need to be exported
	  3.function need to have two arguments both of which exported type
	  e.g GetByName func cant be called as rpc because it has only one argument,
	  however editItem func satifies becoz it takes two arguments strin and Item and item can be exported
	  4.second argument for the function must be a pointer
	  5. return type for rpc function must be error type
	*/
}

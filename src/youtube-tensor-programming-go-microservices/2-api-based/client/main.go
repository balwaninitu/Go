package main

import (
	"fmt"
	"log"
	"net/rpc"
)

/* we need to define same type here to catche receiving data
for larger application its common to shared types like creating
some kind of libraries where server and client depend on*/
type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error:", err)
	}

	a := Item{"first", "a first item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}
	//create item which can manipulate between client and server
	/*to execute method from client we will use client call method and pass
	  string representatio of method we need to execute follow by the two args for the method
	   each of the item which are passing to server we will get that item in reply variable reference
	   unlike our local programme we cant directly access the database to see if these
	   items push back to database*/
	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("database", db)
	//in below we cant change item title becaue of set up we can only change body
	client.Call("API.EditItem", Item{"second", "A new second item"}, &reply)
	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("database", db)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("first item:", reply)

}

package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

//api use to convert func to methods so all methods will on that api
type API int

var database []Item

/*below function allow to grabbed the data from client
first args is just for sake for rpc speculations of pasing two args
we dont need to pass anything to this func
it grabs the database and pass it to pointer reply which then sends back to client */
func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

/*add two args, first args represent which is passing by through caller and
second args represent result of calling this func, 2nd args is result that is
returning to the client thats calling this api  */
func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == title {
			getItem = val
		}
	}

	*reply = getItem

	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil

}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.Title == edit.Title {
			database[idx] = Item{edit.Title, edit.Body}
			changed = database[idx]
		}
	}
	*reply = changed
	return nil

}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {

			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}

	}
	*reply = del
	return nil

}

func main() {
	/*create new api type variable so that we call all method on it
	add new func to do it */
	var api = new(API)
	/*we need to use api variable in rpc register method
	so creating variable is important so that we can register its type in resgister method
	and can call it remotely*/
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}
	//register handle
	rpc.HandleHTTP()
	//open the connection
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}

	//serve the listener which we created
	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving:", err)
	}
}

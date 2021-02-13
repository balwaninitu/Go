package main

import "fmt"

func main() {

	fmt.Println("Struct")
	structs()

}

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}

	addRooms(myHouse)
	fmt.Printf("%+v\n", myHouse)

	addRoomsPtr(&myHouse)
	fmt.Printf("%+v\n", myHouse)

}

func addRooms(h house) {

	h.rooms++
}

func addRoomsPtr(h *house) {

	h.rooms++
}

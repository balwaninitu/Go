package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {

	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	px := []person{p1, p2}

	//json marshal returns byte slice and err and takes interface(can be anything)
	bs, err := json.Marshal(px)
	if err != nil {
		log.Panic(err) //if any error in marshaling is programmer error so panic  to unstack
	}
	fmt.Println("Print Marshal", string(bs))

	px2 := []person{}
	//unmarshal returns error and takes byte slice and pointer to interface
	//unmarshal is back to go data structure
	err = json.Unmarshal(bs, &px2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Print Unmarshal", px2)
}

package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	fName        string
	lName        string
	userID       int
	invoiceTotal float64
}

func inspect(n interface{}) {
	refType := reflect.TypeOf(n)
	refValue := reflect.ValueOf(n)
	fmt.Println("Num of fields", refType.NumField())
	for i := 0; i < refType.NumField(); i++ {
		fmt.Println(refType.Field(i).Name, "value:", refValue.Field(i), "type:", refType.Field(i).Type)
	}

}

func main() {
	custome1 :=

		customer{

			fName:        "John",
			lName:        "Wick",
			userID:       123123123,
			invoiceTotal: 10000,
		}
	inspect(custome1)

}

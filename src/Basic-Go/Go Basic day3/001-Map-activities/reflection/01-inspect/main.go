package main

import (
	"fmt"
	"reflect"
)

func inspect(n interface{}) {
	refType := reflect.TypeOf(n)

	fmt.Println("Content:", n, "Name:", refType.Name(), "Kind:", refType.Kind())
}

func main() {

	msg := "this is a string"
	inspect(msg)
	number := 12345
	inspect(number)
	numberFloat := 1.2345
	inspect(numberFloat)
	condition := true
	inspect(condition)

}

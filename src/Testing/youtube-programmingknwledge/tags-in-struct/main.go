package main

import (
	"fmt"
	"reflect"
)

//User can be exported
type User struct {
	Username string `maxsize: "10"`
	Password string `minsize: "8"`
}

func main() {

	u := User{
		Username: "",
		Password: "",
	}

	t := reflect.TypeOf(u)
	//field := t.Field(0)
	//alternative
	field, _ := t.FieldByName("Password")

	fmt.Println("t", t)
	fmt.Println(field.Tag)
	//fmt.Println(field.Tag.Get("maxsize"))

}

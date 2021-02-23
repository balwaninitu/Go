package main

import (
	"fmt"
	"text/template"
)

func main() {
	//new undefined template with given name
	tpl := template.New("hi")

	fmt.Println(tpl)

}

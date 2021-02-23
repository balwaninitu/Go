package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type family struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	myfamily := family{
		Name: "Riddham",
		Age:  8,
	}

	err := tpl.Execute(os.Stdout, myfamily)
	if err != nil {
		log.Fatalln(err)
	}
}

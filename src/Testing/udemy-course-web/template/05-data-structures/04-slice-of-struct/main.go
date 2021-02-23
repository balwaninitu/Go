package main

import (
	"log"
	"os"
	"text/template"
)

type family struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	daddy := family{
		Name: "Santosh",
		Age:  39,
	}

	mummy := family{
		Name: "Nitu",
		Age:  37,
	}

	baby := family{
		Name: "Riddham",
		Age:  8,
	}

	myFamily := []family{daddy, mummy, baby}

	err := tpl.Execute(os.Stdout, myFamily)
	if err != nil {
		log.Fatalln(err)
	}
}

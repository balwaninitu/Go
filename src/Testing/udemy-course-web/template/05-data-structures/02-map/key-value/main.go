package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	family := map[string]string{
		"daddy": "Santosh",
		"mummy": "Nitu",
		"baby":  "Riddham",
	}

	err := tpl.Execute(os.Stdout, family)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"
	"os"
	"text/template"
)

//took dot from template and put it into variable and variable replace instead of dot in html doc
var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {

	//tpl = template.Must(template.ParseFiles("tpl.gohtml"))

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "focus")
	if err != nil {
		log.Fatalln(err)
	}

}

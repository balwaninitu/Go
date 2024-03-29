package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "nice!!") //data into template
	if err != nil {
		log.Fatalln(err)
	}
}

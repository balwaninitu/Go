package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "mummy.hi", nil)
	err = tpl.ExecuteTemplate(os.Stdout, "daddy.hi", nil)
}

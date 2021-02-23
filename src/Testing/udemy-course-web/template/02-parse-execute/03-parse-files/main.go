package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("riddham.hi")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = template.ParseFiles("santosh.hi")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = template.ParseFiles("nitu.hi")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = template.ParseFiles("riddham.hi", "nitu.hi", "santosh.hi")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "nitu.hi", nil)
	err = tpl.ExecuteTemplate(os.Stdout, "riddham.hi", nil)
}

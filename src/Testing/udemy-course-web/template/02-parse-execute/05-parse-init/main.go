package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	//mustfunc use to avoid error checking as it returns if there ia any error
	//parseGlob returns *template and error whereas must take it as input
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "daddy.hi", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "mummy.hi", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "baby.hi", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

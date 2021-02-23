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

type friends struct {
	Name   string
	Weight float64
	IsGood bool
}

type myInfo struct {
	MyFamily  []family
	MyFriends []friends
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

	best := friends{
		Name:   "Kani",
		Weight: 20.8,
		IsGood: true,
	}

	nice := friends{
		Name:   "Gauri",
		Weight: 20.1,
		IsGood: true,
	}

	good := friends{
		Name:   "Sadhu",
		Weight: 29.8,
		IsGood: false,
	}

	myFriends := []friends{best, nice, good}

	me := myInfo{
		MyFamily:  myFamily,
		MyFriends: myFriends,
	}

	err := tpl.Execute(os.Stdout, me)
	if err != nil {
		log.Fatalln(err)
	}

}

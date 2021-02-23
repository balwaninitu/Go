package main

import (
	"log"
	"os"
	"strings"
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

// var fm = template.FuncMap{
// 	"uc": strings.ToUpper,
// 	"ft": firstThree,
// }

var fm = template.FuncMap{
	"lc": strings.ToLower,
	"tc": strings.ToTitle,
}

var tpl *template.Template

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {

	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

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

	myInfo := struct {
		MyFamily  []family
		MyFriends []friends
	}{

		MyFamily:  myFamily,
		MyFriends: myFriends,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", myInfo)
	if err != nil {
		log.Fatalln(err)
	}

}

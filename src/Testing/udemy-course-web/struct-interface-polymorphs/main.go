package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	isFunny bool
}

type human interface {
	speak()
}

func saysomething(h human) {
	h.speak()

}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says "Good morning Nitu"`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says "you are shaken but not stirred"`)
}

func main() {

	p1 := person{fname: "miss", lname: "moneypenny"}

	fmt.Println(p1)
	p1.speak()

	m1 := secretAgent{
		person{fname: "James", lname: "Bond"},
		true,
	}
	fmt.Println(m1)
	saysomething(p1) //polymorphism meaning many human changing
	saysomething(m1)

}

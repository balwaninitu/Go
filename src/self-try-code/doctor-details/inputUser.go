package main

import "fmt"

func inputUser(pn string, dn string, d string, t int) {
	var name string
	var doctorName string
	var day string
	var time int
	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	fmt.Println("Enter doctor name")
	fmt.Scanln(&doctorName)
	fmt.Println("Enter day of appointment")
	fmt.Scanln(&day)
	fmt.Println("Enter time of appointment")
	fmt.Scanln(&time)
}

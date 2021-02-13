package main

import "fmt"

type customer struct {
	firstName string
	lastName  string
	username  string
	password  string
	email     string
	phone     int
	address   string
}

func (c customer) printUserCredential() (string, string) {
	return c.username, c.password
}

func (c customer) printAddress() string {
	return c.address
}

func (c customer) printAllInfo() {
	fmt.Println(c.firstName, c.lastName, c.username, c.password, c.email, c.phone, c.address)
}

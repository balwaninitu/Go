package main

import "fmt"

type details struct {
	id           int
	name         string
	availableDay string
	next         *details
}

type appointmentList struct {
	name string
	head *details
	book *details
	size int
}

func createAppointmentList(name string) *appointmentList {
	return &appointmentList{
		name: name,
	}
}

func (a *appointmentList) addDetails(id int, name string, day string) error {
	d := &details{
		id:           id,
		name:         name,
		availableDay: day,
	}
	if a.head == nil {
		a.head = d
	} else {
		currentNode := a.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = d
	}
	return nil
}

func (a *appointmentList) showAllDetails() error {
	currentNode := a.head
	if currentNode == nil {
		fmt.Println("No Details found")
		return nil
	}
	fmt.Printf("%s, has appointment with %d on %s\n", *&currentNode.name, currentNode.id, currentNode.availableDay)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode)
	}
	return nil
}

func (a *appointmentList) startBooking(pn string, id int) *details {
	a.book = a.head
	return a.book
}

func (a *appointmentList) nextBooking(id int) *details {
	a.book = a.book.next
	return a.book
}

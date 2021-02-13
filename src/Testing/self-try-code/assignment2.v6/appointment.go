package main

import (
	"fmt"
	"time"
)

type appointmentDetails struct {
	appointmentID int
	patientName   string
	doctor        doctorDetails
	next          *appointmentDetails
}

type appointmentList struct {
	name string
	head *appointmentDetails
	size int
}

func createAppointmentList(name string) *appointmentList {
	return &appointmentList{
		name: name,
	}
}

func (a *appointmentList) addAppointmentDetails(*appointmentDetails) error {
	apt := &appointmentDetails{
		appointmentID: 1,
		patientName:   "Pt1",
		doctorDetails{
			id:            1,
			name:          "dr1",
			availableTime: time.Time,
			availability:  true,
		},
	}
	if a.head == nil {
		a.head = apt
	} else {
		currentNode := a.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = apt
	}
	return nil
}

func (a *appointmentList) showAllDetails() error {
	currentNode := a.head
	if currentNode == nil {
		fmt.Println("No Details found")
		return nil
	}

	fmt.Printf("%s, your appointmnet with %s has created\n", currentNode.patientName, currentNode.doctorName)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode)
	}
	return nil
}

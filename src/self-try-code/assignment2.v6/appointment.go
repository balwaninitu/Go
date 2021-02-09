package main

import (
	"fmt"
	"time"
)

type appointmentDetails struct {
	patientName string
	doctorName  string
	doctorID    int
	aptTime     time.Time
	next        *appointmentDetails
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

func (a *appointmentList) addAppointmentDetails(pn string, dn string, dID int, aptTime time.Location) error {
	apt := &appointmentDetails{
		patientName: pn,
		doctorName:  dn,
		doctorID:    dID,
		aptTime:     time.Date(2021, time.February, time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Local),
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

	fmt.Printf("%s, your appointmnet with %s on %v has created\n", *&currentNode.patientName, *&currentNode.doctorName, *&currentNode.aptTime)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode)
	}
	return nil
}

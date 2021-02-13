package main

import "fmt"

type appointment struct {
	patientName       string
	dayOfAppointment  string
	timeOfAppointment string
	doctorName        string
	next              *appointment
}

type doctorList struct {
	head *appointment
	name string
}

func createAppointment(n string) *doctorList {
	return &doctorList{
		name: n,
	}
}

func (d *doctorList) addAppointmentDetails(pn string, dn string, day string, t string) error {
	appt := &appointment{
		patientName:       pn,
		doctorName:        dn,
		dayOfAppointment:  day,
		timeOfAppointment: t,
	}
	if d.head == nil {
		d.head = appt
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = appt
	}
	return nil
}

func (d *doctorList) showAllAppointment() error {
	fmt.Printf("\nCurrent Appointments:\n")
	fmt.Println("-------------")
	currentNode := d.head
	if currentNode == nil {
		fmt.Println("Empty list")
		return nil
	}
	fmt.Printf("Patien Name: %s Appointment Day: %s Appointment Time: %s Doctor Name: %s\n", *&currentNode.patientName, currentNode.dayOfAppointment, currentNode.timeOfAppointment, currentNode.doctorName)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("Patien Name: %s Appointment Day: %s Appointment Time: %s Doctor Name: %s\n", *&currentNode.patientName, currentNode.dayOfAppointment, currentNode.timeOfAppointment, currentNode.doctorName)
	}
	fmt.Println("-------------")
	return nil
}

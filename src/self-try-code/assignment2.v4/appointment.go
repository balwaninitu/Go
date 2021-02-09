package main

import (
	"fmt"
	"sync"
)

type appointment struct {
	patientName      string
	dayOfAppointment string
	doctorID         int
	next             *appointment
}

type doctorList struct {
	head *appointment
	now  *appointment
	name string
	size int
	lock sync.RWMutex
}

func createAppointment(n string) *doctorList {
	return &doctorList{
		name: n,
	}
}

func (d *doctorList) addAppointmentDetails(pn string, dID int, day string) error {
	d.lock.Lock()
	appt := appointment{
		patientName:      pn,
		doctorID:         dID,
		dayOfAppointment: day,
	}
	if d.head == nil {
		d.head = &appt
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &appt
	}
	d.size++
	d.lock.Unlock()
	return nil
}

func (d *doctorList) Head() *appointment {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.head
}

func (d *doctorList) show() *appointment {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.head
}

func (d *doctorList) showAllAppointment() error {
	fmt.Printf("\nCurrent Appointments:\n")
	fmt.Println("-------------")
	currentNode := d.head
	if currentNode == nil {
		fmt.Println("Empty list")
		return nil
	}
	fmt.Printf("Patien Name: %s Appointment Day: %s Doctor ID: %d\n", *&currentNode.patientName, currentNode.dayOfAppointment, currentNode.doctorID)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("Patien Name: %s Appointment Day: %s Doctor ID: %d\n", *&currentNode.patientName, currentNode.dayOfAppointment, currentNode.doctorID)
	}
	fmt.Println("-------------")
	return nil
}

func (d *doctorList) startBooking() *appointment {
	d.now = d.head
	return d.now
}

func (d *doctorList) nextApt() *appointment {
	d.now = d.now.next
	return d.now
}

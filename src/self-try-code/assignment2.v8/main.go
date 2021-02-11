package main

import (
	"fmt"
	"time"
)

type appointment struct {
	patientName string
	next        *appointment
}

type aptList struct {
	head *appointment
	len  int
}

type doctorDetails struct {
	drID       int
	doctorName string
	DayTime    time.Time
	available  bool
}

func main() {

	t1, _ := time.Parse(time.RFC822, "17 FEB 21 10:00 SGT")
	t2, _ := time.Parse(time.RFC822, "18 FEB 21 10:00 SGT")

	d1 := doctorDetails{drID: 1, doctorName: "Dr1", DayTime: t1, available: true}
	d2 := doctorDetails{drID: 2, doctorName: "Dr2", DayTime: t2, available: true}
	d3 := doctorDetails{drID: 3, doctorName: "Dr3", DayTime: t1, available: true}

	doctorList := []doctorDetails{d1, d2, d3}
	appointmentList := []appointment{}
	fmt.Println(doctorList)
	fmt.Println(appointmentList)

	var patientName string
	var doctorName string
	//var dv doctorDetails
	fmt.Println("enter patient name")
	fmt.Scanln(&patientName)
	//show list of doctor name & ask to select
	fmt.Println("enter doctor name")
	fmt.Scanln(&doctorName)
	for _, dv := range doctorList {
		if dv.doctorName == doctorName {
			fmt.Printf("%d. %s is available on %v\n", dv.drID, dv.doctorName, dv.DayTime)
			break
		}
		fmt.Println("doctor not avaialble")
		break
	}
	clinicList := &aptList{nil, 0}
	clinicList.addAppointmentDetails(patientName)
	clinicList.PrintNode()
	fmt.Printf("there is %d patient in the list\n", clinicList.len)

	// var d = doctorDetails{
	// 	drID:       dv.drID,
	// 	doctorName: dv.doctorName,
	// 	DayTime:    dv.DayTime,
	// }

	// var a = appointment{
	// 	aptID:      len(appointmentList),
	// 	doctorName: dv.doctorName,
	// }
	//fmt.Println("enter to confirm appointment")
	//fmt.Scanln()
	//fmt.Printf("%s, your appointment with %s is confirmed on %v\n", patientName, dv.doctorName, dv.DayTime)

	//myList.PrintNode()
}

func (a *aptList) addAppointmentDetails(patientName string) {
	newPatient := &appointment{patientName, nil}
	if a.head == nil {
		a.head = newPatient
	} else {
		currentPatient := a.head
		for currentPatient.next != nil {
			currentPatient = currentPatient.next
		}
		currentPatient.next = newPatient
	}
	a.len++
}

func (a *aptList) PrintNode() {
	currentPatient := a.head
	if currentPatient == nil {
		fmt.Println("List is empty")
	}
	fmt.Printf("%s\n", currentPatient.patientName)
	for currentPatient.next != nil {
		currentPatient = currentPatient.next
		fmt.Printf("%s\n", currentPatient.patientName)
	}
}

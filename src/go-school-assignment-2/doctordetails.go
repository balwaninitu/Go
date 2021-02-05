package main

// import (
// 	"fmt"
// )

// type doctorDetails struct {
// 	num           int
// 	name          string
// 	availableDay  []string
// 	availableTime int
// }

// type appointment struct {
// 	patientName       string
// 	dayOfAppointment  string
// 	timeOfAppointment int
// 	doctorName        string
// 	next              *appointment
// }

// type doctorList struct {
// 	name string
// 	head *appointment
// }

// func createAppointment(n string) *doctorList {
// 	return &doctorList{
// 		name: n,
// 	}
// }

// func (d *doctorList) addAppointmentDetails(pn string, dn string, day string, t int) error {
// 	appt := &appointment{
// 		patientName:       pn,
// 		doctorName:        dn,
// 		dayOfAppointment:  day,
// 		timeOfAppointment: t,
// 	}

// 	if d.head == nil {
// 		d.head = appt
// 	} else {
// 		currentNode := d.head
// 		for currentNode.next != nil {
// 			currentNode = currentNode.next
// 		}
// 		currentNode.next = appt
// 	}
// 	return nil
// }

// var m map[int]doctorDetails

// func main() {

// 	m = make(map[int]doctorDetails)

// 	m[1] = doctorDetails{"Sunny Je", "Mon, Wed", "2pm, 4pm, 5pm"}
// 	m[2] = doctorDetails{"Joyce Tan", "Tue, Thurs", "1pm, 2pm, 4pm"}
// 	m[3] = doctorDetails{"Anthony Lim", "Wed, Fri", "2 pm, 3pm, 4pm"}

// 	for _, v := range m {
// 		fmt.Printf("Doctor name: %s Avaialble days:%s Available Time: %s\n", v.name, v.availableDay, v.availableTime)
// 	}

// 	//myAppointment := createAppointment("myAppointment")

// 	//make appointment
// 	var name string
// 	var id int
// 	var date int
// 	var time int
// 	fmt.Println("Enter your name")
// 	fmt.Scanln(&name)
// 	fmt.Println("Enter doctor ID")
// 	fmt.Scanln(&id)
// 	fmt.Println("Enter date of appointment")
// 	fmt.Scanln(&date)
// 	fmt.Println("Enter date of appointment")
// 	fmt.Scanln(&time)
// 	fmt.Printf("Hello %s. You have requested for appointment on date:%d time:%d\n", name, date, time)
// 	fmt.Println("Below are the availble doctor on date and time selected by you")

// }

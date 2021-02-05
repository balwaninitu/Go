package main

import "fmt"

type details struct {
	availableDays []string
	availableTime []string
}

var k string
var v details

func main() {
	m := make(map[string]details)
	m["Dr1"] = details{availableDays: []string{"Mon", "Tue", "Fri"}, availableTime: []string{"1pm", "2pm", "4pm"}}
	m["Dr2"] = details{availableDays: []string{"Mon", "Wed", "Thurs"}, availableTime: []string{"1pm", "2pm", "3pm"}}
	m["Dr3"] = details{availableDays: []string{"Mon", "Tue", "Wed", "Fri"}, availableTime: []string{"3pm", "4pm", "5pm"}}
	for k, v = range m {
		fmt.Printf("Doctor:%s Avalable on:%s time:%s\n", k, v.availableDays, v.availableTime)
	}

	myAppointment := createAppointment("myAppointment")
	fmt.Println("Created appointmentList")

	var aptName string
	var aptDoctorName string
	var aptDay string
	var aptTime string
	//before this add welcome message and display doctors name along with details of their specialisation
	//make appointment section based on doctor availability
	fmt.Println("Enter your name")
	fmt.Scanln(&aptName)
	fmt.Println("Enter doctor name")
	fmt.Scanln(&aptDoctorName)
	fmt.Printf("%s is available on Days:%s Time:%s\n", aptDoctorName, v.availableDays, v.availableTime) //list shows doctor available days & time to book
	fmt.Println("Pleae select day of appointmen from available slots")
	fmt.Scanln(&aptDay)
	fmt.Println("Pleae select time of appointmen from available slots")
	fmt.Scanln(&aptTime)

	myAppointment.addAppointmentDetails(aptName, aptDoctorName, aptDay, aptTime)
	//fmt.Printf("Thank you %s!\nYour Appointment has been added", name)
	fmt.Println(myAppointment.showAllAppointment())

	//check if selected dr available on selected day & time
	// var found bool

	// v, found = m[day]

	// if found == true {
	// 	fmt.Println("we Found")
	// 	fmt.Println("%s is available on %s\n", m[day], v.days)
	// } else {
	// 	fmt.Println("not found")
	// }

}

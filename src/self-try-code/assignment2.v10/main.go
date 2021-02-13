package main

import (
	"time"
)

const (
	title = `
Welcome to "GO" Dental Clinic!
You are at "Appointment Booking Application"`

	listDisplay = `
1        :Make Appointment
2        :Lists all Avaialble Doctors
3        :Search available Doctor by Name(eg.Name DR1)
4        :Show all booked Appointment
5        :Delete Appointment(admin only)
6        :Exit
select your choice:`
)

type doctorDetails struct {
	drID       int
	doctorName string
	DayTime    time.Time
	available  bool
}

//Appointment can be exported to another package
type Appointment struct {
	aptID       int
	patientName string
	doctor      doctorDetails
	next        *Appointment
}

//ClinicAppointmentList can be exported to another package
type ClinicAppointmentList struct {
	start  *Appointment
	length int
}

func main() {
	t1, _ := time.Parse(time.RFC822, "17 FEB 21 10:00 SGT")
	t2, _ := time.Parse(time.RFC822, "17 FEB 21 02:00 SGT")
	t3, _ := time.Parse(time.RFC822, "18 FEB 21 10:00 SGT")
	t4, _ := time.Parse(time.RFC822, "18 FEB 21 02:00 SGT")

	d1 := doctorDetails{drID: 1, doctorName: "Dr1", DayTime: t1, available: true}
	d2 := doctorDetails{drID: 2, doctorName: "Dr2", DayTime: t2, available: true}
	d3 := doctorDetails{drID: 3, doctorName: "Dr3", DayTime: t3, available: true}
	d4 := doctorDetails{drID: 4, doctorName: "Dr4", DayTime: t4, available: true}

	doctorList := []doctorDetails{d1, d2, d3, d4}

	appointmentList := &ClinicAppointmentList{}

	for {
		userAppointmentListInput := displayList()
		userChoiceAction(userAppointmentListInput, &doctorList, appointmentList)
	}
}

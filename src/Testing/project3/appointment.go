package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// This struct wil be used to define doctor details
type doctorDetails struct {
	drID          int
	appointmentID int
	doctorName    string
	DayTime       time.Time
	available     bool
}

//Appointment can be exported to another package
// Node Appointment for cretion of Appointment
type Appointment struct {
	aptID       int
	patientName string
	doctor      doctorDetails
	next        *Appointment
}

//ClinicAppointmentList can be exported to another package
// ClinicAppointmentList linked List will be used for storing all appointment.
type ClinicAppointmentList struct {
	start  *Appointment
	length int
}

func makeAppointment(doctorList *[]doctorDetails, appointmentList *ClinicAppointmentList) []doctorDetails {
	var patientName, doctorName string
	fmt.Println()
	fmt.Println("******Appointment Booking System******")
	fmt.Print("Enter Patient Name ")
	_, err := fmt.Scanln(&patientName)
	if err != nil {
		fmt.Println(errors.New("Error:Unexpected new line"))
	}

	fmt.Print("Enter Doctor Name(eg. DR1) ")
	_, err = fmt.Scanln(&doctorName)
	if err != nil {
		fmt.Println(errors.New("Error:Unexpected new line"))
	}

	var tempDoctorName = strings.ToUpper(strings.TrimSpace(doctorName))
	var tempPatientName = strings.ToUpper(strings.TrimSpace(patientName))
	fmt.Println()
	fmt.Printf("Available time slots for doctor %s\n", tempDoctorName)
	available := searchDoctorByName(&(*doctorList), tempDoctorName)

	if available {
		var slot int
		fmt.Print("Enter time slot(number display beside doctor name) for booking (eg.1 to book Timeslot 1): ")
		_, err := fmt.Scanln(&slot)
		if err != nil {
			fmt.Println(errors.New("Error:Unexpected new line"))
		}
		for index, doctorValue := range *doctorList {

			//if strings.ToUpper(strings.TrimSpace(doctorValue.doctorName)) == tempDoctorName && doctorValue.available {
			if index+1 == slot && (*doctorList)[index].available {
				var d1 = doctorDetails{
					drID:       doctorValue.drID,
					doctorName: doctorValue.doctorName,
					DayTime:    doctorValue.DayTime,
					available:  false,
				}

				a1 := Appointment{
					aptID:       appointmentList.length + 1,
					patientName: tempPatientName,
					doctor:      d1,
				}

				appointmentList.Append(&a1)
				(*doctorList)[slot-1].available = false
				fmt.Println()
				fmt.Printf("Timeslot %d booked\n", slot)
				fmt.Println("\n***Booking successful***")
				fmt.Printf("\tAppointment Id:- %d\n", appointmentList.length)
				fmt.Printf("\tAppointment booked for patient %s with doctor %s on %s\n", tempPatientName, tempDoctorName, doctorValue.DayTime.Format(time.ANSIC))
				break
			}
		}

	} else {
		fmt.Printf("Doctor %s not found or invalid timeslot selected\n", tempDoctorName)
	}
	return *doctorList
}

//Append function can be exported
//linked list method will receive appointments and get append into the list
func (c *ClinicAppointmentList) Append(newAppointment *Appointment) {

	if c.length == 0 {
		c.start = newAppointment
	} else {
		currentAppointment := c.start
		for currentAppointment.next != nil {
			currentAppointment = currentAppointment.next
		}
		currentAppointment.next = newAppointment
	}
	c.length++
}

//linked list of booked appointment list will get display by this function, all the current book appointment can be seen
func displayAllBookedAppointments(appointmentList *ClinicAppointmentList) {
	if appointmentList.length == 0 {
		fmt.Println("Appointment list is Empty")
		return
	}
	index := 1
	appointmentListValue := appointmentList.start
	fmt.Println()
	fmt.Println("List of booked Appointments:")
	fmt.Println(strings.Repeat("=", 40))
	for appointmentListValue != nil {
		fmt.Printf("%d) %s %s %v\n", appointmentListValue.aptID, appointmentListValue.patientName, appointmentListValue.doctor.doctorName, appointmentListValue.doctor.DayTime.Format(time.ANSIC))
		index = index + 1
		appointmentListValue = appointmentListValue.next
	}
}

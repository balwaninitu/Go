package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func makeAppointment(doctorList *[]doctorDetails, appointmentList *ClinicAppointmentList) []doctorDetails {
	var patientName, doctorName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("******Appointment Booking System******")
	fmt.Print("Enter Patient Name: ")
	patientName, _ = reader.ReadString('\n')

	//conver CRLF to LF
	patientName = strings.Replace(patientName, "\n", "", -1)

	fmt.Print("Enter Doctor Name(eg. dr1): ")
	doctorName, _ = reader.ReadString('\n')
	//conver CRLF to LF
	doctorName = strings.Replace(doctorName, "\n", "", -1)

	var tempDoctorName = strings.ToUpper(strings.TrimSpace(doctorName))
	var tempPatientName = strings.ToUpper(strings.TrimSpace(patientName))

	fmt.Printf("Here are available time slots for doctor %s\n", tempDoctorName)
	available := searchDoctorByName(&(*doctorList), doctorName)
	if available {
		var slot int
		fmt.Print("Enter time slot for booking (eg. 1 to book Timeslot 1): ")
		fmt.Scanf("%d", &slot)
		for _, doctorValue := range *doctorList {

			if strings.ToUpper(strings.TrimSpace(doctorValue.doctorName)) == tempDoctorName && doctorValue.available {
				var d1 = doctorDetails{
					drID:       doctorValue.drID,
					doctorName: doctorValue.doctorName,
					DayTime:    doctorValue.DayTime,
					available:  true}

				a1 := Appointment{
					aptID:       appointmentList.length + 1,
					patientName: tempPatientName,
					doctor:      d1,
				}
				available = true
				appointmentList.Append(&a1)
				(*doctorList)[slot-1].available = false
				fmt.Printf("Timeslot %d booked", slot)
				fmt.Println("\n***Booking successful***")
				fmt.Printf("\tAppointment Id:- %d\n", appointmentList.length)
				fmt.Printf("\tAppointment booked for patient %s with doctor %s on %s\n", tempPatientName, tempDoctorName, doctorValue.DayTime.Format(time.ANSIC))

			}
		}

	} else {
		fmt.Printf("%s not found or invalid timeslot selected\n", doctorName)
	}
	return *doctorList
}

//Append function can be exported
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

func displayAllBookedAppointments(appointmentList *ClinicAppointmentList) {
	if appointmentList.length == 0 {
		fmt.Println("Appointment list is Empty")
		return
	}
	index := 1
	appointmentListValue := appointmentList.start
	for appointmentListValue != nil {
		fmt.Printf("%d) %s %s %v\n", appointmentListValue.aptID, appointmentListValue.patientName, appointmentListValue.doctor.doctorName, appointmentListValue.doctor.DayTime.Format(time.ANSIC))
		index = index + 1
		appointmentListValue = appointmentListValue.next
	}
}

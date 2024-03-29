package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

/*#5 and #6 are for admin only which are protected by password
func readAdminPassword is used to securely open sensitive document which for admin use only
password information can get from chaneel running concurrently*/
func readAdminPassword(ch chan string) {
	defer recoverFromPanic()
	text, err := ioutil.ReadFile("C:/Projects/Go/src/project3/password.txt")
	if err != nil {
		panic(errors.New("Wrong file name or path"))
	}
	//convert byte to string
	password := string(text)
	//println(password)
	ch <- password
}

//password enter function, get error if there is no input.
func enterAdminPassword() string {
	var adminPassword string
	fmt.Print("Enter Admin Password ")
	_, err := fmt.Scanln(&adminPassword)
	if err != nil {
		fmt.Println(errors.New("Error:No input"))
	}
	return adminPassword
}

/*appointments can be search by its ID
linked list will help to track available appointments
only access to admin and supported by remove function of linked list*/
func deleteAppointment(ClinicAppointmentList *ClinicAppointmentList) {
	var aptID int
	fmt.Print("Enter Appointment Id to be delete ")
	_, err := fmt.Scanln(&aptID)
	if err != nil {
		fmt.Println(errors.New("Error:No input"))
	} else {
		available := ClinicAppointmentList.Remove(aptID)
		if available {
			fmt.Printf("Appointment id %d deleted successfully!\n", aptID)
		} else {
			fmt.Printf("Appointment id %d not found!\n", aptID)
		}
	}
}

// Remove method can be exported
//Remove function is access by admin only to delete past or current appointments if need arise.
func (c *ClinicAppointmentList) Remove(aptID int) bool {

	/* Appointment list is empty
	when empty list panics, the deffered function will
	be called which uses recover to stop the panicking sequence*/
	defer recoverFromPanic()
	if c.length == 0 {
		panic(errors.New("Appointment list is empty"))
	}

	available := false
	//To delete first appointment - move head pointer
	currentAppointment := c.start
	if c.start.aptID == aptID {
		c.start = currentAppointment.next
		c.length--
		available = true

	} else {

		//To delete middle appointment - need two pointers (previous and next)
		previousAppointment := c.start
		currentAppointment = c.start.next
		for currentAppointment.next != nil {
			if currentAppointment.aptID == aptID {
				previousAppointment.next = currentAppointment.next
				c.length--
				available = true

			}
			currentAppointment = currentAppointment.next
			previousAppointment = previousAppointment.next
		}

		//To delete last appointment need to confirm if node is nil or not
		if currentAppointment.next == nil && currentAppointment.aptID == aptID {
			previousAppointment.next = nil
			c.length--
		}
	}
	return available
}

//below function will help admin to add docotr schedule for future bookings
//date and time slot of doctor are to be added in the given format
func creatDoctorSchedule(doctorList *[]doctorDetails) {
	var doctorName, doctorDateSlot, doctorTimeSlot string
	fmt.Println()
	fmt.Print("Enter Doctor Name ")
	_, err := fmt.Scanln(&doctorName)
	if err != nil {
		fmt.Println(errors.New("Error:No input"))
	}
	var tempDoctorName = strings.ToUpper(strings.TrimSpace(doctorName))

	fmt.Print("Enter Doctor available Date in YYYY-MM-DD format [example :- 2020-02-13] : ")
	_, err = fmt.Scanln(&doctorDateSlot)
	if err != nil {
		fmt.Println(errors.New("Error:No input"))
	}
	_, err = time.Parse("2006-01-02", doctorDateSlot)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter Doctor available time in HH:MM format [example :- 16:00] ")
	_, err = fmt.Scanln(&doctorTimeSlot)
	if err != nil {
		fmt.Println(errors.New("Error:No input"))
	}
	doctorTimeSlot = doctorTimeSlot + ":00.000"

	_, err = time.Parse("15:04:05.000", doctorTimeSlot)
	if err != nil {
		fmt.Println(err)
		return
	}

	doctorDateTimeSlot := doctorDateSlot + " " + doctorTimeSlot
	fmt.Println(doctorDateTimeSlot)
	dt, err := time.Parse("2006-01-02 15:04:05.000", doctorDateTimeSlot)
	fmt.Println(dt)
	if err != nil {
		fmt.Println(err)
		return
	}

	var d1 = doctorDetails{
		drID:          len(*doctorList) + 1,
		appointmentID: 0,
		doctorName:    tempDoctorName,
		DayTime:       dt,
		available:     true,
	}

	*doctorList = append(*doctorList, d1)
	displayAllDoctorAvailableTime(*doctorList)

}

//check error and return appropriate message if there is any error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//to regain control of a panicking program
func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from panic", r)
	}
}

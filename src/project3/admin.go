package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func readAdminPassword(ch chan string) {
	text, err := ioutil.ReadFile("C:/Projects/Go/src/project2/password.txt")
	check(err)
	//convert byte to string
	password := string(text)
	//println(password)
	ch <- password
}

func enterAdminPassword() string {
	var adminPassword string
	fmt.Print("Enter Admin Password: ")
	fmt.Scanln(&adminPassword)
	return adminPassword
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Append function can be exported
//append function will add booked appointmets into the list
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

//Remove func can be exported
//Remove function is access by admin only to delete past or current appointments if need arise, appointments can be search by its ID
//linked list will help to track
func (c *ClinicAppointmentList) Remove(aptID int) bool {

	// Appointment list is empty
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

//only access to admin and supported by remove function of linked list
func deleteAppointment(ClinicAppointmentList *ClinicAppointmentList) {
	var aptID int
	fmt.Print("Enter Appointment Id to be delete: ")
	_, err := fmt.Scanln(&aptID)
	if err != nil {
		log.Print("Scan for id failed, due to", err)
	} else {
		available := ClinicAppointmentList.Remove(aptID)
		if available {
			fmt.Printf("Appointment id %d deleted successfully!\n", aptID)
		} else {
			fmt.Printf("Appointment id %d not found!\n", aptID)
		}
	}
}

func creatDoctorSchedule(doctorList *[]doctorDetails) {
	var doctorName, doctorDateSlot string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Doctor Name: ")
	doctorName, _ = reader.ReadString('\n')
	doctorDateSlot = strings.Replace(doctorName, "\n", "", -1)

	fmt.Println()
	fmt.Print("Enter Doctor available Date in YYYY-MMM-DD format [example :- 2021-FEB-12] :")
	doctorDateSlot, _ = reader.ReadString('\n')
	fmt.Print(doctorDateSlot)
	doctorDateSlot = strings.Replace(doctorDateSlot, "\r\n", "", -1)
	_, err := time.Parse("2006-FEB-02", doctorDateSlot)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Print("Enter Doctor available time in HH:MM format [example :- 16:00] ")
	doctorTimeSlot, _ := reader.ReadString('\n')
	doctorTimeSlot = strings.Replace(doctorTimeSlot, "\r\n", "", -1)
	doctorTimeSlot = doctorTimeSlot + ":00.000"

	_, err = time.Parse("15:04:05.000", doctorTimeSlot)
	if err != nil {
		fmt.Println(err)
		return
	}
	doctorDateTimeSlot := doctorDateSlot + " " + doctorTimeSlot
	dt, err := time.Parse("2006-FEB-02 15:04:05.000", doctorDateTimeSlot)

	if err != nil {
		fmt.Println(err)
		return
	}

	var d1 = doctorDetails{
		drID:          len(*doctorList) + 1,
		appointmentID: 0,
		doctorName:    doctorName,
		DayTime:       dt,
		available:     true,
	}

	*doctorList = append(*doctorList, d1)
	displayAllDoctorAvailableTime(*doctorList)

}

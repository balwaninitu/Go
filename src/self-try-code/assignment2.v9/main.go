package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
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
		userAppointmentListInput := displayAppointmentList()
		userChoiceAction(userAppointmentListInput, &doctorList, appointmentList)
	}
}
func displayAppointmentList() int {
	fmt.Println()
	fmt.Println(strings.TrimSpace(title))
	fmt.Println(strings.Repeat("=", 40))
	fmt.Print(strings.TrimSpace(listDisplay))
	var userAppointmentListInput int
	fmt.Scanln(&userAppointmentListInput)
	fmt.Println()
	return userAppointmentListInput
}

func userChoiceAction(userAppointmentListInput int, doctorList *[]doctorDetails, appointmentList *ClinicAppointmentList) {
	switch userAppointmentListInput {
	case 1:
		*doctorList = makeAppointment(&(*doctorList), appointmentList)
	case 2:
		displayAllDoctorAvailableTime(*doctorList)
	case 3:
		var doctorName string
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Doctor Name: ")
		doctorName, _ = reader.ReadString('\n')
		//convert CRLF to LF
		doctorName = strings.Replace(doctorName, "\n", "", -1)
		_ = searchDoctorByName(&(*doctorList), doctorName)
	case 4:
		displayAllBookedAppointments(appointmentList)

	case 5:
		deleteAppointment(appointmentList)
	case 6:
		fmt.Println("Bye Bye!")
		os.Exit(0)
	}
}

func displayAllDoctorAvailableTime(doctorList []doctorDetails) {
	fmt.Println(strings.Repeat("=", 50))
	for index, doctorValue := range doctorList {
		if doctorValue.available {
			fmt.Printf("%d) %s %v\n", index, doctorValue.doctorName, doctorValue.DayTime.Format(time.ANSIC))
		}
	}
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
	available := searchDoctorByName(&(*doctorList), tempDoctorName)

	if available {
		var slot int

		fmt.Print("Enter time slot for booking (eg.1 to book Timeslot 1): ")
		fmt.Scanf("%d", &slot)
		for _, doctorValue := range *doctorList {

			if strings.ToUpper(strings.TrimSpace(doctorValue.doctorName)) == tempDoctorName && doctorValue.available {
				var d1 = doctorDetails{
					drID:       doctorValue.drID,
					doctorName: doctorValue.doctorName,
					DayTime:    doctorValue.DayTime,
					available:  false}

				a1 := Appointment{
					aptID:       appointmentList.length + 1,
					patientName: tempPatientName,
					doctor:      d1,
				}
				appointmentList.Append(&a1)
				(*doctorList)[slot-1].available = false
				fmt.Printf("Timeslot %d booked", slot)
				fmt.Println("\n***Booking successful***")
				fmt.Printf("\tAppointment Id:- %d\n", appointmentList.length)
				fmt.Printf("\tAppointment booked for patient %s with doctor %s on %s\n", tempPatientName, tempDoctorName, doctorValue.DayTime.Format(time.ANSIC))
				break
			}
		}

	} else {
		fmt.Printf("%s not found or invalid timeslot selected\n", tempDoctorName)
	}
	return *doctorList
}

//sequential search
func searchDoctorByName(doctorList *[]doctorDetails, doctorName string) bool {
	var available bool = false
	var tempDoctorName = strings.ToUpper(strings.TrimSpace(doctorName))
	for index, doctorValue := range *doctorList {
		if strings.ToUpper(strings.TrimSpace(doctorValue.doctorName)) == tempDoctorName && doctorValue.available {

			fmt.Printf("%d) %s %v\n", index+1, doctorValue.doctorName, doctorValue.DayTime.Format(time.ANSIC))
			available = true

		}
	}
	if !available {
		fmt.Printf("%s not found\n", tempDoctorName)
		return available
	}

	return available
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

//Remove func can be exported
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

func deleteAppointment(ClinicAppointmentList *ClinicAppointmentList) {
	var aptID int
	fmt.Print("Enter Appointment Id to be delete: ")
	_, err := fmt.Scanf("%d", &aptID)
	if err != nil {
		log.Print("Scan for id failed, due to", err)
	} else {
		available := ClinicAppointmentList.Remove(aptID)
		if available {
			fmt.Printf("Appointment id %d deleted successfully!", aptID)
		} else {
			fmt.Printf("Appointment id %d not found!", aptID)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	title = "Doctor Appointment Booking Application"

	doctorListDisplay = `
1        :Lists all Avaialble Doctors
2        :Show all booked Appointment
3        :Make Appointment
4        :Search available Doctor by Name(eg.Name DR1)
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

func main() {
	t1, _ := time.Parse(time.RFC822, "17 FEB 21 10:00 SGT")
	t2, _ := time.Parse(time.RFC822, "18 FEB 21 10:00 SGT")

	d1 := doctorDetails{drID: 1, doctorName: "Dr1", DayTime: t1, available: true}
	d2 := doctorDetails{drID: 2, doctorName: "Dr2", DayTime: t2, available: true}
	d3 := doctorDetails{drID: 3, doctorName: "Dr3", DayTime: t1, available: true}

	doctorList := []doctorDetails{d1, d2, d3}

	appointmentList := &ClinicAppointmentList{}

	a1 := Appointment{
		aptID:       1,
		patientName: "p1",
		doctor:      d1,
	}

	appointmentList.Append(&a1)

	for {
		userAppointmentListInput := displayAppointmentList()
		userChoiceAction(userAppointmentListInput, &doctorList, appointmentList)
	}
}
func displayAppointmentList() int {
	fmt.Println()
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", 40))
	fmt.Print(strings.TrimSpace(doctorListDisplay))
	var userAppointmentListInput int
	fmt.Scanln(&userAppointmentListInput)
	return userAppointmentListInput
}

func userChoiceAction(userAppointmentListInput int, doctorList *[]doctorDetails, appointmentList *ClinicAppointmentList) {
	switch userAppointmentListInput {
	case 1:
		dispplayAllDoctorAvailableTime(*doctorList)
	case 2:
		displayAllBookedAppointments(appointmentList)
	case 3:
		*doctorList, _ = makeAppointment(&(*doctorList), appointmentList)
	case 4:
		var doctorName string
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Doctor Name: ")
		doctorName, _ = reader.ReadString('\n')
		//convert CRLF to LF
		doctorName = strings.Replace(doctorName, "\n", "", -1)
		searchDoctorByName(&(*doctorList), doctorName)
	case 5:
		//deleteAppointment(1)
	case 6:
		fmt.Println("Bye Bye!")
		os.Exit(0)
	}
}

func dispplayAllDoctorAvailableTime(doctorList []doctorDetails) {
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

func makeAppointment(doctorList *[]doctorDetails, appointmentList *ClinicAppointmentList) ([]doctorDetails, bool) {
	var available bool = false
	var patientName, doctorName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("******Appointment Booking System******")
	fmt.Print("Enter Patient Name: ")
	patientName, _ = reader.ReadString('\n')
	//conver CRLF to LF
	patientName = strings.Replace(patientName, "\n", "", -1)

	fmt.Print("Enter Doctor Name: ")
	doctorName, _ = reader.ReadString('\n')
	//conver CRLF to LF
	doctorName = strings.Replace(doctorName, "\n", "", -1)

	var tempDoctorName = strings.ToUpper(strings.TrimSpace(doctorName))
	var tempPatientName = strings.ToUpper(strings.TrimSpace(patientName))

	fmt.Printf("Here are available time slots for doctor %s\n", tempDoctorName)
	searchDoctorByName(&(*doctorList), doctorName)
	for index, doctorValue := range *doctorList {

		if strings.ToUpper(strings.TrimSpace(doctorValue.doctorName)) == tempDoctorName && doctorValue.available {
			t1, _ := time.Parse(time.RFC822, "16 FEB 21 10:00 SGT")

			var d1 = doctorDetails{
				drID:       doctorValue.drID,
				doctorName: doctorValue.doctorName,
				DayTime:    t1,
				available:  true}

			a1 := Appointment{
				aptID:       appointmentList.length + 1,
				patientName: tempPatientName,
				doctor:      d1,
			}
			available = true
			appointmentList.Append(&a1)

			(*doctorList)[index].available = false
			fmt.Println("\n***Booking successful***")
			fmt.Printf("\tAppointment Id:- %d\n", appointmentList.length)
			fmt.Printf("\tAppointment booked for patient %s with doctor %s on %s\n", tempPatientName, tempDoctorName, doctorValue.DayTime.Format(time.ANSIC))
			break
		}
	}

	if !available {
		fmt.Printf("%s not found", tempDoctorName)
	}
	return *doctorList, available
}

//sequential search
func searchDoctorByName(doctorList *[]doctorDetails, doctorName string) {
	var available bool = false
	var tempDoctorName = strings.ToUpper(strings.TrimSpace(doctorName))
	for index, doctorValue := range *doctorList {
		if strings.ToUpper(strings.TrimSpace(doctorValue.doctorName)) == tempDoctorName && doctorValue.available {
			fmt.Printf("%d) %s %v\n", index, doctorValue.doctorName, doctorValue.DayTime.Format(time.ANSIC))
			available = true
			break
		}
	}
	if !available {
		fmt.Printf("%s not found\n", tempDoctorName)
		return
	}

}

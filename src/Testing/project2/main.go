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

const (
	title = `
Welcome to "GO" Dental Clinic!
You are at "Appointment Booking Application"`

	listDisplay = `
1        :Make appointment
2        :Lists all available doctors
3        :Search available Doctor by Name(eg: DR1)
4        :Show all booked appointment
5        :Delete appointment(Admin only)
6        :Add Doctor schedule(Admin only)
7        :Exit
Enter your choice [1-6]: `
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

//Start of main program
func main() {

	//Defualt Time Data
	t1, _ := time.Parse(time.RFC822, "17 FEB 21 10:00 SGT")
	t2, _ := time.Parse(time.RFC822, "18 FEB 21 11:00 SGT")
	t3, _ := time.Parse(time.RFC822, "19 FEB 21 03:00 SGT")
	t4, _ := time.Parse(time.RFC822, "20 FEB 21 04:00 SGT")

	// This will be default data for doctor schedule at program startup.
	// Schedule can be entered by Admin using program option #6
	d1 := doctorDetails{drID: 1, doctorName: "DR1", DayTime: t1, available: true}
	d2 := doctorDetails{drID: 1, doctorName: "DR1", DayTime: t2, available: true}

	d3 := doctorDetails{drID: 2, doctorName: "DR2", DayTime: t3, available: true}
	d4 := doctorDetails{drID: 2, doctorName: "DR2", DayTime: t4, available: true}

	// This is Fist Data Structure (Slice) of the program for storing Doctor details and their available time slot
	// Schedule can be entered by Admin using program option #6
	doctorList := []doctorDetails{d1, d2, d3, d4}

	// This is Second Data Structure (Singly Linked List) of the program
	appointmentList := &ClinicAppointmentList{}

	//Infinite loop to show list options till user exits the program
	for {

		//Display list and accept user input
		userListInput := displayList()

		//Call appropriate function based on user input
		userChoiceAction(userListInput, &doctorList, appointmentList)
	}
}

// Function #1 :- Display user choice list and accept user input
func displayList() int {
	fmt.Println()
	fmt.Println(strings.TrimSpace(title))
	fmt.Println(strings.Repeat("=", 40))
	fmt.Print(strings.TrimSpace(listDisplay))
	var userListInput int
	fmt.Scanln(&userListInput)
	return userListInput
}

// Function#2 :- Call appropriate sub-function based on user Input
func userChoiceAction(userListInput int, doctorList *[]doctorDetails, appointmentList *ClinicAppointmentList) {
	ch := make(chan string)
	go readAdminPassword(ch)
	switch userListInput {
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
		readPassword := <-ch
		enteredPassword := enterAdminPassword()
		if enteredPassword == readPassword {
			fmt.Println("Success")
			deleteAppointment(appointmentList)
		} else {
			fmt.Println("Incorrect password")
		}

	case 6:
		readPassword := <-ch
		enteredPassword := enterAdminPassword()
		if enteredPassword == readPassword {
			fmt.Println("Success")
			displayAllDoctorAvailableTime(*doctorList)
			creatDoctorSchedule(&(*doctorList))
		} else {
			fmt.Println("Incorrect password")
		}

	case 7:
		fmt.Println("Bye! Bye!")
		os.Exit(0)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

//function to display available doctors time slot to book for appointment.
//Doctor deatils is avaialble in slice, loop will range over to find all available doctors to display
func displayAllDoctorAvailableTime(doctorList []doctorDetails) {
	fmt.Println(strings.Repeat("=", 50))
	for index, doctorValue := range doctorList {
		if doctorValue.available {
			fmt.Printf("%d) %s %s %d\n", index+1, doctorValue.doctorName, doctorValue.DayTime.Format(time.ANSIC), doctorValue.appointmentID)
		}
	}
}

//linked list of booked appointment list will get display by this function, all the current book appointment cab be seen
func displayAllBookedAppointments(appointmentList *ClinicAppointmentList) {
	if appointmentList.length == 0 {
		fmt.Println("Appointment list is Empty")
		return
	}
	index := 1
	appointmentListValue := appointmentList.start
	fmt.Println("List of booked Appointments:")
	for appointmentListValue != nil {
		fmt.Printf("%d) %s %s %v\n", appointmentListValue.aptID, appointmentListValue.patientName, appointmentListValue.doctor.doctorName, appointmentListValue.doctor.DayTime.Format(time.ANSIC))
		index = index + 1
		appointmentListValue = appointmentListValue.next
	}
}

//function will help user to make appointment once time slot of selected doctor book.
//Available status of selected doctor become false & booked appointment ca n be seen in #4
//user get prompted based on availability of slot
func makeAppointment(doctorList *[]doctorDetails, appointmentList *ClinicAppointmentList) []doctorDetails {
	var patientName, doctorName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println("******Appointment Booking System******")
	fmt.Print("Enter Patient Name: ")
	patientName, _ = reader.ReadString('\n')

	//conver CRLF to LF
	patientName = strings.Replace(patientName, "\n", "", -1)

	fmt.Print("Enter Doctor Name(eg. DR1): ")
	doctorName, _ = reader.ReadString('\n')
	//conver CRLF to LF
	doctorName = strings.Replace(doctorName, "\n", "", -1)

	var tempDoctorName = strings.ToUpper(strings.TrimSpace(doctorName))
	var tempPatientName = strings.ToUpper(strings.TrimSpace(patientName))
	fmt.Println()
	fmt.Printf("Here are available time slots for doctor %s\n", tempDoctorName)
	available := searchDoctorByName(&(*doctorList), tempDoctorName)

	if available {
		var slot int
		fmt.Print("Enter time slot for booking (eg.1 to book Timeslot 1): ")
		fmt.Scanln(&slot)
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

//sequential search algorithm is used to search available docotor selected by user
//user will get prompted based on availability of doctor
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

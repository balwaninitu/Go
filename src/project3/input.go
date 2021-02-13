package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

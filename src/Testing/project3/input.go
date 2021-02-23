package main

import (
	"errors"
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
	_, err := fmt.Scanln(&userListInput)
	if err != nil {
		fmt.Println(errors.New("Error:Unexpected new line"))
	}
	return userListInput
}

// Function#2 :- Call appropriate sub-function based on user Input
func userChoiceAction(userListInput int, doctorList *[]doctorDetails, appointmentList *ClinicAppointmentList) {
	ch := make(chan string)
	go readAdminPassword(ch) //go routine created to receive and transfer password securely through channel for admin only features
	switch userListInput {
	case 1:
		*doctorList = makeAppointment(&(*doctorList), appointmentList)
	case 2:
		displayAllDoctorAvailableTime(*doctorList)
	case 3:
		var doctorName string

		fmt.Print("Enter Doctor Name: ")
		_, err := fmt.Scanln(&doctorName)
		if err != nil {
			fmt.Println(errors.New("Error:Unexpected new line"))
		}

		_ = searchDoctorByName(&(*doctorList), doctorName)
	case 4:
		displayAllBookedAppointments(appointmentList)
		//admin features only accessible after entering password which is concurrentltly running through goroutines channel
		//password will be send by channel
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

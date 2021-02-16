package main

import (
	"errors"
	"fmt"
	"os"
)

//  Call appropriate sub-function based on user Input
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
			fmt.Println(errors.New("Error:No input"))
		}

		_ = searchDoctorByName(&(*doctorList), doctorName)
	case 4:
		displayAllBookedAppointments(appointmentList)
		//admin features only accessible after entering password which is concurrentltly running through goroutines channel
		//password will be send by channel
	case 5: //reading password from channel and entering to access admin feature
		readPassword := <-ch
		enteredPassword := enterAdminPassword()
		if enteredPassword == readPassword {
			fmt.Println("Success")
			deleteAppointment(appointmentList)
		} else {
			fmt.Println("Incorrect password")
		}

	case 6: //reading password from channel and entering to access admin feature
		readPassword := <-ch
		enteredPassword := enterAdminPassword()
		if enteredPassword == readPassword {
			fmt.Println("Success")
			displayAllDoctorAvailableTime(*doctorList)
			creatDoctorSchedule(&(*doctorList))
		} else {
			fmt.Println("Incorrect password")
		}

	case 7: //exit from infinite loop
		fmt.Println("Bye! Bye!")
		os.Exit(0)
	}
}

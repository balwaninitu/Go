package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func displayList() int {
	fmt.Println()
	fmt.Println(strings.TrimSpace(title))
	fmt.Println(strings.Repeat("=", 40))
	fmt.Print(strings.TrimSpace(listDisplay))
	var userListInput int
	fmt.Scanln(&userListInput)
	fmt.Println()
	return userListInput
}

func userChoiceAction(userListInput int, doctorList *[]doctorDetails, appointmentList *ClinicAppointmentList) {
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
		deleteAppointment(appointmentList)
	case 6:
		fmt.Println("Bye Bye!")
		os.Exit(0)
	}
}

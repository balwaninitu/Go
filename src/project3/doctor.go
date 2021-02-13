package main

import (
	"fmt"
	"strings"
	"time"
)

//sequential search algorithm is used to search available docotor selected by user
//user will get prompted based on availability of doctor
func searchDoctorByName(doctorList *[]doctorDetails, doctorName string) bool {
	var available bool = false
	var tempDoctorName = strings.ToUpper(strings.TrimSpace(doctorName))
	fmt.Println()
	for index, doctorValue := range *doctorList {
		if strings.ToUpper(strings.TrimSpace(doctorValue.doctorName)) == tempDoctorName && doctorValue.available {
			fmt.Printf("%d) %s %v\n", index+1, doctorValue.doctorName, doctorValue.DayTime.Format(time.ANSIC))
			available = true
		}
	}
	if !available {
		fmt.Printf("Docotr %s not found\n", tempDoctorName)
		return available
	}

	return available
}

//function to display available doctors time slot to book for appointment.
//Doctor deatils is avaialble in slice, loop will range over to find all available doctors to display
func displayAllDoctorAvailableTime(doctorList []doctorDetails) {
	fmt.Println()
	fmt.Println("***List of available docotrs***")
	fmt.Println(strings.Repeat("=", 50))
	for index, doctorValue := range doctorList {
		if doctorValue.available {
			fmt.Printf("%d) %s %s %d\n", index+1, doctorValue.doctorName, doctorValue.DayTime.Format(time.ANSIC), doctorValue.appointmentID)
		}
	}
}

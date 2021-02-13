package main

import (
	"fmt"
	"strings"
	"time"
)

func displayAllDoctorAvailableTime(doctorList []doctorDetails) {
	fmt.Println(strings.Repeat("=", 50))
	for index, doctorValue := range doctorList {
		if doctorValue.available {
			fmt.Printf("%d) %s %v\n", index, doctorValue.doctorName, doctorValue.DayTime.Format(time.ANSIC))
		}
	}
}

//sequential search
func searchDoctorByName(doctorList *[]doctorDetails, doctorName string) bool {
	var available bool = false
	var tempDoctorName = strings.ToUpper(strings.TrimSpace(doctorName))

	for index, doctorValue := range *doctorList {
		if strings.ToUpper(strings.TrimSpace(doctorValue.doctorName)) == tempDoctorName && doctorValue.available {

			fmt.Printf("Id:%d) %s %v\n", index+1, doctorValue.doctorName, doctorValue.DayTime.Format(time.ANSIC))
			available = true

		}
	}
	if !available {
		fmt.Printf("%s not found\n", tempDoctorName)
		return false
	}

	return available
}

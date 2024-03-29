package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	fmt.Println(strings.Repeat("=", 40))
	for index, doctorValue := range doctorList {
		if doctorValue.available {
			fmt.Printf("%d) %s %s %d\n", index+1, doctorValue.doctorName, doctorValue.DayTime.Format(time.ANSIC), doctorValue.appointmentID)
		}
	}
}

//ReadCsv can be exported
//return docotr details slice
func ReadCsv(filename string) ([][]string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

// ReadDoctorList can be exported
func readDoctorList(doctorList *[]doctorDetails) []doctorDetails {
	defer recoverFromPanic()
	var lines, err = ReadCsv("C:/Projects/Go/src/Assignments/goschool-assignment-2/Dental-Appointment-Syastem/doctor.csv")
	if err != nil {
		panic(errors.New("Wrong file name or path"))
	} //data in slice avaialble in string
	for _, line := range lines {
		drID, err := strconv.Atoi(line[0]) //convert string to int
		check(err)
		doctorName := line[1]
		DayTime, err := time.Parse(time.RFC822, line[2]) //string to time format
		check(err)
		available, err := strconv.ParseBool(line[3]) //string to bool
		check(err)
		data := doctorDetails{
			drID:       drID,
			doctorName: doctorName,
			DayTime:    DayTime,
			available:  available,
		}

		*doctorList = append(*doctorList, data)
	}
	return *doctorList
}

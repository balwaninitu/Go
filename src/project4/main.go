package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type doctorDetails struct {
	drID       int
	doctorName string
	DayTime    time.Time
	available  bool
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//ReadCsv can be exported
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

func main() {

	doctorList := []doctorDetails{}

	displayAllDoctorAvailableTime(doctorList)

	err := ReadDoctorList(doctorList)

	check(err)

}

func displayAllDoctorAvailableTime(doctorList []doctorDetails) {
	fmt.Println()
	fmt.Println("***List of available docotrs***")
	fmt.Println(strings.Repeat("=", 50))
	for index, doctorValue := range doctorList {
		if doctorValue.available {
			fmt.Printf("%d) %s %s\n", index+1, doctorValue.doctorName, doctorValue.DayTime.Format(time.ANSIC))
		}
	}
}

// ReadDoctorList can be exported
func ReadDoctorList(doctorList []doctorDetails) error {
	var lines, err = ReadCsv("C:/Projects/Go/src/project4/doctor.csv")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		drID, err := strconv.Atoi(line[0])
		check(err)
		doctorName := line[1]
		DayTime, err := time.Parse(time.RFC822, line[2])
		check(err)
		available, err := strconv.ParseBool(line[3])
		check(err)
		data := doctorDetails{
			drID:       drID,
			doctorName: doctorName,
			DayTime:    DayTime,
			available:  available,
		}

		doctorList = append(doctorList, data)
		fmt.Println(data.drID, data.doctorName, data.DayTime, data.available)
	}
	return nil
}

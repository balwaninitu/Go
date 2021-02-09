package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type doctorDetails struct {
	id            int
	name          string
	availableTime time.Time
}

func main() {

	doctors := []doctorDetails{
		{id: 1, name: "dr1", availableTime: time.Date(2021, time.February, time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Local)},
		{id: 2, name: "dr2", availableTime: time.Date(2021, time.February, time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Local)},
		{id: 3, name: "dr3", availableTime: time.Date(2021, time.February, time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Local)},
		{id: 4, name: "dr4", availableTime: time.Date(2021, time.February, time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Local)},
		{id: 5, name: "dr5", availableTime: time.Date(2021, time.February, time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Local)},
	}

	//fmt.Println(doctors)

	//index the doctor by id

	byID := make(map[int]doctorDetails)

	for _, d := range doctors {
		byID[d.id] = d
	}

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
>list         : lists all the doctors
>book         : make Appointment
>search       : search available doctor by name
>delete       : delete Appointment(admin only)
>exit         : Exit
`)
		if !in.Scan() {
			break
		}
		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}
		switch cmd[0] {
		case "exit":
			fmt.Println("bye bye!")
			return

		case "list":
			for _, d := range doctors {
				fmt.Printf("%d. %-10q  Available: %-10q\n", d.id, d.name, d.availableTime)
			}

		case "search":
			var name string
			fmt.Println("enter name of doctor")
			fmt.Scanln(&name)
			search(doctors, len(doctors), name)

		case "book":

			myAppointment := createAppointmentList("myAppointment")
			var patientName string
			var doctorID int
			var aptTime time.Location
			//var availableDay string
			fmt.Println("enter patient name")
			fmt.Scanln(&patientName)
			fmt.Println("find below details of doctor and their available day")
			for _, d := range doctors {
				fmt.Printf("%d. %-10q  Available: %-10q\n", d.id, d.name, d.availableTime)
			}
			fmt.Println("enter doctor id")
			fmt.Scanln(&doctorID)
			d, ok := byID[doctorID]
			if !ok {
				fmt.Println("Sorry, We dont have the Doctor with this ID")

			}
			fmt.Printf("%d is %s and is available on %v", doctorID, d.name, d.availableTime)
			myAppointment.addAppointmentDetails(patientName, d.name, doctorID, aptTime)
			fmt.Println("Appointment created")
			myAppointment.showAllDetails()

		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type details struct {
	id           int
	name         string
	availableDay string
}

func main() {
	doctors := []details{
		details{id: 1, name: "dr1", availableDay: "Mon"},
		details{id: 2, name: "dr2", availableDay: "Tue"},
		details{id: 3, name: "dr3", availableDay: "Wed"},
		details{id: 4, name: "dr4", availableDay: "Thu"},
		details{id: 5, name: "dr5", availableDay: "Thu"},
	}

	//fmt.Println(doctors)

	//index the doctor by id

	byID := make(map[int]details)

	for _, d := range doctors {
		byID[d.id] = d
	}

	fmt.Printf("Dental clinic has %d doctors.\n", len(doctors))
	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
>list         : lists all the doctors
>book         : make Appointment(eg. )
>id N         : search available doctor by ID(eg. id 1)
>delete       :delete Appointment(admin only)
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

		case "book":
			myAppointment := createAppointment("myAppointment")
			var name string
			fmt.Println("enter patient name")
			fmt.Scanln(&name)
			fmt.Printf("Hi %s Available doctors are\n", name)
			for _, d := range doctors {
				fmt.Printf("Id :%d  Name: %s on %s\n", d.id, d.name, d.availableDay)
			}
			fmt.Println("enter docter id")
			fmt.Scanln(&drID)
			d, ok := byID[drID]
			if !ok {
				fmt.Println("Sorry, We dont have the Doctor with this ID")
				continue
			} else {
				fmt.Printf("Id:%d  Name: %-10q  Available: %-10q  Specialisation: %-10q\n", d.id, d.name, d.availableDay, d.specialisation)
			}
			fmt.Printf("%s is available on: %s\n", d.name, d.availableDay)
			fmt.Println("Please proceed pressing enter")
			fmt.Printf("Appointment 1: %s has appointment with %d on %s", myAppointment.name, myAppointment.now.doctorID, myAppointment.now.dayOfAppointment)
			// var patientName, aptDay string
			// var doctorID int
			// var d doctor
			// fmt.Println("Enter your name")
			// fmt.Scanln(&patientName)
			// fmt.Printf("Available doctors are\n")
			// for _, d := range doctors {
			// 	fmt.Printf("%d. Name: %s Specialisation: %s\n", d.id, d.name, d.specialisation)
			// }
			// fmt.Println("Enter doctor Id")
			// fmt.Scanln(&doctorID)
			// d, ok := byID[doctorID]
			// if !ok {
			// 	fmt.Println("Sorry, We dont have the Doctor with this ID")
			// 	continue
			// } else {
			// 	fmt.Printf("%d. Name: %-10q  Available: %-10q  Specialisation: %-10q\n", d.id, d.name, d.availableDay, d.specialisation)
			// }
			// fmt.Printf("%s is available on: %s\n", d.name, d.availableDay)
			// fmt.Println("Please proceed by entering given day")
			// fmt.Scanln(&aptDay)

			// myAppointment := createAppointment("myAppointment")
			// //fmt.Println("Created appointmentList")

			// myAppointment.addAppointmentDetails(patientName, doctorID, aptDay)
			// created := myAppointment.Head()
			// fmt.Printf("Appointment for %s has booked with Doctor: %d on %s\n", created.patientName, created.doctorID, created.dayOfAppointment)

			// fmt.Printf("Thank you %s!\nYour Appointment has been added\n", patientName)
			// //fmt.Println(myAppointment.showAllAppointment())

		case "edit":
			var created *appointment
			var name string
			fmt.Println("enter your name")
			fmt.Scanln(&name)
			if name == created.patientName {
				fmt.Printf("Appointment for %s has booked with Doctor: %d on %s\n", created.patientName, created.doctorID, created.dayOfAppointment)

			} else {
				fmt.Println("Sorry!we dont have apt with this name")
			}

		}
	}
}

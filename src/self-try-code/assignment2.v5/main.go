package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	doctors := []details{
		{id: 1, name: "dr1", availableDay: "Mon"},
		{id: 2, name: "dr2", availableDay: "Tue"},
		{id: 3, name: "dr3", availableDay: "Wed"},
		{id: 4, name: "dr4", availableDay: "Thu"},
		{id: 5, name: "dr5", availableDay: "Fri"},
	}

	//fmt.Println(doctors)

	//index the doctor by id

	byID := make(map[int]details)

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
				fmt.Printf("%d. %-10q  Available: %-10q\n", d.id, d.name, d.availableDay)
			}

		case "search":
			var name string
			fmt.Println("enter name of doctor")
			fmt.Scanln(&name)
			search(doctors, len(doctors), name)

		case "book":

			myAppointment := createAppointmentList("myAppointment")

			myAppointment.addDetails(1, "Dr1", "Mon")
			myAppointment.addDetails(2, "Dr2", "Tue")
			myAppointment.addDetails(3, "Dr3", "Wed")
			myAppointment.addDetails(4, "Dr4", "Thu")
			myAppointment.addDetails(5, "Dr5", "Fri")
			// make appointment
			var patientName string
			var id int
			//var anotherID int
			fmt.Println("enter name")
			fmt.Scanln(&patientName)
			fmt.Println("find below details of doctor and their available day")
			myAppointment.showAllDetails()
			fmt.Println()
			fmt.Println("enter doctor id to book appointment")
			fmt.Scanln(&id)
			myAppointment.startBooking(patientName, id)
			d, ok := byID[id]
			if !ok {
				fmt.Println("Sorry, We dont have the Doctor with this ID")
				continue
			} else {
				fmt.Printf("%s, your appointment with Doctor, ID : %d has confirmed on %s\n", patientName, d.id, d.availableDay)
				fmt.Println()
				// fmt.Println("do you want to edit appointment?")
				// fmt.Println("if no, you can continue to use other services")
				// fmt.Println("if yes, press enter")

				// fmt.Scanln(&anotherID)
				// if anotherID == id {
				// 	fmt.Println("you have already confirmed appointment with same doctor")
				// 	return
				// } //clinic allows max. two appointments at one time
				// myAppointment.nextBooking(anotherID)
				// d, ok := byID[anotherID]
				// if !ok {
				// 	fmt.Println("Sorry, We dont have the Doctor with this ID")
				// 	continue
				// } else {
				// 	fmt.Printf("%s, your appointment with Doctor, ID : %d has confirmed on %s\n", patientName, d.id, d.availableDay)

			}

		}

	}
}

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
		case "book":

			myAppointment := createAppointmentList("myAppointment")
			var name string
			var drID int
			fmt.Println("enter patient name")
			fmt.Scanln(&name)
			fmt.Printf("Hi %s Available doctors are\n", name)
			for _, d := range doctors {
				fmt.Printf("Id :%d  Name: %s on %s\n", d.id, d.name, d.availableDay)
			}
			fmt.Println("enter docter id to book appointment(eg.1)")
			fmt.Scanln(&drID)
			d, ok := byID[drID]
			if !ok {
				fmt.Println("Sorry, We dont have the Doctor with this ID")
				continue
			} else {
				fmt.Printf("Your appointment with %s is confirmed on %s\n", d.name, d.availableDay)
			}
			myAppointment.addDetails(d.id, name, d.availableDay)
			//myAppointment.showAllDetails()
		}
	}
}

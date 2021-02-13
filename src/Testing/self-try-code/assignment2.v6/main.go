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
	availability  bool
}

func main() {

	t1, _ := time.Parse(time.RFC822, "17 FEB 21 10:00 SGT")
	t2, _ := time.Parse(time.RFC822, "18 FEB 21 12:00 SGT")
	t3, _ := time.Parse(time.RFC822, "17 FEB 21 14:00 SGT")
	t4, _ := time.Parse(time.RFC822, "18 FEB 21 16:00 SGT")
	t5, _ := time.Parse(time.RFC822, "19 FEB 21 18:00 SGT")

	doctors := []doctorDetails{
		{id: 1, name: "dr1", availableTime: t1, availability: true},
		{id: 2, name: "dr2", availableTime: t2, availability: true},
		{id: 3, name: "dr3", availableTime: t3, availability: true},
		{id: 4, name: "dr4", availableTime: t4, availability: true},
		{id: 5, name: "dr5", availableTime: t5, availability: true},
	}

	//fmt.Println(doctors)

	//index the doctor by id

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
				fmt.Printf("%d. %-10q  Available On: %-10q Availability:%t\n", d.id, d.name, d.availableTime.Format(time.ANSIC), d.availability)
			}

		case "search":
			var name string
			fmt.Println("enter name of doctor")
			fmt.Scanln(&name)
			search(doctors, len(doctors), name)

		case "book":

			myAppointment := createAppointmentList("myAppointment")
			var patientName string
			var doctorName string
			//var aptTime time.Location
			var i int
			var d doctorDetails
			//var availableDay string
			fmt.Println("enter patient name")
			fmt.Scanln(&patientName)
			fmt.Println("find below details of doctor and their availability")
			for i, d = range doctors {
				i = i + 1
				fmt.Printf("%d. %-10q  Available on: %-10q, Available:%t\n", d.id, d.name, d.availableTime.Format(time.ANSIC), d.availability)
			}
			fmt.Println("choose doctor name from above list")
			fmt.Scanln(&doctorName)
			err := search(doctors, len(doctors), doctorName) //sequential search
			if err != nil {
				fmt.Println("Invalid name")
				continue
			}
			fmt.Println("press eneter to make appointment")
			fmt.Scanln()
			myAppointment.addAppointmentDetails(patientName, doctorName)
			fmt.Println()
			myAppointment.showAllDetails()
			continue

		}

	}
}

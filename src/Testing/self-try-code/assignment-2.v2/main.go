package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type details struct {
	id           int
	name         string
	availableDay string
}
type doctor struct {
	details
	specialisation string
}

func main() {
	doctors := []doctor{
		{details: details{id: 1, name: "dr1", availableDay: "Mon"}, specialisation: "Periodontist"},
		{details: details{id: 2, name: "dr2", availableDay: "Tue"}, specialisation: "Endodontist"},
		{details: details{id: 3, name: "dr3", availableDay: "Wed"}, specialisation: "Orthodontist"},
		{details: details{id: 4, name: "dr4", availableDay: "Thu"}, specialisation: "General Dentist"},
		{details: details{id: 5, name: "dr5", availableDay: "Thu"}, specialisation: "Prosthodontist"},
	}

	//fmt.Println(doctors)

	//index the doctor by id

	byID := make(map[int]doctor)

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
		case "exit":
			fmt.Println("bye bye!")
			return

		case "list":
			for _, d := range doctors {
				fmt.Printf("%d. Name: %-10q  Available: %-10q  Specialisation: %-10q\n", d.id, d.name, d.availableDay, d.specialisation)
			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			// d, ok := byID[id]
			// if !ok {
			// 	fmt.Println("Sorry, We dont have the Doctor with this ID")
			// 	continue
			// } else {
			// 	fmt.Printf("%d. Name: %-10q  Available: %-10q  Specialisation: %-10q\n", d.id, d.name, d.availableDay, d.specialisation)
			// }
			searchDrSlice := appendSlice(doctors)
			search(searchDrSlice, len(searchDrSlice), id)
			//fmt.Printf("%d. Name: %-10q  Available: %-10q  Specialisation: %-10q\n", d.id, d.name, d.availableDay, d.specialisation)

		case "delete":
			//only admin can delete appointment based on dequeue algorithm FIFO
			fmt.Println("sorry you dont have access to this feature")

			// lock this section by password

		case "book":
			var patientName, aptDay string
			var doctorID int
			var d doctor
			fmt.Println("Enter your name")
			fmt.Scanln(&patientName)
			fmt.Printf("Available doctors are\n")
			for _, d := range doctors {
				fmt.Printf("%d. Name: %s Specialisation: %s\n", d.id, d.name, d.specialisation)
			}
			fmt.Println("Enter doctor Id (eg. 1)")
			fmt.Scanln(&doctorID)
			d, ok := byID[doctorID]
			if !ok {
				fmt.Println("Sorry, We dont have the Doctor with this ID")
				continue
			} else {
				fmt.Printf("%d. Name: %-10q  Available: %-10q  Specialisation: %-10q\n", d.id, d.name, d.availableDay, d.specialisation)
			}

			fmt.Printf("%s is available on: %s\n", d.name, d.availableDay)
			fmt.Println("Please proceed by entering given day")
			fmt.Scanln(&aptDay)

			myAppointment := createAppointment("myAppointment")
			fmt.Println("Created appointmentList")

			myAppointment.addAppointmentDetails(patientName, doctorID, aptDay)
			fmt.Printf("Thank you %s!\nYour Appointment has been added", patientName)
			fmt.Println(myAppointment.showAllAppointment())

		}

	}
}

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
		{details: details{id: 2, name: "dr1", availableDay: "Tue"}, specialisation: "Endodontist"},
		{details: details{id: 3, name: "dr1", availableDay: "Wed"}, specialisation: "Orthodontist"},
		{details: details{id: 4, name: "dr1", availableDay: "Thu"}, specialisation: "General Dentist"},
		{details: details{id: 5, name: "dr1", availableDay: "Thu"}, specialisation: "Prosthodontist"},
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
>appointment  : make Appointment
>id N         : search available doctor by ID(eg. id 1)
>edit         : edit Appointment
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
		case "quit":
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

			d, ok := byID[id]
			if !ok {
				fmt.Println("Sorry, We dont have the Doctor with this ID")
				continue
			}
			fmt.Printf("%d. Name: %-10q  Available: %-10q  Specialisation: %-10q\n", d.id, d.name, d.availableDay, d.specialisation)

		}
	}
}

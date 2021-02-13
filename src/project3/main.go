package main

import "time"

const (
	title = `
Welcome to "GO" Dental Clinic!
You are at "Appointment Booking Application"`

	listDisplay = `
1        :Make appointment
2        :Lists all available doctors
3        :Search available Doctor by Name(eg: DR1)
4        :Show all booked appointment
5        :Delete appointment(Admin only)
6        :Add Doctor schedule(Admin only)
7        :Exit
Enter your choice [1-6]: `
)

//Start of main program
func main() {

	//Defualt Time Data
	t1, _ := time.Parse(time.RFC822, "17 FEB 21 10:00 SGT")
	t2, _ := time.Parse(time.RFC822, "18 FEB 21 11:00 SGT")
	t3, _ := time.Parse(time.RFC822, "19 FEB 21 03:00 SGT")
	t4, _ := time.Parse(time.RFC822, "20 FEB 21 04:00 SGT")

	// This will be default data for doctor schedule at program startup.
	// Schedule can be entered by Admin using program option #6
	d1 := doctorDetails{drID: 1, doctorName: "DR1", DayTime: t1, available: true}
	d2 := doctorDetails{drID: 1, doctorName: "DR1", DayTime: t2, available: true}

	d3 := doctorDetails{drID: 2, doctorName: "DR2", DayTime: t3, available: true}
	d4 := doctorDetails{drID: 2, doctorName: "DR2", DayTime: t4, available: true}

	// This is Fist Data Structure (Slice) of the program for storing Doctor details and their available time slot
	// Schedule can be entered by Admin using program option #6
	doctorList := []doctorDetails{d1, d2, d3, d4}

	// This is Second Data Structure (Singly Linked List) of the program
	appointmentList := &ClinicAppointmentList{}

	//Infinite loop to show list options till user exits the program
	for {

		//Display list and accept user input
		userListInput := displayList()

		//Call appropriate function based on user input
		userChoiceAction(userListInput, &doctorList, appointmentList)
	}
}

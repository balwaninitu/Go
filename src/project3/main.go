package main

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
Enter your choice [1-7]: `
)

//Start of main program
func main() {

	// This is Fist Data Structure (Slice) of the program for storing Doctor details and their available time slot
	doctorList := []doctorDetails{}
	//doctor details will get read from csv file
	//function details available in doctor.go file
	doctorList = readDoctorList(&doctorList)
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

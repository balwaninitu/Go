package main

import (
	"errors"
	"fmt"
	"log"
)

//Remove func can be exported
func (c *ClinicAppointmentList) Remove(aptID int) bool {

	// Appointment list is empty
	if c.length == 0 {
		panic(errors.New("Appointment list is empty"))
	}

	available := false
	//To delete first appointment - move head pointer
	currentAppointment := c.start
	if c.start.aptID == aptID {
		c.start = currentAppointment.next
		c.length--
		available = true

	} else {

		//To delete middle appointment - need two pointers (previous and next)
		previousAppointment := c.start
		currentAppointment = c.start.next
		for currentAppointment.next != nil {
			if currentAppointment.aptID == aptID {
				previousAppointment.next = currentAppointment.next
				c.length--
				available = true

			}
			currentAppointment = currentAppointment.next
			previousAppointment = previousAppointment.next
		}

		//To delete last appointment need to confirm if node is nil or not
		if currentAppointment.next == nil && currentAppointment.aptID == aptID {
			previousAppointment.next = nil
			c.length--
		}
	}
	return available
}

func deleteAppointment(ClinicAppointmentList *ClinicAppointmentList) {
	var aptID int
	fmt.Print("Enter Appointment Id to be delete: ")
	_, err := fmt.Scanf("%d", &aptID)
	if err != nil {
		log.Print("Scan for id failed, due to", err)
	} else {
		available := ClinicAppointmentList.Remove(aptID)
		if available {
			fmt.Printf("Appointment id %d deleted successfully!", aptID)
		} else {
			fmt.Printf("Appointment id %d not found!", aptID)
		}
	}
}

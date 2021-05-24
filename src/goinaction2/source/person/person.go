// Package Person declares the Person struct,vcontaining user information such as username, identification, password, first and last name, date of birth, phone, address, email, and vaccination information. It contains all methods that manipulate it.
package person

import (
	"fmt"
	"strings"
	"time"
)

const (
	//Time Format constant
	layoutDateTime = "2006-01-02 15:04:05"

	// number of available appointments shown at a time
	numOfApptsShown = 15
)

// Person data structure to store all data related to the user/person
type Person struct {
	Identification     string
	Username           string
	Password           string
	Firstname          string
	Lastname           string
	Dob                string
	Phone              string
	Address            string
	Email              string
	VaccinationQualify bool
	FirstVaccineDone   bool
	FirstVaccineDate   string
	FirstVaccineTime   string
	SecondVaccineDone  bool
	SecondVaccineDate  string
	SecondVaccineTime  string
}

// This is a method for Person struct. It returns all the appointments information for the Person as a slice of string.
// It gets the first and second vaccination date, time and qualification information stored in Person struct.
// Based on current time, and vaccination criteria, the appointment related message is created and returned.
func (p *Person) PrintAllAppt() ([]string, error) {
	message := []string{}
	message = append(message, fmt.Sprintf("Appointment information for %s %s\n", p.Firstname, p.Lastname))
	timeFirst, ferr := datePlusTime(p.FirstVaccineDate, p.FirstVaccineTime)
	timeSecond, serr := datePlusTime(p.SecondVaccineDate, p.SecondVaccineTime)
	if ferr != nil {
		return message, ferr
	}
	if serr != nil {
		return message, serr
	}

	timeFirst29 := timeFirst.Add(time.Hour * 24 * time.Duration(29))
	timeNow := time.Now()

	if timeFirst.IsZero() {
		message = append(message, fmt.Sprintf("You have no appointments yet. Please make new appointments."))
	} else {
		if timeFirst.After(time.Now()) {
			message = append(message, fmt.Sprintf("First Vaccine Appointment is on: %s", timeFirst.Format(layoutDateTime)))
		} else {
			if p.FirstVaccineDone {
				message = append(message, fmt.Sprintf("First Vaccine was given on: %s", timeFirst.Format(layoutDateTime)))

			} else {
				message = append(message, fmt.Sprintf("You missed your first Vaccine appointment dated: %s", timeFirst.Format(layoutDateTime)))
				message = append(message, fmt.Sprintf("Please delete first appointment and make new appointments"))
			}
		}

		if timeSecond.IsZero() {
			if p.FirstVaccineDone {
				//check if first vaccine is  done more than 29 days before now
				if timeNow.After(timeFirst29) {
					message = append(message, fmt.Sprintf("28 days have already elapsed after your first vaccination dose."))
					message = append(message, fmt.Sprintf("You can no longer make a second vaccination appointment. Contact MOH immediately!"))
					return message, nil
				}
			} else {
				message = append(message, fmt.Sprintf("You have no appointment for second vaccine. Please make new appointment."))
			}
		} else {
			if timeSecond.After(timeNow) {
				message = append(message, fmt.Sprintf("Second Vaccine Appointment is on: %s", timeSecond.Format(layoutDateTime)))
			} else {
				if p.SecondVaccineDone {
					message = append(message, fmt.Sprintf("Second Vaccine was given on: %s", timeSecond.Format(layoutDateTime)))
				} else {
					message = append(message, fmt.Sprintf("You missed your second Vaccine appointment dated: %s", timeSecond.Format(layoutDateTime)))
					if p.FirstVaccineDone {
						if timeNow.After(timeFirst29) {
							message = append(message, fmt.Sprintf("28 days have already elapsed after your first vaccination dose."))
							message = append(message, fmt.Sprintf("You can no longer make a second vaccination appointment. Contact MOH immediately!"))
						} else {
							message = append(message, fmt.Sprintf("You can still make a new appointment for second vaccination."))
							message = append(message, fmt.Sprintf("Delete second appointment before making new appointment"))
							message = append(message, fmt.Sprintf("Hurry up! Slots are limited. Or contact MOH immediately!"))
						}
					}
				}
			}
		}
	}
	return message, nil
}

// This is a method for Person struct. It returns appointment information for template makeappt.gohtml.
// It returns first and last name information, message to be displayed, appointment type, and a list of possible appointments.
// It uses apptArray to extract information about available appointments.
func (p *Person) MakeNewAppt(apptArray []time.Time) (err error, fn string, ln string, msg []string, aptI string, posApts []string) {

	var firstname string
	var lastname string
	var message []string
	var apptinfo string
	var possibleAppts []string

	firstname = p.Firstname
	lastname = p.Lastname

	timeFirst, ferr := datePlusTime(p.FirstVaccineDate, p.FirstVaccineTime)
	timeSecond, serr := datePlusTime(p.SecondVaccineDate, p.SecondVaccineTime)
	if ferr != nil {
		return ferr, firstname, lastname, message, apptinfo, possibleAppts
	}
	if serr != nil {
		return serr, firstname, lastname, message, apptinfo, possibleAppts
	}

	timeFirst20 := timeFirst.Add(time.Hour * 24 * time.Duration(20))
	timeFirst29 := timeFirst.Add(time.Hour * 24 * time.Duration(29))
	timeNow := time.Now()
	// Apptinfo : none = no appt being made, first = first appt being made, second = second appt being made
	apptinfo = "none"

	if p.FirstVaccineDone {
		if p.SecondVaccineDone {
			message = append(message, fmt.Sprintf("Your COVID-19 Vaccination is complete. Take good rest and be safe."))
			return nil, firstname, lastname, message, apptinfo, possibleAppts
		} else {
			// new appt for second vaccine. Must be 21-28 days within first vaccine
			// else contact MOH hotline
			if timeSecond.IsZero() {
				message = append(message, fmt.Sprintf("Your first vaccination was done on %s", timeFirst.Format(layoutDateTime)))
				message = append(message, fmt.Sprintf("Available dates and time for next appointment are:"))
				// Apptinfo : none = no appt being made, first = first appt being made, second = second appt being made
				apptinfo = "second"

				counter := 0
				for _, apptItem := range apptArray {
					if timeNow.Before(apptItem) {
						if timeFirst29.After(apptItem) && timeFirst20.Before(apptItem) {
							counter++
							possibleAppts = append(possibleAppts, fmt.Sprintf("%s", apptItem.Format(layoutDateTime)))
						}
						if counter >= numOfApptsShown {
							break
						}
					}
				}

				if counter == 0 {
					message = append(message, fmt.Sprintf("\nNo available appointments found! Please contact MOH Hotline\n "))
					return nil, firstname, lastname, message, apptinfo, possibleAppts
				} else {
					message = append(message, fmt.Sprintf("Please choose your appointment date/time:"))
					return nil, firstname, lastname, message, apptinfo, possibleAppts
				}
			} else {
				message = append(message, fmt.Sprintf("Your first vaccination was done on %s", timeFirst.Format(layoutDateTime)))
				if timeNow.After(timeSecond) {
					message = append(message, fmt.Sprintf("You missed your second appointment on %s", timeSecond.Format(layoutDateTime)))
				} else {
					message = append(message, fmt.Sprintf("You already have an appointment on %s", timeSecond.Format(layoutDateTime)))
				}
				message = append(message, fmt.Sprintf("Delete this appointment before making new appointment"))
				return nil, firstname, lastname, message, apptinfo, possibleAppts
			}
		}
	} else {
		// make both appts together - difference of 21-28 days
		//no first appt made
		if timeFirst.IsZero() {
			message = append(message, fmt.Sprintf("Available dates and time for first appointment are:"))
			// Apptinfo : none = no appt being made, first = first appt being made, second = second appt being made
			apptinfo = "first"
			counter := 0
			for _, apptItem := range apptArray {
				if timeNow.Before(apptItem) {
					counter++
					possibleAppts = append(possibleAppts, fmt.Sprintf("%s", apptItem.Format(layoutDateTime)))
				}
				if counter >= numOfApptsShown {
					break
				}
			}
			if counter == 0 {
				message = append(message, fmt.Sprintf("\nNo available appointments found! Please contact MOH Hotline\n "))
				return nil, firstname, lastname, message, apptinfo, possibleAppts
			} else {
				message = append(message, fmt.Sprintf("Please choose your appointment date/time:"))
				return nil, firstname, lastname, message, apptinfo, possibleAppts
			}
		} else { // first appt done. check if second appt is done
			if timeSecond.IsZero() {
				// For second appointment after choosing first appointment date
				if timeNow.After(timeFirst) {
					message = append(message, fmt.Sprintf("You missed your first appointment on %s", timeFirst.Format(layoutDateTime)))
					message = append(message, fmt.Sprintf("Delete this appointment, and make new appointments"))
				} else {
					message = append(message, fmt.Sprintf("You first vaccination appointment is on %s", timeFirst.Format(layoutDateTime)))
					message = append(message, fmt.Sprintf("Available dates and time for second appointment are:"))
					// Apptinfo : none = no appt being made, first = first appt being made, second = second appt being made
					apptinfo = "second"

					counter := 0
					for _, apptItem := range apptArray {
						if timeNow.Before(apptItem) {
							if timeFirst29.After(apptItem) && timeFirst20.Before(apptItem) {
								counter++
								possibleAppts = append(possibleAppts, fmt.Sprintf("%s", apptItem.Format(layoutDateTime)))
							}
							if counter >= numOfApptsShown {
								break
							}
						}
					}
					if counter == 0 {
						message = append(message, fmt.Sprintf("\nNo available appointments found! Please contact MOH Hotline\n "))
						return nil, firstname, lastname, message, apptinfo, possibleAppts
					} else {
						message = append(message, fmt.Sprintf("Please choose your appointment date/time:"))
						return nil, firstname, lastname, message, apptinfo, possibleAppts
					}
				}
			}
		}
	}
	return nil, firstname, lastname, message, apptinfo, possibleAppts
}

// This is a method for Person struct. It takes the time for the appointment selected by user.
// It also takes in information whether appointment time is for first or second vacination.
// This is used to update the Person data with appointment time for first/second vaccination.
// It also takes in the appointment array to be updated, the selected appointment time will be deleted from the appointment array.
// It returns the updated appointment array.
func (p *Person) UpdateNewAppt(apptTime string, apptFor string, apptArray []time.Time) (error, []time.Time) {

	timeOfAppt, err := dateTimeFormat(apptTime) // convert string to time.Time format
	if err != nil {
		return err, apptArray
	}

	apptArray = deleteFromApptArray(apptArray, timeOfAppt)
	d := timeOfAppt.Format(layoutDateTime)
	dateArray := strings.Split(d, " ")

	if apptFor == "first" {
		p.FirstVaccineDate = dateArray[0]
		p.FirstVaccineTime = dateArray[1]
		p.FirstVaccineDone = false
	} else if apptFor == "second" {
		p.SecondVaccineDate = dateArray[0]
		p.SecondVaccineTime = dateArray[1]
		p.SecondVaccineDone = false
	}
	return nil, apptArray
}

// This is a method for Person struct. It returns user appt information for template deleteappt.gohtml.
// It returns first and last name information, message to be displayed, first/second/both appointments to be deleted.
func (p *Person) DeleteAppt() (err error, fn string, ln string, msg []string, aptI string) {

	var firstname string
	var lastname string
	var message []string
	var apptinfo string

	timeFirst, ferr := datePlusTime(p.FirstVaccineDate, p.FirstVaccineTime)
	timeSecond, serr := datePlusTime(p.SecondVaccineDate, p.SecondVaccineTime)
	if ferr != nil {
		return ferr, firstname, lastname, message, apptinfo
	}
	if serr != nil {
		return serr, firstname, lastname, message, apptinfo
	}

	timeNow := time.Now()

	firstname = p.Firstname
	lastname = p.Lastname

	if p.FirstVaccineDone { // cannot delete if done
		if p.SecondVaccineDone { // cannot delete if done
			message = append(message, fmt.Sprintf("Your COVID-19 Vaccination is complete. Take good rest and be safe."))
		} else {
			//no second appt made
			if timeSecond.IsZero() {
				message = append(message, fmt.Sprintf("Your first vaccination was done on %s", timeFirst.Format(layoutDateTime)))
				message = append(message, fmt.Sprintf("You have no appointment for second vaccination. Please make new appointment"))
			} else {
				message = append(message, fmt.Sprintf("Your first vaccination was done on %s", timeFirst.Format(layoutDateTime)))

				if timeNow.After(timeSecond) {
					message = append(message, fmt.Sprintf("You missed your second appointment dated %s", timeSecond.Format(layoutDateTime)))
				} else {
					message = append(message, fmt.Sprintf("Your second appointment is scheduled for %s", timeSecond.Format(layoutDateTime)))
				}

				apptinfo = "second"
				message = append(message, fmt.Sprintf("Click confirm to delete second appointment"))

			}
		}
	} else {
		//no first appt made
		if timeFirst.IsZero() {
			//no second appt made
			if timeSecond.IsZero() {
				message = append(message, fmt.Sprintf("You have no appointments made. Please make new appointments"))
			} else {
				message = append(message, fmt.Sprintf("You have no appointment for first vaccination. Deleting second appointment"))
				//corner case, will not happen if everything goes well
				p.SecondVaccineDate = ""
				p.SecondVaccineTime = ""
				p.SecondVaccineDone = false
			}
		} else {
			//no second appt made
			if timeSecond.IsZero() {
				if timeNow.After(timeFirst) {
					message = append(message, fmt.Sprintf("You missed your first appointment dated %s", timeFirst.Format(layoutDateTime)))
				} else {
					message = append(message, fmt.Sprintf("Your first appointment is scheduled for %s", timeFirst.Format(layoutDateTime)))
				}
				message = append(message, fmt.Sprintf("You have no appointments for second vaccination."))
				apptinfo = "first"
				message = append(message, fmt.Sprintf("Click confirm to delete first appointment"))

			} else {
				if timeNow.After(timeFirst) {
					message = append(message, fmt.Sprintf("You missed your first appointment dated %s", timeFirst.Format(layoutDateTime)))
				} else {
					message = append(message, fmt.Sprintf("Your first appointment is scheduled for %s", timeFirst.Format(layoutDateTime)))
				}
				if timeNow.After(timeSecond) {
					message = append(message, fmt.Sprintf("You missed your second appointment dated %s", timeSecond.Format(layoutDateTime)))
				} else {
					message = append(message, fmt.Sprintf("Your second appointment is scheduled for %s", timeSecond.Format(layoutDateTime)))
				}
				apptinfo = "both"
				message = append(message, fmt.Sprintf("Click confirm to delete both appointments"))

			}
		}
	}
	return nil, firstname, lastname, message, apptinfo
}

// This is a method for Person struct.
// It takes in the appointment to be deleted i.e. first/second or both, and the appointment array.
// It inserts the deleted appointments back into the appointement array in a sorted fashion.
// It return the updated appointment array.
func (p *Person) UpdatedeleteAppt(appttoDelete string, apptArray []time.Time) ([]time.Time, error) {

	timeFirst, ferr := datePlusTime(p.FirstVaccineDate, p.FirstVaccineTime)
	timeSecond, serr := datePlusTime(p.SecondVaccineDate, p.SecondVaccineTime)

	if ferr != nil {
		return apptArray, ferr
	}
	if serr != nil {
		return apptArray, serr
	}

	timeNow := time.Now()

	if appttoDelete == "first" {
		p.FirstVaccineDate = ""
		p.FirstVaccineTime = ""
		p.FirstVaccineDone = false
		//if appt is after now, add to apptArray
		if timeNow.Before(timeFirst) {
			apptArray = insertApptArray(apptArray, timeFirst)
		}

	} else if appttoDelete == "second" {
		p.SecondVaccineDate = ""
		p.SecondVaccineTime = ""
		p.SecondVaccineDone = false
		//if appt is after now, add to apptArray
		if timeNow.Before(timeSecond) {
			apptArray = insertApptArray(apptArray, timeSecond)
		}
	} else if appttoDelete == "both" {
		p.FirstVaccineDate = ""
		p.FirstVaccineTime = ""
		p.FirstVaccineDone = false
		p.SecondVaccineDate = ""
		p.SecondVaccineTime = ""
		p.SecondVaccineDone = false
		//if appt is after now, add to apptArray
		if timeNow.Before(timeFirst) {
			apptArray = insertApptArray(apptArray, timeFirst)
		}
		//if appt is after now, add to apptArray
		if timeNow.Before(timeSecond) {
			apptArray = insertApptArray(apptArray, timeSecond)
		}

	}
	return apptArray, nil
}

// This function concatenates date string and time string, and converts into time.Time
// It uses function dateTimeFormat() to convert string to time.Time
// It returns the converted time.Time
func datePlusTime(date, timeOfDay string) (time.Time, error) {
	if date == "" {
		var t time.Time

		return t, nil
	}
	dateWithTime := date + " " + timeOfDay
	return dateTimeFormat(dateWithTime)
}

// This function convert string to time.Time format
// It uses time.Parse function to convert string into time using layoutDateTime (2006-01-02 15:04:05) format
func dateTimeFormat(dateWithTime string) (time.Time, error) {
	dateTemp, err := time.Parse(layoutDateTime, dateWithTime)
	if err != nil {
		return dateTemp, err
	}
	return dateTemp, nil
}

// This function takes in appointment array as a slice of time.Time, and an appointment to be added in time.Time
// It inserts the appointment to be added into the slice, and returns the updated appointment array slice
// Insertion is done such that the array is always sorted, and any duplicate entries are removed.
func insertApptArray(apptArray []time.Time, addDate time.Time) []time.Time {

	for i := 0; i < len(apptArray); i++ {
		arrayTime := apptArray[i]
		if i+1 < len(apptArray) { // i+1 is not more than array size
			arrayTimeNext := apptArray[i+1]
			if arrayTime.Before(addDate) && arrayTimeNext.After(addDate) {
				tempArray := append(apptArray[:i+1], addDate)
				tempArray = append(tempArray, apptArray[i+1:]...)
				return tempArray
			} else if arrayTime.Equal(addDate) {
				// Duplicate item, no need to add
				return apptArray
			}
		} else { // last item reached
			return append(apptArray, addDate)
		}
	}
	return apptArray
}

// This function takes in appointment array as a slice of time.Time, and an appointment to be deleted in time.Time
// It deletes the appointment to be deleted from the slice, and returns the updated appointment array slice
func deleteFromApptArray(apptArray []time.Time, deleteThis time.Time) []time.Time {
	for i := 0; i < len(apptArray); i++ {
		arrayTime := apptArray[i]
		if arrayTime.Sub(deleteThis) == 0 {
			return append(apptArray[:i], apptArray[i+1:]...)
		}
	}
	return apptArray
}

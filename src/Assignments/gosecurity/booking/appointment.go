package booking

import (
	"database/sql"
	"gosecurity/config"
	"gosecurity/logger"
	"net/http"
	"strconv"
)

//Order and field name of struct Appointments matches with table in database.
type Appointments struct {
	AppointmentID int
	PatientID     int
	DoctorID      int
	ScheduleID    int
}

//IndexA route the request towards appointment booking page.
func IndexA(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/appointments", http.StatusSeeOther)
}

//AppointmentIndex will display appointment information.
//Admin can look for available appointments.
//It will display which patient ID has appointment with which doctor ID and on what shedule ID.
//      For Example: Patient ID 1 has appointment with Doctor ID 1 and their schedule ID is 1.
//Each appointment has its own ID.
func AppointmentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	} //return rows of table in database which are fields of struct in go
	rows, err := config.DB.Query("SELECT * FROM appointments")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	//make slice of appointment struct
	apts := make([]Appointments, 0)
	for rows.Next() { //loop over the slice elements untill return value is true
		apt := Appointments{}
		//scan will get data from the query into struct//
		err := rows.Scan(&apt.AppointmentID, &apt.PatientID, &apt.DoctorID, &apt.ScheduleID)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		apts = append(apts, apt)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	config.TPL.ExecuteTemplate(w, "appointments.gohtml", apts)
}

//ShowAppointment will display all the booked appointments.
/*If select on individual appointment ID, it will display information related to
that selected appointment ID.*/
func ShowAppointments(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	aptid := r.FormValue("appointmentid")
	if aptid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	//show selected appointment, it pull out only one record
	row := config.DB.QueryRow("SELECT * FROM appointments WHERE appointmentid = $1", aptid)
	//$1 represents the first arguement in postgres it is equivalent to first field in the struct
	apt := Appointments{}
	err := row.Scan(&apt.AppointmentID, &apt.PatientID, &apt.DoctorID, &apt.ScheduleID)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "showAppointment.gohtml", apt)
}

//CreateAppointment will facilitate in booking appointment.
func CreateAppointments(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "createAppointment.gohtml", nil)
}

//CreateAppointmentProcess will allow to enter information in the given fields.
//Once information like patient id etc has been added the appointment will get create.
func CreateAppointmentProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	apt := Appointments{}
	aptid := r.FormValue("appointmentid")
	ptid := r.FormValue("patientid")
	drid := r.FormValue("doctorid")
	schid := r.FormValue("scheduleid")

	// convert string to int
	str, err := strconv.Atoi(aptid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Please enter a number for the appointment id", http.StatusNotAcceptable)
		logger.ErrorLog.Println("Wrong input")
		return
	}
	apt.AppointmentID = int(str)

	str, err = strconv.Atoi(ptid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Please enter a number from the available Patients ID", http.StatusNotAcceptable)
		logger.ErrorLog.Println("Wrong input")
		return
	}
	apt.PatientID = int(str)

	str, err = strconv.Atoi(drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Please enter a number from the avialable Doctors ID", http.StatusNotAcceptable)
		logger.ErrorLog.Println("Wrong input")
		return
	}
	apt.DoctorID = int(str)

	str, err = strconv.Atoi(schid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Please enter a number from the available Schedule ID", http.StatusNotAcceptable)
		logger.ErrorLog.Println("Wrong input")
		return
	}
	apt.ScheduleID = int(str)

	if aptid == "" || ptid == "" || drid == "" || schid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	//Insert id entered into application, Exec func executes query.
	_, err = config.DB.Exec("INSERT INTO appointments (appointmentid, patientid, doctorid, scheduleid) VALUES ($1, $2, $3, $4)", apt.AppointmentID, apt.PatientID, apt.DoctorID, apt.ScheduleID)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- ID not avaialble, please input available ID only", http.StatusNotAcceptable)
		logger.WarningLog.Println("Input ID not available")
		return
	}

	// confirm insertion
	config.TPL.ExecuteTemplate(w, "appointmentProcessCreated.gohtml", apt)
	logger.TraceLog.Println("Appointment created")
}

//DeleteAppointment will facilitate in deleting booked appointment.
/*Once the appointment delete doctor and DoctorSchedule ID will be available
to book for susequent appointments.*/
func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	aptid := r.FormValue("appointmentid")
	if aptid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// delete appointment
	_, err := config.DB.Exec("DELETE FROM appointments WHERE appointmentid=$1;", aptid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/appointments", http.StatusSeeOther)
	logger.InfoLog.Println("Appointment details deleted from database")

}

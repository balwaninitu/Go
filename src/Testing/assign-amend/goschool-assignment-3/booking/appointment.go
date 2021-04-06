package booking

import (
	"Assignments/goschool-assignment-3/util"
	"database/sql"
	"net/http"
	"strconv"
)

//order and naming of struct matches with table in database
type Appointments struct {
	AppointmentID int
	PatientID     int
	DoctorID      int
	ScheduleID    int
}

func IndexA(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/appointments", http.StatusSeeOther)
}

func AppointmentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	} //return rows of table in database which are fields of struct in go
	rows, err := util.DB.Query("SELECT * FROM appointments")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	//make slice of appointment struct
	apts := make([]Appointments, 0)
	for rows.Next() { //loop over the slice elements untill return value is true
		apt := Appointments{}
		err := rows.Scan(&apt.AppointmentID, &apt.PatientID, &apt.DoctorID, &apt.ScheduleID) /*scan will
		get data from the query into struct */
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
	util.TPL.ExecuteTemplate(w, "appointments.gohtml", apts)
}
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
	row := util.DB.QueryRow("SELECT * FROM appointments WHERE appointmentid = $1", aptid)
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

	util.TPL.ExecuteTemplate(w, "showAppointment.gohtml", apt)
}

func CreateAppointments(w http.ResponseWriter, r *http.Request) {
	util.TPL.ExecuteTemplate(w, "createAppointment.gohtml", nil)
}

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
		return
	}
	apt.AppointmentID = int(str)

	str, err = strconv.Atoi(ptid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Please enter a number from the available Patients ID", http.StatusNotAcceptable)
		return
	}
	apt.PatientID = int(str)

	str, err = strconv.Atoi(drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Please enter a number from the avialable Doctors ID", http.StatusNotAcceptable)
		return
	}
	apt.DoctorID = int(str)

	str, err = strconv.Atoi(schid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Please enter a number from the available Schedule ID", http.StatusNotAcceptable)
		return
	}
	apt.ScheduleID = int(str)

	if aptid == "" || ptid == "" || drid == "" || schid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert id entered into application, Exec func executes query
	_, err = util.DB.Exec("INSERT INTO appointments (appointmentid, patientid, doctorid, scheduleid) VALUES ($1, $2, $3, $4)", apt.AppointmentID, apt.PatientID, apt.DoctorID, apt.ScheduleID)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- ID not avaialble to add, please add available ID only", http.StatusNotAcceptable)
		return
	}

	// confirm insertion
	util.TPL.ExecuteTemplate(w, "appointmentProcessCreated.gohtml", apt)
}

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
	_, err := util.DB.Exec("DELETE FROM appointments WHERE appointmentid=$1;", aptid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/appointments", http.StatusSeeOther)

}

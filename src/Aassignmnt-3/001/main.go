package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB
var tpl *template.Template
var err error

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:password@localhost/dentalclinic1?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You are connected to your database.")
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

//Appointments can be exported
type Appointments struct {
	AppointmentID int
	PatientID     int
	DoctorID      int
	ScheduleID    int
}

//DoctorDetails can be exported
type DoctorDetails struct {
	DoctorID   int
	DoctorName string
}

//DoctorSchedule can be exported
type DoctorSchedule struct {
	ScheduleID    int
	DoctorID      string
	AvailableTime time.Time
	AvailableFlag bool
}

//PatientDetails can be exported
type PatientDetails struct {
	PatientID   int
	PatientName string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/a", index1)
	http.HandleFunc("/appointments", appointmentIndex)
	http.HandleFunc("/appointments/show", showAppointments)
	http.HandleFunc("/appointments/create", createAppointments)
	http.HandleFunc("/appointments/create/process", createAppointmentProcess)
	http.HandleFunc("/appointments/delete/process", deleteAppointment)

	http.HandleFunc("/d", index2)
	http.HandleFunc("/doctordetails", doctorIndex)
	http.HandleFunc("/doctordetails/show", showDoctor)
	http.HandleFunc("/doctordetails/create", createDoctor)
	http.HandleFunc("/doctordetails/delete/process", deleteDoctor)

	http.HandleFunc("/s", index3)
	http.HandleFunc("/doctorschedule", scheduleIndex)
	http.HandleFunc("/doctorschedule/show", showDoctorSchedule)
	http.HandleFunc("/doctorschedule/create", createDoctorSchedule)
	http.HandleFunc("/doctorschedule/delete/process", deleteDoctorSchedule)

	http.HandleFunc("/p", index4)
	http.HandleFunc("/patientdetails", patientIndex)
	http.HandleFunc("/patientdetails/show", showPatient)
	http.HandleFunc("/patientdetails/create", createPatient)
	http.HandleFunc("/patientdetails/delete/process", deletePatient)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "***Welcome to Dental Appointment system***\n\n/a for Appointments\n\n/d for Doctordetails\n\n/s for Doctor's Schedule\n\n/p for Patients Details")
}

func index1(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/appointments", http.StatusSeeOther)
}

func index2(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doctordetails", http.StatusSeeOther)

}

func index3(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doctorschedule", http.StatusSeeOther)

}

func index4(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/patientdetails", http.StatusSeeOther)

}

func appointmentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT * FROM appointments")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	apts := make([]Appointments, 0)
	for rows.Next() {
		apt := Appointments{}
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
	tpl.ExecuteTemplate(w, "appointments.gohtml", apts)
}
func showAppointments(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	aptid := r.FormValue("appointmentid")
	if aptid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM appointments WHERE appointmentid = $1", aptid)

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

	tpl.ExecuteTemplate(w, "showAppointment.gohtml", apt)
}

func createAppointments(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "createAppointment.gohtml", nil)
}

func createAppointmentProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form valuess
	apt := Appointments{}
	aptid := r.FormValue("appointmentid")
	ptid := r.FormValue("patientid")
	drid := r.FormValue("doctorid")
	schid := r.FormValue("scheduleid")

	// convert form values
	str, err := strconv.Atoi(aptid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the appointment id", http.StatusNotAcceptable)
		return
	}
	apt.AppointmentID = int(str)

	str, err = strconv.Atoi(ptid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the patient id", http.StatusNotAcceptable)
		return
	}
	apt.PatientID = int(str)

	str, err = strconv.Atoi(drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the doctor id", http.StatusNotAcceptable)
		return
	}
	apt.DoctorID = int(str)

	str, err = strconv.Atoi(schid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the schedule id", http.StatusNotAcceptable)
		return
	}
	apt.ScheduleID = int(str)

	if aptid == "" || ptid == "" || drid == "" || schid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("INSERT INTO appointments (appointmentid, patientid, doctorid, scheduleid) VALUES ($1, $2, $3, $4)", apt.AppointmentID, apt.PatientID, apt.DoctorID, apt.ScheduleID)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "appointmentProcessCreated.gohtml", apt)
}

func deleteAppointment(w http.ResponseWriter, r *http.Request) {
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
	_, err := db.Exec("DELETE FROM appointments WHERE appointmentid=$1;", aptid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/appointments", http.StatusSeeOther)

}

func doctorIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM doctordetails")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	drs := make([]DoctorDetails, 0)
	for rows.Next() {
		dr := DoctorDetails{}
		err := rows.Scan(&dr.DoctorID, &dr.DoctorName)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		drs = append(drs, dr)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	tpl.ExecuteTemplate(w, "doctor.gohtml", drs)

}

func showDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	drid := r.FormValue("doctorid")
	if drid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM doctordetails WHERE doctorid = $1", drid)

	dr := DoctorDetails{}
	err := row.Scan(&dr.DoctorID, &dr.DoctorName)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "showDoctor.gohtml", dr)
}

func createDoctor(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "createDoctor.gohtml", nil)
}

func deleteDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	drid := r.FormValue("doctorid")
	if drid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// delete doctor
	_, err := db.Exec("DELETE FROM doctordetails WHERE doctorid=$1;", drid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/doctordetails", http.StatusSeeOther)

}

func scheduleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM doctorschedule")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	schs := make([]DoctorSchedule, 0)
	for rows.Next() {
		sch := DoctorSchedule{}
		err := rows.Scan(&sch.ScheduleID, &sch.DoctorID, &sch.AvailableTime, &sch.AvailableFlag)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		schs = append(schs, sch)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	tpl.ExecuteTemplate(w, "doctorschedule.gohtml", schs)

}

func showDoctorSchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	schid := r.FormValue("scheduleid")
	if schid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM doctorschedule WHERE scheduleid = $1", schid)

	sch := DoctorSchedule{}
	err := row.Scan(&sch.ScheduleID, &sch.DoctorID, sch.AvailableTime, sch.AvailableFlag)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "showDoctorSchedule.gohtml", sch)
}

func createDoctorSchedule(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "createDoctorSchedule.gohtml", nil)
}

func patientIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM patientdetails")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	pts := make([]PatientDetails, 0)
	for rows.Next() {
		pt := PatientDetails{}
		err := rows.Scan(&pt.PatientID, &pt.PatientName)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		pts = append(pts, pt)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	tpl.ExecuteTemplate(w, "patient.gohtml", pts)

}

func deleteDoctorSchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	schid := r.FormValue("scheduleid")
	if schid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// delete doctorschedule
	_, err := db.Exec("DELETE FROM doctorschedule WHERE schedulid=$1;", schid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/doctorschedule", http.StatusSeeOther)

}

func showPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	ptid := r.FormValue("patientid")
	if ptid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM patientdetails WHERE patientid = $1", ptid)

	pt := PatientDetails{}
	err := row.Scan(&pt.PatientID, &pt.PatientName)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "showPatient.gohtml", pt)
}

func createPatient(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "createPatient.gohtml", nil)
}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	ptid := r.FormValue("patientid")
	if ptid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// delete patient
	_, err := db.Exec("DELETE FROM patientdetails WHERE patientid=$1;", ptid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/patientdetails", http.StatusSeeOther)

}

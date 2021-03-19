package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
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

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/appointments", appointmentIndex)
	r.HandleFunc("/appointments/show", showAppointments)
	r.HandleFunc("/appointments/create", createAppointments)
	r.HandleFunc("/appointments/create/process", createAppointmentProcess)
	//http.HandleFunc("/appointments/update", updateAppointment)
	r.HandleFunc("/appointments/update/process", updateAppointmentProcess)
	r.HandleFunc("/appointments/delete/process", deleteAppointment)
	log.Fatal(http.ListenAndServeTLS(":8082", "C:/Users/Lenovo/cert.pem", "C:/Users/Lenovo/key.pem", r))

}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/appointments", http.StatusSeeOther)
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

// func updateAppointment(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
// 		return
// 	}

// 	aptid := r.FormValue("appointmentid")
// 	if aptid == "" {
// 		http.Error(w, http.StatusText(400), http.StatusBadRequest)
// 		return
// 	}

// 	row := db.QueryRow("SELECT * FROM appointments WHERE appointmentid = $1", aptid)

// 	apt := Appointments{}
// 	err := row.Scan(&apt.AppointmentID, &apt.PatientID, &apt.DoctorID, apt.ScheduleID)
// 	switch {
// 	case err == sql.ErrNoRows:
// 		http.NotFound(w, r)
// 		return
// 	case err != nil:
// 		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
// 		return
// 	}
// 	tpl.ExecuteTemplate(w, "updateAppointment.gohtml", apt)
// }

func updateAppointmentProcess(w http.ResponseWriter, r *http.Request) {
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

	// convert form values
	str1, err := strconv.Atoi(aptid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the appointment id", http.StatusNotAcceptable)
		return
	}
	apt.AppointmentID = int(str1)

	str2, err := strconv.Atoi(ptid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the patient id", http.StatusNotAcceptable)
		return
	}
	apt.PatientID = int(str2)

	str3, err := strconv.Atoi(drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the doctor id", http.StatusNotAcceptable)
		return
	}
	apt.DoctorID = int(str3)

	str4, err := strconv.Atoi(schid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the schedule id", http.StatusNotAcceptable)
		return
	}
	apt.ScheduleID = int(str4)

	if aptid == "" || ptid == "" || drid == "" || schid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("UPDATE appointments SET appointmentid = $1, patientid=$2, doctorid=$3, scheduleid=$4 WHERE appointmentid=$1;", apt.AppointmentID, apt.PatientID, apt.DoctorID, apt.ScheduleID)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "appointmentProcessUpdate.gohtml", apt)

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

	// delete doctor
	_, err := db.Exec("DELETE FROM appointments WHERE appointmentid=$1;", aptid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/appointments", http.StatusSeeOther)

}

package main

import (
	"database/sql"
	"fmt"
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

//DoctorSchedule can be exported
type DoctorSchedule struct {
	ScheduleID    int
	DoctorID      string
	AvailableTime time.Time
	AvailableFlag bool
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/doctorschedule", scheduleIndex)
	http.HandleFunc("/doctorschedule/show", showDoctorSchedule)
	http.HandleFunc("/doctorschedule/create", createDoctorSchedule)
	http.HandleFunc("/doctorschedule/create/process", createDoctorScheduleProcess)
	http.HandleFunc("/doctorschedule/update", updateDoctorSchedule)
	http.HandleFunc("/doctorschedule/update/process", updateDoctorScheduleProcess)
	http.HandleFunc("/doctorschedule/delete/process", deleteDoctorSchedule)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doctorschedule", http.StatusSeeOther)

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

func createDoctorScheduleProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form valuess
	sch := DoctorSchedule{}
	schid := r.FormValue("scheduleid")
	drid := r.FormValue("doctorid")
	avlbTime := r.FormValue("availabletime")
	avlbFlag := r.FormValue("availableflag")

	// convert form values
	str1, err := strconv.Atoi(schid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the Schedule id", http.StatusNotAcceptable)
		return
	}
	sch.ScheduleID = int(str1)

	str2, err := strconv.Atoi(drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the doctor id", http.StatusNotAcceptable)
		return
	}
	sch.ScheduleID = int(str2)

	str3, err := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", avlbTime)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter in format[2021-03-13 01:00:00+08]  for the available time", http.StatusNotAcceptable)
		return
	}
	sch.AvailableTime = (str3)

	str4, err := strconv.ParseBool(avlbFlag)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a bool value for flag", http.StatusNotAcceptable)
		return
	}
	sch.AvailableFlag = bool(str4)

	if schid == "" || drid == "" || avlbTime == "" || avlbFlag == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("INSERT INTO doctorschedule (scheduleid, doctorid, availabletime, availableflag) VALUES ($1, $2, $3, $4", sch.ScheduleID, sch.DoctorID, sch.AvailableTime, sch.AvailableFlag)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "doctorScheduleProcessCreated.gohtml", sch)
}

func updateDoctorSchedule(w http.ResponseWriter, r *http.Request) {
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
	err := row.Scan(&sch.ScheduleID, &sch.DoctorID, &sch.AvailableTime, &sch.AvailableFlag)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "updateDoctorSchedule.gohtml", sch)
}

func updateDoctorScheduleProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	sch := DoctorSchedule{}
	schid := r.FormValue("scheduleid")
	drid := r.FormValue("doctorid")
	avlbTime := r.FormValue("availabletime")
	avlbFlag := r.FormValue("availableflag")

	// convert form values
	str1, err := strconv.Atoi(schid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the Schedule id", http.StatusNotAcceptable)
		return
	}
	sch.ScheduleID = int(str1)

	str2, err := strconv.Atoi(drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the doctor id", http.StatusNotAcceptable)
		return
	}
	sch.ScheduleID = int(str2)

	str3, err := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", avlbTime)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter in format[2021-03-13 01:00:00+08]  for the available time", http.StatusNotAcceptable)
		return
	}
	sch.AvailableTime = (str3)

	str4, err := strconv.ParseBool(avlbFlag)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a bool value for flag", http.StatusNotAcceptable)
		return
	}
	sch.AvailableFlag = bool(str4)

	if schid == "" || drid == "" || avlbTime == "" || avlbFlag == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("UPDATE doctorschedule SET scheduleid = $1 doctorid = $2, availabletime=$3,availableflag=$4  WHERE scheduleid=$1;", sch.ScheduleID, sch.DoctorID, sch.AvailableTime, sch.AvailableFlag)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	// confirm insertion
	tpl.ExecuteTemplate(w, "doctorScheduleProcessUpdate.gohtml", sch)

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

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

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

//PatientDetails can be exported
type PatientDetails struct {
	PatientID   int
	PatientName string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/patientdetails", patientIndex)
	http.HandleFunc("/patientdetails/show", showPatient)
	http.HandleFunc("/patientdetails/create", createPatient)
	http.HandleFunc("/patientdetails/create/process", createPatientProcess)
	http.HandleFunc("/patientdetails/update", updatePatient)
	http.HandleFunc("/patientdetails/update/process", updatePatientProcess)
	http.HandleFunc("/patientdetails/delete/process", deletePatient)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/patientdetails", http.StatusSeeOther)

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

func createPatientProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form valuess
	pt := PatientDetails{}
	ptid := r.FormValue("patientid")
	pt.PatientName = r.FormValue("patientname")

	// convert form values
	str, err := strconv.Atoi(ptid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the patient id", http.StatusNotAcceptable)
		return
	}
	pt.PatientID = int(str)

	if ptid == "" || pt.PatientName == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("INSERT INTO patientdetails (patientid, patientname) VALUES ($1, $2", pt.PatientID, pt.PatientName)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "patientProcessCreated.gohtml", pt)
}

func updatePatient(w http.ResponseWriter, r *http.Request) {
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
	tpl.ExecuteTemplate(w, "updatePatient.gohtml", pt)
}

func updatePatientProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	pt := PatientDetails{}
	ptid := r.FormValue("patientid")
	pt.PatientName = r.FormValue("patientname")

	// convert form values
	str, err := strconv.Atoi(ptid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the patient id", http.StatusNotAcceptable)
		return
	}
	pt.PatientID = int(str)

	if ptid == "" || pt.PatientName == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("UPDATE patientdetails SET patientid = $1, patientname=$2 WHERE patientid=$1;", pt.PatientID, pt.PatientName)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	// confirm insertion
	tpl.ExecuteTemplate(w, "patientProcessUpdate.gohtml", pt)

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

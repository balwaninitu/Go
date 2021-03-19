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

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:password@localhost/dentalclinic?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You are connected to your database.")
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

//PatientDetails exported
type PatientDetails struct {
	ID   int
	Name string
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

	rows, err := db.Query("SELECT * FROM patientdetails;")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	pts := make([]PatientDetails, 0)
	for rows.Next() {
		pt := PatientDetails{}
		err := rows.Scan(&pt.ID, &pt.Name)
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

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM patientdetails WHERE id = $1", id)

	pt := PatientDetails{}
	err := row.Scan(&pt.ID, &pt.Name)
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
	i := r.FormValue("id")
	pt.Name = r.FormValue("name")

	// convert form values
	str, err := strconv.Atoi(i)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the id", http.StatusNotAcceptable)
		return
	}
	pt.ID = int(str)

	// validate form values
	if i == "" || pt.Name == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("INSERT INTO patientdetails (id, name) VALUES ($1, $2)", pt.ID, pt.Name)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "processCreatedpatient.gohtml", pt)
}

func updatePatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM patientdetails WHERE id = $1", id)

	pt := PatientDetails{}
	err := row.Scan(&pt.ID, &pt.Name)
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
	i := r.FormValue("id")
	pt.Name = r.FormValue("name")

	// convert form values
	str, err := strconv.Atoi(i)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the id", http.StatusNotAcceptable)
		return
	}
	pt.ID = int(str)

	// validate form values
	if i == "" || pt.Name == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("UPDATE patientdetails SET id = $1, name=$2 WHERE id=$1;", pt.ID, pt.Name)
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

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// delete doctor
	_, err := db.Exec("DELETE FROM patientdetails WHERE id=$1;", id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/patientdetails", http.StatusSeeOther)

}

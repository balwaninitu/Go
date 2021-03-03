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

//DoctorDetails can be exported
type DoctorDetails struct {
	DoctorID   int
	DoctorName string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/doctordetails", doctorIndex)
	http.HandleFunc("/doctordetails/show", showDoctor)
	http.HandleFunc("/doctordetails/create", createDoctor)
	http.HandleFunc("/doctordetails/create/process", createDoctorProcess)
	http.HandleFunc("/doctordetails/update", updateDoctor)
	http.HandleFunc("/doctordetails/update/process", updateDoctorProcess)
	http.HandleFunc("/doctordetails/delete/process", deleteDoctor)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doctordetails", http.StatusSeeOther)

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

func createDoctorProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form valuess
	dr := DoctorDetails{}
	drid := r.FormValue("doctorid")
	dr.DoctorName = r.FormValue("doctorname")

	// convert form values
	str, err := strconv.Atoi(drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the doctor id", http.StatusNotAcceptable)
		return
	}
	dr.DoctorID = int(str)

	if drid == "" || dr.DoctorName == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("INSERT INTO doctordetails (doctorid, doctorname) VALUES ($1, $2", dr.DoctorID, dr.DoctorName)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "doctorProcessCreated.gohtml", dr)
}

func updateDoctor(w http.ResponseWriter, r *http.Request) {
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
	tpl.ExecuteTemplate(w, "updateDoctor.gohtml", dr)
}

func updateDoctorProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	dr := DoctorDetails{}
	drid := r.FormValue("doctorid")
	dr.DoctorName = r.FormValue("doctorname")

	// convert form values
	str, err := strconv.Atoi(drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the doctor id", http.StatusNotAcceptable)
		return
	}
	dr.DoctorID = int(str)

	if drid == "" || dr.DoctorName == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("UPDATE doctordetails SET doctorid = $1, doctorname=$2 WHERE doctorid=$1;", dr.DoctorID, dr.DoctorName)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	// confirm insertion
	tpl.ExecuteTemplate(w, "doctorProcessUpdate.gohtml", dr)

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

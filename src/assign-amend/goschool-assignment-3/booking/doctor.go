package booking

import (
	"Assignments/goschool-assignment-3/util"
	"database/sql"
	"net/http"
	"strconv"
)

//order and naming of struct matches with table in database
type DoctorDetails struct {
	DoctorID   int
	DoctorName string
}

func IndexD(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doctordetails", http.StatusSeeOther)

}

func DoctorIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := util.DB.Query("SELECT * FROM doctordetails")
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
	util.TPL.ExecuteTemplate(w, "doctor.gohtml", drs)

}

func ShowDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	drid := r.FormValue("doctorid")
	if drid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	//show available Doctor's details when any id selected on client location
	row := util.DB.QueryRow("SELECT * FROM doctordetails WHERE doctorid = $1", drid)

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

	util.TPL.ExecuteTemplate(w, "showDoctor.gohtml", dr)
}

func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	drid := r.FormValue("doctorid")
	if drid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := util.DB.QueryRow("SELECT * FROM doctordetails WHERE doctorid = $1", drid)

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
	util.TPL.ExecuteTemplate(w, "updateDoctor.gohtml", dr)
}

//only allow to change doctor name
func UpdateDoctorProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	dr := DoctorDetails{}
	drid := r.FormValue("doctorid")
	dr.DoctorName = r.FormValue("doctorname")

	// convert form values string to int
	str, err := strconv.Atoi(drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Please enter a number for the doctor id", http.StatusNotAcceptable)
		return
	}
	dr.DoctorID = int(str)

	if drid == "" || dr.DoctorName == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = util.DB.Exec("UPDATE doctordetails SET doctorid = $1, doctorname=$2 WHERE doctorid=$1;", dr.DoctorID, dr.DoctorName)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	// confirm insertion
	util.TPL.ExecuteTemplate(w, "doctorProcessUpdate.gohtml", dr)

}

func DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	drid := r.FormValue("doctorid")
	if drid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// delete doctor details only when Doctor Id is not taken up in doctor's schedule
	_, err := util.DB.Exec("DELETE FROM doctordetails WHERE doctorid=$1;", drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Doctor ID can not delete if it is scheduled already", http.StatusAlreadyReported)
		return
	}

	http.Redirect(w, r, "/doctordetails", http.StatusSeeOther)

}

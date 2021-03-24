package booking

import (
	"database/sql"
	"gosecurity/config"
	"gosecurity/logger"
	"net/http"
	"strconv"
)

//Order and field name of DoctorDetails struct matches with table in database.
type DoctorDetails struct {
	DoctorID   int
	DoctorName string
}

//IndexD route the request towards the page where doctors details are available.
func IndexD(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doctordetails", http.StatusSeeOther)

}

/*DoctorIndex will display available doctor details such doctor ID
and doctor name in database. */
func DoctorIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query("SELECT * FROM doctordetails")
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
	config.TPL.ExecuteTemplate(w, "doctor.gohtml", drs)

}

//ShowDoctor will display doctor details available in database.
//If click on individual doctor ID, it will display information of each selected ID.
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
	row := config.DB.QueryRow("SELECT * FROM doctordetails WHERE doctorid = $1", drid)

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

	config.TPL.ExecuteTemplate(w, "showDoctor.gohtml", dr)
}

/*UpdateDoctor func allow admin to update the doctor details
such as doctor name for the available ID*/
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

	row := config.DB.QueryRow("SELECT * FROM doctordetails WHERE doctorid = $1", drid)

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
	config.TPL.ExecuteTemplate(w, "updateDoctor.gohtml", dr)
}

//UpdateDoctorProcess can allow admin to change doctor name for the give Doctor ID.
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
	_, err = config.DB.Exec("UPDATE doctordetails SET doctorid = $1, doctorname=$2 WHERE doctorid=$1;", dr.DoctorID, dr.DoctorName)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	// confirm insertion
	config.TPL.ExecuteTemplate(w, "doctorProcessUpdate.gohtml", dr)

}

//DeleteDoctor func can delete doctors details only if doctor ID is not schedule.
//However, doctor details can be deleted if it is not scheduled.
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

	//Delete doctor details only when Doctor Id is not taken up in doctor's schedule
	_, err := config.DB.Exec("DELETE FROM doctordetails WHERE doctorid=$1;", drid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Doctor ID can not delete if it is scheduled already", http.StatusAlreadyReported)
		logger.WarningLog.Println("Doctor ID can not delete if it is scheduled already for appointment")
		return
	}

	http.Redirect(w, r, "/doctordetails", http.StatusSeeOther)

}

package booking

import (
	"database/sql"
	"gosecurity/config"
	"gosecurity/logger"
	"net/http"
)

//Order and field name of struct PatientDetails matches with table in database.
type PatientDetails struct {
	PatientID   int
	PatientName string
}

//IndexD route the request towards the page where patients details are available.
func IndexP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/patientdetails", http.StatusSeeOther)

}

/*PatientIndex will display available patient details such patient ID
and patient name in database. */
func PatientIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query("SELECT * FROM patientdetails")
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
	config.TPL.ExecuteTemplate(w, "patient.gohtml", pts)

}

//ShowPatient will display patient details available in database.
//If click on individual patient ID, it will display information of each selected ID.
func ShowPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	ptid := r.FormValue("patientid")
	if ptid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := config.DB.QueryRow("SELECT * FROM patientdetails WHERE patientid = $1", ptid)

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

	config.TPL.ExecuteTemplate(w, "showPatient.gohtml", pt)
}

//DeletePatient func can delete patients details only if patient ID doesnt have any booked appointment.
//However, patient details can be deleted if it is not booked for any appointment.
func DeletePatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	ptid := r.FormValue("patientid")
	if ptid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// delete patients details only if corrensponding id doesnt have any book appointment
	_, err := config.DB.Exec("DELETE FROM patientdetails WHERE patientid=$1;", ptid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Patient Id can not delete as appointment already booked with this ID", http.StatusNotAcceptable)
		logger.WarningLog.Println("Patient ID can not delete if it is scheduled already for appointment")
		return
	}

	http.Redirect(w, r, "/patientdetails", http.StatusSeeOther)
	logger.InfoLog.Println("Patient's details deleted from database")

}

package booking

import (
	"Assignments/goschool-assignment-3/util"
	"database/sql"
	"net/http"
)

//PatientDetails can be exported
type PatientDetails struct {
	PatientID   int
	PatientName string
}

func IndexP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/patientdetails", http.StatusSeeOther)

}

func PatientIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := util.DB.Query("SELECT * FROM patientdetails")
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
	util.TPL.ExecuteTemplate(w, "patient.gohtml", pts)

}

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

	row := util.DB.QueryRow("SELECT * FROM patientdetails WHERE patientid = $1", ptid)

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

	util.TPL.ExecuteTemplate(w, "showPatient.gohtml", pt)
}

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

	// delete patients details
	_, err := util.DB.Exec("DELETE FROM patientdetails WHERE patientid=$1;", ptid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/patientdetails", http.StatusSeeOther)

}

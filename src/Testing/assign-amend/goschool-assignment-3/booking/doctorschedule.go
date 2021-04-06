package booking

import (
	"Assignments/goschool-assignment-3/util"
	"database/sql"
	"net/http"
	"time"
)

//DoctorSchedule order and naming of struct matches with table in database
type DoctorSchedule struct {
	ScheduleID    int
	DoctorID      string
	AvailableTime time.Time
	AvailableFlag bool
}

func IndexS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doctorschedule", http.StatusSeeOther)

}

func ScheduleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := util.DB.Query("SELECT * FROM doctorschedule")
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
	util.TPL.ExecuteTemplate(w, "doctorschedule.gohtml", schs)

}

func ShowDoctorSchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	schid := r.FormValue("scheduleid")
	if schid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	//show Docotr's available schedule details
	row := util.DB.QueryRow("SELECT * FROM doctorschedule WHERE scheduleid = $1", schid)

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

	util.TPL.ExecuteTemplate(w, "showDoctorSchedule.gohtml", sch)
}

func DeleteDoctorSchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	schid := r.FormValue("scheduleid")
	if schid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	_, err := util.DB.Exec("DELETE FROM doctorschedule WHERE schedulid=$1;", schid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Doctor ID can not delete if it is scheduled already", http.StatusAlreadyReported)
		return
	}

	http.Redirect(w, r, "/doctorschedule", http.StatusSeeOther)

}

package booking

import (
	"database/sql"
	"gosecurity/config"
	"gosecurity/logger"
	"net/http"
	"time"
)

//Order and field name of struct DoctorSchedule matches with table in database.
type DoctorSchedule struct {
	ScheduleID    int
	DoctorID      string
	AvailableTime time.Time
	AvailableFlag bool
}

//IndexS route the request towards the page where each doctor schedule is available.
func IndexS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doctorschedule", http.StatusSeeOther)

}

func ScheduleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query("SELECT * FROM doctorschedule")
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
	config.TPL.ExecuteTemplate(w, "doctorschedule.gohtml", schs)

}

//ShowDoctorSchedule will only display the doctors schedule available in database
//It will not allow any changes in details as if the ID is scheduled already for appointment.
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
	row := config.DB.QueryRow("SELECT * FROM doctorschedule WHERE scheduleid = $1", schid)

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

	config.TPL.ExecuteTemplate(w, "showDoctorSchedule.gohtml", sch)
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

	_, err := config.DB.Exec("DELETE FROM doctorschedule WHERE schedulid=$1;", schid)
	if err != nil {
		http.Error(w, http.StatusText(406)+"- Doctor ID can not delete if it is scheduled already", http.StatusAlreadyReported)
		logger.WarningLog.Println("Doctor ID can not delete if it is scheduled already for appointment")
		return
	}

	http.Redirect(w, r, "/doctorschedule", http.StatusSeeOther)

}

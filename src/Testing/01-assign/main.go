package main

import (
	"Testing/01-assign/logger"
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

//DoctorsDetails exported
type DoctorsDetails struct {
	ID      int
	Name    string
	DayTime time.Time
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/doctorsdetails", doctorsIndex)
	http.HandleFunc("/doctorsdetails/show", showDoctors)
	http.HandleFunc("/doctorsdetails/create", createDoctors)
	http.HandleFunc("/doctorsdetails/create/process", createDoctorsProcess)
	http.HandleFunc("/doctorsdetails/update", updateDoctors)
	http.HandleFunc("/doctorsdetails/update/process", updateDoctorsProcess)
	http.HandleFunc("/doctorsdetails/delete/process", deleteDoctor)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doctorsdetails", http.StatusSeeOther)

}

func doctorsIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM doctorsdetails;")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	drs := make([]DoctorsDetails, 0)
	for rows.Next() {
		dr := DoctorsDetails{}
		err := rows.Scan(&dr.ID, &dr.Name, &dr.DayTime)
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
	tpl.ExecuteTemplate(w, "doctors.gohtml", drs)

}

func showDoctors(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM doctorsdetails WHERE id = $1", id)

	dr := DoctorsDetails{}
	err := row.Scan(&dr.ID, &dr.Name, &dr.DayTime)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)

		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "showDoctor.gohtml", dr)
	logger.CommonLog.Println("Displaying Doctors details")
}

func createDoctors(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "createDoctor.gohtml", nil)
}

func createDoctorsProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form valuess
	dr := DoctorsDetails{}
	i := r.FormValue("id")
	dr.Name = r.FormValue("name")
	timeStr := dr.DayTime.Format("220902 050316")
	timeStr = r.FormValue("dayTime")

	// convert form values
	str, err := strconv.Atoi(i)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the id", http.StatusNotAcceptable)
		logger.ErrorLog.Println("Wrong input")
		return
	}
	dr.ID = int(str)

	// timeStr := "2019 05 27 11:23:45"
	layout := "020106 150405"
	_, err = time.Parse(layout, timeStr)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a format (220902 050316)  for the dateTime", http.StatusNotAcceptable)
		return
	}

	//dr.DayTime = timeStr.Format(layout)

	// validate form values
	if i == "" || dr.Name == "" || timeStr == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("INSERT INTO doctorsdetails (id, name, dayTime) VALUES ($1, $2, $3)", dr.ID, dr.Name, dr.DayTime)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "processCreated.gohtml", dr)
}

func updateDoctors(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM doctorsdetails WHERE id = $1", id)

	dr := DoctorsDetails{}
	err := row.Scan(&dr.ID, &dr.Name, &dr.DayTime)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "update.gohtml", dr)
}

func updateDoctorsProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	dr := DoctorsDetails{}
	i := r.FormValue("id")
	dr.Name = r.FormValue("name")
	timeStr := dr.DayTime.Format("220902 050316")
	timeStr = r.FormValue("dayTime")

	// convert form values
	str, err := strconv.Atoi(i)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a number for the id", http.StatusNotAcceptable)
		return
	}
	dr.ID = int(str)
	// You need to parse 050316 = 05:03:16
	// This means that you look at 15:04:05 of the first template and you change it so that it matches your value.
	// In this case you just remove ":" so 15:04:05 becomes 150405. This is your time layout.
	//
	// You need to parse 220902 = September 22, 2002.
	// You look at the second template 01/02 03:04:05PM '06 -0700.
	// All you need to keep is 01/02 '06 (month, day, year).
	// Your value is 220902 (day, month, year).
	// You remove "/" and " '" so 01/02 '06 becomes 010206.
	// But this is (month, day, year) so you swap 01 and 02 and 010206 becomes 020106. This is your date layout.
	//
	// To put them all together, in order to parse the datetime values:
	// "220902 050316" you will need the layout:
	// "020106 150405".
	layout := "020106 150405"
	_, err = time.Parse(layout, timeStr)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please enter a format (220902 050316) for the dateTime", http.StatusNotAcceptable)
		return
	}
	// validate form values
	if i == "" || dr.Name == "" || timeStr == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	_, err = db.Exec("UPDATE doctorsdetails SET id = $1, name=$2, dayTime=$3 WHERE id=$1;", dr.ID, dr.Name, dr.DayTime)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "processUpdate.gohtml", dr)

}

func deleteDoctor(w http.ResponseWriter, r *http.Request) {
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
	_, err := db.Exec("DELETE FROM doctorsdetails WHERE id=$1;", id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/doctorsdetails", http.StatusSeeOther)

}

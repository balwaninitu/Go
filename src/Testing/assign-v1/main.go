package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

//DoctorDetails is exported
type DoctorDetails struct {
	//drID          int
	//appointmentID int
	DoctorName string
	//dayTime       time.Time
	//available     bool
}

type patientDetails struct {
	patientName string
}

type apppointment struct {
	DoctorDetails
	patientDetails
}

// func init() {

// 	//d1 := DoctorDetails{DoctorName: "dr1"}
// }

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/doctors", doctorsIndex)
	http.HandleFunc("/doctors/show", doctorsShow)
	http.HandleFunc("/doctors/create", doctorsCreateForm)
	http.HandleFunc("/doctors/create/process", doctorsCreateProcess)
	http.HandleFunc("/doctors/update", doctorsUpdateForm)
	http.HandleFunc("/doctors/update/process", doctorsUpdateProcess)
	http.HandleFunc("/doctors/delete/process", doctorsDeleteProcess)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/doctors", http.StatusSeeOther)
}

func doctorsIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	drs := make([]DoctorDetails, 0)

	dk := DoctorDetails{DoctorName: "dr1"}
	drs = append(drs, dk)
	fmt.Println(drs)
	tpl.ExecuteTemplate(w, "doctors.gohtml", drs)
}

func doctorsShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	doctorName := r.FormValue("doctorName")
	if doctorName == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	dk := DoctorDetails{DoctorName: "dr1"}

	tpl.ExecuteTemplate(w, "show.gohtml", dk)
}

func doctorsCreateForm(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "create.gohtml", nil)
}

func doctorsCreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	dk := DoctorDetails{DoctorName: "dr1"}
	dk.DoctorName = r.FormValue("doctorName")

	// validate form values
	if dk.DoctorName == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "created.gohtml", dk)
}

func doctorsUpdateForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	doctorName := r.FormValue("doctorName")
	if doctorName == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	dk := DoctorDetails{}

	tpl.ExecuteTemplate(w, "update.gohtml", dk)
}

func doctorsUpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	dk := DoctorDetails{}
	dk.DoctorName = r.FormValue("doctorName")

	// validate form values
	if dk.DoctorName == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "updated.gohtml", dk)
}

func doctorsDeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	doctorName := r.FormValue("doctorName")
	if doctorName == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/doctors", http.StatusSeeOther)
}

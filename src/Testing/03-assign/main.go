package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"text/template"

// 	_ "github.com/lib/pq"
// )

// var db *sql.DB
// var tpl *template.Template

// func init() {
// 	var err error
// 	db, err = sql.Open("postgres", "postgres://goschool:password@localhost/dentalclinic?sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err = db.Ping(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("You are connected to your database.")
// 	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

// }

//
// type AppointmentDetails struct {
// 	ID   int
// 	drID int
// 	ptID int
// }

// func main() {
// 	http.HandleFunc("/", index)
// 	http.HandleFunc("/patientdetails", patientIndex)
// 	http.HandleFunc("/patientdetails/show", showPatient)
// 	http.HandleFunc("/patientdetails/create", createPatient)
// 	http.HandleFunc("/patientdetails/create/process", createPatientProcess)
// 	http.HandleFunc("/patientdetails/update", updatePatient)
// 	http.HandleFunc("/patientdetails/update/process", updatePatientProcess)
// 	http.HandleFunc("/patientdetails/delete/process", deletePatient)

// 	http.ListenAndServe(":8080", nil)

// }

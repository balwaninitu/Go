package main

import (
	"gosecurity/booking"
	"gosecurity/user"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	//Routers for user signup/login
	r.HandleFunc("/", user.Index)
	r.HandleFunc("/signup", user.Signup)
	r.HandleFunc("/login", user.Login)
	r.HandleFunc("/logout", user.Logout)
	//Routers for appointment booking
	r.HandleFunc("/a", booking.IndexA)
	r.HandleFunc("/appointments", booking.AppointmentIndex)
	r.HandleFunc("/appointments/show", booking.ShowAppointments)
	r.HandleFunc("/appointments/create", booking.CreateAppointments)
	r.HandleFunc("/appointments/create/process", booking.CreateAppointmentProcess)
	r.HandleFunc("/appointments/delete/process", booking.DeleteAppointment)
	//Routers to show and update doctor details
	r.HandleFunc("/d", booking.IndexD)
	r.HandleFunc("/doctordetails", booking.DoctorIndex)
	r.HandleFunc("/doctordetails/show", booking.ShowDoctor)
	r.HandleFunc("/doctordetails/update", booking.UpdateDoctor)
	r.HandleFunc("/doctordetails/update/process", booking.UpdateDoctorProcess)
	r.HandleFunc("/doctordetails/delete/process", booking.DeleteDoctor)
	//Routers to display available docotr schedule
	r.HandleFunc("/s", booking.IndexS)
	r.HandleFunc("/doctorschedule", booking.ScheduleIndex)
	r.HandleFunc("/doctorschedule/show", booking.ShowDoctorSchedule)
	r.HandleFunc("/doctorschedule/delete/process", booking.DeleteDoctorSchedule)
	//Routers to show and delete available patients
	r.HandleFunc("/p", booking.IndexP)
	r.HandleFunc("/patientdetails", booking.PatientIndex)
	r.HandleFunc("/patientdetails/show", booking.ShowPatient)
	r.HandleFunc("/patientdetails/delete/process", booking.DeletePatient)

	s := &http.Server{
		Addr:              ":8081",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	err := s.ListenAndServeTLS("C:/Users/Lenovo/cert.pem", "C:/Users/Lenovo/key.pem")
	if err != nil {
		log.Fatal("ListenAndServe", err)

	}

}

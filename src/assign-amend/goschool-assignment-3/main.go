package main

import (
	"Assignments/goschool-assignment-3/booking"
	"io"
	"net/http"
)

func main() {
	//Routers for appointment booking
	http.HandleFunc("/", index)
	http.HandleFunc("/a", booking.IndexA)
	http.HandleFunc("/appointments", booking.AppointmentIndex)
	http.HandleFunc("/appointments/show", booking.ShowAppointments)
	http.HandleFunc("/appointments/create", booking.CreateAppointments)
	http.HandleFunc("/appointments/create/process", booking.CreateAppointmentProcess)
	http.HandleFunc("/appointments/delete/process", booking.DeleteAppointment)
	//Routers to show and update doctor details
	http.HandleFunc("/d", booking.IndexD)
	http.HandleFunc("/doctordetails", booking.DoctorIndex)
	http.HandleFunc("/doctordetails/show", booking.ShowDoctor)
	http.HandleFunc("/doctordetails/update", booking.UpdateDoctor)
	http.HandleFunc("/doctordetails/update/process", booking.UpdateDoctorProcess)
	http.HandleFunc("/doctordetails/delete/process", booking.DeleteDoctor)
	//Routers to display available docotr schedule
	http.HandleFunc("/s", booking.IndexS)
	http.HandleFunc("/doctorschedule", booking.ScheduleIndex)
	http.HandleFunc("/doctorschedule/show", booking.ShowDoctorSchedule)
	http.HandleFunc("/doctorschedule/delete/process", booking.DeleteDoctorSchedule)
	//Routers to show and delete available patients
	http.HandleFunc("/p", booking.IndexP)
	http.HandleFunc("/patientdetails", booking.PatientIndex)
	http.HandleFunc("/patientdetails/show", booking.ShowPatient)
	http.HandleFunc("/patientdetails/delete/process", booking.DeletePatient)

	http.ListenAndServe(":5221", nil)
}

//handler to display welcome page
func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "*****Welcome to Dental Appointment system*****\n\n/a for Appointments  (Create Read and Delete only)\n\n/d for Doctordetails  (Read and Update Doctor Name Only)\n\n/s for Doctor's Schedule (Read Only)\n\n/p for Patients Details  (Read Only)")
}

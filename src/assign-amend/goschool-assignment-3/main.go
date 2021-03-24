package main

import (
	"Assignments/goschool-assignment-3/booking"
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	timeout    = time.Duration(1 * time.Second)
	timeoutMsg = "your request has timed out"
)

func main() {
	r := mux.NewRouter()
	//Routers for appointment booking
	r.HandleFunc("/", index)
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

	http.ListenAndServe(":5221", r)
	//h := timeoutHandler{}
	// rootHandler := &MyHandler{}
	// srv := &http.Server{
	// 	ReadTimeout:       1 * time.Second,
	// 	WriteTimeout:      1 * time.Second,
	// 	IdleTimeout:       30 * time.Second,
	// 	ReadHeaderTimeout: 30 * time.Second,
	// 	Handler:           rootHandler,
	// 	Addr:              ":5221",
	// }

	// if err := srv.ListenAndServe(); err != nil {
	// 	fmt.Printf("Server Failed: %s\n", err)
	// }
}

//handler to display welcome page
func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "*****Welcome to Dental Appointment system*****\n\n/a for Appointments  (Create Read and Delete only)\n\n/d for Doctordetails  (Read and Update Doctor Name Only)\n\n/s for Doctor's Schedule (Read Only)\n\n/p for Patients Details  (Read Only)")
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this request will always timeout!
	time.Sleep(timeout)
}

type MyHandler struct{}

type timeoutHandler struct{}

func (h timeoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	timer := time.AfterFunc(5*time.Second, func() {
		r.Body.Close()
	})
	bodyBytes := make([]byte, 0)
	for {
		//We reset the timer, for the variable time
		timer.Reset(1 * time.Second)

		_, err := io.CopyN(bytes.NewBuffer(bodyBytes), r.Body, 256)
		if err == io.EOF {
			// This is not an error in the common sense
			// io.EOF tells us, that we did read the complete body
			break
		} else if err != nil {
			//You should do error handling here
			break
		}
	}
}

package main

import (
	"courses_api/src/application"
	"log"
	"net/http"
	"time"

	"github.com/balwaninitu/courses_rest_api/src/controllers"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func mapUrls() {
	router.HandleFunc("/courses", controllers.CourseController.Create).Methods("POST")
}
func StartApplication() {

	mapUrls()

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 500 * time.Second,
		ReadTimeout:  500 * time.Second,
		IdleTimeout:  600 * time.Second,
		Handler:      router,
	}
	log.Println(" Listening on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)

	}

}

func main() {
	application.StartApplication()

}

package application

import (
	"github.com/balwaninitu/courses_rest_api/src/controllers"
)

func mapUrls() {
	router.HandleFunc("/courses", controllers.CourseController.Create).Methods("POST")
	router.HandleFunc("/courses/{id}", controllers.CourseController.Get).Methods("GET")
	router.HandleFunc("/courses/{id}", controllers.CourseController.Update).Methods("PUT")
	router.HandleFunc("/courses/{id}", controllers.CourseController.Delete).Methods("DELETE")
}

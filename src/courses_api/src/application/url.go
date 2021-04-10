package application

import (
	"courses_api/src/controllers"
)

func mapUrls() {
	router.POST("/courses", controllers.Create)
	//router.GET("/courses/{id}", controllers.CourseController.Get)
	//router.PUT("/courses/{id}", controllers.CourseController.Update)
	//router.DELETE("/courses/{id}", controllers.CourseController.Delete)
}

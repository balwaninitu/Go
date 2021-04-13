package application

import (
	"courses_api/src/controllers"
)

//mapUrls func will handle router with method which get invoke.
func mapUrls() {
	router.POST("/courses", controllers.Create)
	router.GET("/courses/:course_id", controllers.Get)
	router.PUT("/courses/:course_id", controllers.Update)
	router.DELETE("/courses/:course_id", controllers.Delete)

}

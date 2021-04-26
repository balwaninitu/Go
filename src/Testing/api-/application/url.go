package application

import "api-/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/courses/:course_id", controllers.Get)
	//router.POST("/courses", controllers.Create)
	router.PUT("/courses:course_id", controllers.Update)
	router.DELETE("/courses:course_id", controllers.Delete)
}

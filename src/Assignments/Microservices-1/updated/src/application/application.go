package application

import (
	"log"

	"github.com/gin-gonic/gin"
)

//default return engine with logger and recovery middleware
//every request get handle by gingonic router
//router will create different goroutine for different request/handle
var (
	router = gin.Default()
)

/*startApplication func is implementing mapUrl func and
Run attaches the router to a http.Server and starts listening and serving HTTP requests.*/
func StartApplication() {
	mapUrls()

	log.Println(" Listening on port 8080")
	router.Run(":8080")
}

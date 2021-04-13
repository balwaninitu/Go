package application

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	mapUrls()

	log.Println("Listening on port 5000")
	router.Run(":5000")
}

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

	log.Println(" Listening on port 8080")
	router.Run(":8080")
}

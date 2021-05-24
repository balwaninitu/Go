package start

import (
	"log"
	"net/http"

	config "projectGoLive/application/config"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {

	mapUrls()
	log.Println(" Listening on port ", config.PortNum)
	http.ListenAndServe(config.PortNum, router)
	// need to use https, and self generated cert - in cert folder - CertPath
}

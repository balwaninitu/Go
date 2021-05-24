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
	//log.Fatal(http.ListenAndServeTLS(":5221", certPath+"cert.pem", certPath+"key.pem", router))

	//log.Fatal(http.ListenAndServeTLS(config.PortNum, config.CertPath+"cert.pem", config.CertPath+"key.pem", router))
	//err := http.ListenAndServeTLS(":5221", config.CertPath+"cert.pem", config.CertPath+"key.pem", router)
	//fmt.Println(err)
}

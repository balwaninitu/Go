package application

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

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

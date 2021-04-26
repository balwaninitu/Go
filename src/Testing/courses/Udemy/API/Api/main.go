package main

import (
	"Api/controllers"
	"Api/driver"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDB()
	controller := controllers.Controller{}

	r := mux.NewRouter()
	r.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	r.HandleFunc("/login", controller.Login(db)).Methods("POST")
	r.HandleFunc("/protected", controller.TokenVerifyMiddleware(controller.ProtectedEndpoint())).Methods("GET")
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}

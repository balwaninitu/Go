package main

import (
	"courses_api/src/application"
	"database/sql"
	"log"
)

var (
	Client *sql.DB
)

func main() {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	log.Println(" Connected to Database ")

	application.StartApplication()

}

package main

import (
	"projectGoLive/application/start"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Handle function for all server functions

	//http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	start.StartApplication()
}

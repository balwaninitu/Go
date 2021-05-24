package main

import (
	"projectGoLive/application/start"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	start.StartApplication()
}

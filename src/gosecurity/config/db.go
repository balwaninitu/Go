package config

import (
	"database/sql"
	"fmt"
	"gosecurity/logger"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	/*init func will execute before main func and open database by using given link of driver name
	and Ping func will provide connection */
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:password@localhost/dentalclinic1?sslmode=disable")
	if err != nil {
		panic(err)
	} //check the connection
	if err = DB.Ping(); err != nil {
		panic(err) //system will panic if no connection to database
	}
	fmt.Println("You are connected to your database.")
	logger.TraceLog.Println("Connected to database")

}

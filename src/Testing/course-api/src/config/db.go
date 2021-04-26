package config

import (
	"database/sql"
	"fmt"
	"log"

	"course-api/src/logger"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

/*init func will invoke before main func and will open databse to
perform operation */
func init() {
	//to run application locally
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"user", "password", "127.0.0.1:3306", "users_db")

	//to run application on docker
	//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
	//	"root", "password", "127.0.0.1:62958", "users_db")

	DB, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to database successfully")
	logger.TraceLog.Println("Connected to database")
}

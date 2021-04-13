package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
	err    error
)

/*init func will invoke before main func and will open databse to
perform operation */
func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"user", "password", "127.0.0.1:3306", "users_db")
	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to database successfully")
}

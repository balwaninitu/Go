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

	// username = os.Getenv("mysql_username")
	// password = os.Getenv("mysql_password")
	// host     = os.Getenv("mysql_host")
	// schema   = os.Getenv("mysql_schema")
)

func init() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"user", "password", "127.0.0.1:3306", "users_db")

	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}
	//defer Client.Close()
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println(" Connected to Database ")

}

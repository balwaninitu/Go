package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	var err error
	//user:password@tcp(127.0.0.1:3306)/my_db
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
	// 	"user", "password", "127.0.0.1:3306", "my_db")
	Client, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database connected")
	fmt.Println("Database connected")
}

package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB

	// username = os.Getenv("mysql_username")
	// password = os.Getenv("mysql_password")
	// host     = os.Getenv("mysql_host")
	// schema   = os.Getenv("mysql_schema")
)

// func init() {

// 	//gotenv.Load()

// 	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
// 	// 	"user", "password", "127.0.0.1:3306", "my_db")

// 	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	log.Println(" Connected to Database ")

// }

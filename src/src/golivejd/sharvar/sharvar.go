//Package sharvar purpose is to share  variables acrossa all pkgs of this application
//
//
package sharvar

import (
	"database/sql"
	"fmt"
	"html/template"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

//postgresDB database connection
var MySqlDB *sql.DB

var Mu sync.Mutex
var Tpl *template.Template
var CurLoginUsers map[string]string

func init() {

	Tpl = template.Must(template.ParseGlob("C:/Projects/Go/src/golivejd/templates/*"))

	//	Tpl = template.Must(template.ParseGlob("templates/*"))
	MySqlDB, _ = sql.Open("mysql", "root:apamss@tcp(127.0.0.1:3306)/batch2jd")

	// if there is an error opening the connection, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }

	defer MySqlDB.Close()

}

func openDatabaseCon() (*sql.DB, error) {

	MySqlDB, err := sql.Open("mysql", "root:apamss@tcp(127.0.0.1:3306)/batch2jd")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err)
	}

	// defer the close till after the main function has finished
	// executing
	defer MySqlDB.Close()

	err = MySqlDB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("<< sharvar connecion ok = ", err)
	fmt.Println("<< sharvar open connecion ok = ", MySqlDB.Stats().OpenConnections)

	return MySqlDB, err
}

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {

	db, err := sql.Open("mysql", "root:apamss@tcp(127.0.0.1:3306)/batch2jd")

	if err != nil {
		panic(err.Error())
	}

	return db
}

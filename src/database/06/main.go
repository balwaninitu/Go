package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Courses struct {
	Id    int64
	Title string
}

func InsertRecord(db *sql.DB, Id int, Title string) {
	query := fmt.Sprintf("INSERT INTO courses VALUES (%d,'%s')",
		Id, Title)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/users_db")
	//db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:52572)/my_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Database opened")

	InsertRecord(db, 120, "action1")

}

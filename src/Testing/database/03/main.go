package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

func InsertRecord(db *sql.DB, ID int, FN string, LN string, Age int) {
	query := fmt.Sprintf("INSERT INTO Persons VALUES (%d,' %s', '%s', %d)",
		ID, FN, LN, Age)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")
	//db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:52572)/my_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Database opened")

	InsertRecord(db, 10, "TARA", "Tan", 23)
}

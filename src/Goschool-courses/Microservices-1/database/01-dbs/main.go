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

func DeleteRecord(db *sql.DB, ID int) {
	query := fmt.Sprintf(
		"DELETE FROM Persons WHERE ID ='%d'", ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

}

func InsertRecord(db *sql.DB, ID int, FN string, LN string, Age int) {
	query := fmt.Sprintf("INSERT INTO Persons VALUES (%d,' %s', '%s', %d)",
		ID, FN, LN, Age)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func EditRecord(db *sql.DB, ID int, FN string, LN string, Age int) {
	query := fmt.Sprintf(
		"UPDATE Persons SET FirstName= '%s', LastName = '%s', Age=%d WHERE ID=%d",
		FN, LN, Age, ID)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM my_db.Persons")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var person Person
		err = results.Scan(&person.ID, &person.FirstName,
			&person.LastName, &person.Age)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(person.ID, person.FirstName,
			person.LastName, person.Age)
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

	InsertRecord(db, 11, "Ryan", "Tan", 23)
	//EditRecord(db, 2, "Ariana", "Grande", 32)
	//DeleteRecord(db, 2)
	GetRecords(db)

}

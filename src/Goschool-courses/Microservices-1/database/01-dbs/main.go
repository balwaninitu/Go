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

func EdittRecord(db *sql.DB, ID int, FN string, LN string, Age int) {
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

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Database opened")

	//InsertRecord(db, 2, "Ryan", "Tan", 23)
	//EditRecord(db, 2, )
	GetRecords(db)

}

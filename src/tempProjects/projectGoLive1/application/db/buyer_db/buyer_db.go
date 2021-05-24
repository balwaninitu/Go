package buyer_db

// Functions to access buyer_db and authenticate buyer.
// Only db related functions here

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BuyerDetails struct {
	Username string
	Password string
	Location string
}

// Function to get all records from the MYSQL database.
// The function takes in the handle to the database.
// It returns all the info of all courses as an array of type BuyerDetails.
// It returns true when retrieval of records from the database is successful.
// It returns false when there is any error encountered and retrieval of records is not successful.
func GetRecords(db *sql.DB) ([]BuyerDetails, bool) {
	var bd []BuyerDetails

	results, err := db.Query("SELECT * FROM buyer_db.BuyerDetails")
	if err != nil {
		log.Println("Not able to get buyer details")
		log.Println(err)
	}

	for results.Next() {
		// map this type to the record in the table
		var bi BuyerDetails
		err = results.Scan(&bi.Username, &bi.Password)
		if err != nil {
			log.Println("Unable to get records")
			log.Println(err)
			return bd, false
		}
		bd = append(bd, bi)
	}
	return bd, true
}

// Following functions need to be updated for buyer_db

// Function to get one record from the MYSQL database.
// The function takes in the handle to the database.
// Its also takes in the name of the course to search for, of type string.
// It returns all the info of one course of type BuyerDetails.
// It returns true when retrieval of the record from the database is successful.
// It returns false when there is any error encountered, and retrieval of record is not successful.
func GetARecord(db *sql.DB, BN string) (BuyerDetails, bool) {
	var bd BuyerDetails
	query := fmt.Sprintf("SELECT * FROM buyer_db.BuyerDetails WHERE Username='%s'", BN)
	results, err := db.Query(query)
	if err != nil {
		log.Println("Unable to find a record")
		log.Println(err)
		return bd, false
	}

	for results.Next() {
		// map this type to the record in the table
		err = results.Scan(&bd.Username, &bd.Password, &bd.Location)
		if err != nil {
			log.Println("Unable to get the record")
			log.Println(err)
			return bd, false
		}
	}
	if bd.Username == "" {
		return bd, false
	}
	return bd, true
}

// Function to insert one record into the MYSQL database.
// The function takes in the handle to the database.
// Its also takes in the course to insert of type BuyerDetails.
// It returns true when the course is inserted into the database successfully.
// It returns false when there is any error encountered, and course is not inserted successfully.
func InsertRecord(db *sql.DB, bd BuyerDetails) bool {
	query := fmt.Sprintf("INSERT INTO BuyerDetails (Username, Password, Location) VALUES ('%s','%s',%s)", bd.Username, bd.Password, bd.Location)
	_, err := db.Query(query)

	if err != nil {
		log.Println("Unable to insert the record")
		log.Println(err)
		return false
	}
	return true
}

// Function to update an existing record in the MYSQL database.
// The function takes in the handle to the database.
// It also takes in the name of the user to update of type string, and the new details of the user of type BuyerDetails.
// It returns true when the user details is updated in the database successfully.
// It returns false when there is any error encountered, and user details is not updated successfully.
func EditRecord(db *sql.DB, BN string, bd BuyerDetails) bool {
	query := fmt.Sprintf("UPDATE BuyerDetails SET Username='%s', Password='%s', Location='%s' WHERE Username='%s'", bd.Username, bd.Password, bd.Location, BN)
	_, err := db.Query(query)
	if err != nil {
		log.Println("Unable to edit the record")
		log.Println(err)
		return false
	}
	return true
}

// Function to delete an existing record in the MYSQL database.
// The function takes in the handle to the database.
// It takes in the name of the user to delete of type string.
// It returns true when the user is deleted from the database successfully.
// It returns false when there is any error encountered, and user is not deleted successfully.
func DeleteRecord(db *sql.DB, BN string) bool {
	query := fmt.Sprintf(
		"DELETE FROM BuyerDetails WHERE Username='%s'", BN)
	_, err := db.Query(query)
	if err != nil {
		log.Println("Unable to delete the record")
		log.Println(err)
		return false
	}
	return true
}

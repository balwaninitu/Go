package seller_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type SellerDetails struct {
	Username string
	Password string
	Location string
}

// Function to get all records from the MYSQL database.
// The function takes in the handle to the database.
// It returns all the info of all users as an array of type SellerDetails.
// It returns true when retrieval of records from the database is successful.
// It returns false when there is any error encountered and retrieval of records is not successful.
func GetRecords(db *sql.DB) ([]SellerDetails, bool) {
	var sd []SellerDetails

	results, err := db.Query("SELECT * FROM seller_db.SellerDetails")
	if err != nil {
		log.Println("Not able to get seller details")
		log.Println(err)
	}

	for results.Next() {
		// map this type to the record in the table
		var si SellerDetails
		err = results.Scan(&si.Username, &si.Password, &si.Location)
		if err != nil {
			log.Println("Unable to get records")
			log.Println(err)
			return sd, false
		}
		sd = append(sd, si)
	}
	return sd, true
}

// Following functions need to be updated for seller_db

// Function to get one record from the MYSQL database.
// The function takes in the handle to the database.
// Its also takes in the name of the user to search for, of type string.
// It returns all the info of one user of type SellerDetails.
// It returns true when retrieval of the record from the database is successful.
// It returns false when there is any error encountered, and retrieval of record is not successful.
func GetARecord(db *sql.DB, SN string) (SellerDetails, bool) {
	var sd SellerDetails
	query := fmt.Sprintf("SELECT * FROM seller_db.SellerDetails WHERE Username='%s'", SN)
	results, err := db.Query(query)
	if err != nil {
		log.Println("Unable to find a record")
		log.Println(err)
		return sd, false
	}

	for results.Next() {
		// map this type to the record in the table
		err = results.Scan(&sd.Username, &sd.Password, &sd.Location)
		if err != nil {
			log.Println("Unable to get the record")
			log.Println(err)
			return sd, false
		}
	}
	if sd.Username == "" {
		return sd, false
	}
	return sd, true
}

// Function to insert one record into the MYSQL database.
// The function takes in the handle to the database.
// Its also takes in the user details to insert of type SellerDetails.
// It returns true when the user is inserted into the database successfully.
// It returns false when there is any error encountered, and user is not inserted successfully.
func InsertRecord(db *sql.DB, sd SellerDetails) bool {
	query := fmt.Sprintf("INSERT INTO SellerDetails (Username, Password, Location)	VALUES ('%s','%s',%s)", sd.Username, sd.Password, sd.Location)
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
// It also takes in the name of the user to update of type string, and the new details of the user of type SellerDetails.
// It returns true when the user details is updated in the database successfully.
// It returns false when there is any error encountered, and user details is not updated successfully.
func EditRecord(db *sql.DB, SN string, sd SellerDetails) bool {
	query := fmt.Sprintf("UPDATE SellerDetails SET Username='%s', Password='%s', Location='%s' WHERE Username='%s'", sd.Username, sd.Password, sd.Location, SN)
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
func DeleteRecord(db *sql.DB, SN string) bool {
	query := fmt.Sprintf(
		"DELETE FROM SellerDetails WHERE Username='%s'", SN)
	_, err := db.Query(query)
	if err != nil {
		log.Println("Unable to delete the record")
		log.Println(err)
		return false
	}
	return true
}

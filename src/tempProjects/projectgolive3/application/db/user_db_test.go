package user_db

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

var uinfo UserDetails

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error '%s' while opening database not expected", err)
	}
	return db, mock
}

func TestGetRecords(t *testing.T) {

	db, mock := NewMock()
	query := "SELECT * FROM `Userdetails`"
	rows := sqlmock.NewRows([]string{"Username", "Password", "Location"}).
		AddRow(uinfo.Username, uinfo.Password, uinfo.Isbuyer, uinfo.Location)

	mock.ExpectQuery(query).WillReturnRows(rows)
	actual, _ := GetRecords(db)

	var expected []UserDetails
	expected = append(expected, uinfo)
	assert.Equal(t, expected, actual)
}

func TestGetRecord(t *testing.T) {
	db, mock := NewMock()
	query := "SELECT * FROM `Userdetails` WHERE Username='%s'"
	rows := sqlmock.NewRows([]string{"Username", "Password", "Location"}).
		AddRow(uinfo.Username, uinfo.Password, uinfo.Isbuyer, uinfo.Location)

	mock.ExpectQuery(query).WithArgs(uinfo.Username).WillReturnRows(rows)
	actual, _ := GetARecord(db, uinfo.Username)
	var expected UserDetails
	assert.Equal(t, expected, actual)
}

func TestInsertRecord(t *testing.T) {
	db, mock := NewMock()
	query := "INSERT INTO `Userdetails`(Username, Password, Isbuyer, Location) VALUES ('%s','%s',%t,'%s')"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(uinfo.Username, uinfo.Password, uinfo.Isbuyer, uinfo.Location).
		WillReturnResult(sqlmock.NewResult(0, 1))
	_ = InsertRecord(db, uinfo)
}
func TestEditRecord(t *testing.T) {
	db, mock := NewMock()
	query := "UPDATE `Userdetails` SET Username='%s', Password='%s', Isbuyer=%t, Location='%s' WHERE Username='%s'"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(uinfo.Username, uinfo.Password, uinfo.Isbuyer, uinfo.Location).
		WillReturnResult(sqlmock.NewResult(0, 1))
	_ = EditRecord(db, uinfo.Username, uinfo)
}

func TestDeleteRecord(t *testing.T) {
	db, mock := NewMock()
	query := "DELETE FROM `Userdetails` WHERE Username='%s'"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(uinfo.Username).
		WillReturnResult(sqlmock.NewResult(0, 1))
	_ = DeleteRecord(db, uinfo.Username)
}

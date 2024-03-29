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

func TestGetRecord(t *testing.T) {
	db, mock := NewMock()
	query := "SELECT * FROM `Userdetails` WHERE Username='%s'"
	rows := sqlmock.NewRows([]string{"Username", "Password", "Fullname", "Phone", "Address", "Email"}).
		AddRow(uinfo.Username, uinfo.Password, uinfo.Fullname, uinfo.Phone, uinfo.Address, uinfo.Email)

	mock.ExpectQuery(query).WithArgs(uinfo.Username).WillReturnRows(rows)
	actual, _ := GetARecord(db, uinfo.Username)
	var expected UserDetails
	assert.Equal(t, expected, actual)
}

func TestInsertRecord(t *testing.T) {
	db, mock := NewMock()
	query := "INSERT INTO `userdetails`(Username, Password, Fullname, Phone, Address, Email) VALUES ('%s','%s','%s',%t,'%s','%s','%s')"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(uinfo.Username, uinfo.Password, uinfo.Fullname, uinfo.Phone, uinfo.Address, uinfo.Email).
		WillReturnResult(sqlmock.NewResult(0, 1))
	_ = InsertRecord(db, uinfo)
}
func TestEditRecord(t *testing.T) {
	db, mock := NewMock()
	query := "UPDATE `userdetails` SET Username='%s', Password='%s', Firstname='%s', Isbuyer=%t, Phone='%s', Address='%s', Email='%s') WHERE Username='%s'"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(uinfo.Username, uinfo.Password, uinfo.Fullname, uinfo.Phone, uinfo.Address, uinfo.Email).
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

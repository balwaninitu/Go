package main

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var iteminfo ItemsDetails

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error '%s' while opening database not expected", err)
	}
	return db, mock
}

func TestGetRecordsSeller(t *testing.T) {
	db, mock := NewMock()
	query := "SELECT * FROM sellerAPIdb.itemsdetails WHERE Username='%s';"
	rows := sqlmock.NewRows([]string{"Item", "Username"}).
		AddRow(iteminfo.Item, iteminfo.Quantity, iteminfo.Cost, iteminfo.Username)

	mock.ExpectQuery(query).WillReturnRows(rows)
	actual, _ := GetRecordsSeller(db, iteminfo.Username)

	var expected []ItemsDetails
	expected = append(expected, iteminfo)
	assert.Equal(t, expected, actual)
}

func TestGetARecordSeller(t *testing.T) {
	db, mock := NewMock()
	query := "SELECT * FROM sellerAPIdb.itemsdetails WHERE Username='%s' AND Item='%s';"
	rows := sqlmock.NewRows([]string{"Item", "Username"}).
		AddRow(iteminfo.Item, iteminfo.Quantity, iteminfo.Cost, iteminfo.Username)

	mock.ExpectQuery(query).WithArgs(iteminfo.Item, iteminfo.Username).WillReturnRows(rows)
	actual, _ := GetARecordSeller(db, iteminfo.Item, iteminfo.Username)

	var expected ItemsDetails
	assert.Equal(t, expected, actual)
}

func TestInsertRecordSeller(t *testing.T) {
	db, mock := NewMock()
	query := "INSERT INTO `itemsdetails` (Item, Quantity, Cost, Username) VALUES ('%s',%d, %f,'%s');"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(iteminfo.Item, iteminfo.Quantity, iteminfo.Cost, iteminfo.Username).
		WillReturnResult(sqlmock.NewResult(0, 1))
	_ = InsertRecordSeller(db, iteminfo)
}

func TestEditRecordSeller(t *testing.T) {
	db, mock := NewMock()
	query := "UPDATE `itemsdetails` SET Item='%s', Quantity= %d, Cost=%f, Username='%s' WHERE Item='%s' AND Username='%s';"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(iteminfo.Item, iteminfo.Quantity, iteminfo.Cost, iteminfo.Username).
		WillReturnResult(sqlmock.NewResult(0, 1))
	_ = EditRecordSeller(db, iteminfo.Item, iteminfo.Username, iteminfo)
}

func TestDeleteRecord(t *testing.T) {
	db, mock := NewMock()
	query := "DELETE FROM `Userdetails` WHERE Username='%s'"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(iteminfo.Username).
		WillReturnResult(sqlmock.NewResult(0, 1))
	_ = DeleteRecordSeller(db, iteminfo.Item, iteminfo.Username)
}

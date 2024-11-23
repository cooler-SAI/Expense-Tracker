package cmd

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddExpense(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectExec("INSERT INTO expenses").
		WithArgs(
			sqlmock.AnyArg(),
			"Test Description",
			100.50,
			"Test Category",
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectClose()

	description := "Test Description"
	amount := 100.50
	category := "Test Category"
	addExpenseTest(db, description, amount, category)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func addExpenseTest(db *sql.DB, description string, amount float64, category string) {
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	query := `INSERT INTO expenses (date, description, amount, category) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(query, time.Now().Format("2006-01-02"), description, amount, category)
	if err != nil {
		panic(err)
	}
}

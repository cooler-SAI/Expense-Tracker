package cmd

import (
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestQueryExpenses(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer func() {
		err := db.Close()
		assert.NoError(t, err)
	}()

	mock.ExpectQuery("SELECT date, description, amount, category FROM expenses").
		WillReturnRows(sqlmock.NewRows([]string{"date", "description", "amount", "category"}).
			AddRow("2024-11-22", "Test Description 1", 100.50, "Test Category 1").
			AddRow("2024-11-23", "Test Description 2", 200.75, "Test Category 2"))

	result := captureOutput(func() {
		queryExpenses(db)
	})

	assert.Contains(t, result, "2024-11-22 - Test Description 1: $100.50 (Test Category 1)")
	assert.Contains(t, result, "2024-11-23 - Test Description 2: $200.75 (Test Category 2)")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()
	err := w.Close()
	if err != nil {
		return ""
	}
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

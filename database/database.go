package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const DbName = "expenses.db"

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", DbName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT NOT NULL,
		description TEXT NOT NULL,
		amount REAL NOT NULL,
		category TEXT
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	fmt.Println("Database initialized successfully.")
	return db

}

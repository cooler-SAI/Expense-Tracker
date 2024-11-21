package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const SQLBaseName = "expense-tracker.db"

func initLog() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msgf("DATABASE: Init here...")
}

func InitDB() *sql.DB {
	initLog()
	db, err := sql.Open("sqlite3", SQLBaseName)
	if err != nil {
		log.Fatal().Msgf("Failed to open database: %v", err)
	}
	log.Info().Msgf("DATABASE: Using database file: %s", SQLBaseName)

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
		log.Fatal().Msgf("Error creating table: %v", err)
	}

	log.Info().Msg("DATABASE: Database initialized successfully.")
	return db

}

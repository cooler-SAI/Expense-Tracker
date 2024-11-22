package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const SQLBaseName = "expense-tracker.db"

var dbInstance *sql.DB

func initLog() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("DATABASE: Init here...")
}

func InitDB() *sql.DB {
	initLog()

	db, err := sql.Open("sqlite3", SQLBaseName)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open database")
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
		log.Fatal().Err(err).Msg("Error creating table")
	}

	log.Info().Msg("DATABASE: Database initialized successfully.")
	dbInstance = db
	return db
}

func GetDB() *sql.DB {
	if dbInstance == nil {
		log.Fatal().Msg("DATABASE: Instance is nil. Did you call InitDB?")
	}
	return dbInstance
}

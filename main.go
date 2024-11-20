package main

import (
	"Expense-Tracker/database"
	"database/sql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msgf("Hello Expense-Tracker")

	db := database.InitDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	log.Info().Msgf("Expense Tracker is ready to use!")
}

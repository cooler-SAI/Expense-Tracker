package main

import (
	"Expense-Tracker/cmd"
	"Expense-Tracker/database"
	"database/sql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msgf("Starting Expense-Tracker ...")

	db := database.InitDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing database")
		}
	}(db)

	cmd.Execute()

	log.Info().Msgf("Expense Tracker's work finished")
}

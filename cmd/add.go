package cmd

import (
	"Expense-Tracker/database"
	"database/sql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func initLog() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msgf("DATABASE: Init here...")
}

var description string
var amount float64
var category string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		addExpense(description, amount, category)
	},
}

func init() {
	initLog()
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&description, "description", "d", "",
		"Description of the expense")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0,
		"Amount of the expense")
	addCmd.Flags().StringVarP(&category, "category", "c", "Other",
		"Category of the expense")

	err := addCmd.MarkFlagRequired("description")
	if err != nil {
		log.Fatal().Msgf("Error marking flag: %v", err)
	}
	err2 := addCmd.MarkFlagRequired("amount")
	if err2 != nil {
		log.Fatal().Msgf("Error marking flag: %v", err)
	}
}

func addExpense(description string, amount float64, category string) {
	log.Info().Msgf("ADD COMMAND: description = %s, amount = %.2f, category = %s",
		description, amount, category)

	db, err := sql.Open("sqlite3", database.SQLBaseName)
	if err != nil {
		log.Fatal().Msgf("Error opening database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal().Msgf("Error closing database: %v", err)
		}
	}(db)

	query := `INSERT INTO expenses (date, description, amount, category) VALUES (?, ?, ?, ?)`
	log.Info().Msgf("ADD COMMAND: Executing query: %s", query)
	log.Info().Msgf("ADD COMMAND: Parameters: date = %s, description = %s, amount = %.2f, category = %s",
		time.Now().Format("2006-01-02"), description, amount, category)

	_, err = db.Exec(query, time.Now().Format("2006-01-02"), description, amount, category)
	if err != nil {
		log.Fatal().Msgf("Error adding expense: %v", err)
	}

	log.Info().Msgf("ADD COMMAND: Expense added successfully: %s - $%.2f (%s)\n", description, amount, category)
}

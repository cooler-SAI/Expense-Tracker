package cmd

import (
	"Expense-Tracker/database"
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("LIST COMMAND: Starting execution")

		db := database.GetDB()
		if db == nil {
			log.Fatal().Msg("DATABASE: Connection is nil")
		}

		log.Info().Msg("DATABASE: Querying expenses")
		queryExpenses(db)

		log.Info().Msg("LIST COMMAND: Finished execution")
	},
}

func queryExpenses(db *sql.DB) {
	var rows, err = db.Query("SELECT date, description, amount, category FROM expenses")
	if err != nil {
		log.Fatal().Err(err).Msg("DATABASE: Query failed")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	fmt.Println("Expenses:")
	for rows.Next() {
		var date, description, category string
		var amount float64

		if err := rows.Scan(&date, &description, &amount, &category); err != nil {
			log.Fatal().Err(err).Msg("DATABASE: Failed to scan row")
		}

		fmt.Printf("%s - %s: $%.2f (%s)\n", date, description, amount, category)
	}

	if err := rows.Err(); err != nil {
		log.Fatal().Err(err).Msg("DATABASE: Error during row iteration")
	}
}

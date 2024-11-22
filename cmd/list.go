package cmd

import (
	"Expense-Tracker/database"
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

		// Инициализация базы данных
		db := database.GetDB()
		if db == nil {
			log.Fatal().Msg("DATABASE: Connection is nil")
		}

		log.Info().Msg("DATABASE: Querying expenses")
		queryExpenses(db)

		log.Info().Msg("LIST COMMAND: Finished execution")
	},
}

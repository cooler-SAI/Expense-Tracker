package cmd

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	debug   bool
	version = "1.0.0"
)

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "A CLI tool for tracking expenses",
}

func init() {

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false,
		"Enable debug logging")

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of Expense-Tracker",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().Msgf("Expense Tracker %s", version)
		},
	})
}

func Execute() {

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Debug mode enabled")
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msg("Error executing command")
		os.Exit(1)
	}
}

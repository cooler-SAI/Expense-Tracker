package cmd

import (
	"Expense-Tracker/database"
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"time"
)

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
	rootCmd.AddCommand(addCmd) // Регистрация команды в rootCmd

	addCmd.Flags().StringVarP(&description, "description", "d", "",
		"Description of the expense")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0,
		"Amount of the expense")
	addCmd.Flags().StringVarP(&category, "category", "c", "Other",
		"Category of the expense")

	err := addCmd.MarkFlagRequired("description")
	if err != nil {
		log.Fatalf("Error marking flag: %v", err)
	}
	err2 := addCmd.MarkFlagRequired("amount")
	if err2 != nil {
		log.Fatalf("Error marking flag: %v", err)
	}
}

func addExpense(description string, amount float64, category string) {
	log.Printf("addExpense called with: description=%s, amount=%.2f, category=%s", description, amount, category)

	db, err := sql.Open("sqlite3", database.SQLBaseName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}(db)

	query := `INSERT INTO expenses (date, description, amount, category) VALUES (?, ?, ?, ?)`
	log.Printf("Executing query: %s", query)
	log.Printf("Parameters: date=%s, description=%s, amount=%.2f, category=%s",
		time.Now().Format("2006-01-02"), description, amount, category)

	_, err = db.Exec(query, time.Now().Format("2006-01-02"), description, amount, category)
	if err != nil {
		log.Fatalf("Error adding expense: %v", err)
	}

	fmt.Printf("Expense added successfully: %s - $%.2f (%s)\n", description, amount, category)
}

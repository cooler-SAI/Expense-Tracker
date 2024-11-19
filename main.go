package main

import (
	"Expense-Tracker/database"
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Hello Expense-Tracker")

	db := database.InitDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	fmt.Println("Expense Tracker is ready to use!")
}

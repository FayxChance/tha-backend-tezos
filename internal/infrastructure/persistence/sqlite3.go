package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite3Database struct {
	// Database sql.
	DB *sql.DB
}

func (sql3db *SQLite3Database) SetupDatabase() error {

	db, err := sql.Open("sqlite3", "./store-transactions.db")
	if err != nil {
		return err
	}

	// Check if the table already exists
	var tableName string
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='Delegations';").Scan(&tableName)

	// If table does not exist, create it
	if err != nil {
		if err == sql.ErrNoRows {
			// Table doesn't exist, proceed with creation
			fmt.Println("Table 'Delegations' does not exist, creating...")

			sqlStmt := `
				CREATE TABLE IF NOT EXISTS Delegations (
					id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Auto-incrementing primary key for SQLite
					timestamp TIMESTAMP NOT NULL,         -- Equivalent to time.Time in Go
					amount VARCHAR(255) NOT NULL,         -- String type for amount
					delegator VARCHAR(255) NOT NULL,      -- String type for delegator
					level VARCHAR(255) NOT NULL           -- String type for level
				);
			`
			_, err = db.Exec(sqlStmt)
			if err != nil {
				return err
			}
		} else {
			// Some other error occurred
			return err
		}
	} else {
		// Table already exists
		fmt.Println("Table 'Delegations' already exists.")
	}

	// Assign the DB connection to the struct
	sql3db.DB = db
	return nil
}

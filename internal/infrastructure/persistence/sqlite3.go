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
				CREATE TABLE Delegations (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					tzkt_id BIGINT UNIQUE,             
					timestamp TIMESTAMP NOT NULL,
					amount VARCHAR(255) NOT NULL,
					delegator VARCHAR(255) NOT NULL,
					level VARCHAR(255) NOT NULL
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

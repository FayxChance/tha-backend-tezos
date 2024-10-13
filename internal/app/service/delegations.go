package service

import (
	"fmt"
	"log"
	"time"

	"github.com/FayxChance/tha-backend-tezos/internal/infrastructure/persistence"
	"github.com/FayxChance/tha-backend-tezos/model"
)

type DelegationsService struct {
	DB persistence.SQLite3Database
}

func (svc *DelegationsService) Delegations() ([]model.Delegation, error) {

	if svc.DB.DB == nil {
		log.Printf("Database connection is nil")
		return nil, fmt.Errorf("database connection is nil")
	}

	delegations := make([]model.Delegation, 0)

	// Log the SQL query being executed
	query := "SELECT id, timestamp, amount, delegator, level FROM Delegations"
	log.Printf("Executing query: %s", query)

	// Execute the query
	rows, err := svc.DB.DB.Query(query)
	if err != nil {
		// Log the error if the query fails
		log.Printf("Error executing query: %v", err)
		return nil, err
	}

	defer rows.Close()
	var id int
	var timestamp time.Time
	var amount string
	var delegator string
	var level string

	// Iterate through the rows and scan the results
	for rows.Next() {
		err := rows.Scan(&id, &timestamp, &amount, &delegator, &level)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		log.Printf("Fetched row: id=%d, timestamp=%v, amount=%s, delegator=%s, level=%s", id, timestamp, amount, delegator, level)

		delegations = append(delegations, model.Delegation{
			Timestamp: timestamp,
			Amount:    amount,
			Delegator: delegator,
			Level:     level,
		})
	}

	// Log the total number of delegations fetched
	log.Printf("Total delegations fetched: %d", len(delegations))

	return delegations, nil
}

package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/FayxChance/tha-backend-tezos/internal/infrastructure/persistence"
	"github.com/FayxChance/tha-backend-tezos/model"
)

// DelegationsService handles the logic for fetching and managing delegations.
type DelegationsService struct {
	DB persistence.SQLite3Database // SQLite3 database connection
}

// Delegations fetches delegations from the database.
func (svc *DelegationsService) Delegations() ([]model.Delegation, error) {
	var delegations []model.Delegation

	// Query to select delegations ordered by the timestamp (most recent first)
	query := "SELECT id, timestamp, amount, delegator, level FROM Delegations ORDER BY timestamp DESC;"
	log.Printf("Executing query: %s", query)

	rows, err := svc.DB.DB.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var id int
	var timestamp time.Time
	var amount, delegator, level string

	// Iterate through the query results and append to delegations slice
	for rows.Next() {
		err := rows.Scan(&id, &timestamp, &amount, &delegator, &level)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		delegations = append(delegations, model.Delegation{
			Timestamp: timestamp,
			Amount:    amount,
			Delegator: delegator,
			Level:     level,
		})
	}

	log.Printf("Total delegations fetched: %d", len(delegations))
	return delegations, nil
}

// FetchFirst1000Delegations retrieves the first 1000 delegations from the Tezos API.
func (svc *DelegationsService) FetchFirst1000Delegations(apiURL string) error {
	params := url.Values{}
	params.Add("limit", "1000")
	params.Add("sort.desc", "id")
	fullURL := apiURL + "?" + params.Encode()

	fmt.Printf("URL requested: %s\n", fullURL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}
	defer resp.Body.Close()

	// Read and unmarshal the response body into the Transaction struct
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	var transactions []model.Transaction
	err = json.Unmarshal(body, &transactions)
	if err != nil {
		fmt.Println("Error unmarshaling response body:", err)
		return err
	}

	// Save the transactions as delegations
	return svc.SaveTransactions(transactions)
}

// FetchContinuousDelegations continuously polls the Tezos API for new delegations.
func (svc *DelegationsService) FetchContinuousDelegations(apiURL string) error {
	params := url.Values{}
	params.Add("limit", "10000")
	params.Add("sort.desc", "id")

	// Get the last stored tzkt_id to avoid fetching duplicates
	lastTzktIDInserted, err := svc.LastDelegationTzktID()
	if err != nil {
		return err
	}
	params.Add("id.gt", fmt.Sprintf("%d", lastTzktIDInserted))

	fullURL := apiURL + "?" + params.Encode()
	fmt.Printf("URL requested: %s\n", fullURL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	// If the response body is empty (length of 64), return
	if len(body) == 64 {
		return nil
	}

	// Unmarshal the response into the Transaction struct
	var transactions []model.Transaction
	err = json.Unmarshal(body, &transactions)
	if err != nil {
		fmt.Println("Error unmarshaling response body:", err)
		return err
	}

	fmt.Printf("[ContinuousDelegations] [%s] Delegation retrieved = %d\n", time.Now(), len(transactions))
	return svc.SaveTransactions(transactions)
}

// SaveTransactions converts transactions to delegations and stores them in the database.
func (svc *DelegationsService) SaveTransactions(transactions []model.Transaction) error {
	var delegations []model.Delegation

	for _, t := range transactions {
		delegations = append(delegations, model.TransactionToDelegation(t))
	}

	tx, err := svc.DB.DB.Begin() // Start a transaction
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO Delegations (tzkt_id, timestamp, amount, delegator, level) VALUES (?, ?, ?, ?, ?)`
	stmt, err := tx.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert each delegation into the database
	for _, delegation := range delegations {
		_, err = stmt.Exec(delegation.TzKTID, delegation.Timestamp, delegation.Amount, delegation.Delegator, delegation.Level)
		if err != nil {
			tx.Rollback() // Rollback if there's an error
			return fmt.Errorf("failed to insert delegation with ID %d: %v", delegation.TzKTID, err)
		}
	}

	err = tx.Commit() // Commit the transaction
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

// LastDelegationTzktID retrieves the most recent tzkt_id from the database.
func (svc *DelegationsService) LastDelegationTzktID() (int64, error) {
	var lastTzktID int64
	query := "SELECT tzkt_id FROM Delegations ORDER BY tzkt_id DESC LIMIT 1;"

	rows, err := svc.DB.DB.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&lastTzktID)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return 0, err
		}
	}

	return lastTzktID, nil
}

// StartFetch begins a background process to continuously fetch new delegations every 15 seconds.
func (svc *DelegationsService) StartFetch(apiURL string) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			svc.FetchContinuousDelegations(apiURL)
		}
	}
}

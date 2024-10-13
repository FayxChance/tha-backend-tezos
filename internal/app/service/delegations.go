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

type DelegationsService struct {
	DB persistence.SQLite3Database
}

func (svc *DelegationsService) Delegations() ([]model.Delegation, error) {
	var delegations []model.Delegation

	// Log the SQL query being executed
	query := "SELECT id, timestamp, amount, delegator, level FROM Delegations ORDER BY timestamp DESC;"
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

		// log.Printf("Fetched row: id=%d, timestamp=%v, amount=%s, delegator=%s, level=%s", id, timestamp, amount, delegator, level)

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

	// Unmarshal the response body into the Transaction struct
	var transactions []model.Transaction
	err = json.Unmarshal(body, &transactions)
	if err != nil {
		fmt.Println("Error unmarshaling response body:", err)
		return err
	}

	err = svc.SaveTransactions(transactions)
	if err != nil {
		return err
	}

	return nil
}

func (svc *DelegationsService) FetchContinuousDelegations(apiURL string) error {
	params := url.Values{}

	params.Add("limit", "10000")
	params.Add("sort.desc", "id")

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

	// If no new delegation, then exit the function. Len of 64 is an empty array
	if len(body) == 64 {
		return nil
	}

	// Unmarshal the response body into the Transaction struct
	var transactions []model.Transaction
	err = json.Unmarshal(body, &transactions)
	if err != nil {
		fmt.Println("Error unmarshaling response body:", err)
		return err
	}

	fmt.Printf("[ContinuousDelegations] [%s] Delegation retrieved = %d\n", time.Now(), len(transactions))
	err = svc.SaveTransactions(transactions)
	if err != nil {
		return err
	}

	return nil
}

func (svc *DelegationsService) SaveTransactions(transactions []model.Transaction) error {
	var delegations []model.Delegation

	for _, t := range transactions {
		delegations = append(delegations, model.TransactionToDelegation(t))
	}

	tx, err := svc.DB.DB.Begin()
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO Delegations (tzkt_id, timestamp, amount, delegator, level) VALUES (?, ?, ?, ?, ?)`
	stmt, err := tx.Prepare(insertQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, delegation := range delegations {
		// fmt.Printf("Inserting delegation - ID: %d, Level: %s, Timestamp: %s, Amount: %s\n", delegation.TzKTID, delegation.Level, delegation.Timestamp, delegation.Amount)
		_, err = stmt.Exec(delegation.TzKTID, delegation.Timestamp, delegation.Amount, delegation.Delegator, delegation.Level)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert transaction with ID %d: %v", delegation.TzKTID, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

func (svc *DelegationsService) LastDelegationTzktID() (int64, error) {
	var lastTzktID int64
	query := "SELECT tzkt_id FROM Delegations ORDER BY tzkt_id DESC LIMIT 1;"

	rows, err := svc.DB.DB.Query(query)
	if err != nil {
		// Log the error if the query fails
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

func (svc *DelegationsService) StartFetch(apiURL string) {
	// Poll every 30 seconds
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			svc.FetchContinuousDelegations(apiURL)
		}
	}
}

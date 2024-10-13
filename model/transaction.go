package model

import (
	"fmt"
	"time"
)

// Address struct represents a simple address model.
type Address struct {
	Address string `json:"address"` // The actual address, serialized in JSON as "address"
}

// Transaction struct represents a transaction model, used to capture details of a Tezos transaction.
type Transaction struct {
	ID        int64     `json:"id"`        // Unique identifier for the transaction
	Level     int       `json:"level"`     // Block level (height) at which the transaction occurred
	Timestamp time.Time `json:"timestamp"` // The time the transaction occurred
	Amount    int64     `json:"amount"`    // The amount transacted (as an integer for precision)
	Sender    Address   `json:"sender"`    // The sender's address, wrapped in an Address struct
}

// TransactionToDelegation converts a Transaction struct into a Delegation struct.
// This is useful for converting transaction data into the format required for delegation records.
func TransactionToDelegation(tx Transaction) Delegation {
	delegation := Delegation{
		TzKTID:    tx.ID,                        // Use the transaction ID as the TzKT ID
		Timestamp: tx.Timestamp,                 // Pass the transaction timestamp
		Amount:    fmt.Sprintf("%d", tx.Amount), // Convert amount to string for the Delegation model
		Delegator: tx.Sender.Address,            // Set the sender's address as the delegator
		Level:     fmt.Sprintf("%d", tx.Level),  // Convert block level to string
	}

	// Return the newly created delegation
	return delegation
}

// GenerateFalseTransactions creates some mock transaction data for testing or debugging purposes.
// This function generates fake transactions with different levels, timestamps, and senders.
func GenerateFalseTransactions() []Transaction {
	transactions := []Transaction{
		{
			ID:        1,
			Level:     1000,
			Timestamp: time.Now(),
			Sender:    Address{Address: "tz1sender1"},
		},
		{
			ID:        2,
			Level:     1010,
			Timestamp: time.Now().Add(-10 * time.Minute),
			Sender:    Address{Address: "tz1sender2"},
		},
		{
			ID:        3,
			Level:     1020,
			Timestamp: time.Now().Add(-20 * time.Minute),
			Sender:    Address{Address: "tz1sender3"},
		},
		{
			ID:        4,
			Level:     1030,
			Timestamp: time.Now().Add(-30 * time.Minute),
			Sender:    Address{Address: "tz1sender4"},
		},
		{
			ID:        5,
			Level:     1040,
			Timestamp: time.Now().Add(-40 * time.Minute),
			Sender:    Address{Address: "tz1sender5"},
		},
	}
	return transactions
}

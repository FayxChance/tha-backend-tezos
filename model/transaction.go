package model

import (
	"fmt"
	"time"
)

type Address struct {
	Address string `json:"address"`
}

type Transaction struct {
	ID        int64     `json:"id"`
	Level     int       `json:"level"`
	Timestamp time.Time `json:"timestamp"`
	Amount    int64     `json:"amount"`
	Sender    Address   `json:"sender"`
}

func TransactionToDelegation(tx Transaction) Delegation {
	delegation := Delegation{
		TzKTID:    tx.ID,
		Timestamp: tx.Timestamp,
		Amount:    fmt.Sprintf("%d", tx.Amount),
		Delegator: tx.Sender.Address,
		Level:     fmt.Sprintf("%d", tx.Level),
	}

	// Print the delegation for debugging
	// fmt.Printf("Converted Delegation: TzKTID: %d, Timestamp: %s, Amount: %s, Delegator: %s, Level: %s\n",
	// 	delegation.TzKTID, delegation.Timestamp, delegation.Amount, delegation.Delegator, delegation.Level)

	return delegation
}

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

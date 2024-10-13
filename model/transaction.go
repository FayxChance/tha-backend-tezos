package model

import "time"

// Error represents the structure of error in the JSON
type Error struct {
	Type string `json:"type"`
}

// Quote represents the structure of currency quotes in the JSON
type Quote struct {
	BTC float64 `json:"btc"`
	EUR float64 `json:"eur"`
	USD float64 `json:"usd"`
	CNY float64 `json:"cny"`
	JPY float64 `json:"jpy"`
	KRW float64 `json:"krw"`
	ETH float64 `json:"eth"`
	GBP float64 `json:"gbp"`
}

// Address represents structures for initiator, sender, prevDelegate, and newDelegate
type Address struct {
	Alias   string `json:"alias"`
	Address string `json:"address"`
}

// Transaction represents the main structure of the JSON
type Transaction struct {
	Type                 string    `json:"type"`
	ID                   int       `json:"id"`
	Level                int       `json:"level"`
	Timestamp            time.Time `json:"timestamp"`
	Block                string    `json:"block"`
	Hash                 string    `json:"hash"`
	Counter              int       `json:"counter"`
	Initiator            Address   `json:"initiator"`
	Sender               Address   `json:"sender"`
	SenderCodeHash       int       `json:"senderCodeHash"`
	Nonce                int       `json:"nonce"`
	GasLimit             int       `json:"gasLimit"`
	GasUsed              int       `json:"gasUsed"`
	StorageLimit         int       `json:"storageLimit"`
	BakerFee             int       `json:"bakerFee"`
	Amount               int       `json:"amount"`
	StakingUpdatesCount  int       `json:"stakingUpdatesCount"`
	PrevDelegate         Address   `json:"prevDelegate"`
	NewDelegate          Address   `json:"newDelegate"`
	Status               string    `json:"status"`
	Errors               []Error   `json:"errors"`
	Quote                Quote     `json:"quote"`
	UnstakedPseudotokens int       `json:"unstakedPseudotokens"`
	UnstakedBalance      int       `json:"unstakedBalance"`
	UnstakedRewards      int       `json:"unstakedRewards"`
}

func GenerateFalseTransactions() []Transaction {
	transactions := []Transaction{
		{
			Type:         "delegation",
			ID:           1,
			Level:        1000,
			Timestamp:    time.Now(),
			Block:        "B1",
			Hash:         "H1",
			Counter:      1,
			Initiator:    Address{Alias: "Init1", Address: "tz1address1"},
			Sender:       Address{Alias: "Sender1", Address: "tz1sender1"},
			PrevDelegate: Address{Alias: "Delegate1", Address: "tz1delegate1"},
			NewDelegate:  Address{Alias: "Delegate2", Address: "tz1delegate2"},
			Status:       "applied",
			Quote:        Quote{BTC: 0.1, EUR: 1000, USD: 1200, CNY: 8000, JPY: 130000, KRW: 1500000, ETH: 0.02, GBP: 900},
		},
		{
			Type:         "delegation",
			ID:           2,
			Level:        1010,
			Timestamp:    time.Now().Add(-10 * time.Minute),
			Block:        "B2",
			Hash:         "H2",
			Counter:      2,
			Initiator:    Address{Alias: "Init2", Address: "tz1address2"},
			Sender:       Address{Alias: "Sender2", Address: "tz1sender2"},
			PrevDelegate: Address{Alias: "Delegate3", Address: "tz1delegate3"},
			NewDelegate:  Address{Alias: "Delegate4", Address: "tz1delegate4"},
			Status:       "failed",
			Quote:        Quote{BTC: 0.2, EUR: 2000, USD: 2400, CNY: 16000, JPY: 260000, KRW: 3000000, ETH: 0.04, GBP: 1800},
		},
		{
			Type:         "delegation",
			ID:           3,
			Level:        1020,
			Timestamp:    time.Now().Add(-20 * time.Minute),
			Block:        "B3",
			Hash:         "H3",
			Counter:      3,
			Initiator:    Address{Alias: "Init3", Address: "tz1address3"},
			Sender:       Address{Alias: "Sender3", Address: "tz1sender3"},
			PrevDelegate: Address{Alias: "Delegate5", Address: "tz1delegate5"},
			NewDelegate:  Address{Alias: "Delegate6", Address: "tz1delegate6"},
			Status:       "applied",
			Quote:        Quote{BTC: 0.15, EUR: 1500, USD: 1800, CNY: 12000, JPY: 195000, KRW: 2250000, ETH: 0.03, GBP: 1350},
		},
		{
			Type:         "delegation",
			ID:           4,
			Level:        1030,
			Timestamp:    time.Now().Add(-30 * time.Minute),
			Block:        "B4",
			Hash:         "H4",
			Counter:      4,
			Initiator:    Address{Alias: "Init4", Address: "tz1address4"},
			Sender:       Address{Alias: "Sender4", Address: "tz1sender4"},
			PrevDelegate: Address{Alias: "Delegate7", Address: "tz1delegate7"},
			NewDelegate:  Address{Alias: "Delegate8", Address: "tz1delegate8"},
			Status:       "failed",
			Quote:        Quote{BTC: 0.05, EUR: 500, USD: 600, CNY: 4000, JPY: 65000, KRW: 750000, ETH: 0.01, GBP: 450},
		},
		{
			Type:         "delegation",
			ID:           5,
			Level:        1040,
			Timestamp:    time.Now().Add(-40 * time.Minute),
			Block:        "B5",
			Hash:         "H5",
			Counter:      5,
			Initiator:    Address{Alias: "Init5", Address: "tz1address5"},
			Sender:       Address{Alias: "Sender5", Address: "tz1sender5"},
			PrevDelegate: Address{Alias: "Delegate9", Address: "tz1delegate9"},
			NewDelegate:  Address{Alias: "Delegate10", Address: "tz1delegate10"},
			Status:       "applied",
			Quote:        Quote{BTC: 0.3, EUR: 3000, USD: 3600, CNY: 24000, JPY: 390000, KRW: 4500000, ETH: 0.06, GBP: 2700},
		},
	}
	return transactions

}

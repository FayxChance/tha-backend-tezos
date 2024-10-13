package model

import "time"

// Delegation struct defines the model for storing delegation data.
// This struct maps the fields received from the Tezos API.
type Delegation struct {
	TzKTID    int64     `json:"-"`         // Internal ID used by the tzKT API, not exposed in the JSON response
	ID        int       `json:"-"`         // Internal database ID, not exposed in the JSON response
	Timestamp time.Time `json:"timestamp"` // Timestamp of the delegation event
	Amount    string    `json:"amount"`    // Amount of XTZ delegated (as a string to avoid precision issues)
	Delegator string    `json:"delegator"` // The address of the delegator (sender)
	Level     string    `json:"level"`     // The block level (height) at which the delegation occurred
}

// Data struct wraps an array of Delegation structs
// This is used for formatting API responses
type Data struct {
	Data []Delegation `json:"data"` // List of delegations to be returned in the API response
}

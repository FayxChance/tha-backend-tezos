package model

import "time"

type Delegation struct {
	Timestamp time.Time `json:"timestamp"`
	Amount    string    `json:"amount"`
	Delegator string    `json:"delegator"`
	Level     string    `json:"level"`
}

type Data struct {
	Data []Delegation `json:"data"`
}

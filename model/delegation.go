package model

import "time"

type Delegation struct {
	TzKTID    int64     `json:"-"`
	ID        int       `json:"-"`
	Timestamp time.Time `json:"timestamp"`
	Amount    string    `json:"amount"`
	Delegator string    `json:"delegator"`
	Level     string    `json:"level"`
}

type Data struct {
	Data []Delegation `json:"data"`
}

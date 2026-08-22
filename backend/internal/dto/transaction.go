package dto

import "time"

type TransactionInput struct {
	Type            string
	Amount          float64
	Category        string
	AccountID       uint
	Title           string
	Note            string
	TransactionDate time.Time
}
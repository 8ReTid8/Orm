package dto

import (
	"backend/internal/models"
	
	"time"
)

type TransactionInput struct {
	Type            string
	Amount          float64
	Category        string
	AccountID       uint
	Title           string
	Note            string
	TransactionDate time.Time
}

type BudgetInput struct {
	Amount    float64
	Category  string
	AccountID uint
	StartDate time.Time
	EndDate   time.Time
}

type BudgetResponse struct {
	ID        uint      `json:"id"`
	AccountID uint      `json:"accountId"`
	Category  string    `json:"category"`
	Amount    float64   `json:"amount"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`

	Spent     float64 `json:"spent"`
	Remaining float64 `json:"remaining"`
	IsActive  bool    `json:"isActive"`

	Account models.Account `json:"account"`
}
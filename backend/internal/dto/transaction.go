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

type TransactionFilter struct {
	StartDate time.Time
	EndDate   time.Time // วันสิ้นสุดแบบ exclusive
	AccountID *uint
	Category  string
	Type      string
}

type TransactionResponse struct {
	ID              uint      `json:"id"`
	AccountID       uint      `json:"accountId"`
	Type            string    `json:"type"`
	Amount          float64   `json:"amount"`
	Category        string    `json:"category"`
	Title           string    `json:"title"`
	Note            string    `json:"note"`
	Image           string    `json:"image"`
	TransactionDate time.Time `json:"transactionDate"`
	CreatedAt       time.Time `json:"createdAt"`
}

func ToTransactionResponse(
	transaction models.Transaction,
) TransactionResponse {
	return TransactionResponse{
		ID:              transaction.ID,
		AccountID:       transaction.AccountID,
		Type:            transaction.Type,
		Amount:          transaction.Amount,
		Category:        transaction.Category,
		Title:           transaction.Title,
		Note:            transaction.Note,
		Image:           transaction.Image,
		TransactionDate: transaction.TransactionDate,
		CreatedAt:       transaction.CreatedAt,
	}
}


package dto

import (
	"backend/internal/models"
	
	"time"
)
type AccountResponse struct {
	ID              uint      `json:"id"`
	UserID          uint      `json:"userId"`
	Name            string    `json:"name"`
	Balance         float64   `json:"balance"`
	CreatedAt       time.Time `json:"createdAt"`
}

func ToAccountResponse(
	account models.Account,
) AccountResponse {
	return AccountResponse{
		ID:              account.ID,
		UserID:          account.UserID,
		Name:            account.Name,
		Balance:         account.Balance,
		CreatedAt:       account.CreatedAt,
	}
}

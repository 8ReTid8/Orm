package dto

import (
	"backend/internal/models"

	"time"
)

type BudgetInput struct {
	Amount    float64
	Category  string
	AccountID uint
	StartDate time.Time
	EndDate   time.Time
}

type BudgetFilter struct {
	Status    string
	Year      *int
	Month     *int
	AccountID *uint
}
type BudgetResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"userId"`
	AccountID uint      `json:"accountId"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	CreatedAt time.Time `json:"createdAt"`
}

type BudgetOverviewResponse struct {
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

func ToBudgetResponse(
	budget models.Budget,
) BudgetResponse {
	return BudgetResponse{
		ID:        budget.ID,
		UserID:    budget.UserID,
		AccountID: budget.AccountID,
		Amount:    budget.Amount,
		Category:  budget.Category,
		StartDate: budget.StartDate,
		EndDate:   budget.EndDate,
		CreatedAt: budget.CreatedAt,
	}
}
func ToBudgetOverviewResponse(
	budget models.Budget,
	spent float64,
	isActive bool,
) BudgetOverviewResponse {
	return BudgetOverviewResponse{
		ID:        budget.ID,
		AccountID: budget.AccountID,
		Category:  budget.Category,
		Amount:    budget.Amount,
		StartDate: budget.StartDate,
		EndDate:   budget.EndDate,

		Spent:     spent,
		Remaining: budget.Amount - spent,
		IsActive:  isActive,

		Account: budget.Account,
	}
}

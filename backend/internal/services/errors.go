package services

import "errors"

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrTransactionNotFound = errors.New("transaction not found")

	ErrInsufficientBalance    = errors.New("insufficient balance")
	ErrInvalidTransactionType = errors.New("invalid transaction type")

	ErrBudgetNotFound     = errors.New("budget not found")
	ErrInvalidBudgetInput = errors.New("invalid budget input")

	ErrInvalidAccountName = errors.New("invalid account name")

	ErrEmailAlreadyUsed   = errors.New("email already used")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
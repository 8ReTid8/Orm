// internal/services/account_service.go
package services

import (
	"context"
	"errors"
	"strings"

	// "backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repositories"

	// "gorm.io/gorm"
)

var (
	// ErrAccountNotFound     = errors.New("account not found")
	// ErrTransactionNotFound = errors.New("transaction not found")
)
var ErrInvalidAccountName = errors.New("invalid account name")

type AccountService struct {
	// db *gorm.DB
	accountRepo *repositories.AccountRepository
}

// func NewAccountService(db *gorm.DB) *AccountService {
// 	return &AccountService{
// 		db: db,
// 	}
// }

func NewAccountService(
	accountRepo *repositories.AccountRepository,
) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}
func (s *AccountService) Create(
	ctx context.Context,
	userID uint,
	name string,
) (*models.Account, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, ErrInvalidAccountName
	}

	account := &models.Account{
		UserID:  userID,
		Name:    name,
		Balance: 0,
	}

	if err := s.accountRepo.Create(ctx, account); err != nil {
		return nil, err
	}

	return account, nil
}
func (s *AccountService) Get(
	ctx context.Context,
	userID uint,
) ([]models.Account, error) {

	// accountRepo := repositories.NewAccountRepository(
	// 	s.db,
	// )

	return s.accountRepo.ListByUser(
		ctx,
		userID,
	)
}
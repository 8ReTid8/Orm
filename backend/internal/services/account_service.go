// internal/services/account_service.go
package services

import (
	"context"
	"errors"
	"strings"

	"backend/internal/models"
	"backend/internal/repositories"

	"gorm.io/gorm"
)

type AccountService struct {
	db              *gorm.DB
	accountRepo     *repositories.AccountRepository
	transactionRepo *repositories.TransactionRepository
}

func NewAccountService(
	db *gorm.DB,
	accountRepo *repositories.AccountRepository,
	transactionRepo *repositories.TransactionRepository,
) *AccountService {
	return &AccountService{
		db:              db,
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
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

	return s.accountRepo.ListByUser(
		ctx,
		userID,
	)
}

func (s *AccountService) Delete(
	ctx context.Context,
	userID uint,
	accountID uint,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		accountRepo := s.accountRepo.WithTx(tx)
		transactionRepo := s.transactionRepo.WithTx(tx)

		account, err := accountRepo.FindOwnedForUpdate(ctx, userID, accountID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAccountNotFound
		}
		if err != nil {
			return err
		}

		if err := transactionRepo.DeleteByAccountID(ctx, account.ID); err != nil {
			return err
		}

		return accountRepo.Delete(ctx, account)
	})
}
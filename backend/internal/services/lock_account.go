package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"context"
)
func lockAccountsForUpdate(
	ctx context.Context,
	accountRepo *repositories.AccountRepository,
	userID uint,
	oldAccountID uint,
	newAccountID uint,
) (*models.Account, *models.Account, error) {
	if oldAccountID == newAccountID {
		account, err := accountRepo.FindOwnedForUpdate(
			ctx,
			userID,
			oldAccountID,
		)
		return account, account, err
	}

	// lock ตาม ID น้อย -> มาก เสมอ เพื่อลดโอกาส deadlock
	if oldAccountID < newAccountID {
		oldAccount, err := accountRepo.FindOwnedForUpdate(
			ctx,
			userID,
			oldAccountID,
		)
		if err != nil {
			return nil, nil, err
		}

		newAccount, err := accountRepo.FindOwnedForUpdate(
			ctx,
			userID,
			newAccountID,
		)
		if err != nil {
			return nil, nil, err
		}

		return oldAccount, newAccount, nil
	}

	newAccount, err := accountRepo.FindOwnedForUpdate(
		ctx,
		userID,
		newAccountID,
	)
	if err != nil {
		return nil, nil, err
	}
	
	oldAccount, err := accountRepo.FindOwnedForUpdate(
		ctx,
		userID,
		oldAccountID,
	)
	if err != nil {
		return nil, nil, err
	}

	return oldAccount, newAccount, nil
}

package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"context"
)
func lockAccountsForUpdate(
	// tx *gorm.DB,
	ctx context.Context,
	accountRepo *repositories.AccountRepository,
	userID uint,
	oldAccountID uint,
	newAccountID uint,
) (*models.Account, *models.Account, error) {
	if oldAccountID == newAccountID {
		// account, err := findAccountForUpdate(
		// 	tx,
		// 	userID,
		// 	oldAccountID,
		// )
		account, err := accountRepo.FindOwnedForUpdate(
			ctx,
			userID,
			oldAccountID,
		)
		return account, account, err
	}

	// lock ตาม ID น้อย -> มาก เสมอ เพื่อลดโอกาส deadlock
	if oldAccountID < newAccountID {
		// oldAccount, err := findAccountForUpdate(
		// 	tx,
		// 	userID,
		// 	oldAccountID,
		// )
		oldAccount, err := accountRepo.FindOwnedForUpdate(
			ctx,
			userID,
			oldAccountID,
		)
		if err != nil {
			return nil, nil, err
		}

		// newAccount, err := findAccountForUpdate(
		// 	tx,
		// 	userID,
		// 	newAccountID,
		// )

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

	// newAccount, err := findAccountForUpdate(
	// 	tx,
	// 	userID,
	// 	newAccountID,
	// )
	newAccount, err := accountRepo.FindOwnedForUpdate(
		ctx,
		userID,
		newAccountID,
	)
	if err != nil {
		return nil, nil, err
	}

	// oldAccount, err := findAccountForUpdate(
	// 	tx,
	// 	userID,
	// 	oldAccountID,
	// )
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

// func findAccountForUpdate(
// 	tx *gorm.DB,
// 	userID uint,
// 	accountID uint,
// ) (*models.Account, error) {
// 	var account models.Account

// 	err := tx.
// 		Clauses(clause.Locking{Strength: "UPDATE"}).
// 		Where(
// 			"id = ? AND user_id = ?",
// 			accountID,
// 			userID,
// 		).
// 		First(&account).
// 		Error

// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, ErrAccountNotFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &account, nil
// }

// func updateBalance(
// 	tx *gorm.DB,
// 	account *models.Account,
// ) error {
// 	return tx.
// 		Model(&models.Account{}).
// 		Where("id = ?", account.ID).
// 		Update("balance", account.Balance).
// 		Error
// }
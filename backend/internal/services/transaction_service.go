// internal/services/transaction_service.go
package services

import (
	"context"
	"errors"

	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repositories"

	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrTransactionNotFound = errors.New("transaction not found")
)

type TransactionService struct {
	db              *gorm.DB
	accountRepo     *repositories.AccountRepository
	transactionRepo *repositories.TransactionRepository
}

func NewTransactionService(
	db *gorm.DB,
	accountRepo *repositories.AccountRepository,
	transactionRepo *repositories.TransactionRepository,
) *TransactionService {
	return &TransactionService{
		db:              db,
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *TransactionService) Create(
	ctx context.Context,
	userID uint,
	input dto.TransactionInput,
	slipPath string,
) (*models.Transaction, error) {
	transaction := &models.Transaction{
		AccountID:       input.AccountID,
		Category:        input.Category,
		Type:            input.Type,
		Amount:          input.Amount,
		Title:           input.Title,
		Note:            input.Note,
		Image:           slipPath,
		TransactionDate: input.TransactionDate,
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		accountRepo := s.accountRepo.WithTx(tx)
		transactionRepo := s.transactionRepo.WithTx(tx)
		// var account models.Account

		// lock บัญชีภายใน transaction
		// err := tx.
		// 	Clauses(clause.Locking{Strength: "UPDATE"}).
		// 	Where(
		// 		"id = ? AND user_id = ?",
		// 		input.AccountID,
		// 		userID,
		// 	).
		// 	First(&account).
		// 	Error
		account, err := accountRepo.FindOwnedForUpdate(
			ctx,
			userID,
			input.AccountID,
		)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAccountNotFound
		}

		if err != nil {
			return err
		}

		newBalance, err := ApplyTransactionToBalance(
			account.Balance,
			input.Type,
			input.Amount,
		)
		if err != nil {
			return err
		}

		account.Balance = newBalance

		// if err := tx.
		// 	Model(&models.Account{}).
		// 	Where("id = ?", account.ID).
		// 	Update("balance", newBalance).
		// 	Error; err != nil {
		// 	return err
		// }

		// // สร้างรายการ
		// if err := tx.Create(transaction).Error; err != nil {
		// 	return err
		// }

		// return nil
		if err := accountRepo.UpdateBalance(ctx, account); err != nil {
			return err
		}

		return transactionRepo.Create(ctx, transaction)
	})

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *TransactionService) Update(
	ctx context.Context,
	userID uint,
	transactionID uint,
	input dto.TransactionInput,
	slipPath string,
) (*models.Transaction, error) {
	var updatedTransaction *models.Transaction

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		accountRepo := s.accountRepo.WithTx(tx)
		transactionRepo := s.transactionRepo.WithTx(tx)

		oldTransaction, err := transactionRepo.FindOwnedForUpdate(
			ctx,
			userID,
			transactionID,
		)

		// var oldTransaction models.Transaction

		// // lock transaction เก่า และเช็ก ownership
		// err := tx.
		// 	Clauses(clause.Locking{
		// 		Strength: "UPDATE",
		// 		Table:    clause.Table{Name: "transactions"},
		// 	}).
		// 	Joins("JOIN accounts ON accounts.id = transactions.account_id").
		// 	Where(
		// 		"transactions.id = ? AND accounts.user_id = ?",
		// 		transactionID,
		// 		userID,
		// 	).
		// 	First(&oldTransaction).
		// 	Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrTransactionNotFound
		}
		if err != nil {
			return err
		}

		// lock account เก่าและ account ใหม่
		// oldAccount, newAccount, err := lockAccountsForUpdate(
		// 	tx,
		// 	userID,
		// 	oldTransaction.AccountID,
		// 	input.AccountID,
		// )
		oldAccount, newAccount, err := lockAccountsForUpdate(
			ctx,
			accountRepo,
			userID,
			oldTransaction.AccountID,
			input.AccountID,
		)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAccountNotFound
		}
		if err != nil {
			return err
		}

		// 1. คืนผลของ transaction เก่า
		oldBalance, err := ReverseTransactionFromBalance(
			oldAccount.Balance,
			oldTransaction.Type,
			oldTransaction.Amount,
		)
		if err != nil {
			return err
		}
		oldAccount.Balance = oldBalance

		// 2. ใช้ผลของ transaction ใหม่
		newBalance, err := ApplyTransactionToBalance(
			newAccount.Balance,
			input.Type,
			input.Amount,
		)
		if err != nil {
			return err
		}
		newAccount.Balance = newBalance

		// account เดิมและใหม่อาจเป็นบัญชีเดียวกัน
		if err := accountRepo.UpdateBalance(ctx, oldAccount); err != nil {
			return err
		}

		if oldAccount.ID != newAccount.ID {
			if err := accountRepo.UpdateBalance(ctx, newAccount); err != nil {
				return err
			}
		}
		// if oldAccount.ID == newAccount.ID {
		// 	if err := updateBalance(tx, oldAccount); err != nil {
		// 		return err
		// 	}
		// } else {
		// 	if err := updateBalance(tx, oldAccount); err != nil {
		// 		return err
		// 	}

		// 	if err := updateBalance(tx, newAccount); err != nil {
		// 		return err
		// 	}
		// }

		// ถ้าไม่ได้ส่งรูปใหม่ ใช้รูปเดิม
		imagePath := oldTransaction.Image
		if slipPath != "" {
			imagePath = slipPath
		}

		oldTransaction.AccountID = input.AccountID
		oldTransaction.Type = input.Type
		oldTransaction.Amount = input.Amount
		oldTransaction.Category = input.Category
		oldTransaction.Title = input.Title
		oldTransaction.Note = input.Note
		oldTransaction.Image = imagePath
		oldTransaction.TransactionDate = input.TransactionDate

		if err := tx.Save(&oldTransaction).Error; err != nil {
			return err
		}

		updatedTransaction = oldTransaction
		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedTransaction, nil
	// return &updatedTransaction, nil
}

func (s *TransactionService) Delete(
	ctx context.Context,
	userID uint,
	transactionID uint,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// var transaction models.Transaction
		
		accountRepo := s.accountRepo.WithTx(tx)

		transactionRepo := s.transactionRepo.WithTx(tx)
		transaction, err := transactionRepo.FindOwnedForUpdate(
			ctx,
			userID,
			transactionID,
		)
		// หาและ lock transaction พร้อมตรวจว่าเป็นของ user นี้
		// err := tx.
		// 	Clauses(clause.Locking{
		// 		Strength: "UPDATE",
		// 		Table:    clause.Table{Name: "transactions"},
		// 	}).
		// 	Joins("JOIN accounts ON accounts.id = transactions.account_id").
		// 	Where(
		// 		"transactions.id = ? AND accounts.user_id = ?",
		// 		transactionID,
		// 		userID,
		// 	).
		// 	First(&transaction).
		// 	Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrTransactionNotFound
		}
		if err != nil {
			return err
		}

		// lock account ที่เกี่ยวข้อง
		// account, err := findAccountForUpdate(
		// 	tx,
		// 	userID,
		// 	transaction.AccountID,
		// )
		account, err := accountRepo.FindOwnedForUpdate(
			ctx,
			userID,
			transaction.AccountID,
		)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAccountNotFound
		}
		if err != nil {
			return err
		}

		// คืนยอดจาก transaction เดิม
		newBalance, err := ReverseTransactionFromBalance(
			account.Balance,
			transaction.Type,
			transaction.Amount,
		)
		if err != nil {
			return err
		}

		account.Balance = newBalance

		// if err := updateBalance(tx, account); err != nil {
		// 	return err
		// }

		// // ลบ transaction เมื่อคืนยอดสำเร็จแล้ว
		// return tx.Delete(&transaction).Error
		if err := accountRepo.UpdateBalance(ctx, account); err != nil {
			return err
		}

		return transactionRepo.Delete(ctx, transaction)
	})
}

func (s *TransactionService) Get(
	ctx context.Context,
	userID uint,
	filter dto.TransactionFilter,
) ([]models.Transaction, error) {
	return s.transactionRepo.FindAllByUser(
		ctx,
		userID,
		filter,
	)
	// return transactionRepo.FindAllByUser(
	// 	ctx,
	// 	userID,
	// 	filter,
	// )
	// query := s.db.WithContext(ctx).
	// 	Joins("JOIN accounts ON accounts.id = transactions.account_id").
	// 	Where("accounts.user_id = ?", userID).
	// 	Where(
	// 		"transactions.transaction_date >= ? AND transactions.transaction_date < ?",
	// 		filter.StartDate,
	// 		filter.EndDate,
	// 	)

	// if filter.AccountID != nil {
	// 	query = query.Where(
	// 		"transactions.account_id = ?",
	// 		*filter.AccountID,
	// 	)
	// }

	// if filter.Category != "" {
	// 	query = query.Where(
	// 		"transactions.category = ?",
	// 		filter.Category,
	// 	)
	// }

	// var transactions []models.Transaction

	// if err := query.
	// 	Order("transactions.created_at DESC").
	// 	Find(&transactions).
	// 	Error; err != nil {
	// 	return nil, err
	// }

	// return transactions, nil
}

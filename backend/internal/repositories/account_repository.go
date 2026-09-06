package repositories

import (
	"context"

	"backend/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(
	db *gorm.DB,
) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) WithTx(
	tx *gorm.DB,
) *AccountRepository {
	return NewAccountRepository(tx)
}

// ใช้ตอน Create / Update / Delete
// ต้อง lock account เพื่อกัน request พร้อมกันแก้ balance
func (r *AccountRepository) FindOwnedForUpdate(
	ctx context.Context,
	userID uint,
	accountID uint,
) (*models.Account, error) {
	var account models.Account

	err := r.db.WithContext(ctx).
		Clauses(clause.Locking{
			Strength: "UPDATE",
		}).
		Where(
			"id = ? AND user_id = ?",
			accountID,
			userID,
		).
		First(&account).
		Error

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepository) FindOwned(
	ctx context.Context,
	userID uint,
	accountID uint,
) (*models.Account, error) {
	var account models.Account

	err := r.db.WithContext(ctx).
		Where(
			"id = ? AND user_id = ?",
			accountID,
			userID,
		).
		First(&account).
		Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepository) UpdateBalance(
	ctx context.Context,
	account *models.Account,
) error {
	return r.db.WithContext(ctx).
		Model(&models.Account{}).
		Where(
			"id = ? AND user_id = ?",
			account.ID,
			account.UserID,
		).
		Update(
			"balance",
			account.Balance,
		).
		Error
}

func (r *AccountRepository) Create(
	ctx context.Context,
	account *models.Account,
) error {
	return r.db.WithContext(ctx).
		Create(account).
		Error
}

func (r *AccountRepository) ListByUser(
	ctx context.Context,
	userID uint,
) ([]models.Account, error) {

	var accounts []models.Account

	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at ASC").
		Find(&accounts).
		Error; err != nil {
		return nil, err
	}

	return accounts, nil
}

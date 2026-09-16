package repositories

import (
	"context"

	"backend/internal/dto"
	"backend/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(
	db *gorm.DB,
) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) WithTx(
	tx *gorm.DB,
) *TransactionRepository {
	return NewTransactionRepository(tx)
}

func (r *TransactionRepository) Create(
	ctx context.Context,
	transaction *models.Transaction,
) error {
	return r.db.WithContext(ctx).
		Create(transaction).
		Error
}

// หา transaction พร้อมเช็ก ownership และ lock แถวไว้
func (r *TransactionRepository) FindOwnedForUpdate(
	ctx context.Context,
	userID uint,
	transactionID uint,
) (*models.Transaction, error) {
	var transaction models.Transaction

	err := r.db.WithContext(ctx).
		Clauses(clause.Locking{
			Strength: "UPDATE",
			Table:    clause.Table{Name: "transactions"},
		}).
		Joins(
			"JOIN accounts ON accounts.id = transactions.account_id",
		).
		Where(
			"transactions.id = ? AND accounts.user_id = ?",
			transactionID,
			userID,
		).
		First(&transaction).
		Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) Update(
	ctx context.Context,
	transaction *models.Transaction,
) error {
	return r.db.WithContext(ctx).
		Save(transaction).
		Error
}

func (r *TransactionRepository) Delete(
	ctx context.Context,
	transaction *models.Transaction,
) error {
	return r.db.WithContext(ctx).
		Delete(transaction).
		Error
}

func (r *TransactionRepository) FindAllByUser(
	ctx context.Context,
	userID uint,
	filter dto.TransactionFilter,
) ([]models.Transaction, error) {
	query := r.db.WithContext(ctx).
		Joins(
			"JOIN accounts ON accounts.id = transactions.account_id",
		).
		Where("accounts.user_id = ?", userID).
		Where(
			"transactions.transaction_date >= ? AND transactions.transaction_date < ?",
			filter.StartDate,
			filter.EndDate,
		)

	if filter.AccountID != nil {
		query = query.Where(
			"transactions.account_id = ?",
			*filter.AccountID,
		)
	}

	if filter.Type != "" {
		query = query.Where(
			"transactions.type = ?",
			filter.Type,
		)
	}

	if filter.Category != "" {
		query = query.Where(
			"transactions.category = ?",
			filter.Category,
		)
	}

	var transactions []models.Transaction

	if err := query.
		Order("transactions.created_at DESC").
		Find(&transactions).
		Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) GetAvailableYears(
	ctx context.Context,
	userID uint,
) ([]int, error) {

	var years []int

	err := r.db.WithContext(ctx).
		Raw(`
			SELECT DISTINCT EXTRACT(YEAR FROM transactions.transaction_date)::int AS year
			FROM transactions
			JOIN accounts ON accounts.id = transactions.account_id
			WHERE accounts.user_id = ?
			ORDER BY year DESC
		`, userID).
		Scan(&years).Error

	if err != nil {
		return nil, err
	}

	return years, nil
}

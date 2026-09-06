package repositories

import (
	"context"
	"time"

	"backend/internal/dto"
	"backend/internal/models"

	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type BudgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(
	db *gorm.DB,
) *BudgetRepository {
	return &BudgetRepository{
		db: db,
	}
}

func (r *BudgetRepository) Create(
	ctx context.Context,
	budget *models.Budget,
) error {
	return r.db.WithContext(ctx).
		Create(budget).
		Error
}

func (r *BudgetRepository) Update(
	ctx context.Context,
	budget *models.Budget,
) error {
	return r.db.WithContext(ctx).
		Save(budget).
		Error
}

func (r *BudgetRepository) Delete(
	ctx context.Context,
	budget *models.Budget,
) error {
	return r.db.WithContext(ctx).
		Delete(budget).
		Error
}

func (r *BudgetRepository) ListByUser(
	ctx context.Context,
	userID uint,
	filter dto.BudgetFilter,
	today time.Time,
) ([]models.Budget, error) {
	query := r.db.WithContext(ctx).
		Model(&models.Budget{}).
		Preload("Account").
		Where("budgets.user_id = ?", userID)

	switch filter.Status {
	case "active":
		query = query.Where(
			"budgets.start_date <= ? AND budgets.end_date >= ?",
			today,
			today,
		)

	case "ended":
		query = query.Where(
			"budgets.end_date < ?",
			today,
		)

		// เหมือน filter เดิม: year/month ใช้กับ ended
		if filter.Year != nil {
			var rangeStart time.Time
			var rangeEnd time.Time

			if filter.Month != nil {
				rangeStart = time.Date(
					*filter.Year,
					time.Month(*filter.Month),
					1,
					0, 0, 0, 0,
					today.Location(),
				)

				rangeEnd = rangeStart.AddDate(0, 1, 0)
			} else {
				rangeStart = time.Date(
					*filter.Year,
					1,
					1,
					0, 0, 0, 0,
					today.Location(),
				)

				rangeEnd = rangeStart.AddDate(1, 0, 0)
			}

			// budget ต้อง overlap กับช่วงที่เลือก
			query = query.Where(
				"budgets.start_date < ? AND budgets.end_date >= ?",
				rangeEnd,
				rangeStart,
			)
		}
	}

	if filter.AccountID != nil {
		query = query.Where(
			"budgets.account_id = ?",
			*filter.AccountID,
		)
	}

	var budgets []models.Budget

	if err := query.
		Order("budgets.start_date ASC").
		Find(&budgets).
		Error; err != nil {
		return nil, err
	}

	return budgets, nil
}

func (r *BudgetRepository) SumExpense(
	ctx context.Context,
	budget *models.Budget,
) (float64, error) {
	var spent float64

	err := r.db.WithContext(ctx).
		Model(&models.Transaction{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("account_id = ?", budget.AccountID).
		Where("category = ?", budget.Category).
		Where("type = ?", "expense").
		Where("transaction_date >= ?", budget.StartDate).
		Where("transaction_date <= ?", budget.EndDate).
		Scan(&spent).
		Error

	if err != nil {
		return 0, err
	}

	return spent, nil
}
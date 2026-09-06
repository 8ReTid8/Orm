package services

import (
	"context"
	"errors"
	"time"

	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repositories"

	"gorm.io/gorm"
)

// var ErrBudgetNotFound = errors.New("budget not found")

type BudgetSummary struct {
	Budget   models.Budget
	Spent    float64
	IsActive bool
}

type BudgetService struct {
	budgetRepo  *repositories.BudgetRepository
	accountRepo *repositories.AccountRepository
}

func NewBudgetService(
	budgetRepo *repositories.BudgetRepository,
	accountRepo *repositories.AccountRepository,
) *BudgetService {
	return &BudgetService{
		budgetRepo:  budgetRepo,
		accountRepo: accountRepo,
	}
}

func (s *BudgetService) Create(
	ctx context.Context,
	userID uint,
	input dto.BudgetInput,
) (*models.Budget, error) {

	account, err := s.accountRepo.FindOwned(
		ctx,
		userID,
		input.AccountID,
	)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	budget := &models.Budget{
		UserID:    userID,
		AccountID: input.AccountID,
		Amount:    input.Amount,
		Category:  input.Category,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	}

	if err := s.budgetRepo.Create(ctx, budget); err != nil {
		return nil, err
	}

	// ให้ handler สามารถสร้าง BudgetResponse พร้อม Account ได้
	budget.Account = *account

	return budget, nil
}

func (s *BudgetService) Update(
	ctx context.Context,
	userID uint,
	budgetID uint,
	input dto.BudgetInput,
) (*models.Budget, error) {
	budget, err := s.budgetRepo.FindOwned(
		ctx,
		userID,
		budgetID,
	)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrBudgetNotFound
	}
	if err != nil {
		return nil, err
	}

	// เช็กว่า account ใหม่เป็นของ user
	account, err := s.accountRepo.FindOwned(
		ctx,
		userID,
		input.AccountID,
	)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	budget.AccountID = input.AccountID
	budget.Amount = input.Amount
	budget.Category = input.Category
	budget.StartDate = input.StartDate
	budget.EndDate = input.EndDate

	if err := s.budgetRepo.Update(ctx, budget); err != nil {
		return nil, err
	}

	// ใช้เฉพาะเพื่อ response ไม่ได้บันทึก account ซ้ำ
	budget.Account = *account

	return budget, nil
}

func (s *BudgetService) Get(
	ctx context.Context,
	userID uint,
	filter dto.BudgetFilter,
) ([]BudgetSummary, error) {
	now := time.Now()

	// Budget ใช้วันที่ จึงตัดเวลาออก
	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	budgets, err := s.budgetRepo.ListByUser(
		ctx,
		userID,
		filter,
		today,
	)
	if err != nil {
		return nil, err
	}

	summaries := make(
		[]BudgetSummary,
		0,
		len(budgets),
	)

	for index := range budgets {
		budget := &budgets[index]

		spent, err := s.budgetRepo.SumExpense(ctx, budget)
		if err != nil {
			return nil, err
		}

		// isActive := !today.Before(budget.StartDate) &&
		// 	!today.After(budget.EndDate)
		isActive := !today.After(budget.EndDate)

		summaries = append(
			summaries,
			BudgetSummary{
				Budget:   *budget,
				Spent:    spent,
				IsActive: isActive,
			},
		)
	}

	return summaries, nil
}

func (s *BudgetService) Delete(
	ctx context.Context,
	userID uint,
	budgetID uint,
) error {
	budget, err := s.budgetRepo.FindOwned(
		ctx,
		userID,
		budgetID,
	)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrBudgetNotFound
	}
	if err != nil {
		return err
	}

	return s.budgetRepo.Delete(ctx, budget)
}

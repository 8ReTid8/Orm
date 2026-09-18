package services

import (
	"context"
	"time"

	"backend/internal/dto"
	"backend/internal/repositories"
)

type SummaryService struct {
	transactionRepo *repositories.TransactionRepository
}

func NewSummaryService(
	transactionRepo *repositories.TransactionRepository,
) *SummaryService {
	return &SummaryService{
		transactionRepo: transactionRepo,
	}
}

func (s *SummaryService) Get(
	ctx context.Context,
	userID uint,
	filter dto.SummaryFilter,
) (*dto.SummaryResult, error) {
	location := time.Local

	yearStart := time.Date(
		filter.Year,
		time.January,
		1,
		0, 0, 0, 0,
		location,
	)

	yearEnd := yearStart.AddDate(1, 0, 0)

	// ช่วงที่ใช้กับ cards และ pie chart
	startDate := yearStart
	endDate := yearEnd

	if filter.Month != nil {
		startDate = time.Date(
			filter.Year,
			time.Month(*filter.Month),
			1,
			0, 0, 0, 0,
			location,
		)

		endDate = startDate.AddDate(0, 1, 0)
	}

	totals, err := s.transactionRepo.GetSummaryTotals(
		ctx,
		userID,
		startDate,
		endDate,
		filter.AccountID,
	)
	if err != nil {
		return nil, err
	}

	incomeByCategory, err := s.transactionRepo.ListCategoryTotals(
		ctx,
		userID,
		"income",
		startDate,
		endDate,
		filter.AccountID,
	)
	if err != nil {
		return nil, err
	}

	expenseByCategory, err := s.transactionRepo.ListCategoryTotals(
		ctx,
		userID,
		"expense",
		startDate,
		endDate,
		filter.AccountID,
	)
	if err != nil {
		return nil, err
	}

	comparisonPeriod := "month"
	comparisonLength := 12
	var comparisonRows []dto.ComparisonTotal

	if filter.Month != nil {
		comparisonPeriod = "day"
		comparisonLength = endDate.AddDate(0, 0, -1).Day()
		comparisonRows, err = s.transactionRepo.ListDailyTotals(ctx, userID, startDate, endDate, filter.AccountID)
	} else {
		comparisonRows, err = s.transactionRepo.ListMonthlyTotals(ctx, userID, yearStart, yearEnd, filter.AccountID)
	}
	if err != nil {
		return nil, err
	}

	// ส่งทุกวันหรือเดือนให้ chart แม้ช่วงนั้นไม่มีรายการ
	comparison := make([]dto.ComparisonTotal, comparisonLength)
	for period := 1; period <= comparisonLength; period++ {
		comparison[period-1] = dto.ComparisonTotal{Period: period}
	}
	for _, row := range comparisonRows {
		comparison[row.Period-1] = row
	}

	netBalance := totals.Income - totals.Expense

	savingsRate := float64(0)

	if totals.Income > 0 {
		savingsRate = (netBalance / totals.Income) * 100

		if savingsRate < 0 {
			savingsRate = 0
		}
	}

	return &dto.SummaryResult{
		TotalIncome:       totals.Income,
		TotalExpense:      totals.Expense,
		NetBalance:        netBalance,
		SavingsRate:       savingsRate,
		IncomeByCategory:  incomeByCategory,
		ExpenseByCategory: expenseByCategory,
		ComparisonPeriod:  comparisonPeriod,
		Comparison:        comparison,
	}, nil
}

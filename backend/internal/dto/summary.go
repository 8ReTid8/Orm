package dto

type SummaryFilter struct {
	Year      int
	Month     *int
	AccountID *uint
}

type ComparisonFilter struct {
    Year      int
	Month     *int
    AccountID *uint
    // Type      *string
    Category  *string
}

type ComparisonTotal struct {
	Period  int     `json:"period"`
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
}
type ComparisonResult struct {
    Period string
    Items  []ComparisonTotal
}
type SummaryTotals struct {
	Income  float64
	Expense float64
}

type CategoryTotal struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type SummaryResult struct {
	TotalIncome       float64
	TotalExpense      float64
	NetBalance        float64
	SavingsRate       float64
	IncomeByCategory  []CategoryTotal
	ExpenseByCategory []CategoryTotal
	// ComparisonPeriod  string
	// Comparison        []ComparisonTotal
}

type SummaryResponse struct {
	TotalIncome       float64           `json:"totalIncome"`
	TotalExpense      float64           `json:"totalExpense"`
	NetBalance        float64           `json:"netBalance"`
	SavingsRate       float64           `json:"savingsRate"`
	IncomeByCategory  []CategoryTotal   `json:"incomeByCategory"`
	ExpenseByCategory []CategoryTotal   `json:"expenseByCategory"`
	// ComparisonPeriod  string            `json:"comparisonPeriod"`
	// Comparison        []ComparisonTotal `json:"comparison"`
}

type ComparisonResponse struct {
    Period string            `json:"period"`
    Items  []ComparisonTotal `json:"items"`
}
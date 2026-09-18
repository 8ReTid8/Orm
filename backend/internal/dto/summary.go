package dto

type SummaryFilter struct {
	Year      int
	Month     *int
	AccountID *uint
}

type SummaryTotals struct {
	Income  float64
	Expense float64
}

type CategoryTotal struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type ComparisonTotal struct {
	Period  int     `json:"period"`
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
}

type SummaryResult struct {
	TotalIncome       float64
	TotalExpense      float64
	NetBalance        float64
	SavingsRate       float64
	IncomeByCategory  []CategoryTotal
	ExpenseByCategory []CategoryTotal
	ComparisonPeriod  string
	Comparison        []ComparisonTotal
}

type SummaryResponse struct {
	TotalIncome       float64           `json:"totalIncome"`
	TotalExpense      float64           `json:"totalExpense"`
	NetBalance        float64           `json:"netBalance"`
	SavingsRate       float64           `json:"savingsRate"`
	IncomeByCategory  []CategoryTotal   `json:"incomeByCategory"`
	ExpenseByCategory []CategoryTotal   `json:"expenseByCategory"`
	ComparisonPeriod  string            `json:"comparisonPeriod"`
	Comparison        []ComparisonTotal `json:"comparison"`
}

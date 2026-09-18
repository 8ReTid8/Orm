
export interface CategoryTotal{
    name: string
    amount: number
}
export type ComparisonPeriod = "month" | "day"

export interface ComparisonTotal{
    period: number
    income: number
    expense: number
}
export interface Summary{
    totalIncome: number
    totalExpense: number
    netBalance: number
    savingsRate: number
    incomeByCategory: CategoryTotal[]
    expenseByCategory: CategoryTotal[]
    comparisonPeriod: ComparisonPeriod
    comparison: ComparisonTotal[]
}

export interface Budget {
  id: number
  accountId: number
  category: string
  amount: number
  startDate: string // "YYYY-MM-DD"
  endDate: string   // "YYYY-MM-DD"
  spent: number
}

export interface BudgetForm {
  category: string
  amount: number | null
  accountId: number | null
  startDate: string // "YYYY-MM-DD"
  endDate: string   // "YYYY-MM-DD"
}

export interface CreateBudgetResponse {
  message: string
  budget: Budget
}

export interface GetBudgetsResponse {
  budgets: Budget[]
}

export interface DeleteBudgetResponse {
  message: string
}

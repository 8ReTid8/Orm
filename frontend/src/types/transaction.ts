export type TransactionType = "income" | "expense"

export interface TransactionForm {
  type: TransactionType
  amount: number | null
  category: string
  title: string
  note: string
  transactionDate: string
}
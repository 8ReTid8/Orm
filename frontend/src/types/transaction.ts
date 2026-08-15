export type TransactionType = "income" | "expense"

export interface TransactionForm {
  type: TransactionType
  amount: number | null
  category: string
  bankId: number | null
  title: string
  note: string
  transactionDate: string
  slipImage: File | null
}

export interface Transaction {
  id: number
  type: TransactionType
  amount: number
  category: string
  bankName: string
  title: string
  note: string
  transactionDate: string
  slipImage?: string | null
}


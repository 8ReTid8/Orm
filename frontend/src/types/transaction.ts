import type { Account } from "./account"

export type TransactionType = "income" | "expense"

export interface TransactionForm {
  type: TransactionType
  amount: number | null
  category: string
  accountId: number | null
  // title: string
  note: string
  transactionDate: string
  slipImage: File | null
}

export interface Transaction {
  id: number
  // account: Account
  accountId: number
  type: TransactionType
  amount: number
  category: string
  bankName: string
  // title: string
  note: string
  transactionDate: string
  slipImage?: string | null
}

export interface TransactionFilter {
    accountId: number | null
    year?: number | null
    month?: number | null
    startDate?: string | null
    endDate?: string | null
    category?: string | null
}

export interface CreateTransactionResponse {
  message: string
  transaction: Transaction
}

export interface GetTransactionsResponse {
  transactions: Transaction[]
}

export interface DeleteTransactionResponse {
  message: string
}

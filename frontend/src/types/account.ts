import type { PageMeta } from "./pagination"
import type { Transaction } from "./transaction"

export interface Account {
  id: number
  userId: number
  name: string
  balance: number
}

export interface CreateAccountResponse {
  message: string
  account: Account
}

export interface CreateAccountInput {
  name: string
}

export interface DeleteAccountResponse {
  message: string
}



export interface AccountTransactionSummary {
  income: number
  expense: number
  netFlow: number
}

export interface AccountTransactionsResponse {
  transactions: Transaction[]
  meta: PageMeta
  summary: AccountTransactionSummary
}
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
import api from "./api"
import type { Account } from "@/types/account"

export interface CreateAccountInput {
  name: string
}

interface CreateAccountResponse {
  message: string
  account: Account
}

export async function createAccount(
  input: CreateAccountInput,
) {
  const response =
    await api.post<CreateAccountResponse>(
      "/accounts",
      input,
    )

  return response.data
}

export async function getAccounts() {
  const response =
    await api.get<Account[]>("/accounts")

  return response.data
}
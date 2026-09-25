import api from "./api"
import type { Account, AccountTransactionsResponse, CreateAccountInput, CreateAccountResponse, DeleteAccountResponse } from "@/types/account"

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
  const response = await api.get<Account[]>("/accounts")
  return response.data
}

export async function deleteAccount(
  id: number
) {
  const response = await api.delete<DeleteAccountResponse>(
    `/accounts/${id}`,
  )

  return response.data
}

export async function getAccountTransactions(
  accountId: number,
  year: number,
  month: number | null,
  page: number,
  limit = 20,
) {
  const response = await api.get<AccountTransactionsResponse>(
    `/accounts/${accountId}/transactions`,
    {
      params: {
        year,
        month,
        page,
        limit,
      },
    },
  )

  return response.data
}
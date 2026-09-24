import api from "./api"

import type {
  TransactionForm,
  CreateTransactionResponse,
  GetTransactionsResponse,
  DeleteTransactionResponse,
} from "@/types/transaction"


export async function createTransaction(
  form: TransactionForm,
) {
  console.log("CREATE TRANSACTION SERVICE CALLED")
  const payload = new FormData()

  payload.append("type", form.type)
  payload.append("amount", String(form.amount))
  payload.append("category", form.category)
  payload.append("accountId", String(form.accountId))
  // payload.append("title", form.title)
  payload.append("note", form.note)
  payload.append("transactionDate",form.transactionDate)

  if (form.slipImage) {
    payload.append(
      "slipImage",
      form.slipImage,
    )
  }
  for (const [key, value] of payload.entries()) {
    console.log(key, value)
  }
  const response =
    await api.post<CreateTransactionResponse>(
      "/transactions",
      payload,
    )

  return response.data
}

export async function updateTransaction(
  id: number,
  form: TransactionForm,
) {
  const payload = new FormData()
  console.log("UPDATE TRANSACTION SERVICE CALLED")
  payload.append("type", form.type)
  payload.append("amount", String(form.amount))
  payload.append("category", form.category)
  payload.append("accountId", String(form.accountId),)
  // payload.append("title", form.title)
  payload.append("note", form.note)
  payload.append("transactionDate",form.transactionDate)

  if (form.slipImage) {
    payload.append(
      "slipImage",
      form.slipImage,
    )
  }

  const response =
    await api.patch<CreateTransactionResponse>(
      `/transactions/${id}`,
      payload,
    )

  return response.data
}

export async function getTransactions(
  year: number | null,
  month: number | null,
  startDate: string | null,
  endDate: string | null,
  accountId: number | null,
  category: string | null
) {
  const response =
    await api.get<GetTransactionsResponse>(
      "/transactions",
      {
        params: {
          startDate,
          endDate,
          accountId,
          category,
          year,
          month,
        },
      },
    )

  return response.data.transactions
}

export async function deleteTransaction(
  id: number
) {
  const response = await api.delete<DeleteTransactionResponse>(
    `/transactions/${id}`,
  )

  return response.data
}

export async function getTransactionYears(): Promise<number[]> {
  const response = await api.get<{ years: number[] }>("/transactions/years")
  return response.data.years
}
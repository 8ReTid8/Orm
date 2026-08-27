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
  const payload = new FormData()

  payload.append("type", form.type)
  payload.append("amount", String(form.amount))
  payload.append("category", form.category)
  payload.append("accountId", String(form.accountId))
  payload.append("title", form.title)
  payload.append("note", form.note)
  payload.append(
    "transactionDate",
    form.transactionDate,
  )

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

  payload.append("type", form.type)
  payload.append("amount", String(form.amount))
  payload.append("category", form.category)
  payload.append("accountId",String(form.accountId),)
  payload.append("title", form.title)
  payload.append("note", form.note)
  payload.append(
    "transactionDate",
    form.transactionDate,
  )

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
  year: number,
  month: number,
) {
  const response =
    await api.get<GetTransactionsResponse>(
      "/transactions",
      {
        params: {
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
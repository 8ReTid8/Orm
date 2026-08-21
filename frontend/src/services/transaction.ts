import api from "./api"

import type {
  TransactionForm,
  Transaction,
} from "@/types/transaction"

interface CreateTransactionResponse {
  message: string
  transaction: Transaction
}

interface GetTransactionsResponse {
  transactions: Transaction[]
}


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
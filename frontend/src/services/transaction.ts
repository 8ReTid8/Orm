import api from "./api"

import type {
  TransactionForm,
  Transaction,
} from "@/types/transaction"

interface CreateTransactionResponse {
  message: string
  transaction: Transaction
}

export async function createTransaction(
  form: TransactionForm,
) {
  const payload = new FormData()

  payload.append("type", form.type)
  payload.append("amount", String(form.amount))
  payload.append("category", form.category)
  payload.append("bankId", String(form.bankId))
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
    await api.post<CreateTransactionResponse>(
      "/transactions",
      payload,
    )

  return response.data
}
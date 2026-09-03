
import { ref } from "vue"

import type {
  Transaction,
  TransactionForm,
} from "@/types/transaction"

export function useTransactionForm() {
  const form = ref<TransactionForm>({
    type: "expense",
    amount: null,
    category: "",
    accountId: null,
    title: "",
    note: "",
    transactionDate: "",
    slipImage: null,
  })

  function resetForm(date: string) {
    form.value = {
      type: "expense",
      amount: null,
      category: "",
      accountId: null,
      title: "",
      note: "",
      transactionDate: date,
      slipImage: null,
    }
  }

  function setEditForm(transaction: Transaction) {
    form.value = {
      type: transaction.type,
      amount: transaction.amount,
      category: transaction.category,
      accountId: transaction.accountId,
      title: transaction.title,
      note: transaction.note,
      transactionDate: transaction.transactionDate.slice(0, 10),
      slipImage: null,
    }
  }

  return {
    form,
    resetForm,
    setEditForm,
  }
}
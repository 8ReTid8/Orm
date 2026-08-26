import { ref } from "vue"

import type {
  Transaction,
  TransactionForm,
} from "@/types/transaction"

import {
  createTransaction,
  getTransactions,
  updateTransaction,
} from "@/services/transaction"

export function useTransactions() {
  const transactions = ref<Transaction[]>([])
  const isLoadingTransactions = ref(false)
  const editingTransactionId = ref<number | null>(null)

  async function loadTransactions(
    year: number,
    month: number,
  ) {
    try {
      isLoadingTransactions.value = true

      transactions.value = await getTransactions(
        year,
        month,
      )
    } catch (error) {
      console.error(
        "Failed to load transactions:",
        error,
      )
    } finally {
      isLoadingTransactions.value = false
    }
  }

  async function saveTransaction(
    form: TransactionForm,
  ) {
    if (
      !form.amount ||
      !form.title ||
      !form.category ||
      !form.accountId
    ) {
      return
    }

    if (editingTransactionId.value !== null) {
      await updateTransaction(
        editingTransactionId.value,
        form,
      )
    } else {
      await createTransaction(form)
    }
  }

  function startEdit(transaction: Transaction) {
    editingTransactionId.value = transaction.id
  }

  function cancelEdit() {
    editingTransactionId.value = null
  }

  return {
    transactions,
    isLoadingTransactions,
    editingTransactionId,

    loadTransactions,
    saveTransaction,
    startEdit,
    cancelEdit,
  }
}
import { ref } from "vue"

import type {
  Transaction,
  TransactionFilter,
  TransactionForm,
} from "@/types/transaction"

import {
  createTransaction,
  getTransactions,
  updateTransaction,
  deleteTransaction as deleteTransactionApi,
} from "@/services/transaction"

export function useTransactions() {
  const transactions = ref<Transaction[]>([])
  const isLoadingTransactions = ref(false)
  const editingTransactionId = ref<number | null>(null)

  async function loadTransactions({
    year = null,
    month = null,
    startDate = null,
    endDate = null,
    accountId = null,
    category = null
  }: TransactionFilter) {
    try {
      isLoadingTransactions.value = true

      transactions.value = await getTransactions(
        year,
        month,
        startDate,
        endDate,
        accountId,
        category
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
    console.log("FORM:", form)
    console.log(
      "EDITING ID:",
      editingTransactionId.value,
    )

    if (
      !form.amount ||
      !form.title ||
      !form.category ||
      !form.accountId
    ) {
      console.log("VALIDATION FAILED")
      return
    }

    if (editingTransactionId.value !== null) {
      await updateTransaction(
        editingTransactionId.value,
        form,
      )
    } else {
      console.log("CREATE")
      await createTransaction(form)
    }
  }

  async function deleteTransaction(id: number) {
    try {
      await deleteTransactionApi(id)
    } catch (error) {
      console.error("Delete transaction failed:", error)
      throw error
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
    deleteTransaction,
    saveTransaction,
    startEdit,
    cancelEdit,
  }
}
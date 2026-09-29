import { ref } from "vue"

import type {
  Transaction,
  TransactionFilter,
  TransactionForm,
  TransactionTotals,
} from "@/types/transaction"

import {
  createTransaction,
  getTransactions,
  updateTransaction,
  deleteTransaction as deleteTransactionApi,
  getTransactionPage,
} from "@/services/transaction"
import type { PageMeta } from "@/types/pagination"
import { getAccountTransactions } from "@/services/account"

export function useTransactions() {
  const transactions = ref<Transaction[]>([])
  const isLoadingTransactions = ref(false)
  const editingTransactionId = ref<number | null>(null)
  const error = ref<string | null>(null)

  const meta = ref<PageMeta>({
    page: 1,
    limit: 20,
    total: 0,
    totalPages: 0,
  })

  const totals = ref<TransactionTotals>({
    income: 0,
    expense: 0,
    netFlow: 0,
  })

  async function loadTransactions({
    year = null,
    month = null,
    startDate = null,
    endDate = null,
    accountId = null,
    category = null,
  }: TransactionFilter) {
    try {
      isLoadingTransactions.value = true
      error.value = null

      transactions.value = await getTransactions(
        year,
        month,
        startDate,
        endDate,
        accountId,
        category,
      )
    } catch (err) {
      console.error("Failed to load transactions:", err)
      error.value = "ไม่สามารถโหลดรายการธุรกรรมได้"
      throw err
    } finally {
      isLoadingTransactions.value = false
    }
  }

  async function loadTransactionPage(
    filter: TransactionFilter,
    page = 1,
    limit = 5,
  ) {
    try {
      isLoadingTransactions.value = true
      error.value = null

      const response = await getTransactionPage(
        filter,
        page,
        limit,
      )

      transactions.value = response.transactions
      meta.value = response.meta
    } catch (err) {
      console.error("Failed to load transaction page:", err)
      error.value = "ไม่สามารถโหลดหน้ารายการธุรกรรมได้"
      throw err
    } finally {
      isLoadingTransactions.value = false
    }
  }

  async function loadAccountTransactions(
    accountId: number,
    year: number,
    month: number | null,
    page = 1,
    limit = 20,
  ) {
    try {
      isLoadingTransactions.value = true
      error.value = null

      const response = await getAccountTransactions(
        accountId,
        year,
        month,
        page,
        limit,
      )

      transactions.value = response.transactions
      meta.value = response.meta
      totals.value = response.summary
    } catch (err) {
      console.error("Failed to load account transactions:", err)
      error.value = "ไม่สามารถโหลดรายการธุรกรรมของบัญชีได้"
      throw err
    } finally {
      isLoadingTransactions.value = false
    }
  }

  async function saveTransaction(form: TransactionForm) {
    if (!form.amount || !form.category || !form.accountId || !form.transactionDate) {
      const validationError = new Error("กรุณากรอกข้อมูลธุรกรรมให้ครบถ้วน")
      console.error("Validation failed:", validationError.message)
      throw validationError
    }

    try {
      error.value = null
      if (editingTransactionId.value !== null) {
        const updated = await updateTransaction(
          editingTransactionId.value,
          form,
        )
        editingTransactionId.value = null
        return updated
      } else {
        return await createTransaction(form)
      }
    } catch (err) {
      console.error("Failed to save transaction:", err)
      error.value = "ไม่สามารถบันทึกรายการธุรกรรมได้"
      throw err
    }
  }

  async function deleteTransaction(id: number) {
    try {
      error.value = null
      await deleteTransactionApi(id)
      // transactions.value = transactions.value.filter((t) => t.id !== id)
    } catch (err) {
      console.error("Delete transaction failed:", err)
      error.value = "ไม่สามารถลบรายการธุรกรรมได้"
      throw err
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
    meta,
    totals,
    isLoadingTransactions,
    editingTransactionId,
    error,

    loadTransactions,
    loadTransactionPage,
    loadAccountTransactions,
    deleteTransaction,
    saveTransaction,
    startEdit,
    cancelEdit,
  }
}

// Alias for naming consistency
// export const useTransaction = useTransactions
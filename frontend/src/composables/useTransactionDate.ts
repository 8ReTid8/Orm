import { computed, ref } from "vue"
import type { Transaction } from "@/types/transaction"

export function useTransactionDate(
  transactions: {
    value: Transaction[]
  },
) {
  const selectedDate = ref<Date>(new Date())

  function formatDate(date: Date): string {
    const year = date.getFullYear()
    const month = String(
      date.getMonth() + 1,
    ).padStart(2, "0")
    const day = String(
      date.getDate(),
    ).padStart(2, "0")

    return `${year}-${month}-${day}`
  }

  const selectedDateText = computed(() => {
    return selectedDate.value.toLocaleDateString(
      "th-TH",
      {
        day: "numeric",
        month: "long",
        year: "numeric",
      },
    )
  })

  const selectedDayTransactions = computed(() => {
    const date = formatDate(selectedDate.value)

    return transactions.value.filter(
      transaction =>
        transaction.transactionDate.startsWith(date),
    )
  })

  function selectDate(date: Date) {
    selectedDate.value = date
  }

  return {
    selectedDate,
    selectedDateText,
    selectedDayTransactions,
    formatDate,
    selectDate,
  }
}
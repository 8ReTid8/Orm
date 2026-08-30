import { computed, ref } from "vue"
import type { Transaction } from "@/types/transaction"
import { formatDate } from "@/utils/format"

export function useTransactionFilter(
  transactions: {
    value: Transaction[]
  },
) {
  const selectedDate = ref<Date>(new Date())
  const selectedAccountId = ref<number | null>(null)
  const selectedCategory = ref<string | null>(null)
  // const filteredTransactions = computed(() => {
  //   if (selectedAccountId.value === null) {
  //     return transactions.value
  //   }

  //   return transactions.value.filter(
  //     transaction =>
  //       transaction.account.id === selectedAccountId.value
  //   )
  // })
  const filteredTransactions = computed(() => {
    return transactions.value.filter(transaction => {

      const accountMatch =
        selectedAccountId.value === null ||
        transaction.account.id === selectedAccountId.value

      const categoryMatch =
        selectedCategory.value === null ||
        transaction.category === selectedCategory.value

      return accountMatch && categoryMatch
    })
  })

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

    // return transactions.value.filter(
    return filteredTransactions.value.filter(
      transaction =>
        transaction.transactionDate.startsWith(date),
    )
  })

  function selectDate(date: Date) {
    selectedDate.value = date
  }

  return {
    selectedDate,
    selectedAccountId,
    selectedCategory,
    filteredTransactions,
    selectedDayTransactions,
    selectedDateText,
    selectDate,
  }
}
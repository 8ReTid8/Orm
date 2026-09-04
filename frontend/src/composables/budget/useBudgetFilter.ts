import { ref, computed, watch } from "vue"
import type { Ref } from "vue"

type BudgetStatus = "active" | "ended"

interface LoadBudgetParams {
  status: BudgetStatus
  accountId: number | null
  year: number | null
  month: number | null
}

export function useBudgetFilter(
  loadBudgets: (params: LoadBudgetParams) => Promise<void>,
) {
  const status = ref<BudgetStatus>("active")
  const selectedAccountId = ref<number | null>(null)
  const selectedYear = ref<number | null>(null)
  const selectedMonth = ref<number | null>(null)

  const months = [
    { value: 1, label: "มกราคม" },
    { value: 2, label: "กุมภาพันธ์" },
    { value: 3, label: "มีนาคม" },
    { value: 4, label: "เมษายน" },
    { value: 5, label: "พฤษภาคม" },
    { value: 6, label: "มิถุนายน" },
    { value: 7, label: "กรกฎาคม" },
    { value: 8, label: "สิงหาคม" },
    { value: 9, label: "กันยายน" },
    { value: 10, label: "ตุลาคม" },
    { value: 11, label: "พฤศจิกายน" },
    { value: 12, label: "ธันวาคม" },
  ]

  const currentYear = new Date().getFullYear()

  const years = Array.from(
    { length: 5 },
    (_, index) => currentYear - index,
  )

  async function fetchBudgets() {
    if (selectedAccountId.value === null) {
      return
    }

    await loadBudgets({
      status: status.value,
      accountId: selectedAccountId.value,
      year: status.value === "ended"
        ? selectedYear.value
        : null,
      month: status.value === "ended"
        ? selectedMonth.value
        : null,
    })
  }

  async function selectStatus(value: BudgetStatus) {
    if (status.value === value) {
      return
    }

    status.value = value
    selectedYear.value = null
    selectedMonth.value = null

    await fetchBudgets()
  }

  async function handleYearChange() {
    selectedMonth.value = null
    await fetchBudgets()
  }

  async function handleMonthChange() {
    if (selectedYear.value === null) {
      return
    }

    await fetchBudgets()
  }

  watch(selectedAccountId, async (newValue, oldValue) => {
    if (newValue === null || newValue === oldValue) {
      return
    }

    await fetchBudgets()
  })

  return {
    status,
    selectedAccountId,
    selectedYear,
    selectedMonth,

    months,
    years,
    fetchBudgets,
    selectStatus,
    handleYearChange,
    handleMonthChange,
  }
}
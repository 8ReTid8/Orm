import type { BudgetFilter } from "@/types/budget"
import { THAI_MONTHS } from "@/utils/date"
import { ref, watch, type Ref } from "vue"

type BudgetStatus = "active" | "ended"

interface LoadBudgetParams {
  status: BudgetStatus
  accountId: number | null
  year: number | null
  month: number | null
}

export function useBudgetFilter(
  loadBudgets: (params: BudgetFilter) => Promise<void>,
  selectedYear: Ref<number | null>,
  selectedMonth: Ref<number | null>,
) {
  const status = ref<BudgetStatus>("active")
  const selectedAccountId = ref<number | null>(null)
  // const selectedYear = ref<number | null>(null)
  // const selectedMonth = ref<number | null>(null)


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

    fetchBudgets,
    selectStatus,
    handleYearChange,
    handleMonthChange,
  }
}
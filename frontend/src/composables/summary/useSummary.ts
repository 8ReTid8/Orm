// frontend/src/composables/summary/useSummary.ts
import { ref } from "vue"
import { getComparison, getSummary } from "@/services/summary"
import type { ComparisonResponse, Summary } from "@/types/summary"

export function useSummary() {
  const summary = ref<Summary | null>(null)
  const isLoadingSummary = ref(false)
  const comparison = ref<ComparisonResponse | null>(null)
  const isLoadingComparison = ref(false)

  async function loadSummary(
    year: number,
    month: number | null,
    accountId: number | null,
  ) {
    try {
      isLoadingSummary.value = true

      summary.value = await getSummary(
        year,
        month,
        accountId,
      )
    } finally {
      isLoadingSummary.value = false
    }
  }

  async function loadComparison(
    year: number,
    month: number | null,
    accountId: number | null,
    category: string | null,
  ) {

    try {
      isLoadingComparison.value = true

      comparison.value = await getComparison(
        year,
        month,
        accountId,
        category,
      )
    } finally {
      isLoadingComparison.value = false
    }
  }
  return {
    summary,
    comparison,
    isLoadingComparison,
    isLoadingSummary,
    loadSummary,
    loadComparison
  }
}
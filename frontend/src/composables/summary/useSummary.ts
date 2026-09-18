// frontend/src/composables/summary/useSummary.ts
import { ref } from "vue"
import { getSummary } from "@/services/summary"
import type { Summary } from "@/types/summary"

export function useSummary() {
  const summary = ref<Summary | null>(null)
  const isLoadingSummary = ref(false)

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

  return {
    summary,
    isLoadingSummary,
    loadSummary,
  }
}
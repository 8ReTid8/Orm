// frontend/src/composables/summary/useSummary.ts
import { ref } from "vue"
import { getComparison, getSummary } from "@/services/summary"
import type { ComparisonResponse, Summary } from "@/types/summary"

export function useSummary() {
  const summary = ref<Summary | null>(null)
  const isLoadingSummary = ref(false)
  const comparison = ref<ComparisonResponse | null>(null)
  const isLoadingComparison = ref(false)
  const error = ref<string | null>(null)

  async function loadSummary(
    year: number,
    month: number | null,
    accountId: number | null,
  ) {
    try {
      isLoadingSummary.value = true
      error.value = null

      summary.value = await getSummary(
        year,
        month,
        accountId,
      )
    } catch (err) {
      console.error("Failed to load summary:", err)
      error.value = "ไม่สามารถโหลดข้อมูลสรุปได้"
      throw err
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
      error.value = null

      comparison.value = await getComparison(
        year,
        month,
        accountId,
        category,
      )
    } catch (err) {
      console.error("Failed to load comparison:", err)
      error.value = "ไม่สามารถโหลดข้อมูลเปรียบเทียบได้"
      throw err
    } finally {
      isLoadingComparison.value = false
    }
  }

  return {
    summary,
    comparison,
    isLoadingComparison,
    isLoadingSummary,
    error,
    loadSummary,
    loadComparison,
  }
}
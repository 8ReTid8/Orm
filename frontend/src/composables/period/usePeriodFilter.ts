import { ref } from "vue"
import { THAI_MONTHS } from "@/utils/date"
import { getPeriodYears, type PeriodResource } from "@/services/period"

// 👈 รับ resource เข้ามาตอนเรียกใช้ composable
export function usePeriodFilter(resource?: PeriodResource) {
  const selectedYear = ref<number | null>(null)
  const selectedMonth = ref<number | null>(null)
  const years = ref<number[]>([])
  const months = THAI_MONTHS

  async function loadYears() {
    try {
      // 👈 ส่ง resource ไปยัง service เพื่อดึงเฉพาะตารางนั้น
      years.value = await getPeriodYears(resource)
    } catch (error) {
      console.error(`Failed to load ${resource || "all"} period years:`, error)
    }
  }

  function handleYearChange() {
    selectedMonth.value = null
  }

  return {
    selectedYear,
    selectedMonth,
    years,
    months,
    loadYears,
    handleYearChange,
  }
}
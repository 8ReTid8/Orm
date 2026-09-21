import { ref, computed, watch } from "vue"
import type { Ref } from "vue"
import type { Transaction, TransactionFilter } from "@/types/transaction"
import { formatDate } from "@/utils/format"
// พารามิเตอร์สำหรับส่งไปหา Backend API
export interface TransactionServerParams {
  year?: number | null
  month?: number | null
  startDate?: string | null
  endDate?: string | null
  accountId?: number | null
  category?: string | null
}
export function useTransactionFilter(
  transactions: Ref<Transaction[]>,
  selectedYear: Ref<number | null>,
  selectedMonth: Ref<number | null>,
  loadTransactions?: (params: TransactionFilter) => Promise<void>,
) {
  // ==========================================
  // 1. Backend Filter State (ยิงไป Server)
  // ==========================================
  const selectedAccountId = ref<number | null>(null)
  const serverStartDate = ref<string | null>(null)
  const serverEndDate = ref<string | null>(null)
  // const serverAccountId = ref<number | null>(null)
  const serverCategory = ref<string | null>(null)
  // รายการเดือนและปีย้อนหลัง สำหรับ Dropdown
  // ฟังก์ชัน Fetch ข้อมูลจาก Backend (เหมือนใน useBudgetFilter)
  async function fetchTransactions() {
    if (!loadTransactions) return
    await loadTransactions({
      year: selectedYear.value,
      month: selectedMonth.value,
      startDate: serverStartDate.value,
      endDate: serverEndDate.value,
      accountId: selectedAccountId.value,
      category: serverCategory.value,
    })
  }
  // เปลี่ยนเดือน/ปี แล้วยิง fetch ใหม่อัตโนมัติ
  async function handleMonthChange(year: number, month: number) {
    selectedYear.value = year
    selectedMonth.value = month
    await fetchTransactions()
  }

  watch(selectedAccountId, async (newValue, oldValue) => {
    if (newValue === null || newValue === oldValue) {
      return
    }
    await fetchTransactions()
  })

  // ==========================================
  // 2. Frontend Filter State (กรองในเครื่องแบบ Instant)
  // ==========================================
  const selectedDate = ref<Date>(new Date()) // วันที่เลือกบนปฏิทิน
  const clientCategory = ref<string | null>(null)
  const clientType = ref<"income" | "expense" | null>(null)
  // กรองในเครื่องจาก transactions ทั้งก้อนที่โหลดมาแล้ว
  const filteredTransactions = computed(() => {
    return transactions.value.filter((transaction) => {
      // กรองบัญชี (Client)
      
      const categoryMatch =
        clientCategory.value === null ||
        transaction.category === clientCategory.value
      // กรองประเภท รายรับ/รายจ่าย (Client)
      const typeMatch =
        clientType.value === null || transaction.type === clientType.value
      return categoryMatch && typeMatch
    })
  })
  // กรองเฉพาะวันนั้นๆ สำหรับ Daily List
  const selectedDayTransactions = computed(() => {
    const dateStr = formatDate(selectedDate.value)
    return filteredTransactions.value.filter((t) =>
      t.transactionDate.startsWith(dateStr),
    )
  })
  function selectDate(date: Date) {
    selectedDate.value = date
  }
  // ล้างตัวกรอง Frontend
  function resetClientFilters() {
    clientCategory.value = null
    clientType.value = null
  }
  return {
    // Backend filters & actions
    selectedYear,
    selectedMonth,
    serverStartDate,
    serverEndDate,
    selectedAccountId,
    serverCategory,
    fetchTransactions,
    handleMonthChange,
    // Frontend filters
    selectedDate,
    clientCategory,
    clientType,
    filteredTransactions,
    selectedDayTransactions,
    selectDate,
    resetClientFilters,
  }
}
// import { computed, ref } from "vue"
// import type { Transaction } from "@/types/transaction"
// import { formatDate} from "@/utils/format"

// export function useTransactionFilter(
//   transactions: {
//     value: Transaction[]
//   },
// ) {
//   const selectedDate = ref<Date>(new Date())
//   const selectedAccountId = ref<number | null>(null)
//   const selectedCategory = ref<string | null>(null)
//   const filteredTransactions = computed(() => {
//     return transactions.value.filter(transaction => {

//       const accountMatch =
//         selectedAccountId.value === null ||
//         transaction.accountId === selectedAccountId.value

//       const categoryMatch =
//         selectedCategory.value === null ||
//         transaction.category === selectedCategory.value

//       return accountMatch && categoryMatch
//     })
//   })


//   const selectedDayTransactions = computed(() => {
//     const date = formatDate(selectedDate.value)

//     return filteredTransactions.value.filter(
//       transaction =>
//         transaction.transactionDate.startsWith(date),
//     )
//   })

//   function selectDate(date: Date) {
//     selectedDate.value = date
//   }

//   return {
//     selectedDate,
//     selectedAccountId,
//     selectedCategory,
//     filteredTransactions,
//     selectedDayTransactions,
//     selectDate,
//   }
// }

import { ref, computed, watch } from "vue"
import type { Ref } from "vue"
import type { Transaction, TransactionFilter } from "@/types/transaction"
import { formatDate } from "@/utils/format"
import { THAI_MONTHS } from "@/utils/date"
import { usePeriodFilter } from "../period/usePeriodFilter"
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
  // const { selectedYear, selectedMonth } = usePeriodFilter("transaction")
  // const now = new Date()
  // selectedYear.value = now.getFullYear()
  // selectedMonth.value = now.getMonth() + 1
  // const selectedYear = ref<number>(now.getFullYear())
  // const selectedMonth = ref<number>(now.getMonth() + 1)
  const serverStartDate = ref<string | null>(null)
  const serverEndDate = ref<string | null>(null)
  const serverAccountId = ref<number | null>(null)
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
      accountId: serverAccountId.value,
      category: serverCategory.value,
    })
  }
  // เปลี่ยนเดือน/ปี แล้วยิง fetch ใหม่อัตโนมัติ
  async function handleMonthChange(year: number, month: number) {
    selectedYear.value = year
    selectedMonth.value = month
    await fetchTransactions()
  }

  watch(serverAccountId, async (newValue, oldValue) => {
    if (newValue === null || newValue === oldValue) {
      return
    }
    await fetchTransactions()
  })

  // ==========================================
  // 2. Frontend Filter State (กรองในเครื่องแบบ Instant)
  // ==========================================
  const selectedDate = ref<Date>(new Date()) // วันที่เลือกบนปฏิทิน
  const clientAccountId = ref<number | null>(null)
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
    clientAccountId.value = null
    clientCategory.value = null
    clientType.value = null
  }
  return {
    // Backend filters & actions
    selectedYear,
    selectedMonth,
    serverStartDate,
    serverEndDate,
    serverAccountId,
    serverCategory,
    fetchTransactions,
    handleMonthChange,
    // Frontend filters
    selectedDate,
    clientAccountId,
    clientCategory,
    clientType,
    filteredTransactions,
    selectedDayTransactions,
    selectDate,
    resetClientFilters,
  }
}
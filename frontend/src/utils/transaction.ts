// src/utils/transaction.ts
import type { Transaction } from "@/types/transaction"
import { formatThaiDateLong } from "./format"

export interface TransactionDateGroup {
  dateKey: string // "YYYY-MM-DD"
  formattedDate: string // "10 กันยายน 2569"
  totalExpense: number
  totalIncome: number
  transactions: Transaction[]
}

/**
 * จัดกลุ่ม Transactions ตามวันที่ และเรียงลำดับจากวันล่าสุดไปวันแรก
 */
export function groupTransactionsByDate(
  transactions: Transaction[]
): TransactionDateGroup[] {
  if (!transactions || transactions.length === 0) return []

  const groups = new Map<string, Transaction[]>()

  // 1. เรียงลำดับจากวันล่าสุดลงมา
  const sorted = [...transactions].sort(
    (a, b) =>
      new Date(b.transactionDate).getTime() -
      new Date(a.transactionDate).getTime()
  )

  // 2. จัดกลุ่มตาม "YYYY-MM-DD"
  for (const t of sorted) {
    const dateKey = t.transactionDate.slice(0, 10)
    if (!groups.has(dateKey)) {
      groups.set(dateKey, [])
    }
    groups.get(dateKey)!.push(t)
  }

  // 3. แปลงเป็น Array พร้อมคำนวณยอดเงินของแต่ละวัน
  return Array.from(groups.entries()).map(([dateKey, items]) => {
    const totalExpense = items
      .filter((t) => t.type === "expense")
      .reduce((sum, t) => sum + t.amount, 0)

    const totalIncome = items
      .filter((t) => t.type === "income")
      .reduce((sum, t) => sum + t.amount, 0)

    return {
      dateKey,
      formattedDate: formatThaiDateLong(new Date(dateKey)),
      totalExpense,
      totalIncome,
      transactions: items,
    }
  })
}
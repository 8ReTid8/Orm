// src/utils/budget.ts
import type { Budget } from "@/types/budget"

export interface BudgetPeriodGroup {
    startDate: string
    endDate: string
    budgets: Budget[]
}

/**
 * จัดกลุ่ม Budget ที่มีช่วงเวลา (startDate - endDate) เดียวกัน
 */
export function groupBudgetsByPeriod(budgets: Budget[]): BudgetPeriodGroup[] {
    if (!budgets || budgets.length === 0) return []

    const groups = new Map<string, Budget[]>()

    for (const b of budgets) {
        const key = `${b.startDate}_${b.endDate}`
        if (!groups.has(key)) {
            groups.set(key, [])
        }
        groups.get(key)!.push(b)
    }

    return Array.from(groups.values()).flatMap((budgetGroup) => {
        const firstBudget = budgetGroup[0]
        if (!firstBudget) return []

        return {
            startDate: firstBudget.startDate,
            endDate: firstBudget.endDate,
            budgets: budgetGroup
        }
    })
   
}
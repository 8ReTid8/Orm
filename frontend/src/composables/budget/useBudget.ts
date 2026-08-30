import { ref } from "vue";
import type { Budget, BudgetForm } from "@/types/budget";
import { createBudget, getBudgets } from "@/services/budget";


export function useBudget() {
    const budgets = ref<Budget[]>([])
    const editingBudgetId = ref<number | null>(null)
    const isLoadingBudgets = ref(false)
   
    async function loadBudgets(
        year: number,
        month: number,
        status: string
    ) {
        try {
            isLoadingBudgets.value = true

            budgets.value = await getBudgets(
                year,
                month,
                status
            )
        } catch (error) {
            console.error(
                "Failed to load budgets:",
                error,
            )
        } finally {
            isLoadingBudgets.value = false
        }
    }
    
    async function saveBudget(
        form: BudgetForm
    ) {
        if (!form.amount || !form.category || !form.accountId || !form.startDate || !form.endDate) {
            console.log("VALIDATION FAILED")
            return
        }
        if (editingBudgetId.value !== null) {
            // update budget
        } else {
            await createBudget(form)
        }
    }
    return {
        budgets,
        isLoadingBudgets,
        editingBudgetId,

        loadBudgets,
        saveBudget,
    }
}
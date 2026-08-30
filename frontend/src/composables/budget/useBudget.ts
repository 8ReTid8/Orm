import { ref } from "vue";
import type { Budget, BudgetForm } from "@/types/budget";
import {
    createBudget,
    getBudgets,
    updateBudget,
    deleteBudget as deleteBudgetApi
} from "@/services/budget";


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
            await updateBudget(
                editingBudgetId.value,
                form
            )

        } else {
            await createBudget(form)
        }
    }

    async function deleteBudget(id: number) {
        try {
            await deleteBudgetApi(id)
        } catch (error) {
            console.error("Delete budget failed:", error)
            throw error
        }
    }
    function startEdit(budget: Budget) {
        editingBudgetId.value = budget.id
    }

    function cancelEdit() {
        editingBudgetId.value = null
    }

    return {
        budgets,
        isLoadingBudgets,
        editingBudgetId,

        loadBudgets,
        saveBudget,
        deleteBudget,
        startEdit,
        cancelEdit,
    }
}
import { ref } from "vue";
import type { Budget, BudgetFilter, BudgetForm, GetBudgetDetailResponse } from "@/types/budget";
import {
    createBudget,
    getBudgets,
    updateBudget,
    deleteBudget as deleteBudgetApi,
    getBudgetDetail,
} from "@/services/budget";

export function useBudget() {
    const budgets = ref<Budget[]>([])
    const budgetDetail = ref<GetBudgetDetailResponse | null>(null)
    const editingBudgetId = ref<number | null>(null)
    const isLoadingBudgets = ref(false)
    const error = ref<string | null>(null)

    async function loadBudgets({
        year = null,
        month = null,
        status,
        accountId = null,
    }: BudgetFilter) {
        try {
            isLoadingBudgets.value = true
            error.value = null

            budgets.value = await getBudgets(
                year,
                month,
                status,
                accountId
            )
        } catch (err) {
            console.error("Failed to load budgets:", err)
            error.value = "ไม่สามารถโหลดข้อมูลงบประมาณได้"
            throw err
        } finally {
            isLoadingBudgets.value = false
        }
    }

    async function saveBudget(form: BudgetForm) {
        if (!form.amount || !form.category || !form.accountId || !form.startDate || !form.endDate) {
            const validationError = new Error("กรุณากรอกข้อมูลงบประมาณให้ครบถ้วน")
            console.error("Validation failed:", validationError.message)
            throw validationError
        }

        try {
            error.value = null
            if (editingBudgetId.value !== null) {
                const updated = await updateBudget(editingBudgetId.value, form)
                editingBudgetId.value = null
                return updated
            } else {
                return await createBudget(form)
            }
        } catch (err) {
            console.error("Failed to save budget:", err)
            error.value = "ไม่สามารถบันทึกงบประมาณได้"
            throw err
        }
    }

    async function deleteBudget(id: number) {
        try {
            error.value = null
            await deleteBudgetApi(id)
            // budgets.value = budgets.value.filter((b) => b.id !== id)
        } catch (err) {
            console.error("Delete budget failed:", err)
            error.value = "ไม่สามารถลบงบประมาณได้"
            throw err
        }
    }

    async function loadBudgetDetail(id: number) {
        try {
            isLoadingBudgets.value = true
            error.value = null
            budgetDetail.value = await getBudgetDetail(id)
        } catch (err) {
            console.error("Failed to load budget detail:", err)
            error.value = "ไม่สามารถโหลดรายละเอียดงบประมาณได้"
            throw err
        } finally {
            isLoadingBudgets.value = false
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
        budgetDetail,
        error,

        loadBudgetDetail,
        loadBudgets,
        saveBudget,
        deleteBudget,
        startEdit,
        cancelEdit,
    }
}
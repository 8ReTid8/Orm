
import type { Budget, BudgetForm } from "@/types/budget"
import { ref } from "vue"

export function useBudgetForm() {

    const form = ref<BudgetForm>({
        category: "",
        amount: null,
        accountId: null,
        startDate: "",
        endDate: "",
    })

    function resetForm() {
        form.value = {
            category: "",
            amount: null,
            accountId: null,
            startDate: "",
            endDate: "",
        }
    }

    function setEditForm(budget: Budget) {
        form.value = {
            category: budget.category,
            amount: budget.amount,
            accountId: budget.accountId,
            startDate: budget.startDate.slice(0, 10),
            endDate: budget.endDate.slice(0, 10),
        }
    }

    return {
        form,
        resetForm,
        setEditForm,
    }
}
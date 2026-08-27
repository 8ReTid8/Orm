
import type { Budget, BudgetForm } from "@/types/budget"
import { ref } from "vue"

export function useBudgetForm() {
    //   const form = ref<TransactionForm>({
    //     type: "expense",
    //     amount: null,
    //     category: "",
    //     accountId: null,
    //     title: "",
    //     note: "",
    //     transactionDate: "",
    //     slipImage: null,
    //   })
    const form = ref<BudgetForm>({
        category: "",
        amount: null,
        accountId: null,
        startDate: "",
        endDate: "",
    })

    function resetForm(date: string) {
        form.value = {
            category: "",
            amount: null,
            accountId: null,
            startDate: date,
            endDate: date,
        }
    }

    function setEditForm(budget: Budget) {
        form.value = {
            category: budget.category,
            amount: budget.amount,
            accountId: budget.accountId,
            startDate: budget.startDate,
            endDate: budget.endDate,
        }
    }
    //   function resetForm(date: string) {
    //     form.value = {
    //       type: "expense",
    //       amount: null,
    //       category: "",
    //       accountId: null,
    //       title: "",
    //       note: "",
    //       transactionDate: date,
    //       slipImage: null,
    //     }
    //   }

    //   function setEditForm(transaction: Transaction) {
    //     form.value = {
    //       type: transaction.type,
    //       amount: transaction.amount,
    //       category: transaction.category,
    //       accountId: transaction.account.id,
    //       title: transaction.title,
    //       note: transaction.note,
    //       transactionDate:
    //         transaction.transactionDate.slice(0, 10),
    //       slipImage: null,
    //     }
    // }
    return {
        form,
        resetForm,
        setEditForm,
    }
}
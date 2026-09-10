<script setup lang="ts">
import TransactionForm from "@/components/transaction/TransactionForm.vue";
import TransactionList from "@/components/transaction/TransactionList.vue";
import { useBudget } from "@/composables/budget/useBudget";
import { useTransactions } from "@/composables/transaction/useTransaction";
import { useTransactionForm } from "@/composables/transaction/useTransactionForm";
import { useAccountStore } from "@/stores/account";
import { useCategoryStore } from "@/stores/category";
import type { Transaction } from "@/types/transaction";
import { ReceiptText } from "lucide-vue-next";
import { ref, onMounted } from "vue"

const isDialogOpen = ref(false)
const accountStore = useAccountStore()
const categoryStore = useCategoryStore()

const {
    deleteTransaction,
    saveTransaction: saveTransactionApi,
    startEdit,
    cancelEdit,
} = useTransactions()

const {
    form,
    resetForm,
    setEditForm,
} = useTransactionForm()

const {
    isLoadingBudgets,
    budgetDetail,
    loadBudgetDetail,
} = useBudget()

function openEditTransaction(transaction: Transaction) {
    startEdit(transaction)
    setEditForm(transaction)
    isDialogOpen.value = true
}

async function handleDeleteTransaction(id: number) {
    try {
        await deleteTransaction(id)
        // fetchTransactions()
    } catch (error) {
        console.error(
            "Delete transaction failed:",
            error,
        )
    }
}async function saveTransaction() {
  try {
    await saveTransactionApi(form.value)
  
    // await loadTransactions(
    //   selectedDate.value.getFullYear(),
    //   selectedDate.value.getMonth() + 1,
    // )
    cancelEdit()

    isDialogOpen.value = false

    // resetForm(
    //   formatDate(selectedDate.value),
    // )
  } catch (error) {
    console.error(
      "Transaction failed:",
      error,
    )
  }
}
function closeDialog() {
  isDialogOpen.value = false
}
onMounted(async () => {
    await loadBudgetDetail(7) // ตัวอย่างการโหลด Budget Detail ของ Budget ที่มี id = 1
    console.log("Budget Detail:", budgetDetail.value)
})

</script>
<template>
    <div class="space-y-6 p-6">

        <!-- Header -->
        <div>
            <button class="btn btn-ghost btn-sm mb-3">
                ← Back
            </button>

            <div class="flex items-center justify-between">
                <div>
                    <h1 class="text-2xl font-bold">
                        Food Budget
                    </h1>

                    <p class="text-sm text-base-content/60">
                        SCB • 1 Sep 2026 - 30 Sep 2026
                    </p>
                </div>

                <div class="text-right">
                    <p class="text-sm text-base-content/60">
                        Budget
                    </p>

                    <p class="text-2xl font-bold">
                        ฿10,000
                    </p>
                </div>
            </div>
        </div>


        <!-- Budget Usage -->
        <div class="grid grid-cols-1 gap-4 md:grid-cols-3">

            <!-- Spent -->
            <div class="card bg-base-100 shadow-sm">
                <div class="card-body">

                    <p class="text-sm text-base-content/60">
                        Spent
                    </p>

                    <p class="text-2xl font-bold">
                        ฿6,450
                    </p>

                    <p class="text-sm">
                        64.5% of budget
                    </p>

                </div>
            </div>


            <!-- Remaining -->
            <div class="card bg-base-100 shadow-sm">
                <div class="card-body">

                    <p class="text-sm text-base-content/60">
                        Remaining
                    </p>

                    <p class="text-2xl font-bold">
                        ฿3,550
                    </p>

                    <p class="text-sm">
                        35.5% remaining
                    </p>

                </div>
            </div>


            <!-- Average -->
            <div class="card bg-base-100 shadow-sm">
                <div class="card-body">

                    <p class="text-sm text-base-content/60">
                        Average / Day
                    </p>

                    <p class="text-2xl font-bold">
                        ฿215
                    </p>

                    <p class="text-sm text-base-content/60">
                        Based on 30 days
                    </p>

                </div>
            </div>

        </div>


        <!-- Spending Graph -->
        <div class="card bg-base-100 shadow-sm">

            <div class="card-body">

                <div>
                    <h2 class="text-lg font-bold">
                        Spending
                    </h2>

                    <p class="text-sm text-base-content/60">
                        การใช้เงินในช่วง Budget
                    </p>
                </div>


                <!-- Chart -->
                <div class="h-80 w-full">

                    <canvas id="budget-spending-chart"></canvas>

                </div>

            </div>

        </div>


        <!-- Transactions -->
        <div class="card bg-base-100 shadow-sm">

            <div class="card-body">

                <div class="flex items-center justify-between">

                    <div>
                        <h2 class="text-lg font-bold">
                            Transactions
                        </h2>

                        <p class="text-sm text-base-content/60">
                            รายการใช้เงินของ Budget นี้
                        </p>
                    </div>

                    <span class="badge badge-neutral">
                        8 transactions
                    </span>

                </div>


                <!-- Transaction List -->
                <div class="mt-2 divide-y divide-base-200">
                    <!-- <div v-if="loading" class="flex justify-center py-12">
                        <span class="loading loading-spinner loading-lg" />
                    </div> -->

                    <!-- Empty -->
                    <!-- <EmptyState v-if="budgetDetail.transactions.length === 0" :icon="ReceiptText" title="ยังไม่มีรายการในวันนี้"
                        description="กดเพิ่มรายการเพื่อบันทึกรายรับหรือรายจ่าย" /> -->
                    <div class="mt-4 space-y-2">
                        <TransactionList v-if="budgetDetail" :transactions="budgetDetail.transactions"
                            :accounts="accountStore.accounts" :categories="categoryStore.categories"
                            @edit="openEditTransaction" @delete="handleDeleteTransaction" />
                    </div>
                </div>
            </div>
        </div>
    </div>
    <TransactionForm :open="isDialogOpen" :form="form" :accounts="accountStore.accounts"
        :categories="categoryStore.categories" :selected-date="selectedDate" @close="closeDialog"
        @save="saveTransaction" />
</template>
<script setup lang="ts">
import BudgetSpendingChart from "@/components/budget/BudgetSpendingChart.vue";
import BudgetSummary from "@/components/budget/BudgetSummary.vue";
import Pagination from "@/components/common/Pagination.vue";
import TotalList from "@/components/common/TotalList.vue";
import GroupedTransactionList from "@/components/transaction/GroupedTransactionList.vue";
import TransactionForm from "@/components/transaction/TransactionForm.vue";
import { useBudget } from "@/composables/budget/useBudget";
import { useTransactions } from "@/composables/transaction/useTransaction";
import { useTransactionForm } from "@/composables/transaction/useTransactionForm";
import { useAccountStore } from "@/stores/account";
import { useCategoryStore } from "@/stores/category";
import type { Transaction } from "@/types";
import { resolveCategoryIcon } from "@/utils/categoryIcons";
import { formatDate, formatMoney, formatThaiDateLong } from "@/utils/format";
import { groupTransactionsByDate } from "@/utils/transaction";
import { AlertCircle, ArrowLeft, Calendar, ReceiptText, Wallet } from "lucide-vue-next";
import { ref, onMounted, computed, watch } from "vue"
import { useRoute, useRouter } from "vue-router";

const route = useRoute()
const router = useRouter()

const isDialogOpen = ref(false)
const accountStore = useAccountStore()
const categoryStore = useCategoryStore()

const budgetId = computed(() => {
    const id = route.params.id || route.query.id
    return id ? Number(id) : null
})

const {
    deleteTransaction,
    saveTransaction: saveTransactionApi,
    startEdit,
    cancelEdit,
} = useTransactions()

const {
    form,
    setEditForm,
} = useTransactionForm()

const {
    budgetDetail,
    loadBudgetDetail,
} = useBudget()

const isLoading = ref(false)
const budget = computed(() => budgetDetail.value?.budget)
const transactions = computed(() => budgetDetail.value?.transactions ?? [])
const currentPage = ref(1)
const pageSize = ref(10) // กำหนดจำนวนรายการต่อหน้า (เช่น 10 รายการ)
const totalPages = computed(() =>
    Math.ceil(transactions.value.length / pageSize.value) || 1
)

// รายการ transaction เฉพาะหน้าปัจจุบัน
const paginatedTransactions = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value
    return transactions.value.slice(start, start + pageSize.value)
    // return sortedTransactions.value.slice(start, start + pageSize.value)
})
// ป้องกันกรณีลบรายการแล้วหน้าปัจจุบันเกิน totalPages
watch(totalPages, (newTotal) => {
    if (currentPage.value > newTotal) {
        currentPage.value = Math.max(1, newTotal)
    }
})
function getCategoryIcon(categoryName?: string) {
    if (!categoryName) return null
    const cat = categoryStore.categories.find((c) => c.name === categoryName)
    return resolveCategoryIcon(cat?.icon)
}
function getAccountName(accountId?: number) {
    if (!accountId) return "ทุกบัญชี"
    const acc = accountStore.accounts.find((a) => a.id === accountId)
    return acc?.name ?? "ไม่พบบัญชี"
}
function getDaysRemaining(endDateStr?: string) {
    if (!endDateStr) return ""
    const end = new Date(endDateStr)
    const today = new Date()
    today.setHours(0, 0, 0, 0)
    end.setHours(0, 0, 0, 0)
    const diffDays = Math.ceil((end.getTime() - today.getTime()) / (1000 * 60 * 60 * 24))
    return diffDays > 0 ? `เหลือ ${diffDays} วัน` : "สิ้นสุดแล้ว"
}

function openEditTransaction(transaction: Transaction) {
    startEdit(transaction)
    setEditForm(transaction)
    isDialogOpen.value = true
}


async function handleDeleteTransaction(id: number) {
    try {
        await deleteTransaction(id)
        if (budgetId.value) {
            await loadBudgetDetail(budgetId.value)
        }
    } catch (error) {
        console.error(
            "Delete transaction failed:",
            error,
        )
    }
}

async function saveTransaction() {
    try {
        await saveTransactionApi(form.value)

        if (budgetId.value) {
            await loadBudgetDetail(budgetId.value)
        }
        cancelEdit()

        isDialogOpen.value = false

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
    isLoading.value = true
    try {
        await Promise.all([
            accountStore.loadAccounts(),
            categoryStore.loadCategories(),
        ])
        if (budgetId.value) {
            await loadBudgetDetail(budgetId.value)
        }
    } finally {
        isLoading.value = false
    }
})
</script>
<template>
    <section class="space-y-6!">
        <!-- Loading State -->
        <div v-if="isLoading" class="flex justify-center py-16">
            <span class="loading loading-spinner loading-lg text-primary" />
        </div>
        <!-- Not Found State -->
        <div v-else-if="!budget" class="text-center py-16">
            <p class="text-lg font-semibold">ไม่พบข้อมูลงบประมาณนี้</p>
            <button class="btn btn-outline btn-sm mt-4" @click="router.back()">
                ย้อนกลับ
            </button>
        </div>
        <!-- Content -->
        <template v-else>
            <!-- Header -->
            <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
                <div class="flex items-center gap-3 min-w-0">
                    <!-- Back Button -->
                    <button type="button" class="btn btn-circle btn-ghost btn-sm shrink-0" title="ย้อนกลับ"
                        @click="router.back()">
                        <ArrowLeft class="size-5" />
                    </button>
                    <!-- Category Icon -->
                    <div class="flex size-12 shrink-0 items-center justify-center rounded-2xl bg-error/15 text-error">
                        <component :is="getCategoryIcon(budget.category)" class="size-6" />
                    </div>
                    <div class="min-w-0">
                        <div class="flex flex-wrap items-center gap-2">
                            <h1 class="text-2xl font-bold truncate">
                                งบ{{ budget.category }}
                            </h1>
                            <span class="badge badge-sm badge-ghost font-medium text-base">
                                {{ getDaysRemaining(budget.endDate) }}
                            </span>
                        </div>
                        <!-- Account & Date Details -->
                        <div class="mt-1 flex flex-wrap items-center gap-x-2 text-base text-base-content/60">
                            <span class="flex items-center gap-1">
                                <Wallet class="size-3.5" />
                                {{ getAccountName(budget.accountId) }}
                            </span>
                            <span>•</span>
                            <span class="flex items-center gap-1">
                                <Calendar class="size-3.5" />
                                {{ formatDate(new Date(budget.startDate)) }} ถึง {{ formatDate(new Date(budget.endDate))
                                }}
                            </span>
                        </div>
                    </div>
                </div>
                <!-- Total Budget Badge / Text -->
                <div class="rounded-2xl border border-base-200 bg-base-100 p-3 sm:text-right shadow-xs">
                    <p class="text-xs text-base-content/50">งบประมาณที่ตั้งไว้</p>
                    <p class="text-2xl font-bold">฿{{ formatMoney(budget.amount) }}</p>
                </div>
            </div>
            <!-- Overview Cards (3 Cards Grid) -->
            <BudgetSummary :budget="budget" />

            <BudgetSpendingChart v-if="budget" :budget="budget" :transactions="transactions" />
            <!-- Transactions Section -->
            <div class="card border border-base-300 bg-base-100 shadow-sm">
                <div class="card-body p-5 sm:p-6">
                    <div class="flex items-center justify-between">
                        <div>
                            <h2 class="text-xl font-bold">รายการใช้เงิน</h2>
                            <p class="text-sm text-base-content/60">
                                ประวัติรายการที่อยู่ในงบประมาณช่วงเวลานี้
                            </p>
                        </div>
                        <TotalList :total="transactions.length" />
                    </div>
                    <!-- 👈 ใช้งาน GroupedTransactionList แทนลูปเดิมทั้งหมด -->
                    <GroupedTransactionList :transactions="paginatedTransactions" :accounts="accountStore.accounts"
                        :categories="categoryStore.categories" empty-title="ยังไม่มีรายการใช้จ่าย"
                        empty-description="ยังไม่มีการบันทึกรายการในหมวดหมู่นี้ในช่วงเวลาที่กำหนด"
                        @edit="openEditTransaction" @delete="handleDeleteTransaction" />
                    <Pagination v-model:current-page="currentPage" :total-pages="totalPages" :disabled="isLoading" />
                </div>
            </div>
        </template>
    </section>
    <TransactionForm :open="isDialogOpen" :form="form" :accounts="accountStore.accounts"
        :categories="categoryStore.categories" @close="closeDialog" @save="saveTransaction" />
</template>
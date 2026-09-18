<script setup lang="ts">
import { ref, computed, onMounted } from "vue"
import {
    ArrowDownLeft,
    ArrowUpRight,
    Wallet,
    Percent,
    Award,
} from "lucide-vue-next"
import { useAccountStore } from "@/stores/account"
import { usePeriodFilter } from "@/composables/period/usePeriodFilter"
import { formatMoney } from "@/utils/format"
import PeriodFilter from "@/components/filter/PeriodFilter.vue"
import AccountFilter from "@/components/filter/accountFilter.vue"
import CategoryPieChart from "@/components/category/CategoryPieChart.vue"
import MonthlyComparisonChart from "@/components/category/MonthlyComparisonChart.vue"
import { useSummary } from "@/composables/summary/useSummary"

const accountStore = useAccountStore()
const selectedAccountId = ref<number | null>(null)
const viewMode = ref<"monthly" | "yearly">("monthly")

// 👈 ใช้ usePeriodFilter ตัวกลาง
const {
    years,
    months,
    selectedYear,
    selectedMonth,
    loadYears,
} = usePeriodFilter("transaction")

const {
    summary,
    isLoadingSummary,
    loadSummary,
} = useSummary()

// ฟังก์ชันดึงข้อมูลตาม Filter
async function fetchSummary() {
    if (!selectedYear.value) {
        return
    }

    await loadSummary(
        selectedYear.value,
        viewMode.value === "monthly"
            ? selectedMonth.value
            : null,
        selectedAccountId.value,
    )
}

async function updateYear(year: number | null) {
    selectedYear.value = year
    await fetchSummary()
}

async function updateMonth(month: number | null) {
    selectedMonth.value = month
    await fetchSummary()
}

async function updateAccount(accountId: number | null) {
    selectedAccountId.value = accountId
    await fetchSummary()
}

// เมื่อสลับโหมด รายเดือน / รายปี
async function setViewMode(mode: "monthly" | "yearly") {
    viewMode.value = mode
    if (mode === "monthly" && !selectedMonth.value) {
        selectedMonth.value = new Date().getMonth() + 1
    }
    await fetchSummary()
}

// 5 อันดับแรกที่มีรายจ่ายสูงสุด
const topExpenseCategories = computed(() => summary.value?.expenseByCategory.slice(0, 5) ?? [],)

onMounted(async () => {
    const now = new Date()
    selectedYear.value = now.getFullYear()
    selectedMonth.value = now.getMonth() + 1

    await Promise.all([
        accountStore.loadAccounts(),
        loadYears(),
    ])

    const firstAccount = accountStore.accounts[0]
    if (firstAccount) {
        selectedAccountId.value = firstAccount.id
    }

    await fetchSummary()
})
</script>

<template>
    <section class="space-y-6!">
        <!-- 1. Header & Filters -->
        <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
            <div>
                <h1 class="text-2xl font-bold">Summary</h1>
                <p class="text-base-content/60">ภาพรวมสถิติรายรับ-รายจ่าย และพฤติกรรมการใช้เงิน</p>
            </div>

            <!-- Controls & Filters -->
            <div class="flex flex-wrap items-center gap-2">

                <!-- ตัวเลือก ปี / เดือน -->
                <PeriodFilter :years="years" :months="viewMode === 'monthly' ? months : []" :model-year="selectedYear"
                    :model-month="viewMode === 'monthly' ? selectedMonth : null"
                    @update:model-year="updateYear"
                    @update:model-month="updateMonth" />

                <AccountFilter :model-value="selectedAccountId" :accounts="accountStore.accounts"
                    @update:model-value="updateAccount" />
            </div>
        </div>

        <!-- Loading State -->
        <div v-if="isLoadingSummary" class="flex justify-center py-12">
            <span class="loading loading-spinner loading-lg text-primary" />
        </div>

        <!-- <template v-else> -->
        <template v-else-if="summary">
            <!-- 2. Overview Stat Cards (4 ใบ) -->
            <div class="grid grid-cols-1 gap-3 sm:grid-cols-2 lg:grid-cols-4">
                <!-- รายรับรวม -->
                <div class="rounded-2xl bg-success/10 p-5 border border-success/20">
                    <div class="flex items-center justify-between text-success">
                        <span class="text-sm font-medium">รายรับรวม</span>
                        <ArrowDownLeft class="size-5" />
                    </div>
                    <p class="mt-3 text-2xl font-bold text-success">
                        +฿{{ formatMoney(summary.totalIncome) }}
                    </p>
                </div>

                <!-- รายจ่ายรวม -->
                <div class="rounded-2xl bg-error/10 p-5 border border-error/20">
                    <div class="flex items-center justify-between text-error">
                        <span class="text-sm font-medium">รายจ่ายรวม</span>
                        <ArrowUpRight class="size-5" />
                    </div>
                    <p class="mt-3 text-2xl font-bold text-error">
                        -฿{{ formatMoney(summary.totalExpense) }}
                    </p>
                </div>

                <!-- สุทธิคงเหลือ -->
                <div class="rounded-2xl bg-base-200 p-5 border border-base-300">
                    <div class="flex items-center justify-between text-base-content/70">
                        <span class="text-sm font-medium">สุทธิคงเหลือ</span>
                        <Wallet class="size-5" />
                    </div>
                    <p class="mt-3 text-2xl font-bold" :class="summary.netBalance >= 0 ? 'text-success' : 'text-error'">
                        {{ summary.netBalance >= 0 ? '+' : '' }}฿{{ formatMoney(summary.netBalance) }}
                    </p>
                </div>

                <!-- อัตราการออม -->
                <div class="rounded-2xl bg-base-200 p-5 border border-base-300">
                    <div class="flex items-center justify-between text-base-content/70">
                        <span class="text-sm font-medium">อัตราการออม</span>
                        <Percent class="size-5" />
                    </div>
                    <div class="mt-3 flex items-baseline gap-2">
                        <p class="text-2xl font-bold text-primary">
                            {{ summary.savingsRate.toFixed(1) }}%
                        </p>
                        <span class="text-xs text-base-content/60">ของรายรับ</span>
                    </div>
                </div>
            </div>

            <!-- 3. กราฟ 2 ฝั่ง (โดนัท + แท่งเปรียบเทียบ) -->
            <div class="grid grid-cols-1 gap-6 lg:grid-cols-2">
                <!-- สัดส่วนรายรับ -->
                <div class="card bg-base-100 p-5 border border-base-200 shadow-sm rounded-2xl">
                    <div class="flex items-center justify-between mb-4">
                        <h2 class="font-bold text-lg text-success">สัดส่วนรายรับตามหมวดหมู่</h2>
                        <span class="text-sm font-semibold text-success">+฿{{ formatMoney(summary.totalIncome) }}</span>
                    </div>
                    <CategoryPieChart :items="summary.incomeByCategory" empty-text="ไม่มีข้อมูลรายรับในช่วงเวลานี้" />
                </div>
                <!-- สัดส่วนรายจ่าย -->
                <div class="card bg-base-100 p-5 border border-base-200 shadow-sm rounded-2xl">
                    <div class="flex items-center justify-between mb-4">
                        <h2 class="font-bold text-lg text-error">สัดส่วนรายจ่ายตามหมวดหมู่</h2>
                        <span class="text-sm font-semibold text-error">-฿{{ formatMoney(summary.totalExpense) }}</span>
                    </div>
                    <CategoryPieChart :items="summary.expenseByCategory" empty-text="ไม่มีข้อมูลรายจ่ายในช่วงเวลานี้" />
                </div>
            </div>
            <!-- กราฟเปรียบเทียบรายรับและรายจ่ายตามช่วงเวลาที่เลือก -->
            <div class="card bg-base-100 p-5 border border-base-200 shadow-sm rounded-2xl">
                <h2 class="font-bold text-lg mb-4">
                    เปรียบเทียบรายรับ vs รายจ่าย{{ summary.comparisonPeriod === "day" ? "รายวัน" : "รายเดือน" }}
                    ({{ summary.comparisonPeriod === "day" ? "เดือน " + selectedMonth : "ปี " + (selectedYear ? selectedYear + 543 : "") }})
                </h2>
                <MonthlyComparisonChart :items="summary.comparison" :period="summary.comparisonPeriod" />
            </div>

            <!-- 4. 5 อันดับหมวดหมู่ที่มีรายจ่ายสูงสุด -->
            <div class="card bg-base-100 p-5 border border-base-200 shadow-sm rounded-2xl">
                <div class="flex items-center gap-2 mb-4">
                    <Award class="size-5 text-warning" />
                    <h2 class="font-bold text-lg">หมวดหมู่ที่มีรายจ่ายสูงสุด</h2>
                </div>

                <div v-if="topExpenseCategories.length === 0" class="py-8 text-center text-base-content/50">
                    ไม่มีรายการค่าใช้จ่ายในช่วงเวลานี้
                </div>

                <div v-else class="space-y-4">
                    <div v-for="(item, index) in topExpenseCategories" :key="item.name" class="space-y-1.5">
                        <div class="flex items-center justify-between text-sm">
                            <div class="flex items-center gap-2">
                                <span class="font-semibold text-base-content/50 w-5">#{{ index + 1 }}</span>
                                <span class="font-medium">{{ item.name }}</span>
                            </div>
                            <div class="flex items-center gap-3">
                                <span class="font-bold text-error">-฿{{ formatMoney(item.amount) }}</span>
                                <span class="text-xs text-base-content/60 w-12 text-right">
                                    {{ summary.totalExpense > 0 ? ((item.amount / summary.totalExpense) *
                                    100).toFixed(1) : 0 }}%
                                </span>
                            </div>
                        </div>

                        <!-- Progress bar สัดส่วน -->
                        <progress class="progress progress-error w-full h-2" :value="item.amount"
                            :max="summary.totalExpense" />
                    </div>
                </div>
            </div>
        </template>
    </section>
</template>

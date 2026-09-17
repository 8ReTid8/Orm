<script setup lang="ts">
import { ref, computed, onMounted } from "vue"
import {
  ArrowDownLeft,
  ArrowUpRight,
  Wallet,
  Percent,
  TrendingUp,
  Award,
} from "lucide-vue-next"
import { useAccountStore } from "@/stores/account"
import { useTransactions } from "@/composables/transaction/useTransaction"
import { usePeriodFilter } from "@/composables/period/usePeriodFilter"
import { formatMoney } from "@/utils/format"
import { resolveCategoryIcon } from "@/utils/categoryIcons"
import PeriodFilter from "@/components/filter/PeriodFilter.vue"
import AccountFilter from "@/components/filter/accountFilter.vue"
// import MonthlyComparisonChart from "@/components/summary/MonthlyComparisonChart.vue"
import CategoryPieChart from "@/components/category/CategoryPieChart.vue"
import MonthlyComparisonChart from "@/components/category/MonthlyComparisonChart.vue"

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
  handleYearChange,
} = usePeriodFilter("transaction")

const { transactions, isLoadingTransactions, loadTransactions } = useTransactions()

// ฟังก์ชันดึงข้อมูลตาม Filter
async function fetchSummary() {
  await loadTransactions({
    year: selectedYear.value,
    month: viewMode.value === "monthly" ? selectedMonth.value : null,
    accountId: selectedAccountId.value,
  })
}

// เมื่อสลับโหมด รายเดือน / รายปี
async function setViewMode(mode: "monthly" | "yearly") {
  viewMode.value = mode
  if (mode === "monthly" && !selectedMonth.value) {
    selectedMonth.value = new Date().getMonth() + 1
  }
  await fetchSummary()
}

// คำนวณยอดสรุปต่างๆ
const totalIncome = computed(() =>
  transactions.value.filter((t) => t.type === "income").reduce((sum, t) => sum + t.amount, 0)
)

const totalExpense = computed(() =>
  transactions.value.filter((t) => t.type === "expense").reduce((sum, t) => sum + t.amount, 0)
)

const netBalance = computed(() => totalIncome.value - totalExpense.value)

const savingsRate = computed(() => {
  if (totalIncome.value <= 0) return 0
  return Math.max(0, (netBalance.value / totalIncome.value) * 100)
})

// รวมยอดรายจ่ายแยกตามหมวดหมู่ (เรียงจากมากไปน้อย)
const expenseByCategory = computed(() => {
  const map = new Map<string, number>()
  for (const t of transactions.value) {
    if (t.type === "expense") {
      map.set(t.category, (map.get(t.category) || 0) + t.amount)
    }
  }
  return Array.from(map.entries())
    .map(([name, amount]) => ({ name, amount }))
    .sort((a, b) => b.amount - a.amount)
})

// 5 อันดับแรกที่มีรายจ่ายสูงสุด
const topExpenseCategories = computed(() => expenseByCategory.value.slice(0, 5))

onMounted(async () => {
  const now = new Date()
  selectedYear.value = now.getFullYear()
  selectedMonth.value = now.getMonth() + 1

  await Promise.all([
    accountStore.loadAccounts(),
    loadYears(),
  ])

  if (accountStore.accounts.length > 0) {
    selectedAccountId.value = accountStore.accounts[0].id
  }

  await fetchSummary()
})
</script>

<template>
  <section class="space-y-6">
    <!-- 1. Header & Filters -->
    <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
      <div>
        <h1 class="text-2xl font-bold">สรุปผลการใช้เงิน</h1>
        <p class="text-base-content/60">ภาพรวมสถิติรายรับ-รายจ่าย และพฤติกรรมการใช้เงิน</p>
      </div>

      <!-- Controls & Filters -->
      <div class="flex flex-wrap items-center gap-2">
        <!-- สลับโหมด รายเดือน / รายปี -->
        <div class="join bg-base-200 p-1 rounded-xl">
          <button
            class="btn btn-sm join-item border-none"
            :class="viewMode === 'monthly' ? 'btn-neutral text-white shadow-sm' : 'btn-ghost'"
            @click="setViewMode('monthly')"
          >
            รายเดือน
          </button>
          <button
            class="btn btn-sm join-item border-none"
            :class="viewMode === 'yearly' ? 'btn-neutral text-white shadow-sm' : 'btn-ghost'"
            @click="setViewMode('yearly')"
          >
            รายปี
          </button>
        </div>

        <!-- ตัวเลือก ปี / เดือน -->
        <PeriodFilter
          :years="years"
          :months="viewMode === 'monthly' ? months : []"
          :model-year="selectedYear"
          :model-month="viewMode === 'monthly' ? selectedMonth : null"
          @update:model-year="selectedYear = $event; fetchSummary()"
          @update:model-month="selectedMonth = $event; fetchSummary()"
          @year-change="handleYearChange(); fetchSummary()"
          @month-change="fetchSummary()"
        />

        <AccountFilter
          v-model="selectedAccountId"
          :accounts="accountStore.accounts"
          @update:model-value="fetchSummary()"
        />
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="isLoadingTransactions" class="flex justify-center py-12">
      <span class="loading loading-spinner loading-lg text-primary" />
    </div>

    <template v-else>
      <!-- 2. Overview Stat Cards (4 ใบ) -->
      <div class="grid grid-cols-1 gap-3 sm:grid-cols-2 lg:grid-cols-4">
        <!-- รายรับรวม -->
        <div class="rounded-2xl bg-success/10 p-5 border border-success/20">
          <div class="flex items-center justify-between text-success">
            <span class="text-sm font-medium">รายรับรวม</span>
            <ArrowDownLeft class="size-5" />
          </div>
          <p class="mt-3 text-2xl font-bold text-success">
            +฿{{ formatMoney(totalIncome) }}
          </p>
        </div>

        <!-- รายจ่ายรวม -->
        <div class="rounded-2xl bg-error/10 p-5 border border-error/20">
          <div class="flex items-center justify-between text-error">
            <span class="text-sm font-medium">รายจ่ายรวม</span>
            <ArrowUpRight class="size-5" />
          </div>
          <p class="mt-3 text-2xl font-bold text-error">
            -฿{{ formatMoney(totalExpense) }}
          </p>
        </div>

        <!-- สุทธิคงเหลือ -->
        <div class="rounded-2xl bg-base-200 p-5 border border-base-300">
          <div class="flex items-center justify-between text-base-content/70">
            <span class="text-sm font-medium">สุทธิคงเหลือ</span>
            <Wallet class="size-5" />
          </div>
          <p
            class="mt-3 text-2xl font-bold"
            :class="netBalance >= 0 ? 'text-success' : 'text-error'"
          >
            {{ netBalance >= 0 ? '+' : '' }}฿{{ formatMoney(netBalance) }}
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
              {{ savingsRate.toFixed(1) }}%
            </p>
            <span class="text-xs text-base-content/60">ของรายรับ</span>
          </div>
        </div>
      </div>

      <!-- 3. กราฟ 2 ฝั่ง (โดนัท + แท่งเปรียบเทียบ) -->
      <div class="grid grid-cols-1 gap-6 lg:grid-cols-2">
        <!-- กราฟสัดส่วนรายจ่ายตามหมวดหมู่ -->
        <div class="card bg-base-100 p-5 border border-base-200 shadow-sm rounded-2xl">
          <h2 class="font-bold text-lg mb-4">สัดส่วนรายจ่ายตามหมวดหมู่</h2>
          <CategoryPieChart :items="expenseByCategory" />
        </div>

        <!-- กราฟเปรียบเทียบ รายรับ vs รายจ่าย ตลอดปี -->
        <div class="card bg-base-100 p-5 border border-base-200 shadow-sm rounded-2xl">
          <h2 class="font-bold text-lg mb-4">เปรียบเทียบ รายรับ vs รายจ่าย (ปี {{ selectedYear ? selectedYear + 543 : '' }})</h2>
          <MonthlyComparisonChart :transactions="transactions" />
        </div>
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
          <div
            v-for="(item, index) in topExpenseCategories"
            :key="item.name"
            class="space-y-1.5"
          >
            <div class="flex items-center justify-between text-sm">
              <div class="flex items-center gap-2">
                <span class="font-semibold text-base-content/50 w-5">#{{ index + 1 }}</span>
                <span class="font-medium">{{ item.name }}</span>
              </div>
              <div class="flex items-center gap-3">
                <span class="font-bold text-error">-฿{{ formatMoney(item.amount) }}</span>
                <span class="text-xs text-base-content/60 w-12 text-right">
                  {{ totalExpense > 0 ? ((item.amount / totalExpense) * 100).toFixed(1) : 0 }}%
                </span>
              </div>
            </div>

            <!-- Progress bar สัดส่วน -->
            <progress
              class="progress progress-error w-full h-2"
              :value="item.amount"
              :max="totalExpense"
            />
          </div>
        </div>
      </div>
    </template>
  </section>
</template>
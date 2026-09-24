<script setup lang="ts">
import { ref, computed, onMounted } from "vue"
import { useRoute, useRouter } from "vue-router"
import {
  ArrowLeft,
  Wallet,
  ArrowDownLeft,
  ArrowUpRight,
  ReceiptText,
} from "lucide-vue-next"
import { useAccountStore } from "@/stores/account"
import { useCategoryStore } from "@/stores/category"
import { useTransactions } from "@/composables/transaction/useTransaction"
import { useTransactionForm } from "@/composables/transaction/useTransactionForm"
import { usePeriodFilter } from "@/composables/period/usePeriodFilter"
import { formatMoney } from "@/utils/format"

import PeriodFilter from "@/components/filter/PeriodFilter.vue"
import GroupedTransactionList from "@/components/transaction/GroupedTransactionList.vue"
import TransactionForm from "@/components/transaction/TransactionForm.vue"
import TotalList from "@/components/common/TotalList.vue"
import type { Transaction } from "@/types/transaction"

const route = useRoute()
const router = useRouter()
const accountStore = useAccountStore()
const categoryStore = useCategoryStore()

const accountId = computed(() => Number(route.params.id))
const account = computed(() =>
  accountStore.accounts.find((a) => a.id === accountId.value)
)

const viewMode = ref<"monthly" | "yearly">("monthly")
const isDialogOpen = ref(false)

const {
  years,
  months,
  selectedYear,
  selectedMonth,
  loadYears,
} = usePeriodFilter("transaction")

const {
  transactions,
  isLoadingTransactions,
  loadTransactions,
  deleteTransaction,
  saveTransaction: saveTransactionApi,
  startEdit,
  cancelEdit,
} = useTransactions()

const { form, setEditForm, resetForm } = useTransactionForm()

// ดึงรายการเดินบัญชีของบัญชีนี้
async function fetchTransactions() {
  if (!accountId.value || !selectedYear.value) return

  await loadTransactions({
    accountId: accountId.value,
    year: selectedYear.value,
    month: viewMode.value === "monthly" ? selectedMonth.value : null,
  })
}

async function updateYear(year: number | null) {
  selectedYear.value = year
  await fetchTransactions()
}

async function updateMonth(month: number | null) {
  selectedMonth.value = month
  await fetchTransactions()
}

async function setViewMode(mode: "monthly" | "yearly") {
  viewMode.value = mode
  if (mode === "monthly" && !selectedMonth.value) {
    selectedMonth.value = new Date().getMonth() + 1
  }
  await fetchTransactions()
}

// คำนวณกระแสเงินสดเฉพาะช่วงเวลานี้
const totalIncome = computed(() =>
  transactions.value
    .filter((t) => t.type === "income")
    .reduce((sum, t) => sum + t.amount, 0)
)

const totalExpense = computed(() =>
  transactions.value
    .filter((t) => t.type === "expense")
    .reduce((sum, t) => sum + t.amount, 0)
)

const netFlow = computed(() => totalIncome.value - totalExpense.value)

// แก้ไขและลบรายการ
function openEditTransaction(transaction: Transaction) {
  startEdit(transaction)
  setEditForm(transaction)
  isDialogOpen.value = true
}

function closeDialog() {
  isDialogOpen.value = false
  cancelEdit()
//   resetForm()
}

async function saveTransaction() {
  try {
    await saveTransactionApi(form.value)
    closeDialog()
    await Promise.all([
      fetchTransactions(),
      accountStore.loadAccounts(true), // รีเฟรชยอดคงเหลือรวมของบัญชี
    ])
  } catch (error) {
    console.error("Save transaction failed:", error)
  }
}

async function handleDeleteTransaction(id: number) {
  try {
    await deleteTransaction(id)
    await Promise.all([
      fetchTransactions(),
      accountStore.loadAccounts(true), // รีเฟรชยอดคงเหลือรวมของบัญชี
    ])
  } catch (error) {
    console.error("Delete transaction failed:", error)
  }
}

onMounted(async () => {
  const now = new Date()
  selectedYear.value = now.getFullYear()
  selectedMonth.value = now.getMonth() + 1

  await Promise.all([
    accountStore.loadAccounts(),
    categoryStore.loadCategories(),
    loadYears(),
  ])

  await fetchTransactions()
})
</script>

<template>
  <section class="space-y-6!">
    <!-- Not Found State -->
    <div v-if="!account && !accountStore.isLoading" class="text-center py-16">
      <p class="text-lg font-semibold">ไม่พบข้อมูลบัญชีนี้</p>
      <button class="btn btn-outline btn-sm mt-4" @click="router.back()">
        ย้อนกลับ
      </button>
    </div>

    <template v-else-if="account">
      <!-- 1. Header บัญชี & ยอดคงเหลือปัจจุบัน -->
      <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
        <div class="flex items-center gap-3 min-w-0">
          <!-- ปุ่มย้อนกลับ -->
          <button
            type="button"
            class="btn btn-circle btn-ghost btn-sm shrink-0"
            title="ย้อนกลับ"
            @click="router.back()"
          >
            <ArrowLeft class="size-5" />
          </button>

          <!-- ไอคอนบัญชี -->
          <div class="flex size-12 shrink-0 items-center justify-center rounded-2xl bg-lime-300 text-green-700">
            <Wallet class="size-6" />
          </div>

          <div class="min-w-0">
            <h1 class="text-2xl font-bold truncate">
              {{ account.name }}
            </h1>
            <p class="text-xs text-base-content/60">
              สมุดบัญชีและประวัติการเดินบัญชี
            </p>
          </div>
        </div>

        <!-- การ์ดยอดคงเหลือปัจจุบันของบัญชี -->
        <div class="rounded-2xl border border-base-200 bg-base-100 p-4 sm:text-right shadow-xs">
          <p class="text-xs text-base-content/50">ยอดคงเหลือปัจจุบัน</p>
          <p class="text-2xl font-bold text-success">
            ฿{{ formatMoney(account.balance) }}
          </p>
        </div>
      </div>

      <!-- 2. Cashflow Cards (3 ใบ สรุปเงินเข้า-ออกรอบนี้) -->
      <div class="grid grid-cols-1 gap-3 sm:grid-cols-3">
        <!-- เงินเข้า -->
        <div class="rounded-2xl bg-success/10 p-5 border border-success/20">
          <div class="flex items-center justify-between text-success">
            <span class="text-sm font-medium">เงินเข้า</span>
            <ArrowDownLeft class="size-5" />
          </div>
          <p class="mt-3 text-2xl font-bold text-success">
            +฿{{ formatMoney(totalIncome) }}
          </p>
        </div>

        <!-- เงินออก -->
        <div class="rounded-2xl bg-error/10 p-5 border border-error/20">
          <div class="flex items-center justify-between text-error">
            <span class="text-sm font-medium">เงินออก</span>
            <ArrowUpRight class="size-5" />
          </div>
          <p class="mt-3 text-2xl font-bold text-error">
            -฿{{ formatMoney(totalExpense) }}
          </p>
        </div>

        <!-- สุทธิรอบนี้ -->
        <div class="rounded-2xl bg-base-200 p-5 border border-base-300">
          <div class="flex items-center justify-between text-base-content/70">
            <span class="text-sm font-medium">ส่วนต่างเงินในรอบนี้</span>
            <Wallet class="size-5" />
          </div>
          <p
            class="mt-3 text-2xl font-bold"
            :class="netFlow >= 0 ? 'text-success' : 'text-error'"
          >
            {{ netFlow >= 0 ? '+' : '' }}฿{{ formatMoney(netFlow) }}
          </p>
        </div>
      </div>

      <!-- 3. แถบควบคุมตัวกรอง (สลับ รายเดือน / รายปี) -->
      <div class="flex flex-wrap items-center justify-between gap-3">
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

        <!-- PeriodFilter -->
        <PeriodFilter
          :years="years"
          :months="viewMode === 'monthly' ? months : []"
          :model-year="selectedYear"
          :model-month="viewMode === 'monthly' ? selectedMonth : null"
          @update:model-year="updateYear"
          @update:model-month="updateMonth"
        />
      </div>

      <!-- 4. Section รายการเดินบัญชี (Statement) -->
      <div class="card border border-base-300 bg-base-100 shadow-sm rounded-2xl">
        <div class="card-body p-5 sm:p-6">
          <div class="flex items-center justify-between mb-4">
            <div>
              <h2 class="text-xl font-bold">รายการเดินบัญชี (Statement)</h2>
              <p class="text-sm text-base-content/60">
                ประวัติรายการเงินเข้าและเงินออกของบัญชีนี้
              </p>
            </div>
            <TotalList :total="transactions.length" />
          </div>

          <!-- ใช้ GroupedTransactionList component -->
          <GroupedTransactionList
            :transactions="transactions"
            :accounts="accountStore.accounts"
            :categories="categoryStore.categories"
            :loading="isLoadingTransactions"
            empty-title="ยังไม่มีรายการเดินบัญชี"
            empty-description="ไม่มีรายการเงินเข้าหรือเงินออกในช่วงเวลาที่เลือก"
            @edit="openEditTransaction"
            @delete="handleDeleteTransaction"
          />
        </div>
      </div>
    </template>
  </section>

  <!-- ฟอร์มแก้ไขรายการ -->
  <TransactionForm
    :open="isDialogOpen"
    :form="form"
    :accounts="accountStore.accounts"
    :categories="categoryStore.categories"
    @close="closeDialog"
    @save="saveTransaction"
  />
</template>
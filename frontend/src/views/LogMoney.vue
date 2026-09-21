<script setup lang="ts">
import { nextTick, onMounted, ref } from "vue"
import { Plus } from "lucide-vue-next"
import type { Transaction } from "@/types/transaction"
import { formatDate } from "@/utils/format"
import { useAccountStore } from "@/stores/account"
import { useCategoryStore } from "@/stores/category"
import { useTransactions } from "@/composables/transaction/useTransaction"
import { useTransactionForm } from "@/composables/transaction/useTransactionForm"
import { useTransactionFilter } from "@/composables/transaction/useTransactionFilter"
import DailyTransactionList from "@/components/transaction/DailyTransactionList.vue"
import TransCalendar from "@/components/transaction/TransCalendar.vue"
import TransactionSummary from "@/components/transaction/TransactionSummary.vue"
import TransactionForm from "@/components/transaction/TransactionForm.vue"
import AccountFilter from "@/components/filter/accountFilter.vue"
import CategoryFilter from "@/components/filter/categoryFilter.vue"
import { usePeriodFilter } from "@/composables/period/usePeriodFilter"

const isDialogOpen = ref(false)
const dailyTransactionSection = ref<HTMLElement | null>(null)
const accountStore = useAccountStore()
const categoryStore = useCategoryStore()
const {
  years,
  months,
  selectedYear,
  selectedMonth,
  loadYears,
} = usePeriodFilter("transaction")

const now = new Date()
selectedYear.value = now.getFullYear()
selectedMonth.value = now.getMonth() + 1

const {
  transactions,
  isLoadingTransactions,
  editingTransactionId,
  loadTransactions,
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
  fetchTransactions,
  handleMonthChange,
  selectedDate,
  selectedAccountId,
  // selectedCategory,
  // clientAccountId,
 
  clientCategory,
  filteredTransactions,
  selectedDayTransactions,
  selectDate,
} = useTransactionFilter(transactions,selectedYear,selectedMonth,loadTransactions)

function openEditTransaction(transaction: Transaction) {
  startEdit(transaction)
  setEditForm(transaction)
  isDialogOpen.value = true
}

function openAddTransaction() {
  cancelEdit()
  resetForm(
    formatDate(selectedDate.value),
  )

  form.value.transactionDate =
    formatDate(selectedDate.value)

  isDialogOpen.value = true
}

async function handleSelectDate(date: Date) {
  selectDate(date)

  form.value.transactionDate =
    formatDate(date)

  await nextTick()

  dailyTransactionSection.value?.scrollIntoView({
    behavior: "smooth",
    block: "start",
  })
}

function openTodayTransaction() {
  const today = new Date()
  handleSelectDate(today)
  openAddTransaction()
}

async function saveTransaction() {
  try {
    await saveTransactionApi(form.value)
    fetchTransactions()
    // await loadTransactions(
    //   selectedDate.value.getFullYear(),
    //   selectedDate.value.getMonth() + 1,
    // )
    cancelEdit()

    isDialogOpen.value = false

    resetForm(
      formatDate(selectedDate.value),
    )
  } catch (error) {
    console.error(
      "Transaction failed:",
      error,
    )
  }
}

async function handleDeleteTransaction(id: number) {
  try {
    await deleteTransaction(id)
    fetchTransactions()
  } catch (error) {
    console.error(
      "Delete transaction failed:",
      error,
    )
  }
}


function closeDialog() {
  isDialogOpen.value = false
}

onMounted(async () => {
  // 1. โหลดข้อมูลบัญชี และ หมวดหมู่ พร้อมกัน
  await Promise.all([
    accountStore.loadAccounts(),
    categoryStore.loadCategories(),
  ])

  // 2. ตั้งค่า Default เป็นบัญชีแรกที่มี (ถ้ามี)
  const firstAccount = accountStore.accounts[0]
  if (firstAccount) {
    // clientAccountId.value = firstAccount.id
    selectedAccountId.value = firstAccount.id
  }

  // 3. ยิงดึงข้อมูล Transaction ประจำเดือน
  await fetchTransactions()
})

</script>

<template>
  <section class="space-y-6!">
    <!-- Header -->
    <div class="flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
      <!-- Title -->
      <div>
        <h1 class="text-2xl font-bold">
          Finance Calendar
        </h1>

        <p class="text-base-content/60">
          เลือกวันที่เพื่อบันทึกรายรับหรือรายจ่าย
        </p>
      </div>

      <!-- Actions -->
      <div class="flex flex-col gap-2 sm:flex-row sm:items-end">
        <CategoryFilter v-model="clientCategory" :categories="categoryStore.categories" />
        <AccountFilter v-model="selectedAccountId" :accounts="accountStore.accounts" /> 
        <button class="btn text-white bg-green-700" type="button" @click="openTodayTransaction">
          <Plus class="size-4" />
          เพิ่มรายการวันนี้
        </button>
      </div>
    </div>

    <!-- Monthly Summary -->
    <TransactionSummary :transactions="filteredTransactions" />

    <!-- Calendar -->
    <TransCalendar :selected-date="selectedDate" :transactions="filteredTransactions" @select="handleSelectDate" @month-change="handleMonthChange" />

    <div ref="dailyTransactionSection">
      <DailyTransactionList :selected-date="selectedDate" :categories="categoryStore.categories" :transactions="selectedDayTransactions"
        :accounts="accountStore.accounts" :loading="isLoadingTransactions" @add="openAddTransaction" @edit="openEditTransaction" @delete="handleDeleteTransaction"/>
    </div>
  </section>

  <TransactionForm :open="isDialogOpen" :form="form" :accounts="accountStore.accounts"
    :categories="categoryStore.categories" @close="closeDialog"
    @save="saveTransaction" />

</template>

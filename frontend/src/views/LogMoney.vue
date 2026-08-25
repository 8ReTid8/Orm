<script setup lang="ts">
import { computed, nextTick, onMounted, ref } from "vue"
import { Plus } from "lucide-vue-next"
import type { Transaction } from "@/types/transaction"
import LogMoneyForm from "@/components/logMoney/LogMoneyForm.vue"
import DailyTransactionList from "@/components/logMoney/DailyTransactionList.vue"
import { useAccountStore } from "@/stores/account"
import { useCategoryStore } from "@/stores/category"
import TransCalendar from "@/components/logMoney/TransCalendar.vue"

import { useTransactions } from "@/composables/useTransaction"
import { useTransactionForm } from "@/composables/useTransactionForm"
import { useTransactionFilter } from "@/composables/useTransactionFilter"
import { formatDate } from "@/utils/date"
import TransactionSummary from "@/components/logMoney/TransactionSummary.vue"
import AccountFilter from "@/components/filter/accountFilter.vue"

// const selectedDate = ref<Date>(new Date())
// const selectedAccountId = ref<number | null>(null)
const isDialogOpen = ref(false)
const dailyTransactionSection = ref<HTMLElement | null>(null)
const accountStore = useAccountStore()
const categoryStore = useCategoryStore()
const {
  transactions,
  isLoadingTransactions,
  editingTransactionId,
  loadTransactions,
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
  selectedDate,
  selectedAccountId,
  filteredTransactions,
  selectedDateText,
  selectedDayTransactions,
  selectDate,
} = useTransactionFilter(transactions)

function openEditTransaction(transaction: Transaction) {
  // editingTransactionId.value = transaction.id
  startEdit(transaction)
  setEditForm(transaction)
  isDialogOpen.value = true
}

function openAddTransaction() {
  // editingTransactionId.value = null
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

// async function saveTransaction() {
//   if (!form.value.amount || !form.value.title || !form.value.category || !form.value.accountId) {
//     return
//   }
//   try {
//     if (editingTransactionId.value) {
//       await updateTransaction(
//         editingTransactionId.value,
//         form.value,
//       )
//     } else {
//       await createTransaction(
//         form.value,
//       )
//     }

//     await Promise.all([
//       loadTransactions(
//         selectedDate.value.getFullYear(),
//         selectedDate.value.getMonth() + 1,
//       ),

//     ])

//     editingTransactionId.value = null
//     isDialogOpen.value = false
//     resetForm()
//   } catch (error) {
//     console.error("Create transaction failed:", error)
//   }
// }

async function saveTransaction() {
  try {
    await saveTransactionApi(form.value)

    await loadTransactions(
      selectedDate.value.getFullYear(),
      selectedDate.value.getMonth() + 1,
    )

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

function closeDialog() {
  isDialogOpen.value = false
}

onMounted(() => {
  const now = new Date()

  loadTransactions(
    now.getFullYear(),
    now.getMonth() + 1,
  )

  categoryStore.loadCategories()
  accountStore.loadAccounts()
})

</script>

<template>
  <section class="!space-y-6">
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
        <AccountFilter v-model="selectedAccountId" :accounts="accountStore.accounts" />

        <button class="btn btn-primary" type="button" @click="openTodayTransaction">
          <Plus class="size-4" />
          เพิ่มรายการวันนี้
        </button>
      </div>
    </div>

    <!-- Monthly Summary -->
    <TransactionSummary :transactions="filteredTransactions" />

    <!-- Calendar -->
    <TransCalendar :selected-date="selectedDate" :transactions="filteredTransactions" @select="handleSelectDate" />

    <div ref="dailyTransactionSection">
      <DailyTransactionList :selected-date-text="selectedDateText" :transactions="selectedDayTransactions"
        :loading="isLoadingTransactions" @add="openAddTransaction" @edit="openEditTransaction" />
    </div>
  </section>

  <LogMoneyForm :open="isDialogOpen" :form="form" :accounts="accountStore.accounts"
    :categories="categoryStore.categories" :selected-date-text="selectedDateText" @close="closeDialog"
    @save="saveTransaction" />

</template>

<script setup lang="ts">
import { computed, onMounted, ref } from "vue"
import { Plus } from "lucide-vue-next"
import type { TransactionForm, Transaction } from "@/types/transaction"
import LogMoneyForm from "@/components/logMoney/LogMoneyForm.vue"
import DailyTransactionList from "@/components/logMoney/DailyTransactionList.vue"
import { useAccountStore } from "@/stores/account"
import { useCategoryStore } from "@/stores/category"
import { createTransaction, getTransactions } from "@/services/transaction"
import TransCalendar from "@/components/logMoney/TransCalendar.vue"
const selectedDate = ref<Date>(new Date())
const isDialogOpen = ref(false)
const transactions = ref<Transaction[]>([])
const accountStore = useAccountStore()
const categoryStore = useCategoryStore()
const editingTransactionId = ref<number | null>(null)
const isLoadingTransactions = ref(false)
const form = ref<TransactionForm>({
  type: "expense",
  amount: null,
  category: "",
  accountId: null,
  title: "",
  note: "",
  transactionDate: formatDate(new Date()),
  slipImage: null,
})

const selectedDateText = computed(() => {
  return selectedDate.value.toLocaleDateString("th-TH", {
    day: "numeric",
    month: "long",
    year: "numeric",
  })
})

const selectedDayTransactions = computed(() => {
  const date = formatDate(selectedDate.value)

  return transactions.value.filter(
    transaction =>
      transaction.transactionDate.startsWith(date)
  )
})
function openEditTransaction(transaction: Transaction) {
  editingTransactionId.value = transaction.id
  form.value = {
    type: transaction.type,
    amount: transaction.amount,
    category: transaction.category,
    accountId: transaction.account.id,
    title: transaction.title,
    note: transaction.note,
    transactionDate: transaction.transactionDate.slice(0, 10),
    slipImage: null,
  }
  isDialogOpen.value = true
}
function openAddTransaction() {
  editingTransactionId.value = null
  resetForm()

  form.value.transactionDate =
    formatDate(selectedDate.value)

  isDialogOpen.value = true
}
function formatDate(date: Date): string {
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, "0")
  const day = String(date.getDate()).padStart(2, "0")

  return `${year}-${month}-${day}`
}

function selectDate(date: Date) {
  selectedDate.value = date
  form.value.transactionDate = formatDate(date)
}

function addTodayTransaction() {
  selectedDate.value = new Date()
  form.value.transactionDate = formatDate(new Date())
  isDialogOpen.value = true
}

async function saveTransaction() {
  if (!form.value.amount || !form.value.title || !form.value.category || !form.value.accountId) {
    return
  }

  try {
    console.log("form before send:", form.value)
    const result = await createTransaction(form.value)
    console.log("Transaction:", result.transaction)

    isDialogOpen.value = false
    resetForm()
  } catch (error) {
    console.error("Create transaction failed:", error)
  }
  // console.log("Transaction:", form.value)

  // isDialogOpen.value = false
  // resetForm()
}

function resetForm() {
  form.value = {
    type: "expense",
    amount: null,
    category: "",
    accountId: null,
    title: "",
    note: "",
    transactionDate: formatDate(selectedDate.value),
    slipImage: null,
  }
}

function closeDialog() {
  isDialogOpen.value = false
}

async function loadTransactions(
  year: number,
  month: number,
) {
  try {
    isLoadingTransactions.value = true

    transactions.value = await getTransactions(
      year,
      month,
    )
  } catch (error) {
    console.error(
      "Failed to load transactions:",
      error,
    )
  } finally {
    isLoadingTransactions.value = false
  }
}

console.log(localStorage.getItem("token"))
onMounted(() => {
  const now = new Date()

  loadTransactions(
    now.getFullYear(),
    now.getMonth() + 1,
  )
  // loadFormOptions()
  categoryStore.loadCategories()
  accountStore.loadAccounts()
})

</script>

<template>
  <section class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold">Finance Calendar</h1>
        <p class="text-base-content/60">
          เลือกวันที่เพื่อบันทึกรายรับหรือรายจ่าย
        </p>
      </div>

      <button class="btn btn-primary" type="button" @click="selectDate(new Date())">
        <Plus class="size-4" />
        เพิ่มรายการวันนี้
      </button>
    </div>
    <TransCalendar :selected-date="selectedDate" :transactions="transactions" @select="selectDate" />
    <DailyTransactionList :selected-date-text="selectedDateText" :transactions="selectedDayTransactions"
      :loading="isLoadingTransactions" @add="openAddTransaction" @edit="openEditTransaction" />
  </section>
  <LogMoneyForm :open="isDialogOpen" :form="form" :accounts="accountStore.accounts"
    :categories="categoryStore.categories" :selected-date-text="selectedDateText" @close="closeDialog"
    @save="saveTransaction" />
</template>

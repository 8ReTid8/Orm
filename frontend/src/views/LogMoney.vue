<script setup lang="ts">
import { computed, onMounted, ref } from "vue"
import { CalendarDays, Plus } from "lucide-vue-next"
import type { TransactionForm, Transaction } from "@/types/transaction"
import LogMoneyForm from "@/components/logMoney/LogMoneyForm.vue"
import DailyTransactionList from "@/components/logMoney/DailyTransactionList.vue"
import type { Category } from "@/types/category"
import { getBanks } from "@/services/bank"
import { getCategories } from "@/services/category"
import type { Bank } from "@/types/bank"
const selectedDate = ref<Date>(new Date())
// const dialogRef = ref<HTMLDialogElement | null>(null)
const isDialogOpen = ref(false)
// const transactions = ref<Transaction[]>([])
const transactions = ref<Transaction[]>([
  {
    id: 1,
    type: "expense",
    amount: 120,
    category: "อาหาร",
    bankName: "KBank",
    title: "ข้าวกลางวัน",
    note: "",
    transactionDate: formatDate(new Date()),
  },
  {
    id: 2,
    type: "income",
    amount: 500,
    category: "รายได้",
    bankName: "SCB",
    title: "ค่าขนม",
    note: "",
    transactionDate: formatDate(new Date()),
  },
])
const banks = ref<Bank[]>([])
const categories = ref<Category[]>([])
const isLoadingTransactions = ref(false)
const form = ref<TransactionForm>({
  type: "expense",
  amount: null,
  category: "",
  bankId: null,
  title: "",
  note: "",
  transactionDate: formatDate(new Date()),
  slipImage: null,
})
async function loadFormOptions() {
  try {
    const [bankResult, categoryResult] =
      await Promise.all([
        getBanks(),
        getCategories(),
      ])

    banks.value = bankResult
    categories.value = categoryResult

  } catch (error) {
    console.error(
      "Failed to load form options:",
      error
    )
  }
}
onMounted(() => {
  loadFormOptions()
})
const selectedDateText = computed(() => {
  return selectedDate.value.toLocaleDateString("th-TH", {
    day: "numeric",
    month: "long",
    year: "numeric",
  })
})

function formatDate(date: Date): string {
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, "0")
  const day = String(date.getDate()).padStart(2, "0")

  return `${year}-${month}-${day}`
}

function selectDate(date: Date) {
  selectedDate.value = date
  form.value.transactionDate = formatDate(date)
  // dialogRef.value?.showModal()
  // isDialogOpen.value = true
}
function addTodayTransaction() {
  selectedDate.value = new Date()
  form.value.transactionDate = formatDate(new Date())
  isDialogOpen.value = true
}
function saveTransaction() {
  if (!form.value.amount || !form.value.title || !form.value.category) {
    return
  }

  console.log("Transaction:", form.value)

  // dialogRef.value?.close()
  isDialogOpen.value = false
  resetForm()
}

function resetForm() {
  form.value = {
    type: "expense",
    amount: null,
    category: "",
    bankId: null,
    title: "",
    note: "",
    transactionDate: formatDate(selectedDate.value),
    slipImage: null,
  }
}
function closeDialog() {
  isDialogOpen.value = false
}

console.log(localStorage.getItem("token"))
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

    <!-- Calendar card -->
    <div class="card border border-base-300 bg-base-100 shadow-sm">
      <div class="card-body">
        <div class="mb-3 flex items-center gap-2">
          <CalendarDays class="size-5" />
          <h2 class="card-title">ปฏิทินรายรับรายจ่าย</h2>
        </div>
        <VCalendar class="finance-calendar w-full" expanded borderless>
          <template #day-content="{ day }">
            <button type="button" class="
        relative
        block
        min-h-28
        w-full
        cursor-pointer
        bg-base-100
        p-3
        pt-10
        text-left
        transition-colors
        hover:bg-base-200
      " @click="selectDate(day.date)">
              <span class="
          absolute
          right-3
          top-2
          flex
          size-7
          items-center
          justify-center
          rounded-full
          text-sm
          font-medium
        ">
                {{ day.day }}
              </span>
            </button>
          </template>
        </VCalendar>
      </div>
    </div>
    <DailyTransactionList :selected-date-text="selectedDateText" :transactions="transactions"
      :loading="isLoadingTransactions" @add="isDialogOpen = true" />

    <!-- Transaction dialog -->
  </section>
  <LogMoneyForm :open="isDialogOpen" :form="form" :banks="banks" :categories="categories"
    :selected-date-text="selectedDateText" @close="closeDialog" @save="saveTransaction" />
</template>

<style>
.finance-calendar .vc-weeks {
  border-top: 1px solid var(--color-base-300);
  /* border-left: 1px solid var(--color-base-300); */
}

.finance-calendar .vc-day {
  min-height: 112px;
  padding: 0 !important;
  border-right: 1px solid var(--color-base-300);
  border-bottom: 1px solid var(--color-base-300);
  box-sizing: border-box;
}
</style>

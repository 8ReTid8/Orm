<script setup lang="ts">
import { computed, ref } from "vue"
import { CalendarDays, Plus } from "lucide-vue-next"
import type { TransactionForm } from "@/types/transaction"
import LogMoneyForm from "@/components/logMoney/LogMoneyForm.vue"

const selectedDate = ref<Date>(new Date())
// const dialogRef = ref<HTMLDialogElement | null>(null)
const isDialogOpen = ref(false)

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

        <!-- <VCalendar expanded borderless :attributes="[
          {
            key: 'selected-day',
            highlight: true,
            dates: selectedDate,
          },
        ]" @dayclick="selectDate($event.date)" /> -->
        <!-- <VCalendar class="finance-calendar" expanded borderless :attributes="[
          {
            key: 'selected-day',
            highlight: true,
            dates: selectedDate,
          },
        ]" @dayclick="selectDate($event.date)" /> -->
        <!-- <VCalendar class="finance-calendar" expanded borderless :attributes="[
          {
            key: 'selected-day',
            highlight: true,
            dates: selectedDate,
          },
        ]">
          <template #day-content="{ day }">
            <button type="button" class="calendar-day-cell" @click="selectDate(day.date)">
              <span class="calendar-day-number">
                {{ day.day }}
              </span>

             
              <div class="calendar-day-details">
               
              </div>
            </button>
          </template>
</VCalendar> -->
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

    <!-- Transaction dialog -->
    <LogMoneyForm :open="isDialogOpen" :form="form" :selected-date-text="selectedDateText" @close="closeDialog"
      @save="saveTransaction" />

  </section>
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
<!-- <style>
.finance-calendar .vc-weeks {
  border-top: 1px solid var(--color-base-300);
  border-left: 1px solid var(--color-base-300);
}

.finance-calendar .vc-day {
  padding: 0 !important;
}
</style> -->

<!-- <style>
.finance-calendar.vc-container {
  width: 100%;
  border: 0;
}

/* ตารางวันที่ */
.finance-calendar .vc-weeks {
  border-top: 1px solid #e5e7eb;
  border-left: 1px solid #e5e7eb;
}

/* ช่องแต่ละวัน */
.finance-calendar .vc-day {
  height: 100px !important;
  min-height: 100px !important;
  padding: 8px;

  border-right: 1px solid #e5e7eb;
  border-bottom: 1px solid #e5e7eb;

  box-sizing: border-box;
  cursor: pointer;
}

/* Hover ทั้งช่อง */
.finance-calendar .vc-day:hover {
  background-color: #f5f5f5;
}

/* เลขวันที่ */
.finance-calendar .vc-day-content {
  width: 32px;
  height: 32px;
}

/* ชื่อวัน อา จ อ ... */
.finance-calendar .vc-weekday {
  padding: 12px 0;
  font-weight: 600;
}

.finance-calendar .vc-day-content {
  position: absolute;

  top: 8px;
  right: 8px;

  width: 28px;
  height: 28px;

  display: flex;
  align-items: center;
  justify-content: center;

  border-radius: 9999px;
}
</style> -->
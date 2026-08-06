<script setup lang="ts">
import { computed, ref } from "vue"
import { CalendarDays, Plus } from "lucide-vue-next"
import type { TransactionForm } from "@/types/transaction"

const selectedDate = ref<Date>(new Date())
const dialogRef = ref<HTMLDialogElement | null>(null)

const form = ref<TransactionForm>({
  type: "expense",
  amount: null,
  category: "",
  title: "",
  note: "",
  transactionDate: formatDate(new Date()),
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
  dialogRef.value?.showModal()
}

function saveTransaction() {
  if (!form.value.amount || !form.value.title || !form.value.category) {
    return
  }

  console.log("Transaction:", form.value)

  dialogRef.value?.close()
  resetForm()
}

function resetForm() {
  form.value = {
    type: "expense",
    amount: null,
    category: "",
    title: "",
    note: "",
    transactionDate: formatDate(selectedDate.value),
  }
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

        <VCalendar expanded borderless :attributes="[
          {
            key: 'selected-day',
            highlight: true,
            dates: selectedDate,
          },
        ]" @dayclick="selectDate($event.date)" />
      </div>
    </div>

    <!-- Transaction dialog -->
    <dialog ref="dialogRef" class="modal">
      <div class="modal-box max-w-lg">
        <h2 class="text-xl font-bold">เพิ่มรายการ</h2>

        <p class="mt-1 text-sm text-base-content/60">
          วันที่ {{ selectedDateText }}
        </p>

        <form class="mt-6 space-y-4" @submit.prevent="saveTransaction">
          <!-- Type -->
          <fieldset>
            <legend class="mb-2 font-medium">ประเภทรายการ</legend>

            <div class="grid grid-cols-2 gap-3">
              <label class="btn" :class="form.type === 'income'
                ? 'btn-success'
                : 'btn-outline'">
                <input v-model="form.type" type="radio" value="income" class="hidden" />
                รายรับ
              </label>

              <label class="btn" :class="form.type === 'expense'
                ? 'btn-error'
                : 'btn-outline'">
                <input v-model="form.type" type="radio" value="expense" class="hidden" />
                รายจ่าย
              </label>
            </div>
          </fieldset>

          <!-- Title -->
          <fieldset class="fieldset">
            <legend class="fieldset-legend">ชื่อรายการ</legend>

            <input v-model.trim="form.title" type="text" class="input w-full" placeholder="เช่น ค่าอาหาร" required />
          </fieldset>

          <!-- Amount -->
          <fieldset class="fieldset">
            <legend class="fieldset-legend">จำนวนเงิน</legend>

            <input v-model.number="form.amount" type="number" min="0.01" step="0.01" class="input w-full"
              placeholder="0.00" required />
          </fieldset>

          <!-- Category -->
          <fieldset class="fieldset">
            <legend class="fieldset-legend">หมวดหมู่</legend>

            <select v-model="form.category" class="select w-full" required>
              <option disabled value="">เลือกหมวดหมู่</option>
              <option value="food">อาหาร</option>
              <option value="transport">การเดินทาง</option>
              <option value="salary">เงินเดือน</option>
              <option value="shopping">ช้อปปิ้ง</option>
              <option value="other">อื่น ๆ</option>
            </select>
          </fieldset>

          <!-- Note -->
          <fieldset class="fieldset">
            <legend class="fieldset-legend">หมายเหตุ</legend>

            <textarea v-model.trim="form.note" class="textarea w-full" placeholder="รายละเอียดเพิ่มเติม" />
          </fieldset>

          <div class="modal-action">
            <button type="button" class="btn btn-ghost" @click="dialogRef?.close()">
              ยกเลิก
            </button>

            <button type="submit" class="btn btn-primary">
              บันทึก
            </button>
          </div>
        </form>
      </div>

      <form method="dialog" class="modal-backdrop">
        <button aria-label="Close dialog">close</button>
      </form>
    </dialog>
  </section>
</template>
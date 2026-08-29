<script setup lang="ts">
import { ref } from "vue"
import {
  Plus,
  ReceiptText,
  Pencil,
  Trash2,
  ArrowUpRight,
  ArrowDownLeft,
  FileText,
  AlertTriangle
} from "lucide-vue-next"

import type { Transaction } from "@/types/transaction"
import { formatMoney } from "@/utils/number.ts"
import TransactionSummary from "./TransactionSummary.vue"
import { resolveCategoryIcon } from "@/utils/categoryIcons.ts"
import type { Category } from "@/types/category.ts"

interface Props {
  selectedDateText: string
  transactions: Transaction[]
  categories: Category[]
  loading?: boolean
}

// เก็บ transaction ที่กำลังจะถูกลบ (ถ้าเป็น null แปลว่า modal ปิดอยู่)
const transactionToDelete = ref<Transaction | null>(null)
// เมื่อกดปุ่มถังขยะ ให้เปิด modal ยืนยัน
function openConfirmDelete(transaction: Transaction) {
  transactionToDelete.value = transaction
}
// เมื่อกดยกเลิก
function closeConfirmDelete() {
  transactionToDelete.value = null
}
// เมื่อกดยืนยันการลบ
function handleConfirmDelete() {
  if (transactionToDelete.value) {
    emit("delete", transactionToDelete.value.id)
    closeConfirmDelete()
  }
}

const props = defineProps<Props>()
// const dailyTransactionSection = ref<HTMLElement | null>(null)
const emit = defineEmits<{
  add: []
  edit: [transaction: Transaction]
  delete: [id: number]
}>()

function getCategoryIcon(categoryName: string) {
  const cat = props.categories.find(c => c.name === categoryName)
  return resolveCategoryIcon(cat?.icon)
}

</script>

<template>
  <dialog :class="{ 'modal-open': transactionToDelete !== null }" class="modal">
    <div class="modal-box max-w-sm text-center">
      <!-- Icon เตือน -->
      <!-- <div class="mx-auto mb-4 flex size-14 items-center justify-center rounded-full bg-error/10 text-error">
        <AlertTriangle class="size-7" />
      </div> -->
      <h3 class="text-lg font-bold">ยืนยันการลบรายการ?</h3>

      <p v-if="transactionToDelete" class="mt-2 text-sm text-base-content/70">
        คุณต้องการลบรายการ
        <span class="font-semibold text-base-content">"{{ transactionToDelete.title }}"</span>
        จำนวน <span class="font-semibold text-error">฿{{ formatMoney(transactionToDelete.amount) }}</span> ใช่หรือไม่?
      </p>
      <div class="modal-action justify-center gap-3 mt-6">
        <button type="button" class="btn btn-ghost" @click="closeConfirmDelete">
          ยกเลิก
        </button>
        <button type="button" class="btn btn-error text-white" @click="handleConfirmDelete">
          ลบรายการ
        </button>
      </div>
    </div>
    <!-- Backdrop คลิกข้างนอกเพื่อปิด -->
    <form method="dialog" class="modal-backdrop" @submit.prevent="closeConfirmDelete">
      <button>close</button>
    </form>
  </dialog>
  <div class="card border border-base-300 bg-base-100 shadow-sm">
    <div class="card-body">
      <!-- Header -->
      <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <h2 class="text-xl font-bold">
            รายการวันที่ {{ selectedDateText }}
          </h2>

          <p class="text-sm text-base-content/60">
            รายรับและรายจ่ายทั้งหมดของวันที่เลือก
          </p>
        </div>

        <button type="button" class="btn text-white bg-green-700" @click="emit('add')">
          <Plus class="size-4" />
          เพิ่มรายการ
        </button>
      </div>

      <!-- Summary -->
      <TransactionSummary :transactions="transactions" />

      <!-- Loading -->
      <div v-if="loading" class="flex justify-center py-12">
        <span class="loading loading-spinner loading-lg" />
      </div>

      <!-- Empty -->
      <div v-else-if="transactions.length === 0"
        class="mt-6 flex flex-col items-center justify-center rounded-xl border border-dashed border-base-300 py-12 text-center">
        <ReceiptText class="size-10 text-base-content/30" />

        <p class="mt-3 font-medium">
          ยังไม่มีรายการในวันนี้
        </p>

        <p class="mt-1 text-sm text-base-content/50">
          กดเพิ่มรายการเพื่อบันทึกรายรับหรือรายจ่าย
        </p>
      </div>

      <!-- Transactions -->
      <!-- <div v-else class="mt-6 divide-y divide-base-300">
        <div v-for="transaction in transactions" :key="transaction.id"
          class="flex items-center justify-between gap-4 py-4">
          <div class="min-w-0">
            <p class="font-semibold">
              {{ transaction.title }}
            </p>

            <div class="mt-1 flex flex-wrap items-center gap-x-2 text-sm text-base-content/50">
              <span>{{ transaction.category }}</span>
              <span>•</span>
              <span>{{ transaction.account.name }}</span>
            </div>
          </div>

          <div class="shrink-0 text-right">
            <p class="font-bold" :class="transaction.type === 'income'
                ? 'text-success'
                : 'text-error'
              ">
              {{ transaction.type === "income" ? "+" : "-" }}
              ฿{{ formatMoney(transaction.amount) }}
            </p>
            <button type="button" class="btn btn-square btn-ghost btn-sm" aria-label="แก้ไขรายการ"
              @click="emit('edit', transaction)">
              <Pencil class="size-4" />
            </button>
          </div>
        </div>
      </div> -->
      <!-- Transactions -->
      <div v-else class="mt-4 space-y-2">
        <div v-for="transaction in transactions" :key="transaction.id"
          class="group flex items-center justify-between gap-4 rounded-xl border border-transparent p-3.5 transition-all hover:border-base-300 hover:bg-base-200/50">
          <!-- Left: Icon & Info -->
          <div class="flex items-center gap-3.5 min-w-0">
            <!-- Type Icon Circle -->
            <div class="flex size-10 shrink-0 items-center justify-center rounded-xl font-bold" :class="transaction.type === 'income'
              ? 'bg-success/15 text-success'
              : 'bg-error/15 text-error'
              ">
              <!-- <ArrowDownLeft v-if="transaction.type === 'income'" class="size-5" /> -->
              <!-- <ArrowUpRight v-else class="size-5" /> -->
              <!-- Render dynamic Lucide icon ตามชื่อที่ดึงมา -->
              <component :is="getCategoryIcon(transaction.category)" class="size-5" />
            </div>

            <!-- Title & Details -->
            <div class="min-w-0">
              <div class="flex items-center gap-2">
                <p class="truncate font-semibold text-base">
                  {{ transaction.title }}
                </p>
                <!-- Badge ถ้ามี note หรือ slip -->
                <FileText v-if="transaction.note" class="size-3.5 text-base-content/40 shrink-0"
                  title="มีบันทึกข้อความ" />
              </div>

              <div class="mt-0.5 flex flex-wrap items-center gap-1.5 text-xs text-base-content/60">
                <span class="badge badge-sm badge-ghost font-normal">{{ transaction.category }}</span>
                <span>•</span>
                <span class="text-sm">{{ transaction.account.name }}</span>
              </div>
            </div>
          </div>

          <!-- Right: Amount & Action Buttons -->
          <div class="flex items-center gap-3 shrink-0">
            <p class="text-right font-bold text-base md:text-lg" :class="transaction.type === 'income'
              ? 'text-success'
              : 'text-error'
              ">
              {{ transaction.type === 'income' ? '+' : '-' }}฿{{ formatMoney(transaction.amount) }}
            </p>

            <!-- Actions -->
            <div class="flex items-center gap-1">
              <!-- Edit Button -->
              <button type="button"
                class="btn btn-ghost btn-xs sm:btn-sm btn-square text-base-content/60 hover:bg-base-300 hover:text-base-content"
                title="แก้ไขรายการ" @click="emit('edit', transaction)">
                <Pencil class="size-4" />
              </button>

              <!-- Delete Button -->
              <button type="button"
                class="btn btn-ghost btn-xs sm:btn-sm btn-square text-base-content/60 hover:bg-error/10 hover:text-error"
                title="ลบรายการ" @click="openConfirmDelete(transaction)">
                <Trash2 class="size-4" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
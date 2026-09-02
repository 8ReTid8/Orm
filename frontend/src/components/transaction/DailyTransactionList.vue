<script setup lang="ts">
import { ref } from "vue"
import {
  Plus,
  ReceiptText,
  FileText,
  Wallet,
} from "lucide-vue-next"
import type { Transaction } from "@/types/transaction"
import { formatMoney } from "@/utils/format"
import TransactionSummary from "./TransactionSummary.vue"
import { resolveCategoryIcon } from "@/utils/categoryIcons.ts"
import type { Category } from "@/types/category.ts"
import ConfirmModal from "../common/ConfirmModal.vue"
import ActionButton from "../common/ActionButton.vue"
import type { Account } from "@/types/account.ts"

interface Props {
  selectedDateText: string
  transactions: Transaction[]
  accounts: Account[]
  categories: Category[]
  loading?: boolean
}

// เก็บ transaction ที่ต้องการลบ
const targetTransaction = ref<Transaction | null>(null)

const props = defineProps<Props>()
const emit = defineEmits<{
  add: []
  edit: [transaction: Transaction]
  delete: [id: number]
}>()

function getCategoryIcon(categoryName: string) {
  const cat = props.categories.find(c => c.name === categoryName)
  return resolveCategoryIcon(cat?.icon)
}

function getAccountName(accountId: number) {
  const acc = props.accounts.find((a) => a.id === accountId)
  return acc ? acc.name : "ไม่พบบัญชี"
}

function openDeleteModal(transaction: Transaction) {
  targetTransaction.value = transaction
}
function handleConfirmDelete() {
  if (targetTransaction.value) {
    emit("delete", targetTransaction.value.id)
    targetTransaction.value = null
  }
}
</script>

<template>
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
                <span class="text-sm flex items-center gap-1">
                  <Wallet class="size-3.5" />
                  {{ getAccountName(transaction.accountId) }}
                </span>
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
            <!-- <div class="flex items-center gap-1"> -->
            <!-- Edit Button -->
            <!-- <button type="button"
                class="btn btn-ghost btn-xs sm:btn-sm btn-square text-base-content/60 hover:bg-base-300 hover:text-base-content"
                title="แก้ไขรายการ" @click="emit('edit', transaction)">
                <Pencil class="size-4" />
              </button> -->

            <!-- Delete Button -->
            <!-- <button type="button"
                class="btn btn-ghost btn-xs sm:btn-sm btn-square text-base-content/60 hover:bg-error/10 hover:text-error"
                title="ลบรายการ" @click="openDeleteModal(transaction)">
                <Trash2 class="size-4" />
              </button> -->
            <!-- </div> -->
            <ActionButton edit-title="แก้ไขรายการ" delete-title="ลบรายการ" @edit="emit('edit', transaction)"
              @delete="openDeleteModal(transaction)" />
          </div>
        </div>
      </div>
    </div>
  </div>
  <ConfirmModal :open="targetTransaction !== null" title="ยืนยันการลบรายการ?" confirm-text="ลบรายการ" type="danger"
    @close="targetTransaction = null" @confirm="handleConfirmDelete" />
</template>
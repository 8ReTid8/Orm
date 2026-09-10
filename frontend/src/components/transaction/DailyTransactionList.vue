<script setup lang="ts">
import { ref } from "vue"
import {
  Plus,
  ReceiptText,
} from "lucide-vue-next"
import type { Transaction } from "@/types/transaction"
import { formatThaiDateLong } from "@/utils/format"
import TransactionSummary from "./TransactionSummary.vue"
import type { Category } from "@/types/category.ts"
import ConfirmModal from "../common/ConfirmModal.vue"
import type { Account } from "@/types/account.ts"
import EmptyState from "../common/EmptyState.vue"
import TransactionList from "./TransactionList.vue"

interface Props {
  selectedDate: Date
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
            รายการวันที่ {{ formatThaiDateLong(selectedDate) }}
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
      <EmptyState v-else-if="transactions.length === 0" :icon="ReceiptText" title="ยังไม่มีรายการในวันนี้"
        description="กดเพิ่มรายการเพื่อบันทึกรายรับหรือรายจ่าย" />

      <!-- Transactions -->
      <div v-else class="mt-4 space-y-2">
        <TransactionList :transactions="transactions" :accounts="accounts" :categories="categories"
          @edit="emit('edit', $event)" @delete="emit('delete', $event)" />
      </div>
    </div>
  </div>
  <ConfirmModal :open="targetTransaction !== null" title="ยืนยันการลบรายการ?" confirm-text="ลบรายการ" type="danger"
    @close="targetTransaction = null" @confirm="handleConfirmDelete" />
</template>
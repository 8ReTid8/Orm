<script setup lang="ts">
import { computed, ref } from "vue"
import {
  Plus,
  ReceiptText,
  Pencil
} from "lucide-vue-next"

import type { Transaction } from "@/types/transaction"
import { formatMoney } from "@/utils/number.ts"
import TransactionSummary from "./TransactionSummary.vue"

interface Props {
  selectedDateText: string
  transactions: Transaction[]
  loading?: boolean
}

const props = defineProps<Props>()
const dailyTransactionSection = ref<HTMLElement | null>(null)
const emit = defineEmits<{
  add: []
  edit: [transaction: Transaction]
}>()

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
      <div v-else class="mt-6 divide-y divide-base-300">
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
      </div>
    </div>
  </div>
</template>
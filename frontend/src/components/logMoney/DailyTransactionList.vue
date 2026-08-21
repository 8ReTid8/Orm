<script setup lang="ts">
import { computed } from "vue"
import {
  ArrowDownLeft,
  ArrowUpRight,
  Plus,
  ReceiptText,
} from "lucide-vue-next"

import type { Transaction } from "@/types/transaction"

interface Props {
  selectedDateText: string
  transactions: Transaction[]
  loading?: boolean
}

const props = defineProps<Props>()

const emit = defineEmits<{
  add: []
}>()

const totalIncome = computed(() => {
  return props.transactions
    .filter((item) => item.type === "income")
    .reduce((total, item) => total + item.amount, 0)
})

const totalExpense = computed(() => {
  return props.transactions
    .filter((item) => item.type === "expense")
    .reduce((total, item) => total + item.amount, 0)
})

const balance = computed(() => {
  return totalIncome.value - totalExpense.value
})

function formatMoney(amount: number) {
  return new Intl.NumberFormat("th-TH", {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  }).format(amount)
}
</script>

<template>
  <div class="card border border-base-300 bg-base-100 shadow-sm">
    <div class="card-body">
      <!-- Header -->
      <div
        class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between"
      >
        <div>
          <h2 class="text-xl font-bold">
            รายการวันที่ {{ selectedDateText }}
          </h2>

          <p class="text-sm text-base-content/60">
            รายรับและรายจ่ายทั้งหมดของวันที่เลือก
          </p>
        </div>

        <button
          type="button"
          class="btn btn-primary"
          @click="emit('add')"
        >
          <Plus class="size-4" />
          เพิ่มรายการ
        </button>
      </div>

      <!-- Summary -->
      <div class="mt-6 grid grid-cols-1 gap-3 sm:grid-cols-3">
        <div class="rounded-xl bg-success/10 p-4">
          <div class="flex items-center gap-2 text-success">
            <ArrowDownLeft class="size-5" />
            <span class="text-sm font-medium">
              รายรับ
            </span>
          </div>

          <p class="mt-2 text-xl font-bold text-success">
            +฿{{ formatMoney(totalIncome) }}
          </p>
        </div>

        <div class="rounded-xl bg-error/10 p-4">
          <div class="flex items-center gap-2 text-error">
            <ArrowUpRight class="size-5" />
            <span class="text-sm font-medium">
              รายจ่าย
            </span>
          </div>

          <p class="mt-2 text-xl font-bold text-error">
            -฿{{ formatMoney(totalExpense) }}
          </p>
        </div>

        <div class="rounded-xl bg-base-200 p-4">
          <p class="text-sm font-medium text-base-content/60">
            สุทธิ
          </p>

          <p
            class="mt-2 text-xl font-bold"
            :class="balance >= 0 ? 'text-success' : 'text-error'"
          >
            ฿{{ formatMoney(balance) }}
          </p>
        </div>
      </div>

      <!-- Loading -->
      <div
        v-if="loading"
        class="flex justify-center py-12"
      >
        <span class="loading loading-spinner loading-lg" />
      </div>

      <!-- Empty -->
      <div
        v-else-if="transactions.length === 0"
        class="mt-6 flex flex-col items-center justify-center rounded-xl border border-dashed border-base-300 py-12 text-center"
      >
        <ReceiptText class="size-10 text-base-content/30" />

        <p class="mt-3 font-medium">
          ยังไม่มีรายการในวันนี้
        </p>

        <p class="mt-1 text-sm text-base-content/50">
          กดเพิ่มรายการเพื่อบันทึกรายรับหรือรายจ่าย
        </p>
      </div>

      <!-- Transactions -->
      <div
        v-else
        class="mt-6 divide-y divide-base-300"
      >
        <div
          v-for="transaction in transactions"
          :key="transaction.id"
          class="flex items-center justify-between gap-4 py-4"
        >
          <div class="min-w-0">
            <p class="font-semibold">
              {{ transaction.title }}
            </p>

            <div
              class="mt-1 flex flex-wrap items-center gap-x-2 text-sm text-base-content/50"
            >
              <span>{{ transaction.category }}</span>
              <span>•</span>
              <span>{{ transaction.account.name }}</span>
            </div>
          </div>

          <div class="shrink-0 text-right">
            <p
              class="font-bold"
              :class="
                transaction.type === 'income'
                  ? 'text-success'
                  : 'text-error'
              "
            >
              {{ transaction.type === "income" ? "+" : "-" }}
              ฿{{ formatMoney(transaction.amount) }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
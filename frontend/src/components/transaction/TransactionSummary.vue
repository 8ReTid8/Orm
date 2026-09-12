<script setup lang="ts">
import { computed } from "vue"
import {
  ArrowDownLeft,
  ArrowUpRight,
} from "lucide-vue-next"

import type { Transaction } from "@/types/transaction"
import { formatMoney } from "@/utils/format.ts"

interface Props {
  transactions: Transaction[]
}

const props = defineProps<Props>()

const totalIncome = computed(() => {
  return props.transactions
    .filter(item => item.type === "income")
    .reduce((total, item) => total + item.amount, 0)
})

const totalExpense = computed(() => {
  return props.transactions
    .filter(item => item.type === "expense")
    .reduce((total, item) => total + item.amount, 0)
})

const balance = computed(() => {
  return totalIncome.value - totalExpense.value
})

</script>

<template>
  <div class="grid grid-cols-1 gap-3 sm:grid-cols-3">

    <!-- Income -->
    <div class="rounded-xl bg-success/10 p-4">
      <div class="flex items-center gap-2 text-success">
        <ArrowDownLeft class="size-5" />
        <span class="text-lg font-medium">
          รายรับ
        </span>
      </div>

      <p class="mt-2 text-2xl font-bold text-success">
        +฿{{ formatMoney(totalIncome) }}
      </p>
    </div>

    <!-- Expense -->
    <div class="rounded-xl bg-error/10 p-4">
      <div class="flex items-center gap-2 text-error">
        <ArrowUpRight class="size-5" />
        <span class="text-lg font-medium">
          รายจ่าย
        </span>
      </div>

      <p class="mt-2 text-2xl font-bold text-error">
        -฿{{ formatMoney(totalExpense) }}
      </p>
    </div>

    <!-- Balance -->
    <div class="rounded-xl bg-base-200 p-4">
      <p class="text-lg font-medium text-base-content/60">
        สุทธิ
      </p>

      <p
        class="mt-2 text-2xl font-bold"
        :class="balance >= 0 ? 'text-success' : 'text-error'"
      >
        ฿{{ formatMoney(balance) }}
      </p>
    </div>

  </div>
</template>
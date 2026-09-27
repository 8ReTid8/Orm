<!-- <script setup lang="ts">
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
</template> -->

<script setup lang="ts">
import { computed } from "vue"
import {
  ArrowDownLeft,
  ArrowUpRight,
  Wallet,
  Percent,
} from "lucide-vue-next"
import type { Transaction } from "@/types/transaction"
import { formatMoney } from "@/utils/format"

interface Props {
  // รองรับการส่งตัวเลขตรงๆ (สำหรับ AccountDetail และ SummaryView)
  income?: number
  expense?: number
  net?: number
  savingsRate?: number // ถ้าส่งมาจะแสดงการ์ดใบที่ 4 ให้อัตโนมัติ

  // ปรับแต่งข้อความหัวการ์ดได้ (Default เป็น รายรับ, รายจ่าย, สุทธิ)
  incomeLabel?: string
  expenseLabel?: string
  netLabel?: string

  // หรือส่ง transactions มาให้คำนวณเอง (สำหรับ LogMoney และ DailyTransactionList)
  transactions?: Transaction[]
}

const props = withDefaults(defineProps<Props>(), {
  incomeLabel: "รายรับ",
  expenseLabel: "รายจ่าย",
  netLabel: "สุทธิ",
})

// คำนวณรายรับ (ถ้ามีส่ง income มาให้ใช้ค่านั้นก่อน ถ้าไม่มีค่อยวนลูป transactions)
const totalIncome = computed(() => {
  if (props.income !== undefined) return props.income
  return (
    props.transactions
      ?.filter((item) => item.type === "income")
      .reduce((sum, item) => sum + item.amount, 0) ?? 0
  )
})

// คำนวณรายจ่าย
const totalExpense = computed(() => {
  if (props.expense !== undefined) return props.expense
  return (
    props.transactions
      ?.filter((item) => item.type === "expense")
      .reduce((sum, item) => sum + item.amount, 0) ?? 0
  )
})

// คำนวณสุทธิ
const balance = computed(() => {
  if (props.net !== undefined) return props.net
  return totalIncome.value - totalExpense.value
})
</script>

<template>
  <div
    class="grid grid-cols-1 gap-3"
    :class="savingsRate !== undefined ? 'sm:grid-cols-2 lg:grid-cols-4' : 'sm:grid-cols-3'"
  >
    <!-- 1. รายรับ / เงินเข้า -->
    <div class="rounded-2xl bg-success/10 p-5 border border-success/20">
      <div class="flex items-center justify-between text-success">
        <span class="text-sm font-medium">{{ incomeLabel }}</span>
        <ArrowDownLeft class="size-5" />
      </div>
      <p class="mt-3 text-2xl font-bold text-success">
        +฿{{ formatMoney(totalIncome) }}
      </p>
    </div>

    <!-- 2. รายจ่าย / เงินออก -->
    <div class="rounded-2xl bg-error/10 p-5 border border-error/20">
      <div class="flex items-center justify-between text-error">
        <span class="text-sm font-medium">{{ expenseLabel }}</span>
        <ArrowUpRight class="size-5" />
      </div>
      <p class="mt-3 text-2xl font-bold text-error">
        -฿{{ formatMoney(totalExpense) }}
      </p>
    </div>

    <!-- 3. สุทธิคงเหลือ -->
    <div class="rounded-2xl bg-base-200 p-5 border border-base-300">
      <div class="flex items-center justify-between text-base-content/70">
        <span class="text-sm font-medium">{{ netLabel }}</span>
        <Wallet class="size-5" />
      </div>
      <p
        class="mt-3 text-2xl font-bold"
        :class="balance >= 0 ? 'text-success' : 'text-error'"
      >
        {{ balance >= 0 ? '+' : '' }}฿{{ formatMoney(balance) }}
      </p>
    </div>

    <!-- 4. อัตราการออม (จะแสดงเฉพาะเมื่อมีการส่ง savingsRate มา) -->
    <div
      v-if="savingsRate !== undefined"
      class="rounded-2xl bg-base-200 p-5 border border-base-300"
    >
      <div class="flex items-center justify-between text-base-content/70">
        <span class="text-sm font-medium">อัตราการออม</span>
        <Percent class="size-5" />
      </div>
      <div class="mt-3 flex items-baseline gap-2">
        <p class="text-2xl font-bold text-primary">
          {{ savingsRate.toFixed(1) }}%
        </p>
        <span class="text-xs text-base-content/60">ของรายรับ</span>
      </div>
    </div>
  </div>
</template>
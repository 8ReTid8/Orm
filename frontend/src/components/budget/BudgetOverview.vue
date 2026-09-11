<script setup lang="ts">
import type { Budget } from "@/types/budget"
import { formatDate, formatMoney, formatThaiDateLong } from "@/utils/format"
import { computed, ref } from "vue"
import type { Category } from "@/types/category.ts"
import type { Account } from "@/types/account.ts"
import { ChevronDown } from "lucide-vue-next"
import BudgetList from "./BudgetList.vue"

interface Props {
  budgets: Budget[]
  startDate: string
  endDate: string
  categories: Category[]
  accounts: Account[]
}
const props = defineProps<Props>()
const emit = defineEmits<{
  edit: [budget: Budget]
  delete: [id: number]
}>()

const isOpen = ref(false)
const totalBudget = computed(() => {
  return props.budgets.reduce(
    (sum, budget) => sum + budget.amount,
    0
  )
})

const totalSpent = computed(() => {
  return props.budgets.reduce(
    (sum, budget) => sum + budget.spent,
    0
  )
})

const remaining = computed(() => {
  return totalBudget.value - totalSpent.value
})

const percent = computed(() => {
  if (totalBudget.value === 0) return 0

  return Math.min(
    (totalSpent.value / totalBudget.value) * 100,
    100
  )
})

const statusColor = computed(() => {
  const p = (totalSpent.value / totalBudget.value) * 100
  if (p >= 100) return 'error'    // over/at budget — red
  if (p >= 80) return 'warning'   // getting close — yellow/orange
  return 'success'                // healthy — green
})
</script>

<template>
  <div class="card border border-base-300 bg-base-100 shadow-sm">
    <button type="button" class="flex w-full items-center justify-between px-6 pt-5 text-left"
      @click="isOpen = !isOpen">
      <div>
        <p class="font-semibold text-lg">
          {{ formatThaiDateLong(new Date(startDate)) }}
          -
          {{ formatThaiDateLong(new Date(endDate)) }}
        </p>
      </div>
      <ChevronDown class="size-5 transition-transform duration-200" :class="{ 'rotate-180': isOpen }" />
    </button>

    <div class="card-body">
      <div class="grid gap-6 md:grid-cols-2">

        <!-- Progress -->
        <div class="flex flex-col items-center justify-center">
          <div class="relative flex size-48 items-center justify-center">

            <!-- base track -->
            <div class="absolute inset-0 rounded-full bg-base-200" />

            <!-- progress fill, shaped via mask instead of background -->
            <div class="absolute inset-0 rounded-full transition-colors duration-300" :class="{
              'bg-success': statusColor === 'success',
              'bg-warning': statusColor === 'warning',
              'bg-error': statusColor === 'error',
            }" :style="{
              mask: `conic-gradient(#000 ${percent}%, transparent ${percent}% 100%)`,
              WebkitMask: `conic-gradient(#000 ${percent}%, transparent ${percent}% 100%)`
            }" />
            <!-- inner mask to turn the filled disc into a ring -->
            <div class="absolute inset-4.5 rounded-full bg-base-100" />

            <div class="relative text-center">
              <p class="text-sm text-base-content/50">
                ใช้ไป
              </p>

              <!-- <p class="text-2xl font-bold"> -->
              <p class="text-2xl font-bold transition-colors duration-300" :class="{
                'text-success': statusColor === 'success',
                'text-warning': statusColor === 'warning',
                'text-error': statusColor === 'error',
              }">
                ฿{{ formatMoney(totalSpent) }}
              </p>

              <p class="text-sm text-base-content/50">
                จาก ฿{{ formatMoney(totalBudget) }}
              </p>
            </div>

          </div>

          <p class="mt-3! text-sm text-base-content/50">
            ใช้ไป {{ percent.toFixed(0) }}% ของงบประมาณ
          </p>

        </div>

        <!-- Summary -->
        <div class="grid grid-cols-2 gap-3">

          <div class="rounded-xl bg-base-200 p-4">
            <p class="text-sm text-base-content/50">
              งบประมาณ
            </p>

            <p class="mt-2 text-xl font-bold">
              ฿{{ formatMoney(totalBudget) }}
            </p>
          </div>

          <div class="rounded-xl bg-error/10 p-4">
            <p class="text-sm text-error">
              ใช้ไป
            </p>

            <p class="mt-2 text-xl font-bold text-error">
              ฿{{ formatMoney(totalSpent) }}
            </p>
          </div>

          <div class="col-span-2 rounded-xl bg-success/10 p-4">
            <p class="text-sm text-success">
              เหลือ
            </p>

            <p class="mt-2 text-2xl font-bold text-success">
              ฿{{ formatMoney(remaining) }}
            </p>
          </div>

        </div>
      </div>
      <div v-if="isOpen" class="border-t border-base-300 px-6 pb-6">
        <div class="flex items-center justify-between mt-3!">
          <div>
            <h2 class="text-xl font-bold">
              งบประมาณแยกตามหมวด
            </h2>
            <p class="text-sm text-base-content/60">
              ติดตามและควบคุมการใช้จ่ายตามช่วงเวลาที่กำหนด
            </p>
          </div>
          <span class="badge badge-neutral font-medium">
            {{ budgets.length }} รายการ
          </span>
        </div>
        <!-- <BudgetList :budgets="budgets" :categories="categories" :accounts="accounts" @edit="emit('edit', $event)"
          @delete="emit('delete', $event)" /> -->
        <BudgetList v-for="budget in budgets" :key="budget.id" :budget="budget" :categories="categories" :accounts="accounts" @edit="emit('edit', $event)"
          @delete="emit('delete', $event)" />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import type { Budget } from "@/types/budget"
import { formatMoney } from "@/utils/format"
import { computed } from "vue"

interface Props {
  budgets: Budget[]
}
const props = defineProps<Props>()

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
    <div class="card-body">

      <div class="grid gap-6 md:grid-cols-2">

        <!-- Progress -->
        <div class="flex flex-col items-center justify-center">

          <div class="relative flex size-48 items-center justify-center">

            <!-- base track -->
            <div class="absolute inset-0 rounded-full bg-base-200" />

            <!-- progress fill, shaped via mask instead of background -->
            <!-- <div class="absolute inset-0 rounded-full bg-success" :style="{
              mask: `conic-gradient(#000 ${percent}%, transparent ${percent}% 100%)`,
              WebkitMask: `conic-gradient(#000 ${percent}%, transparent ${percent}% 100%)`
            }" /> -->
            <div class="absolute inset-0 rounded-full transition-colors duration-300" :class="{
              'bg-success': statusColor === 'success',
              'bg-warning': statusColor === 'warning',
              'bg-error': statusColor === 'error',
            }" :style="{
              mask: `conic-gradient(#000 ${percent}%, transparent ${percent}% 100%)`,
              WebkitMask: `conic-gradient(#000 ${percent}%, transparent ${percent}% 100%)`
            }" />
            <!-- inner mask to turn the filled disc into a ring -->
            <div class="absolute inset-[18px] rounded-full bg-base-100" />

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

              <p class="text-xs text-base-content/50">
                จาก ฿{{ formatMoney(totalBudget) }}
              </p>
            </div>

          </div>

          <p class="mt-3 text-sm text-base-content/50">
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

    </div>
  </div>
</template>
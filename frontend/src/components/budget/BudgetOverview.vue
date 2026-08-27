<script setup lang="ts">
import { computed } from "vue"

const budget = 20000
const spent = 14088.15

const remaining = computed(() => {
  return budget - spent
})

const percent = computed(() => {
  if (budget === 0) return 0

  return Math.min((spent / budget) * 100, 100)
})

function formatMoney(value: number) {
  return new Intl.NumberFormat("th-TH", {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  }).format(value)
}
</script>

<template>
  <div class="card border border-base-300 bg-base-100 shadow-sm">
    <div class="card-body">

      <div class="grid gap-6 md:grid-cols-2">

        <!-- Progress -->
        <div class="flex flex-col items-center justify-center">

          <div class="relative flex size-48 items-center justify-center">

            <!-- background circle -->
            <div
              class="absolute inset-0 rounded-full border-[18px] border-base-200"
            />

            <!-- progress -->
            <div
              class="absolute inset-0 rounded-full border-[18px] border-success"
              :style="{
                clipPath: `inset(0 ${100 - percent}% 0 0)`
              }"
            />

            <div class="text-center">
              <p class="text-sm text-base-content/50">
                ใช้ไป
              </p>

              <p class="text-2xl font-bold">
                ฿{{ formatMoney(spent) }}
              </p>

              <p class="text-xs text-base-content/50">
                จาก ฿{{ formatMoney(budget) }}
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
              ฿{{ formatMoney(budget) }}
            </p>
          </div>

          <div class="rounded-xl bg-error/10 p-4">
            <p class="text-sm text-error">
              ใช้ไป
            </p>

            <p class="mt-2 text-xl font-bold text-error">
              ฿{{ formatMoney(spent) }}
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

<script setup lang="ts">
import { ref, watch } from "vue"
import type { Account } from "@/types/account"
import type { Category } from "@/types/category"
import type { BudgetForm } from "@/types/budget"
interface Props {
  open: boolean
  accounts: Account[]
  categories: Category[]
}
const props = defineProps<Props>()
const emit = defineEmits<{
  close: []
  save: [form: BudgetForm]
}>()
const dialogRef = ref<HTMLDialogElement | null>(null)
const form = ref<BudgetForm>({
  accountId: null,
  category: "",
  amount: null,
  startDate: "",
  endDate: "",
})
watch(
  () => props.open,
  (isOpen) => {
    if (isOpen) {
      dialogRef.value?.showModal()
    } else {
      dialogRef.value?.close()
    }
  },
)
function submitForm() {
  if (!form.value.category || !form.value.amount || !form.value.startDate || !form.value.endDate || !form.value.accountId) {
    return
  }
  emit("save", { ...form.value })
  resetForm()
}
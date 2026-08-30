<!-- <script setup lang="ts">
const categories = [
  {
    name: "อาหาร",
    spent: 6379,
    budget: 8000,
  },
  {
    name: "เดินทาง",
    spent: 7368,
    budget: 6000,
  },
  {
    name: "ช้อปปิ้ง",
    spent: 2341,
    budget: 4000,
  },
  {
    name: "บันเทิง",
    spent: 1200,
    budget: 2000,
  },
]

function getPercent(spent: number, budget: number) {
  if (budget === 0) return 0

  return Math.min((spent / budget) * 100, 100)
}

function formatMoney(value: number) {
  return new Intl.NumberFormat("th-TH").format(value)
}
</script>

<template>
  <div class="card border border-base-300 bg-base-100 shadow-sm">

    <div class="card-body">

      <div>
        <h2 class="card-title">
          งบประมาณแต่ละหมวด
        </h2>

        <p class="text-sm text-base-content/50">
          ติดตามการใช้จ่ายของแต่ละหมวด
        </p>
      </div>

      <div class="mt-4 divide-y divide-base-300">

        <div
          v-for="category in categories"
          :key="category.name"
          class="py-5"
        >

          <div class="flex items-center justify-between">

            <div>
              <p class="font-semibold">
                {{ category.name }}
              </p>

              <p class="text-sm text-base-content/50">
                ฿{{ formatMoney(category.spent) }}
                /
                ฿{{ formatMoney(category.budget) }}
              </p>
            </div>

            <p
              class="text-sm font-semibold"
              :class="
                category.spent > category.budget
                  ? 'text-error'
                  : 'text-base-content'
              "
            >
              {{ getPercent(category.spent, category.budget).toFixed(0) }}%
            </p>

          </div>

          <div class="mt-3 h-2 overflow-hidden rounded-full bg-base-200">

            <div
              class="h-full rounded-full transition-all"
              :class="
                category.spent > category.budget
                  ? 'bg-error'
                  : 'bg-success'
              "
              :style="{
                width:
                  getPercent(
                    category.spent,
                    category.budget
                  ) + '%'
              }"
            />

          </div>

        </div>

      </div>

    </div>
  </div>
</template> -->

<script setup lang="ts">
import { ref } from "vue"
import { Calendar, WalletCards, AlertCircle, PiggyBank } from "lucide-vue-next"
import { useCategoryStore } from "@/stores/category"
import { useAccountStore } from "@/stores/account"
import { resolveCategoryIcon } from "@/utils/categoryIcons"
import type { Budget } from "@/types/budget"
import { formatDate } from "@/utils/format"
import ActionButton from "../common/ActionButton.vue"
import ConfirmModal from "../common/ConfirmModal.vue"

interface Props {
  budgets: Budget[]
}

const targetBudget = ref<Budget | null>(null)

const props = defineProps<Props>()
const emit = defineEmits<{
  // add: []
  edit: [budget: Budget]
  delete: [id: number]
}>()

const categoryStore = useCategoryStore()
const accountStore = useAccountStore()

// const budgets = ref<Budget[]>([
//   {
//     id: 1,
//     accountId: 1,
//     category: "อาหาร",
//     amount: 8000,
//     spent: 6379,
//     startDate: "2026-08-01",
//     endDate: "2026-08-31",
//   },
//   {
//     id: 2,
//     accountId: 1,
//     category: "เดินทาง",
//     amount: 6000,
//     spent: 7368, // ตัวอย่างใช้เกินงบ
//     startDate: "2026-08-01",
//     endDate: "2026-08-31",
//   },
//   {
//     id: 3,
//     accountId: 1,
//     category: "ช้อปปิ้ง",
//     amount: 4000,
//     spent: 2341,
//     startDate: "2026-08-01",
//     endDate: "2026-08-31",
//   },
//   {
//     id: 4,
//     accountId: 1,
//     category: "บันเทิง",
//     amount: 2000,
//     spent: 1200,
//     startDate: "2026-08-15",
//     endDate: "2026-08-20", // ตัวอย่าง Custom Date Range
//   },
// ])

// คำนวณเปอร์เซ็นต์ (ไม่เกิน 100 สำหรับความกว้าง progress bar)
function getPercent(spent: number, amount: number) {
  if (!amount || amount === 0) return 0
  return Math.min((spent / amount) * 100, 100)
}

// เปอร์เซ็นต์จริง (อาจเกิน 100% ถ้าใช้เกินงบ)
function getRealPercent(spent: number, amount: number) {
  if (!amount || amount === 0) return 0
  return (spent / amount) * 100
}

function formatMoney(value: number) {
  return new Intl.NumberFormat("th-TH").format(value)
}

// หาชื่อบัญชีจาก accountId
function getAccountName(accountId: number | null) {
  if (!accountId) return "ทุกบัญชี"
  const acc = accountStore.accounts.find((a) => a.id === accountId)
  return acc?.name ?? "บัญชี"
}

// หา Icon หมวดหมู่
function getCategoryIcon(categoryName: string) {
  const cat = categoryStore.categories.find((c) => c.name === categoryName)
  return resolveCategoryIcon(cat?.icon)
}

// คำนวณจำนวนวันที่เหลือ
function getDaysRemaining(endDateStr: string) {
  const end = new Date(endDateStr)
  const today = new Date()
  today.setHours(0, 0, 0, 0)
  end.setHours(0, 0, 0, 0)
  const diffTime = end.getTime() - today.getTime()
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  return diffDays > 0 ? `เหลือ ${diffDays} วัน` : "สิ้นสุดแล้ว"
}

function openDeleteModal(budget: Budget) {
  targetBudget.value = budget
}
function handleConfirmDelete() {
  if (targetBudget.value) {
    emit("delete", targetBudget.value.id)
    targetBudget.value = null
  }
}
</script>

<template>
  <div class="card border border-base-300 bg-base-100 shadow-sm">
    <div class="card-body p-5 sm:p-6">

      <!-- Header -->
      <div class="flex items-center justify-between">
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

      <!-- Empty State -->
      <div v-if="budgets.length === 0"
        class="mt-6 flex flex-col items-center justify-center rounded-2xl border border-dashed border-base-300 py-12 text-center">
        <div class="flex size-14 items-center justify-center rounded-2xl bg-base-200 text-base-content/40">
          <PiggyBank class="size-7" />
        </div>
        <p class="mt-3 font-semibold text-base">
          ยังไม่มีงบประมาณในช่วงนี้
        </p>
        <p class="mt-1 text-sm text-base-content/50">
          กดปุ่ม "สร้าง Budget" ด้านบนเพื่อเริ่มวางแผนการเงิน
        </p>
      </div>

      <!-- Budget Cards Grid / List -->
      <div v-else class="mt-6 space-y-4">
        <div v-for="budget in budgets" :key="budget.id"
          class="group rounded-2xl border border-base-200 bg-base-100 p-4 transition-all hover:border-base-300 hover:shadow-md">
          <!-- Top Row: Icon + Category + Badges + Actions -->
          <div class="flex items-start justify-between gap-3">

            <!-- Category & Account Info -->
            <div class="flex items-center gap-3 min-w-0">
              <!-- Category Icon -->
              <div class="flex size-11 shrink-0 items-center justify-center rounded-xl bg-primary/10 text-primary">
                <component :is="getCategoryIcon(budget.category)" class="size-6" />
              </div>

              <div class="min-w-0">
                <div class="flex flex-wrap items-center gap-2">
                  <h3 class="font-bold text-base truncate">
                    {{ budget.category }}
                  </h3>
                  <!-- Status / Days badge -->
                  <span class="badge badge-sm badge-ghost text-xs">
                    {{ getDaysRemaining(budget.endDate) }}
                  </span>
                </div>

                <!-- Account & Date details -->
                <div class="mt-1 flex flex-wrap items-center gap-x-2 text-xs text-base-content/60">
                  <span class="flex items-center gap-1">
                    <WalletCards class="size-3.5" />
                    {{ getAccountName(budget.accountId) }}
                  </span>
                  <span>•</span>
                  <span class="flex items-center gap-1">
                    <Calendar class="size-3.5" />
                    {{ formatDate(new Date(budget.startDate)) }} ถึง {{ formatDate(new Date(budget.endDate)) }}
                  </span>
                </div>
              </div>
            </div>

            <!-- Action Buttons (Edit / Delete) -->
            <div class="flex items-center gap-1 shrink-0">
              <ActionButton edit-title="แก้ไขรายการ" delete-title="ลบรายการ" @edit="emit('edit', budget)"
                @delete="openDeleteModal(budget)" />
            </div>
          </div>

          <!-- Middle Row: Spent vs Amount Money Display -->
          <div class="mt-4 flex items-baseline justify-between text-sm">
            <div>
              <span class="text-xs text-base-content/50">ใช้ไปแล้ว: </span>
              <span class="font-bold text-base"
                :class="budget.spent > budget.amount ? 'text-error' : 'text-base-content'">
                ฿{{ formatMoney(budget.spent) }}
              </span>
              <span class="text-xs text-base-content/50"> / ฿{{ formatMoney(budget.amount) }}</span>
            </div>

            <!-- Percentage / Over-budget warning -->
            <div class="text-right">
              <span v-if="budget.spent > budget.amount"
                class="inline-flex items-center gap-1 text-xs font-bold text-error">
                <AlertCircle class="size-3.5" />
                เกินงบ ฿{{ formatMoney(budget.spent - budget.amount) }}
              </span>
              <span v-else class="text-xs font-semibold text-base-content/70">
                เหลือ ฿{{ formatMoney(budget.amount - budget.spent) }}
              </span>
              <span class="ml-2 font-bold" :class="budget.spent > budget.amount ? 'text-error' : 'text-base-content'">
                ({{ getRealPercent(budget.spent, budget.amount).toFixed(0) }}%)
              </span>
            </div>
          </div>

          <!-- Bottom Row: Progress Bar (เปลี่ยนสีตาม % ที่ใช้ไป) -->
          <div class="mt-2 h-2.5 w-full overflow-hidden rounded-full bg-base-200">
            <div class="h-full rounded-full transition-all duration-300" :class="{
              'bg-error': budget.spent > budget.amount,
              'bg-warning': budget.spent <= budget.amount && getRealPercent(budget.spent, budget.amount) >= 80,
              'bg-success': getRealPercent(budget.spent, budget.amount) < 80,
            }" :style="{ width: getPercent(budget.spent, budget.amount) + '%' }" />
          </div>
        </div>
      </div>
    </div>
  </div>
  <ConfirmModal :open="targetBudget !== null" title="ยืนยันการลบรายการ?" confirm-text="ลบรายการ" type="danger"
    @close="targetBudget = null" @confirm="handleConfirmDelete" />
</template>
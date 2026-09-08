<script setup lang="ts">
import { ref } from "vue"
import { AlertCircle, PiggyBank, Wallet,BarChart3 } from "lucide-vue-next"
import { resolveCategoryIcon } from "@/utils/categoryIcons"
import type { Budget } from "@/types/budget"
import { formatMoney } from "@/utils/format"
import ActionButton from "../common/ActionButton.vue"
import ConfirmModal from "../common/ConfirmModal.vue"
import type { Category } from "@/types/category.ts"
import type { Account } from "@/types/account.ts"
import { useRouter } from "vue-router"

interface Props {
  budgets: Budget[]
  categories: Category[]
  accounts: Account[]
}
const router = useRouter()
// กดแล้วส่งข้อมูลของ Budget ตัวนั้นไปหน้า Detail
function goToBudgetDetail(budget: Budget) {
  // router.push({
  //   path: "/budgets/detail", // หรือจะใช้แบบ params เช่น `/budgets/${budget.id}`
  //   query: {
  //     id: budget.id,
  //     category: budget.category,
  //     startDate: budget.startDate,
  //     endDate: budget.endDate,
  //     accountId: budget.accountId,
  //   },
  // })
  router.push(`/budgets/${budget.id}`)
}

const targetBudget = ref<Budget | null>(null)

const props = defineProps<Props>()
const emit = defineEmits<{
  edit: [budget: Budget]
  delete: [id: number]
}>()

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

// หาชื่อบัญชีจาก accountId
function getAccountName(accountId: number) {
  const acc = props.accounts.find((a) => a.id === accountId)
  return acc ? acc.name : "ไม่พบบัญชี"
}

// หา Icon หมวดหมู่
function getCategoryIcon(categoryName: string) {
  const cat = props.categories.find((c) => c.name === categoryName)
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
  <!-- <div class="card border border-base-300 bg-base-100 shadow-sm">
    <div class="card-body p-5 sm:p-6"> -->
  <div>
    <!-- Header -->
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

    <!-- Budget Cards Grid / List -->
    <div class="mt-2! space-y-4!">
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
                <span class="text-sm flex items-center gap-1">
                  <Wallet class="size-3.5" />
                  {{ getAccountName(budget.accountId) }}
                </span>
                <!-- <span>•</span>
                <span class="text-sm flex items-center gap-1">
                  <Calendar class="size-3.5" />
                  {{ formatDate(new Date(budget.startDate)) }} ถึง {{ formatDate(new Date(budget.endDate)) }}
                
                </span> -->
              </div>
            </div>
          </div>

          <!-- Action Buttons (Edit / Delete) -->
          <div class="flex items-center gap-1 shrink-0">
            <button type="button"
              class="btn btn-ghost btn-xs sm:btn-sm btn-square text-base-content/60 hover:text-primary hover:bg-primary/10"
              title="ดูรายละเอียดการใช้จ่าย" @click.stop="goToBudgetDetail(budget)">
              <BarChart3 class="size-4" /> <!-- หรือใช้ <Eye class="size-4" /> -->
            </button>
            <ActionButton edit-title="แก้ไขรายการ" delete-title="ลบรายการ" @edit="emit('edit', budget)"
              @delete="openDeleteModal(budget)" />
          </div>
        </div>

        <!-- Middle Row: Spent vs Amount Money Display -->
        <div class="mt-4 flex items-baseline justify-between text-sm">
          <div>
            <span class="text-sm text-base-content/50">ใช้ไปแล้ว: </span>
            <span class="font-bold text-base"
              :class="budget.spent > budget.amount ? 'text-error' : 'text-base-content'">
              ฿{{ formatMoney(budget.spent) }}
            </span>
            <span class="text-base text-base-content/50"> / ฿{{ formatMoney(budget.amount) }}</span>
          </div>

          <!-- Percentage / Over-budget warning -->
          <div class="text-right">
            <span v-if="budget.spent > budget.amount"
              class="inline-flex items-center gap-1 text-base font-bold text-error">
              <AlertCircle class="size-3.5" />
              เกินงบ ฿{{ formatMoney(budget.spent - budget.amount) }}
            </span>
            <span v-else class="text-base font-semibold text-base-content/70">
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
  <!-- </div>
  </div> -->
  <ConfirmModal :open="targetBudget !== null" title="ยืนยันการลบรายการ?" confirm-text="ลบรายการ" type="danger"
    @close="targetBudget = null" @confirm="handleConfirmDelete" />
</template>
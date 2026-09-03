<script setup lang="ts">
import BudgetCategoryList from "@/components/budget/BudgetCategoryList.vue";
import { Plus, WalletCards } from "lucide-vue-next"
import { computed, onMounted, ref, watch } from "vue";
import { useAccountStore } from "@/stores/account";
import { useCategoryStore } from "@/stores/category";
import { useBudgetForm } from "@/composables/budget/useBudgetForm";
import BudgetForm from "@/components/budget/BudgetForm.vue";
import BudgetOverview from "@/components/budget/BudgetOverview.vue";
import { useBudget } from "@/composables/budget/useBudget";
import AccountFilter from "@/components/filter/accountFilter.vue";
import type { Budget } from "@/types/budget";
import EmptyState from "@/components/common/EmptyState.vue";
const isDialogOpen = ref(false)
const accountStore = useAccountStore()
const categoryStore = useCategoryStore()
const selectedAccountId = ref<number | null>(null)
const status = ref<"active" | "ended">("active")
const selectedYear = ref<number | null>(null)
const selectedMonth = ref<number | null>(null)
const months = [
  { value: 1, label: "มกราคม" },
  { value: 2, label: "กุมภาพันธ์" },
  { value: 3, label: "มีนาคม" },
  { value: 4, label: "เมษายน" },
  { value: 5, label: "พฤษภาคม" },
  { value: 6, label: "มิถุนายน" },
  { value: 7, label: "กรกฎาคม" },
  { value: 8, label: "สิงหาคม" },
  { value: 9, label: "กันยายน" },
  { value: 10, label: "ตุลาคม" },
  { value: 11, label: "พฤศจิกายน" },
  { value: 12, label: "ธันวาคม" },
]

const currentYear = new Date().getFullYear()

const years = Array.from(
  { length: 5 },
  (_, index) => currentYear - index,
)
const {
  form,
  resetForm,
  setEditForm,
} = useBudgetForm()

const {
  budgets,
  isLoadingBudgets,
  editingBudgetId,
  loadBudgets,
  deleteBudget,
  saveBudget: saveBudgetApi,
  startEdit,
  cancelEdit,
} = useBudget()

function openEditBudget(budget: Budget) {
  console.log("openEditBudget", budget)
  startEdit(budget)
  setEditForm(budget)
  isDialogOpen.value = true
}

function openAddBudget() {

  cancelEdit()
  resetForm()

  // form.value.transactionDate =
  //   formatDate(selectedDate.value)

  isDialogOpen.value = true
}

async function saveBudget() {
  try {
    await saveBudgetApi(form.value)

    cancelEdit()

    isDialogOpen.value = false

    resetForm()
    await fetchBudgets()
  } catch (error) {
    console.error("Error saving budget:", error)
  }
}

// async function fetchBudgets() {
//   // const date = selectedMonth.value
//   const now = new Date()
//   await loadBudgets(
//     now.getFullYear(),
//     now.getMonth() + 1,
//     status.value,
//     selectedAccountId.value,
//   )
// }

async function fetchBudgets() {
  if (selectedAccountId.value === null) {
    return
  }

  await loadBudgets({
    status: status.value,
    accountId: selectedAccountId.value,

    year:
      status.value === "ended"
        ? selectedYear.value
        : null,

    month:
      status.value === "ended"
        ? selectedMonth.value
        : null,
  })
}

// async function selectStatus(
//   value: "active" | "ended"
// ) {
//   status.value = value
//   await fetchBudgets()
// }
async function handleYearChange() {
  // ถ้าเปลี่ยนปี ให้ reset เดือนก่อน
  selectedMonth.value = null

  await fetchBudgets()
}
async function handleMonthChange() {
  if (selectedYear.value === null) {
    return
  }

  await fetchBudgets()
}
async function selectStatus(value: "active" | "ended") {
  if (status.value === value) {
    return
  }

  status.value = value

  selectedYear.value = null
  selectedMonth.value = null

  await fetchBudgets()
}

async function handleDeleteBudget(id: number) {
  const now = new Date()
  try {
    await deleteBudget(id)
    await fetchBudgets()
    // await loadBudgets(
    //   now.getFullYear(),
    //   now.getMonth() + 1,
    //   "active",
    //   // status.value,
    //   // selectedAccountId.value
    // )
  } catch (error) {
    console.error(
      "Delete budget failed:",
      error,
    )
  }
}

const budgetPeriods = computed(() => {
  const groups = new Map<string, Budget[]>()

  for (const budget of budgets.value) {
    const key = `${budget.startDate}_${budget.endDate}`

    if (!groups.has(key)) {
      groups.set(key, [])
    }

    groups.get(key)!.push(budget)
  }

  return Array.from(groups.values()).flatMap((budgetGroup) => {
    const firstBudget = budgetGroup[0]

    if (!firstBudget) {
      return []
    }

    return [
      {
        startDate: firstBudget.startDate,
        endDate: firstBudget.endDate,
        budgets: budgetGroup,
      },
    ]
  })
})

function closeDialog() {
  isDialogOpen.value = false
}

// onMounted(() => {
//   const now = new Date()
//   loadBudgets(
//     now.getFullYear(),
//     now.getMonth() + 1,
//     "active",
//     // status.value, 
//     // selectedAccountId.value
//   )
//   categoryStore.loadCategories()
//   accountStore.loadAccounts()
// })

watch(selectedAccountId, async (newAccountId, oldAccountId) => {
  if (newAccountId === null) {
    return
  }

  if (newAccountId === oldAccountId) {
    return
  }

  await fetchBudgets()
})

onMounted(async () => {
  await Promise.all([
    accountStore.loadAccounts(),
    categoryStore.loadCategories(),
  ])

  const firstAccount = accountStore.accounts[0]

  if (!firstAccount) {
    return
  }

  selectedAccountId.value = firstAccount.id

  // await fetchBudgets()
})

</script>

<template>
  <section class="space-y-6!">

    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold">
          Budget
        </h1>

        <p class="text-base-content/60">
          วางแผนและติดตามค่าใช้จ่ายของคุณ
        </p>
      </div>

      <button class="btn text-white bg-green-700" type="button" @click="openAddBudget">
        <Plus class="size-4" />
        สร้าง Budget
      </button>
    </div>

    <!-- 2. Controls & Filters -->
    <div class="flex flex-wrap items-center justify-between gap-3 mt-4">
      <!-- Status Toggle (สลับ Active / Ended) -->
      <div class="join bg-base-200 p-1 rounded-xl">
        <button class="btn btn-sm join-item border-none"
          :class="status === 'active' ? 'btn-success text-white shadow-sm' : 'btn-ghost'"
          @click="selectStatus('active')">
          🟢 กำลังใช้งาน
        </button>
        <button class="btn btn-sm join-item border-none"
          :class="status === 'ended' ? 'btn-neutral text-white shadow-sm' : 'btn-ghost'" @click="selectStatus('ended')">
          📁 สิ้นสุดแล้ว
        </button>
      </div>
      <!-- Account Filter -->
      <!-- <AccountFilter v-model="selectedAccountId" :accounts="accountStore.accounts" /> -->
      <!-- Filters -->
      <div class="flex items-center gap-2">

        <!-- Ended Date Filters -->
        <template v-if="status === 'ended'">
          <select v-model="selectedYear" class="select select-bordered select-sm" @change="handleYearChange">
            <option :value="null">
              ทุกปี
            </option>

            <option v-for="year in years" :key="year" :value="year">
              {{ year + 543 }}
            </option>
          </select>

          <select v-model="selectedMonth" class="select select-bordered select-sm" :disabled="selectedYear === null"
            @change="handleMonthChange">
            <option :value="null">
              ทุกเดือน
            </option>

            <option v-for="month in months" :key="month.value" :value="month.value">
              {{ month.label }}
            </option>
          </select>
        </template>

        <!-- Account -->
        <AccountFilter v-model="selectedAccountId" :accounts="accountStore.accounts" />

      </div>
    </div>
    <!-- Overview -->
    <!-- <BudgetOverview :budgets="budgets" /> -->
    <EmptyState v-if="budgetPeriods.length === 0" :icon="WalletCards" title="ยังไม่มีงบประมาณ"
      description='กดปุ่ม "สร้าง Budget" ด้านบนเพื่อเริ่มวางแผนการเงิน' />
    <div v-else class="space-y-4!">
      <BudgetOverview v-for="period in budgetPeriods" :key="`${period.startDate}-${period.endDate}`"
        :budgets="period.budgets" :start-date="period.startDate" :end-date="period.endDate"
        :categories="categoryStore.categories" :accounts="accountStore.accounts" @edit="openEditBudget"
        @delete="handleDeleteBudget" />
    </div>
    <!-- Categories -->
    <!-- <BudgetCategoryList :budgets="budgets" :categories="categoryStore.categories" :accounts="accountStore.accounts" @delete="handleDeleteBudget" @edit="openEditBudget" /> -->

  </section>
  <BudgetForm :open="isDialogOpen" :form="form" :accounts="accountStore.accounts" :categories="categoryStore.categories"
    @close="closeDialog" @save="saveBudget" />
</template>

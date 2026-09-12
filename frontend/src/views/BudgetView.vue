<script setup lang="ts">
import { Plus, WalletCards } from "lucide-vue-next"
import { computed, onMounted, ref } from "vue";
import { useAccountStore } from "@/stores/account";
import { useCategoryStore } from "@/stores/category";
import { useBudgetForm } from "@/composables/budget/useBudgetForm";
import BudgetForm from "@/components/budget/BudgetForm.vue";
import BudgetOverview from "@/components/budget/BudgetOverview.vue";
import { useBudget } from "@/composables/budget/useBudget";
import AccountFilter from "@/components/filter/accountFilter.vue";
import type { Budget } from "@/types/budget";
import EmptyState from "@/components/common/EmptyState.vue";
import { useBudgetFilter } from "@/composables/budget/useBudgetFilter";
import { groupBudgetsByPeriod } from "@/utils/budget";

const isDialogOpen = ref(false)
const accountStore = useAccountStore()
const categoryStore = useCategoryStore()

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

const {
  status,
  selectedAccountId,
  selectedYear,
  selectedMonth,
  months,
  years,
  fetchBudgets,
  selectStatus,
  handleYearChange,
  handleMonthChange,
} = useBudgetFilter(loadBudgets)

function openEditBudget(budget: Budget) {
  startEdit(budget)
  setEditForm(budget)
  isDialogOpen.value = true
}

function openAddBudget() {
  cancelEdit()
  resetForm()
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

async function handleDeleteBudget(id: number) {
  try {
    await deleteBudget(id)
    await fetchBudgets()

  } catch (error) {
    console.error(
      "Delete budget failed:",
      error,
    )
  }
}
const budgetPeriods = computed(() => groupBudgetsByPeriod(budgets.value))

function closeDialog() {
  isDialogOpen.value = false
}

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
    <div v-if="isLoadingBudgets" class="flex justify-center py-12">
      <span class="loading loading-spinner loading-lg" />
    </div>
    <EmptyState v-else-if="budgetPeriods.length === 0" :icon="WalletCards" title="ยังไม่มีงบประมาณ"
      description='กดปุ่ม "สร้าง Budget" ด้านบนเพื่อเริ่มวางแผนการเงิน' />
    <div v-else class="space-y-4!">
      <BudgetOverview v-for="period in budgetPeriods" :key="`${period.startDate}-${period.endDate}`"
        :budgets="period.budgets" :start-date="period.startDate" :end-date="period.endDate"
        :categories="categoryStore.categories" :accounts="accountStore.accounts" @edit="openEditBudget"
        @delete="handleDeleteBudget" />
    </div>

  </section>
  <BudgetForm :open="isDialogOpen" :form="form" :accounts="accountStore.accounts" :categories="categoryStore.categories"
    @close="closeDialog" @save="saveBudget" />
</template>

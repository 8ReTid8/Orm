<script setup lang="ts">
import BudgetCategoryList from "@/components/budget/BudgetCategoryList.vue";
import { ChevronLeft, ChevronRight, Plus } from "lucide-vue-next"
import { onMounted, ref } from "vue";
import { useAccountStore } from "@/stores/account";
import { useCategoryStore } from "@/stores/category";
import { useBudgetForm } from "@/composables/budget/useBudgetForm";
import BudgetForm from "@/components/budget/BudgetForm.vue";
import BudgetFilter from "@/components/filter/budgetFilter.vue";
import BudgetOverview from "@/components/budget/BudgetOverview.vue";
import { useBudget } from "@/composables/budget/useBudget";
import AccountFilter from "@/components/filter/accountFilter.vue";
import type { Budget } from "@/types/budget";
const isDialogOpen = ref(false)
const accountStore = useAccountStore()
const categoryStore = useCategoryStore()
const startDate = ref("")
const endDate = ref("")
const selectedAccountId = ref<number | null>(null)
const status = ref("")
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
  resetForm(  )

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
  } catch (error) {
    console.error("Error saving budget:", error)
  }
}

async function handleDeleteBudget(id: number) {
  const now = new Date()
  try {
    await deleteBudget(id)

    await loadBudgets(
      now.getFullYear(),
      now.getMonth() + 1,
      "active",
      // status.value,
      // selectedAccountId.value
    )
  } catch (error) {
    console.error(
      "Delete budget failed:",
      error,
    )
  }
}

function closeDialog() {
  isDialogOpen.value = false
}

onMounted(() => {
  const now = new Date()
  loadBudgets(
    now.getFullYear(),
    now.getMonth() + 1,
    "active",
    // status.value, 
    // selectedAccountId.value
  )
  categoryStore.loadCategories()
  accountStore.loadAccounts()
})

</script>

<template>
  <section class="!space-y-6">

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

    <!-- <BudgetFilter v-model:start-date="startDate" v-model:end-date="endDate"
      v-model:selected-account-id="selectedAccountId" v-model:status="status" :accounts="accountStore.accounts" /> -->
    <!-- 1. Period Navigator (เลือกเดือน) -->
    <div class="flex items-center justify-center">
      <button class="btn btn-ghost btn-sm btn-circle" type="button" @click="prevMonth">
        <ChevronLeft class="size-4" />
      </button>
      <div class="mx-6 min-w-36 text-center">
        <p class="text-lg font-bold">
          {{ selectedMonthText }} <!-- เช่น "สิงหาคม 2569" -->
        </p>
      </div>
      <button class="btn btn-ghost btn-sm btn-circle" type="button" @click="nextMonth">
        <ChevronRight class="size-4" />
      </button>
    </div>
    <!-- 2. Controls & Filters -->
    <div class="flex flex-wrap items-center justify-between gap-3 mt-4">
      <!-- Status Toggle (สลับ Active / Ended) -->
      <div class="join bg-base-200 p-1 rounded-xl">
        <button class="btn btn-sm join-item border-none"
          :class="status === 'active' ? 'btn-success text-white shadow-sm' : 'btn-ghost'" @click="status = 'active'">
          🟢 กำลังใช้งาน
        </button>
        <button class="btn btn-sm join-item border-none"
          :class="status === 'ended' ? 'btn-neutral text-white shadow-sm' : 'btn-ghost'" @click="status = 'ended'">
          📁 สิ้นสุดแล้ว
        </button>
      </div>
      <!-- Account Filter -->
      <!-- <select v-model="selectedAccountId" class="select select-bordered select-sm">
        <option :value="null">ทุกบัญชี</option>
        <option v-for="acc in accountStore.accounts" :key="acc.id" :value="acc.id">
          {{ acc.name }}
        </option>
      </select> -->
      <AccountFilter v-model="selectedAccountId" :accounts="accountStore.accounts" />
    </div>
    <!-- Overview -->
    <BudgetOverview />

    <!-- Categories -->
    <BudgetCategoryList :budgets="budgets" @delete="handleDeleteBudget" @edit="openEditBudget" />

  </section>
  <BudgetForm :open="isDialogOpen" :form="form" :accounts="accountStore.accounts" :categories="categoryStore.categories"
    @close="closeDialog" @save="saveBudget" />
</template>

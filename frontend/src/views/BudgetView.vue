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

function closeDialog() {
  isDialogOpen.value = false
}
function openAddBudget() {

  // cancelEdit()
  // resetForm(
  //   formatDate(selectedDate.value),
  // )

  // form.value.transactionDate =
  //   formatDate(selectedDate.value)

  isDialogOpen.value = true
}
async function saveBudget() {
  // await saveTransactionApi()
  isDialogOpen.value = false
}

onMounted(() => {
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

    <!-- Month selector -->
    <!-- <div class="flex items-center justify-center gap-4">
      <button class="btn btn-ghost btn-sm btn-circle">
        <ChevronLeft class="size-4" />
      </button>

      <div class="text-center">
        <p class="font-semibold">
          สิงหาคม 2569
        </p>

        <p class="text-xs text-base-content/50">
          1 ส.ค. - 31 ส.ค.
        </p>
      </div>

      <button class="btn btn-ghost btn-sm btn-circle">
        <ChevronRight class="size-4" />
      </button>
    </div> -->
    <!-- Period Navigation -->
    <div class="flex items-center justify-center">
      <button class="btn btn-ghost btn-sm btn-circle" type="button">
        <ChevronLeft class="size-4" />
      </button>

      <div class="mx-6 min-w-36 text-center">
        <p class="font-semibold">
          สิงหาคม 2569
        </p>

        <p class="mt-1 text-xs text-base-content/50">
          Budgets ในเดือนนี้
        </p>
      </div>

      <button class="btn btn-ghost btn-sm btn-circle" type="button">
        <ChevronRight class="size-4" />
      </button>
    </div>

    <!-- Filters -->
    <div class="flex flex-wrap items-center justify-between gap-3">

      <div>
        <h2 class="text-lg font-semibold">
          Budgets
        </h2>

        <p class="text-sm text-base-content/50">
          ดูงบประมาณที่อยู่ในช่วงเวลานี้
        </p>
      </div>
      <div class="flex flex-wrap gap-2">

        <!-- Account -->
        <select class="select select-bordered select-sm">
          <option value="">
            ทุกบัญชี
          </option>

          <option>
            เงินสด
          </option>

          <option>
            KBank
          </option>
        </select>

        <!-- Status -->
        <select class="select select-bordered select-sm">
          <option value="">
            ทุกสถานะ
          </option>

          <option value="active">
            Active
          </option>

          <option value="upcoming">
            Upcoming
          </option>

          <option value="completed">
            Completed
          </option>
        </select>

      </div>
    </div>
    <BudgetFilter v-model:start-date="startDate" v-model:end-date="endDate"
      v-model:selected-account-id="selectedAccountId" v-model:status="status" :accounts="accountStore.accounts" />
    <!-- Overview -->
    <BudgetOverview />

    <!-- Categories -->
    <BudgetCategoryList />

  </section>
  <BudgetForm :open="isDialogOpen" :form="form" :accounts="accountStore.accounts"
    :categories="categoryStore.categories" :selected-date-text="selectedDateText" @close="closeDialog"
    @save="saveBudget" />
</template>

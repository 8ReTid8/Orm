<script setup lang="ts">
import BudgetCategoryList from "@/components/budget/BudgetCategoryList.vue";
import BudgetOverview from "@/components/budget/BudgetOverview.vue";
import { ChevronLeft, ChevronRight, Plus } from "lucide-vue-next"
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

      <button class="btn text-white bg-green-700" type="button">
        <Plus class="size-4" />
        สร้าง Budget
      </button>
    </div>

    <!-- Month selector -->
    <div class="flex items-center justify-center gap-4">
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
    </div>

    <!-- Overview -->
    <BudgetOverview />

    <!-- Categories -->
    <BudgetCategoryList />

  </section>
</template>
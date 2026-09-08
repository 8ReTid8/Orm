<script setup lang="ts">
import { useBudget } from "@/composables/budget/useBudget";
import { ref, onMounted } from "vue"
const {
    budgets,
    isLoadingBudgets,
    editingBudgetId,
    budgetDetail,
    loadBudgetDetail,
    loadBudgets,
    deleteBudget,
    saveBudget: saveBudgetApi,
    startEdit,
    cancelEdit,
} = useBudget()
onMounted(async () => {
    await loadBudgetDetail(7) // ตัวอย่างการโหลด Budget Detail ของ Budget ที่มี id = 1
    console.log("Budget Detail:", budgetDetail.value)
})

</script>
<template>
    <div class="space-y-6 p-6">

        <!-- Header -->
        <div>
            <button class="btn btn-ghost btn-sm mb-3">
                ← Back
            </button>

            <div class="flex items-center justify-between">
                <div>
                    <h1 class="text-2xl font-bold">
                        Food Budget
                    </h1>

                    <p class="text-sm text-base-content/60">
                        SCB • 1 Sep 2026 - 30 Sep 2026
                    </p>
                </div>

                <div class="text-right">
                    <p class="text-sm text-base-content/60">
                        Budget
                    </p>

                    <p class="text-2xl font-bold">
                        ฿10,000
                    </p>
                </div>
            </div>
        </div>


        <!-- Budget Usage -->
        <div class="grid grid-cols-1 gap-4 md:grid-cols-3">

            <!-- Spent -->
            <div class="card bg-base-100 shadow-sm">
                <div class="card-body">

                    <p class="text-sm text-base-content/60">
                        Spent
                    </p>

                    <p class="text-2xl font-bold">
                        ฿6,450
                    </p>

                    <p class="text-sm">
                        64.5% of budget
                    </p>

                </div>
            </div>


            <!-- Remaining -->
            <div class="card bg-base-100 shadow-sm">
                <div class="card-body">

                    <p class="text-sm text-base-content/60">
                        Remaining
                    </p>

                    <p class="text-2xl font-bold">
                        ฿3,550
                    </p>

                    <p class="text-sm">
                        35.5% remaining
                    </p>

                </div>
            </div>


            <!-- Average -->
            <div class="card bg-base-100 shadow-sm">
                <div class="card-body">

                    <p class="text-sm text-base-content/60">
                        Average / Day
                    </p>

                    <p class="text-2xl font-bold">
                        ฿215
                    </p>

                    <p class="text-sm text-base-content/60">
                        Based on 30 days
                    </p>

                </div>
            </div>

        </div>


        <!-- Spending Graph -->
        <div class="card bg-base-100 shadow-sm">

            <div class="card-body">

                <div>
                    <h2 class="text-lg font-bold">
                        Spending
                    </h2>

                    <p class="text-sm text-base-content/60">
                        การใช้เงินในช่วง Budget
                    </p>
                </div>


                <!-- Chart -->
                <div class="h-80 w-full">

                    <canvas id="budget-spending-chart"></canvas>

                </div>

            </div>

        </div>


        <!-- Transactions -->
        <div class="card bg-base-100 shadow-sm">

            <div class="card-body">

                <div class="flex items-center justify-between">

                    <div>
                        <h2 class="text-lg font-bold">
                            Transactions
                        </h2>

                        <p class="text-sm text-base-content/60">
                            รายการใช้เงินของ Budget นี้
                        </p>
                    </div>

                    <span class="badge badge-neutral">
                        8 transactions
                    </span>

                </div>


                <!-- Transaction List -->
                <div class="mt-2 divide-y divide-base-200">


                    <!-- Transaction -->
                    <div class="flex items-center justify-between py-4">

                        <div class="flex items-center gap-3">

                            <div class="flex size-10 items-center justify-center rounded-full bg-base-200">
                                🍜
                            </div>

                            <div>

                                <p class="font-medium">
                                    Lunch
                                </p>

                                <p class="text-sm text-base-content/60">
                                    8 Sep 2026 • Food
                                </p>

                            </div>

                        </div>

                        <p class="font-semibold">
                            -฿250
                        </p>

                    </div>


                    <!-- Transaction -->
                    <div class="flex items-center justify-between py-4">

                        <div class="flex items-center gap-3">

                            <div class="flex size-10 items-center justify-center rounded-full bg-base-200">
                                ☕
                            </div>

                            <div>

                                <p class="font-medium">
                                    Coffee
                                </p>

                                <p class="text-sm text-base-content/60">
                                    8 Sep 2026 • Food
                                </p>

                            </div>

                        </div>

                        <p class="font-semibold">
                            -฿80
                        </p>

                    </div>


                    <!-- Transaction -->
                    <div class="flex items-center justify-between py-4">

                        <div class="flex items-center gap-3">

                            <div class="flex size-10 items-center justify-center rounded-full bg-base-200">
                                🍔
                            </div>

                            <div>

                                <p class="font-medium">
                                    Dinner
                                </p>

                                <p class="text-sm text-base-content/60">
                                    7 Sep 2026 • Food
                                </p>

                            </div>

                        </div>

                        <p class="font-semibold">
                            -฿350
                        </p>

                    </div>


                    <!-- Transaction -->
                    <div class="flex items-center justify-between py-4">

                        <div class="flex items-center gap-3">

                            <div class="flex size-10 items-center justify-center rounded-full bg-base-200">
                                🛒
                            </div>

                            <div>

                                <p class="font-medium">
                                    Groceries
                                </p>

                                <p class="text-sm text-base-content/60">
                                    5 Sep 2026 • Food
                                </p>

                            </div>

                        </div>

                        <p class="font-semibold">
                            -฿1,200
                        </p>

                    </div>


                    <!-- Transaction -->
                    <div class="flex items-center justify-between py-4">

                        <div class="flex items-center gap-3">

                            <div class="flex size-10 items-center justify-center rounded-full bg-base-200">
                                🍱
                            </div>

                            <div>

                                <p class="font-medium">
                                    Dinner
                                </p>

                                <p class="text-sm text-base-content/60">
                                    4 Sep 2026 • Food
                                </p>

                            </div>

                        </div>

                        <p class="font-semibold">
                            -฿420
                        </p>

                    </div>


                </div>

            </div>

        </div>

    </div>
</template>
<script setup lang="ts">
import type { Transaction } from '@/types/transaction';
import { CalendarDays } from 'lucide-vue-next';
interface Props {
    selectedDate: Date
}

interface Props {
    selectedDate: Date
    transactions: Transaction[]
}

const props = defineProps<Props>()
// defineProps<Props>()

const emit = defineEmits<{
    select: [date: Date]
}>()
function formatDate(date: Date): string {
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, "0")
    const day = String(date.getDate()).padStart(2, "0")

    return `${year}-${month}-${day}`
}

function getDaySummary(date: Date) {
    const targetDate = formatDate(date)

    const dayTransactions = props.transactions.filter(
        (transaction) =>
            transaction.transactionDate.startsWith(targetDate)
    )

    const income = dayTransactions
        .filter((transaction) => transaction.type === "income")
        .reduce((sum, transaction) => sum + transaction.amount, 0)

    const expense = dayTransactions
        .filter((transaction) => transaction.type === "expense")
        .reduce((sum, transaction) => sum + transaction.amount, 0)

    const total = income + expense

    return {
        income,
        expense,

        incomePercent:
            total === 0 ? 0 : (income / total) * 100,

        expensePercent:
            total === 0 ? 0 : (expense / total) * 100,
    }
}
</script>

<template>
    <div class="card border border-base-300 bg-base-100 shadow-sm">
        <div class="card-body">
            <div class="mb-3 flex items-center gap-2">
                <CalendarDays class="size-5" />
                <h2 class="card-title">ปฏิทินรายรับรายจ่าย</h2>
            </div>
            <VCalendar class="finance-calendar w-full" expanded borderless>
                <!-- <template #day-content="{ day }">
                    <button type="button"
                        class="relative block min-h-28 w-full cursor-pointer bg-base-100 p-3 pt-10 text-left transition-colors hover:bg-base-200"
                        @click="emit('select', day.date)">
                        <span
                            class="absolute right-3 top-2 flex size-7 items-center justify-center rounded-full text-sm font-medium">
                            {{ day.day }}
                        </span>
                    </button>
                </template> -->
                <template #day-content="{ day }">
                    <button type="button"
                        class="relative block min-h-28 w-full cursor-pointer bg-base-100 p-3 pt-10 text-left transition-colors hover:bg-base-200"
                        @click="emit('select', day.date)">
                        <!-- เลขวัน -->
                        <span
                            class="absolute right-3 top-2 flex size-7 items-center justify-center rounded-full text-sm font-medium">
                            {{ day.day }}
                        </span>

                        <!-- Income / Expense -->
                        <template v-if="
                            getDaySummary(day.date).income > 0 ||
                            getDaySummary(day.date).expense > 0
                        ">
                            <div class="absolute bottom-3 left-3 right-3">
                                <!-- Label -->
                                <div class="mb-1 flex justify-between text-xs">
                                    <span class="font-medium text-success">
                                        รับ
                                    </span>

                                    <span class="font-medium text-error">
                                        จ่าย
                                    </span>
                                </div>
                                <!-- Bar -->
                                <div class="flex h-2 w-full overflow-hidden rounded-full bg-base-200">

                                    <div class="bg-success" :style="{
                                        width:
                                            getDaySummary(day.date).incomePercent + '%'
                                    }" />

                                    <div class="bg-error" :style="{
                                        width:
                                            getDaySummary(day.date).expensePercent + '%'
                                    }" />

                                </div>

                                <!-- Text -->
                                <div class="mt-1 flex justify-between text-[11px] text-base-content/50">
                                    <span>
                                        +{{ getDaySummary(day.date).income.toLocaleString() }}
                                    </span>

                                    <span>
                                        -{{ getDaySummary(day.date).expense.toLocaleString() }}
                                    </span>
                                </div>

                            </div>
                        </template>
                    </button>
                </template>
            </VCalendar>
        </div>
    </div>
</template>

<style>
.finance-calendar .vc-weeks {
    border-top: 1px solid var(--color-base-300);
}

.finance-calendar .vc-day {
    min-height: 112px;
    padding: 0 !important;
    border-right: 1px solid var(--color-base-300);
    border-bottom: 1px solid var(--color-base-300);
    box-sizing: border-box;
}
</style>
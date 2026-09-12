<script setup lang="ts">
import type { Budget } from '@/types/budget';
import { formatMoney } from '@/utils/format';
import { computed } from 'vue';
interface Props {
    budget: Budget
}
const props = defineProps<Props>()
// 1. คำนวณยอดคงเหลือ
const remaining = computed(() => {
    return props.budget.amount - props.budget.spent
})
// 2. คำนวณเปอร์เซ็นต์จริง
const realPercent = computed(() => {
    if (props.budget.amount === 0) return 0
    return (props.budget.spent / props.budget.amount) * 100
})
// 3. คำนวณจำนวนวันทั้งหมดของงบนี้
const totalDays = computed(() => {
    const start = new Date(props.budget.startDate).getTime()
    const end = new Date(props.budget.endDate).getTime()
    const diff = Math.ceil((end - start) / (1000 * 60 * 60 * 24)) + 1
    return diff > 0 ? diff : 1
})
// 4. คำนวณค่าใช้จ่ายเฉลี่ยต่อวัน
const averagePerDay = computed(() => {
    return props.budget.spent / totalDays.value
})
</script>
<template>
    <div class="grid grid-cols-1 gap-3 sm:grid-cols-3">
        <!-- Spent -->
        <div class="rounded-xl bg-error/10 p-4">
            <p class="text-lg font-medium text-error">ใช้ไปแล้ว</p>
            <p class="mt-2 text-2xl font-bold text-error">
                ฿{{ formatMoney(budget.spent) }}
            </p>
            <p class="mt-1 text-sm text-base-content/60">
                {{ realPercent.toFixed(1) }}% ของงบประมาณ
            </p>
        </div>
        <!-- Remaining -->
        <div class="rounded-xl p-4" :class="remaining >= 0 ? 'bg-success/10' : 'bg-error/10'">
            <p class="text-lg font-medium" :class="remaining >= 0 ? 'text-success' : 'text-error'">
                {{ remaining >= 0 ? "คงเหลือ" : "ใช้เกินงบ" }}
            </p>
            <p class="mt-2 text-2xl font-bold" :class="remaining >= 0 ? 'text-success' : 'text-error'">
                ฿{{ formatMoney(Math.abs(remaining)) }}
            </p>
            <p class="mt-1 text-sm text-base-content/60">
                {{ remaining >= 0 ? `เหลืออีก ${(100 - realPercent).toFixed(1)}%` : "เกินเป้าหมายที่ตั้งไว้" }}
            </p>
        </div>
        <!-- Average per day -->
        <div class="rounded-xl bg-base-200 p-4">
            <p class="text-lg font-medium text-base-content/60">เฉลี่ย / วัน</p>
            <p class="mt-2 text-2xl font-bold">
                ฿{{ formatMoney(averagePerDay) }}
            </p>
            <p class="mt-1 text-sm text-base-content/60">
                จากระยะเวลาทั้งหมด {{ totalDays }} วัน
            </p>
        </div>
    </div>
</template>
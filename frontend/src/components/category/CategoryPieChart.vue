<script setup lang="ts">
import { computed } from "vue"
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js"
import { Doughnut } from "vue-chartjs"
import { formatMoney } from "@/utils/format"

ChartJS.register(ArcElement, Tooltip, Legend)

interface CategoryItem {
  name: string
  amount: number
  color?: string
}

interface Props {
  items: CategoryItem[]
  emptyText?: string // 👈 เพิ่ม prop นี้
}

const props = withDefaults(defineProps<Props>(), {
  emptyText: "ไม่มีข้อมูลในช่วงเวลานี้", // 👈 กำหนดค่าเริ่มต้น
})
// const props = defineProps<Props>()

// จานสีสวยงามสำหรับหมวดหมู่ต่างๆ
const palette = [
  "#22c55e", "#3b82f6", "#f59e0b", "#ef4444", "#8b5cf6",
  "#ec4899", "#14b8a6", "#f97316", "#06b6d4", "#84cc16"
]

const chartData = computed(() => ({
  labels: props.items.map(i => i.name),
  datasets: [{
    data: props.items.map(i => i.amount),
    backgroundColor: props.items.map((_, index) => palette[index % palette.length]),
    borderWidth: 2,
    borderColor: "#ffffff",
  }],
}))

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: { position: "right" as const },
    tooltip: {
      callbacks: {
        label: (ctx: any) => ` ฿${formatMoney(ctx.raw)}`,
      },
    },
  },
}
</script>

<template>
  <div class="h-64 w-full">
    <Doughnut v-if="items.length > 0" :data="chartData" :options="chartOptions" />
    <div v-else class="flex h-full items-center justify-center text-base-content/50">
      {{ emptyText }}
    </div>
  </div>
</template>
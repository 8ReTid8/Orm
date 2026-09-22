<script setup lang="ts">
import { computed } from "vue"
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
  type ChartOptions,
} from "chart.js"
import { Bar } from "vue-chartjs"
import { formatMoney } from "@/utils/format"
import { THAI_MONTHS } from "@/utils/date"
import type { ComparisonPeriod, ComparisonTotal } from "@/types/summary"

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend)

interface Props {
  items: ComparisonTotal[]
  period: ComparisonPeriod
}

const props = defineProps<Props>()

const chartData = computed(() => ({
  labels: props.period === "month"
    ? THAI_MONTHS.map((m) => m.shortLabel)
    : props.items.map((item) => String(item.period)),
  datasets: [
    {
      label: "รายรับ",
      backgroundColor: "#22c55e",
      borderRadius: 6,
      data: props.items.map((item) => item.income),
    },
    {
      label: "รายจ่าย",
      backgroundColor: "#ef4444",
      borderRadius: 6,
      data: props.items.map((item) => item.expense),
    },
  ],
}))

const chartOptions: ChartOptions<"bar"> = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: { position: "top" },
    tooltip: {
      callbacks: {
        label: (ctx) => ` ${ctx.dataset.label}: ฿${formatMoney(Number(ctx.raw) || 0)}`,
      },
    },
  },
  scales: {
    y: {
      beginAtZero: true,
      ticks: {
        callback: (val) => `฿${Number(val).toLocaleString()}`,
      },
    },
  },
}
</script>

<template>
  <div class="h-72 w-full">
    <Bar :data="chartData" :options="chartOptions" />
  </div>
</template>

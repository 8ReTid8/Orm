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
import type { Transaction } from "@/types/transaction"

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend)

interface Props {
  transactions: Transaction[]
}

const props = defineProps<Props>()

// รวมยอดรายรับ และรายจ่าย แยกตาม 12 เดือน
const monthlyData = computed(() => {
  const income = new Array(12).fill(0)
  const expense = new Array(12).fill(0)

  for (const t of props.transactions) {
    const monthIndex = new Date(t.transactionDate).getMonth() // 0 - 11
    if (monthIndex >= 0 && monthIndex < 12) {
      if (t.type === "income") {
        income[monthIndex] += t.amount
      } else if (t.type === "expense") {
        expense[monthIndex] += t.amount
      }
    }
  }

  return { income, expense }
})

const chartData = computed(() => ({
  labels: THAI_MONTHS.map((m) => m.shortLabel),
  datasets: [
    {
      label: "รายรับ",
      backgroundColor: "#22c55e",
      borderRadius: 6,
      data: monthlyData.value.income,
    },
    {
      label: "รายจ่าย",
      backgroundColor: "#ef4444",
      borderRadius: 6,
      data: monthlyData.value.expense,
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
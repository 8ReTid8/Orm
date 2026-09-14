<script setup lang="ts">
import { computed } from "vue"
import type { Budget } from "@/types/budget"
import type { Transaction } from "@/types/transaction"
import { formatMoney } from "@/utils/format"

// นำเข้า Chart.js และ Component Line จาก vue-chartjs
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler,
  type ChartOptions,
  type ChartData,
} from "chart.js"
import { Line } from "vue-chartjs"

// ลงทะเบียนโมดูลที่ต้องใช้ (Tree-shaking ช่วยให้ไฟล์ไม่หนัก)
ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
)

interface Props {
  budget: Budget
  transactions: Transaction[]
}

const props = defineProps<Props>()

// 1. เตรียมข้อมูล Labels (วันที่) และ Datasets สำหรับกราฟ
const chartData = computed<ChartData<"line">>(() => {
  const start = new Date(props.budget.startDate)
  const end = new Date(props.budget.endDate)
  const today = new Date()

  start.setHours(0, 0, 0, 0)
  end.setHours(0, 0, 0, 0)
  today.setHours(0, 0, 0, 0)

  const plotEnd = end < today ? end : today

  // รวบรวมยอดใช้จ่ายรายวัน
  const dailyMap = new Map<string, number>()
  for (const t of props.transactions) {
    if (t.type === "expense") {
      const d = t.transactionDate.slice(0, 10)
      dailyMap.set(d, (dailyMap.get(d) ?? 0) + t.amount)
    }
  }

  const labels: string[] = []
  const spentData: number[] = []
  const budgetLimitData: number[] = []

  let runningTotal = 0
  const cur = new Date(start)

  // วนลูปตั้งแต่วันแรกจนถึงวันสิ้นสุดงบ
  while (cur <= end) {
    const dateStr = cur.toISOString().slice(0, 10)
    // Label แสดงแบบสั้น เช่น "1 ก.ย."
    labels.push(
      cur.toLocaleDateString("th-TH", { day: "numeric", month: "short" })
    )

    // เพดานงบประมาณ (มีค่าเท่ากันทุกวันเป็นเส้นตรง)
    budgetLimitData.push(props.budget.amount)

    // ยอดใช้สะสม (วาดเฉพาะวันที่มีข้อมูลถึงปัจจุบัน)
    if (cur <= plotEnd) {
      runningTotal += dailyMap.get(dateStr) ?? 0
      spentData.push(runningTotal)
    }

    cur.setDate(cur.getDate() + 1)
  }

  return {
    labels,
    datasets: [
      {
        label: "ยอดใช้สะสม",
        data: spentData,
        borderColor: "#2e8b57", // สีเส้นหลัก (Primary Indigo)
        // backgroundColor: "rgba(99, 102, 241, 0.12)", // พื้นที่ไล่สีใต้กราฟ
        backgroundColor: "rgba(46, 139, 87, 0.15)",
        fill: true,
        tension: 0.35, // เส้นโค้งนุ่มนวล
        borderWidth: 3,
        pointRadius: 4,
        pointHoverRadius: 7,
        pointBackgroundColor: "#2e8b57",
        pointBorderColor: "#ffffff",
        pointBorderWidth: 2,
        pointHoverBackgroundColor: "#05472a", // เมื่อชี้เมาส์: เปลี่ยนเป็นสีเขียวเข้มลึก
        pointHoverBorderColor: "#ffffff",
        pointHoverBorderWidth: 2,
      },
      {
        label: `เพดานงบ (฿${formatMoney(props.budget.amount)})`,
        data: budgetLimitData,
        borderColor: "#ef4444", // สีแดงเตือน
        borderDash: [6, 6], // เส้นประ
        borderWidth: 2,
        pointRadius: 0, // ไม่ต้องโชว์จุดบนเส้นเพดาน
        fill: false,
      },
    ],
  }
})

// 2. ปรับแต่ง Option การแสดงผล (สเกล, แกน X/Y, Tooltip)
const chartOptions = computed<ChartOptions<"line">>(() => ({
  responsive: true,
  maintainAspectRatio: false, // ขยายเต็มความสูงของการ์ด
  interaction: {
    mode: "index",
    intersect: false,
  },
  plugins: {
    legend: {
      position: "top",
      align: "end",
      labels: {
        boxWidth: 12,
        boxHeight: 12,
        usePointStyle: true,
        font: {
          family: "inherit",
          size: 12,
        },
      },
    },
    tooltip: {
      backgroundColor: "rgba(15, 23, 42, 0.9)",
      titleFont: { size: 13, weight: "bold" },
      bodyFont: { size: 12 },
      padding: 12,
      cornerRadius: 10,
      callbacks: {
        label: (context) => {
          return ` ${context.dataset.label}: ฿${formatMoney(Number(context.raw))}`
        },
      },
    },
  },
  scales: {
    x: {
      grid: {
        display: false,
      },
      ticks: {
        maxTicksLimit: 8, // จำกัดไม่ให้ป้ายวันที่เบียดกันเกินไป
        font: { size: 11 },
      },
    },
    y: {
      beginAtZero: true,
      grid: {
        color: "rgba(156, 163, 175, 0.15)", // เส้นตารางแนวนอนสีจางๆ
      },
      ticks: {
        callback: (val) => `฿${formatMoney(Number(val))}`,
        font: { size: 11 },
      },
    },
  },
}))
</script>

<template>
  <div class="card border border-base-300 bg-base-100 shadow-sm">
    <div class="card-body p-5 sm:p-6">
      
      <div>
        <h2 class="text-xl font-bold">กราฟการใช้เงินสะสม</h2>
        <p class="text-sm text-base-content/60">
          แนวโน้มการใช้จ่ายจริงเทียบกับเพดานงบประมาณ
        </p>
      </div>

      <!-- กำหนดความสูงของกราฟให้ใหญ่เต็มตา (h-80 = 320px, sm:h-96 = 384px) -->
      <div class="mt-4 h-80 sm:h-96 w-full">
        <Line :data="chartData" :options="chartOptions" />
      </div>

    </div>
  </div>
</template>
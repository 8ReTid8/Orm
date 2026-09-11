<script setup lang="ts">
import { ref, computed } from "vue"
import type { Budget } from "@/types/budget"
import type { Transaction } from "@/types/transaction"
import { formatMoney } from "@/utils/format"

interface Props {
  budget: Budget
  transactions: Transaction[]
}

const props = defineProps<Props>()

// State สำหรับ Tooltip เมื่อเอาเมาส์ชี้ที่จุด
const hoveredPoint = ref<{
  date: string
  daily: number
  total: number
  x: number
  y: number
} | null>(null)

// คำนวณพิกัด X, Y และเส้น Path ของกราฟ
const chartData = computed(() => {
  const start = new Date(props.budget.startDate)
  const end = new Date(props.budget.endDate)
  const today = new Date()

  // วันสิ้นสุดของการพล็อต (ไม่เกินวันสิ้นสุดงบ)
  const plotEnd = end < today ? end : today

  // รวมยอดใช้จ่ายแยกตามวัน (YYYY-MM-DD)
  const dailyMap = new Map<string, number>()
  for (const t of props.transactions) {
    if (t.type === "expense") {
      const d = t.transactionDate.slice(0, 10)
      dailyMap.set(d, (dailyMap.get(d) ?? 0) + t.amount)
    }
  }

  // สร้างข้อมูลรายวันตั้งแต่วันแรกจนถึงปัจจุบัน
  const dailyPoints: { date: string; daily: number; total: number }[] = []
  let runningTotal = 0

  const cur = new Date(start)
  while (cur <= plotEnd) {
    const dateStr = cur.toISOString().slice(0, 10)
    const daily = dailyMap.get(dateStr) ?? 0
    runningTotal += daily
    dailyPoints.push({
      date: dateStr,
      daily,
      total: runningTotal,
    })
    cur.setDate(cur.getDate() + 1)
  }

  // กำหนดขนาด SVG Grid
  const width = 600
  const height = 220
  const padLeft = 55
  const padRight = 20
  const padTop = 25
  const padBottom = 35

  const chartW = width - padLeft - padRight
  const chartH = height - padTop - padBottom

  // หาค่าสูงสุดของแกน Y (เผื่อยอดใช้เกินงบไว้ 15%)
  const maxVal = Math.max(props.budget.amount, runningTotal) * 1.15 || 100

  const getY = (val: number) => padTop + chartH - (val / maxVal) * chartH
  const getX = (index: number) => {
    if (dailyPoints.length <= 1) return padLeft + chartW / 2
    return padLeft + (index / (dailyPoints.length - 1)) * chartW
  }

  const points = dailyPoints.map((dp, i) => ({
    ...dp,
    x: getX(i),
    y: getY(dp.total),
  }))

  let pathD = ""
  let areaD = ""
  const bottomY = padTop + chartH

  if (points.length > 0) {
    pathD = `M ${points[0].x} ${points[0].y}`
    for (let i = 1; i < points.length; i++) {
      pathD += ` L ${points[i].x} ${points[i].y}`
    }
    areaD = `${pathD} L ${points[points.length - 1].x} ${bottomY} L ${points[0].x} ${bottomY} Z`
  }

  const budgetLineY = getY(props.budget.amount)

  return {
    width,
    height,
    padLeft,
    padRight,
    points,
    pathD,
    areaD,
    budgetLineY,
    bottomY,
  }
})
</script>

<template>
  <div class="card border border-base-300 bg-base-100 shadow-sm">
    <div class="card-body p-5 sm:p-6">
      
      <!-- Header กราฟ & Legend -->
      <div class="flex flex-wrap items-center justify-between gap-2">
        <div>
          <h2 class="text-lg font-bold">กราฟการใช้เงินสะสม</h2>
          <p class="text-xs text-base-content/60">
            แนวโน้มการใช้จ่ายเทียบกับเพดานงบประมาณ
          </p>
        </div>

        <!-- Legend คำอธิบายเส้น -->
        <div class="flex items-center gap-4 text-xs">
          <div class="flex items-center gap-1.5">
            <span class="size-2.5 rounded-full bg-primary" />
            <span class="text-base-content/70">ใช้สะสม</span>
          </div>
          <div class="flex items-center gap-1.5">
            <span class="h-0.5 w-3 border-t-2 border-dashed border-error" />
            <span class="text-base-content/70">เพดานงบ (฿{{ formatMoney(budget.amount) }})</span>
          </div>
        </div>
      </div>

      <!-- SVG Chart Container -->
      <div class="relative mt-4 w-full select-none">
        <svg
          :viewBox="`0 0 ${chartData.width} ${chartData.height}`"
          class="w-full h-56 sm:h-64 overflow-visible"
        >
          <defs>
            <!-- ไล่เฉดสีพื้นที่ใต้กราฟ -->
            <linearGradient id="spendingGradient" x1="0" y1="0" x2="0" y2="1">
              <stop offset="0%" stop-color="var(--color-primary, #10b981)" stop-opacity="0.25" />
              <stop offset="100%" stop-color="var(--color-primary, #10b981)" stop-opacity="0.0" />
            </linearGradient>
          </defs>

          <!-- เส้นฐานแกน X แนวนอน -->
          <line
            :x1="chartData.padLeft"
            :y1="chartData.bottomY"
            :x2="chartData.width - chartData.padRight"
            :y2="chartData.bottomY"
            class="stroke-base-300"
            stroke-width="1"
          />

          <!-- เส้นประเพดานงบประมาณ (Budget Limit Line) -->
          <line
            v-if="chartData.budgetLineY"
            :x1="chartData.padLeft"
            :y1="chartData.budgetLineY"
            :x2="chartData.width - chartData.padRight"
            :y2="chartData.budgetLineY"
            class="stroke-error"
            stroke-width="1.5"
            stroke-dasharray="4 4"
          />

          <!-- ข้อความบอกเพดานงบแกน Y -->
          <text
            v-if="chartData.budgetLineY"
            :x="chartData.padLeft - 8"
            :y="chartData.budgetLineY + 3"
            text-anchor="end"
            class="fill-error text-[10px] font-semibold"
          >
            ฿{{ formatMoney(budget.amount) }}
          </text>

          <!-- พื้นที่ใต้กราฟ (Area Fill) -->
          <path
            v-if="chartData.areaD"
            :d="chartData.areaD"
            fill="url(#spendingGradient)"
          />

          <!-- เส้นกราฟยอดใช้จ่าย (Line Path) -->
          <path
            v-if="chartData.pathD"
            :d="chartData.pathD"
            fill="none"
            class="stroke-primary"
            stroke-width="2.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />

          <!-- จุดในกราฟแต่ละวัน (Points) -->
          <g v-for="pt in chartData.points" :key="pt.date">
            <circle
              :cx="pt.x"
              :cy="pt.y"
              :r="pt.daily > 0 ? 4 : 2"
              class="transition-all duration-150 cursor-pointer"
              :class="pt.daily > 0 ? 'fill-primary stroke-base-100 stroke-2 hover:r-6' : 'fill-primary/40'"
              @mouseenter="hoveredPoint = pt"
              @mouseleave="hoveredPoint = null"
            />
          </g>

          <!-- ป้ายวันที่เริ่มต้น และ วันสิ้นสุด -->
          <text
            :x="chartData.padLeft"
            :y="chartData.height - 8"
            class="fill-base-content/50 text-[10px]"
          >
            {{ budget.startDate }}
          </text>
          <text
            :x="chartData.width - chartData.padRight"
            :y="chartData.height - 8"
            text-anchor="end"
            class="fill-base-content/50 text-[10px]"
          >
            {{ budget.endDate }}
          </text>
        </svg>

        <!-- Floating Tooltip เมื่อ Hover โดนจุดในกราฟ -->
        <div
          v-if="hoveredPoint"
          class="pointer-events-none absolute z-10 -translate-x-1/2 -translate-y-full rounded-xl bg-neutral text-neutral-content px-3 py-2 text-xs shadow-lg transition-all duration-75"
          :style="{
            left: `${(hoveredPoint.x / chartData.width) * 100}%`,
            top: `${(hoveredPoint.y / chartData.height) * 100 - 10}%`,
          }"
        >
          <p class="font-bold">{{ hoveredPoint.date }}</p>
          <p v-if="hoveredPoint.daily > 0" class="text-error font-semibold">
            วันนี้ใช้: ฿{{ formatMoney(hoveredPoint.daily) }}
          </p>
          <p class="text-neutral-content/80">
            ยอดสะสม: ฿{{ formatMoney(hoveredPoint.total) }}
          </p>
        </div>
      </div>

    </div>
  </div>
</template>
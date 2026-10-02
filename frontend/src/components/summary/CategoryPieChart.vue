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

// จานสีสวยงามสำหรับหมวดหมู่ต่างๆ
const palette = [
  "#22c55e", "#3b82f6", "#f59e0b", "#ef4444", "#8b5cf6",
  "#ec4899", "#14b8a6", "#f97316", "#06b6d4", "#84cc16"
]

const emit = defineEmits<{
  select: [categoryName: string]
}>()

// ฟังก์ชันคำนวณสีอัตโนมัติด้วย HSL
function getCategoryColor(index: number, total: number): string {
  const color = palette[index]
  if (color) {
    return color
  }
  const hue = Math.round((index * (360 / Math.max(total, 1)))) % 360
  return `hsl(${hue}, 65%, 55%)`
}

const chartData = computed(() => {
  const total = props.items.length
  return {
    labels: props.items.map(i => i.name),
    datasets: [{
      data: props.items.map(i => i.amount),
      // 👈 เรียกใช้ฟังก์ชันกำหนดสี
      backgroundColor: props.items.map((item, index) => 
        item.color || getCategoryColor(index, total)
      ),
      borderWidth: 2,
      borderColor: "#ffffff",
    }],
  }
})
// const chartData = computed(() => ({
//   labels: props.items.map(i => i.name),
//   datasets: [{
//     data: props.items.map(i => i.amount),
//     backgroundColor: props.items.map((_, index) => palette[index % palette.length]),
//     borderWidth: 2,
//     borderColor: "#ffffff",
//   }],
// }))

const chartOptions = computed(() => ({
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
  // เมื่อชี้โดนชิ้นกราฟ ให้เมาส์เปลี่ยนเป็นรูปมือ (Pointer)
  onHover: (event: any, elements: any[]) => {
    const target = event?.native?.target as HTMLElement
    if (target) {
      target.style.cursor = elements.length > 0 ? "pointer" : "default"
    }
  },
  // เมื่อคลิกที่ชิ้นกราฟ ให้ส่งชื่อหมวดหมู่ออกไป
  onClick: (_event: any, elements: any[]) => {
    if (elements.length > 0) {
      const index = elements[0].index
      const selectedItem = (props.items ?? [])[index]
      if (selectedItem) {
        emit("select", selectedItem.name)
      }
    }
  },
}))
</script>

<template>
  <div class="h-64 w-full">
    <Doughnut v-if="items.length > 0" :data="chartData" :options="chartOptions" />
    <div v-else class="flex h-full items-center justify-center text-base-content/50">
      {{ emptyText }}
    </div>
  </div>
</template>
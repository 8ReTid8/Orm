<script setup lang="ts">
const categories = [
  {
    name: "อาหาร",
    spent: 6379,
    budget: 8000,
  },
  {
    name: "เดินทาง",
    spent: 7368,
    budget: 6000,
  },
  {
    name: "ช้อปปิ้ง",
    spent: 2341,
    budget: 4000,
  },
  {
    name: "บันเทิง",
    spent: 1200,
    budget: 2000,
  },
]

function getPercent(spent: number, budget: number) {
  if (budget === 0) return 0

  return Math.min((spent / budget) * 100, 100)
}

function formatMoney(value: number) {
  return new Intl.NumberFormat("th-TH").format(value)
}
</script>

<template>
  <div class="card border border-base-300 bg-base-100 shadow-sm">

    <div class="card-body">

      <div>
        <h2 class="card-title">
          งบประมาณแต่ละหมวด
        </h2>

        <p class="text-sm text-base-content/50">
          ติดตามการใช้จ่ายของแต่ละหมวด
        </p>
      </div>

      <div class="mt-4 divide-y divide-base-300">

        <div
          v-for="category in categories"
          :key="category.name"
          class="py-5"
        >

          <div class="flex items-center justify-between">

            <div>
              <p class="font-semibold">
                {{ category.name }}
              </p>

              <p class="text-sm text-base-content/50">
                ฿{{ formatMoney(category.spent) }}
                /
                ฿{{ formatMoney(category.budget) }}
              </p>
            </div>

            <p
              class="text-sm font-semibold"
              :class="
                category.spent > category.budget
                  ? 'text-error'
                  : 'text-base-content'
              "
            >
              {{ getPercent(category.spent, category.budget).toFixed(0) }}%
            </p>

          </div>

          <div class="mt-3 h-2 overflow-hidden rounded-full bg-base-200">

            <div
              class="h-full rounded-full transition-all"
              :class="
                category.spent > category.budget
                  ? 'bg-error'
                  : 'bg-success'
              "
              :style="{
                width:
                  getPercent(
                    category.spent,
                    category.budget
                  ) + '%'
              }"
            />

          </div>

        </div>

      </div>

    </div>
  </div>
</template>
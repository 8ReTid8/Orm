<script setup lang="ts">
import { toBuddhistYear, THAI_MONTHS, type MonthOption } from '@/utils/date'

interface Props {
  years: number[]
  months?: readonly MonthOption[]
  modelYear: number | null
  modelMonth: number | null
}

const props = withDefaults(defineProps<Props>(), {
  months: () => THAI_MONTHS,
})

const emit = defineEmits<{
  "update:modelYear": [value: number | null]
  "update:modelMonth": [value: number | null]
  "yearChange": []
  "monthChange": []
}>()

function handleYearChange(event: Event) {
  const select = event.currentTarget as HTMLSelectElement
  const value = select.value ? Number(select.value) : null
  emit("update:modelYear", value)
  emit("yearChange")
}

function handleMonthChange(event: Event) {
  const select = event.currentTarget as HTMLSelectElement
  const value = select.value ? Number(select.value) : null
  emit("update:modelMonth", value)
  emit("monthChange")
}
</script>

<template>
  <div class="flex items-center gap-2">
    <select
      :value="modelYear ?? ''"
      class="select select-bordered min-w-max"
      @change="handleYearChange"
    >
      <option value="">ทุกปี</option>
      <option v-for="year in years" :key="year" :value="year">
        <!-- {{ year + 543 }} -->
          {{ toBuddhistYear(year) }}
      </option>
    </select>

    <select
      :value="modelMonth ?? ''"
      class="select select-bordered min-w-max"
      :disabled="modelYear === null"
      @change="handleMonthChange"
    >
      <option value="">ทุกเดือน</option>
      <option v-for="month in months" :key="month.value" :value="month.value">
        {{ month.label }}
      </option>
    </select>
  </div>
</template>

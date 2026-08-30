<script setup lang="ts">
import type { Category } from "@/types/category";

interface Props {
  categories: Category[]
  modelValue: string | null
}

defineProps<Props>()

const emit = defineEmits<{
  "update:modelValue": [value: string | null]
}>()
function handleChange(event: Event) {
  const select = event.currentTarget as HTMLSelectElement

  emit(
    "update:modelValue",
    // select.value ? Number(select.value) : null
    select.value ? select.value : null
  )
}
</script>

<template>
  <div class=" flex items-center gap-2">
    <label class="shrink-0 text-sm font-medium">
      หมวดหมู่
    </label>

    <select
      :value="modelValue ?? ''"
      class="select select-bordered"
      @change="handleChange"
    >
      <option value="">
        ทุกหมวดหมู่
      </option>

      <option
        v-for="category in categories"
        :key="category.id"
        :value="category.name"
      >
        {{ category.name }}
      </option>
    </select>
  </div>
</template>
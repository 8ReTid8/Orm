<script setup lang="ts">
import type { Account } from "@/types/account"

interface Props {
  accounts: Account[]
  modelValue: number | null
}

defineProps<Props>()

const emit = defineEmits<{
  "update:modelValue": [value: number | null]
}>()
function handleChange(event: Event) {
  const select = event.currentTarget as HTMLSelectElement

  emit(
    "update:modelValue",
    select.value ? Number(select.value) : null
  )
}
</script>

<template>
  <div class="flex items-center gap-2">
    <label class="text-sm font-medium">
      บัญชี
    </label>

    <select
      :value="modelValue ?? ''"
      class="select select-bordered min-w-max  "
      @change="handleChange"
    >
      <option value="">
        ทุกบัญชี
      </option>

      <option
        v-for="account in accounts"
        :key="account.id"
        :value="account.id"
      >
        {{ account.name }}
      </option>
    </select>
  </div>
</template>
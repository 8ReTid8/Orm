<script setup lang="ts">
import type { Account } from "@/types/account"
import { Wallet } from "lucide-vue-next";

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
    <!-- <label class="text-sm font-medium">
      บัญชี
    </label> -->

    <!-- <select
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
    </select> -->
    <div class="relative flex items-center">
      <Wallet class="pointer-events-none absolute left-3 size-4 text-base-content/60 z-10" />
      <select :value="modelValue ?? ''"
        class="select select-bordered select-sm h-10 pl-9 min-w-max text-sm font-normal" @change="handleChange">
        <option value="">
          ทุกบัญชี
        </option>
        <option v-for="account in accounts" :key="account.id" :value="account.id">
          {{ account.name }}
        </option>
      </select>
    </div>
  </div>
</template>
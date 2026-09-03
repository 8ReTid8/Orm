<script setup lang="ts">
import type { Component } from "vue"

interface Props {
  icon: Component
  title: string
  description?: string
}

defineProps<Props>()
</script>

<template>
  <div
    class="flex flex-col items-center justify-center rounded-xl border border-dashed border-base-300 py-12 text-center"
  >
    <!-- <component
      :is="icon"
      class="size-10 text-base-content/30"
    /> -->
    <component
      :is="icon"
      class="size-10 text-base-content/30"
    />

    <p class="mt-3 font-medium">
      {{ title }}
    </p>

    <p
      v-if="description"
      class="mt-1 text-sm text-base-content/50"
    >
      {{ description }}
    </p>
  </div>
</template>
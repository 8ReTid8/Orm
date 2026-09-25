<script setup lang="ts">
import { ChevronLeft, ChevronRight } from "lucide-vue-next"

interface Props {
  currentPage: number
  totalPages: number
  disabled?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  disabled: false,
})

const emit = defineEmits<{
  "update:currentPage": [page: number]
  change: [page: number]
}>()

function changePage(newPage: number) {
  if (
    newPage < 1 ||
    newPage > props.totalPages ||
    newPage === props.currentPage ||
    props.disabled
  ) {
    return
  }

  emit("update:currentPage", newPage)
  emit("change", newPage)
}
</script>

<template>
  <div v-if="totalPages > 1" class="mt-6 flex items-center justify-center gap-3">
    <!-- ปุ่มก่อนหน้า -->
    <button
      type="button"
      class="btn btn-sm gap-1"
      :disabled="currentPage <= 1 || disabled"
      @click="changePage(currentPage - 1)"
    >
      <ChevronLeft class="size-4" />
      <span>ก่อนหน้า</span>
    </button>

    <!-- ตัวเลขแสดงหน้า -->
    <span class="text-sm font-medium text-base-content/70">
      หน้า {{ currentPage }} / {{ totalPages }}
    </span>

    <!-- ปุ่มถัดไป -->
    <button
      type="button"
      class="btn btn-sm gap-1"
      :disabled="currentPage >= totalPages || disabled"
      @click="changePage(currentPage + 1)"
    >
      <span>ถัดไป</span>
      <ChevronRight class="size-4" />
    </button>
  </div>
</template>
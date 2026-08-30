<script setup lang="ts">
import { ref, watch } from "vue"
import { AlertTriangle, Info, AlertCircle } from "lucide-vue-next"

interface Props {
  open: boolean
  title?: string
  message?: string
  confirmText?: string
  cancelText?: string
  type?: "danger" | "warning" | "info"
  loading?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  title: "ยืนยันการทำรายการ?",
  message: "คุณแน่ใจหรือไม่ที่จะดำเนินการนี้ การกระทำนี้ไม่สามารถย้อนกลับได้",
  confirmText: "ยืนยัน",
  cancelText: "ยกเลิก",
  type: "danger",
  loading: false,
})

const emit = defineEmits<{
  confirm: []
  close: []
}>()

const dialogRef = ref<HTMLDialogElement | null>(null)

watch(
  () => props.open,
  (isOpen) => {
    if (isOpen) {
      dialogRef.value?.showModal()
    } else {
      dialogRef.value?.close()
    }
  },
  { immediate: true }
)

function handleConfirm() {
  emit("confirm")
}

function handleClose() {
  if (!props.loading) {
    emit("close")
  }
}
</script>

<template>
  <dialog ref="dialogRef" class="modal" @close="handleClose">
    <div class="modal-box max-w-sm text-center">
      <!-- Icon ตามประเภท (danger / warning / info) -->
      <!-- <div
        class="mx-auto mb-4 flex size-14 items-center justify-center rounded-full"
        :class="{
          'bg-error/15 text-error': type === 'danger',
          'bg-warning/15 text-warning': type === 'warning',
          'bg-info/15 text-info': type === 'info',
        }"
      >
        <AlertTriangle v-if="type === 'danger' || type === 'warning'" class="size-7" />
        <Info v-else class="size-7" />
      </div> -->

      <!-- Title -->
      <h3 class="text-lg font-bold">
        {{ title }}
      </h3>

      <!-- Message / Slot รายละเอียดเพิ่มเติม -->
      <div class="mt-2 text-sm text-base-content/70">
        <slot>
          <p>{{ message }}</p>
        </slot>
      </div>

      <!-- Action Buttons -->
      <div class="modal-action mt-6 justify-center gap-3">
        <button
          type="button"
          class="btn btn-ghost"
          :disabled="loading"
          @click="handleClose"
        >
          {{ cancelText }}
        </button>

        <button
          type="button"
          class="btn"
          :class="type === 'danger' ? 'btn-error text-white' : 'btn-primary'"
          :disabled="loading"
          @click="handleConfirm"
        >
          <span v-if="loading" class="loading loading-spinner loading-sm" />
          {{ confirmText }}
        </button>
      </div>
    </div>

    <!-- Backdrop สำหรับกดปิดข้างนอก -->
    <form method="dialog" class="modal-backdrop" @submit.prevent="handleClose">
      <button :disabled="loading">close</button>
    </form>
  </dialog>
</template>
<script setup lang="ts">
import { ref, watch } from "vue"
import { WalletCards } from "lucide-vue-next"

interface Props {
  open: boolean
}

const props = defineProps<Props>()

const emit = defineEmits<{
  close: []
  save: [
    {
      name: string
    //   balance: number
    }
  ]
}>()

const dialogRef = ref<HTMLDialogElement | null>(null)

const name = ref("")
// const balance = ref<number | null>(null)

watch(
  () => props.open,
  (isOpen) => {
    if (isOpen) {
      dialogRef.value?.showModal()
    } else {
      dialogRef.value?.close()
    }
  },
)

function submitForm() {
  if (!name.value.trim()) {
    return
  }

  emit("save", {
    name: name.value.trim(),
    // balance: balance.value,
  })

  resetForm()
}

function closeDialog() {
  emit("close")
}

function resetForm() {
  name.value = ""
//   balance.value = null
}
</script>

<template>
  <dialog ref="dialogRef" class="modal">
    <div class="modal-box max-w-md">
      <div class="mb-6 flex items-center gap-3">
        <div
          class="flex size-11 items-center justify-center rounded-xl bg-primary/15 text-primary"
        >
          <WalletCards class="size-6" />
        </div>

        <div>
          <h2 class="text-xl font-bold">
            เพิ่มบัญชี
          </h2>

          <p class="text-sm text-base-content/60">
            สร้างบัญชีสำหรับบันทึกรายรับรายจ่าย
          </p>
        </div>
      </div>

      <form
        class="flex flex-col gap-5"
        @submit.prevent="submitForm"
      >
        <fieldset class="fieldset">
          <legend class="fieldset-legend">
            ชื่อบัญชี
          </legend>

          <input
            v-model.trim="name"
            type="text"
            class="input w-full"
            placeholder="เช่น KBank เงินเดือน"
            required
          />
        </fieldset>

        <!-- <fieldset class="fieldset">
          <legend class="fieldset-legend">
            ยอดเริ่มต้น
          </legend>

          <input
            v-model.number="balance"
            type="number"
            min="0"
            step="0.01"
            class="input w-full"
            placeholder="0.00"
            required
          />
        </fieldset> -->

        <div class="modal-action">
          <button
            type="button"
            class="btn btn-ghost"
            @click="closeDialog"
          >
            ยกเลิก
          </button>

          <button
            type="submit"
            class="btn btn-primary"
          >
            เพิ่มบัญชี
          </button>
        </div>
      </form>
    </div>

    <form
      method="dialog"
      class="modal-backdrop"
      @submit="closeDialog"
    >
      <button aria-label="close">
        close
      </button>
    </form>
  </dialog>
</template>
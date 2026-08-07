<script setup lang="ts">
import type { Bank, TransactionForm } from "@/types/transaction"
import { ref, watch } from "vue";

interface Props {
    open: boolean
    form: TransactionForm
    selectedDateText: string
}
const props = defineProps<Props>()
const emit = defineEmits<{
    close: []
    save: []
}>()
const dialogRef = ref<HTMLDialogElement | null>(null)
const banks: Bank[] = [
    { id: 1, name: "เงินสด" },
    { id: 2, name: "KBank" },
    { id: 3, name: "SCB" },
    { id: 4, name: "Krungthai" },
    { id: 5, name: "Bangkok Bank" },
    { id: 6, name: "TrueMoney" },
]
watch(
    () => props.open,
    (isOpen) => {
        if (isOpen) {
            dialogRef.value?.showModal()
        } else {
            dialogRef.value?.close()
        }
    },
    {
        immediate: true,
    },
)
const slipPreview = ref<string | null>(null)

function handleSlipChange(event: Event) {
    const input = event.target as HTMLInputElement
    const file = input.files?.[0]

    if (!file) {
        return
    }

    if (!file.type.startsWith("image/")) {
        input.value = ""
        return
    }

    const maxSize = 5 * 1024 * 1024

    if (file.size > maxSize) {
        input.value = ""
        return
    }

    if (slipPreview.value) {
        URL.revokeObjectURL(slipPreview.value)
    }

    props.form.slipImage = file
    slipPreview.value = URL.createObjectURL(file)
}

function removeSlip() {
    props.form.slipImage = null

    if (slipPreview.value) {
        URL.revokeObjectURL(slipPreview.value)
        slipPreview.value = null
    }
}
function submitForm() {
    emit("save")
}

function closeDialog() {
    emit("close")
}

</script>

<template>
    <dialog ref="dialogRef" class="modal">
        <div class="modal-box max-w-lg">
            <h2 class="text-xl font-bold">เพิ่มรายการ</h2>

            <p class="mt-1 text-sm text-base-content/60">
                วันที่ {{ selectedDateText }}
            </p>

            <form class="mt-6 space-y-4" @submit.prevent="submitForm">
                <!-- Type -->
                <fieldset>
                    <legend class="mb-2 font-medium">ประเภทรายการ</legend>

                    <div class="grid grid-cols-2 gap-3">
                        <label class="btn" :class="form.type === 'income'
                            ? 'btn-success'
                            : 'btn-outline'">
                            <input v-model="form.type" type="radio" value="income" class="hidden" />
                            รายรับ
                        </label>

                        <label class="btn" :class="form.type === 'expense'
                            ? 'btn-error'
                            : 'btn-outline'">
                            <input v-model="form.type" type="radio" value="expense" class="hidden" />
                            รายจ่าย
                        </label>
                    </div>
                </fieldset>

                <!-- Title -->
                <fieldset class="fieldset">
                    <legend class="fieldset-legend">ชื่อรายการ</legend>

                    <input v-model.trim="form.title" type="text" class="input w-full" placeholder="เช่น ค่าอาหาร"
                        required />
                </fieldset>

                <!-- Amount -->
                <fieldset class="fieldset">
                    <legend class="fieldset-legend">จำนวนเงิน</legend>

                    <input v-model.number="form.amount" type="number" min="0.01" step="0.01" class="input w-full"
                        placeholder="0.00" required />
                </fieldset>

                <!-- Category -->
                <fieldset class="fieldset">
                    <legend class="fieldset-legend">หมวดหมู่</legend>
                    <select v-model="form.category" class="select w-full" required>
                        <option disabled value="">เลือกหมวดหมู่</option>
                        <option value="food">อาหาร</option>
                        <option value="transport">การเดินทาง</option>
                        <option value="salary">เงินเดือน</option>
                        <option value="shopping">ช้อปปิ้ง</option>
                        <option value="other">อื่น ๆ</option>
                    </select>
                </fieldset>
                <fieldset class="fieldset">
                    <legend class="fieldset-legend">
                        ธนาคาร
                    </legend>

                    <select v-model.number="form.bankId" class="select w-full" required>
                        <option disabled :value="null">
                            เลือกธนาคาร
                        </option>

                        <option v-for="bank in banks" :key="bank.id" :value="bank.id">
                            {{ bank.name }}
                        </option>
                    </select>
                </fieldset>
                <!-- Note -->
                <fieldset class="fieldset">
                    <legend class="fieldset-legend">หมายเหตุ</legend>

                    <textarea v-model.trim="form.note" class="textarea w-full" placeholder="รายละเอียดเพิ่มเติม" />
                </fieldset>
                <fieldset class="fieldset">
                    <legend class="fieldset-legend">
                        รูปสลิป
                    </legend>

                    <label class="
      flex
      min-h-32
      cursor-pointer
      flex-col
      items-center
      justify-center
      rounded-box
      border
      border-dashed
      border-base-300
      bg-base-200/50
      p-4
      text-center
      transition
      hover:border-primary
      hover:bg-primary/5
    ">
                        <input type="file" accept="image/png,image/jpeg,image/webp" class="hidden"
                            @change="handleSlipChange" />

                        <template v-if="!slipPreview">
                            <ImagePlus class="mb-2 size-8 text-base-content/50" />

                            <p class="text-sm font-medium">
                                คลิกเพื่ออัปโหลดสลิป
                            </p>

                            <p class="mt-1 text-xs text-base-content/50">
                                PNG, JPG หรือ WebP ไม่เกิน 5 MB
                            </p>
                        </template>

                        <div v-else class="relative w-full">
                            <img :src="slipPreview" alt="ตัวอย่างสลิป"
                                class="mx-auto max-h-52 rounded-lg object-contain" />

                            <button type="button" class="btn btn-circle btn-error btn-sm absolute right-0 top-0"
                                aria-label="ลบรูปสลิป" @click.prevent.stop="removeSlip">
                                <X class="size-4" />
                            </button>
                        </div>
                    </label>
                </fieldset>

                <div class="modal-action">
                    <!-- <button type="button" class="btn btn-ghost" @click="dialogRef?.close()"> -->
                    <button type="button" class="btn btn-ghost" @click="closeDialog">
                        ยกเลิก
                    </button>

                    <!-- <button type="submit" class="btn btn-primary"> -->
                    <button type="submit" class="btn btn-primary">
                        บันทึก
                    </button>
                </div>
            </form>
        </div>

        <form method="dialog" class="modal-backdrop" @submit="closeDialog">
            <button aria-label="Close dialog">close</button>
        </form>
    </dialog>
</template>

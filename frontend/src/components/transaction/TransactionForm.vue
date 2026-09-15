<script setup lang="ts">
import type { Account } from "@/types/account";
import type { Category } from "@/types/category";
import type { TransactionForm } from "@/types/transaction"
import { resolveCategoryIcon } from "@/utils/categoryIcons";
import { formatThaiDateLong } from "@/utils/format";
import { ImagePlus, Tag } from "lucide-vue-next";
import { computed, ref, watch } from "vue";

interface Props {
    open: boolean
    form: TransactionForm
    accounts: Account[]
    categories: Category[]
    // selectedDate: Date
}
const props = defineProps<Props>()
const emit = defineEmits<{
    close: []
    save: []
}>()
const dialogRef = ref<HTMLDialogElement | null>(null)

const filteredCategories = computed(() => {
    return props.categories.filter(
        (category) => category.type === props.form.type
    )
})


watch(
    () => props.open,
    (isOpen) => {
        if (isOpen) {
            dialogRef.value?.showModal()
        } else {
            dialogRef.value?.close()
        }
    },
    { immediate: true },
)

watch(
    () => props.form.type,
    (newType, oldType) => {
        // ทำงานเฉพาะตอนที่มีการสลับ type จริงๆ
        if (oldType && newType !== oldType) {
            // เช็คว่า category ปัจจุบัน มีอยู่ในหมวดหมู่ของประเภทใหม่หรือไม่
            const isValidForNewType = props.categories.some(
                (c) => c.type === newType && c.name === props.form.category
            )
            // ถ้า category เดิมไม่ตรงกับประเภทใหม่ ถึงจะล้างค่า
            if (!isValidForNewType) {
                props.form.category = ""
            }
        }
    }
)
const slipPreview = ref<string | null>(null)

const selectedCategoryIcon = computed(() => {
    const selectedCat = filteredCategories.value.find((c) => c.name === props.form.category)
    return selectedCat?.icon ? resolveCategoryIcon(selectedCat.icon) : Tag
})

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
    if (!props.form.accountId) {
        alert("กรุณาเลือกบัญชีก่อนบันทึก")
        return 
    }
    emit("save")
}

function closeDialog() {
    emit("close")
}

</script>

<template>
    <dialog ref="dialogRef" class="modal">
        <div class="modal-box max-w-lg ">
            <h2 class="text-xl font-bold">เพิ่มรายการ</h2>

            <p class="mt-1 text-base text-base-content/60">
                <!-- วันที่ {{ selectedDateText }} -->
                <!-- วันที่ {{ formatThaiDateLong(selectedDate) }} -->
                วันที่ {{ formatThaiDateLong(new Date(form.transactionDate)) }}
            </p>

            <form class="mt-6 space-y-4 " @submit.prevent="submitForm">
                <!-- <fieldset class="fieldset gap-0.5">
                    <label class="label text-base">วันที่ทำรายการ</label>
                    <input v-model="form.transactionDate" type="date" class="input w-full" required />
                </fieldset> -->
                <!-- Type -->
                <fieldset class="fieldset gap-0.5">
                    <!-- <legend class="mb-2 font-medium">ประเภทรายการ</legend> -->
                    <label class="label text-base">ประเภทรายการ</label>
                    <div class="grid grid-cols-2 gap-3">
                        <label class="btn" :class="form.type === 'income'
                            ? 'btn-success text-white'
                            : 'btn-outline'">
                            <input v-model="form.type" type="radio" value="income" class="hidden" />
                            รายรับ
                        </label>

                        <label class="btn" :class="form.type === 'expense'
                            ? 'btn-error text-white'
                            : 'btn-outline'">
                            <input v-model="form.type" type="radio" value="expense" class="hidden" />
                            รายจ่าย
                        </label>
                    </div>
                </fieldset>

                <!-- Title -->
                <fieldset class="fieldset gap-0.5">
                    <label class="label text-base">ชื่อรายการ</label>
                    <input v-model.trim="form.title" type="text" class="input w-full" placeholder="เช่น ค่าอาหาร"
                        required />
                </fieldset>

                <!-- Amount -->
                <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
                    <fieldset class="fieldset gap-0.5">
                        <!-- <legend class="fieldset-legend text-lg">จำนวนเงิน</legend> -->
                        <label class="label text-base">จำนวนเงิน</label>
                        <input v-model.number="form.amount" type="number" min="0.01" step="0.01" class="input w-full"
                            placeholder="0.00" required />
                    </fieldset>

                    <!-- Category -->
                    <fieldset class="fieldset gap-0.5">
                        <label class="label text-base">หมวดหมู่</label>
                        <div class="relative flex items-center">
                            <!-- Icon ฝั่งซ้าย (จะเปลี่ยนตาม Category ที่เลือก) -->
                            <component :is="selectedCategoryIcon"
                                class="pointer-events-none absolute left-3 size-5 text-base-content/50 z-10" />
                            <!-- เพิ่ม pl-10 เพื่อเว้นที่ให้ Icon ฝั่งซ้าย -->
                            <select v-model="form.category" class="select w-full pl-10" required>
                                <option disabled value="">เลือกหมวดหมู่</option>
                                <option v-for="category in filteredCategories" :key="category.id"
                                    :value="category.name">
                                    {{ category.name }}
                                </option>
                            </select>
                        </div>
                    </fieldset>
                </div>
                <fieldset class="fieldset gap-0.5">
                    <label class="label text-base">บัญชี</label>
                    <select v-model.number="form.accountId" class="select w-full" required>
                        <option disabled :value="null">
                            เลือกบัญชี
                        </option>
                        <option v-for="account in accounts" :key="account.id" :value="account.id">
                            {{ account.name }}
                        </option>
                    </select>
                </fieldset>

                <!-- Note -->
                <fieldset class="fieldset gap-0.5">
                    <label class="label text-base">หมายเหตุ</label>
                    <textarea v-model.trim="form.note" class="textarea w-full" placeholder="รายละเอียดเพิ่มเติม" />
                </fieldset>

                <fieldset class="fieldset gap-0.5">
                    <label class="label text-base">รูปสลิป</label>
                    <label
                        class="flex min-h-32 cursor-pointer flex-col items-center justify-center rounded-box border border-dashed border-base-300 bg-base-200/50 p-4 text-center transition hover:border-primary hover:bg-primary/5">
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

                <div class="modal-action mt-2!">
                    <button type="button" class="btn btn-ghost" @click="closeDialog">
                        ยกเลิก
                    </button>

                    <button type="submit" class="btn text-white bg-green-700">
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

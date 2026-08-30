<script setup lang="ts">
import { ref, watch } from "vue"
import type { Account } from "@/types/account"
import type { Category } from "@/types/category"
import type { BudgetForm } from "@/types/budget"

interface Props {
    open: boolean
    form: BudgetForm
    accounts: Account[]
    categories: Category[]
}

const props = defineProps<Props>()
const emit = defineEmits<{
    close: []
    // save: [form: BudgetForm]
    save: []
}>()
const dialogRef = ref<HTMLDialogElement | null>(null)

// const form = ref<BudgetForm>({
//     accountId: null,
//     category: "",
//     amount: null,
//     startDate: "",
//     endDate: "",
// })
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
    if (!props.form.category || !props.form.amount || !props.form.startDate || !props.form.endDate || !props.form.accountId) {
        return
    }
    emit("save")
    // emit("save", { ...form.value })
    // resetForm()
}
function closeDialog() {
    emit("close")
}
// function resetForm() {
//     form.value = {
//         accountId: null,
//         category: "",
//         amount: null,
//         startDate: "",
//         endDate: "",
//     }
// }
</script>
<template>
    <dialog ref="dialogRef" class="modal">
        <div class="modal-box max-w-lg">
            <h2 class="text-xl font-bold">สร้าง Budget</h2>
            <p class="mt-1 text-sm text-base-content/60">
                กำหนดงบประมาณตามช่วงเวลาที่ต้องการ
            </p>
            <form class="mt-6 space-y-4" @submit.prevent="submitForm">

                <!-- Account -->
                <fieldset class="fieldset gap-0.5">
                    <label class="label text-base font-medium">บัญชี</label>
                    <select v-model.number="form.accountId" class="select select-bordered w-full" required>
                        <option disabled :value="null">เลือกบัญชี</option>
                        <option v-for="account in accounts" :key="account.id" :value="account.id">
                            {{ account.name }}
                        </option>
                    </select>
                </fieldset>
                <!-- Category -->
                <fieldset class="fieldset gap-0.5">
                    <label class="label text-base font-medium">หมวดหมู่ (รายจ่าย)</label>
                    <select v-model="form.category" class="select select-bordered w-full" required>
                        <option disabled value="">เลือกหมวดหมู่</option>
                        <option v-for="category in categories.filter(c => c.type === 'expense')" :key="category.id"
                            :value="category.name">
                            {{ category.name }}
                        </option>
                    </select>
                </fieldset>
                <!-- Amount -->
                <fieldset class="fieldset gap-0.5">
                    <label class="label text-base font-medium">จำนวนเงิน</label>
                    <input v-model.number="form.amount" type="number" min="0" step="0.01"
                        class="input input-bordered w-full" placeholder="0.00" required />
                </fieldset>
                <!-- Date Range -->
                <div class="grid grid-cols-2 gap-4">
                    <fieldset class="fieldset gap-0.5">
                        <label class="label text-base font-medium">วันเริ่มต้น</label>
                        <input v-model="form.startDate" type="date" class="input input-bordered w-full" required />
                    </fieldset>
                    <fieldset class="fieldset gap-0.5">
                        <label class="label text-base font-medium">วันสิ้นสุด</label>
                        <input v-model="form.endDate" type="date" class="input input-bordered w-full"
                            :min="form.startDate" required />
                    </fieldset>
                </div>
                <div class="modal-action !mt-2">
                    <button type="button" class="btn btn-ghost" @click="closeDialog">
                        ยกเลิก
                    </button>
                    <button type="submit" class="btn text-white bg-green-700">
                        บันทึก Budget
                    </button>
                </div>
            </form>
        </div>
        <form method="dialog" class="modal-backdrop" @submit="closeDialog">
            <button aria-label="Close dialog">close</button>
        </form>
    </dialog>
</template>
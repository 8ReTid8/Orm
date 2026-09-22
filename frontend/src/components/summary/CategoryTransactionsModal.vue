<script setup lang="ts">
import { ref, watch, computed } from "vue"
import { X, Calendar, ReceiptText } from "lucide-vue-next"
import type { Account } from "@/types/account"
import type { Category } from "@/types/category"
import type { Transaction } from "@/types/transaction"
import { useTransactions } from "@/composables/transaction/useTransaction"
import { useTransactionForm } from "@/composables/transaction/useTransactionForm"
import { resolveCategoryIcon } from "@/utils/categoryIcons"
import { formatMoney } from "@/utils/format"
import TransactionForm from "@/components/transaction/TransactionForm.vue"
import TotalList from "@/components/common/TotalList.vue"
import GroupedTransactionList from "../transaction/GroupedTransactionList.vue"

interface Props {
    open: boolean
    categoryName: string
    categoryType?: "expense" | "income"
    year: number | null
    month: number | null
    accountId: number | null
    accounts: Account[]
    categories: Category[]
}

const props = withDefaults(defineProps<Props>(), {
    categoryType: "expense",
})

const emit = defineEmits<{
    close: []
    refresh: [] // แจ้ง SummaryView ให้โหลด useSummary ใหม่เมื่อมีการแก้ไข/ลบ
}>()

const dialogRef = ref<HTMLDialogElement | null>(null)
const isEditFormOpen = ref(false)

const {
    transactions,
    isLoadingTransactions,
    loadTransactions,
    deleteTransaction,
    saveTransaction: saveTransactionApi,
    startEdit,
    cancelEdit,
} = useTransactions()

const { form, setEditForm, resetForm } = useTransactionForm()

// รวมยอดเงินทั้งหมดในหมวดนี้
const totalAmount = computed(() =>
    transactions.value.reduce((sum, t) => sum + t.amount, 0)
)

// ดึง Icon หมวดหมู่
const categoryIcon = computed(() => {
    const cat = props.categories.find((c) => c.name === props.categoryName)
    return resolveCategoryIcon(cat?.icon)
})

// ดึงข้อมูลเมื่อ Modal เปิด
async function fetchCategoryTransactions() {
    if (!props.categoryName || !props.year) return

    await loadTransactions({
        year: props.year,
        month: props.month,
        accountId: props.accountId,
        category: props.categoryName,
    })
}

// ควบคุมการเปิด/ปิด Dialog
watch(
    () => props.open,
    async (isOpen) => {
        if (isOpen) {
            dialogRef.value?.showModal()
            await fetchCategoryTransactions()
        } else {
            dialogRef.value?.close()
        }
    }
)

function handleClose() {
    emit("close")
}

// จัดการ Edit / Delete
function openEditTransaction(transaction: Transaction) {
    startEdit(transaction)
    setEditForm(transaction)
    isEditFormOpen.value = true
}

function closeEditForm() {
    isEditFormOpen.value = false
    cancelEdit()
    //   resetForm()
}

async function handleSaveTransaction() {
    try {
        await saveTransactionApi(form.value)
        closeEditForm()
        await fetchCategoryTransactions()
        emit("refresh")
    } catch (error) {
        console.error("Save transaction failed:", error)
    }
}

async function handleDeleteTransaction(id: number) {
    try {
        await deleteTransaction(id)
        await fetchCategoryTransactions()
        emit("refresh")
    } catch (error) {
        console.error("Delete transaction failed:", error)
    }
}
</script>

<template>
    <dialog ref="dialogRef" class="modal modal-bottom sm:modal-middle" @close="handleClose">
        <div class="modal-box max-w-2xl p-0 flex flex-col max-h-[85vh] bg-base-100">

            <!-- Header -->
            <div class="flex items-center justify-between p-5 border-b border-base-200 sticky top-0 bg-base-100 z-10">
                <div class="flex items-center gap-3">
                    <div class="flex size-11 shrink-0 items-center justify-center rounded-xl font-bold"
                        :class="categoryType === 'income' ? 'bg-success/15 text-success' : 'bg-error/15 text-error'">
                        <component :is="categoryIcon" class="size-6" />
                    </div>
                    <div>
                        <div class="flex items-center gap-2">
                            <h3 class="font-bold text-lg">หมวดหมู่: {{ categoryName }}</h3>
                            <span class="badge badge-sm font-semibold"
                                :class="categoryType === 'income' ? 'badge-success text-white' : 'badge-error text-white'">
                                {{ categoryType === 'income' ? 'รายรับ' : 'รายจ่าย' }}
                            </span>
                        </div>
                        <p class="text-xs text-base-content/60 mt-0.5">
                            ยอดรวม: <span class="font-bold"
                                :class="categoryType === 'income' ? 'text-success' : 'text-error'">
                                ฿{{ formatMoney(totalAmount) }}
                            </span>
                        </p>
                    </div>
                </div>

                <div class="flex items-center gap-2">
                    <TotalList :total="transactions.length" />
                    <button type="button" class="btn btn-ghost btn-circle btn-sm" @click="handleClose">
                        <X class="size-5" />
                    </button>
                </div>
            </div>

            <!-- Body: Content & List -->
            <div class="p-5 overflow-y-auto flex-1">
                <GroupedTransactionList :transactions="transactions" :accounts="accounts" :categories="categories"
                    :loading="isLoadingTransactions" empty-title="ไม่พบรายการในหมวดนี้"
                    empty-description="ไม่มีการบันทึกรายการสำหรับหมวดหมู่นี้ในช่วงเวลาที่เลือก"
                    @edit="openEditTransaction" @delete="handleDeleteTransaction" />
            </div>

        </div>

        <!-- Backdrop ปิดเมื่อคลิกข้างนอก -->
        <form method="dialog" class="modal-backdrop" @submit.prevent="handleClose">
            <button>close</button>
        </form>
    </dialog>

    <!-- ฟอร์มแก้ไข (เมื่อกดแก้ไขจากการ์ด) -->
    <TransactionForm :open="isEditFormOpen" :form="form" :accounts="accounts" :categories="categories"
        @close="closeEditForm" @save="handleSaveTransaction" />
</template>
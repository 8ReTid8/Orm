<script setup lang="ts">
import type { Account } from '@/types/account'
import type { Category } from '@/types/category'
import type { Transaction } from '@/types/transaction'
import { resolveCategoryIcon } from '@/utils/categoryIcons'
import { formatMoney } from '@/utils/format'
import { ref } from 'vue'
import ActionButton from '../common/ActionButton.vue'
import ConfirmModal from '../common/ConfirmModal.vue'
import { FileText, Wallet } from 'lucide-vue-next'

interface Props {
    transaction: Transaction
    categories: Category[]
    accounts: Account[]
}
const targetTransaction = ref<Transaction | null>(null)

const props = defineProps<Props>()
const emit = defineEmits<{
    edit: [transaction: Transaction]
    delete: [id: number]
}>()

function getCategoryIcon(categoryName: string) {
    const cat = props.categories.find(c => c.name === categoryName)
    return resolveCategoryIcon(cat?.icon)
}
function getAccountName(accountId: number) {
    const acc = props.accounts.find((a) => a.id === accountId)
    return acc ? acc.name : "ไม่พบบัญชี"
}
function openDeleteModal(transaction: Transaction) {
    targetTransaction.value = transaction
}
function handleConfirmDelete() {
    if (targetTransaction.value) {
        emit("delete", targetTransaction.value.id)
        targetTransaction.value = null
    }
}
</script>
<template>
    <!-- <div v-for="transaction in transactions" :key="transaction.id"
        class="group flex items-center justify-between gap-4 rounded-xl border border-transparent p-3.5 transition-all hover:border-base-300 hover:bg-base-200/50"> -->
    <div class="group flex items-center justify-between gap-4 rounded-xl border border-transparent p-3.5 transition-all hover:border-base-300 hover:bg-base-200/50">
        <!-- Left: Icon & Info -->
        <div class="flex items-center gap-3.5 min-w-0">
            <!-- Type Icon Circle -->
            <div class="flex size-10 shrink-0 items-center justify-center rounded-xl font-bold" :class="transaction.type === 'income'
                ? 'bg-success/15 text-success'
                : 'bg-error/15 text-error'
                ">
                <component :is="getCategoryIcon(transaction.category)" class="size-5" />
            </div>

            <!-- Title & Details -->
            <div class="min-w-0">
                <div class="flex items-center gap-2">
                    <p class="truncate font-semibold text-base">
                        {{ transaction.title }}
                    </p>
                    <!-- Badge ถ้ามี note หรือ slip -->
                    <FileText v-if="transaction.note" class="size-3.5 text-base-content/40 shrink-0"
                        title="มีบันทึกข้อความ" />
                </div>

                <div class="mt-0.5 flex flex-wrap items-center gap-1.5 text-sm text-base-content/60">
                    <span class="badge badge-sm badge-ghost font-normal">{{ transaction.category }}</span>
                    <span>•</span>
                    <span class="flex items-center gap-1">
                        <Wallet class="size-3.5" />
                        {{ getAccountName(transaction.accountId) }}
                    </span>
                </div>
            </div>
        </div>

        <!-- Right: Amount & Action Buttons -->
        <div class="flex items-center gap-3 shrink-0">
            <p class="text-right font-bold text-base md:text-lg" :class="transaction.type === 'income'
                ? 'text-success'
                : 'text-error'
                ">
                {{ transaction.type === 'income' ? '+' : '-' }}฿{{ formatMoney(transaction.amount) }}
            </p>
            <ActionButton edit-title="แก้ไขรายการ" delete-title="ลบรายการ" @edit="emit('edit', transaction)"
                @delete="openDeleteModal(transaction)" />
        </div>
    </div>
    <ConfirmModal :open="targetTransaction !== null" title="ยืนยันการลบรายการ?" confirm-text="ลบรายการ" type="danger"
        @close="targetTransaction = null" @confirm="handleConfirmDelete" />
</template>
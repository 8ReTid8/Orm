<script setup lang="ts">
import { computed } from "vue"
import { Calendar, ReceiptText } from "lucide-vue-next"
import type { Account } from "@/types/account"
import type { Category } from "@/types/category"
import type { Transaction } from "@/types/transaction"
import { groupTransactionsByDate } from "@/utils/transaction"
import TransactionCard from "./TransactionCard.vue"
import EmptyState from "../common/EmptyState.vue"

interface Props {
  transactions: Transaction[]
  accounts: Account[]
  categories: Category[]
  loading?: boolean
  emptyTitle?: string
  emptyDescription?: string
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
  emptyTitle: "ยังไม่มีรายการธุรกรรม",
  emptyDescription: "ไม่มีการบันทึกรายการในช่วงเวลาที่กำหนด",
})

const emit = defineEmits<{
  edit: [transaction: Transaction]
  delete: [id: number]
}>()

// รวมกลุ่มรายการตามวัน
const groupedTransactions = computed(() =>
  groupTransactionsByDate(props.transactions)
)
</script>

<template>
  <div>
    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center py-12">
      <span class="loading loading-spinner loading-md text-primary" />
    </div>

    <!-- Empty State -->
    <EmptyState
      v-else-if="transactions.length === 0"
      :icon="ReceiptText"
      :title="emptyTitle"
      :description="emptyDescription"
    />

    <!-- Grouped Transaction List -->
    <div v-else class="space-y-4">
      <div
        v-for="group in groupedTransactions"
        :key="group.dateKey"
        class="space-y-2"
      >
        <!-- 📅 วันที่กำกับด้านบน -->
        <div class="px-1 text-sm text-base-content/60 flex items-center gap-1.5 font-medium">
          <Calendar class="size-3.5 text-base-content/50" />
          <span>{{ group.formattedDate }}</span>
        </div>

        <!-- Cards รายการในแต่ละวัน -->
        <div class="space-y-1.5">
          <TransactionCard
            v-for="transaction in group.transactions"
            :key="transaction.id"
            :transaction="transaction"
            :accounts="accounts"
            :categories="categories"
            @edit="emit('edit', $event)"
            @delete="emit('delete', $event)"
          />
        </div>
      </div>
    </div>
  </div>
</template>
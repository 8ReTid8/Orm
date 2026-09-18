<script setup lang="ts">
import { MoreVertical, Wallet } from "lucide-vue-next"
import type { Account } from "@/types/account"
import { formatMoney } from "@/utils/format.ts";
import ActionButton from "../common/ActionButton.vue";
import { ref } from "vue";
import ConfirmModal from "../common/ConfirmModal.vue";

interface Props {
  account: Account
}

const props = defineProps<Props>()

const emit = defineEmits<{
  detail: [accountId: number]
  edit: [account: Account]     // 👈 เพิ่ม emit edit
  delete: [accountId: number]  // 👈 เพิ่ม emit de
}>()

const isDeleteModalOpen = ref(false)
function handleConfirmDelete() {
  emit("delete", props.account.id)
  isDeleteModalOpen.value = false
}

</script>

<template>
  <!-- <div class="card border border-base-300 bg-base-100 shadow-sm transition hover:-translate-y-1 hover:shadow-md"> -->
  <div
    class="card min-h-52 border border-base-300 bg-base-100 shadow-sm transition hover:-translate-y-1 hover:shadow-md">

    <div class="card-body">

      <!-- Account -->
      <div class="flex items-center gap-3">
        <div class="flex size-12 shrink-0 items-center justify-center rounded-xl bg-lime-300 text-green-700">
          <Wallet class="size-6" />
        </div>

        <div>
          <p class="text-base text-base-content/50">
            บัญชี
          </p>

          <h2 class="text-lg font-semibold">
            {{ account.name }}
          </h2>
        </div>
      </div>

      <!-- Balance -->
      <div class="mt-5 text-right">
        <p class="text-sm text-base-content/50">
          ยอดคงเหลือ
        </p>

        <p class="mt-1 text-2xl font-bold">
          ฿{{ formatMoney(account.balance) }}
        </p>
      </div>

      <!-- Footer -->
      <div class="mt-5 flex items-center justify-between border-t border-base-300 pt-4">
        <button type="button" class="link link-primary text-sm font-medium" @click="emit('detail', account.id)">
          ดูรายละเอียด
        </button>

        <!-- <button
        type="button"
        class="btn btn-square btn-ghost btn-sm"
      >
        <MoreVertical class="size-5" />
      </button> -->
        <ActionButton edit-title="แก้ไขบัญชี" delete-title="ลบบัญชี" @edit="emit('edit', account)"
          @delete="isDeleteModalOpen = true" />
      </div>

    </div>
  </div>
  <ConfirmModal :open="isDeleteModalOpen" title="ยืนยันการลบบัญชี?"
    description="หากลบบัญชีนี้ ข้อมูลที่เกี่ยวข้องกับบัญชีนี้อาจได้รับผลกระทบ" confirm-text="ลบบัญชี" type="danger"
    @close="isDeleteModalOpen = false" @confirm="handleConfirmDelete" />
</template>
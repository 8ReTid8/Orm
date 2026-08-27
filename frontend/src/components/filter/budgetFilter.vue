<script setup lang="ts">
import type { Account } from "@/types/account"

interface Props {
  startDate: string
  endDate: string
  accounts: Account[]
  selectedAccountId: number | null
  status: string
}

defineProps<Props>()

const emit = defineEmits<{
  "update:startDate": [value: string]
  "update:endDate": [value: string]
  "update:selectedAccountId": [value: number | null]
  "update:status": [value: string]
}>()
</script>

<template>
  <div class="rounded-xl border border-base-300 bg-base-100 p-4">

    <div class="mb-3">
      <h2 class="font-semibold">
        ตัวกรอง
      </h2>

      <p class="text-sm text-base-content/50">
        เลือกช่วงเวลาและบัญชีที่ต้องการดู
      </p>
    </div>

    <div class="flex flex-wrap items-end gap-3">

      <!-- Start Date -->
      <div class="form-control">
        <label class="label px-0 pb-1">
          <span class="label-text text-xs">
            ตั้งแต่
          </span>
        </label>

        <input
          type="date"
          :value="startDate"
          class="input input-bordered input-sm"
          @input="
            emit(
              'update:startDate',
              ($event.target as HTMLInputElement).value
            )
          "
        />
      </div>

      <!-- End Date -->
      <div class="form-control">
        <label class="label px-0 pb-1">
          <span class="label-text text-xs">
            ถึง
          </span>
        </label>

        <input
          type="date"
          :value="endDate"
          class="input input-bordered input-sm"
          @input="
            emit(
              'update:endDate',
              ($event.target as HTMLInputElement).value
            )
          "
        />
      </div>

      <!-- Account -->
      <div class="form-control">
        <label class="label px-0 pb-1">
          <span class="label-text text-xs">
            บัญชี
          </span>
        </label>

        <select
          :value="selectedAccountId ?? ''"
          class="select select-bordered select-sm"
          @change="
            emit(
              'update:selectedAccountId',
              ($event.target as HTMLSelectElement).value
                ? Number(($event.target as HTMLSelectElement).value)
                : null
            )
          "
        >
          <option value="">
            ทุกบัญชี
          </option>

          <option
            v-for="account in accounts"
            :key="account.id"
            :value="account.id"
          >
            {{ account.name }}
          </option>
        </select>
      </div>

      <!-- Status -->
      <div class="form-control">
        <label class="label px-0 pb-1">
          <span class="label-text text-xs">
            สถานะ
          </span>
        </label>

        <select
          :value="status"
          class="select select-bordered select-sm"
          @change="
            emit(
              'update:status',
              ($event.target as HTMLSelectElement).value
            )
          "
        >
          <option value="">
            ทุกสถานะ
          </option>

          <option value="active">
            กำลังใช้งาน
          </option>

          <option value="upcoming">
            กำลังจะเริ่ม
          </option>

          <option value="completed">
            สิ้นสุดแล้ว
          </option>
        </select>
      </div>

    </div>
  </div>
</template>
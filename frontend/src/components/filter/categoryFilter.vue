<!-- <script setup lang="ts">
import type { Category } from "@/types/category";

interface Props {
  categories: Category[]
  modelValue: string | null
}

defineProps<Props>()

const emit = defineEmits<{
  "update:modelValue": [value: string | null]
}>()
function handleChange(event: Event) {
  const select = event.currentTarget as HTMLSelectElement

  emit(
    "update:modelValue",
    // select.value ? Number(select.value) : null
    select.value ? select.value : null
  )
}
</script>

<template>
  <div class=" flex items-center gap-2">
    <label class="shrink-0 text-sm font-medium">
      หมวดหมู่
    </label>

    <select
      :value="modelValue ?? ''"
      class="select select-bordered"
      @change="handleChange"
    >
      <option value="">
        ทุกหมวดหมู่
      </option>

      <option
        v-for="category in categories"
        :key="category.id"
        :value="category.name"
      >
        {{ category.name }}
      </option>
    </select>
  </div>
</template> -->

<script setup lang="ts">
import { ref, computed } from "vue";
import type { Category } from "@/types/category";
import { resolveCategoryIcon } from "@/utils/categoryIcons";
import { Tag, X, ChevronDown, Check } from "lucide-vue-next";

interface Props {
  categories: Category[]
  modelValue: string | null
}

const props = defineProps<Props>()

const emit = defineEmits<{
  "update:modelValue": [value: string | null]
}>()

const dialogRef = ref<HTMLDialogElement | null>(null)

// แยกหมวดหมู่ตามประเภท
const incomeCategories = computed(() =>
  props.categories.filter((c) => c.type === "income")
)

const expenseCategories = computed(() =>
  props.categories.filter((c) => c.type === "expense")
)

// หา Category ที่เลือกอยู่ปัจจุบันเพื่อเอา Icon มาโชว์ที่ปุ่มด้านนอก
const selectedCategory = computed(() =>
  props.categories.find((c) => c.name === props.modelValue)
)

function openModal() {
  dialogRef.value?.showModal()
}

function closeModal() {
  dialogRef.value?.close()
}

function selectCategory(categoryName: string | null) {
  emit("update:modelValue", categoryName)
  closeModal()
}
</script>

<template>
  <div>
    <!-- ปุ่มกดเปิด Modal หน้าเว็บ -->
    <div class="flex items-center gap-2">
      <label class="shrink-0 text-sm font-medium">หมวดหมู่</label>
      <!-- <div class="flex flex-col gap-1">
      
      <label class="text-xs font-medium text-base-content/70">
        หมวดหมู่
      </label> -->
      <button type="button" class="btn btn-sm btn-outline gap-2 font-normal h-10 self-start" :class="{
        'btn-success text-success hover:text-white': selectedCategory?.type === 'income',
        'btn-error text-error hover:text-white': selectedCategory?.type === 'expense'
      }" @click="openModal">
        <component :is="selectedCategory?.icon ? resolveCategoryIcon(selectedCategory.icon) : Tag" class="size-4" />
        <span>{{ modelValue || "ทุกหมวดหมู่" }}</span>

        <X v-if="modelValue" class="size-3.5 hover:text-error ml-1" @click.stop="selectCategory(null)" />
        <ChevronDown v-else class="size-3.5 text-base-content/50" />
      </button>
    </div>

    <!-- Modal แนวนอน -->
    <dialog ref="dialogRef" class="modal">
      <div class="modal-box max-w-xl">
        <div class="flex items-center justify-between pb-3 mb-4">
          <h3 class="font-bold text-lg">เลือกหมวดหมู่</h3>

          <div class="flex items-center gap-2">
            <button type="button" class="btn btn-sm" :class="!modelValue ? 'btn-neutral' : 'btn-ghost'"
              @click="selectCategory(null)">
              ทุกหมวดหมู่
            </button>
            <form method="dialog">
              <button class="btn btn-sm btn-circle btn-ghost">✕</button>
            </form>
          </div>
        </div>

        <!-- แถวบน: รายรับ / แถวล่าง: รายจ่าย -->
        <div class="space-y-4! my-2!">
          <!-- รายรับ -->
          <div>
            <div class="flex items-center gap-2 font-semibold text-success text-sm mb-2.5!">
              <span class="badge badge-success badge-xs"></span>
              รายรับ
            </div>
            <div class="flex flex-wrap gap-2">
              <button v-for="cat in incomeCategories" :key="cat.id" type="button"
                class="btn btn-base font-normal gap-2 rounded-full transition-all"
                :class="modelValue === cat.name ? 'btn-success text-white shadow-sm' : 'btn-ghost bg-base-200/80 hover:bg-base-300'"
                @click="selectCategory(cat.name)">
                <component :is="resolveCategoryIcon(cat.icon)" class="size-4 shrink-0" />
                <span>{{ cat.name }}</span>
                <Check v-if="modelValue === cat.name" class="size-3.5" />
              </button>
            </div>
          </div>

          <div class="divider my-1"></div>

          <!-- รายจ่าย -->
          <div>
            <div class="flex items-center gap-2 font-semibold text-error text-sm mb-2.5!">
              <span class="badge badge-error badge-xs"></span>
              รายจ่าย
            </div>
            <div class="flex flex-wrap gap-2">
              <button v-for="cat in expenseCategories" :key="cat.id" type="button"
                class="btn btn-base font-normal gap-2 rounded-full transition-all"
                :class="modelValue === cat.name ? 'btn-error text-white shadow-sm' : 'btn-ghost bg-base-200/80 hover:bg-base-300'"
                @click="selectCategory(cat.name)">
                <component :is="resolveCategoryIcon(cat.icon)" class="size-4 shrink-0" />
                <span>{{ cat.name }}</span>
                <Check v-if="modelValue === cat.name" class="size-3.5" />
              </button>
            </div>
          </div>
        </div>
      </div>

      <form method="dialog" class="modal-backdrop">
        <button>close</button>
      </form>
    </dialog>
  </div>
</template>
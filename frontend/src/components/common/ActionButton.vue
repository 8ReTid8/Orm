<script setup lang="ts">
import { Pencil, Trash2 } from "lucide-vue-next"

interface Props {
    editTitle?: string
    deleteTitle?: string
    size?: "xs" | "sm" | "md"
    disabled?: boolean
    showEdit?: boolean
    showDelete?: boolean
}

withDefaults(defineProps<Props>(), {
    editTitle: "แก้ไข",
    deleteTitle: "ลบ",
    size: "sm",
    disabled: false,
    showEdit: true,
    showDelete: true,
})

const emit = defineEmits<{
    edit: []
    delete: []
}>()
</script>

<template>
    <div class="flex items-center gap-1">
        <!-- ปุ่ม Edit -->
        <button v-if="showEdit" type="button"
            class="btn btn-ghost btn-square text-base-content/60 hover:bg-base-300 hover:text-base-content" :class="{
                'btn-xs': size === 'xs',
                'btn-xs sm:btn-sm': size === 'sm',
                'btn-md': size === 'md',
            }" :title="editTitle" :disabled="disabled" @click.stop="emit('edit')">
            <Pencil class="size-4" />
        </button>

        <!-- ปุ่ม Delete -->
        <button v-if="showDelete" type="button"
            class="btn btn-ghost btn-square text-base-content/60 hover:bg-error/10 hover:text-error" :class="{
                'btn-xs': size === 'xs',
                'btn-xs sm:btn-sm': size === 'sm',
                'btn-md': size === 'md',
            }" :title="deleteTitle" :disabled="disabled" @click.stop="emit('delete')">
            <Trash2 class="size-4" />
        </button>
    </div>
</template>
<script setup lang="ts">
interface Props {
    total: number
}
defineProps<Props>()
</script>
<template>
    <span class="badge badge-neutral font-medium">
        {{ total }} รายการ
    </span>
</template>
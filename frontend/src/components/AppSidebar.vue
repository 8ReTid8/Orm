<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const isCollapsed = ref(false)

const navItems = [
  {
    label: 'ปฏิทิน',
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"/><line x1="16" y1="2" x2="16" y2="6"/><line x1="8" y1="2" x2="8" y2="6"/><line x1="3" y1="10" x2="21" y2="10"/></svg>`,
    to: '/',
  },
  {
    label: 'saving goal',
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="7" height="7"/><rect x="14" y="3" width="7" height="7"/><rect x="14" y="14" width="7" height="7"/><rect x="3" y="14" width="7" height="7"/></svg>`,
    to: '/dashboard',
  },
  {
    label: 'รายรับ-รายจ่าย',
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="1" x2="12" y2="23"/><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/></svg>`,
    to: '/transactions',
  },
  {
    label: 'งบประมาณ',
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/></svg>`,
    to: '/budget',
  },
  {
    label: 'รายงาน',
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="20" x2="18" y2="10"/><line x1="12" y1="20" x2="12" y2="4"/><line x1="6" y1="20" x2="6" y2="14"/></svg>`,
    to: '/reports',
  },
  {
    label: 'การลงทุน',
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="22 7 13.5 15.5 8.5 10.5 2 17"/><polyline points="16 7 22 7 22 13"/></svg>`,
    to: '/investments',
  },
]

const bottomItems = [
  {
    label: 'ตั้งค่า',
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83-2.83l.06-.06A1.65 1.65 0 0 0 4.68 15a1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 2.83-2.83l.06.06A1.65 1.65 0 0 0 9 4.68a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 2.83l-.06.06A1.65 1.65 0 0 0 19.4 9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"/></svg>`,
    to: '/settings',
  },
]

function isActive(to: string) {
  if (to === '/') return route.path === '/'
  return route.path.startsWith(to)
}
</script>

<template>
  <aside
    :class="[
      'flex flex-col min-h-screen bg-gradient-to-b from-[#1a1d2e] to-[#141624]',
      'border-r border-white/[0.06] flex-shrink-0 relative z-50',
      'transition-all duration-300 ease-in-out overflow-hidden',
      isCollapsed ? 'w-[68px]' : 'w-[260px]',
    ]"
  >
    <!-- Logo / Brand -->
    <div class="flex items-center justify-between px-4 pt-5 pb-4 border-b border-white/[0.06]">
      <div class="flex items-center gap-3 overflow-hidden">
        <div
          class="w-9 h-9 bg-gradient-to-br from-[#6c63ff] to-[#00d4aa] rounded-[10px] flex items-center justify-center text-white flex-shrink-0 shadow-[0_4px_12px_rgba(108,99,255,0.4)]"
        >
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="20" height="20">
            <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1.41 16.09V20h-2.67v-1.93c-1.71-.36-3.16-1.46-3.27-3.4h1.96c.1 1.05.82 1.87 2.65 1.87 1.96 0 2.4-.98 2.4-1.59 0-.83-.44-1.61-2.67-2.14-2.48-.6-4.18-1.62-4.18-3.67 0-1.72 1.39-2.84 3.11-3.21V4h2.67v1.95c1.86.45 2.79 1.86 2.85 3.39H14.3c-.05-1.11-.64-1.87-2.22-1.87-1.5 0-2.4.68-2.4 1.64 0 .84.65 1.39 2.67 1.91s4.18 1.39 4.18 3.91c-.01 1.83-1.38 2.83-3.12 3.16z"/>
          </svg>
        </div>
        <div v-if="!isCollapsed" class="flex flex-col overflow-hidden">
          <span class="text-[15px] font-bold whitespace-nowrap bg-gradient-to-r from-white to-indigo-300 bg-clip-text text-transparent">FinanceORM</span>
          <span class="text-[10px] text-slate-500 whitespace-nowrap">บริหารการเงินส่วนตัว</span>
        </div>
      </div>
      <button
        class="text-slate-500 hover:text-white hover:bg-white/[0.08] p-1 rounded-md flex items-center justify-center flex-shrink-0 transition-all duration-200 cursor-pointer border-none bg-transparent"
        @click="isCollapsed = !isCollapsed"
        :title="isCollapsed ? 'ขยาย' : 'ย่อ'"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none"
          stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
          class="w-4 h-4 transition-transform duration-300"
          :style="{ transform: isCollapsed ? 'rotate(180deg)' : 'rotate(0deg)' }"
        >
          <polyline points="15 18 9 12 15 6"/>
        </svg>
      </button>
    </div>

    <!-- Balance Card -->
    <div
      v-if="!isCollapsed"
      class="mx-3 my-4 p-4 bg-gradient-to-br from-[rgba(108,99,255,0.15)] to-[rgba(0,212,170,0.08)] border border-[rgba(108,99,255,0.2)] rounded-[14px] backdrop-blur-md"
    >
      <div class="text-[11px] text-slate-400 uppercase tracking-[0.05em] mb-1.5">ยอดเงินคงเหลือ</div>
      <div class="text-[22px] font-bold text-white leading-none mb-2">
        ฿ 124,580<span class="text-sm text-slate-400 font-normal">.00</span>
      </div>
      <div class="text-xs font-medium flex items-center gap-1 text-emerald-500">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" width="12" height="12">
          <polyline points="18 15 12 9 6 15"/>
        </svg>
        +8.2% เดือนนี้
      </div>
    </div>

    <!-- Navigation -->
    <nav class="flex-1 px-2.5 py-2 overflow-y-auto">
      <div v-if="!isCollapsed" class="text-[10px] font-semibold text-slate-600 uppercase tracking-[0.08em] px-2 py-2">
        เมนูหลัก
      </div>

      <router-link
        v-for="item in navItems"
        :key="item.to"
        :to="item.to"
        :title="isCollapsed ? item.label : ''"
        :class="[
          'flex items-center gap-3 px-2.5 py-2.5 rounded-[10px] no-underline text-sm font-medium',
          'transition-all duration-200 relative mb-0.5 overflow-hidden whitespace-nowrap',
          isActive(item.to)
            ? 'bg-gradient-to-r from-[rgba(108,99,255,0.2)] to-[rgba(108,99,255,0.05)] text-indigo-300 border border-[rgba(108,99,255,0.2)]'
            : 'text-slate-400 hover:bg-white/[0.06] hover:text-slate-200 border border-transparent',
        ]"
      >
        <span class="nav-icon flex items-center justify-center w-5 h-5 flex-shrink-0" v-html="item.icon" />
        <span v-if="!isCollapsed" class="flex-1">{{ item.label }}</span>
        <span
          v-if="isActive(item.to) && !isCollapsed"
          class="w-1.5 h-1.5 rounded-full bg-[#6c63ff] shadow-[0_0_8px_#6c63ff] flex-shrink-0"
        />
      </router-link>
    </nav>

    <!-- Bottom -->
    <div class="px-2.5 pb-3 pt-2 border-t border-white/[0.06]">
      <router-link
        v-for="item in bottomItems"
        :key="item.to"
        :to="item.to"
        :title="isCollapsed ? item.label : ''"
        :class="[
          'flex items-center gap-3 px-2.5 py-2.5 rounded-[10px] no-underline text-sm font-medium',
          'transition-all duration-200 mb-0.5 overflow-hidden whitespace-nowrap',
          isActive(item.to)
            ? 'bg-gradient-to-r from-[rgba(108,99,255,0.2)] to-[rgba(108,99,255,0.05)] text-indigo-300 border border-[rgba(108,99,255,0.2)]'
            : 'text-slate-400 hover:bg-white/[0.06] hover:text-slate-200 border border-transparent',
        ]"
      >
        <span class="nav-icon flex items-center justify-center w-5 h-5 flex-shrink-0" v-html="item.icon" />
        <span v-if="!isCollapsed" class="flex-1">{{ item.label }}</span>
      </router-link>

      <!-- User Profile -->
      <div
        :class="[
          'flex items-center gap-2.5 p-2.5 rounded-[10px] mt-2 bg-white/[0.04] overflow-hidden',
          isCollapsed ? 'justify-center' : '',
        ]"
      >
        <div class="w-8 h-8 rounded-full bg-gradient-to-br from-[#6c63ff] to-[#00d4aa] flex items-center justify-center text-sm font-bold text-white flex-shrink-0">
          ก
        </div>
        <div v-if="!isCollapsed" class="flex-1 min-w-0">
          <div class="text-[13px] font-semibold text-slate-200 whitespace-nowrap overflow-hidden text-ellipsis">กัลยา สมิต</div>
          <div class="text-[11px] text-slate-500 whitespace-nowrap overflow-hidden text-ellipsis">kalaya@example.com</div>
        </div>
        <button
          v-if="!isCollapsed"
          title="ออกจากระบบ"
          class="text-slate-500 hover:text-rose-400 hover:bg-rose-500/10 p-1 rounded-md flex items-center justify-center transition-all duration-200 flex-shrink-0 cursor-pointer border-none bg-transparent"
        >
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" width="16" height="16">
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
            <polyline points="16 17 21 12 16 7"/>
            <line x1="21" y1="12" x2="9" y2="12"/>
          </svg>
        </button>
      </div>
    </div>
  </aside>
</template>

<style scoped>
/* :deep() selector required for v-html SVG icon sizing */
.nav-icon :deep(svg) {
  width: 18px;
  height: 18px;
}
</style>

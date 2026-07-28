<script setup lang="ts">
import { Menu } from "lucide-vue-next"
import AppSidebar from "@/components/AppSidebar.vue"
</script>

<template>
  <div class="drawer lg:drawer-open">
    <input id="main-sidebar" type="checkbox" class="drawer-toggle" />

    <!-- เนื้อหาหลัก -->
    <div class="drawer-content flex min-h-screen flex-col bg-base-100">
      <header class="navbar sticky top-0 z-20 border-b border-base-300 bg-base-100 px-4">
        <!-- ปุ่มเปิด Sidebar บนมือถือ -->
        <div class="flex-none lg:hidden">
          <label for="main-sidebar" aria-label="Open sidebar" class="btn btn-square btn-ghost">
            <Menu class="size-5" />
          </label>
        </div>

        <div class="flex-1">
          <span class="text-lg font-semibold">Finance Dashboard</span>
        </div>

        <div class="flex-none">
          <div class="avatar placeholder">
            <div class="w-9 rounded-full bg-neutral text-neutral-content">
              <span class="text-sm">ND</span>
            </div>
          </div>
        </div>
      </header>

      <main class="flex-1 p-4 md:p-6">
        <RouterView />
      </main>
    </div>

    <!-- Sidebar -->
    <div class="drawer-side z-30">
      <label for="main-sidebar" aria-label="Close sidebar" class="drawer-overlay" />

      <AppSidebar />
    </div>
  </div>
</template>
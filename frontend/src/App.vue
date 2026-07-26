<script setup lang="ts">
import AppSidebar from './components/AppSidebar.vue'
</script>

<template>
  <div id="app-layout">
    <AppSidebar />
    <RouterView />
  </div>
</template>

<style scoped>
#app-layout {
  display: flex;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}
</style>

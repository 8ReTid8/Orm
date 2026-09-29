<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useRoute } from "vue-router"
import axios from "axios"
import { CircleCheck, CircleX, MailCheck } from "lucide-vue-next"
import { verifyEmail } from "@/services/auth"

const route = useRoute()

const isLoading = ref(true)
const isSuccess = ref(false)
const message = ref("กำลังยืนยันอีเมล...")

onMounted(async () => {
  const token = route.query.token

  if (typeof token !== "string" || token === "") {
    isLoading.value = false
    message.value = "ไม่พบ token สำหรับยืนยันอีเมล"
    return
  }

  try {
    const result = await verifyEmail(token)

    isSuccess.value = true
    message.value = result.message
  } catch (error: unknown) {
    if (axios.isAxiosError(error)) {
      message.value =
        error.response?.data?.message ??
        "ไม่สามารถยืนยันอีเมลได้"
    } else {
      message.value = "เกิดข้อผิดพลาด"
    }
  } finally {
    isLoading.value = false
  }
})
</script>

<template>
  <main class="flex min-h-screen items-center justify-center bg-base-200 px-4">
    <div class="card w-full max-w-md bg-base-100 shadow-2xl">
      <div class="card-body items-center p-10 text-center">
        <span
          v-if="isLoading"
          class="loading loading-spinner loading-lg text-primary"
        />

        <CircleCheck
          v-else-if="isSuccess"
          class="size-16 text-success"
        />

        <CircleX
          v-else
          class="size-16 text-error"
        />

        <h1 class="mt-4 text-2xl font-bold">
          {{
            isLoading
              ? "กำลังยืนยันอีเมล"
              : isSuccess
                ? "ยืนยันอีเมลสำเร็จ"
                : "ยืนยันอีเมลไม่สำเร็จ"
          }}
        </h1>

        <p class="mt-2 text-base-content/60">
          {{ message }}
        </p>

        <RouterLink
          v-if="!isLoading"
          to="/login"
          class="btn btn-primary mt-6"
        >
          <MailCheck class="size-5" />
          ไปหน้าเข้าสู่ระบบ
        </RouterLink>
      </div>
    </div>
  </main>
</template>
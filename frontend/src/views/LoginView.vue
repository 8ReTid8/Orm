<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"
import axios from "axios"
import { LogIn } from "lucide-vue-next"

import { login } from "@/services/auth"

const router = useRouter()

const email = ref("")
const password = ref("")

const errorMessage = ref("")
const isLoading = ref(false)

async function submitLogin() {
    errorMessage.value = ""

    try {
        isLoading.value = true

        const result = await login({
            email: email.value,
            password: password.value,
        })

        localStorage.setItem("token", result.token)
        localStorage.setItem("user", JSON.stringify(result.user))

        await router.push("/")
    } catch (error: unknown) {
        if (axios.isAxiosError(error)) {
            errorMessage.value =
                error.response?.data?.message ?? "เข้าสู่ระบบไม่สำเร็จ"
        } else {
            errorMessage.value = "เกิดข้อผิดพลาด"
        }
    } finally {
        isLoading.value = false
    }
}
</script>

<template>
    <main class="flex min-h-screen items-center justify-center bg-base-200 p-4">
        <div class="card w-full max-w-md bg-base-100 shadow-xl">
            <div class="card-body">
                <div class="mb-3 flex justify-center">
                    <div class="rounded-full bg-primary p-3 text-primary-content">
                        <LogIn class="size-7" />
                    </div>
                </div>

                <h1 class="text-center text-2xl font-bold">
                    เข้าสู่ระบบ
                </h1>

                <p class="text-center text-sm text-base-content/60">
                    เข้าสู่ระบบเพื่อจัดการข้อมูลการเงิน
                </p>

                <div v-if="errorMessage" role="alert" class="alert alert-error mt-4">
                    <span>{{ errorMessage }}</span>
                </div>

                <form class="mt-4 space-y-4" @submit.prevent="submitLogin">
                    <label class="form-control">
                        <span class="label-text mb-2">อีเมล</span>
                        <input v-model.trim="email" type="email" class="input input-bordered w-full"
                            autocomplete="email" required />
                    </label>

                    <label class="form-control">
                        <span class="label-text mb-2">รหัสผ่าน</span>
                        <input v-model="password" type="password" class="input input-bordered w-full"
                            autocomplete="current-password" required />
                    </label>

                    <button type="submit" class="btn btn-primary w-full" :disabled="isLoading">
                        <span v-if="isLoading" class="loading loading-spinner loading-sm" />

                        {{ isLoading ? "กำลังเข้าสู่ระบบ..." : "เข้าสู่ระบบ" }}
                    </button>
                </form>

                <p class="mt-3 text-center text-sm">
                    ยังไม่มีบัญชี?

                    <RouterLink to="/register" class="link link-primary">
                        สมัครสมาชิก
                    </RouterLink>
                </p>
            </div>
        </div>
    </main>
</template>
<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"
import axios from "axios"
import { LogIn, Mail, Lock } from "lucide-vue-next"
import { login, resendVerificationEmail } from "@/services/auth"
import { useAuthStore } from "@/stores/auth"

const router = useRouter()
const authStore = useAuthStore()
const email = ref("")
const password = ref("")
const errorMessage = ref("")
const isLoading = ref(false)

const canResend = ref(false)
const isResending = ref(false)
const resendSuccessMessage = ref("")

// async function submitLogin() {
//     errorMessage.value = ""

//     try {
//         isLoading.value = true

//         const result = await login({
//             email: email.value,
//             password: password.value,
//         })

//         // authStore.setAuth(result.token, result.user)
//         authStore.setAuth(result.accessToken, result.refreshToken, result.user)
//         await router.push("/")
//     } catch (error: unknown) {
//         if (axios.isAxiosError(error)) {
//             errorMessage.value =
//                 error.response?.data?.message ?? "เข้าสู่ระบบไม่สำเร็จ"
//         } else {
//             errorMessage.value = "เกิดข้อผิดพลาด"
//         }
//     } finally {
//         isLoading.value = false
//     }
// }

async function submitLogin() {
    errorMessage.value = ""
    resendSuccessMessage.value = ""
    canResend.value = false
    try {
        isLoading.value = true
        const result = await login({ email: email.value, password: password.value })
        authStore.setAuth(result.accessToken, result.refreshToken, result.user)
        await router.push("/")
    } catch (error: unknown) {
        if (axios.isAxiosError(error)) {
            errorMessage.value = error.response?.data?.message ?? "เข้าสู่ระบบไม่สำเร็จ"
            // 👈 ถ้าเจอ status 403 (ยังไม่ยืนยันอีเมล) ให้เปิดปุ่ม Resend
            if (error.response?.status === 403) {
                canResend.value = true
            }
        } else {
            errorMessage.value = "เกิดข้อผิดพลาด"
        }
    } finally {
        isLoading.value = false
    }
}
// ฟังก์ชันกดส่งอีเมลใหม่
async function handleResend() {
    if (!email.value) return
    try {
        isResending.value = true
        errorMessage.value = ""
        const res = await resendVerificationEmail(email.value)
        resendSuccessMessage.value = res.message
        canResend.value = false
    } catch (err: any) {
        errorMessage.value = err.response?.data?.message ?? "ส่งอีเมลไม่สำเร็จ"
    } finally {
        isResending.value = false
    }
}
</script>

<template>
    <main class="flex min-h-screen items-center justify-center bg-base-200 px-4">
        <div class="card w-full max-w-lg bg-base-100 shadow-2xl">
            <div class="card-body p-10">
                <!-- <div class="mb-3 flex justify-center">
                    <div class="rounded-full bg-[#99e550] p-3 ">
                        <LogIn class="size-7" />
                    </div>
                </div> -->
                <div class="flex justify-center mb-4">
                    <div
                        class="w-20 h-20 rounded-full bg-[#99e550] text-success-content flex items-center justify-center shadow-lg">
                        <LogIn class="w-10 h-10" />
                    </div>
                </div>

                <div class="text-center mb-8">
                    <h1 class="text-3xl font-bold">
                        เข้าสู่ระบบ
                    </h1>

                    <p class="text-base-content/60 mt-2">
                        เข้าสู่ระบบเพื่อจัดการข้อมูลการเงิน
                    </p>
                </div>


                <!-- <div v-if="errorMessage" role="alert" class="alert alert-error mt-4">
                    <span>{{ errorMessage }}</span>
                </div> -->
                <div v-if="resendSuccessMessage" class="alert alert-success text-sm mb-4">
                    {{ resendSuccessMessage }}
                </div>
                <!-- กล่องแจ้งเตือน Error + ปุ่มกดส่งเมลใหม่ -->
                <div v-if="errorMessage" class="alert alert-error text-sm mb-4 flex flex-col items-start gap-2">
                    <div>{{ errorMessage }}</div>

                    <!-- ปุ่มกดส่งอีเมลใหม่ จะโผล่ขึ้นมาเฉพาะตอน 403 (ยังไม่ยืนยันอีเมล) -->
                    <button v-if="canResend" type="button"
                        class="btn btn-xs btn-outline border-white text-white hover:bg-white hover:text-error"
                        :disabled="isResending" @click="handleResend">
                        {{ isResending ? "กำลังส่งอีเมล..." : "ส่งลิงก์ยืนยันใหม่อีกครั้ง" }}
                    </button>
                </div>
                <!-- <form class="mt-4 space-y-4" @submit.prevent="submitLogin"> -->
                <form class="flex flex-col gap-3" @submit.prevent="submitLogin">
                    <!-- <label class="form-control">
                        <span class="label-text mb-2">อีเมล</span>
                        <input v-model.trim="email" type="email" class="input input-bordered w-full"
                            autocomplete="email" required />
                    </label>

                    <label class="form-control">
                        <span class="label-text mb-2">รหัสผ่าน</span>
                        <input v-model="password" type="password" class="input input-bordered w-full"
                            autocomplete="current-password" required />
                    </label> -->

                    <div class="flex flex-col gap-1">
                        <label for="email" class="text-sm font-medium">
                            อีเมล
                        </label>

                        <label class="input input-bordered flex w-full items-center gap-3">
                            <Mail class="size-5 shrink-0 text-[#99e550]" />

                            <input id="email" v-model.trim="email" type="email" class="grow"
                                placeholder="example@email.com" autocomplete="email" required />
                        </label>
                    </div>

                    <div class="flex flex-col gap-1">
                        <label for="password" class="text-sm font-medium">
                            รหัสผ่าน
                        </label>

                        <label class="input input-bordered flex w-full items-center gap-3">
                            <Lock class="size-5 shrink-0 text-[#99e550]" />

                            <input id="password" v-model="password" type="password" class="grow" placeholder="********"
                                autocomplete="new-password" minlength="8" required />
                        </label>
                    </div>

                    <button type="submit" class="btn bg-[#99e550] hover:bg-[#6abe30] w-full h-12 text-base"
                        :disabled="isLoading">
                        <span v-if="isLoading" class="loading loading-spinner loading-sm" />

                        {{ isLoading ? "กำลังเข้าสู่ระบบ..." : "เข้าสู่ระบบ" }}
                    </button>
                </form>

                <p class="mt-3 text-center text-sm">
                    ยังไม่มีบัญชี?

                    <RouterLink to="/register" class="link link-success">
                        สมัครสมาชิก
                    </RouterLink>
                </p>
            </div>
        </div>
    </main>
</template>
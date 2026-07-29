<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"
import axios from "axios"
import { UserPlus } from "lucide-vue-next"

import { register } from "@/services/auth"

const router = useRouter()

const email = ref("")
const password = ref("")
const confirmPassword = ref("")

const errorMessage = ref("")
const isLoading = ref(false)

async function submitRegister() {
    errorMessage.value = ""

    if (password.value !== confirmPassword.value) {
        errorMessage.value = "รหัสผ่านทั้งสองช่องไม่ตรงกัน"
        return
    }

    if (password.value.length < 8) {
        errorMessage.value = "รหัสผ่านต้องมีอย่างน้อย 8 ตัวอักษร"
        return
    }

    try {
        isLoading.value = true

        const result = await register({
            email: email.value,
            password: password.value,
        })

        localStorage.setItem("token", result.token)
        localStorage.setItem("user", JSON.stringify(result.user))

        await router.push("/")
    } catch (error: unknown) {
        if (axios.isAxiosError(error)) {
            errorMessage.value =
                error.response?.data?.message ?? "สมัครสมาชิกไม่สำเร็จ"
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
                        <UserPlus class="size-7" />
                    </div>
                </div>

                <h1 class="text-center text-2xl font-bold">
                    สมัครสมาชิก
                </h1>

                <p class="text-center text-sm text-base-content/60">
                    สร้างบัญชีเพื่อเริ่มบันทึกรายรับรายจ่าย
                </p>

                <div v-if="errorMessage" role="alert" class="alert alert-error mt-4">
                    <span>{{ errorMessage }}</span>
                </div>

                <form class="mt-4 space-y-4" @submit.prevent="submitRegister">
                    <!-- <label class="form-control">
                        <span class="label-text mb-2">ชื่อ</span>
                        <input v-model.trim="name" type="text" class="input input-bordered w-full" autocomplete="name"
                            required />
                    </label> -->

                    <label class="form-control">
                        <span class="label-text mb-2">อีเมล</span>
                        <input v-model.trim="email" type="email" class="input input-bordered w-full"
                            autocomplete="email" required />
                    </label>

                    <label class="form-control">
                        <span class="label-text mb-2">รหัสผ่าน</span>
                        <input v-model="password" type="password" class="input input-bordered w-full"
                            autocomplete="new-password" minlength="8" required />
                    </label>

                    <label class="form-control">
                        <span class="label-text mb-2">ยืนยันรหัสผ่าน</span>
                        <input v-model="confirmPassword" type="password" class="input input-bordered w-full"
                            autocomplete="new-password" minlength="8" required />
                    </label>

                    <button type="submit" class="btn btn-primary w-full" :disabled="isLoading">
                        <span v-if="isLoading" class="loading loading-spinner loading-sm" />
                        {{ isLoading ? "กำลังสมัคร..." : "สมัครสมาชิก" }}
                    </button>
                </form>

                <p class="mt-3 text-center text-sm">
                    มีบัญชีแล้ว?

                    <RouterLink to="/login" class="link link-primary">
                        เข้าสู่ระบบ
                    </RouterLink>
                </p>
            </div>
        </div>
    </main>
</template>
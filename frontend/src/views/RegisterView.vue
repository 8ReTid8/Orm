<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"
import axios from "axios"
import { UserPlus, Lock, Mail, ShieldCheck } from "lucide-vue-next"
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

<!-- <template>
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
</template> -->
<template>
    <main class="min-h-screen bg-base-200 flex items-center justify-center px-4">
        <div class="card w-full max-w-lg bg-base-100 shadow-2xl">
            <div class="card-body p-10">

                <div class="flex justify-center mb-4">
                    <div
                        class="w-20 h-20 rounded-full bg-success text-success-content flex items-center justify-center shadow-lg">
                        <UserPlus class="w-10 h-10" />
                    </div>
                </div>


                <div class="text-center mb-8">
                    <h1 class="text-3xl font-bold">
                        สมัครสมาชิก
                    </h1>

                    <p class="text-base-content/60 mt-2">
                        สร้างบัญชีเพื่อเริ่มจัดการการเงินของคุณ
                    </p>
                </div>


                <div v-if="errorMessage" class="alert alert-error mb-4">
                    {{ errorMessage }}
                </div>

                <form class="space-y-5" @submit.prevent="submitRegister">

                    <label class="form-control">
                        <span class="label-text font-medium mb-2">
                            อีเมล
                        </span>
                        <label class="input input-bordered w-full flex items-center gap-3">
                            <Mail class="w-5 h-5 text-success" />
                            <input v-model.trim="email" type="email" class="grow" placeholder="example@email.com"
                                autocomplete="email" />
                        </label>
                    </label>


                    <label class="form-control">
                        <span class="label-text font-medium mb-2">
                            รหัสผ่าน
                        </span>

                        <label class="input input-bordered flex w-full items-center gap-3">
                            <Lock class="w-5 h-5 text-success" />
                            <input v-model="password" type="password" class="grow" placeholder="********" />
                        </label>
                    </label>


                    <label class="form-control">
                        <span class="label-text font-medium mb-2">
                            ยืนยันรหัสผ่าน
                        </span>
                        <label class="input input-bordered flex w-full items-center gap-3">
                            <ShieldCheck class="w-5 h-5 text-success" />
                            <input v-model="confirmPassword" type="password" class="grow" placeholder="********" />
                        </label>
                    </label>


                    <button class="btn btn-success w-full h-12 text-base" :disabled="isLoading">
                        <span v-if="isLoading" class="loading loading-spinner" />
                        {{ isLoading ? "กำลังสมัคร..." : "สมัครสมาชิก" }}
                    </button>

                </form>


                <div class="divider my-6">หรือ</div>
                <p class="text-center text-sm">
                    มีบัญชีอยู่แล้ว ?
                    <RouterLink to="/login" class="link link-success font-semibold">
                        เข้าสู่ระบบ
                    </RouterLink>
                </p>
            </div>
        </div>
    </main>
</template>
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

        // localStorage.setItem("token", result.token)
        // localStorage.setItem("user", JSON.stringify(result.user))

        // await router.push("/")
        await router.push({
            name: "login",
            query: {
                registered: "true",
            },
        })
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
    <main class="min-h-screen bg-base-200 flex items-center justify-center px-4">
        <div class="card w-full max-w-lg bg-base-100 shadow-2xl">
            <div class="card-body p-10">

                <div class="flex justify-center mb-4">
                    <div
                        class="w-20 h-20 rounded-full bg-[#99e550] text-success-content flex items-center justify-center shadow-lg">
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

                <form class="flex flex-col gap-3" @submit.prevent="submitRegister">
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

                    <!-- Confirm password -->
                    <div class="flex flex-col gap-1">
                        <label for="confirmPassword" class="text-sm font-medium">
                            ยืนยันรหัสผ่าน
                        </label>

                        <label class="input input-bordered flex w-full items-center gap-3">
                            <ShieldCheck class="size-5 shrink-0 text-[#99e550]" />

                            <input id="confirmPassword" v-model="confirmPassword" type="password" class="grow"
                                placeholder="********" autocomplete="new-password" minlength="8" required />
                        </label>
                    </div>

                    <button class="btn bg-[#99e550] hover:bg-[#6abe30] w-full h-12 text-base" :disabled="isLoading">
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
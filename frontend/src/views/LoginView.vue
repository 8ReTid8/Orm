<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"
import axios from "axios"
import { LogIn,Mail,Lock } from "lucide-vue-next"
import { login } from "@/services/auth"
import { useAuthStore } from "@/stores/auth"

const router = useRouter()
const authStore = useAuthStore()
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

        // localStorage.setItem("token", result.token)
        // localStorage.setItem("user", JSON.stringify(result.user))
        authStore.setAuth(result.token, result.user)
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


                <div v-if="errorMessage" role="alert" class="alert alert-error mt-4">
                    <span>{{ errorMessage }}</span>
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

                    <button type="submit" class="btn bg-[#99e550] hover:bg-[#6abe30] w-full h-12 text-base" :disabled="isLoading">
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
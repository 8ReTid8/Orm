<script setup lang="ts">
import logo from "@/assets/logo.svg"
import { useRouter } from "vue-router"
import { useAuthStore } from "@/stores/auth";
import {
    ArrowLeftRight,
    WalletCards,
    PiggyBank,
    Tags,
    User,
    LogOut,
    CalendarDays,
    WalletMinimal,
} from "lucide-vue-next"
const authStore = useAuthStore()
const router = useRouter()
const menuItems = [
    {
        name: "Log Money",
        path: "/",
        icon: CalendarDays,
    },
    {
        name: "Accounts",
        path: "/accounts",
        icon: WalletMinimal,
    },
    {
        name: "Budgets",
        path: "/budgets",
        icon: WalletCards,
    },
    {
        name: "Saving Goals",
        path: "/saving-goals",
        icon: PiggyBank,
    },
    {
        name: "Categories",
        path: "/categories",
        icon: Tags,
    },
]
async function logout() {
    authStore.clearAuth()
    await router.push("/login")
}
</script>

<template>
    <aside class="flex min-h-full w-64 flex-col bg-base-200">
        <!-- Logo -->
        <div class="flex h-20 items-center gap-3 border-b border-base-300 px-5 bg-white">
            <div>
                <!-- <h1 class="font-bold">ORM</h1> -->
                <img :src="logo" alt="ORM Logo" class="h-10 w-auto" />
                <p class="text-xs text-base-content/60">Money management</p>
            </div>
        </div>

        <!-- Menu -->
        <nav class="flex-1 p-3">


            <ul class="menu w-full gap-2 p-0 ">
                <li v-for="item in menuItems" :key="item.path">
                    <RouterLink :to="item.path" 
                        exact-active-class="menu-active bg-[#99e550] text-white"
                        class="hover:bg-[#99e550] hover:text-white">
                        <component :is="item.icon" class="size-6" />
                        <span class="text-base te">{{ item.name }}</span>
                    </RouterLink>
                </li>
            </ul>
        </nav>

        <!-- Footer -->
        <div class="border-t border-base-300 p-3">
            <ul class="menu w-full gap-1 p-0" v-if="authStore.isAuthenticated">
                <li>
                    <RouterLink to="/profile" active-class="menu-active">
                        <User class="size-6" />
                        <span class="text-base ">Profile</span>
                    </RouterLink>
                </li>

                <li>
                    <button type="button" @click="logout">
                        <LogOut class="size-6" />
                        <span class="text-base">Logout</span>
                    </button>
                </li>
            </ul>
            <ul v-else class="menu w-full gap-1 p-0">
                <li>
                    <RouterLink to="/login" active-class="menu-active" class="flex items-center gap-3">
                        <LogIn class="size-6 shrink-0" />
                        <span class="text-base leading-none">Login</span>
                    </RouterLink>
                </li>

                <li>
                    <RouterLink to="/register" active-class="menu-active" class="flex items-center gap-3">
                        <UserPlus class="size-6 shrink-0" />
                        <span class="text-base leading-none">Register</span>
                    </RouterLink>
                </li>
            </ul>
        </div>
    </aside>
</template>
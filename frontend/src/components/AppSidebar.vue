<script setup lang="ts">
import {
    LayoutDashboard,
    ArrowLeftRight,
    WalletCards,
    PiggyBank,
    Tags,
    User,
    LogOut,
    CalendarDays,
} from "lucide-vue-next"

const menuItems = [
    {
        name: "Log Money",
        path: "/",
        icon: CalendarDays,
    },
    {
        name: "Transactions",
        path: "/transactions",
        icon: ArrowLeftRight,
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
</script>

<template>
    <aside class="flex min-h-full w-64 flex-col bg-base-200">
        <!-- Logo -->
        <div class="flex h-16 items-center gap-3 border-b border-base-300 px-5">
            <!-- <div class="flex size-9 items-center justify-center rounded-lg bg-primary text-primary-content">
                <WalletCards class="size-5" />
            </div> -->

            <div>
                <h1 class="font-bold">ORM</h1>
                <p class="text-xs text-base-content/60">Money management</p>
            </div>
        </div>

        <!-- Menu -->
        <nav class="flex-1 p-3">
            

            <ul class="menu w-full gap-1 p-0">
                <li v-for="item in menuItems" :key="item.path">
                    <RouterLink :to="item.path" active-class="menu-active"
                        :exact-active-class="item.path === '/' ? 'menu-active' : ''">
                        <component :is="item.icon" class="size-5" />
                        <span>{{ item.name }}</span>
                    </RouterLink>
                </li>
            </ul>
        </nav>

        <!-- Footer -->
        <div class="border-t border-base-300 p-3">
            <ul class="menu w-full gap-1 p-0">
                <li>
                    <RouterLink to="/profile" active-class="menu-active">
                        <User class="size-5" />
                        <span>Profile</span>
                    </RouterLink>
                </li>

                <li>
                    <button type="button">
                        <LogOut class="size-5" />
                        <span>Logout</span>
                    </button>
                </li>
            </ul>
        </div>
    </aside>
</template>
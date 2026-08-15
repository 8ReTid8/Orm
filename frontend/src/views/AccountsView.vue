<script setup lang="ts">
import { onMounted, ref } from "vue"
import {
    Plus,
    Wallet,
    Landmark,
    MoreVertical,
} from "lucide-vue-next"
import type { Account } from "@/types/account"
import AddAccountForm from "@/components/account/AddAccountForm.vue"
import { createAccount, getAccounts } from "@/services/account"


// const accounts = ref<Account[]>([
//     {
//         id: 1,
//         name: "เงินสด",
//         balance: 2500,
//         // transactionCount: 12,
//     },
//     {
//         id: 2,
//         name: "KBank เงินเดือน",
//         balance: 18200,
//         // transactionCount: 31,
//     },
//     {
//         id: 3,
//         name: "SCB เงินออม",
//         balance: 42000,
//         // transactionCount: 8,
//     },
// ])
const accounts = ref<Account[]>([])
const isLoading = ref(false)
const isAddAccountOpen = ref(false)

function formatMoney(value: number) {
    return new Intl.NumberFormat("th-TH", {
        minimumFractionDigits: 2,
        maximumFractionDigits: 2,
    }).format(value)
}

function openAddAccount() {
    isAddAccountOpen.value = true
}

function closeAddAccount() {
    isAddAccountOpen.value = false
}

async function saveAccount(data: {
    name: string
}) {
    try {
        const result = await createAccount({
            name: data.name,
        })

        console.log(result.account)

        isAddAccountOpen.value = false

        // ขั้นต่อไปค่อย reload accounts
    } catch (error) {
        console.error(
            "Create account failed:",
            error,
        )
    }
}
async function loadAccounts() {
    try {
        isLoading.value = true

        accounts.value = await getAccounts()
    } catch (error) {
        console.error(
            "Failed to load accounts:",
            error,
        )
    } finally {
        isLoading.value = false
    }
}
onMounted(() => {
    loadAccounts()
})
</script>

<template>
    <section class="space-y-6">
        <!-- Header -->
        <div class="flex items-center justify-between">
            <div>
                <h1 class="text-2xl font-bold">
                    Accounts
                </h1>

                <p class="mt-1 text-base-content/60">
                    จัดการบัญชีและแหล่งเงินทั้งหมดของคุณ
                </p>
            </div>

            <button type="button" class="btn btn-primary" @click="openAddAccount">
                <Plus class="size-5" />
                เพิ่มบัญชี
            </button>
        </div>

        <!-- Account grid -->
        <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 xl:grid-cols-3">
            <!-- Account cards -->
            <div v-for="account in accounts" :key="account.id"
                class="card border border-base-300 bg-base-100 shadow-sm transition hover:-translate-y-1 hover:shadow-md">
                <div class="card-body">
                    <!-- Top -->
                    <div class="flex items-start justify-between">
                        <div class="flex size-11 items-center justify-center rounded-xl bg-primary/15 text-primary">
                            <Wallet class="size-6" />
                        </div>

                        <button type="button" class="btn btn-square btn-ghost btn-sm">
                            <MoreVertical class="size-5" />
                        </button>
                    </div>

                    <!-- Account name -->
                    <div class="mt-4">
                        <p class="text-sm text-base-content/50">
                            บัญชี
                        </p>

                        <h2 class="mt-1 text-lg font-semibold">
                            {{ account.name }}
                        </h2>
                    </div>

                    <!-- Balance -->
                    <div class="mt-5">
                        <p class="text-sm text-base-content/50">
                            ยอดคงเหลือ
                        </p>

                        <p class="mt-1 text-2xl font-bold">
                            ฿{{ formatMoney(account.balance) }}
                        </p>
                    </div>

                    <!-- Footer -->
                    <div
                        class="mt-5 flex items-center justify-between border-t border-base-300 pt-4 text-sm text-base-content/60">
                        <!-- <span>
                            {{ account.transactionCount }} รายการ
                        </span> -->

                        <button type="button" class="link link-primary font-medium">
                            ดูรายละเอียด
                        </button>
                    </div>
                </div>
            </div>
            <!-- <div v-for="account in accounts" :key="account.id"
                class="card border border-base-300 bg-base-100 shadow-sm">
                <div class="card-body">
                    <h2 class="text-lg font-semibold">
                        {{ account.name }}
                    </h2>

                    <div class="mt-4">
                        <p class="text-sm text-base-content/50">
                            ยอดคงเหลือ
                        </p>

                        <p class="text-2xl font-bold">
                            ฿{{ account.balance.toLocaleString() }}
                        </p>
                    </div>
                </div>
            </div> -->
            <!-- Add account card -->
            <button type="button"
                class="flex min-h-64 flex-col items-center justify-center rounded-2xl border border-dashed border-base-300 bg-base-100 text-base-content/50 transition hover:border-primary hover:bg-primary/5 hover:text-primary"
                @click="openAddAccount">
                <div class="flex size-12 items-center justify-center rounded-full bg-base-200">
                    <Plus class="size-6" />
                </div>

                <p class="mt-3 font-semibold">
                    เพิ่มบัญชีใหม่
                </p>

                <p class="mt-1 text-sm">
                    เงินสด ธนาคาร หรือ Wallet
                </p>
            </button>
        </div>
    </section>
    <AddAccountForm :open="isAddAccountOpen" @close="closeAddAccount" @save="saveAccount" />
</template>
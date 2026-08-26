<script setup lang="ts">
import { onMounted, ref } from "vue"
import {
    Plus,
} from "lucide-vue-next"
import AddAccountForm from "@/components/account/AddAccountForm.vue"
import { createAccount, getAccounts } from "@/services/account"
import AccountCard from "@/components/account/AccountCard.vue"
import { useAccountStore } from "@/stores/account"

// const accounts = ref<Account[]>([])
const accountStore = useAccountStore()
const isLoading = ref(false)
const isAddAccountOpen = ref(false)
console.log(localStorage.getItem("token"))
function openAddAccount() {
    isAddAccountOpen.value = true
}

function openAccountDetail(accountId: number) {
    console.log("account:", accountId)
}

function closeAddAccount() {
    isAddAccountOpen.value = false
}

async function saveAccount(data: {
    name: string
}) {
    try {
        const result = await createAccount(data)

        console.log(result.account)
        accountStore.addAccount(result.account)
        isAddAccountOpen.value = false

        // ขั้นต่อไปค่อย reload accounts
    } catch (error) {
        console.error(
            "Create account failed:",
            error,
        )
    }
}

onMounted(() => {
    accountStore.loadAccounts(true)
})
</script>

<template>
    <section class="!space-y-6">
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
        </div>

        <!-- Account grid -->
        <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 xl:grid-cols-3">

            <!-- <div v-for="account in accounts" :key="account.id"
                class="card border border-base-300 bg-base-100 shadow-sm transition hover:-translate-y-1 hover:shadow-md">
                <div class="card-body">
                 
                    <div class="flex items-start justify-between">
                        <div class="flex size-11 items-center justify-center rounded-xl bg-primary/15 text-primary">
                            <Wallet class="size-6" />
                        </div>

                        <button type="button" class="btn btn-square btn-ghost btn-sm">
                            <MoreVertical class="size-5" />
                        </button>
                    </div>

                 
                    <div class="mt-4">
                        <p class="text-sm text-base-content/50">
                            บัญชี
                        </p>

                        <h2 class="mt-1 text-lg font-semibold">
                            {{ account.name }}
                        </h2>
                    </div>

                    <div class="mt-5">
                        <p class="text-sm text-base-content/50">
                            ยอดคงเหลือ
                        </p>

                        <p class="mt-1 text-2xl font-bold">
                            ฿{{ formatMoney(account.balance) }}
                        </p>
                    </div>

                   
                    <div
                        class="mt-5 flex items-center justify-between border-t border-base-300 pt-4 text-sm text-base-content/60">
                        <span>
                            {{ account.transactionCount }} รายการ
                        </span>

                        <button type="button" class="link link-primary font-medium">
                            ดูรายละเอียด
                        </button>
                    </div>
                </div>
            </div> -->

            <AccountCard v-for="account in accountStore.accounts" :key="account.id" :account="account" @detail="openAccountDetail" />

            <!-- Add account card -->
            <button type="button"
                class="flex min-h-52 flex-col items-center justify-center rounded-2xl border border-dashed border-base-300 bg-base-100 text-base-content/50 transition hover:border-lime-600 hover:bg-lime-600/5 hover:text-lime-600"
                @click="openAddAccount">
                <div class="flex size-12 items-center justify-center rounded-full bg-base-200">
                    <Plus class="size-6" />
                </div>

                <p class="mt-3 font-semibold">
                    เพิ่มบัญชีใหม่
                </p>

            </button>
        </div>
    </section>
    <AddAccountForm :open="isAddAccountOpen" @close="closeAddAccount" @save="saveAccount" />
</template>
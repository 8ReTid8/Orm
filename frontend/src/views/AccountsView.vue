<script setup lang="ts">
import { onMounted, ref } from "vue"
import { Plus } from "lucide-vue-next"
import AddAccountForm from "@/components/account/AddAccountForm.vue"
import { createAccount } from "@/services/account"
import AccountCard from "@/components/account/AccountCard.vue"
import { useAccountStore } from "@/stores/account"
import { useAccount } from "@/composables/account/useAccount"

// const accountStore = useAccountStore()
// const isAddAccountOpen = ref(false)
// function openAddAccount() {
//     isAddAccountOpen.value = true
// }

// function openAccountDetail(accountId: number) {
//     console.log("account:", accountId)
// }

// function closeAddAccount() {
//     isAddAccountOpen.value = false
// }

// async function saveAccount(data: {
//     name: string
// }) {
//     try {
//         const result = await createAccount(data)

//         console.log(result.account)
//         accountStore.addAccount(result.account)
//         isAddAccountOpen.value = false

//         // ขั้นต่อไปค่อย reload accounts
//     } catch (error) {
//         console.error(
//             "Create account failed:",
//             error,
//         )
//     }
// }
// onMounted(() => {
//     accountStore.loadAccounts(true)
// })

const {
    accounts,
    isAddAccountOpen,
    openAddAccount,
    closeAddAccount,
    openAccountDetail,
    saveAccount,
} = useAccount()

</script>

<template>
    <section class="space-y-6!">
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

            <AccountCard v-for="account in accounts" :key="account.id" :account="account"
                @detail="openAccountDetail" />

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
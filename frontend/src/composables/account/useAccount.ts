import { onMounted, ref } from "vue"
import { storeToRefs } from "pinia"
import { useAccountStore } from "@/stores/account"
import { createAccount } from "@/services/account"

export function useAccount() {
  const accountStore = useAccountStore()

  const { accounts } = storeToRefs(accountStore)

  const isAddAccountOpen = ref(false)

  function openAddAccount() {
    isAddAccountOpen.value = true
  }

  function closeAddAccount() {
    isAddAccountOpen.value = false
  }

  function openAccountDetail(accountId: number) {
    console.log("account:", accountId)
  }

  async function saveAccount(data: { name: string }) {
    try {
      const result = await createAccount(data)

      accountStore.addAccount(result.account)

      closeAddAccount()
    } catch (error) {
      console.error("Create account failed:", error)
    }
  }

  onMounted(() => {
    accountStore.loadAccounts(true)
  })

  return {
    accounts,
    isAddAccountOpen,
    openAddAccount,
    closeAddAccount,
    openAccountDetail,
    saveAccount,
  }
}
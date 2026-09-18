import { onMounted, ref } from "vue"
import { storeToRefs } from "pinia"
import { useAccountStore } from "@/stores/account"
import { 
  createAccount,
  deleteAccount as deleteAccountApi
 } from "@/services/account"

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
  async function deleteAccount(id: number) {
    try {
      await deleteAccountApi(id)
      // accountStore.removeAccount(id)
      accountStore.loadAccounts(true)
    } catch (error) {
      console.error("Delete account failed:", error)
      throw error
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
    deleteAccount,
  }
}
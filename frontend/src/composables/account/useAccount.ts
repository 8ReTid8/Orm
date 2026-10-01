import { onMounted, ref } from "vue"
import { useRouter } from "vue-router"
import { storeToRefs } from "pinia"
import { useAccountStore } from "@/stores/account"
import {
  createAccount,
  deleteAccount as deleteAccountApi
} from "@/services/account"

export function useAccount() {
  const router = useRouter()
  const accountStore = useAccountStore()

  const { accounts, isLoading } = storeToRefs(accountStore)

  const isAddAccountOpen = ref(false)
  const isSubmitting = ref(false)
  const error = ref<string | null>(null)

  function openAddAccount() {
    isAddAccountOpen.value = true
  }

  function closeAddAccount() {
    isAddAccountOpen.value = false
  }

  function openAccountDetail(accountId: number) {
    router.push({ name: "account-detail", params: { id: accountId } })
  }

  async function saveAccount(data: { name: string }) {
    if (!data.name || !data.name.trim()) {
      const validationError = new Error("กรุณากรอกชื่อบัญชี")
      console.error("Validation failed:", validationError.message)
      throw validationError
    }

    try {
      isSubmitting.value = true
      error.value = null

      const result = await createAccount({ name: data.name.trim() })

      accountStore.addAccount(result.account)
      closeAddAccount()
      return result.account
    } catch (err) {
      console.error("Create account failed:", err)
      error.value = "ไม่สามารถเพิ่มบัญชีได้"
      throw err
    } finally {
      isSubmitting.value = false
    }
  }

  async function deleteAccount(id: number) {
    try {
      error.value = null
      await deleteAccountApi(id)
      await accountStore.loadAccounts(true)
    } catch (err) {
      console.error("Delete account failed:", err)
      error.value = "ไม่สามารถลบบัญชีได้"
      throw err
    }
  }

  onMounted(() => {
    accountStore.loadAccounts(true)
  })

  return {
    accounts,
    isLoading,
    isSubmitting,
    isAddAccountOpen,
    error,
    openAddAccount,
    closeAddAccount,
    openAccountDetail,
    saveAccount,
    deleteAccount,
  }
}
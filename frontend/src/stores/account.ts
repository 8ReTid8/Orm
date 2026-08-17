import { ref } from "vue"
import { defineStore } from "pinia"

import { getAccounts } from "@/services/account"
import type { Account } from "@/types/account"

export const useAccountStore = defineStore("account", () => {
  const accounts = ref<Account[]>([])
  const isLoading = ref(false)
  const isLoaded = ref(false)

  async function loadAccounts(force = false) {
    if (isLoaded.value && !force) {
      return
    }

    try {
      isLoading.value = true

      accounts.value = await getAccounts()
      isLoaded.value = true
    } finally {
      isLoading.value = false
    }
  }

  function addAccount(account: Account) {
    accounts.value.push(account)
  }

  return {
    accounts,
    isLoading,
    isLoaded,
    loadAccounts,
    addAccount,
  }
})
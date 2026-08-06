import type { AuthUser } from "@/types/auth"
import { defineStore } from "pinia"
import { computed, ref } from "vue"

function loadUser(): AuthUser | null {
    const savedUser = localStorage.getItem("user")

    if (!savedUser) {
        return null
    }

    try {
        return JSON.parse(savedUser) as AuthUser
    } catch {
        localStorage.removeItem("user")
        return null
    }
}

export const useAuthStore = defineStore("auth", () => {
    const token = ref<string | null>(localStorage.getItem("token"))
    const user = ref<AuthUser | null>(loadUser())

    const isAuthenticated = computed(() => Boolean(token.value))

    function setAuth(newToken: string, newUser: AuthUser) {
        token.value = newToken
        user.value = newUser

        localStorage.setItem("token", newToken)
        localStorage.setItem("user", JSON.stringify(newUser))
    }

    function clearAuth() {
        token.value = null
        user.value = null

        localStorage.removeItem("token")
        localStorage.removeItem("user")
    }

    return {
        token,
        user,
        isAuthenticated,
        setAuth,
        clearAuth,
    }
})
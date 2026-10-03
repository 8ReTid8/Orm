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
    // const token = ref<string | null>(localStorage.getItem("token"))
    const accessToken = ref<string | null>(localStorage.getItem("accessToken"))
    const refreshToken = ref<string | null>(localStorage.getItem("refreshToken"))
    const user = ref<AuthUser | null>(loadUser())

    // const isAuthenticated = computed(() => Boolean(token.value))
    const isAuthenticated = computed(() => Boolean(accessToken.value))

    // function setAuth(newToken: string, newUser: AuthUser) {
    //     token.value = newToken
    //     user.value = newUser

    //     localStorage.setItem("token", newToken)
    //     localStorage.setItem("user", JSON.stringify(newUser))
    // }
    function setAuth(newAccessToken: string, newRefreshToken: string, newUser: AuthUser) {
        accessToken.value = newAccessToken
        refreshToken.value = newRefreshToken
        user.value = newUser
        localStorage.setItem("accessToken", newAccessToken)
        localStorage.setItem("refreshToken", newRefreshToken)
        localStorage.setItem("user", JSON.stringify(newUser))
    }
    function setAccessToken(newAccessToken: string) {
        accessToken.value = newAccessToken
        localStorage.setItem("accessToken", newAccessToken)
    }
    // function clearAuth() {
    //     token.value = null
    //     user.value = null

    //     localStorage.removeItem("token")
    //     localStorage.removeItem("user")
    // }
    function clearAuth() {
        accessToken.value = null
        refreshToken.value = null
        user.value = null
        localStorage.removeItem("accessToken")
        localStorage.removeItem("refreshToken")
        localStorage.removeItem("user")
    }

    return {
        accessToken,
        refreshToken,
        user,
        isAuthenticated,
        setAuth,
        clearAuth,
        setAccessToken,
    }
})
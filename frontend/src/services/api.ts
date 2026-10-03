import { useAuthStore } from "@/stores/auth"
import axios from "axios"
import { refreshAccessToken } from "./auth"

const api = axios.create({
  baseURL: "http://localhost:8080/api",
  // baseURL: "http://192.168.1.102:8080/api",
  // headers: {
  //   "Content-Type": "application/json",
  // },
})

// api.interceptors.request.use((config) => {
//   const token = localStorage.getItem("token")

//   if (token) {
//     config.headers.Authorization = `Bearer ${token}`
//   }

//   return config
// })

// Request: แนบ access token ทุก request
api.interceptors.request.use((config) => {
  const authStore = useAuthStore()
  const token = authStore.accessToken
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Response: ถ้าได้ 401 → ลอง refresh แล้ว retry อัตโนมัติ
let isRefreshing = false
let failedQueue: Array<{
  resolve: (value: string) => void
  reject: (reason: unknown) => void
}> = []
function processQueue(error: unknown, token: string | null = null) {
  failedQueue.forEach((p) => {
    if (error) {
      p.reject(error)
    } else {
      p.resolve(token!)
    }
  })
  failedQueue = []
}
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config
    if (error.response?.status !== 401 || originalRequest._retry) {
      return Promise.reject(error)
    }
    
    if (originalRequest.url?.includes('/auth/refresh')) {
      const authStore = useAuthStore()
      authStore.clearAuth()
      window.location.href = '/login'
      return Promise.reject(error)
    }
    
    const authStore = useAuthStore()
    const storedRefreshToken = authStore.refreshToken
    if (!storedRefreshToken) {
      authStore.clearAuth()
      window.location.href = "/login"
      return Promise.reject(error)
    }
    if (isRefreshing) {
      // ถ้ากำลัง refresh อยู่ → รอคิว
      return new Promise((resolve, reject) => {
        failedQueue.push({ resolve, reject })
      }).then((token) => {
        originalRequest.headers.Authorization = `Bearer ${token}`
        return api(originalRequest)
      })
    }
    originalRequest._retry = true
    isRefreshing = true
    try {
      const data = await refreshAccessToken(storedRefreshToken)
      authStore.setAccessToken(data.accessToken)
      processQueue(null, data.accessToken)
      originalRequest.headers.Authorization = `Bearer ${data.accessToken}`
      return api(originalRequest)
    } catch (refreshError) {
      processQueue(refreshError, null)
      authStore.clearAuth()
      window.location.href = "/login"
      return Promise.reject(refreshError)
    } finally {
      isRefreshing = false
    }
  }
)
export default api
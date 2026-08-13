import type { AuthResponse, LoginInput, RegisterInput } from "@/types/auth"
import api from "./api"

export async function register(input: RegisterInput) {
  const response = await api.post<AuthResponse>(
    "/auth/register",
    input,
  )

  return response.data
}

export async function login(input: LoginInput) {
  const response = await api.post<AuthResponse>(
    "/auth/login",
    input,
  )

  return response.data
}
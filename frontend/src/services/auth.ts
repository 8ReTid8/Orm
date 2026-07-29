import api from "./api"

export interface AuthUser {
  id: number
  email: string
}

export interface AuthResponse {
  message: string
  token: string
  user: AuthUser
}

export interface RegisterInput {
  email: string
  password: string
}

export interface LoginInput {
  email: string
  password: string
}

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
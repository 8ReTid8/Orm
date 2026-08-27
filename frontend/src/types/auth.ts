export interface AuthUser {
  id: number
  email: string
  role: "user" | "admin"
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
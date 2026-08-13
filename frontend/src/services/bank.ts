import type { Bank } from "@/types/bank"
import api from "./api"

export async function getBanks() {
  const response = await api.get<Bank[]>(
    "/banks"
  )

  return response.data
}
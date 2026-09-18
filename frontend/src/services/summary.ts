import type { Summary } from "@/types/summary"
import api from "./api"
export async function getSummary(
  year: number,
  month: number | null,
  accountId: number | null,
//   category: string | null
) {
  const response =
    await api.get<Summary>(
      "/summary",
      {
        params: {
          accountId,
          year,
          month,
        },
      },
    )

  return response.data
}

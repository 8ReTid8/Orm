import type { ComparisonResponse, Summary } from "@/types/summary"
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

export async function getComparison(
  year: number,
  month: number | null,
  accountId: number | null,
  category: string | null,
) {
  const response = await api.get<ComparisonResponse>(
    "/summary/comparison",
    {
      params: {
        year,
        month,
        accountId,
        category,
      },
    },
  )

  return response.data
}
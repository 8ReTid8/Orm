import api from "./api"

export type PeriodResource = "budget" | "transaction"

export async function getPeriodYears(resource?: PeriodResource): Promise<number[]> {
  const response = await api.get<{ years: number[] }>("/period/years", {
    params: {
      resource, // 👈 ส่งแนบไปใน URL เป็น ?resource=budget หรือ ?resource=transaction
    },
  })
  return response.data.years
}
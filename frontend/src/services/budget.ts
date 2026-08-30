import type { BudgetForm, CreateBudgetResponse, GetBudgetsResponse } from "@/types/budget";
import api from "./api"
import type { GetTransactionsResponse } from "@/types/transaction";

export async function createBudget(
    form: BudgetForm
) {
    const payload = new FormData()
    payload.append("category", form.category)
    payload.append("amount", String(form.amount))
    payload.append("accountId", String(form.accountId))
    payload.append("startDate", form.startDate)
    payload.append("endDate", form.endDate)

    const response =
        await api.post<CreateBudgetResponse>(
            "/budgets",
            payload
        )
    return response.data
}

export async function getBudgets(
  year: number,
  month: number,
  status: string,
  
) {
  const response =
    await api.get<GetBudgetsResponse>(
      "/budgets",
      {
        params: {
          year,
          month,
          status,
        },
      },
    )

  return response.data.budgets
}

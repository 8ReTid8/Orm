export function formatDate(date: Date): string {
    const year = date.getFullYear()

    const month = String(
        date.getMonth() + 1,
    ).padStart(2, "0")

    const day = String(
        date.getDate(),
    ).padStart(2, "0")

    return `${year}-${month}-${day}`
}

export function formatMoney(
  amount: number,
): string {
  return new Intl.NumberFormat("th-TH", {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  }).format(amount)
}
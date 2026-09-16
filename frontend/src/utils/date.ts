export interface MonthOption {
  value: number
  label: string
  shortLabel: string
}

export const THAI_MONTHS: readonly MonthOption[] = [
  { value: 1, label: "มกราคม", shortLabel: "ม.ค." },
  { value: 2, label: "กุมภาพันธ์", shortLabel: "ก.พ." },
  { value: 3, label: "มีนาคม", shortLabel: "มี.ค." },
  { value: 4, label: "เมษายน", shortLabel: "เม.ย." },
  { value: 5, label: "พฤษภาคม", shortLabel: "พ.ค." },
  { value: 6, label: "มิถุนายน", shortLabel: "มิ.ย." },
  { value: 7, label: "กรกฎาคม", shortLabel: "ก.ค." },
  { value: 8, label: "สิงหาคม", shortLabel: "ส.ค." },
  { value: 9, label: "กันยายน", shortLabel: "ก.ย." },
  { value: 10, label: "ตุลาคม", shortLabel: "ต.ค." },
  { value: 11, label: "พฤศจิกายน", shortLabel: "พ.ย." },
  { value: 12, label: "ธันวาคม", shortLabel: "ธ.ค." },
] as const

export function toBuddhistYear(year: number): number {
  return year + 543
}
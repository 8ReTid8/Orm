import {
  Utensils,
  Car,
  ShoppingBag,
  Film,
  HeartPulse,
  Home,
  Banknote,
  GraduationCap,
  Plane,
  Receipt,
  CircleDollarSign,
  BanknoteArrowDown,
  HandCoins,
  type LucideIcon,
} from "lucide-vue-next"

// Map string ชื่อ icon ที่อยู่ใน DB เข้ากับ Component จริงที่ import มาเฉพาะตัว
export const categoryIconMap: Record<string, LucideIcon> = {
  Utensils,
  Car,
  ShoppingBag,
  Film,
  HeartPulse,
  Home,
  Banknote,
  GraduationCap,
  Plane,
  Receipt,
  CircleDollarSign,
  BanknoteArrowDown,
  HandCoins,
}

// Helper function
export function resolveCategoryIcon(iconName?: string): LucideIcon {
  if (iconName && categoryIconMap[iconName]) {
    return categoryIconMap[iconName]
  }
  return Receipt // default fallback
}
import { createRouter, createWebHistory } from "vue-router"
import { useAuthStore } from "@/stores/auth"

import MainLayOut from "@/layouts/MainLayout.vue"
import LogMoneyView from "@/views/LogMoneyView.vue"
import LoginView from "@/views/LoginView.vue"
import RegisterView from "@/views/RegisterView.vue"
import AccountsView from "@/views/AccountsView.vue"
import BudgetView from "@/views/BudgetView.vue"
import BudgetDetailView from "@/views/BudgetDetailView.vue"
import SummaryView from "@/views/SummaryView.vue"
import AccountDetailView from "@/views/AccountDetailView.vue"
import VerifyEmailView from "@/views/VerifyEmailView.vue"
// import DashboardView from "@/views/DashboardView.vue"
// import TransactionView from "@/views/TransactionView.vue"
// import SavingGoalView from "@/views/SavingGoalView.vue"
// import CategoryView from "@/views/CategoryView.vue"
// import ProfileView from "@/views/ProfileView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),

  routes: [
    {
      path: "/login",
      name: "login",
      component: LoginView,
      meta: {
        guestOnly: true,
      },
    },
    {
      path: "/register",
      name: "register",
      component: RegisterView,
      meta: {
        guestOnly: true,
      },
    },
    {
      path: "/verify-email",
      name: "verify-email",
      component: VerifyEmailView,
      meta: {
        guestOnly: true,
      },
    },
    {
      path: "/",
      component: MainLayOut,
      children: [
        {
          path: "",
          name: "log money",
          component: LogMoneyView,
        },
        {
          path: "accounts",
          name: "accounts",
          component: AccountsView,
        },
        {
          path: "accounts/:id", // 👈 เพิ่ม Route นี้
          name: "account-detail",
          component: AccountDetailView,
        },
        {
          path: "budgets",
          name: "budgets",
          component: BudgetView,
        },
        {
          path: "budgets/:id",
          name: "budget-detail",
          component: BudgetDetailView,
        },
        {
          path: "summary",
          name: "summary",
          component: SummaryView,
        },
        // {
        //   path: "saving-goals",
        //   name: "saving-goals",
        //   component: SavingGoalView,
        // },
        // {
        //   path: "categories",
        //   name: "categories",
        //   component: CategoryView,
        // },
        // {
        //   path: "profile",
        //   name: "profile",
        //   component: ProfileView,
        // },
      ],
    },
  ],
})

router.beforeEach((to, _from, next) => {
  const authStore = useAuthStore()
  const isGuestOnly = to.matched.some((record) => record.meta.guestOnly)

  // 1. ถ้ายังไม่ได้เข้าสู่ระบบ และไม่ใช่หน้าที่เปิดให้บุคคลทั่วไป (Guest) -> บังคับไปหน้า login
  if (!authStore.isAuthenticated && !isGuestOnly) {
    return next({ name: "login" })
  }

  // 2. ถ้าเข้าสู่ระบบแล้ว แต่พยายามจะเปิดหน้า login / register -> พาไปหน้าหลัก
  if (authStore.isAuthenticated && isGuestOnly) {
    return next({ name: "log money" })
  }

  next()
})

export default router
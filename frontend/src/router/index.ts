import { createRouter, createWebHistory } from "vue-router"

import MainLayout from "@/layouts/MainLayout.vue"
import LogMoney from "@/views/LogMoney.vue"
import LoginView from "@/views/LoginView.vue"
import RegisterView from "@/views/RegisterView.vue"
import AccountsView from "@/views/AccountsView.vue"
import BudgetView from "@/views/BudgetView.vue"
import BudgetDetail from "@/views/BudgetDetail.vue"
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
      path: "/",
      component: MainLayout,
      children: [
        {
          path: "",
          name: "log money",
          component: LogMoney,
        },
        {
          path: "accounts",
          name: "accounts",
          component: AccountsView,
        },
        {
          path: "budgets",
          name: "budgets",
          component: BudgetView,
        },
        {
          path: "budgets/:id",
          name: "budget-detail",
          component: BudgetDetail,
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

export default router
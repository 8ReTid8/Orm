// import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/HomeView.vue'

// const router = createRouter({
//   history: createWebHistory(import.meta.env.BASE_URL),
//   routes: [
//     {
//       path: '/',
//       name: 'home',
//       component: HomeView,
//     },
//     {
//       path: '/about',
//       name: 'about',
//       // route level code-splitting
//       // this generates a separate chunk (About.[hash].js) for this route
//       // which is lazy-loaded when the route is visited.
//       component: () => import('../views/AboutView.vue'),
//     },
//   ],
// })

// export default router
import { createRouter, createWebHistory } from "vue-router"

import MainLayout from "@/layouts/MainLayout.vue"
import LogMoney from "@/views/LogMoney.vue"
import LoginView from "@/views/LoginView.vue"
import RegisterView from "@/views/RegisterView.vue"
// import DashboardView from "@/views/DashboardView.vue"
// import TransactionView from "@/views/TransactionView.vue"
// import BudgetView from "@/views/BudgetView.vue"
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
        // {
        //   path: "transactions",
        //   name: "transactions",
        //   component: TransactionView,
        // },
        // {
        //   path: "budgets",
        //   name: "budgets",
        //   component: BudgetView,
        // },
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
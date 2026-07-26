import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('../views/HomeView.vue'), // placeholder
    },
    {
      path: '/transactions',
      name: 'transactions',
      component: () => import('../views/HomeView.vue'), // placeholder
    },
    {
      path: '/budget',
      name: 'budget',
      component: () => import('../views/HomeView.vue'), // placeholder
    },
    {
      path: '/reports',
      name: 'reports',
      component: () => import('../views/HomeView.vue'), // placeholder
    },
    {
      path: '/investments',
      name: 'investments',
      component: () => import('../views/HomeView.vue'), // placeholder
    },
    {
      path: '/settings',
      name: 'settings',
      component: () => import('../views/HomeView.vue'), // placeholder
    },
  ],
})

export default router

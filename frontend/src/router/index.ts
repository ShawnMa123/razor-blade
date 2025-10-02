import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Dashboard',
      component: () => import('@/views/Dashboard.vue')
    },
    {
      path: '/usage-records',
      name: 'UsageRecords',
      component: () => import('@/views/UsageRecords.vue')
    },
    {
      path: '/razors',
      name: 'Razors',
      component: () => import('@/views/Razors.vue')
    },
    {
      path: '/blades',
      name: 'Blades',
      component: () => import('@/views/Blades.vue')
    },
    {
      path: '/statistics',
      name: 'Statistics',
      component: () => import('@/views/Statistics.vue')
    }
  ]
})

export default router
// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
// 导入views中的页面组件
import Login from '@/views/Login.vue';

// 创建路由实例
const router = createRouter({
  history: createWebHistory(),  // 使用 HTML5 历史模式

  routes: [
    { path: '/acct', component: () => import('@/views/Acct.vue') },
    { path: '/acctBank', component: () => import('@/views/AcctBank.vue') },
    { path: '/merchant', component: () => import('@/views/Merchant.vue') },
    { path: '/cust', component: () => import('@/views/Cust.vue') },
    { path: '/bankAccount', component: () => import('@/views/BankAccount.vue') },
    { path: '/invLoc', component: () => import('@/views/InvLoc.vue') },
    { path: '/paymentMethod', component: () => import('@/views/PaymentMethod.vue') },
    { path: '/cat', component: () => import('@/views/Cat.vue') },
    { path: '/brand', component: () => import('@/views/Brand.vue') },
    { path: '/spec', component: () => import('@/views/Spec.vue') },
    { path: '/empl', component: () => import('@/views/Empl.vue') },
    { path: '/prdt', component: () => import('@/views/Prdt.vue') },
    { path: '/unload', component: () => import('@/views/Unload.vue') },
    { path: '/cost', component: () => import('@/views/Cost.vue') },
    { path: '/sale', component: () => import('@/views/Sale.vue') },
    { path: '/buy', component: () => import('@/views/Buy.vue') },
    { path: '/send', component: () => import('@/views/Send.vue') },
    { path: '/purrec', component: () => import('@/views/Purrec.vue') },
    { path: '/shouldIn', component: () => import('@/views/ShouldIn.vue') },
    { path: '/shouldOut', component: () => import('@/views/ShouldOut.vue') },
    { path: '/in', component: () => import('@/views/In.vue') },
    { path: '/out', component: () => import('@/views/Out.vue') },
    {
      path: '/login',  // 显示登录页面
      name: 'Login',
      component: Login,
    },

  ]
});

export default router;

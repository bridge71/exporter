// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
// 导入views中的页面组件
import Login from '@/views/Login.vue';
import Home from '@/views/TestHome.vue';
import Acct from '@/views/Acct.vue';
import AcctBank from '@/views/AcctBank.vue';

// 创建路由实例
const router = createRouter({
  history: createWebHistory(),  // 使用 HTML5 历史模式
  routes: [
    {
      path: '/',
      name: 'acct',
      component: Acct,

    },
    {
      path: '/acct',
      name: 'acct',
      component: Acct,

    },

    {
      path: '/acctBank',
      name: 'acctBank',
      component: AcctBank,

    },
    {
      path: '/login',  // 显示登录页面
      name: 'Login',
      component: Login,
    },
  ]
});

export default router;

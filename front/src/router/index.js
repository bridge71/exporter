// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
// 导入views中的页面组件
import Login from '../views/Login.vue';
import Home from '../views/Home.vue';
import Test from '../views/Test.vue';
import TestHome from '@/views/TestHome.vue';
import Client from '@/views/Client.vue';
import Entity_Information from '@/views/Entity_Information.vue';
import Bank_Information from '@/views/Bank_Information.vue';

// 创建路由实例
const router = createRouter({
  history: createWebHistory(),  // 使用 HTML5 历史模式
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home,
     
    },
    {
      path: '/test',
      name: 'Test',
      component: Test,
    },
    {
      path: '/login',  // 显示登录页面
      name: 'Login',
      component: Login,
    },
    {
      path: '/testHome',
      name: 'TestHome',
      component: TestHome,
      children: [
        {
          path: '/testHome/CL',
          name: 'Client',
          component: Client,
        },
        {
          path: '/testHome/EN_IN',
          name: 'EntityInformation',
          component: Entity_Information,
        },
        {
          path: '/testHome/Bank_IN',
          name: 'BankInformation',
          component: Bank_Information,
        }
      ]
    }
    // 你可以根据需求在这里添加更多的路由
  ]
});

export default router;

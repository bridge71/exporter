// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
// 导入views中的页面组件
import Login from '../views/Login.vue';
import Home from '../views/Home.vue';
import Test from '../views/Test.vue';
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
    // 你可以根据需求在这里添加更多的路由
  ]
});

export default router;

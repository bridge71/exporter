// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
// 导入views中的页面组件
import Login from '../views/Login.vue';  // 登录页面

// 创建路由实例
const router = createRouter({
  history: createWebHistory(),  // 使用 HTML5 历史模式
  routes: [
    {
      path: '/',  // 根路径，显示 Login 组件
      name: 'Home',
      component: Login,  // 默认显示登录页面
    },
    {
      path: '/login',  // 显示登录页面
      name: 'Login',
      component: Login,  // 显示登录页面
    }
    // 你可以根据需求在这里添加更多的路由
  ]
});

export default router;

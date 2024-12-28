<template>
    <div class="menu-container">
      <h2>欢迎来到菜单页面</h2>
      <p>这是登录后的主页。</p>
      <button @click="logout">退出</button>
    </div>
  </template>
  
  <script>
  import { useRouter } from 'vue-router';
  
  export default {
    name: 'menu',
    setup() {
      const router = useRouter();
  
      const logout = () => {
        // 清除 localStorage 中的登录状态
        localStorage.removeItem('isAuthenticated');
        alert('已退出登录');
        // 跳转回登录页
        router.push({ name: 'Login' });
      };
  
      return {
        logout,
      };
    },
  };
  </script>
  
  <style src="../assets/styles/Menu.css"></style>
  
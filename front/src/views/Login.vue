<!-- src/components/Login.vue -->
<template>
    <div class="login-container">
      <h2>登录</h2>
      <form @submit.prevent="handleLogin">
        <!-- 用户名组件 -->
        <Username v-model:username="username" />
        <!-- 密码组件 -->
        <Password v-model:password="password" />
        <button type="submit">登录</button>
      </form>
    </div>
  </template>
  
  <script>
  import Username from './Username.vue';
  import Password from './Password.vue';
  // 引入 Login 组件的样式
  import '../assets/styles/Login.css';
  export default {
    name: 'Login',
    components: {
      Username,
      Password
    },
    data() {
      return {
        username: '',
        password: ''
      };
    },
    methods: {
      handleLogin() {
        // 模拟登录操作
        if (this.username === 'admin' && this.password === 'password') {
          alert('登录成功');
        } else {
          alert('用户名或密码错误');
        }
      }
    }
  };
  </script>
  
  <style src="../assets/styles/Login.css"></style>
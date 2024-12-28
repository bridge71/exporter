<!-- src/views/Login.vue -->
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
import { ref } from 'vue';
import axios from 'axios';
import Username from '../views/Username.vue';
import Password from '../views/Password.vue';
import { useRouter } from 'vue-router';
import '../assets/styles/Login.css';

export default {
  name: 'Login',
  components: {
    Username,
    Password,
  },
  setup() {
    const username = ref('');
    const password = ref('');
    const router = useRouter();

    /*
    const handleLogin = async () => {
      try {
        const response = await axios.post('www.baidu.com', {
          username: username.value,
          password: password.value,
        });

        if (true) { //response.data.success
          // 登录成功，将用户状态存储在 localStorage 中
          localStorage.setItem('isAuthenticated', true);
          router.push({ name: 'Menu' });  // 登录成功后跳转到 Menu 页面
        } else {
          alert('用户名或密码错误');
        }
      } catch (error) {
        console.error('登录请求失败', error);
        alert('登录失败，请稍后再试');
      }
    };
    */
    const handleLogin = () => {
      if(username.value == '1' && password.value== '1')
      {
        router.push({ name: 'menu' });  // 登录成功后跳转到 Menu 页面
      }
    }
    return {
      username,
      password,
      handleLogin,
    };
  },
};
</script>

<style src="../assets/styles/Login.css"></style>

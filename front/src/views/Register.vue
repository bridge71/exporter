<template>
    <div class="register-container">
      <el-card class="register-card">
        <h2 class="register-title">欢迎注册</h2>
        <el-form :model="registerForm" :rules="rules" ref="registerFormRef">
          <!-- 用户名输入框 -->
          <el-form-item prop="username">
            <el-input v-model="registerForm.username" placeholder="请输入用户名" prefix-icon="User" class="custom-input" />
          </el-form-item>
  
          <!-- 邮箱输入框 -->
          <el-form-item prop="email">
            <el-input v-model="registerForm.email" placeholder="请输入邮箱" prefix-icon="Message" class="custom-input" />
          </el-form-item>
  
          <!-- 密码输入框 -->
          <el-form-item prop="password">
            <el-input v-model="registerForm.password" type="password" placeholder="请输入密码" prefix-icon="Lock" show-password
              class="custom-input" />
          </el-form-item>
  
          <!-- 注册按钮 -->
          <el-form-item>
            <el-button type="primary" @click="handleRegister" :loading="loading" class="register-button">
              注册
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted, onUnmounted } from 'vue';
  import { useRouter } from 'vue-router';
  import axios from 'axios';
  import { ElMessage } from 'element-plus';
  
  const router = useRouter();
  
  // 表单数据
  const registerForm = ref({
    username: '',
    email: '',
    password: '',
  });
  
  // 表单验证规则
  const rules = {
    username: [
      { required: true, message: '请输入用户名', trigger: 'blur' },
    ],
    email: [
      { required: true, message: '请输入邮箱', trigger: 'blur' },
      { type: 'email', message: '请输入有效的邮箱地址', trigger: ['blur', 'change'] },
    ],
    password: [
      { required: true, message: '请输入密码', trigger: 'blur' },
      { min: 6, message: '密码长度至少为 6 位', trigger: 'blur' },
    ],
  };
  
  // 注册按钮加载状态
  const loading = ref(false);
  
  // 注册逻辑
  const handleRegister = () => {
    loading.value = true; // 开始加载
    axios
      .post('/register', {
        Username: registerForm.value.username,
        Email: registerForm.value.email,
        Password: registerForm.value.password,
      })
      .then((response) => {
        if (response.status === 200) {
          ElMessage.success('注册成功');
          router.push({ name: 'Login' }); // 跳转到登录页面
        }
      })
      .catch((error) => {
        ElMessage.error(error.response.data.RetMessage || '注册失败'); // 显示错误信息
      })
      .finally(() => {
        loading.value = false; // 结束加载
      });
  };
  
  // 禁止页面滚动
  const disableScroll = () => {
    document.body.style.overflow = 'hidden';
  };
  
  // 恢复页面滚动
  const enableScroll = () => {
    document.body.style.overflow = 'auto';
  };
  
  // 组件挂载时禁止滚动
  onMounted(() => {
    disableScroll();
  });
  
  // 组件卸载时恢复滚动
  onUnmounted(() => {
    enableScroll();
  });
  </script>
  
  <style scoped>
  .register-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    width: 100vw;
    background-image: url('@/assets/a.jpg');
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    margin: 0;
    padding: 0;
    overflow: hidden;
    position: fixed;
    top: 0;
    left: 0;
  }
  
  .register-card {
    width: 400px;
    padding: 30px;
    border-radius: 12px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
    background-color: rgba(0, 50, 0, 0.7);
    backdrop-filter: blur(10px);
    border: 1px solid rgba(255, 255, 255, 0.2);
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
  
  .register-title {
    text-align: center;
    margin-bottom: 24px;
    font-size: 28px;
    color: #fff;
    font-weight: bold;
  }
  
  .custom-input {
    margin-bottom: 16px;
  }
  
  .register-button {
    width: 100%;
    font-size: 16px;
    font-weight: bold;
    background: linear-gradient(135deg, #6a11cb, #2575fc);
    border: none;
    transition: transform 0.2s ease, box-shadow 0.2s ease;
  }
  
  .register-button:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 10px rgba(106, 17, 203, 0.4);
  }
  
  .register-button:active {
    transform: translateY(0);
    box-shadow: 0 2px 5px rgba(106, 17, 203, 0.4);
  }
  </style>
<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 class="login-title">欢迎登录</h2>
      <el-form :model="loginForm" :rules="rules" ref="loginFormRef">
        <!-- 邮箱输入框 -->
        <el-form-item prop="email">
          <el-input v-model="loginForm.email" placeholder="请输入邮箱" prefix-icon="Message" class="custom-input" />
        </el-form-item>

        <!-- 密码输入框 -->
        <el-form-item prop="password">
          <el-input v-model="loginForm.password" type="password" placeholder="请输入密码" prefix-icon="Lock" show-password
            class="custom-input" />
        </el-form-item>

        <!-- 登录按钮 -->
        <el-form-item>
          <el-button type="primary" @click="handleLogin" :loading="loading" class="login-button">
            登录
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
const loginForm = ref({
  email: '',
  password: '',
});

// 表单验证规则
const rules = {
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入有效的邮箱地址', trigger: ['blur', 'change'] },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度至少为 6 位', trigger: 'blur' },
  ],
};

// 登录按钮加载状态
const loading = ref(false);



// 登录逻辑
const handleLogin = () => {
  loading.value = true; // 开始加载

  axios
    .post('/auth', {
      Email: loginForm.value.email,
      Password: loginForm.value.password,
    })
    .then((response) => {
      if (response.status === 200) {
        ElMessage.success('登录成功');
        sessionStorage.setItem('userToken', response.data.User.UserID);
        router.push({ name: 'Acct' }); // 跳转到菜单页面
      }
    })
    .catch((error) => {
      ElMessage.error(error.response.data.RetMessage || '登录失败'); // 显示错误信息
    })
    .finally(() => {
      loading.value = false; // 结束加载
    });
};

// 禁止页面滚动
const disableScroll = () => {
  document.body.style.overflow = 'hidden';
  ElMessage.info('操作开始，等待三秒...');
  setTimeout(() => {
    router.push({ name: 'Acct' }); // 跳转到菜单页面
  }, 3000); // 3000 毫秒 = 3 秒

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

<style>
/* 重置 body 和 html 的默认样式 */
html,
body {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
  overflow: hidden;
  /* 隐藏滚动条 */
}
</style>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  /* 确保宽度铺满 */
  background-image: url('@/assets/a.jpg');
  /* 设置背景图 */
  background-size: cover;
  /* 背景图覆盖整个容器 */
  background-position: center;
  /* 背景图居中 */
  background-repeat: no-repeat;
  /* 背景图不重复 */
  margin: 0;
  /* 去除默认边距 */
  padding: 0;
  /* 去除默认内边距 */
  overflow: hidden;
  /* 隐藏溢出内容 */
  position: fixed;
  /* 固定位置 */
  top: 0;
  /* 从顶部开始 */
  left: 0;
  /* 从左侧开始 */
}

.login-card {
  width: 400px;
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  background-color: rgba(0, 50, 0, 0.7);
  /* 墨绿色透明背景 */
  backdrop-filter: blur(10px);
  /* 毛玻璃效果 */
  border: 1px solid rgba(255, 255, 255, 0.2);
  position: fixed;
  /* 固定位置 */
  top: 50%;
  /* 垂直居中 */
  left: 50%;
  /* 水平居中 */
  transform: translate(-50%, -50%);
  /* 居中偏移 */
}

.login-title {
  text-align: center;
  margin-bottom: 24px;
  font-size: 28px;
  color: #fff;
  /* 白色文字 */
  font-weight: bold;
}

.custom-input {
  margin-bottom: 16px;
}

.login-button {
  width: 100%;
  font-size: 16px;
  font-weight: bold;
  background: linear-gradient(135deg, #6a11cb, #2575fc);
  border: none;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.login-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 10px rgba(106, 17, 203, 0.4);
}

.login-button:active {
  transform: translateY(0);
  box-shadow: 0 2px 5px rgba(106, 17, 203, 0.4);
}

.register-button {
  width: 100%;
  font-size: 16px;
  font-weight: bold;
  background: linear-gradient(135deg, #6a11cb, #2575fc);
  border: none;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  color: #fff;
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

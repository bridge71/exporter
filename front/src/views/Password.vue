<!-- src/components/Password.vue -->
<template>
  <div class="password-input">
    <label for="password">密码:</label>
    <input type="password" id="password" v-model="password" placeholder="请输入密码" required />
  </div>
</template>

<script>
export default {
  name: 'Password',
  data() {
    return {
      password: ''
    };
  },
  watch: {
    password(newValue) {
      this.$emit('update:password', newValue); // 将密码传递给父组件
    }
  }
};
</script>

<style src="../assets/styles/Password.css"></style>

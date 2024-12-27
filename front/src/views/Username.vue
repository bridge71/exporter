<!-- src/components/Username.vue -->
<template>
  <div class="username-input">
    <label for="username">用户名:</label>
    <input type="text" id="username" v-model="username" placeholder="请输入用户名" required />
  </div>
</template>

<script>
export default {
  name: 'Username',
  data() {
    return {
      username: ''
    };
  },
  watch: {
    username(newValue) {
      this.$emit('update:username', newValue); // 将用户名传递给父组件
    }
  }
};
</script>

<style src="../assets/styles/Username.css"></style>

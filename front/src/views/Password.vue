<!-- src/components/Password.vue -->
<template>
  <div class="password-input">
    <label for="password">密码:</label>
    <input type="password" id="password" v-model="password" placeholder="请输入密码" required />
  </div>
</template>

<script>
import { ref, watch } from 'vue';

export default {
  name: 'Password',
  props: {
    modelValue: String, // 对应 v-model:password
  },
  setup(props, { emit }) {
    const password = ref(props.modelValue);

    watch(password, (newValue) => {
      emit('update:password', newValue); // 向父组件传递密码
    });

    return {
      password,
    };
  },
};
</script>

<style src="../assets/styles/Password.css"></style>

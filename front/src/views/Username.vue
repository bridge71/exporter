<!-- src/components/Username.vue -->
<template>
  <div class="username-input">
    <label for="username">用户名:</label>
    <input type="text" id="username" v-model="username" placeholder="请输入用户名" required />
  </div>
</template>

<script>
import { ref, watch } from 'vue';

export default {
  name: 'Username',
  props: {
    modelValue: String, // 对应 v-model:username
  },
  setup(props, { emit }) {
    const username = ref(props.modelValue);

    watch(username, (newValue) => {
      emit('update:username', newValue); // 向父组件传递用户名
    });

    return {
      username,
    };
  },
};
</script>

<style src="../assets/styles/Username.css"></style>

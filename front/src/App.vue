<!-- src/App.vue -->
<template>
  <div id="app">
    <router-view></router-view> 
  </div>
</template>

<script>
// 导入 Login 组件
import Login from './views/Login.vue';
export default {
  name: 'App',
  components: {
    Login
  }
};
</script>

<style>
/* 根组件的样式 */
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

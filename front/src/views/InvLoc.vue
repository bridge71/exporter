<template>
  <div>
    <!-- 侧边栏菜单 -->
    <SideMenu />
    <!-- 主内容区域 -->
    <router-view />
  </div>
</template>

<script setup>
import SideMenu from '../components/SideMenu.vue';
</script>

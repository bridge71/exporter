<template>
  <el-header>
    <h2>{{ headerTitle }}</h2>
    <div>
      搜索：
      <el-input v-model="searchQuery" placeholder="输入要搜索的关键字" style="width: 200px;" />
      <el-button type="primary" @click="toggleMatchMode">
        {{ isExactMatch ? '完全匹配' : '模糊匹配' }}
      </el-button>
      <el-button type="primary" @click="toggleIdMode">
        {{ onlyId ? '只匹配ID' : '全部匹配' }}
      </el-button>
      <el-button @click="navigateToDictionaryManager">数据字典</el-button>
      <el-button type="primary" @click="handleAdd">{{ addButtonText }}</el-button>
    </div>
  </el-header>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';

// 定义 props
const props = defineProps({
  headerTitle: {
    type: String,
    default: '标题',
  },
  addButtonText: {
    type: String,
    default: '新增',
  },
  searchQuery: {
    type: String,
    default: '',
  },
});

// 定义 emits
const emit = defineEmits([
  'update:searchQuery',
  'toggleMatchMode',
  'toggleIdMode',
  'add',
]);

// 匹配模式
const isExactMatch = ref(true);
const onlyId = ref(true);

// 双向绑定 searchQuery
const searchQuery = computed({
  get() {
    return props.searchQuery;
  },
  set(value) {
    emit('update:searchQuery', value);
  },
});

// 切换匹配模式
const toggleMatchMode = () => {
  isExactMatch.value = !isExactMatch.value;
  emit('toggleMatchMode', isExactMatch.value);
};

// 切换 ID 匹配模式
const toggleIdMode = () => {
  onlyId.value = !onlyId.value;
  emit('toggleIdMode', onlyId.value);
};

// 处理新增按钮点击事件
const handleAdd = () => {
  emit('add');
};

// 路由
const router = useRouter();

// 跳转到数据字典管理页面
const navigateToDictionaryManager = () => {
  router.push('/dictionaryManager');
};
</script>

<style scoped>
.el-header {
  background-color: #f5f5f5;
  padding: 0 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

h2 {
  margin: 0;
  font-size: 20px;
  font-weight: bold;
}

.el-input {
  margin-right: 10px;
}

.el-button {
  margin-left: 10px;
}
</style>

<template>
  <div id="dictionary-manager" class="dictionary-manager">
    <el-header class="header">
      <h2>数据字典管理</h2>
      <el-button type="success" plain @click="goBack">返回</el-button>
    </el-header>

    <el-main class="dictionary-container">
      <el-row :gutter="20">
        <el-col :span="8" v-for="(dict, index) in dictionaryData" :key="dict.dictId">
          <el-card class="custom-card">
            <div class="card-header">
              <h4>{{ dict.dictName }} ({{ dict.items.length }})</h4>
              <div class="button-group">
                <el-button size="mini" type="primary" plain @click="editDictionary(dict, index)">编辑</el-button>
                <el-button size="mini" type="danger" plain @click="deleteDictionary(index)">删除</el-button>
              </div>
            </div>
            <div class="card-body">
              <ul class="scrollable-list">
                <li v-for="item in dict.items" :key="item.itemId" class="dictionary-item">
                  <el-tag effect="plain" size="medium" class="dictionary-tag">{{ item.itemName }}</el-tag>
                </li>
              </ul>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </el-main>

    <!-- 编辑数据字典对话框 -->
    <el-dialog v-model="showEditDialog" title="编辑数据字典" width="50%">
      <el-form :model="dictForm" label-width="80px">
        <el-form-item label="字典名称">
          <el-input v-model="dictForm.dictName"></el-input>
        </el-form-item>
        <el-form-item label="字典项">
          <el-button type="success" plain size="small" @click="addItem">+ 添加字典项</el-button>
          <el-table :data="dictForm.items" border style="margin-top: 10px;" size="small">
            <el-table-column prop="itemName" label="字典项名称">
              <template #default="scope">
                <el-input v-model="scope.row.itemName" size="small"></el-input>
              </template>
            </el-table-column>
            <el-table-column label="操作" width="80">
              <template #default="scope">
                <el-button size="mini" type="danger" @click="removeItem(scope.$index)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button size="small" @click="showEditDialog = false">取消</el-button>
        <el-button size="small" type="success" plain @click="saveDictionary">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

interface DictionaryItem {
  itemId: number;
  itemName: string;
  color: string;
}

interface Dictionary {
  dictId: number;
  dictName: string;
  items: DictionaryItem[];
}

// 路由实例
const router = useRouter();

// 示例数据
const dictionaryData = ref<Dictionary[]>([
  {
    dictId: 1,
    dictName: '客户类型',
    items: [
      { itemId: 1, itemName: '国外终端', color: 'blue' },
      { itemId: 2, itemName: '国外贸易商', color: 'yellow' },
      { itemId: 3, itemName: '国内贸易商', color: 'orange' },
      { itemId: 4, itemName: '国内终端', color: 'red' },
      { itemId: 5, itemName: '外贸离岸公司', color: 'pink' },
      { itemId: 6, itemName: '其他分类', color: 'green' },
    ],
  },
]);

const showEditDialog = ref(false);
const dictForm = ref<Dictionary>({
  dictId: 0,
  dictName: '',
  items: [],
});

const currentIndex = ref<number | null>(null);

// 返回按钮
const goBack = () => {
  router.back();
};

// 添加字典项
const addItem = () => {
  dictForm.value.items.push({ itemId: Date.now(), itemName: '', color: 'blue' });
};

// 删除字典项
const removeItem = (index: number) => {
  dictForm.value.items.splice(index, 1);
};

// 保存字典
const saveDictionary = () => {
  if (currentIndex.value !== null) {
    dictionaryData.value[currentIndex.value] = { ...dictForm.value };
  } else {
    dictForm.value.dictId = Date.now();
    dictionaryData.value.push({ ...dictForm.value });
  }
  showEditDialog.value = false;
};

// 编辑字典
const editDictionary = (dict: Dictionary, index: number) => {
  currentIndex.value = index;
  dictForm.value = JSON.parse(JSON.stringify(dict));
  showEditDialog.value = true;
};

// 删除字典
const deleteDictionary = (index: number) => {
  dictionaryData.value.splice(index, 1);
};
</script>

<style scoped>
.dictionary-manager {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  background-color: #f5f5f5;
  padding: 10px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.dictionary-container {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
}

.custom-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 10px;
  margin-bottom: 16px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.1);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.card-body ul {
  list-style: none;
  padding: 0;
  margin: 0;
  max-height: 120px; /* 设置最大高度 */
  overflow-y: auto; /* 超出部分滚动 */
  border-top: 1px solid #ddd;
  padding-top: 10px;
}

.card-body ul li {
  margin-bottom: 10px;
  text-align: center;
}

.dictionary-tag {
  display: flex;
  align-items: center; /* 垂直居中 */
  justify-content: center; /* 水平居中 */
  width: 100%;
  height: 40px; /* 固定高度 */
  font-size: 16px; /* 字体稍大 */
  border: 1px solid #e0e0e0; /* 添加边框 */
  border-radius: 5px;
  background-color: #f9f9f9; /* 背景颜色 */
}
</style>

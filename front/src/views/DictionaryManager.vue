<template>
  <div id="dictionary-manager" class="dictionary-manager">
    <el-header class="header">
      <h2>数据字典管理</h2>
      <el-button type="success" plain @click="goBack">返回</el-button>
    </el-header>

    <el-main class="dictionary-container">
      <el-row :gutter="20">
        <!-- 动态生成多个表格 -->
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
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

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

// 示例数据（增加7-8个表格）
const defaultDictionaryData: Dictionary[] = [
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
  {
    dictId: 2,
    dictName: '供应商类型',
    items: [
      { itemId: 7, itemName: '电子产品', color: 'blue' },
      { itemId: 8, itemName: '机械设备', color: 'yellow' },
      { itemId: 9, itemName: '服装', color: 'orange' },
      { itemId: 10, itemName: '食品', color: 'red' },
      { itemId: 11, itemName: '家居', color: 'pink' },
      { itemId: 12, itemName: '其他', color: 'green' },
    ],
  },
  {
    dictId: 3,
    dictName: '产品类别',
    items: [
      { itemId: 13, itemName: '物流', color: 'blue' },
      { itemId: 14, itemName: '咨询', color: 'yellow' },
      { itemId: 15, itemName: '维修', color: 'orange' },
      { itemId: 16, itemName: '安装', color: 'red' },
      { itemId: 17, itemName: '培训', color: 'pink' },
      { itemId: 18, itemName: '其他', color: 'green' },
    ],
  },
  {
    dictId: 4,
    dictName: '食品添加剂类别',
    items: [
      { itemId: 19, itemName: '制造业', color: 'blue' },
      { itemId: 20, itemName: '零售业', color: 'yellow' },
      { itemId: 21, itemName: '服务业', color: 'orange' },
      { itemId: 22, itemName: '金融业', color: 'red' },
      { itemId: 23, itemName: '教育行业', color: 'pink' },
      { itemId: 24, itemName: '其他', color: 'green' },
    ],
  },
  {
    dictId: 5,
    dictName: '饲料添加剂类别',
    items: [
      { itemId: 25, itemName: '华北', color: 'blue' },
      { itemId: 26, itemName: '华东', color: 'yellow' },
      { itemId: 27, itemName: '华南', color: 'orange' },
      { itemId: 28, itemName: '西南', color: 'red' },
      { itemId: 29, itemName: '西北', color: 'pink' },
      { itemId: 30, itemName: '其他', color: 'green' },
    ],
  },
  {
    dictId: 6,
    dictName: '计量单位',
    items: [
      { itemId: 31, itemName: '中文', color: 'blue' },
      { itemId: 32, itemName: '英文', color: 'yellow' },
      { itemId: 33, itemName: '法文', color: 'orange' },
      { itemId: 34, itemName: '德文', color: 'red' },
      { itemId: 35, itemName: '西班牙文', color: 'pink' },
      { itemId: 36, itemName: '其他', color: 'green' },
    ],
  },
  {
    dictId: 7,
    dictName: '包装类型',
    items: [
      { itemId: 37, itemName: '小学', color: 'blue' },
      { itemId: 38, itemName: '初中', color: 'yellow' },
      { itemId: 39, itemName: '高中', color: 'orange' },
      { itemId: 40, itemName: '本科', color: 'red' },
      { itemId: 41, itemName: '研究生', color: 'pink' },
      { itemId: 42, itemName: '其他', color: 'green' },
    ],
  },
  {
    dictId: 8,
    dictName: '集装箱类型',
    items: [
      { itemId: 43, itemName: '小型', color: 'blue' },
      { itemId: 44, itemName: '中型', color: 'yellow' },
      { itemId: 45, itemName: '大型', color: 'orange' },
    ],
  },
];

// 示例数据
const dictionaryData = ref<Dictionary[]>(defaultDictionaryData);

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
const addItem = async () => {
  const newItem = { itemId: Date.now(), itemName: '', color: 'blue' };
  dictForm.value.items.push(newItem);

  try {
    // 发送新添加的字典项到后端
    await axios.post(`/api/dictionaries/${dictForm.value.dictId}/items`, newItem); // 替换为你的后端 API 地址
  } catch (error) {
    console.error('添加字典项失败:', error);
    // 如果后端未实现，使用默认数据
    // 这里不需要额外操作，因为已经将新项添加到本地数据中
  }
};

// 删除字典项
const removeItem = async (index: number) => {
  const itemId = dictForm.value.items[index].itemId;
  dictForm.value.items.splice(index, 1);

  try {
    // 发送要删除的字典项到后端
    await axios.delete(`/api/dictionaries/${dictForm.value.dictId}/items/${itemId}`); // 替换为你的后端 API 地址
  } catch (error) {
    console.error('删除字典项失败:', error);
    // 如果后端未实现，使用默认数据
    // 这里不需要额外操作，因为已经从本地数据中删除了该项
  }
};

// 保存字典
const saveDictionary = async () => {
  try {
    if (currentIndex.value !== null) {
      // 更新字典
      await axios.put(`/api/dictionaries/${dictForm.value.dictId}`, dictForm.value); // 替换为你的后端 API 地址
      dictionaryData.value[currentIndex.value] = { ...dictForm.value };
    } else {
      // 添加新字典
      const response = await axios.post('/api/dictionaries', dictForm.value); // 替换为你的后端 API 地址
      dictForm.value.dictId = response.data.dictId; // 假设后端返回新的 dictId
      dictionaryData.value.push({ ...dictForm.value });
    }
    showEditDialog.value = false;
  } catch (error) {
    console.error('保存字典失败:', error);
    // 如果后端未实现，使用默认数据
    if (currentIndex.value !== null) {
      dictionaryData.value[currentIndex.value] = { ...dictForm.value };
    } else {
      dictForm.value.dictId = Date.now();
      dictionaryData.value.push({ ...dictForm.value });
    }
    showEditDialog.value = false;
  }
};

// 编辑字典
const editDictionary = async (dict: Dictionary, index: number) => {
  currentIndex.value = index;

  try {
    // 尝试从后端获取最新的字典数据
    const response = await axios.get(`/api/dictionaries/${dict.dictId}`); // 替换为你的后端 API 地址
    dictForm.value = JSON.parse(JSON.stringify(response.data));
  } catch (error) {
    console.error('获取字典数据失败:', error);
    // 如果后端未实现，使用本地数据
    dictForm.value = JSON.parse(JSON.stringify(dict));
  }

  showEditDialog.value = true;
};

// 删除字典
const deleteDictionary = async (index: number) => {
  try {
    await axios.delete(`/api/dictionaries/${dictionaryData.value[index].dictId}`); // 替换为你的后端 API 地址
    dictionaryData.value.splice(index, 1);
  } catch (error) {
    console.error('删除字典失败:', error);
    // 如果后端未实现，使用默认数据
    dictionaryData.value.splice(index, 1);
  }
};

// 获取数据
const fetchData = async () => {
  try {
    const response = await axios.get('/api/dictionaries'); // 替换为你的后端 API 地址
    dictionaryData.value = response.data;
  } catch (error) {
    console.error('获取数据失败:', error);
    // 如果后端未实现，使用默认数据
    dictionaryData.value = defaultDictionaryData;
  }
};

// onMounted(() => {
//   fetchData();
// });



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
  overflow-y: auto; /* 启用垂直滚动 */
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
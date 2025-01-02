<template>
  <el-container>
    <el-aside width="200px">
      <el-menu>
        <el-menu-item index="1" @click="selectDict('MercType')">MercType 数据字典</el-menu-item>
        <!-- 这里可以添加更多的字典项 -->
      </el-menu>
    </el-aside>
    <el-container>
      <el-header>
        <h1>数据字典管理</h1>
      </el-header>
      <el-main>
        <el-table v-if="currentIndex === '1'" :data="currentDict" style="width: 100%" class="dict-table">
          <el-table-column prop="MercType" label="MercType"></el-table-column>
          <el-table-column label="操作">
            <template #default="scope">
              <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
              <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-button v-if="currentIndex === '1'" type="primary" @click="handleAdd">新增</el-button>
        <el-dialog v-model="dialogVisible" :title="dialogTitle" width="30%">
          <el-form :model="form">
            <el-form-item label="MercType">
              <el-input v-model="form.MercType"></el-input>
            </el-form-item>
          </el-form>
          <template #footer>
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="handleSave">保存</el-button>
          </template>
        </el-dialog>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const mercTypes = ref([]);
    const dialogVisible = ref(false);
    const isEdit = ref(false);
    const dialogTitle = ref('');
    const form = ref({
      MercTypeId: '',
      MercType: ''
    });
    const currentDict = ref([]);
    const currentIndex = ref('');

    const fetchMercTypes = async () => {
      const response = await axios.get('/find/mercType');
      mercTypes.value = response.data.MercType;
      currentDict.value = mercTypes.value;
    };

    const handleAdd = () => {
      dialogTitle.value = '新增 MercType';
      isEdit.value = false;
      form.value = { MercTypeId: '', MercType: '' };
      dialogVisible.value = true;
    };

    const handleEdit = (index, row) => {
      dialogTitle.value = '编辑 MercType';
      isEdit.value = true;
      form.value = { ...row };
      dialogVisible.value = true;
    };

    const handleDelete = async (index, row) => {
      await axios.post('/delete/mercType', { MercTypeId: row.MercTypeId });
      fetchMercTypes();
    };

    const handleSave = async () => {
      if (isEdit.value) {
        await axios.post('/save/mercType', form.value);
      } else {
        await axios.post('/save/mercType', { MercType: form.value.MercType });
      }
      dialogVisible.value = false;
      fetchMercTypes();
    };

    const selectDict = (dictName) => {
      currentIndex.value = '1'; // 设置当前 index 为 1
      // 这里可以根据 dictName 来选择显示哪个字典的数据
      // 例如，如果有一个字典列表，可以在这里根据名称来筛选
      currentDict.value = mercTypes.value.filter(item => item.dictName === dictName);
    };

    onMounted(() => {
      fetchMercTypes();
    });

    return {
      mercTypes,
      dialogVisible,
      isEdit,
      dialogTitle,
      form,
      currentDict,
      currentIndex,
      handleAdd,
      handleEdit,
      handleDelete,
      handleSave,
      selectDict
    };
  }
};
</script>

<style scoped>
/* Header styles */
.el-header {
  background-color: #f09eff;
  /* 浅粉色 */
  color: white;
  text-align: center;
  line-height: 60px;
  font-size: 18px;
  border-radius: 4px;
  margin-bottom: 20px;
}

/* Main content styles */
.el-main {
  background-color: white;
  padding: 20px;
  border-radius: 4px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

/* Table styles */
.dict-table {
  margin-top: 20px;
  width: 100%;
}

/* Button styles */
.el-button {
  margin-top: 20px;
}

/* Dialog styles */
.el-dialog__body {
  padding: 20px;
}

/* Form item styles */
.el-form-item {
  margin-bottom: 20px;
}

/* Aside (sidebar) styles */
.el-aside {
  background-color: #933744;
  /* 深紫色 */
  color: #1ff;
  /* 亮绿色 */
}

/* Menu item styles */
.el-menu-item {
  color: #f2f;
  /* 浅灰色 */
}

/* Hover effect for menu items */
.el-menu-item:hover {
  background-color: #9A4A4A;
  /* 暗紫色 */
}
</style>

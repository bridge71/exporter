<template>
  <el-container>
    <!-- 左侧导航栏 -->
    <el-aside width="200px">
      <el-menu default-active="customer" @select="handleMenuSelect">
        <el-menu-item index="customer">客户类型</el-menu-item>
        <el-menu-item index="supplier">供应商类型</el-menu-item>
      </el-menu>
    </el-aside>

    <!-- 右侧内容 -->
    <el-main>
      <!-- 添加按钮 -->
      <el-button type="primary" @click="handleAdd">添加</el-button>

      <!-- 表格 -->
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="type" label="类型" />
        <el-table-column label="操作">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 添加/编辑对话框 -->
      <el-dialog v-model="dialogVisible" :title="dialogTitle">
        <el-form :model="form">
          <el-form-item label="类型">
            <el-input v-model="form.type" />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const activeMenu = ref('customer'); // 当前选中的菜单项
    const tableData = ref([]); // 表格数据
    const dialogVisible = ref(false); // 对话框是否可见
    const dialogTitle = ref(''); // 对话框标题
    const form = ref({ type: '', id: '' }); // 表单数据
    const isEdit = ref(false); // 是否是编辑模式

    // 获取数据
    const fetchData = async () => {
      let url;
      if (activeMenu.value === 'customer') {
        url = '/find/mercType';
      } else if (activeMenu.value === 'supplier') {
        url = '/find/suprType';
      }
      const response = await axios.get(url);
      if (activeMenu.value === 'customer') {
        tableData.value = response.data.MercType.map(item => ({
          type: item.MercType,
          id: item.MercTypeId,
        }));
      } else if (activeMenu.value === 'supplier') {
        tableData.value = response.data.SuprType.map(item => ({
          type: item.SuprType,
          id: item.SuprTypeId,
        }));
      }
    };

    // 处理菜单选择
    const handleMenuSelect = (index) => {
      activeMenu.value = index;
      fetchData();
    };

    // 处理添加按钮点击
    const handleAdd = () => {
      dialogTitle.value = '添加';
      form.value = { type: '', id: '' };
      isEdit.value = false;
      dialogVisible.value = true;
    };

    // 处理编辑按钮点击
    const handleEdit = (row) => {
      dialogTitle.value = '编辑';
      form.value = { type: row.type, id: row.id };
      isEdit.value = true;
      dialogVisible.value = true;
    };

    // 处理删除按钮点击
    const handleDelete = async (row) => {
      let url;
      if (activeMenu.value === 'customer') {
        url = '/delete/mercType';
      } else if (activeMenu.value === 'supplier') {
        url = '/delete/suprType';
      }
      await axios.post(url, { MercTypeId: row.id, SuprTypeId: row.id });
      fetchData();
    };

    // 处理表单提交
    const handleSubmit = async () => {
      let url;
      let data;
      if (activeMenu.value === 'customer') {
        url = '/save/mercType';
        if (isEdit.value) {
          data = { MercType: form.value.type, MercTypeId: form.value.id }; // 编辑时传递 MercType 和 MercTypeId
        } else {
          data = { MercType: form.value.type }; // 添加时只传递 MercType
        }
      } else if (activeMenu.value === 'supplier') {
        url = '/save/suprType';
        if (isEdit.value) {
          data = { SuprType: form.value.type, SuprTypeId: form.value.id }; // 编辑时传递 SuprType 和 SuprTypeId
        } else {
          data = { SuprType: form.value.type }; // 添加时只传递 SuprType
        }
      }
      await axios.post(url, data); // 调用保存接口
      dialogVisible.value = false; // 关闭对话框
      fetchData(); // 重新获取数据
    };

    // 组件挂载时获取数据
    onMounted(() => {
      fetchData();
    });

    return {
      activeMenu,
      tableData,
      dialogVisible,
      dialogTitle,
      form,
      isEdit,
      handleMenuSelect,
      handleAdd,
      handleEdit,
      handleDelete,
      handleSubmit,
    };
  },
};
</script>

<style>
/* 修改全局文字颜色 */
body {
  color: #333;
}

/* 修改表格文字颜色 */
.el-table {
  color: #333;
}

/* 修改对话框文字颜色 */
.el-dialog {
  color: #333;
}

/* 修改导航栏文字颜色 */
.el-menu {
  color: #333;
}

/* 修改按钮文字颜色 */
.el-button {
  color: #333;
}

.el-aside {
  background-color: #f5f5f5;
}

.el-main {
  padding: 20px;
}
</style>

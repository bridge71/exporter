<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-6'" />
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>
        <el-header style="display: flex; justify-content: space-between; align-items: center;">
          <h2>{{ headerTitle }}</h2>
          <div>
            搜索：
            <el-input v-model="searchQuery" placeholder="输入要搜索的关键字" style="width: 200px;" />
            <el-button type="primary" @click="handleAdd">{{ addButtonText }}</el-button>
          </div>
        </el-header>
        <el-main>
          <!-- 支付方式信息表格 -->
          <el-table :data="paginatedSpotData" style="width: 100%" max-height="450">
            <el-table-column prop="InvLocName" label="库存地点名称" width="220%"></el-table-column>
            <el-table-column prop="InvLocAbbr" label="库存地点简称" width="220%"></el-table-column>
            <el-table-column prop="InvAddr" label="库存地址" width="320%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="320%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.ID)" type="text" size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="SpotData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <el-dialog v-model="showSpotDialog" title="库存地点信息" width="80%" @close="resetSpotForm">
      <el-form :model="SpotForm" label-width="150px" :rules="SpotRules" ref="SpotFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="库存地点名称" prop="InvLocName">
              <el-input v-model="SpotForm.InvLocName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="库存地点简称" prop="InvLocAbbr">
              <el-input v-model="SpotForm.InvLocAbbr"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="库存地址" prop="InvLocAddr">
              <el-input v-model="SpotForm.InvAddr"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="SpotForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitSpotForm">保存</el-button>
            <el-button @click="showSpotDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <el-dialog v-model="showViewSpotDialog" title="库存地点信息" width="80%" @close="resetSpotForm">
      <el-form :model="SpotForm" label-width="150px" :rules="SpotRules" ref="SpotFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="库存地点名称" prop="InvLocName">
              <el-input v-model="SpotForm.InvLocName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="库存地点简称" prop="InvLocAbbr">
              <el-input v-model="SpotForm.InvLocAbbr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="库存地址" prop="InvLocAddr">
              <el-input v-model="SpotForm.InvAddr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="SpotForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewSpotDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, customRef } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import axios from 'axios';
import SideMenu from '@/components/SideMenu.vue';

const searchQuery = ref('');
const currentPage = ref(1);
const pageSize = 8;
const SpotData = ref([]);
const showSpotDialog = ref(false);
const showViewSpotDialog = ref(false);
const SpotForm = ref({
  ID: '',
  InvLocName: '',
  InvLocAbbr: '',
  InvAddr: '',
  Notes: '',
});


const SpotFormRef = ref(null);
// 自定义验证函数
const validateNotEmpty = (rule, value, callback) => {
  if (value === '' || value === null || value === undefined) {
    callback(new Error(rule.message));  // 使用 rule.message 作为错误消息
  } else {
    callback();  // 验证通过
  }
};

// 库存地点表单验证规则
const SpotRules = {
  InvLocName: [
    { required: true, validator: validateNotEmpty, message: '请输入库存地点', trigger: 'blur' }
  ],
  // InvLocAbbr: [
  //   { required: true, validator: validateNotEmpty, message: '请输入库存地点简称', trigger: 'blur' }
  // ],
  // InvAddr: [
  //   { required: true, validator: validateNotEmpty, message: '请输入库存地址', trigger: 'blur' }
  // ]
};

const handlePageChange = (page) => {
  currentPage.value = page;
};

const paginatedSpotData = computed(() => {
  let filteredData = SpotData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.InvLocName.includes(searchQuery.value) ||
      item.InvLocAbbr.includes(searchQuery.value) ||
      item.InvAddr.includes(searchQuery.value) ||
      item.Notes.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

const handleAdd = () => {
  showSpotDialog.value = true;
};

const handleEdit = (index, row) => {
  SpotForm.value = { ...row };
  showSpotDialog.value = true;
};

const handleView = (index, row) => {
  SpotForm.value = { ...row };
  showViewSpotDialog.value = true;
};

const handleDelete = (index, SpotID) => {
  ElMessageBox.confirm('确定要删除该库存地点信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/spot', {
      ID: SpotID,
      InvLocName: "ss"
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchSpotData();
        } else {
          ElMessage.error(response.data.RetMessage || '删除失败');
        }
      })
      .catch(error => {
        ElMessage.error(error.response.data.RetMessage);
      });
  }).catch(() => {
    ElMessage.info('已取消删除');
  });
};

const resetSpotForm = () => {
  SpotForm.value = {
    ID: '',
    InvLocName: '',
    InvLocAbbr: '',
    InvAddr: '',
    Notes: '',
  };
};

const submitSpotForm = async () => {
  try {
    const isValid = await SpotFormRef.value.validate();
    if (!isValid) {
      ElMessage.error('请填写所有必填字段');
      console.log('验证不通过');
      return; // 如果验证不通过，阻止提交
    }

    SpotForm.value.ID = parseInt(SpotForm.value.ID, 10)
    console.log("success");
    const response = await axios.post('/save/spot', SpotForm.value);
    if (response.status === 200) {
      ElMessage.success('库存地点信息保存成功');
      showSpotDialog.value = false;
      fetchSpotData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('库存地点信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

onMounted(() => {
  fetchSpotData();
});

const fetchSpotData = async () => {
  try {
    const response = await axios.get('/find/spot');
    SpotData.value = response.data.Spot;
  } catch (error) {
    console.error('获取库存地点信息失败:', error);
    ElMessage.error('获取库存地点信息失败');
  }
};


const headerTitle = computed(() => '库存地点信息');
const addButtonText = computed(() => '添加库存地点信息');
</script>

<style src="../assets/styles/Bottom.css"></style>

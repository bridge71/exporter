<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-9'" />
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
          <!-- 品牌信息表格 -->
          <el-table :data="paginatedBrandData" style="width: 100%" max-height="450">
            <el-table-column prop="BrandName" label="品牌名称" width="220%"></el-table-column>
            <el-table-column prop="BrandEngName" label="品牌英文名称" width="220%"></el-table-column>
            <el-table-column prop="BrandType" label="品牌类型" width="220%"></el-table-column>
            <el-table-column prop="Nation" label="品牌国别" width="220%"></el-table-column>
            <el-table-column prop="BrandYear" label="品牌创建年份" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="320%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="320%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.BrandID)" type="text"
                    size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="brandData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加品牌信息的对话框 -->
    <el-dialog v-model="showBrandDialog" title="品牌信息" width="80%" @close="resetBrandForm">
      <el-form :model="brandForm" label-width="150px" :rules="brandRules" ref="brandFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="品牌名称" prop="BrandName">
              <el-input v-model="brandForm.BrandName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌英文名称" prop="BrandEngName">
              <el-input v-model="brandForm.BrandEngName"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="品牌类型" prop="BrandType">
              <el-input v-model="brandForm.BrandType"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌创建年份" prop="BrandYear">
              <el-input v-model="brandForm.BrandYear"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="品牌国别" prop="Nation">
              <el-select v-model="brandForm.Nation" @change="onNationChange" placeholder="请选择品牌国别">
                <el-option v-for="nation in nationData" :key="nation.Nation" :label="nation.Nation"
                  :value="nation.Nation"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="brandForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!brandForm.FileID" ref="brandUploadRef" action="" :limit="1"
                :on-change="handleBrandFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(brandForm.FileID, brandForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="brandForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitBrandForm">保存</el-button>
            <el-button @click="showBrandDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看品牌信息的对话框 -->
    <el-dialog v-model="showViewBrandDialog" title="品牌信息" width="80%" @close="resetBrandForm">
      <el-form :model="brandForm" label-width="150px" :rules="brandRules" ref="brandFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="品牌名称" prop="BrandName">
              <el-input v-model="brandForm.BrandName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌英文名称" prop="BrandEngName">
              <el-input v-model="brandForm.BrandEngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="品牌类型" prop="BrandType">
              <el-input v-model="brandForm.BrandType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌创建年份" prop="BrandYear">
              <el-input v-model="brandForm.BrandYear" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="品牌国别" prop="Nation">
              <el-select v-model="brandForm.Nation" @change="onNationChange" placeholder="请选择品牌国别" :disabled="true">
                <el-option v-for="nation in nationData" :key="nation.Nation" :label="nation.Nation"
                  :value="nation.Nation"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="brandForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-button v-if="brandForm.FileID" type="success"
                @click="downloadFile(brandForm.FileID, brandForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="brandForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewBrandDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import axios from 'axios';
import SideMenu from '@/components/SideMenu.vue';

const searchQuery = ref('');
const currentPage = ref(1);
const pageSize = 8;
const brandData = ref([]);
const nationData = ref([]);
const showBrandDialog = ref(false);
const showViewBrandDialog = ref(false);
const brandForm = ref({
  BrandName: '',
  BrandEngName: '',
  BrandType: '',
  Nation: '',
  BrandYear: '',
  Notes: '',
  FileID: '',
  FileName: '',
  BrandID: '',
});
const brandUploadRef = ref(null);
const brandFile = ref(null);

const handlePageChange = (page) => {
  currentPage.value = page;
};

const paginatedBrandData = computed(() => {
  let filteredData = brandData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.BrandName.includes(searchQuery.value) ||
      item.BrandEngName.includes(searchQuery.value) ||
      item.BrandType.includes(searchQuery.value) ||
      item.BrandYear.toString().includes(searchQuery.value) ||
      item.Notes.includes(searchQuery.value) ||
      item.FileName.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

const handleAdd = () => {
  showBrandDialog.value = true;
};

const handleEdit = (index, row) => {
  brandForm.value = { ...row };
  showBrandDialog.value = true;
};

const handleView = (index, row) => {
  brandForm.value = { ...row };
  showViewBrandDialog.value = true;
};

const handleDelete = (index, BrandID) => {
  ElMessageBox.confirm('确定要删除该品牌信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/brand', { BrandID })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchBrandData();
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

const resetBrandForm = () => {
  brandForm.value = {
    BrandName: '',
    BrandEngName: '',
    BrandType: '',
    Nation: '',
    BrandYear: '',
    Notes: '',
    FileID: '',
    FileName: '',
    BrandID: '',
  };
  brandFile.value = null;
  if (brandUploadRef.value) {
    brandUploadRef.value.clearFiles();
  }
};

const onNationChange = (value) => {
  const selectedNation = nationData.value.find(nation => nation.Nation === value);
  if (selectedNation) {
    brandForm.value.Nation = selectedNation.Nation;
  }
};

const submitBrandForm = async () => {
  try {

    const isValid = await brandFormRef.value.validate();
    if (!isValid) {
      ElMessage.error('请填写所有必填字段');
      console.log('验证不通过');
      return; // 如果验证不通过，阻止提交
    }

    const formData = new FormData();
    Object.keys(brandForm.value).forEach((key) => {
      formData.append(key, brandForm.value[key]);
    });
    if (brandFile.value) {
      formData.append('file', brandFile.value);
    }
    const response = await axios.post('/save/brand', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    if (response.status === 200) {
      ElMessage.success('品牌信息保存成功');
      showBrandDialog.value = false;
      fetchBrandData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存品牌信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

const handleBrandFileChange = (uploadFile) => {
  brandFile.value = uploadFile.raw;
};

const downloadFile = async (fileID, fileName) => {
  try {
    const response = await axios.post('/file', { FileID: fileID }, { responseType: 'blob' });
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', fileName || `file_${fileID}`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error('文件下载失败:', error);
    ElMessage.error('文件下载失败，请稍后重试');
  }
};

onMounted(() => {
  fetchBrandData();
  fetchNationData();
});

const fetchBrandData = async () => {
  try {
    const response = await axios.get('/find/brand');
    brandData.value = response.data.Brand;
  } catch (error) {
    console.error('获取品牌信息失败:', error);
    ElMessage.error('获取品牌信息失败，请稍后重试');
  }
};

const fetchNationData = async () => {
  try {
    const response = await axios.get('/find/nation');
    nationData.value = response.data.Nation;
  } catch (error) {
    console.error('获取国家信息失败:', error);
    ElMessage.error('获取国家信息失败，请稍后重试');
  }
};

const headerTitle = computed(() => '品牌信息');
const addButtonText = computed(() => '添加品牌信息');


const brandFormRef = ref(null);
// 自定义验证函数
const validateNotEmpty = (rule, value, callback) => {
  if (value === '' || value === null || value === undefined) {
    callback(new Error(rule.message));  // 使用 rule.message 作为错误消息
  } else {
    callback();  // 验证通过
  }
};

// 品牌表单验证规则
const brandRules = {
  BrandName: [
    { required: true, validator: validateNotEmpty, message: '请输入品牌名称', trigger: 'blur' }
  ],
  BrandEngName: [
    { required: true, validator: validateNotEmpty, message: '请输入品牌英文名称', trigger: 'blur' }
  ]
};
</script>

<style src="../assets/styles/Bottom.css"></style>

<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-10'" />
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
          <!-- 包装规格信息表格 -->
          <el-table :data="paginatedPackSpecData" style="width: 100%" max-height="450">
            <el-table-column prop="SpecName" label="规格名称" width="220%"></el-table-column>
            <el-table-column prop="SpecEngName" label="规格英文名称" width="220%"></el-table-column>
            <el-table-column prop="UnitMeas" label="单位" width="220%"></el-table-column>
            <el-table-column prop="PackType" label="包装类型" width="220%"></el-table-column>
            <el-table-column prop="NetWt" label="净重" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="320%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="320%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.PackSpecId)" type="text"
                    size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="packSpecData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加包装规格信息的对话框 -->
    <el-dialog v-model="showPackSpecDialog" title="包装规格信息" width="80%" @close="resetPackSpecForm">
      <el-form :model="packSpecForm" label-width="150px" :rules="packSpecRules" ref="packSpecFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="规格名称" prop="SpecName">
              <el-input v-model="packSpecForm.SpecName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="规格英文名称" prop="SpecEngName">
              <el-input v-model="packSpecForm.SpecEngName"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas">
              <el-select v-model="packSpecForm.UnitMeas" @change="onUnitMeasChange" placeholder="请选择单位">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeas" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装类型" prop="PackType">
              <el-select v-model="packSpecForm.PackType" @change="onPackTypeChange" placeholder="请选择包装类型">
                <el-option v-for="packType in packTypeData" :key="packType.PackType" :label="packType.PackType"
                  :value="packType.PackType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="净重" prop="NetWt">
              <el-input v-model.number="packSpecForm.NetWt" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="packSpecForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!packSpecForm.FileId" ref="packSpecUploadRef" action="" :limit="1"
                :on-change="handlePackSpecFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(packSpecForm.FileId, packSpecForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="packSpecForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitPackSpecForm">保存</el-button>
            <el-button @click="showPackSpecDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看包装规格信息的对话框 -->
    <el-dialog v-model="showViewPackSpecDialog" title="包装规格信息" width="80%" @close="resetPackSpecForm">
      <el-form :model="packSpecForm" label-width="150px" :rules="packSpecRules" ref="packSpecFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="规格名称" prop="SpecName">
              <el-input v-model="packSpecForm.SpecName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="规格英文名称" prop="SpecEngName">
              <el-input v-model="packSpecForm.SpecEngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas">
              <el-select v-model="packSpecForm.UnitMeas" @change="onUnitMeasChange" placeholder="请选择单位"
                :disabled="true">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeas" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装类型" prop="PackType">
              <el-select v-model="packSpecForm.PackType" @change="onPackTypeChange" placeholder="请选择包装类型"
                :disabled="true">
                <el-option v-for="packType in packTypeData" :key="packType.PackType" :label="packType.PackType"
                  :value="packType.PackType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="净重" prop="NetWt">
              <el-input v-model.number="packSpecForm.NetWt" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="packSpecForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-button v-if="packSpecForm.FileId" type="success"
                @click="downloadFile(packSpecForm.FileId, packSpecForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="packSpecForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewPackSpecDialog = false">关闭</el-button>
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
const packSpecData = ref([]);
const unitMeasData = ref([]); // 单位数据
const packTypeData = ref([]); // 包装类型数据
const showPackSpecDialog = ref(false);
const showViewPackSpecDialog = ref(false);
const packSpecForm = ref({
  SpecName: '',
  SpecEngName: '',
  UnitMeas: '',
  PackType: '',
  NetWt: 0,
  Notes: '',
  FileId: '',
  FileName: '',
  PackSpecId: '',
});
const packSpecUploadRef = ref(null);
const packSpecFile = ref(null);

const handlePageChange = (page) => {
  currentPage.value = page;
};

const paginatedPackSpecData = computed(() => {
  let filteredData = packSpecData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.SpecName.includes(searchQuery.value) ||
      item.SpecEngName.includes(searchQuery.value) ||
      item.UnitMeas.includes(searchQuery.value) ||
      item.PackType.includes(searchQuery.value) ||
      item.NetWt.toString().includes(searchQuery.value) ||
      item.Notes.includes(searchQuery.value) ||
      item.FileName.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

const handleAdd = () => {
  showPackSpecDialog.value = true;
};

const handleEdit = (index, row) => {
  packSpecForm.value = { ...row };
  showPackSpecDialog.value = true;
};

const handleView = (index, row) => {
  packSpecForm.value = { ...row };
  showViewPackSpecDialog.value = true;
};

const handleDelete = (index, PackSpecId) => {
  ElMessageBox.confirm('确定要删除该包装规格信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/packSpec', { PackSpecId })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchPackSpecData();
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

const resetPackSpecForm = () => {
  packSpecForm.value = {
    SpecName: '',
    SpecEngName: '',
    UnitMeas: '',
    PackType: '',
    NetWt: 0,
    Notes: '',
    FileId: '',
    FileName: '',
    PackSpecId: '',
  };
  packSpecFile.value = null;
  if (packSpecUploadRef.value) {
    packSpecUploadRef.value.clearFiles();
  }
};

const onUnitMeasChange = (value) => {
  const selectedUnitMeas = unitMeasData.value.find(unitMeas => unitMeas.UnitMeas === value);
  if (selectedUnitMeas) {
    packSpecForm.value.UnitMeas = selectedUnitMeas.UnitMeas;
  }
};

const onPackTypeChange = (value) => {
  const selectedPackType = packTypeData.value.find(packType => packType.PackType === value);
  if (selectedPackType) {
    packSpecForm.value.PackType = selectedPackType.PackType;
  }
};

const submitPackSpecForm = async () => {
  try {
    const formData = new FormData();
    Object.keys(packSpecForm.value).forEach((key) => {
      formData.append(key, packSpecForm.value[key]);
    });
    if (packSpecFile.value) {
      formData.append('file', packSpecFile.value);
    }
    const response = await axios.post('/save/packSpec', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    if (response.status === 200) {
      ElMessage.success('包装规格信息保存成功');
      showPackSpecDialog.value = false;
      fetchPackSpecData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存包装规格信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

const handlePackSpecFileChange = (uploadFile) => {
  packSpecFile.value = uploadFile.raw;
};

const downloadFile = async (fileId, fileName) => {
  try {
    const response = await axios.post('/file', { FileId: fileId }, { responseType: 'blob' });
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', fileName || `file_${fileId}`);
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
  fetchPackSpecData();
  fetchUnitMeasData();
  fetchPackTypeData();
});

const fetchPackSpecData = async () => {
  try {
    const response = await axios.get('/find/packSpec');
    packSpecData.value = response.data.PackSpec;
  } catch (error) {
    console.error('获取包装规格信息失败:', error);
    ElMessage.error('获取包装规格信息失败，请稍后重试');
  }
};

const fetchUnitMeasData = async () => {
  try {
    const response = await axios.get('/find/unitMeas');
    unitMeasData.value = response.data.UnitMeas;
  } catch (error) {
    console.error('获取单位信息失败:', error);
    ElMessage.error('获取单位信息失败，请稍后重试');
  }
};

const fetchPackTypeData = async () => {
  try {
    const response = await axios.get('/find/packType');
    packTypeData.value = response.data.PackType;
  } catch (error) {
    console.error('获取包装类型信息失败:', error);
    ElMessage.error('获取包装类型信息失败，请稍后重试');
  }
};

const headerTitle = computed(() => '包装规格信息');
const addButtonText = computed(() => '添加包装规格信息');

const packSpecRules = {
  SpecName: [{ required: true, message: '请输入规格名称', trigger: 'blur' }],
  SpecEngName: [{ required: true, message: '请输入规格英文名称', trigger: 'blur' }],
  UnitMeas: [{ required: true, message: '请选择单位', trigger: 'blur' }],
  PackType: [{ required: true, message: '请选择包装类型', trigger: 'blur' }],
  NetWt: [{ required: true, message: '请输入净重', trigger: 'blur' }],
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  margin-top: 60px;
}

.el-header {
  background-color: #f5f5f5;
  padding: 0 20px;
}

.el-aside {
  background-color: #304156;
}

.el-menu {
  border-right: none;
}

.el-menu-item,
.el-submenu__title {
  color: #fff;
}

.el-menu-item:hover,
.el-submenu__title:hover {
  background-color: #263445;
}

.el-main {
  padding: 20px;
}
</style>

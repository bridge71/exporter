<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-8'" />
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>
        
        <HeaderComponent :header-title="headerTitle" :add-button-text="addButtonText" v-model:search-query="searchQuery"
          @toggle-match-mode="toggleMatchMode" @toggle-id-mode="toggleIDMode" @add="handleAdd" />
        <el-header height="1px">
        </el-header>
        <el-main>
          <!-- 品类信息表格 -->
          <el-table :data="paginatedCatData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="CatAbbr" label="品类缩写" width="160"></el-table-column>
            <el-table-column prop="CatName" label="品类名称" width="160"></el-table-column>
            <el-table-column prop="CatEngName" label="品类英文名称" width="220"></el-table-column>
            <el-table-column prop="DeclHS" label="报关HS编码" width="160"></el-table-column>
            <el-table-column prop="DocHS" label="单据HS编码" width="160"></el-table-column>
            <el-table-column prop="CAS" label="CAS编码" width="160"></el-table-column>
            <el-table-column prop="StdDoc" label="国标文件" width="220"></el-table-column>
            <el-table-column prop="Component" label="成分含量" width="220"></el-table-column>
            <el-table-column prop="Source" label="底料来源" width="220"></el-table-column>
            <el-table-column prop="Use" label="用途" width="220"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="220"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220"></el-table-column>
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

          <!-- 分页 -->
          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="catData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加/编辑品类信息的对话框 -->
    <el-dialog v-model="showCatDialog" :title="dialogTitle" width="80%" @close="resetCatForm">
      <el-form :model="catForm" label-width="150px" :rules="catRules" ref="catFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="品类缩写" prop="CatAbbr">
              <el-input v-model="catForm.CatAbbr"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品类名称" prop="CatName">
              <el-input v-model="catForm.CatName"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="品类英文名称" prop="CatEngName">
              <el-input v-model="catForm.CatEngName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="报关HS编码" prop="DeclHS">
              <el-input v-model="catForm.DeclHS"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据HS编码" prop="DocHS">
              <el-input v-model="catForm.DocHS"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="CAS编码" prop="CAS">
              <el-input v-model="catForm.CAS"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国标编码" prop="GB">
              <el-input v-model="catForm.GB"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="成分含量" prop="Notes">
              <el-input v-model="catForm.Component" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="底料来源" prop="Source">
              <el-input v-model="catForm.Source" type="textarea"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="用途" prop="Use">
              <el-input v-model="catForm.Use" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="catForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!catForm.FileID" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(catForm.FileID, catForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="catForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitCatForm">保存</el-button>
            <el-button @click="showCatDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看品类信息的对话框 -->
    <el-dialog v-model="showViewCatDialog" title="品类信息" width="80%">
      <el-form :model="catForm" label-width="150px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="品类缩写">
              <el-input v-model="catForm.CatAbbr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品类名称">
              <el-input v-model="catForm.CatName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="品类英文名称">
              <el-input v-model="catForm.CatEngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="报关HS编码">
              <el-input v-model="catForm.DeclHS" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据HS编码">
              <el-input v-model="catForm.DocHS" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="CAS编码">
              <el-input v-model="catForm.CAS" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国标编码" prop="GB">
              <el-input v-model="catForm.GB" :readonly="true"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="成分含量" prop="Notes">
              <el-input v-model="catForm.Component" :readonly="true" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="底料来源" prop="Source">
              <el-input v-model="catForm.Source" :readonly="true" type="textarea"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="用途" prop="Use">
              <el-input v-model="catForm.Use" :readonly="true" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注">
              <el-input v-model="catForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 文件下载 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-input v-model="catForm.FileName" :readonly="true"></el-input>
              <el-button v-if="catForm.FileID" type="success" style="margin-top: 10px;"
                @click="downloadFile(catForm.FileID, catForm.FileName)">下载文件</el-button>
              <span v-else style="color: #999;">无文件</span>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewCatDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, readonly } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import axios from 'axios';
import SideMenu from '@/components/SideMenu.vue';
import HeaderComponent from '@/components/HeaderComponent.vue';

const isExactMatch = ref(true);
const onlyID = ref(true);

const toggleMatchMode = () => {
  console.log("check onlyID", isExactMatch.value)
  isExactMatch.value = !isExactMatch.value;
};

const toggleIDMode = () => {

console.log("check match", onlyID.value)
onlyID.value = !onlyID.value;
};


// 搜索查询字段
const searchQuery = ref('');

// 分页相关
const currentPage = ref(1);
const pageSize = 10;

// 表格数据
const catData = ref([]);

// 对话框显示状态
const showCatDialog = ref(false);
const showViewCatDialog = ref(false);

// 表单数据
const catForm = ref({
  CatAbbr: '',
  CatName: '',
  CatEngName: '',
  DeclHS: '',
  DocHS: '',
  CAS: '',
  GB: '',
  Notes: '',
  Source: '',
  Component: '',
  Use: '',
  FileID: '',
  FileName: '',
  CatID: '',
  ID: ''
});

// 自定义验证函数
const validateNotEmpty = (rule, value, callback) => {
  if (value === '' || value === null || value === undefined) {
    callback(new Error(rule.message));  // 使用 rule.message 作为错误消息
  } else {
    callback();  // 验证通过
  }
};

const catFormRef = ref(null);

// 品类表单验证规则
const catRules = {
  CatAbbr: [
    { required: true, validator: validateNotEmpty, message: '请输入品类缩写', trigger: 'blur' }
  ],
  CatName: [
    { required: true, validator: validateNotEmpty, message: '请输入品类名称', trigger: 'blur' }
  ],
  CatEngName: [
    { required: true, validator: validateNotEmpty, message: '请输入品类英文名称', trigger: 'blur' }
  ]
};

// 分页数据
const paginatedCatData = computed(() => {
  let filteredData = catData.value || [];

  if (searchQuery.value) {
    console.log(isExactMatch.value);
    console.log(onlyID.value);

    if (isExactMatch.value === false) {
      if (onlyID.value === false) {
        filteredData = filteredData.filter(item =>
          item.CatAbbr?.includes(searchQuery.value) ||
          item.CatName?.includes(searchQuery.value) ||
          item.CatEngName?.includes(searchQuery.value) ||
          item.DeclHS?.includes(searchQuery.value) ||
          item.DocHS?.includes(searchQuery.value) ||
          item.CAS?.includes(searchQuery.value) ||
          item.GB?.includes(searchQuery.value) ||
          item.StdDoc?.includes(searchQuery.value) ||
          item.Notes?.includes(searchQuery.value)
        );
      } else {
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value)
        );
      }
    } else {
      if (onlyID.value === false) {
        filteredData = filteredData.filter(item =>
          item.CatAbbr === searchQuery.value ||
          item.CatName === searchQuery.value ||
          item.CatEngName === searchQuery.value ||
          item.DeclHS === searchQuery.value ||
          item.DocHS === searchQuery.value ||
          item.CAS === searchQuery.value ||
          item.GB === searchQuery.value ||
          item.StdDoc === searchQuery.value ||
          item.Notes === searchQuery.value
        );
      } else {
        filteredData = filteredData.filter(item =>
          item.ID.toString() === searchQuery.value
        );
      }
    }
  }

  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

// 页面加载时获取数据
onMounted(() => {
  fetchCatData();
});

// 获取品类数据
const fetchCatData = async () => {
  try {
    const response = await axios.get('/find/cat');
    // 根据提供的 JSON 返回文件，数据在 Cat 字段中
    catData.value = response.data.Cat || [];
  } catch (error) {
    ElMessage.error('获取品类数据失败，请稍后重试');
    catData.value = [];
  }
};

// 添加按钮点击事件
const handleAdd = () => {
  resetCatForm();
  showCatDialog.value = true;
};

// 编辑按钮点击事件
const handleEdit = (index, row) => {
  catForm.value = { ...row };
  showCatDialog.value = true;
};

// 查看按钮点击事件
const handleView = (index, row) => {
  catForm.value = { ...row };
  showViewCatDialog.value = true;
};

// 删除按钮点击事件
const handleDelete = (index, CatID) => {
  ElMessageBox.confirm('确定要删除该品类吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/cat', {
      ID: CatID,
      CatEngName: "ss",
      CatName: "ss",
      CatAbbr: "ss"
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchCatData();
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

// 提交表单
const submitCatForm = async () => {
  try {

    const isValid = await catFormRef.value.validate();
    if (!isValid) {
      ElMessage.error('请填写所有必填字段');
      console.log('验证不通过');
      return; // 如果验证不通过，阻止提交
    }

    const formData = new FormData();
    Object.keys(catForm.value).forEach((key) => {
      formData.append(key, catForm.value[key]);
    });

    if (file.value) {
      formData.append('file', file.value);
    }

    const response = await axios.post('/save/cat', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    if (response.status === 200) {
      ElMessage.success('保存成功');
      showCatDialog.value = false;
      fetchCatData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    ElMessage.error(error.response.data.RetMessage || '保存失败');
  }
};

// 重置表单
const resetCatForm = () => {
  catForm.value = {
    CatAbbr: '',
    CatName: '',
    CatEngName: '',
    DeclHS: '',
    DocHS: '',
    CAS: '',
    GB: '',
    StdDoc: '',
    Notes: '',
    FileID: '',
    FileName: '',
    CatID: '',
  };
  file.value = null;
  if (uploadRef.value) {
    uploadRef.value.clearFiles();
  }
};

// 分页切换
const handlePageChange = (page) => {
  currentPage.value = page;
};

// 文件上传相关
const uploadRef = ref(null);
const file = ref(null);

const handleFileChange = (uploadFile) => {
  file.value = uploadFile.raw;
};

const downloadFile = async (fileID, fileName) => {
  try {
    const response = await axios.post(
      '/file',
      { FileID: fileID },
      { responseType: 'blob' }
    );

    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', fileName || `file_${fileID}`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
  } catch (error) {
    ElMessage.error('文件下载失败，请稍后重试');
  }
};

// 标题和按钮文本
const headerTitle = computed(() => '品类管理');
const addButtonText = computed(() => '添加品类');
</script>

<style src="../assets/styles/Bottom.css"></style>

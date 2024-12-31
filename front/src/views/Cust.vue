<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-4'" />
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
          <!-- 客户信息表格 -->
          <el-table :data="paginatedCustData" style="width: 100%" max-height="450">
            <el-table-column prop="Name" label="客户姓名" width="220%"></el-table-column>
            <el-table-column prop="Gender" label="性别" width="220%"></el-table-column>
            <el-table-column prop="Nation" label="国家" width="220%"></el-table-column>
            <el-table-column prop="Addr" label="地址" width="220%"></el-table-column>
            <el-table-column prop="Email" label="邮箱" width="220%"></el-table-column>
            <el-table-column prop="PhoneNum" label="联系电话" width="220%"></el-table-column>
            <el-table-column prop="QQ" label="QQ" width="220%"></el-table-column>
            <el-table-column prop="Wechat" label="微信" width="220%"></el-table-column>
            <el-table-column label="所属商户" width="320%">
              <template #default="scope">
                <span v-if="scope.row.Merc">{{ scope.row.Merc }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column prop="Post" label="职位" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="220%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>

            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.CustId)" type="text"
                    size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="custData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加客户信息的对话框 -->
    <el-dialog v-model="showCustDialog" title="客户信息" width="80%" @close="resetCustForm">
      <el-form :model="custForm" label-width="150px" :rules="custRules" ref="custFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客户姓名" prop="Name" :required="true">
              <el-input v-model="custForm.Name" placeholder="请输入客户姓名"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="性别" prop="Gender" :required="true">
              <el-select v-model="custForm.Gender" placeholder="请选择性别">
                <el-option label="男" value="男"></el-option>
                <el-option label="女" value="女"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation" :required="true">
              <el-input v-model="custForm.Nation" placeholder="请输入国家"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地址" prop="Addr">
              <el-input v-model="custForm.Addr" placeholder="请输入地址"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="custForm.Email" placeholder="请输入邮箱"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="custForm.PhoneNum" placeholder="请输入联系电话"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="QQ" prop="QQ">
              <el-input v-model="custForm.QQ" placeholder="请输入QQ"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="微信" prop="Wechat">
              <el-input v-model="custForm.Wechat" placeholder="请输入微信"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="所属商户" prop="MercId">
              <el-select v-model="custForm.MercId" @change="onMercChange" placeholder="请选择商户信息">
                <el-option v-for="merc in merchantData" :key="merc.MercId" :label="`${merc.Merc} (${merc.MercCode})`"
                  :value="merc.MercId"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="职位" prop="Post">
              <el-input v-model="custForm.Post" placeholder="请输入职位"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="custForm.Notes" type="textarea" placeholder="请输入备注"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!custForm.FileId" ref="custUploadRef" action="" :limit="1"
                :on-change="handleCustFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(custForm.FileId, custForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="custForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitCustForm">保存</el-button>
            <el-button @click="showCustDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看客户信息的对话框 -->
    <el-dialog v-model="showViewCustDialog" title="客户信息" width="80%" @close="resetCustForm">
      <el-form :model="custForm" label-width="150px" :rules="custRules" ref="custFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客户姓名" prop="Name">
              <el-input v-model="custForm.Name" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="性别" prop="Gender">
              <el-input v-model="custForm.Gender" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="custForm.Nation" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地址" prop="Addr">
              <el-input v-model="custForm.Addr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="custForm.Email" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="custForm.PhoneNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="QQ" prop="QQ">
              <el-input v-model="custForm.QQ" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="微信" prop="Wechat">
              <el-input v-model="custForm.Wechat" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="所属商户" prop="Merc">
              <el-input v-model="custForm.Merc" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="职位" prop="Post">
              <el-input v-model="custForm.Post" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="custForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!custForm.FileId" ref="custUploadRef" action="" :limit="1"
                :on-change="handleCustFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(custForm.FileId, custForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="custForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewCustDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import axios from 'axios';
import SideMenu from '@/components/SideMenu.vue';

const searchQuery = ref('');
const currentPage = ref(1);
const pageSize = 8;
const showCustDialog = ref(false);
const showViewCustDialog = ref(false);
const custForm = ref({
  Name: '',
  Gender: '',
  Nation: '',
  Addr: '',
  Email: '',
  PhoneNum: '',
  QQ: '',
  Wechat: '',
  MercId: '',
  Merc: '',
  Post: '',
  Notes: '',
  FileName: '',
  FileId: ''
});
const custRules = {
  Name: [{ required: true, message: '请输入客户姓名', trigger: 'blur' }],
  Gender: [{ required: true, message: '请选择性别', trigger: 'blur' }],
  Nation: [{ required: true, message: '请输入国家', trigger: 'blur' }]
};

const custData = ref([]);
const merchantData = ref([]);
const custUploadRef = ref(null);
const custFile = ref(null);

// 分页计算
const paginatedCustData = computed(() => {
  let filteredData = custData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.Name.includes(searchQuery.value) ||
      item.Gender.includes(searchQuery.value) ||
      item.Nation.includes(searchQuery.value) ||
      item.Addr.includes(searchQuery.value) ||
      item.Email.includes(searchQuery.value) ||
      item.PhoneNum.includes(searchQuery.value) ||
      item.QQ.includes(searchQuery.value) ||
      item.Wechat.includes(searchQuery.value) ||
      item.Merc.includes(searchQuery.value) ||
      item.Post.includes(searchQuery.value) ||
      item.Notes.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

// 文件选择事件
const handleCustFileChange = (uploadFile) => {
  custFile.value = uploadFile.raw;
};

// 下载文件
const downloadFile = async (fileId, fileName) => {
  try {
    const response = await axios.post(
      '/file',
      { FileId: fileId },
      { responseType: 'blob' }
    );
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

// 查看按钮逻辑
const handleView = (index, row) => {
  custForm.value = { ...row };
  showViewCustDialog.value = true;
};

// 重置表单
const resetCustForm = () => {
  custForm.value = {
    Name: '',
    Gender: '',
    Nation: '',
    Addr: '',
    Email: '',
    PhoneNum: '',
    QQ: '',
    Wechat: '',
    MercId: '',
    Merc: '',
    Post: '',
    Notes: '',
    FileName: '',
    FileId: ''
  };
  custFile.value = null;
  if (custUploadRef.value) {
    custUploadRef.value.clearFiles();
  }
};

// 提交表单
const submitCustForm = async () => {
  try {
    const formData = new FormData();
    Object.keys(custForm.value).forEach((key) => {
      formData.append(key, custForm.value[key]);
    });
    if (custFile.value) {
      formData.append('file', custFile.value);
    }
    const response = await axios.post('/save/cust', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    if (response.status === 200) {
      ElMessage.success('客户信息保存成功');
      showCustDialog.value = false;
      fetchCustData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 删除按钮逻辑
const handleDelete = (index, CustId) => {
  ElMessageBox.confirm('确定要删除该客户信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/cust', {
      CustId: CustId,
      Name: custData.value[index].Name,
      Gender: custData.value[index].Gender,
      Nation: custData.value[index].Nation
    })
      .then(response => {
        if (response.status === 200) {

          ElMessage.success('删除成功');
          fetchCustData(); // 重新获取客户信息数据
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

// 分页切换
const handlePageChange = (page) => {
  currentPage.value = page;
};

// 添加按钮点击事件
const handleAdd = () => {
  showCustDialog.value = true;
};

// 编辑按钮逻辑
const handleEdit = (index, row) => {
  custForm.value = { ...row };
  showCustDialog.value = true;
};

// 获取客户数据
const fetchCustData = async () => {
  try {
    const response = await axios.get('/find/cust');
    custData.value = response.data.Cust;
  } catch (error) {
    console.error('获取客户信息失败:', error);
    ElMessage.error('获取客户信息失败，请稍后重试');
  }
};

// 获取商户数据
const fetchMerchantData = async () => {
  try {
    const response = await axios.get('/find/merchant');
    merchantData.value = response.data.Merchant;
  } catch (error) {
    console.error('获取商户信息失败:', error);
    ElMessage.error('获取商户信息失败，请稍后重试');
  }
};

// 监听商户选择事件
const onMercChange = (value) => {
  const selectedMerc = merchantData.value.find(merc => merc.MercId === value);
  if (selectedMerc) {
    custForm.value.Merc = selectedMerc.Merc;
  }
};

onMounted(() => {
  fetchCustData();
  fetchMerchantData();
});

const headerTitle = computed(() => {
  return '客户信息';
});

const addButtonText = computed(() => {
  return '添加客户信息';
});
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
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.el-input {
  margin-right: 10px;
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

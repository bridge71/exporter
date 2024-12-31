<template>
<<<<<<< HEAD
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-5'" />
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>
        <el-header style="display: flex; justify-content: space-between; align-items: center;">
          <h2>{{ headerTitle }}</h2>
          <div>
            搜索：
            <el-input v-model="searchQuery" placeholder="输入要搜索的关键字" style="width: 200px;" @input="handleSearch" />
            <el-button type="primary" @click="handleAdd">{{ addButtonText }}</el-button>
          </div>
        </el-header>
        <el-main>
          <!-- 银行账户信息表格 -->
          <el-table :data="paginatedBankAccountData" style="width: 100%" max-height="450">
            <el-table-column prop="BankAccName" label="银行账户名称" width="220%"></el-table-column>
            <el-table-column prop="CompName" label="公司名称" width="220%"></el-table-column>
            <el-table-column prop="AcctNum" label="账号" width="220%"></el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="BankName" label="开户行名称" width="220%"></el-table-column>
            <el-table-column prop="BankEng" label="开户行英文名称" width="220%"></el-table-column>
            <el-table-column prop="BankNum" label="开户行号" width="220%"></el-table-column>
            <el-table-column prop="SwiftCode" label="SWIFT CODE" width="220%"></el-table-column>
            <el-table-column prop="BankAddr" label="开户行地址" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="220%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>

            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.BankAccId)" type="text"
                    size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize"
            :total="filteredBankAccountData.length" layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加银行账户信息的对话框 -->
    <el-dialog v-model="showBankAccountDialog" title="银行账户信息" width="80%" @close="resetBankAccountForm">
      <el-form :model="bankAccountForm" label-width="150px" :rules="bankAccountRules" ref="bankAccountFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="银行账户名称" prop="BankAccName">
              <el-input v-model="bankAccountForm.BankAccName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="公司名称" prop="CompName">
              <el-input v-model="bankAccountForm.CompName"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账号" prop="AcctNum">
              <el-input v-model="bankAccountForm.AcctNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-input v-model="bankAccountForm.Currency"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行名称" prop="BankName">
              <el-input v-model="bankAccountForm.BankName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="开户行英文名称" prop="BankEng">
              <el-input v-model="bankAccountForm.BankEng"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行号" prop="BankNum">
              <el-input v-model="bankAccountForm.BankNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="SWIFT CODE" prop="SwiftCode">
              <el-input v-model="bankAccountForm.SwiftCode"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行地址" prop="BankAddr">
              <el-input v-model="bankAccountForm.BankAddr"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="bankAccountForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!bankAccountForm.FileId" ref="bankAccountUploadRef" action="" :limit="1"
                :on-change="handleBankAccountFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(bankAccountForm.FileId, bankAccountForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="bankAccountForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 商户信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联商户信息" prop="MercId">
              <el-select v-model="bankAccountForm.MercId" @change="onMercChange" placeholder="请选择商户信息">
                <el-option v-for="merc in merchantData" :key="merc.MercId" :label="`${merc.Merc} (${merc.MercCode})`"
                  :value="merc.MercId"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitBankAccountForm">保存</el-button>
            <el-button @click="showBankAccountDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看银行账户信息的对话框 -->
    <el-dialog v-model="showViewBankAccountDialog" title="银行账户信息" width="80%" @close="resetBankAccountForm">
      <el-form :model="bankAccountForm" label-width="150px" :rules="bankAccountRules" ref="bankAccountFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="银行账户名称" prop="BankAccName">
              <el-input v-model="bankAccountForm.BankAccName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="公司名称" prop="CompName">
              <el-input v-model="bankAccountForm.CompName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账号" prop="AcctNum">
              <el-input v-model="bankAccountForm.AcctNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-input v-model="bankAccountForm.Currency" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行名称" prop="BankName">
              <el-input v-model="bankAccountForm.BankName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="开户行英文名称" prop="BankEng">
              <el-input v-model="bankAccountForm.BankEng" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行号" prop="BankNum">
              <el-input v-model="bankAccountForm.BankNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="SWIFT CODE" prop="SwiftCode">
              <el-input v-model="bankAccountForm.SwiftCode" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行地址" prop="BankAddr">
              <el-input v-model="bankAccountForm.BankAddr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="bankAccountForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!bankAccountForm.FileId" ref="bankAccountUploadRef" action="" :limit="1"
                :on-change="handleBankAccountFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(bankAccountForm.FileId, bankAccountForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="bankAccountForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 商户信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联商户信息">
              <el-select v-model="bankAccountForm.MercId" @change="onMercChange" placeholder="请选择商户信息">
                <el-option v-for="merc in merchantData" :key="merc.MercId" :label="`${merc.Merc} (${merc.MercCode})`"
                  :value="merc.MercId"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewBankAccountDialog = false">关闭</el-button>
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
const showBankAccountDialog = ref(false);
const showViewBankAccountDialog = ref(false);
const bankAccountForm = ref({
  BankAccName: '',
  CompName: '',
  AcctNum: '',
  Currency: '',
  BankName: '',
  BankEng: '',
  BankNum: '',
  SwiftCode: '',
  BankAddr: '',
  Notes: '',
  FileName: '',
  FileId: '',
  MercId: '',
  Merc: ''
});
const bankAccountRules = {
  BankAccName: [{ required: true, message: '请输入银行账户名称', trigger: 'blur' }],
  CompName: [{ required: true, message: '请输入公司名称', trigger: 'blur' }],
  AcctNum: [{ required: true, message: '请输入账号', trigger: 'blur' }]
};

const bankAccountData = ref([]);
const merchantData = ref([]);
const bankAccountUploadRef = ref(null);
const bankAccountFile = ref(null);

// 分页计算
const paginatedBankAccountData = computed(() => {
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredBankAccountData.value.slice(start, end);
});

// 过滤后的数据
const filteredBankAccountData = computed(() => {
  if (!searchQuery.value) {
    return bankAccountData.value;
  }
  return bankAccountData.value.filter(item =>
    item.BankAccName.includes(searchQuery.value) ||
    item.CompName.includes(searchQuery.value) ||
    item.AcctNum.includes(searchQuery.value) ||
    item.Currency.includes(searchQuery.value) ||
    item.BankName.includes(searchQuery.value) ||
    item.BankEng.includes(searchQuery.value) ||
    item.BankNum.includes(searchQuery.value) ||
    item.SwiftCode.includes(searchQuery.value) ||
    item.BankAddr.includes(searchQuery.value) ||
    item.Notes.includes(searchQuery.value)
    // item.Merc.includes(searchQuery.value)
  );
});

// 搜索事件
const handleSearch = () => {
  currentPage.value = 1; // 重置分页到第一页
};

// 文件选择事件
const handleBankAccountFileChange = (uploadFile) => {
  bankAccountFile.value = uploadFile.raw;
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
  bankAccountForm.value = { ...row };
  showViewBankAccountDialog.value = true;
};

// 重置表单
const resetBankAccountForm = () => {
  bankAccountForm.value = {
    BankAccName: '',
    CompName: '',
    AcctNum: '',
    Currency: '',
    BankName: '',
    BankEng: '',
    BankNum: '',
    SwiftCode: '',
    BankAddr: '',
    Notes: '',
    FileName: '',
    FileId: '',
    MercId: '',
    Merc: ''
  };
  bankAccountFile.value = null;
  if (bankAccountUploadRef.value) {
    bankAccountUploadRef.value.clearFiles();
  }
};

// 提交表单
const submitBankAccountForm = async () => {
  try {
    const formData = new FormData();
    Object.keys(bankAccountForm.value).forEach((key) => {
      formData.append(key, bankAccountForm.value[key]);
    });
    if (bankAccountFile.value) {
      formData.append('file', bankAccountFile.value);
    }
    const response = await axios.post('/save/bankAccount', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    if (response.status === 200) {
      ElMessage.success('银行账户信息保存成功');
      showBankAccountDialog.value = false;
      fetchBankAccountData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 删除按钮逻辑
const handleDelete = (index, BankAccId) => {
  ElMessageBox.confirm('确定要删除该银行账户信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/bankAccount', {
      BankAccId: BankAccId,
      BankAccName: bankAccountData.value[index].BankAccName,
      CompName: bankAccountData.value[index].CompName,
      AcctNum: bankAccountData.value[index].AcctNum
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchBankAccountData();
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
  showBankAccountDialog.value = true;
};

// 编辑按钮逻辑
const handleEdit = (index, row) => {
  bankAccountForm.value = { ...row };
  showBankAccountDialog.value = true;
};

// 获取银行账户数据
const fetchBankAccountData = async () => {
  try {
    const response = await axios.get('/find/bankAccount');
    bankAccountData.value = response.data.BankAccount;
  } catch (error) {
    console.error('获取银行账户信息失败:', error);
    ElMessage.error('获取银行账户信息失败，请稍后重试');
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
    bankAccountForm.value.Merc = selectedMerc.Merc;
  }
};

onMounted(() => {
  fetchBankAccountData();
  fetchMerchantData();
});

const headerTitle = computed(() => {
  return '银行账户信息';
});

const addButtonText = computed(() => {
  return '添加银行账户信息';
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
=======
    <div class="bank-information">
      <p>银行信息</p>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  
  export default defineComponent({
    name: 'bank-information',
  });
  </script>
  
  <style scoped>
  /* 你可以在这里添加样式 */
  </style>
>>>>>>> 593b97d905b86d5274251b6be73a3d8c8270a1cc

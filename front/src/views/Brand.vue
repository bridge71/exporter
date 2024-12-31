<template>
<<<<<<< HEAD
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu />
        </div>
      </el-aside>

      <el-container>
        <el-header style="display: flex; justify-content: space-between; align-items: center;">
          <h2>会计实体银行账户信息</h2>
          <div>
            搜索：
            <el-input v-model="searchQuery" placeholder="输入要搜索的关键字" style="width: 200px;" />
            <el-button type="primary" @click="handleAdd">添加银行账户信息</el-button>
          </div>
        </el-header>
        <el-main>
          <!-- 银行账户信息表格 -->
          <el-table :data="paginatedBankData" style="width: 100%" max-height="450">
            <el-table-column prop="AccName" label="账户名称" width="160"></el-table-column>
            <el-table-column prop="AccNum" label="账户号码" width="160"></el-table-column>
            <el-table-column prop="Currency" label="货币" width="120"></el-table-column>
            <el-table-column prop="BankName" label="银行名称" width="160"></el-table-column>
            <el-table-column prop="BankNum" label="银行代码" width="120"></el-table-column>
            <el-table-column prop="SwiftCode" label="SWIFT代码" width="120"></el-table-column>
            <el-table-column prop="BankAddr" label="银行地址" width="160"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="160"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="160"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.AcctBankId)" type="text"
                    size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="bankData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加银行账户信息的对话框 -->
    <el-dialog v-model="showBankDialog" title="银行账户信息" width="80%" @close="resetBankForm">
      <el-form :model="bankForm" label-width="150px" :rules="bankRules" ref="bankFormRef">
        <!-- 会计实体信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联会计实体信息" prop="AcctId" :required="true">
              <el-select v-model="bankForm.AcctId" @change="onAcctChange" placeholder="请选择会计实体信息">
                <el-option v-for="acct in acctData" :key="acct.AcctCode" :label="`${acct.AcctName} (${acct.AcctCode})`"
                  :value="acct.AcctId"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账户名称" prop="AccName" :required="true">
              <el-input v-model="bankForm.AccName" placeholder="请输入账户名称"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="账户号码" prop="AccNum" :required="true">
              <el-input v-model="bankForm.AccNum" placeholder="请输入账户号码"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="货币" prop="Currency">
              <el-input v-model="bankForm.Currency" placeholder="请输入货币"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="银行名称" prop="BankName">
              <el-input v-model="bankForm.BankName" placeholder="请输入银行名称"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="银行代码" prop="BankNum">
              <el-input v-model="bankForm.BankNum" placeholder="请输入银行代码"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="SWIFT代码" prop="SwiftCode">
              <el-input v-model="bankForm.SwiftCode" placeholder="请输入SWIFT代码"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="银行地址" prop="BankAddr">
              <el-input v-model="bankForm.BankAddr" placeholder="请输入银行地址"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="bankForm.Notes" placeholder="请输入备注"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!bankForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(bankForm.FileId, bankForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="bankForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitBankForm">保存</el-button>
            <el-button @click="showBankDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看银行账户信息的对话框 -->
    <el-dialog v-model="showViewBankDialog" title="银行账户信息" width="80%" @close="resetBankForm">
      <el-form :model="bankForm" label-width="150px" :rules="bankRules" ref="bankFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账户名称" prop="AccName">
              <el-input v-model="bankForm.AccName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="账户号码" prop="AccNum">
              <el-input v-model="bankForm.AccNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="货币" prop="Currency">
              <el-input v-model="bankForm.Currency" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="银行名称" prop="BankName">
              <el-input v-model="bankForm.BankName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="银行代码" prop="BankNum">
              <el-input v-model="bankForm.BankNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="SWIFT代码" prop="SwiftCode">
              <el-input v-model="bankForm.SwiftCode" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="银行地址" prop="BankAddr">
              <el-input v-model="bankForm.BankAddr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="bankForm.Notes" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewBankDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import SideMenu from '@/components/SideMenu.vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import axios from 'axios';

const searchQuery = ref('');
const currentPage = ref(1);
const pageSize = 8;
const bankData = ref([]);
const acctData = ref([]); // 会计实体数据
const showBankDialog = ref(false);
const showViewBankDialog = ref(false);
const bankForm = ref({
  AccName: '',
  AccNum: '',
  Currency: '',
  BankName: '',
  BankNum: '',
  SwiftCode: '',
  BankAddr: '',
  Notes: '',
  FileName: '',
  FileId: null,
  AcctId: null, // 外键，关联会计实体
});
const bankRules = {
  AccName: [{ required: true, message: '请输入账户名称', trigger: 'blur' }],
  AccNum: [{ required: true, message: '请输入账户号码', trigger: 'blur' }],
  AcctId: [{ required: true, message: '请选择会计实体', trigger: 'change' }],
};

// 获取会计实体数据
const fetchAcctData = async () => {
  try {
    const response = await axios.get('/find/acct');
    acctData.value = response.data.Acct;
  } catch (error) {
    console.error('获取会计实体信息失败:', error);
    ElMessage.error('获取会计实体信息失败，请稍后重试');
  }
};

// 获取银行账户数据
const fetchBankData = async () => {
  try {
    const response = await axios.get('/find/acctBank');
    bankData.value = response.data.AcctBank;
  } catch (error) {
    console.error('获取银行账户信息失败:', error);
    ElMessage.error('获取银行账户信息失败，请稍后重试');
  }
};

// 分页数据
const paginatedBankData = computed(() => {
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return bankData.value.slice(start, end);
});

// 添加按钮点击事件
const handleAdd = () => {
  resetBankForm();
  showBankDialog.value = true;
};

// 编辑按钮逻辑
const handleEdit = (index, row) => {
  bankForm.value = { ...row };
  showBankDialog.value = true;
};

// 查看按钮逻辑
const handleView = (index, row) => {
  bankForm.value = { ...row };
  showViewBankDialog.value = true;
};

// 删除按钮逻辑
const handleDelete = (index, AcctBankId) => {
  ElMessageBox.confirm('确定要删除该银行账户信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    axios.post('/delete/acctBank', { AcctBankId })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchBankData();
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

// 提交银行账户表单
const submitBankForm = async () => {
  try {
    const response = await axios.post('/save/acctBank', bankForm.value);
    if (response.status === 200) {
      ElMessage.success('保存成功');
      showBankDialog.value = false;
      fetchBankData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 重置表单
const resetBankForm = () => {
  bankForm.value = {
    AccName: '',
    AccNum: '',
    Currency: '',
    BankName: '',
    BankNum: '',
    SwiftCode: '',
    BankAddr: '',
    Notes: '',
    FileName: '',
    FileId: null,
    AcctId: null,
  };
};

// 会计实体选择事件
const onAcctChange = (acctId) => {
  bankForm.value.AcctId = acctId;
};

onMounted(() => {
  fetchAcctData();
  fetchBankData();
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

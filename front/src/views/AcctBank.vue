<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-2'" />
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

          <!-- 会计实体银行账户信息表格 -->
          <el-table :data="paginatedBankData" style="width: 100%" max-height="450">
            <el-table-column prop="AccName" label="账户名称" width="220%"></el-table-column>
            <el-table-column prop="AccNum" label="账号" width="220%"></el-table-column>

            <el-table-column label="对应的会计实体信息" width="320%">
              <template #default="scope">
                <span v-if="scope.row.Acct.AcctName">{{ scope.row.Acct.AcctName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="BankName" label="开户行名称" width="220%"></el-table-column>
            <el-table-column prop="BankNum" label="开户行号" width="320%"></el-table-column>
            <el-table-column prop="SwiftCode" label="SWIFT CODE" width="220%"></el-table-column>
            <el-table-column prop="BankAddr" label="开户行地址" width="280%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="320%"></el-table-column>

            <el-table-column prop="FileName" label="文件名" width="320%"></el-table-column>
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

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="bankData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>


    <!-- 添加会计实体银行账户信息的对话框 -->
    <el-dialog v-model="showBankDialog" title="会计实体银行账户信息" width="80%" @close="resetBankForm">
      <el-form :model="bankForm" label-width="150px" :rules="bankRules" ref="bankFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账户名称" prop="AccName">
              <el-input v-model="bankForm.AccName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="账号" prop="AccNum">
              <el-input v-model="bankForm.AccNum"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">

            <el-form-item label="币种">
              <el-select v-model="bankForm.Currency" placeholder="请选择币种">
                <el-option v-for="currency in currencyData" :key="currency.CurrencyID" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="开户行名称" prop="BankName">
              <el-input v-model="bankForm.BankName"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行号" prop="BankNum">
              <el-input v-model="bankForm.BankNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="SWIFT CODE" prop="SwiftCode">
              <el-input v-model="bankForm.SwiftCode"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行地址" prop="BankAddr">
              <el-input v-model="bankForm.BankAddr"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="bankForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="附件">
              <el-upload v-if="!bankForm.FileID" ref="bankUploadRef" action="" :limit="1"
                :on-change="handleBankFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(bankForm.FileID, bankForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="bankForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


        <!-- 会计实体信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联会计实体信息">
              <el-select v-model="bankForm.AcctID" @change=onAcctChange placeholder="请选择会计实体信息">
                <el-option v-for="acct in acctData" :key="acct.AcctCode" :label="`${acct.AcctName} (${acct.AcctCode})`"
                  :value="acct.ID"></el-option>
              </el-select>
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


    <!-- 添加会计实体银行账户信息的对话框 -->
    <el-dialog v-model="showshowBankDialog" title="会计实体银行账户信息" width="80%" @close="resetBankForm">
      <el-form :model="bankForm" label-width="150px" :rules="bankRules" ref="bankFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账户名称" prop="AccName">
              <el-input v-model="bankForm.AccName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="账号" prop="AccNum">
              <el-input v-model="bankForm.AccNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-input v-model="bankForm.Currency" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="开户行名称" prop="BankName">
              <el-input v-model="bankForm.BankName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行号" prop="BankNum">
              <el-input v-model="bankForm.BankNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="SWIFT CODE" prop="SwiftCode">
              <el-input v-model="bankForm.SwiftCode" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行地址" prop="BankAddr">
              <el-input v-model="bankForm.BankAddr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="bankForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="附件">
              <el-upload v-if="!bankForm.FileID" ref="bankUploadRef" action="" :limit="1"
                :on-change="handleBankFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(bankForm.FileID, bankForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="bankForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 会计实体信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联会计实体信息">
              <el-select v-model="bankForm.AcctID" :disabled="true" @change=onAcctChange placeholder="请选择会计实体信息">
                <el-option v-for="acct in acctData" :key="acct.AcctCode" :label="`${acct.AcctName} (${acct.AcctCode})`"
                  :value="acct.ID"></el-option>
              </el-select>

            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showshowBankDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>

import { ref, onMounted, computed } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import axios from 'axios'; // 引入 axios
import SideMenu from '@/components/SideMenu.vue'; // 引入 SideMenu 组件
const bankFormRef = ref(null);
const searchQuery = ref(''); // 添加搜索查询字段
const downloadFile = async (fileID, fileName) => {
  try {
    const response = await axios.post(
      '/file',
      { FileID: fileID }, // 提供 FileID 表单
      {
        responseType: 'blob', // 指定响应类型为二进制数据
      }
    );

    // 创建下载链接
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    // link.setAttribute('download', `file_${fileID}`); // 设置下载文件名
    link.setAttribute('download', `${fileName}` || `file_${fileID}`); // 使用 fileName 作为文件名，如果不存在则使用默认文件名
    document.body.appendChild(link);
    link.click(); // 触发下载
    document.body.removeChild(link); // 移除链接
    window.URL.revokeObjectURL(url); // 释放 URL 对象
  } catch (error) {
    console.error('文件下载失败:', error);
    ElMessage.error('文件下载失败，请稍后重试');
  }
};


const handleDelete = (index, ID) => {
  ElMessageBox.confirm('确定要删除该银行账户信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    axios.post('/delete/acctBank', params, {

      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchAcctBankData(); // 重新获取银行账户信息数据
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

// 文件数据（会计实体银行账户信息）
const bankFile = ref(null);

// 文件选择事件（会计实体银行账户信息）
const handleBankFileChange = (uploadFile) => {
  bankFile.value = uploadFile.raw; // 保存选择的文件
};

// 查看按钮逻辑
const handleView = (index, row) => {
  // 填充会计实体银行账户信息表单
  bankForm.value = { ...row }; // 将当前行的数据赋值给 bankForm
  selectedAcct.value = { AcctID: row.ID, AcctName: row.AcctName }; // 填充关联的会计实体信息
  showshowBankDialog.value = true; // 打开会计实体银行账户信息对话框
  // 检查是否已上传文件
  if (row.FileID) {
    bankForm.value.FileID = row.FileID; // 保存 FileID
    bankForm.value.FileName = row.FileName; // 保存 FileID
  }
};
// 获取 el-upload 组件的引用（会计实体银行账户信息）
const bankUploadRef = ref(null);
// 重置会计实体银行账户信息表单
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
    File: '',
    AcctID: '', // 关联的会计实体 ID

  };
  selectedAcct.value = null; // 重置选择的会计实体信息
  bankFile.value = null; // 重置文件数据
  if (bankUploadRef.value) {
    bankUploadRef.value.clearFiles(); // 清空文件列表
  }
};
// 监听 change 事件并更新 AcctName
function onAcctChange(value) {
  const selectedAcct = acctData.value.find(acct => acct.ID === value);
}
const selectedAcct = ref({ AcctID: '', AcctName: '' }); // 默认值为空对象

// 会计实体银行账户信息表单提交逻辑
const submitBankForm = async () => {

  try {
    const isValid = await bankFormRef.value.validate();
    if (!isValid) {
      ElMessage.error('请填写所有必填字段');
      console.log('验证不通过');
      return; // 如果验证不通过，阻止提交
    }

    const formData = new FormData(); // 创建 FormData 对象
    console.log(bankForm.value)

    Object.keys(bankForm.value).forEach((key) => {
      if (key != 'Acct') {
        formData.append(key, bankForm.value[key]);
      }
    });

    // 添加文件
    if (bankFile.value) {
      formData.append('file', bankFile.value);
    }

    formData.append("AcctCode", '23')
    formData.append("EtyAbbr", '23')
    formData.append("AcctName", '23')
    formData.append("AcctAbbr", '23')

    console.log(formData)

    const response = await axios.post('/save/acctBank', formData, {
      headers: {
        'Content-Type': 'multipart/form-data', // 设置请求头
      },
    });

    // 将选择的会计实体信息赋值给 bankForm
    // const response = await axios.post('/save/acctBank', bankForm.value); // 调用保存会计实体银行账户信息接口
    if (response.status === 200) {
      ElMessage.success('会计实体银行账户信息保存成功');
      showBankDialog.value = false; // 关闭对话框
      fetchAcctBankData(); // 重新获取会计实体银行账户信息数据
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存会计实体银行账户信息失败:', error);
    // ElMessage.error(error.response.data.RetMessage);
    ElMessage.error(error.response.data.RetMessage);
  }
};
const handlePageChange = (page) => {
  currentPage.value = page;
};
const currentPage = ref(1); // 当前页码
const pageSize = 8; // 每页显示的行数

const paginatedBankData = computed(() => {
  let filteredData = bankData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.AccName.includes(searchQuery.value) ||
      item.AccNum.includes(searchQuery.value) ||
      item.Acct.AcctName.includes(searchQuery.value) ||
      item.Currency.includes(searchQuery.value) ||
      item.BankName.includes(searchQuery.value) ||
      item.BankNum.includes(searchQuery.value) ||
      item.SwiftCode.includes(searchQuery.value) ||
      item.BankAddr.includes(searchQuery.value) ||
      item.Notes.includes(searchQuery.value) ||
      item.FileName.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});


const currencyData = ref([]); // 存储货币数据

const fetchCurrencyData = async () => {
  try {
    const response = await axios.get('/find/currency'); // 调用货币数据接口
    currencyData.value = response.data.Currency; // 假设返回的数据结构中有 Currency 字段
  } catch (error) {
    console.error('获取货币数据失败:', error);
    ElMessage.error('获取货币数据失败');
  }
};
onMounted(() => {
  fetchAcctData(); // 获取会计实体信息
  fetchAcctBankData(); // 获取会计实体银行账户信息
  fetchCurrencyData();
});
// 定义接口请求函数
const fetchAcctData = async () => {
  try {
    const response = await axios.get('/find/acct'); // 调用会计实体信息接口
    acctData.value = response.data.Acct; // 假设返回的数据结构中有 Acct 字段
    console.log("已经赋值了");
  } catch (error) {
    console.error('获取会计实体信息失败:', error);
    ElMessage.error('获取会计实体信息失败，请稍后重试');
  }
};

const fetchAcctBankData = async () => {
  //fetchAcctData();

  console.log(acctData.value[0]);
  try {
    const response = await axios.get('/find/acctBank'); // 调用会计实体银行账户信息接口
    bankData.value = response.data.AcctBank; // 假设返回的数据结构中有 AcctBank 字段
    // console.log("这是Bank");
    //console.log(acctData.value);
    //console.log(acctData.value[0].AcctBanks[0].AccName);
    // for (let i = 0; i < bankData.value.length; i++)
    //   for (let j = 0; j < acctData.value.length; j++) {
    //     for (let k = 0; k < acctData.value[j].AcctBanks.length; k++) {
    //       if (acctData.value[j].AcctBanks[k].AccName == bankData.value[i].AccName) {
    //         bankData.value[i].AcctName = acctData.value[j].AcctName;
    //       }
    //     }
    //   }
    //console.log(bankData.value);

  } catch (error) {
    ElMessage.error('获取会计实体银行账户信息失败，请稍后重试');
    console.error('获取会计实体银行账户信息失败:', error);
  }
};

// 会计实体银行账户信息对话框显示状态
const showBankDialog = ref(false);

// 会计实体银行账户信息对话框显示状态
const showshowBankDialog = ref(false);
// 会计实体银行账户信息表单数据
const bankForm = ref({
  AccName: '',
  AccNum: '',
  Currency: '',
  BankName: '',
  BankNum: '',
  SwiftCode: '',
  BankAddr: '',
  Notes: '',
  File: '',
  AcctID: '',
  FileID: '', // 添加 FileName 字段
  FileName: '', // 添加 FileName 字段
});


// 自定义验证函数
const validateNotEmpty = (rule, value, callback) => {
  if (value === '' || value === null || value === undefined) {
    callback(new Error(rule.message));
  } else {
    callback();
  }
};

// 会计实体银行账户信息表单验证规则
const bankRules = {
  AccName: [
    { required: true, validator: validateNotEmpty, message: '请输入账户名称', trigger: 'blur' },
  ],
  AccNum: [
    { required: true, validator: validateNotEmpty, message: '请输入账号', trigger: 'blur' },
  ]
};


// 表格数据（初始值为空数组）
const acctData = ref([]); // 会计实体信息
const bankData = ref([]); // 会计实体银行账户信息
// 根据当前选中的菜单项动态更改标题和按钮文本
const headerTitle = computed(() => {
  return '会计实体银行账户信息';
});

const addButtonText = computed(() => {
  return '添加会计实体银行账户信息';
});


// 添加按钮点击事件
const handleAdd = () => {
  showBankDialog.value = true;
};

// 编辑按钮逻辑
const handleEdit = (index, row) => {
  // 填充会计实体银行账户信息表单
  bankForm.value = { ...row }; // 将当前行的数据赋值给 bankForm
  selectedAcct.value = { ID: row.AcctID, AcctName: row.AcctName }; // 填充关联的会计实体信息
  showBankDialog.value = true; // 打开会计实体银行账户信息对话框
};
</script>

<style src="../assets/styles/Bottom.css"></style>

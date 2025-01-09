<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-3'" />
        </div>
      </el-aside>

      <el-container>
        <HeaderComponent :header-title="headerTitle" :add-button-text="addButtonText" v-model:search-query="searchQuery"
          @toggle-match-mode="toggleMatchMode" @toggle-id-mode="toggleIDMode" @add="handleAdd" />
        <el-header height="1px">
        </el-header>
        <el-main>
          <!-- 商户信息表格 -->
          <el-table :data="paginatedMerchantData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="MercCode" label="商户编码" width="160%"></el-table-column>
            <el-table-column prop="MercAbbr" label="商户缩写" width="160%"></el-table-column>
            <el-table-column prop="ShortMerc" label="商户简称" width="160%"></el-table-column>
            <el-table-column prop="Merc" label="商户名称" width="220%"></el-table-column>

            <!-- <el-table-column label="绑定的客户信息" width="220%"> -->
            <!--   <template #default="scope"> -->
            <!--     <span v-if="scope.row.Custs && scope.row.Custs.length > 0"> -->
            <!--       {{ scope.row.Custs.map(cust => cust.Name).join(', ') }} -->
            <!--     </span> -->
            <!--     <span v-else>无</span> -->
            <!--   </template> -->
            <!-- </el-table-column> -->
            <!---->
            <!-- <el-table-column label="绑定的银行账户信息" width="220%"> -->
            <!--   <template #default="scope"> -->
            <!--     <span v-if="scope.row.BankAccounts && scope.row.BankAccounts.length > 0"> -->
            <!--       {{ scope.row.BankAccounts.map(account => account.AcctNum).join(', ') }} -->
            <!--     </span> -->
            <!--     <span v-else>无</span> -->
            <!--   </template> -->
            <!-- </el-table-column> -->

            <el-table-column prop="EngName" label="商户英文名称" width="220%"></el-table-column>
            <el-table-column prop="Address" label="商户地址" width="220%"></el-table-column>
            <el-table-column prop="Nation" label="国家"></el-table-column>
            <el-table-column prop="PhoneNum" label="联系电话" width="220%"></el-table-column>
            <el-table-column prop="Email" label="邮箱" width="220%"></el-table-column>
            <el-table-column prop="Fax" label="传真" width="220%"></el-table-column>
            <el-table-column prop="Website" label="网站" width="220%"></el-table-column>
            <el-table-column prop="TaxType" label="税号类型" width="220%"></el-table-column>
            <el-table-column prop="TaxCode" label="税号" width="220%"></el-table-column>
            <el-table-column prop="MercType" label="商户类型" width="220%"></el-table-column>
            <el-table-column prop="RegCap" label="注册资本" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="220%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>

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

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="merchantData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加商户信息的对话框 -->
    <el-dialog v-model="showMerchantDialog" title="商户信息" width="80%" @close="resetMerchantForm">
      <el-form :model="merchantForm" label-width="150px" :rules="merchantRules" ref="merchantFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="商户编码" prop="MercCode" :required="true">
              <el-input v-model="merchantForm.MercCode" placeholder="商户编码"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="商户缩写" prop="MercAbbr" :required="true">
              <el-input v-model="merchantForm.MercAbbr" placeholder="商户缩写"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="商户简称" prop="ShortMerc">
              <el-input v-model="merchantForm.ShortMerc" placeholder="商户简称"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="商户名称" prop="Merc" :required="true">
              <el-input v-model="merchantForm.Merc" placeholder="商户名称"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="商户英文名称" prop="EngName">
              <el-input v-model="merchantForm.EngName" placeholder="商户英文名称"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="商户地址" prop="Address">
              <el-input v-model="merchantForm.Address" placeholder="商户地址"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-select v-model="merchantForm.Nation" placeholder="请选择国家">
                <el-option v-for="nation in nationData" :key="nation.NationID" :label="nation.Nation"
                  :value="nation.Nation"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="merchantForm.PhoneNum" placeholder="联系电话"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="merchantForm.Email" placeholder="邮箱"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="传真" prop="Fax">
              <el-input v-model="merchantForm.Fax" placeholder="传真"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="网站" prop="Website">
              <el-input v-model="merchantForm.Website" placeholder="网站"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-select v-model="merchantForm.TaxType" placeholder="请选择">
                <el-option label="类型1" value="type1"></el-option>
                <el-option label="类型2" value="type2"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="税号" prop="TaxCode">
              <el-input v-model="merchantForm.TaxCode" placeholder="税号"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="商户类型" prop="MercType">
              <el-input v-model="merchantForm.MercType" placeholder="商户类型"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册资本" prop="RegCap">
              <el-input v-model="merchantForm.RegCap" placeholder="注册资本"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="merchantForm.Notes" type="textarea" placeholder="备注"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="营业执照">
              <el-upload v-if="!merchantForm.FileID" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(merchantForm.FileID, merchantForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="merchantForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 关联客户信息 -->
        <!-- <el-row :gutter="20"> -->
        <!--   <el-col :span="24"> -->
        <!--     <el-form-item label="关联客户信息"> -->
        <!--       <el-input v-model="merchantForm.CustsDisplay" :readonly="true"></el-input> -->
        <!--     </el-form-item> -->
        <!--   </el-col> -->
        <!-- </el-row> -->
        <!---->
        <!-- <!-- 关联银行账户信息 -->
        <!-- <el-row :gutter="20"> -->
        <!--   <el-col :span="24"> -->
        <!--     <el-form-item label="关联银行账户信息"> -->
        <!--       <el-input v-model="merchantForm.BankAccountsDisplay" :readonly="true"></el-input> -->
        <!--     </el-form-item> -->
        <!--   </el-col> -->
        <!-- </el-row> -->

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitMerchantForm">保存</el-button>
            <el-button @click="showMerchantDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看商户信息的对话框 -->
    <el-dialog v-model="showViewMerchantDialog" title="商户信息" width="80%" @close="resetMerchantForm">
      <el-form :model="merchantForm" label-width="150px" :rules="merchantRules" ref="merchantFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="商户编码" prop="MercCode" :required="true">
              <el-input v-model="merchantForm.MercCode" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="商户缩写" prop="MercAbbr" :required="true">
              <el-input v-model="merchantForm.MercAbbr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="商户简称" prop="ShortMerc">
              <el-input v-model="merchantForm.ShortMerc" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="商户名称" prop="Merc" :required="true">
              <el-input v-model="merchantForm.Merc" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="商户英文名称" prop="EngName">
              <el-input v-model="merchantForm.EngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="商户地址" prop="Address">
              <el-input v-model="merchantForm.Address" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="merchantForm.Nation" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="merchantForm.PhoneNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="merchantForm.Email" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="传真" prop="Fax">
              <el-input v-model="merchantForm.Fax" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="网站" prop="Website">
              <el-input v-model="merchantForm.Website" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-input v-model="merchantForm.TaxType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="税号" prop="TaxCode">
              <el-input v-model="merchantForm.TaxCode" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="商户类型" prop="MercType">
              <el-input v-model="merchantForm.MercType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册资本" prop="RegCap">
              <el-input v-model="merchantForm.RegCap" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="merchantForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="营业执照">
              <el-upload v-if="!merchantForm.FileID" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(merchantForm.FileID, merchantForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="merchantForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 关联客户信息 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联客户信息">
              <el-input v-model="merchantForm.CustsDisplay" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 关联银行账户信息 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联银行账户信息">
              <el-input v-model="merchantForm.BankAccountsDisplay" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewMerchantDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import SideMenu from '@/components/SideMenu.vue';
import HeaderComponent from '@/components/HeaderComponent.vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import axios from 'axios';

const isExactMatch = ref(true);
const onlyID = ref(true);
const searchQuery = ref('');
const currentPage = ref(1);
const pageSize = 8;
const showMerchantDialog = ref(false);
const showViewMerchantDialog = ref(false);
const merchantForm = ref({
  MercCode: '',
  MercAbbr: '',
  ShortMerc: '',
  Merc: '',
  EngName: '',
  Address: '',
  Nation: '',
  PhoneNum: '',
  Email: '',
  Fax: '',
  Website: '',
  TaxType: '',
  TaxCode: '',
  MercType: '',
  RegCap: '',
  Notes: '',
  FileName: '',
  ID: '',
  // Custs: [],
  // BankAccounts: [],
  CustsDisplay: '',
  BankAccountsDisplay: ''
});


const merchantFormRef = ref(null);
// 自定义验证函数
const validateNotEmpty = (rule, value, callback) => {
  if (value === '' || value === null || value === undefined) {
    callback(new Error(rule.message));  // 使用 rule.message 作为错误消息
  } else {
    callback();  // 验证通过
  }
};

const toggleMatchMode = () => {
  console.log("check onlyID", isExactMatch.value)
  isExactMatch.value = !isExactMatch.value;
};

const toggleIDMode = () => {

console.log("check match", onlyID.value)
onlyID.value = !onlyID.value;
};

// 商户表单验证规则
const merchantRules = {
  MercCode: [
    { required: true, validator: validateNotEmpty, message: '请输入商户编码', trigger: 'blur' }
  ],
  MercAbbr: [
    { required: true, validator: validateNotEmpty, message: '请输入商户缩写', trigger: 'blur' }
  ],
  // ShortMerc: [
  //   { required: true, validator: validateNotEmpty, message: '请输入商户简称', trigger: 'blur' }
  // ],
  Merc: [
    { required: true, validator: validateNotEmpty, message: '请输入商户名称', trigger: 'blur' }
  ]
};

const merchantData = ref([]);
const uploadRef = ref(null);
const file = ref(null);

// 分页计算
const paginatedMerchantData = computed(() => {
  let filteredData = merchantData.value;
  if (searchQuery.value) {
    if (isExactMatch.value === false) { // 非精确匹配
      if (onlyID.value === false) { // 不是仅按ID匹配
        filteredData = filteredData.filter(item =>
          item.MercCode.includes(searchQuery.value) ||
          item.MercAbbr.includes(searchQuery.value) ||
          item.ShortMerc.includes(searchQuery.value) ||
          item.Merc.includes(searchQuery.value) ||
          item.EngName.includes(searchQuery.value) ||
          item.Address.includes(searchQuery.value) ||
          item.Nation.includes(searchQuery.value) ||
          item.PhoneNum.includes(searchQuery.value) ||
          item.Email.includes(searchQuery.value) ||
          item.Fax.includes(searchQuery.value) ||
          item.Website.includes(searchQuery.value) ||
          item.TaxType.includes(searchQuery.value) ||
          item.TaxCode.includes(searchQuery.value) ||
          item.MercType.includes(searchQuery.value) ||
          item.RegCap.includes(searchQuery.value) ||
          item.Notes.includes(searchQuery.value) ||
          item.FileName.includes(searchQuery.value) ||
          (item.Custs && item.Custs.some(cust => cust.Name.includes(searchQuery.value))) ||
          (item.BankAccounts && item.BankAccounts.some(account => account.AcctNum.includes(searchQuery.value)))
        );
      } else { // 仅按ID匹配
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value)
        );
      }
    } else { // 精确匹配
      if (onlyID.value === false) { // 不是仅按ID匹配
        filteredData = filteredData.filter(item =>
          item.MercCode === searchQuery.value ||
          item.MercAbbr === searchQuery.value ||
          item.ShortMerc === searchQuery.value ||
          item.Merc === searchQuery.value ||
          item.EngName === searchQuery.value ||
          item.Address === searchQuery.value ||
          item.Nation === searchQuery.value ||
          item.PhoneNum === searchQuery.value ||
          item.Email === searchQuery.value ||
          item.Fax === searchQuery.value ||
          item.Website === searchQuery.value ||
          item.TaxType === searchQuery.value ||
          item.TaxCode === searchQuery.value ||
          item.MercType === searchQuery.value ||
          item.RegCap === searchQuery.value ||
          item.Notes === searchQuery.value ||
          item.FileName === searchQuery.value ||
          (item.Custs && item.Custs.some(cust => cust.Name === searchQuery.value)) ||
          (item.BankAccounts && item.BankAccounts.some(account => account.AcctNum === searchQuery.value))
        );
      } else { // 仅按ID匹配
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

// 文件选择事件
const handleFileChange = (uploadFile) => {
  file.value = uploadFile.raw;
};

// 下载文件
const downloadFile = async (ID, fileName) => {
  try {
    const response = await axios.post(
      '/file',
      { ID: ID },
      { responseType: 'blob' }
    );
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', fileName || `file_${ID}`);
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
  merchantForm.value = { ...row };
  merchantForm.value.CustsDisplay = row.Custs && row.Custs.length > 0
    ? row.Custs.map(cust => cust.Name).join(', ')
    : '无';
  merchantForm.value.BankAccountsDisplay = row.BankAccounts && row.BankAccounts.length > 0
    ? row.BankAccounts.map(account => account.AcctNum).join(', ')
    : '无';
  showViewMerchantDialog.value = true;
};

// 重置表单
const resetMerchantForm = () => {
  merchantForm.value = {
    MercCode: '',
    MercAbbr: '',
    ShortMerc: '',
    Merc: '',
    EngName: '',
    Address: '',
    Nation: '',
    PhoneNum: '',
    Email: '',
    Fax: '',
    Website: '',
    TaxType: '',
    TaxCode: '',
    MercType: '',
    RegCap: '',
    Notes: '',
    FileName: '',
    ID: '',
    // Custs: [],
    // BankAccounts: [],
    CustsDisplay: '',
    BankAccountsDisplay: ''
  };
  file.value = null;
  if (uploadRef.value) {
    uploadRef.value.clearFiles();
  }
};

// 提交表单
const submitMerchantForm = async () => {
  try {
    const isValid = await merchantFormRef.value.validate();
    if (!isValid) {
      ElMessage.error('请填写所有必填字段');
      console.log('验证不通过');
      return; // 如果验证不通过，阻止提交
    }

    const formData = new FormData();
    Object.keys(merchantForm.value).forEach((key) => {
      formData.append(key, merchantForm.value[key]);
    });
    if (file.value) {
      formData.append('file', file.value);
    }
    const response = await axios.post('/save/merchant', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    if (response.status === 200) {
      ElMessage.success('商户信息保存成功');
      showMerchantDialog.value = false;
      fetchMerchantData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 删除按钮逻辑
const handleDelete = (index, ID) => {
  ElMessageBox.confirm('确定要删除该商户信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/merchant', {
      ID: ID,
      MercAbbr: merchantData.value[index].MercAbbr,
      MercCode: merchantData.value[index].MercCode,
      Merc: merchantData.value[index].Merc,
      ShortMerc: merchantData.value[index].ShortMerc
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchMerchantData();
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
  showMerchantDialog.value = true;
};

// 编辑按钮逻辑
const handleEdit = (index, row) => {
  merchantForm.value = { ...row };
  showMerchantDialog.value = true;
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

const nationData = ref([]);

const fetchNationData = async () => {
  try {
    const response = await axios.get('/find/nation');
    nationData.value = response.data.Nation;
  } catch (error) {
    console.error('获取国家信息失败:', error);
    ElMessage.error('获取国家信息失败，请稍后重试');
  }
};
onMounted(() => {
  fetchNationData();
  fetchMerchantData();
});

const headerTitle = computed(() => {
  return '商户信息';
});

const addButtonText = computed(() => {
  return '添加商户信息';
});
</script>

<style src="../assets/styles/Bottom.css"></style>

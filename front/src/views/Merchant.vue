<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px" style="background-color: #304156;">
        <div style="height: 100vh; overflow-y: auto;">
          <el-menu :default-active="activeMenu" class="el-menu-vertical-demo" @select="handleMenuSelect"
            background-color="#304156" text-color="#fff" active-text-color="#ffd04b">
            <el-submenu index="1">
              <template #title>会计实体信息</template>
              <el-menu-item index="1-1" @click="pushAcct">会计实体信息</el-menu-item>
              <el-menu-item index="1-2" @click="pushAcctBank">会计实体银行账户信息</el-menu-item>
              <el-menu-item index="1-3" @click="pushMerchant">客商信息</el-menu-item>
              <el-menu-item index="1-4">联系人信息</el-menu-item>
              <el-menu-item index="1-5">银行账户信息</el-menu-item>
              <el-menu-item index="1-6">库存地点信息</el-menu-item>
              <el-menu-item index="1-7">付款方式信息</el-menu-item>
              <el-menu-item index="1-8">品类信息</el-menu-item>
              <el-menu-item index="1-9">品牌信息</el-menu-item>
              <el-menu-item index="1-10">包装规格信息</el-menu-item>
              <el-menu-item index="1-11">员工信息</el-menu-item>
              <el-menu-item index="1-12">产品明细</el-menu-item>
              <el-menu-item index="1-13">装货明细</el-menu-item>
              <el-menu-item index="1-14">费用明细</el-menu-item>
              <el-menu-item index="1-15">销售订单</el-menu-item>
              <el-menu-item index="1-16">采购收单</el-menu-item>
              <el-menu-item index="1-17">应付账款单</el-menu-item>
              <el-menu-item index="1-18">收款单</el-menu-item>
              <el-menu-item index="1-19">付款单</el-menu-item>
            </el-submenu>
          </el-menu>
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>
        <el-header
          style="background-color: #f5f5f5; padding: 0 20px; display: flex; justify-content: space-between; align-items: center;">
          <h2>{{ headerTitle }}</h2>
          <div>
            搜索：
            <el-input v-model="searchQuery" placeholder="输入要搜索的关键字" style="width: 200px;" />
            <el-button type="primary" @click="handleAdd">{{ addButtonText }}</el-button>
          </div>
        </el-header>
        <el-main>
          <!-- 客商信息表格 -->
          <el-table :data="paginatedMerchantData" style="width: 100%" max-height="450" v-if="activeMenu === '1-3'">
            <el-table-column prop="MercCode" label="客商编码" width="160"></el-table-column>
            <el-table-column prop="MercAbbr" label="客商缩写" width="160"></el-table-column>
            <el-table-column prop="ShortMerc" label="客商简称" width="160"></el-table-column>
            <el-table-column prop="Merc" label="客商名称" width="220"></el-table-column>
            <el-table-column prop="EngName" label="英文名称" width="220"></el-table-column>
            <el-table-column prop="Address" label="地址" width="220"></el-table-column>
            <el-table-column prop="Nation" label="国家" width="220"></el-table-column>
            <el-table-column prop="PhoneNum" label="联系电话" width="220"></el-table-column>
            <el-table-column prop="Email" label="邮箱" width="220"></el-table-column>
            <el-table-column prop="Fax" label="传真" width="220"></el-table-column>
            <el-table-column prop="Website" label="网站" width="220"></el-table-column>
            <el-table-column prop="TaxType" label="税号类型" width="220"></el-table-column>
            <el-table-column prop="TaxCode" label="税号" width="220"></el-table-column>
            <el-table-column prop="MercType" label="客商类型" width="220"></el-table-column>
            <el-table-column prop="RegCap" label="注册资本" width="220"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="220"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.MercId)" type="text"
                    size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize"
            :total="activeMenu === '1-3' ? merchantData.length : 0" layout="prev, pager, next"
            @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加客商信息的对话框 -->
    <el-dialog v-model="showMerchantDialog" title="客商信息" width="80%" @close="resetMerchantForm">
      <el-form :model="merchantForm" label-width="150px" :rules="merchantRules" ref="merchantFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商编码" prop="MercCode" :required="true">
              <el-input v-model="merchantForm.MercCode" placeholder="请输入客商编码"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商缩写" prop="MercAbbr" :required="true">
              <el-input v-model="merchantForm.MercAbbr" placeholder="请输入客商缩写"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商简称" prop="ShortMerc" :required="true">
              <el-input v-model="merchantForm.ShortMerc" placeholder="请输入客商简称"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商名称" prop="Merc" :required="true">
              <el-input v-model="merchantForm.Merc" placeholder="请输入客商名称"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="英文名称" prop="EngName">
              <el-input v-model="merchantForm.EngName" placeholder="请输入英文名称"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地址" prop="Address">
              <el-input v-model="merchantForm.Address" placeholder="请输入地址"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="merchantForm.Nation" placeholder="请输入国家"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="merchantForm.PhoneNum" placeholder="请输入联系电话"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="merchantForm.Email" placeholder="请输入邮箱"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="传真" prop="Fax">
              <el-input v-model="merchantForm.Fax" placeholder="请输入传真"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="网站" prop="Website">
              <el-input v-model="merchantForm.Website" placeholder="请输入网站"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-select v-model="merchantForm.TaxType" placeholder="请选择税号类型">
                <el-option label="类型1" value="type1"></el-option>
                <el-option label="类型2" value="type2"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="税号" prop="TaxCode">
              <el-input v-model="merchantForm.TaxCode" placeholder="请输入税号"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商类型" prop="MercType">
              <el-select v-model="merchantForm.MercType" placeholder="请选择客商类型">
                <el-option label="类型1" value="type1"></el-option>
                <el-option label="类型2" value="type2"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册资本" prop="RegCap">
              <el-input v-model="merchantForm.RegCap" placeholder="请输入注册资本"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="merchantForm.Notes" type="textarea" placeholder="请输入备注"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!merchantForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(merchantForm.FileId, merchantForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="merchantForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitMerchantForm">保存</el-button>
            <el-button @click="showMerchantDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看客商信息的对话框 -->
    <el-dialog v-model="showViewMerchantDialog" title="客商信息" width="80%" @close="resetMerchantForm">
      <el-form :model="merchantForm" label-width="150px" :rules="merchantRules" ref="merchantFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商编码" prop="MercCode" :required="true">
              <el-input v-model="merchantForm.MercCode" placeholder="请输入客商编码" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商缩写" prop="MercAbbr" :required="true">
              <el-input v-model="merchantForm.MercAbbr" placeholder="请输入客商缩写" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商简称" prop="ShortMerc" :required="true">
              <el-input v-model="merchantForm.ShortMerc" placeholder="请输入客商简称" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商名称" prop="Merc" :required="true">
              <el-input v-model="merchantForm.Merc" placeholder="请输入客商名称" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="英文名称" prop="EngName">
              <el-input v-model="merchantForm.EngName" placeholder="请输入英文名称" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地址" prop="Address">
              <el-input v-model="merchantForm.Address" placeholder="请输入地址" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="merchantForm.Nation" placeholder="请输入国家" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="merchantForm.PhoneNum" placeholder="请输入联系电话" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="merchantForm.Email" placeholder="请输入邮箱" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="传真" prop="Fax">
              <el-input v-model="merchantForm.Fax" placeholder="请输入传真" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="网站" prop="Website">
              <el-input v-model="merchantForm.Website" placeholder="请输入网站" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-input v-model="merchantForm.TaxType" placeholder="请输入税号类型" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="税号" prop="TaxCode">
              <el-input v-model="merchantForm.TaxCode" placeholder="请输入税号" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商类型" prop="MercType">
              <el-input v-model="merchantForm.MercType" placeholder="请输入客商类型" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册资本" prop="RegCap">
              <el-input v-model="merchantForm.RegCap" placeholder="请输入注册资本" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="merchantForm.Notes" type="textarea" placeholder="请输入备注" :readonly="true"></el-input>


            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!merchantForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(merchantForm.FileId, merchantForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="merchantForm.FileName" :readonly="true"></el-input>
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
import { useRouter } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus';
import axios from 'axios';

const router = useRouter();
const searchQuery = ref('');

// 文件上传相关
const uploadRef = ref(null);
const file = ref(null);

// 分页相关
const currentPage = ref(1);
const pageSize = 8;

// 菜单相关
const activeMenu = ref('1-3'); // 默认选中客商信息菜单

// 对话框显示状态
const showMerchantDialog = ref(false);
const showViewMerchantDialog = ref(false);

// 客商信息表单数据
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
  FileId: '',
});

// 表单验证规则
const merchantRules = {
  MercCode: [{ required: true, message: '请输入客商编码', trigger: 'blur' }],
  MercAbbr: [{ required: true, message: '请输入客商缩写', trigger: 'blur' }],
  ShortMerc: [{ required: true, message: '请输入客商简称', trigger: 'blur' }],
  Merc: [{ required: true, message: '请输入客商名称', trigger: 'blur' }],
};

// 表格数据
const merchantData = ref([]);

// 计算分页数据
const paginatedMerchantData = computed(() => {
  let filteredData = merchantData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(
      (item) =>
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
        item.Notes.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

// 获取客商信息数据
const fetchMerchantData = async () => {
  try {
    const response = await axios.get('/find/merchant');
    merchantData.value = response.data.Merchant;
  } catch (error) {
    console.error('获取客商信息失败:', error);
    ElMessage.error('获取客商信息失败，请稍后重试');
  }
};

// 文件选择事件
const handleFileChange = (uploadFile) => {
  file.value = uploadFile.raw;
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
  merchantForm.value = { ...row };
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
    FileId: '',
  };
  file.value = null;
  if (uploadRef.value) {
    uploadRef.value.clearFiles();
  }
};

// 提交表单
const submitMerchantForm = async () => {
  try {
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
      ElMessage.success('客商信息保存成功');
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
const handleDelete = (index, MercId) => {
  ElMessageBox.confirm('确定要删除该客商信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      axios
        .post('/delete/merchant', {
          MercId: MercId,
          MercCode: 'ss',
          MercAbbr: 'ss',
          Merc: 'ss',
          ShortMerc: 'ss',
        })
        .then((response) => {
          if (response.status === 200) {
            ElMessage.success('删除成功');
            fetchMerchantData();
          } else {
            ElMessage.error(response.data.RetMessage || '删除失败');
          }
        })
        .catch((error) => {
          ElMessage.error(error.response.data.RetMessage);
        });
    })
    .catch(() => {
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

// 菜单项选择事件
const handleMenuSelect = (index) => {
  activeMenu.value = index;
  if (index === '1-3') {
    fetchMerchantData();
  }
};

// 动态标题和按钮文本
const headerTitle = computed(() => {
  return activeMenu.value === '1-3' ? '客商信息' : '';
});

const addButtonText = computed(() => {
  return activeMenu.value === '1-3' ? '添加客商信息' : '';
});

// 初始化数据
onMounted(() => {
  fetchMerchantData();
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

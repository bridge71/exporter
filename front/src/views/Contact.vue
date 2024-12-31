<template>
<<<<<<< HEAD
  <div>
    <!-- 侧边栏菜单 -->
    <SideMenu />
    <!-- 主内容区域 -->
    <router-view />
  </div>
</template>

<script setup>
import SideMenu from './components/SideMenu.vue'; // 引入 SideMenu 组件
</script>
=======
    <div id="app">
      <el-container>
        <!-- 侧边栏 -->
        <el-aside width="200px">
          <div style="height: 100vh; overflow-y: auto;">
            <el-menu :default-active="activeMenu" class="el-menu-vertical-demo" @select="handleMenuSelect">
              <!-- 菜单项内容保持不变 -->
              <el-submenu index="1">
                <template #title>会计实体信息</template>
                <el-menu-item index="1-1" @click="pushAcct">会计实体信息</el-menu-item>
                <el-menu-item index="1-2" @click="pushAcctBank">会计实体银行账户信息</el-menu-item>
                <el-menu-item index="1-3" @click="pushClient">客商信息</el-menu-item>
                <el-menu-item index="1-4" @click="pushContact">联系人信息</el-menu-item>
                <el-menu-item index="1-5" @click="pushBankAccount">银行账户信息</el-menu-item>
                <el-menu-item index="1-6" @click="pushInventoryLocation">库存地点信息</el-menu-item>
                <el-menu-item index="1-7" @click="pushPaymentMethod">付款方式信息</el-menu-item>
                <el-menu-item index="1-8" @click="pushCategory">品类信息</el-menu-item>
                <el-menu-item index="1-9" @click="pushBrand">品牌信息</el-menu-item>
                <el-menu-item index="1-10" @click="pushPackagingSpec">包装规格信息</el-menu-item>
                <el-menu-item index="1-11" @click="pushEmployee">员工信息</el-menu-item>
                <el-menu-item index="1-12" @click="pushProductDetail">产品明细</el-menu-item>
                <el-menu-item index="1-13" @click="pushLoadingDetail">装货明细</el-menu-item>
                <el-menu-item index="1-14" @click="pushCostDetail">费用明细</el-menu-item>
                <el-menu-item index="1-15" @click="pushSalesOrder">销售订单</el-menu-item>
                <el-menu-item index="1-16" @click="pushPurchaseOrder">采购收单</el-menu-item>
                <el-menu-item index="1-17" @click="pushPayableInvoice">应付账款单</el-menu-item>
                <el-menu-item index="1-18" @click="pushReceipt">收款单</el-menu-item>
                <el-menu-item index="1-19" @click="pushPayment">付款单</el-menu-item>
                <!-- 其他菜单项 -->
              </el-submenu>
              <!-- 其他菜单 -->
            </el-menu>
          </div>
        </el-aside>
  
        <!-- 主体内容 -->
        <el-container>
          <el-header style="display: flex; justify-content: space-between; align-items: center;">
            <h2>{{ headerTitle }}</h2>
            <div>
              搜索：
              <el-input
                v-model="searchQuery"
                placeholder="输入要搜索的关键字"
                style="width: 200px;"
              />
              <el-button type="primary" @click="handleAdd">{{ addButtonText }}</el-button>
            </div>
          </el-header>
          <el-main>
            <!-- 联系人信息表格 -->
            <el-table :data="paginatedContactData" style="width: 100%" max-height="450" v-if="activeMenu === '1-4'">
              <el-table-column prop="Name" label="姓名" width="160%"></el-table-column>
              <el-table-column prop="male" label="性别" width="160%"></el-table-column>
              <el-table-column prop="Nation" label="国家" width="160%"></el-table-column>
              <el-table-column prop="Address" label="地址" width="220%"></el-table-column>
              <el-table-column prop="Email" label="邮箱" width="220%"></el-table-column>
              <el-table-column prop="PhoneNum" label="电话" width="220%"></el-table-column>
              <el-table-column prop="QQ" label="QQ" width="220%"></el-table-column>
              <el-table-column prop="wechat" label="微信" width="220%"></el-table-column>
              <el-table-column prop="Post" label="岗位" width="220%"></el-table-column>
  
              <el-table-column label="操作" fixed="right" width="200">
                <template #default="scope">
                  <el-row type="flex" justify="space-between">
                    <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                    <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                    <el-button @click="handleDelete(scope.$index, scope.row.ContactId)" type="text" size="small">删除</el-button>
                  </el-row>
                </template>
              </el-table-column>
            </el-table>
  
            <el-pagination v-model:current-page="currentPage" :page-size="pageSize"
              :total="activeMenu === '1-4' ? contactData.length : 0" layout="prev, pager, next"
              @current-change="handlePageChange" />
          </el-main>
        </el-container>
      </el-container>
  
      <!-- 添加联系人信息的对话框 -->
      <el-dialog v-model="showContactDialog" title="联系人信息" width="80%" @close="resetContactForm">
        <el-form :model="contactForm" label-width="150px" :rules="contactRules" ref="contactFormRef">
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="姓名" prop="Name" :required="true">
                <el-input v-model="contactForm.Name"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="性别" prop="male" :required="true">
                <el-select v-model="contactForm.male" placeholder="请选择">
                  <el-option label="男" value="男"></el-option>
                  <el-option label="女" value="女"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
  
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="国家" prop="Nation">
                <el-input v-model="contactForm.Nation"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="地址" prop="Address">
                <el-input v-model="contactForm.Address"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
  
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="邮箱" prop="Email">
                <el-input v-model="contactForm.Email"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="电话" prop="PhoneNum">
                <el-input v-model="contactForm.PhoneNum"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
  
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="QQ" prop="QQ">
                <el-input v-model="contactForm.QQ"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="微信" prop="wechat">
                <el-input v-model="contactForm.wechat"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
  
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="岗位" prop="Post">
                <el-input v-model="contactForm.Post"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
  
          <el-row :gutter="20">
            <el-col :span="24" style="text-align: right;">
              <el-button type="primary" @click="submitContactForm">保存</el-button>
              <el-button @click="showContactDialog = false">关闭</el-button>
            </el-col>
          </el-row>
        </el-form>
      </el-dialog>
  
      <!-- 查看联系人信息的对话框 -->
      <el-dialog v-model="showshowContactDialog" title="联系人信息" width="80%" @close="resetContactForm">
        <el-form :model="contactForm" label-width="150px" :rules="contactRules" ref="contactFormRef">
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="姓名" prop="Name" :required="true">
                <el-input v-model="contactForm.Name" :readonly="true"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="性别" prop="male" :required="true">
                <el-input v-model="contactForm.male" :readonly="true"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
  
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="国家" prop="Nation">
                <el-input v-model="contactForm.Nation" :readonly="true"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="地址" prop="Address">
                <el-input v-model="contactForm.Address" :readonly="true"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
  
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="邮箱" prop="Email">
                <el-input v-model="contactForm.Email" :readonly="true"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="电话" prop="PhoneNum">
                <el-input v-model="contactForm.PhoneNum" :readonly="true"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
  
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="QQ" prop="QQ">
                <el-input v-model="contactForm.QQ" :readonly="true"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="微信" prop="wechat">
                <el-input v-model="contactForm.wechat" :readonly="true"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
  
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="岗位" prop="Post">
                <el-input v-model="contactForm.Post" :readonly="true"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
  
          <el-row :gutter="20">
            <el-col :span="24" style="text-align: right;">
              <el-button @click="showshowContactDialog = false">关闭</el-button>
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

const pushAcct = () => {
  router.push('/acct');
};

const pushAcctBank = () => {
  router.push('/acctBank');
};

const pushClient = () => {
  router.push('/client');
};

const pushContact = () => {
  router.push('/contact');
};

const pushBankAccount = () => {
  router.push('/bankAccount');
};

const pushInventoryLocation = () => {
  router.push('/inventoryLocation');
};

const pushPaymentMethod = () => {
  router.push('/paymentMethod');
};

const pushCategory = () => {
  router.push('/category');
};

const pushBrand = () => {
  router.push('/brand');
};

const pushPackagingSpec = () => {
  router.push('/packagingSpec');
};

const pushEmployee = () => {
  router.push('/employee');
};

const pushProductDetail = () => {
  router.push('/productDetail');
};

const pushLoadingDetail = () => {
  router.push('/loadingDetail');
};

const pushCostDetail = () => {
  router.push('/costDetail');
};

const pushSalesOrder = () => {
  router.push('/salesOrder');
};

const pushPurchaseOrder = () => {
  router.push('/purchaseOrder');
};

const pushPayableInvoice = () => {
  router.push('/payableInvoice');
};

const pushReceipt = () => {
  router.push('/receipt');
};

const pushPayment = () => {
  router.push('/payment');
};

const downloadFile = async (fileId, fileName) => {
  try {
    const response = await axios.post(
      '/file',
      { FileId: fileId },
      {
        responseType: 'blob',
      }
    );

    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', `${fileName}` || `file_${fileId}`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error('文件下载失败:', error);
    ElMessage.error('文件下载失败，请稍后重试');
  }
};

const uploadRef = ref(null);
const file = ref(null);

const handleFileChange = (uploadFile) => {
  file.value = uploadFile.raw;
};

const handleView = (index, row) => {
  if (activeMenu.value === '1-4') {
    contactForm.value = { ...row };
    showshowContactDialog.value = true;
  }
};

const resetContactForm = () => {
  contactForm.value = {
    Name: '',
    male: '',
    Nation: '',
    Address: '',
    Email: '',
    PhoneNum: '',
    QQ: '',
    wechat: '',
    Post: ''
  };
  file.value = null;
  if (uploadRef.value) {
    uploadRef.value.clearFiles();
  }
};

const submitContactForm = async () => {
  try {
    const formData = new FormData();

    Object.keys(contactForm.value).forEach((key) => {
      formData.append(key, contactForm.value[key]);
    });

    if (file.value) {
      formData.append('file', file.value);
    }

    const response = await axios.post('/save/contact', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    if (response.status === 200) {
      ElMessage.success('联系人信息保存成功');
      showContactDialog.value = false;
      fetchContactData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    ElMessage.error(error.response.data.RetMessage);
  }
};

const handleDelete = (index, ContactId) => {
  ElMessageBox.confirm('确定要删除该联系人信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/contact', {
      ContactId: ContactId
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchContactData();
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

const handlePageChange = (page) => {
  currentPage.value = page;
};

const currentPage = ref(1);
const pageSize = 8;

const paginatedContactData = computed(() => {
  let filteredData = contactData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.Name.includes(searchQuery.value) ||
      item.male.includes(searchQuery.value) ||
      item.Nation.includes(searchQuery.value) ||
      item.Address.includes(searchQuery.value) ||
      item.Email.includes(searchQuery.value) ||
      item.PhoneNum.includes(searchQuery.value) ||
      item.QQ.includes(searchQuery.value) ||
      item.wechat.includes(searchQuery.value) ||
      item.Post.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

onMounted(() => {
  fetchContactData();
});

const fetchContactData = async () => {
  try {
    const response = await axios.get('/find/contact');
    contactData.value = response.data.Contact;
  } catch (error) {
    console.error('获取联系人信息失败:', error);
    ElMessage.error('获取联系人信息失败，请稍后重试');
  }
};

const activeMenu = ref('1-4');

const showContactDialog = ref(false);
const showshowContactDialog = ref(false);

const contactForm = ref({
  Name: '',
  male: '',
  Nation: '',
  Address: '',
  Email: '',
  PhoneNum: '',
  QQ: '',
  wechat: '',
  Post: ''
});

const contactRules = {
  Name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
  male: [{ required: true, message: '请选择性别', trigger: 'blur' }]
};

const contactData = ref([]);

const headerTitle = computed(() => {
  return activeMenu.value === '1-4' ? '联系人信息' : '';
});

const addButtonText = computed(() => {
  return activeMenu.value === '1-4' ? '添加联系人信息' : '';
});

const handleMenuSelect = (index) => {
  activeMenu.value = index;
  if (index === '1-4') {
    fetchContactData();
  }
};

const handleAdd = () => {
  showContactDialog.value = true;
};

const handleEdit = (index, row) => {
  contactForm.value = { ...row };
  showContactDialog.value = true;
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
>>>>>>> 593b97d905b86d5274251b6be73a3d8c8270a1cc

<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-7'" />
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
          <!-- 支付方式信息表格 -->
          <el-table :data="paginatedPayMentMethodData" style="width: 100%" max-height="450">
            <el-table-column prop="PayMtdName" label="支付方式名称" width="220%"></el-table-column>
            <el-table-column prop="PayRat" label="支付费率" width="220%"></el-table-column>
            <el-table-column prop="PayMth" label="后付款转账方式" width="220%"></el-table-column>
            <el-table-column prop="PayLimit" label="后付款转账期限" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="320%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.PayMentMethodId)" type="text"
                    size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="payMentMethodData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加支付方式信息的对话框 -->
    <el-dialog v-model="showPayMentMethodDialog" title="支付方式信息" width="80%" @close="resetPayMentMethodForm">
      <el-form :model="payMentMethodForm" label-width="150px" :rules="payMentMethodRules" ref="payMentMethodFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="支付方式名称" prop="PayMtdName">
              <el-input v-model="payMentMethodForm.PayMtdName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="支付费率" prop="PayRat">
              <el-input v-model="payMentMethodForm.PayRat"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="后付款转账方式" prop="PayMth">
              <el-select v-model="payMentMethodForm.PayMth" @change="onPayMthChange" placeholder="请选择后付款转账方式">
                <el-option v-for="payMth in payMthData" :key="payMth.PayMth" :label="payMth.PayMth"
                  :value="payMth.PayMth"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="后付款转账期限" prop="PayLimit">
              <el-select v-model="payMentMethodForm.PayLimit" @change="onPayLimitChange" placeholder="请选择后付款转账期限">
                <el-option v-for="payLimit in payLimitData" :key="payLimit.PayLimit" :label="payLimit.PayLimit"
                  :value="payLimit.PayLimit"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="payMentMethodForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitPayMentMethodForm">保存</el-button>
            <el-button @click="showPayMentMethodDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看支付方式信息的对话框 -->
    <el-dialog v-model="showViewPayMentMethodDialog" title="支付方式信息" width="80%" @close="resetPayMentMethodForm">
      <el-form :model="payMentMethodForm" label-width="150px" :rules="payMentMethodRules" ref="payMentMethodFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="支付方式名称" prop="PayMtdName">
              <el-input v-model="payMentMethodForm.PayMtdName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="支付费率" prop="PayRat">
              <el-input v-model="payMentMethodForm.PayRat" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="后付款转账方式" prop="PayMth">
              <el-select v-model="payMentMethodForm.PayMth" @change="onPayMthChange" placeholder="请选择后付款转账方式"
                :disabled="true">
                <el-option v-for="payMth in payMthData" :key="payMth.PayMth" :label="payMth.PayMthName"
                  :value="payMth.PayMth"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="后付款转账期限" prop="PayLimit">
              <el-select v-model="payMentMethodForm.PayLimit" @change="onPayLimitChange" placeholder="请选择后付款转账期限"
                :disabled="true">
                <el-option v-for="payLimit in payLimitData" :key="payLimit.PayLimit" :label="payLimit.PayLimitName"
                  :value="payLimit.PayLimit"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="payMentMethodForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewPayMentMethodDialog = false">关闭</el-button>
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
const payMentMethodData = ref([]);
const payMthData = ref([]); // 后付款转账方式数据
const payLimitData = ref([]); // 后付款转账期限数据
const showPayMentMethodDialog = ref(false);
const showViewPayMentMethodDialog = ref(false);
const payMentMethodForm = ref({
  PayMtdName: '',
  PayRat: '',
  PayMth: '',
  PayLimit: '',
  Notes: '',
  PayMentMethodId: '',
});


const payMentMethodFormRef = ref(null);
// 自定义验证函数
const validateNotEmpty = (rule, value, callback) => {
  if (value === '' || value === null || value === undefined) {
    callback(new Error(rule.message));  // 使用 rule.message 作为错误消息
  } else {
    callback();  // 验证通过
  }
};

// 支付方式表单验证规则
const payMentMethodRules = {
  PayMtdName: [
    { required: true, validator: validateNotEmpty, message: '请输入支付方式名称', trigger: 'blur' }
  ],
  PayRat: [
    { required: true, validator: validateNotEmpty, message: '请输入支付费率', trigger: 'blur' }
  ],
  PayMth: [
    { required: true, validator: validateNotEmpty, message: '请选择后付款转账方式', trigger: 'blur' }
  ],
  PayLimit: [
    { required: true, validator: validateNotEmpty, message: '请选择后付款转账期限', trigger: 'blur' }
  ]
};
const handlePageChange = (page) => {
  currentPage.value = page;
};

const paginatedPayMentMethodData = computed(() => {
  let filteredData = payMentMethodData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.PayMtdName.includes(searchQuery.value) ||
      item.PayRat.toString().includes(searchQuery.value) ||
      item.PayMthName.includes(searchQuery.value) ||
      item.PayLimitName.includes(searchQuery.value) ||
      item.Notes.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

const handleAdd = () => {
  showPayMentMethodDialog.value = true;
};

const handleEdit = (index, row) => {
  payMentMethodForm.value = { ...row };
  showPayMentMethodDialog.value = true;
};

const handleView = (index, row) => {
  payMentMethodForm.value = { ...row };
  showViewPayMentMethodDialog.value = true;
};

const handleDelete = (index, PayMentMethodId) => {
  ElMessageBox.confirm('确定要删除该支付方式信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/payMentMethod', { PayMentMethodId })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchPayMentMethodData();
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

const resetPayMentMethodForm = () => {
  payMentMethodForm.value = {
    PayMtdName: '',
    PayRat: '',
    PayMth: '',
    PayLimit: '',
    Notes: '',
    PayMentMethodId: '',
  };
};

const onPayMthChange = (value) => {
  const selectedPayMth = payMthData.value.find(payMth => payMth.PayMth === value);
  if (selectedPayMth) {
    payMentMethodForm.value.PayMth = selectedPayMth.PayMth;
  }
};

const onPayLimitChange = (value) => {
  const selectedPayLimit = payLimitData.value.find(payLimit => payLimit.PayLimit === value);
  if (selectedPayLimit) {
    payMentMethodForm.value.PayLimit = selectedPayLimit.PayLimit;
  }
};

const submitPayMentMethodForm = async () => {
  try {

    const isValid = await payMentMethodFormRef.value.validate();
    if (!isValid) {
      ElMessage.error('请填写所有必填字段');
      console.log('验证不通过');
      return; // 如果验证不通过，阻止提交
    }

    payMentMethodForm.value.PayRat = parseInt(payMentMethodForm.value.PayRat, 10)
    payMentMethodForm.value.PayMentMethodId = parseInt(payMentMethodForm.value.PayMentMethodId, 10)
    const response = await axios.post('/save/payMentMethod', payMentMethodForm.value);
    if (response.status === 200) {
      ElMessage.success('支付方式信息保存成功');
      showPayMentMethodDialog.value = false;
      fetchPayMentMethodData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存支付方式信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

onMounted(() => {
  fetchPayMentMethodData();
  fetchPayMthData();
  fetchPayLimitData();
});

const fetchPayMentMethodData = async () => {
  try {
    const response = await axios.get('/find/payMentMethod');
    payMentMethodData.value = response.data.PayMentMethod;
  } catch (error) {
    console.error('获取支付方式信息失败:', error);
    ElMessage.error('获取支付方式信息失败，请稍后重试');
  }
};

const fetchPayMthData = async () => {
  try {
    const response = await axios.get('/find/payMth');
    payMthData.value = response.data.PayMth;
  } catch (error) {
    console.error('获取后付款转账方式失败:', error);
    ElMessage.error('获取后付款转账方式失败，请稍后重试');
  }
};

const fetchPayLimitData = async () => {
  try {
    const response = await axios.get('/find/payLimit');
    payLimitData.value = response.data.PayLimit;
  } catch (error) {
    console.error('获取后付款转账期限失败:', error);
    ElMessage.error('获取后付款转账期限失败，请稍后重试');
  }
};

const headerTitle = computed(() => '支付方式信息');
const addButtonText = computed(() => '添加支付方式信息');
</script>

<style src="../assets/styles/Bottom.css"></style>

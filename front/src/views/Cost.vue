<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-14'" />
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
          <!-- 费用信息表格 -->
          <el-table :data="paginatedCostData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="ExpType" label="费用类型" width="220%"></el-table-column>
            <el-table-column prop="Rates" label="收费标准" width="220%"></el-table-column>
            <el-table-column prop="UnitPrice" label="单价" width="220%"></el-table-column>
            <el-table-column prop="Number" label="数量" width="220%"></el-table-column>
            <el-table-column prop="Amount" label="金额" width="220%"></el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="CreatedAt" label="创建时间" width="240%"></el-table-column>
            <el-table-column prop="UpdatedAt" label="更新时间" width="240%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row)" type="text" size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="costData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加费用信息的对话框 -->
    <el-dialog v-model="showCostDialog" title="费用信息" width="80%" @close="resetCostForm">
      <el-form :model="costForm" label-width="150px" :rules="costRules" ref="costFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="费用类型" prop="ExpType">
              <el-select v-model="costForm.ExpType" placeholder="请选择费用类型">
                <el-option v-for="item in expTypeData" :key="item.ExpType" :label="item.ExpType"
                  :value="item.ExpType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="收费标准" prop="Rates">
              <el-select v-model="costForm.Rates" placeholder="请选择收费标准">
                <el-option v-for="item in ratesData" :key="item.Rates" :label="item.Rates"
                  :value="item.Rates"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单价" prop="UnitPrice">
              <el-input v-model="costForm.UnitPrice" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="数量" prop="Number">
              <el-input v-model="costForm.Number" type="number"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="金额" prop="Amount">
              <el-input v-model="costForm.Amount" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-select v-model="costForm.Currency" placeholder="请选择币种">
                <el-option v-for="item in currencyData" :key="item.Currency" :label="item.Currency"
                  :value="item.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitCostForm">保存</el-button>
            <el-button @click="showCostDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看费用信息的对话框 -->
    <el-dialog v-model="showViewCostDialog" title="费用信息" width="80%" @close="resetCostForm">
      <el-form :model="costForm" label-width="150px" :rules="costRules" ref="costFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="费用类型" prop="ExpType">
              <el-input v-model="costForm.ExpType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="收费标准" prop="Rates">
              <el-input v-model="costForm.Rates" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单价" prop="UnitPrice">
              <el-input v-model="costForm.UnitPrice" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="数量" prop="Number">
              <el-input v-model="costForm.Number" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="金额" prop="Amount">
              <el-input v-model="costForm.Amount" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-input v-model="costForm.Currency" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="创建时间" prop="CreatedAt">
              <el-input v-model="costForm.CreatedAt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="更新时间" prop="UpdatedAt">
              <el-input v-model="costForm.UpdatedAt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewCostDialog = false">关闭</el-button>
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
const costData = ref([]);
const showCostDialog = ref(false);
const showViewCostDialog = ref(false);
const costForm = ref({
  ExpType: '',
  Rates: '',
  UnitPrice: 0,
  Number: 0,
  Amount: 0,
  Currency: '',
  CostId: '',
});
const costFormRef = ref(null);

// 字典数据
const expTypeData = ref([]); // 费用类型
const ratesData = ref([]); // 收费标准
const currencyData = ref([]); // 币种

const handlePageChange = (page) => {
  currentPage.value = page;
};

const paginatedCostData = computed(() => {
  let filteredData = costData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.ExpType.includes(searchQuery.value) ||
      item.Rates.includes(searchQuery.value) ||
      item.UnitPrice.toString().includes(searchQuery.value) ||
      item.Number.toString().includes(searchQuery.value) ||
      item.Amount.toString().includes(searchQuery.value) ||
      item.Currency.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

const handleAdd = () => {
  showCostDialog.value = true;
};

const handleEdit = (index, row) => {
  costForm.value = { ...row };
  showCostDialog.value = true;
};

const handleView = (index, row) => {
  costForm.value = { ...row };
  showViewCostDialog.value = true;
};

const handleDelete = (index, row) => {
  ElMessageBox.confirm('确定要删除该费用信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    costForm.value = { ...row };
    axios.post('/delete/costInfo', costForm.value)
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchCostData();
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

const resetCostForm = () => {
  costForm.value = {
    ExpType: '',
    Rates: '',
    UnitPrice: 0,
    Number: 0,
    Amount: 0,
    Currency: '',
    CostId: '',
  };
};

const submitCostForm = async () => {
  try {

    costForm.value.Amount = parseInt(costForm.value.Amount, 10);
    costForm.value.Number = parseInt(costForm.value.Number, 10);
    costForm.value.UnitPrice = parseInt(costForm.value.UnitPrice, 10);
    const response = await axios.post('/save/costInfo', costForm.value);
    if (response.status === 200) {
      ElMessage.success('费用信息保存成功');
      showCostDialog.value = fetchCostData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存费用信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 获取字典数据
const fetchDictionaryData = async () => {
  try {
    const [
      expTypeResponse,
      ratesResponse,
      currencyResponse,
    ] = await Promise.all([
      axios.get('/find/expType'), // 费用类型
      axios.get('/find/rates'), // 收费标准
      axios.get('/find/currency'), // 币种
    ]);

    // 提取字典数据
    expTypeData.value = expTypeResponse.data.ExpType || [];
    ratesData.value = ratesResponse.data.Rates || [];
    currencyData.value = currencyResponse.data.Currency || [];
  } catch (error) {
    console.error('获取字典数据失败:', error);
    ElMessage.error('获取字典数据失败，请稍后重试');
  }
};

// 获取费用数据
const fetchCostData = async () => {
  try {
    const response = await axios.get('/find/costInfo');
    costData.value = response.data.CostInfo || [];
    console.log(response.data.CostInfo)
  } catch (error) {
    console.error('获取费用信息失败:', error);
    ElMessage.error('获取费用信息失败，请稍后重试');
  }
};

onMounted(() => {
  fetchCostData();
  fetchDictionaryData();
});

const headerTitle = computed(() => '费用信息');
const addButtonText = computed(() => '添加费用信息');

const costRules = {
  ExpType: [{ required: true, message: '请选择费用类型', trigger: 'blur' }],
  Rates: [{ required: true, message: '请选择收费标准', trigger: 'blur' }],
  UnitPrice: [{ required: true, message: '请输入单价', trigger: 'blur' }],
  Number: [{ required: true, message: '请输入数量', trigger: 'blur' }],
  Amount: [{ required: true, message: '请输入金额', trigger: 'blur' }],
  Currency: [{ required: true, message: '请选择币种', trigger: 'blur' }],
};
</script>

<style src="../assets/styles/Bottom.css"></style>

<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-12'" />
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
          <!-- 产品信息表格 -->
          <el-table :data="paginatedPrdtData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="CatEngName" label="产品" width="220%"></el-table-column>
            <el-table-column prop="BrandEngName" label="品牌" width="220%"></el-table-column>
            <el-table-column prop="Factory" label="生产工厂" width="220%"></el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="UnitPrice" label="单价" width="220%"></el-table-column>
            <el-table-column prop="Unit" label="单位" width="220%"></el-table-column>
            <el-table-column prop="Amount" label="金额" width="220%"></el-table-column>
            <el-table-column prop="ItemNum" label="件数" width="220%"></el-table-column>
            <el-table-column prop="PackSpec" label="包装规格" width="220%"></el-table-column>
            <el-table-column prop="Weight" label="重量" width="220%"></el-table-column>
            <el-table-column prop="WeightUnit" label="重量单位" width="220%"></el-table-column>
            <el-table-column prop="TradeTerm" label="贸易条款" width="220%"></el-table-column>
            <el-table-column prop="DeliveryLoc" label="交货地点" width="220%"></el-table-column>
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

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="prdtData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加产品信息的对话框 -->
    <el-dialog v-model="showPrdtDialog" title="产品信息" width="80%" @close="resetPrdtForm">
      <el-form :model="prdtForm" label-width="150px" :rules="prdtRules" ref="prdtFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="产品" prop="CatEngName">
              <el-select v-model="prdtForm.CatEngName" placeholder="请选择产品">
                <el-option v-for="item in prdtTypeData" :key="item.PrdtType" :label="item.PrdtType"
                  :value="item.PrdtType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌" prop="BrandEngName">
              <el-select v-model="prdtForm.BrandEngName" placeholder="请选择品牌">
                <el-option v-for="item in brandTypeData" :key="item.BrandType" :label="item.BrandType"
                  :value="item.BrandType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="生产工厂" prop="Factory">
              <el-select v-model="prdtForm.Factory" placeholder="请选择生产工厂">
                <el-option v-for="item in suprTypeData" :key="item.SuprType" :label="item.SuprType"
                  :value="item.SuprType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-select v-model="prdtForm.Currency" placeholder="请选择币种">
                <el-option v-for="item in currencyData" :key="item.Currency" :label="item.Currency"
                  :value="item.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单价" prop="UnitPrice">
              <el-input v-model="prdtForm.UnitPrice" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="Unit">
              <el-select v-model="prdtForm.Unit" placeholder="请选择单位">
                <el-option v-for="item in unitMeasData" :key="item.UnitMeas" :label="item.UnitMeas"
                  :value="item.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="金额" prop="Amount">
              <el-input v-model="prdtForm.Amount" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="件数" prop="ItemNum">
              <el-input v-model="prdtForm.ItemNum" type="number"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpec">
              <el-select v-model="prdtForm.PackSpec" placeholder="请选择包装规格">
                <el-option v-for="item in packTypeData" :key="item.PackType" :label="item.PackType"
                  :value="item.PackType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="重量" prop="Weight">
              <el-input v-model="prdtForm.Weight" type="number"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="重量单位" prop="WeightUnit">
              <el-select v-model="prdtForm.WeightUnit" placeholder="请选择重量单位">
                <el-option v-for="item in unitMeasData" :key="item.UnitMeas" :label="item.UnitMeas"
                  :value="item.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="贸易条款" prop="TradeTerm">
              <el-select v-model="prdtForm.TradeTerm" placeholder="请选择贸易条款">
                <el-option v-for="item in tradeTermData" :key="item.TradeTerm" :label="item.TradeTerm"
                  :value="item.TradeTerm"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="交货地点" prop="DeliveryLoc">
              <el-select v-model="prdtForm.DeliveryLoc" placeholder="请选择交货地点">
                <el-option v-for="item in portData" :key="item.Port" :label="item.Port" :value="item.Port"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitPrdtForm">保存</el-button>
            <el-button @click="showPrdtDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看产品信息的对话框 -->
    <el-dialog v-model="showViewPrdtDialog" title="产品信息" width="80%" @close="resetPrdtForm">
      <el-form :model="prdtForm" label-width="150px" :rules="prdtRules" ref="prdtFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="产品" prop="CatEngName">
              <el-input v-model="prdtForm.CatEngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌" prop="BrandEngName">
              <el-input v-model="prdtForm.BrandEngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="生产工厂" prop="Factory">
              <el-input v-model="prdtForm.Factory" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-input v-model="prdtForm.Currency" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单价" prop="UnitPrice">
              <el-input v-model="prdtForm.UnitPrice" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="Unit">
              <el-input v-model="prdtForm.Unit" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="金额" prop="Amount">
              <el-input v-model="prdtForm.Amount" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="件数" prop="ItemNum">
              <el-input v-model="prdtForm.ItemNum" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpec">
              <el-input v-model="prdtForm.PackSpec" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="重量" prop="Weight">
              <el-input v-model="prdtForm.Weight" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="重量单位" prop="WeightUnit">
              <el-input v-model="prdtForm.WeightUnit" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="贸易条款" prop="TradeTerm">
              <el-input v-model="prdtForm.TradeTerm" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="交货地点" prop="DeliveryLoc">
              <el-input v-model="prdtForm.DeliveryLoc" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="创建时间" prop="CreatedAt">
              <el-input v-model="prdtForm.CreatedAt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="更新时间" prop="UpdatedAt">
              <el-input v-model="prdtForm.UpdatedAt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewPrdtDialog = false">关闭</el-button>
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
const prdtData = ref([]);
const showPrdtDialog = ref(false);
const showViewPrdtDialog = ref(false);
const prdtForm = ref({
  CatEngName: '',
  BrandEngName: '',
  Factory: '',
  Currency: '',
  UnitPrice: 0,
  Unit: '',
  Amount: 0,
  ItemNum: 0,
  PackSpec: '',
  Weight: 0,
  WeightUnit: '',
  TradeTerm: '',
  DeliveryLoc: '',
  PrdtId: '',
});
const prdtFormRef = ref(null);

// 字典数据
const prdtTypeData = ref([]); // 产品类型
const brandTypeData = ref([]); // 品牌类型
const suprTypeData = ref([]); // 供应商类型
const currencyData = ref([]); // 币种
const unitMeasData = ref([]); // 单位
const packTypeData = ref([]); // 包装规格
const tradeTermData = ref([]); // 贸易条款
const portData = ref([]); // 交货地点

const handlePageChange = (page) => {
  currentPage.value = page;
};

const paginatedPrdtData = computed(() => {
  let filteredData = prdtData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.CatEngName.includes(searchQuery.value) ||
      item.BrandEngName.includes(searchQuery.value) ||
      item.Factory.includes(searchQuery.value) ||
      item.Currency.includes(searchQuery.value) ||
      item.UnitPrice.toString().includes(searchQuery.value) ||
      item.Unit.includes(searchQuery.value) ||
      item.Amount.toString().includes(searchQuery.value) ||
      item.ItemNum.toString().includes(searchQuery.value) ||
      item.PackSpec.includes(searchQuery.value) ||
      item.ID.toString().includes(searchQuery.value) ||
      item.Weight.toString().includes(searchQuery.value) ||
      item.WeightUnit.includes(searchQuery.value) ||
      item.TradeTerm.includes(searchQuery.value) ||
      item.DeliveryLoc.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

const handleAdd = () => {
  showPrdtDialog.value = true;
};

const handleEdit = (index, row) => {
  prdtForm.value = { ...row };
  showPrdtDialog.value = true;
};

const handleView = (index, row) => {
  prdtForm.value = { ...row };
  showViewPrdtDialog.value = true;
};

const handleDelete = (index, row) => {
  ElMessageBox.confirm('确定要删除该产品信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {

    prdtForm.value = { ...row };
    axios.post('/delete/prdtInfo', prdtForm.value)
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchPrdtData();
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

const resetPrdtForm = () => {
  prdtForm.value = {
    CatEngName: '',
    BrandEngName: '',
    Factory: '',
    Currency: '',
    UnitPrice: 0,
    Unit: '',
    Amount: 0,
    ItemNum: 0,
    PackSpec: '',
    Weight: 0,
    WeightUnit: '',
    TradeTerm: '',
    DeliveryLoc: '',
    PrdtId: '',
  };
};

const submitPrdtForm = async () => {
  try {
    console.log(prdtForm.value)

    prdtForm.value.Weight = parseInt(prdtForm.value.Weight, 10);
    prdtForm.value.UnitPrice = parseInt(prdtForm.value.UnitPrice, 10);
    prdtForm.value.Amount = parseInt(prdtForm.value.Amount, 10);
    prdtForm.value.ItemNum = parseInt(prdtForm.value.ItemNum, 10);
    const response = await axios.post('/save/prdtInfo', prdtForm.value);
    if (response.status === 200) {
      ElMessage.success('产品信息保存成功');
      await fetchPrdtData(); // 等待数据获取完成
      showPrdtDialog.value = false; // 手动关闭对话框
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存产品信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 获取字典数据
const fetchDictionaryData = async () => {
  try {
    const [
      prdtTypeResponse,
      brandTypeResponse,
      suprTypeResponse,
      currencyResponse,
      unitMeasResponse,
      packTypeResponse,
      tradeTermResponse,
      portResponse,
    ] = await Promise.all([
      axios.get('/find/prdtType'), // 产品类型
      axios.get('/find/brandType'), // 品牌类型
      axios.get('/find/suprType'), // 供应商类型
      axios.get('/find/currency'), // 币种
      axios.get('/find/unitMeas'), // 单位
      axios.get('/find/packType'), // 包装规格
      axios.get('/find/tradeTerm'), // 贸易条款
      axios.get('/find/port'), // 交货地点
    ]);

    // 提取字典数据
    prdtTypeData.value = prdtTypeResponse.data.PrdtType || [];
    brandTypeData.value = brandTypeResponse.data.BrandType || [];
    suprTypeData.value = suprTypeResponse.data.SuprType || [];
    currencyData.value = currencyResponse.data.Currency || [];
    unitMeasData.value = unitMeasResponse.data.UnitMeas || [];
    packTypeData.value = packTypeResponse.data.PackType || [];
    tradeTermData.value = tradeTermResponse.data.TradeTerm || [];
    portData.value = portResponse.data.Port || [];
  } catch (error) {
    console.error('获取字典数据失败:', error);
    ElMessage.error('获取字典数据失败，请稍后重试');
  }
};

// 获取产品数据
const fetchPrdtData = async () => {
  try {
    const response = await axios.get('/find/prdtInfo');
    prdtData.value = response.data.PrdtInfo || [];
    console.log(response.data.PrdtInfo)
  } catch (error) {
    console.error('获取产品信息失败:', error);
    ElMessage.error('获取产品信息失败，请稍后重试');
  }
};

onMounted(() => {
  fetchPrdtData();
  fetchDictionaryData();
});

const headerTitle = computed(() => '产品信息');
const addButtonText = computed(() => '添加产品信息');

const prdtRules = {
  CatEngName: [{ required: true, message: '请选择产品', trigger: 'blur' }],
  BrandEngName: [{ required: true, message: '请选择品牌', trigger: 'blur' }],
  Factory: [{ required: true, message: '请选择生产工厂', trigger: 'blur' }],
  Currency: [{ required: true, message: '请选择币种', trigger: 'blur' }],
  UnitPrice: [{ required: true, message: '请输入单价', trigger: 'blur' }],
  Unit: [{ required: true, message: '请选择单位', trigger: 'blur' }],
  Amount: [{ required: true, message: '请输入金额', trigger: 'blur' }],
  ItemNum: [{ required: true, message: '请输入件数', trigger: 'blur' }],
  PackSpec: [{ required: true, message: '请选择包装规格', trigger: 'blur' }],
  Weight: [{ required: true, message: '请输入重量', trigger: 'blur' }],
  WeightUnit: [{ required: true, message: '请选择重量单位', trigger: 'blur' }],
  TradeTerm: [{ required: true, message: '请选择贸易条款', trigger: 'blur' }],
  DeliveryLoc: [{ required: true, message: '请选择交货地点', trigger: 'blur' }],
};
</script>

<style src="../assets/styles/Bottom.css"></style>

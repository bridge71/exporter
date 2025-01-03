<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-13'" />
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
          <!-- 装货信息表格 -->
          <el-table :data="paginatedLoadingData" style="width: 100%" max-height="450">
            <el-table-column prop="Product" label="产品" width="220%"></el-table-column>
            <el-table-column prop="Brand" label="品牌" width="220%"></el-table-column>
            <el-table-column prop="PrdtPlant" label="生产工厂" width="220%"></el-table-column>
            <el-table-column prop="BatNum" label="批次号" width="220%"></el-table-column>
            <el-table-column prop="ItemNum" label="件数" width="220%"></el-table-column>
            <el-table-column prop="PackSpec" label="包装规格" width="220%"></el-table-column>
            <el-table-column prop="NetWeight" label="净重量" width="220%"></el-table-column>
            <el-table-column prop="Unit" label="单位" width="220%"></el-table-column>
            <el-table-column prop="CnrNum" label="集装箱号" width="220%"></el-table-column>
            <el-table-column prop="SealNum" label="铅封号" width="220%"></el-table-column>
            <el-table-column prop="VehNum" label="车辆号" width="220%"></el-table-column>
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

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="loadingData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加装货信息的对话框 -->
    <el-dialog v-model="showLoadingDialog" title="装货信息" width="80%" @close="resetLoadingForm">
      <el-form :model="loadingForm" label-width="150px" :rules="loadingRules" ref="loadingFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="产品" prop="Product">
              <el-select v-model="loadingForm.Product" placeholder="请选择产品">
                <el-option v-for="item in prdtTypeData" :key="item.PrdtType" :label="item.PrdtType"
                  :value="item.PrdtType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌" prop="Brand">
              <el-select v-model="loadingForm.Brand" placeholder="请选择品牌">
                <el-option v-for="item in brandTypeData" :key="item.BrandType" :label="item.BrandType"
                  :value="item.BrandType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="生产工厂" prop="PrdtPlant">
              <el-select v-model="loadingForm.PrdtPlant" placeholder="请选择生产工厂">
                <el-option v-for="item in suprTypeData" :key="item.SuprType" :label="item.SuprType"
                  :value="item.SuprType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="批次号" prop="BatNum">
              <el-input v-model="loadingForm.BatNum"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="件数" prop="ItemNum">
              <el-input v-model="loadingForm.ItemNum" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpec">
              <el-select v-model="loadingForm.PackSpec" placeholder="请选择包装规格">
                <el-option v-for="item in packTypeData" :key="item.PackType" :label="item.PackType"
                  :value="item.PackType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="净重量" prop="NetWeight">
              <el-input v-model="loadingForm.NetWeight" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="Unit">
              <el-select v-model="loadingForm.Unit" placeholder="请选择单位">
                <el-option v-for="item in unitMeasData" :key="item.UnitMeas" :label="item.UnitMeas"
                  :value="item.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="集装箱号" prop="CnrNum">
              <el-input v-model="loadingForm.CnrNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="铅封号" prop="SealNum">
              <el-input v-model="loadingForm.SealNum"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="车辆号" prop="VehNum">
              <el-input v-model="loadingForm.VehNum"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitLoadingForm">保存</el-button>
            <el-button @click="showLoadingDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看装货信息的对话框 -->
    <el-dialog v-model="showViewLoadingDialog" title="装货信息" width="80%" @close="resetLoadingForm">
      <el-form :model="loadingForm" label-width="150px" :rules="loadingRules" ref="loadingFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="产品" prop="Product">
              <el-input v-model="loadingForm.Product" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌" prop="Brand">
              <el-input v-model="loadingForm.Brand" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="生产工厂" prop="PrdtPlant">
              <el-input v-model="loadingForm.PrdtPlant" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="批次号" prop="BatNum">
              <el-input v-model="loadingForm.BatNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="件数" prop="ItemNum">
              <el-input v-model="loadingForm.ItemNum" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpec">
              <el-input v-model="loadingForm.PackSpec" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="净重量" prop="NetWeight">
              <el-input v-model="loadingForm.NetWeight" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="Unit">
              <el-input v-model="loadingForm.Unit" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="集装箱号" prop="CnrNum">
              <el-input v-model="loadingForm.CnrNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="铅封号" prop="SealNum">
              <el-input v-model="loadingForm.SealNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="车辆号" prop="VehNum">
              <el-input v-model="loadingForm.VehNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="创建时间" prop="CreatedAt">
              <el-input v-model="loadingForm.CreatedAt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="更新时间" prop="UpdatedAt">
              <el-input v-model="loadingForm.UpdatedAt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewLoadingDialog = false">关闭</el-button>
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
const loadingData = ref([]);
const showLoadingDialog = ref(false);
const showViewLoadingDialog = ref(false);
const loadingForm = ref({
  Product: '',
  Brand: '',
  PrdtPlant: '',
  BatNum: '',
  ItemNum: 0,
  PackSpec: '',
  NetWeight: 0,
  Unit: '',
  CnrNum: '',
  SealNum: '',
  VehNum: '',
  LoadingInfoId: '',
});
const loadingFormRef = ref(null);

// 字典数据
const prdtTypeData = ref([]); // 产品类型
const brandTypeData = ref([]); // 品牌类型
const suprTypeData = ref([]); // 供应商类型
const unitMeasData = ref([]); // 单位
const packTypeData = ref([]); // 包装规格

const handlePageChange = (page) => {
  currentPage.value = page;
};

const paginatedLoadingData = computed(() => {
  let filteredData = loadingData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.Product.includes(searchQuery.value) ||
      item.Brand.includes(searchQuery.value) ||
      item.PrdtPlant.includes(searchQuery.value) ||
      item.BatNum.includes(searchQuery.value) ||
      item.ItemNum.toString().includes(searchQuery.value) ||
      item.PackSpec.includes(searchQuery.value) ||
      item.NetWeight.toString().includes(searchQuery.value) ||
      item.Unit.includes(searchQuery.value) ||
      item.CnrNum.includes(searchQuery.value) ||
      item.SealNum.includes(searchQuery.value) ||
      item.VehNum.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

const handleAdd = () => {
  showLoadingDialog.value = true;
};

const handleEdit = (index, row) => {
  loadingForm.value = { ...row };
  showLoadingDialog.value = true;
};

const handleView = (index, row) => {
  loadingForm.value = { ...row };
  showViewLoadingDialog.value = true;
};

const handleDelete = (index, row) => {
  ElMessageBox.confirm('确定要删除该装货信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {

    loadingForm.value = { ...row };
    axios.post('/delete/loadingInfo', loadingForm.value)
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchLoadingData();
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

const resetLoadingForm = () => {
  loadingForm.value = {
    Product: '',
    Brand: '',
    PrdtPlant: '',
    BatNum: '',
    ItemNum: 0,
    PackSpec: '',
    NetWeight: 0,
    Unit: '',
    CnrNum: '',
    SealNum: '',
    VehNum: '',
    LoadingInfoId: '',
    CreatedAt: '', // 添加 CreatedAt
    UpdatedAt: '', // 添加 UpdatedAt
  };
};

const submitLoadingForm = async () => {
  try {

    loadingForm.value.ItemNum = parseInt(loadingForm.value.ItemNum, 10);
    loadingForm.value.NetWeight = parseInt(loadingForm.value.NetWeight, 10);
    const response = await axios.post('/save/loadingInfo', loadingForm.value);
    if (response.status === 200) {
      ElMessage.success('装货信息保存成功');
      showLoadingDialog.value = fetchLoadingData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存装货信息失败:', error);
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
      unitMeasResponse,
      packTypeResponse,
    ] = await Promise.all([
      axios.get('/find/prdtType'), // 产品类型
      axios.get('/find/brandType'), // 品牌类型
      axios.get('/find/suprType'), // 供应商类型
      axios.get('/find/unitMeas'), // 单位
      axios.get('/find/packType'), // 包装规格
    ]);

    // 提取字典数据
    prdtTypeData.value = prdtTypeResponse.data.PrdtType || [];
    brandTypeData.value = brandTypeResponse.data.BrandType || [];
    suprTypeData.value = suprTypeResponse.data.SuprType || [];
    unitMeasData.value = unitMeasResponse.data.UnitMeas || [];
    packTypeData.value = packTypeResponse.data.PackType || [];
  } catch (error) {
    console.error('获取字典数据失败:', error);
    ElMessage.error('获取字典数据失败，请稍后重试');
  }
};

// 获取装货数据
const fetchLoadingData = async () => {
  try {
    const response = await axios.get('/find/loadingInfo');
    loadingData.value = response.data.LoadingInfo || [];
    console.log(response.data.LoadingInfo)
  } catch (error) {
    console.error('获取装货信息失败:', error);
    ElMessage.error('获取装货信息失败，请稍后重试');
  }
};

onMounted(() => {
  fetchLoadingData();
  fetchDictionaryData();
});

const headerTitle = computed(() => '装货信息');
const addButtonText = computed(() => '添加装货信息');

const loadingRules = {
  Product: [{ required: true, message: '请选择产品', trigger: 'blur' }],
  Brand: [{ required: true, message: '请选择品牌', trigger: 'blur' }],
  PrdtPlant: [{ required: true, message: '请选择生产工厂', trigger: 'blur' }],
  BatNum: [{ required: true, message: '请输入批次号', trigger: 'blur' }],
  ItemNum: [{ required: true, message: '请输入件数', trigger: 'blur' }],
  PackSpec: [{ required: true, message: '请选择包装规格', trigger: 'blur' }],
  NetWeight: [{ required: true, message: '请输入净重量', trigger: 'blur' }],
  Unit: [{ required: true, message: '请选择单位', trigger: 'blur' }],
  CnrNum: [{ required: true, message: '请输入集装箱号', trigger: 'blur' }],
  SealNum: [{ required: true, message: '请输入铅封号', trigger: 'blur' }],
  VehNum: [{ required: true, message: '请输入车辆号', trigger: 'blur' }],
};
</script>

<style src="../assets/styles/Bottom.css"></style>

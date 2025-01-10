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

        <HeaderComponent :header-title="headerTitle" :add-button-text="addButtonText" v-model:search-query="searchQuery"
          @toggle-match-mode="toggleMatchMode" @toggle-id-mode="toggleIDMode" @add="handleAdd" />
        <el-header height="1px">
        </el-header>

        <el-main>
          <!-- 装货信息表格 -->
          <el-table :data="paginatedLoadingData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column label="产品" width="220%">
              <template #default="scope">
                <span v-if="scope.row.Cat.CatEngName">{{ scope.row.Cat.CatEngName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>

            <el-table-column label="品牌" width="220%">
              <template #default="scope">
                <span v-if="scope.row.Brand.BrandEngName">{{ scope.row.Brand.BrandEngName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column prop="PrdtPlant" label="生产工厂" width="220%"></el-table-column>
            <el-table-column prop="BatNum" label="批次号" width="220%"></el-table-column>
            <el-table-column prop="ItemNum" label="件数" width="220%"></el-table-column>

            <el-table-column label="包装规格" width="220%">
              <template #default="scope">
                <span v-if="scope.row.PackSpec.SpecName">{{ scope.row.PackSpec.SpecName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
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
            <el-form-item label="产品" prop="CatID">
              <el-select v-model="loadingForm.CatID" placeholder="请选择产品">
                <el-option v-for="item in prdtTypeData" :key="item.ID" :label="item.CatEngName"
                  :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌" prop="BrandID">
              <el-select v-model="loadingForm.BrandID" placeholder="请选择品牌">
                <el-option v-for="item in brandTypeData" :key="item.ID" :label="item.BrandEngName"
                  :value="item.ID"></el-option>
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

            <el-form-item label="包装规格" prop="PackSpecID">
              <el-select v-model="loadingForm.PackSpecID" placeholder="请选择包装规格">
                <el-option v-for="item in packTypeData" :key="item.ID" :label="item.SpecName"
                  :value="item.ID"></el-option>
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
            <el-form-item label="包装规格" prop="PackSpecID">
              <el-select v-model="loadingForm.PackSpecID" placeholder="请选择包装规格">
                <el-option v-for="item in packTypeData" :key="item.ID" :label="item.SpecName"
                  :value="item.ID"></el-option>
              </el-select>
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

import HeaderComponent from '@/components/HeaderComponent.vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import axios from 'axios';
import SideMenu from '@/components/SideMenu.vue';

import { useRoute } from 'vue-router';
const route = useRoute();
const isExactMatch = ref(true);
const onlyID = ref(true);
const toggleMatchMode = () => {
  isExactMatch.value = !isExactMatch.value;
};

const toggleIDMode = () => {
  onlyID.value = !onlyID.value;
};
const searchQuery = ref('');
const currentPage = ref(1);
const pageSize = 8;
const loadingData = ref([]);
const showLoadingDialog = ref(false);
const showViewLoadingDialog = ref(false);
const loadingForm = ref({
  CatID: '',
  PackSpecID: '',
  BrandID: '',
  PrdtPlant: '',
  BatNum: '',
  ItemNum: 0,
  NetWeight: 0,
  Unit: '',
  CnrNum: '',
  SealNum: '',

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

  let filteredData = loadingData.value; // 假设 loadingData 是你的装货数据

  if (searchQuery.value) {
    console.log(isExactMatch.value);
    console.log(onlyID.value);

    if (isExactMatch.value === false) {
      if (onlyID.value === false) {
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value) ||
          (item.Cat && item.Cat.CatEngName && item.Cat.CatEngName.includes(searchQuery.value)) ||
          (item.Brand && item.Brand.BrandEngName && item.Brand.BrandEngName.includes(searchQuery.value)) ||
          item.PrdtPlant.includes(searchQuery.value) ||
          item.BatNum.includes(searchQuery.value) ||
          item.ItemNum.toString().includes(searchQuery.value) ||
          (item.PackSpec && item.PackSpec.SpecName && item.PackSpec.SpecName.includes(searchQuery.value)) ||
          item.NetWeight.toString().includes(searchQuery.value) ||
          item.Unit.includes(searchQuery.value) ||
          item.CnrNum.includes(searchQuery.value) ||
          item.SealNum.includes(searchQuery.value) ||
          item.VehNum.includes(searchQuery.value) ||
          item.CreatedAt.includes(searchQuery.value) ||
          item.UpdatedAt.includes(searchQuery.value)
        );
      } else {
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value)
        );
      }
    } else {
      if (onlyID.value === false) {
        filteredData = filteredData.filter(item =>
          item.ID.toString() === searchQuery.value ||
          (item.Cat && item.Cat.CatEngName && item.Cat.CatEngName === searchQuery.value) ||
          (item.Brand && item.Brand.BrandEngName && item.Brand.BrandEngName === searchQuery.value) ||
          item.PrdtPlant === searchQuery.value ||
          item.BatNum === searchQuery.value ||
          item.ItemNum.toString() === searchQuery.value ||
          (item.PackSpec && item.PackSpec.SpecName && item.PackSpec.SpecName === searchQuery.value) ||
          item.NetWeight.toString() === searchQuery.value ||
          item.Unit === searchQuery.value ||
          item.CnrNum === searchQuery.value ||
          item.SealNum === searchQuery.value ||
          item.VehNum === searchQuery.value ||
          item.CreatedAt === searchQuery.value ||
          item.UpdatedAt === searchQuery.value
        );
      } else {
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
          resetLoadingForm();
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
    LoadingInfoID: '',
    CreatedAt: '', // 添加 CreatedAt
    UpdatedAt: '', // 添加 UpdatedAt
  };
};

const submitLoadingForm = async () => {
  try {

    const isValid = await loadingFormRef.value.validate();
    if (!isValid) {
      ElMessage.error('请填写所有必填字段');
      console.log('验证不通过');
      return; // 如果验证不通过，阻止提交
    }

    // loadingForm.value.ItemNum = parseInt(loadingForm.value.ItemNum, 10);
    // loadingForm.value.NetWeight = parseInt(loadingForm.value.NetWeight, 10);

    const formData = new FormData(); // 创建 FormData 对象

    Object.keys(loadingForm.value).forEach((key) => {
      formData.append(key, loadingForm.value[key]);
    });

    console.log(formData)

    const response = await axios.post('/save/loadingInfo', formData, {
      headers: {

        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    });
    if (response.status === 200) {
      ElMessage.success('装货信息保存成功');
      showLoadingDialog.value = false
      fetchLoadingData();
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
      axios.get('/find/cat'), // 产品类型
      axios.get('/find/brand'), // 品牌类型
      axios.get('/find/suprType'), // 供应商类型
      axios.get('/find/unitMeas'), // 单位
      axios.get('/find/packSpec'), // 包装规格
    ]);

    // 提取字典数据
    prdtTypeData.value = prdtTypeResponse.data.Cat || [];
    brandTypeData.value = brandTypeResponse.data.Brand || [];
    suprTypeData.value = suprTypeResponse.data.SuprType || [];
    unitMeasData.value = unitMeasResponse.data.UnitMeas || [];
    packTypeData.value = packTypeResponse.data.PackSpec || [];
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
  searchQuery.value = route.query.searchQuery || '';
});

const headerTitle = computed(() => '装货信息');
const addButtonText = computed(() => '添加装货信息');

// 自定义验证函数：检查非空
const validateNotEmpty = (rule, value, callback) => {
  if (value === '' || value === null || value === undefined) {
    callback(new Error(rule.message));  // 使用 rule.message 作为错误消息
  } else {
    callback();  // 验证通过
  }
};

// 自定义验证函数：检查大于 0
const validateGreaterThanZero = (rule, value, callback) => {
  if (value < 0) {
    callback(new Error(rule.message || '该字段必须大于 0'));  // 提供字段的错误消息
  } else {
    callback();  // 验证通过
  }
};

const loadingRules = {
  CatID: [{ required: true, message: '请选择产品', trigger: 'blur' }],
  BrandID: [{ required: true, message: '请选择品牌', trigger: 'blur' }],
  ItemNum: [
    { validator: validateGreaterThanZero, message: '件数必须大于 0', trigger: 'blur' }  // 新增验证：件数大于 0
  ],
  PackSpecID: [{ required: true, message: '请选择包装规格', trigger: 'blur' }],
  // NetWeight: [
  //   { required: true, validator: validateNotEmpty, message: '请输入净重量', trigger: 'blur' },
  //   { validator: validateGreaterThanZero, message: '净重量必须大于 0', trigger: 'blur' }  // 新增验证：净重量大于 0
  // ],
  // Unit: [{ required: true, message: '请选择单位', trigger: 'blur' }],
  // CnrNum: [{ required: true, message: '请输入集装箱号', trigger: 'blur' }],
  // SealNum: [{ required: true, message: '请输入铅封号', trigger: 'blur' }],
  // VehNum: [{ required: true, message: '请输入车辆号', trigger: 'blur' }]
};
</script>

<style src="../assets/styles/Bottom.css"></style>

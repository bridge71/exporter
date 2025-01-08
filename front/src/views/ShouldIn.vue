<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-19'" />
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>

        <HeaderComponent :header-title="headerTitle" :add-button-text="addButtonText" v-model:search-query="searchQuery"
          @toggle-match-mode="toggleMatchMode" @toggle-id-mode="toggleIdMode" @add="handleAdd" />
        <el-header height="1px">
        </el-header>

        <el-main>

          <!-- 销售订单信息表格 -->
          <el-table :data="paginatedSaleData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="BillReceNum" label="应收账款单号" width="220%"></el-table-column>
            <el-table-column prop="DocDate" label="单据日期" width="220%"></el-table-column>
            <el-table-column prop="ExpReceDate" label="预计收款日期" width="220%"></el-table-column>
            <el-table-column prop="FinaDocType" label="单据类型" width="220%"></el-table-column>
            <el-table-column prop="FinaDocStatus" label="单据状态" width="220%"></el-table-column>
            <el-table-column prop="Merc" label="付款方" width="220%"></el-table-column>
            <el-table-column prop="AcctName" label="收款方" width="220%"></el-table-column>
            <el-table-column prop="BankAccName" label="付款银行账户" width="220%"></el-table-column>
            <el-table-column prop="AccName" label="收款银行账户" width="220%"></el-table-column>
            <el-table-column prop="TotAmt" label="总金额" width="220%"></el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="描述" width="220%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>

            <el-table-column label="操作" fixed="right" width="420%">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.ID)" type="text" size="small">删除</el-button>
                  <el-button @click="fetchSaleData(scope.row.ID)" type="text" size="small">销售订单</el-button>
                  <el-button @click="fetchSendData(scope.row.ID)" type="text" size="small">销售发货单</el-button>
                  <el-button @click="fetchInData(scope.row.ID)" type="text" size="small">收款单</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="saleData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>


    <el-dialog v-model="SaleVisible" title="销售订单" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="SaleId" placeholder="请输入ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addSale(nowId)">添加</el-button>
      </div>

      <!-- 产品明细表格 -->
      <el-table :data="SaleData" height="400" style="width: 100%">
        <el-table-column prop="ID" label="ID" width="60%" />
        <el-table-column prop="OrderNum" label="订单编号" width="150%" />
        <el-table-column prop="Merc" label="购买方" width="150%" />
        <el-table-column prop="AcctName" label="销售方" width="150%" />
        <el-table-column prop="SpecName" label="包装规格" width="150%" />
        <el-table-column prop="TotAmt" label="总金额" width="150%" />
        <el-table-column prop="Currency" label="币种" width="150%" />
        <el-table-column label="操作" fixed="right" width="150">
          <template #default="scope">
            <!-- <el-button type="text" size="small" @click="viewProduct(scope.row)">查看</el-button> -->

            <el-button type="text" size="small" @click="CheckSale(scope.row.ID)">跳转</el-button>
            <el-button type="text" size="small" @click="DeleteSale(scope.$index, nowId, scope.row.ID)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="sendVisible" title="销售发货单" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="sendId" placeholder="请输入销售发货单ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addSend(nowId)">添加</el-button>
      </div>

      <!-- 销售发货表格 -->
      <el-table :data="sendData" height="400" style="width: 100%">
        <el-table-column prop="ID" label="ID" width="60%" />
        <el-table-column prop="SaleInvNum" label="销售发票号" width="150%" />
        <el-table-column prop="AcctName" label="发货人" width="150%" />
        <el-table-column prop="Note1" label="提单货物描述" width="360%" />
        <el-table-column prop="SaleInvDate" label="销售发票日期" width="420%" />
        <el-table-column label="操作" fixed="right" width="150">
          <template #default="scope">
            <!-- <el-button type="text" size="small" @click="viewProduct(scope.row)">查看</el-button> -->
            <el-button type="text" size="small" @click="DeleteSend(scope.$index, nowId, scope.row.ID)">删除</el-button>

            <el-button type="text" size="small" @click="CheckSend(scope.row.ID)">跳转</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="InVisible" title="收款单" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="InId" placeholder="请输入收款单ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addIn(nowId)">添加</el-button>
      </div>
      <el-table :data="InData" style="width: 100%" max-height="450">
        <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
        <el-table-column prop="ReceNum" label="账款单号" width="220%"></el-table-column>
        <el-table-column prop="RealReceDate" label="实际收款日期" width="220%"></el-table-column>
        <el-table-column prop="ExpReceDate" label="预计收款日期" width="220%"></el-table-column>
        <el-table-column prop="FinaDocType" label="单据类型" width="220%"></el-table-column>
        <el-table-column prop="FinaDocStatus" label="单据状态" width="220%"></el-table-column>
        <el-table-column prop="Merc" label="付款方" width="220%"></el-table-column>
        <el-table-column prop="AcctName" label="收款方" width="220%"></el-table-column>
        <el-table-column prop="BankAccName" label="付款银行账户" width="220%"></el-table-column>
        <el-table-column prop="AccName" label="收款银行账户" width="220%"></el-table-column>
        <el-table-column prop="TotAmt" label="收款金额" width="220%"></el-table-column>
        <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
        <el-table-column prop="Notes" label="描述" width="220%"></el-table-column>
        <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>

        <el-table-column label="操作" fixed="right" width="160%">
          <template #default="scope">
            <el-row type="flex" justify="space-between">
              <el-button type="text" size="small" @click="DeleteIn(scope.$index, nowId, scope.row.ID)">删除</el-button>
              <el-button type="text" size="small" @click="CheckIn(scope.row.ID)">跳转</el-button>
            </el-row>
          </template>
        </el-table-column>
      </el-table>
      <!-- 销售发货表格 -->
    </el-dialog>

    <el-dialog v-model="showShouldInDialog" title="应收账款单" width="80%" @close="resetShouldInForm">
      <el-form :model="shouldInForm" label-width="150px" :rules="saleRules" ref="shouldInFormRef">
        <!-- 第一行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="应收账款单号" prop="BillReceNum">
              <el-input v-model="shouldInForm.BillReceNum" placeholder="请输入应收账款单号"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据日期" prop="DocDate">
              <el-date-picker v-model="shouldInForm.DocDate" type="date" placeholder="请选择单据日期"></el-date-picker>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="预计收款日期" prop="ExpReceDate">
              <el-date-picker v-model="shouldInForm.ExpReceDate" type="date" placeholder="请选择预计收款日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 第三行 -->

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据类型" prop="FinaDocType">
              <el-select v-model="shouldInForm.FinaDocType" @change="onFinaDocTypeChange" placeholder="请选择财务单据类型">
                <el-option v-for="type in FinaDocTypeData" :key="type.FinaDocTypeId" :label="type.FinaDocType"
                  :value="type.FinaDocType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单据状态" prop="FinaDocStatus">
              <el-select v-model="shouldInForm.FinaDocStatus" @change="onFinaDocStatusChange" placeholder="请选择财务单据状态">
                <el-option v-for="st in FinaDocStatusData" :key="st.FinaDocStatusId" :label="st.FinaDocStatus"
                  :value="st.FinaDocStatus"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="付款方" prop="MerchantId">
              <el-select v-model="shouldInForm.MerchantId" @change="onMerchantChange" placeholder="请选择付款方">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="收款方" prop="AcctId">
              <el-select v-model="shouldInForm.AcctId" @change="onAcctChange" placeholder="请选择收款方">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="付款银行账户" prop="BankAccountId">
              <el-select v-model="shouldInForm.BankAccountId" @change="onBankAccountChange" placeholder="请选择付款银行账户">
                <el-option v-for="bankAccount in bankAccountData" :key="bankAccount.ID" :label="bankAccount.BankAccName"
                  :value="bankAccount.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="收款银行账户" prop="AcctBankId">
              <el-select v-model="shouldInForm.AcctBankId" @change="onAcctBankChange" placeholder="请选择收款银行账户">
                <el-option v-for="acctBank in acctBankData" :key="acctBank.ID" :label="acctBank.AccName"
                  :value="acctBank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <!-- 第四行 -->

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="总金额" prop="TotAmt">
              <el-input v-model="shouldInForm.TotAmt"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="CurrencyId">
              <el-select v-model="shouldInForm.Currency" placeholder="请选择币种">
                <el-option v-for="currency in currencyData" :key="currency.Currency" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="shouldInForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">


              <el-upload v-if="!shouldInForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(shouldInForm.FileId, shouldInForm.FileName)">
                下载文件
              </el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="shouldInForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showShouldInDialog = false">取消</el-button>
          <el-button type="primary" @click="submitShouldInForm()">提交</el-button>
        </span>
      </template>
    </el-dialog>


    <el-dialog v-model="showshowShouldInDialog" title="应收账款单" width="80%" @close="resetShouldInForm">
      <el-form :model="shouldInForm" label-width="150px" :rules="saleRules" ref="shouldInFormRef">
        <!-- 第一行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="应收账款单号" prop="BillReceNum">
              <el-input v-model="shouldInForm.BillReceNum" placeholder="请输入应收账款单号" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据日期" prop="DocDate">
              <el-date-picker v-model="shouldInForm.DocDate" type="date" placeholder="请选择单据日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="预计收款日期" prop="ExpReceDate">
              <el-date-picker v-model="shouldInForm.ExpReceDate" type="date" placeholder="请选择预计收款日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 第三行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据类型" prop="FinaDocType">
              <el-select v-model="shouldInForm.FinaDocType" @change="onFinaDocTypeChange" placeholder="请选择财务单据类型"
                :disabled="true">
                <el-option v-for="type in FinaDocTypeData" :key="type.FinaDocTypeId" :label="type.FinaDocType"
                  :value="type.FinaDocType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单据状态" prop="FinaDocStatus">
              <el-select v-model="shouldInForm.FinaDocStatus" @change="onFinaDocStatusChange" placeholder="请选择财务单据状态"
                :disabled="true">
                <el-option v-for="st in FinaDocStatusData" :key="st.FinaDocStatusId" :label="st.FinaDocStatus"
                  :value="st.FinaDocStatus"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="付款方" prop="MerchantId">
              <el-select v-model="shouldInForm.MerchantId" @change="onMerchantChange" placeholder="请选择付款方"
                :disabled="true">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="收款方" prop="AcctId">
              <el-select v-model="shouldInForm.AcctId" @change="onAcctChange" placeholder="请选择收款方" :disabled="true">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="付款银行账户" prop="BankAccountId">
              <el-select v-model="shouldInForm.BankAccountId" @change="onBankAccountChange" placeholder="请选择付款银行账户"
                :disabled="true">
                <el-option v-for="bankAccount in bankAccountData" :key="bankAccount.ID" :label="bankAccount.BankAccName"
                  :value="bankAccount.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="收款银行账户" prop="AcctBankId">
              <el-select v-model="shouldInForm.AcctBankId" @change="onAcctBankChange" placeholder="请选择收款银行账户"
                :disabled="true">
                <el-option v-for="acctBank in acctBankData" :key="acctBank.ID" :label="acctBank.AccName"
                  :value="acctBank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 第四行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总金额" prop="TotAmt">
              <el-input v-model="shouldInForm.TotAmt" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="CurrencyId">
              <el-select v-model="shouldInForm.Currency" placeholder="请选择币种" :disabled="true">
                <el-option v-for="currency in currencyData" :key="currency.Currency" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="shouldInForm.Notes" type="textarea" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!shouldInForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true" :disabled="true">
                <el-button type="primary" :disabled="true">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(shouldInForm.FileId, shouldInForm.FileName)">
                下载文件
              </el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="shouldInForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showShouldInDialog = false">取消</el-button>
          <!-- 隐藏提交按钮 -->
          <!-- <el-button type="primary" @click="submitShouldInForm()">提交</el-button> -->
        </span>
      </template>
    </el-dialog>


  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import HeaderComponent from '@/components/HeaderComponent.vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import axios from 'axios'; // 引入 axios
import SideMenu from '@/components/SideMenu.vue'; // 引入 SideMenu 组件
import { useRouter } from 'vue-router';
const router = useRouter();
import { useRoute } from 'vue-router';
const route = useRoute();
// 匹配模式（默认是模糊匹配）
const FinaDocTypeData = ref([])
const FinaDocStatusData = ref([])
const isExactMatch = ref(true);
const onlyId = ref(true);
const toggleMatchMode = () => {
  isExactMatch.value = !isExactMatch.value;
};

const toggleIdMode = () => {
  onlyId.value = !onlyId.value;
};

const InVisible = ref(false)
const InId = ref(null)
const InData = ref([])

const DeleteIn = (index, ID, InId) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该产品信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    console.log('Confirmed delete', ID); // 添加调试信息

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("InId", InId)

    axios.post('/delete/shouldIn/in', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })

      // axios.post('/delete/sale/prdtInfo', {
      //   "ID": ID,
      //   "PrdtInfoId": PrdtInfoId.value
      // })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchInData(nowId.value); // 重新获取会计实体信息数据
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
const addIn = async (ID) => {

  console.log(nowId.value)
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("InId", InId.value)

    const response = await axios.post('/add/shouldIn/in', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchInData(nowId.value)
    InVisible.value = ''
  } catch (error) {
    console.error('添加失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};
const CheckIn = (ID) => {
  try {
    // 确保 ID 是字符串
    const searchQuery = String(ID);

    // 使用路由的 resolve 方法生成完整路径
    const route = router.resolve({
      name: 'In', // 路由名称
      query: { searchQuery }, // 传递的查询参数（对象形式）
    });

    // 在新标签页打开
    window.open(route.href, '_blank');
  } catch (error) {
    ElMessage.error("查看失败");
  }
};
const SaleVisible = ref(false)
const SaleId = ref(null)
const SaleData = ref([])

const DeleteSale = (index, ID, SaleId) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该产品信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    console.log('Confirmed delete', ID); // 添加调试信息

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("SaleId", SaleId)

    axios.post('/delete/shouldIn/sale', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })

      // axios.post('/delete/sale/prdtInfo', {
      //   "ID": ID,
      //   "PrdtInfoId": PrdtInfoId.value
      // })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchSaleData(nowId.value); // 重新获取会计实体信息数据
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
const addSale = async (ID) => {

  console.log(nowId.value)
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("SaleId", SaleId.value)

    const response = await axios.post('/add/shouldIn/sale', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchSaleData(nowId.value)
    SaleVisible.value = ''
  } catch (error) {
    console.error('添加失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};
const CheckSale = (ID) => {
  try {
    // 确保 ID 是字符串
    const searchQuery = String(ID);

    // 使用路由的 resolve 方法生成完整路径
    const route = router.resolve({
      name: 'Sale', // 路由名称
      query: { searchQuery }, // 传递的查询参数（对象形式）
    });

    // 在新标签页打开
    window.open(route.href, '_blank');
  } catch (error) {
    ElMessage.error("查看失败");
  }
};
const shouldInForm = ref({
  BillReceNum: '', // 应收账款单号
  DocDate: '', // 单据日期
  ExpReceDate: '', // 预计收款日期
  FinaDocType: '', // 财务单据类型
  FinaDocStatus: '', // 财务单据状态
  MerchantId: '', // 付款方 ID
  Merc: '', // 付款方名称
  AcctId: '', // 收款方 ID
  AcctName: '', // 收款方名称
  BankAccountId: '', // 付款银行账户 ID
  BankAccName: '', // 付款银行账户名称
  AcctBankId: '', // 收款银行账户 ID
  AccName: '', // 收款银行账户名称
  TotAmt: '', // 总金额
  Currency: '', // 币种
  Notes: '', // 描述
  FileId: '', // 文件 ID
  FileName: '', // 文件名
});

// 表单验证规则
const shouldInRules = {
  BillReceNum: [{ required: true, message: '请输入应收账款单号', trigger: 'blur' }],
  TotAmt: [{ type: 'number', message: '必须为正数', trigger: 'blur' }],
};
// 控制主弹窗显示
const sendVisible = ref(false);

const sendData = ref([]);
const sendId = ref(null);

const CheckSend = (ID) => {
  try {
    // 确保 ID 是字符串
    const searchQuery = String(ID);

    // 使用路由的 resolve 方法生成完整路径
    const route = router.resolve({
      name: 'Send', // 路由名称
      query: { searchQuery }, // 传递的查询参数（对象形式）
    });

    // 在新标签页打开
    window.open(route.href, '_blank');
  } catch (error) {
    ElMessage.error("查看失败");
  }
};
// 删除按钮逻辑
const DeleteSend = (index, ID, SendId) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("SendId", SendId)

    axios.post('/delete/shouldIn/send', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchSendData(nowId.value); // 重新获取会计实体信息数据
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
const addSend = async (ID) => {

  console.log(nowId.value)
  // acctForm.value.ID = parseInt(acctForm.value.ID, 10);

  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("SendId", sendId.value)

    const response = await axios.post('/add/shouldIn/send', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchSendData(nowId.value)
    sendId.value = ''
  } catch (error) {
    console.error('添加产品明细失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

const fetchSendData = async (ID) => {
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/shouldIn/send', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    sendData.value = response.data.Send; // 假设返回的数据结构中有 PrdtInfo 字段
    nowId.value = ID
    sendVisible.value = true;
    console.log("nowid", nowId.value)
  } catch (error) {
    console.error('获取产品明细失败:', error);
    ElMessage.error('获取产品明细失败');
  }
};
const nowId = ref(null);

const file = ref(null);
const searchQuery = ref(''); // 添加搜索查询字段
const currentPage = ref(1); // 当前页码
const pageSize = 8; // 每页显示的行数



const shouldInData = ref([])

const fetchShouldInData = async () => {
  try {
    const response = await axios.get('/find/shouldIn'); // 调用货币数据接口
    shouldInData.value = response.data.ShouldIn; // 假设返回的数据结构中有 Currency 字段
    console.log(shouldInData.value)
  } catch (error) {
    console.error('获取货币数据失败:', error);
    ElMessage.error('获取货币数据失败，请稍后重试');
  }
};
const fetchFinaDocTypeData = async () => {
  try {
    const response = await axios.get('/find/finaDocType'); // 调用货币数据接口
    FinaDocTypeData.value = response.data.FinaDocType; // 假设返回的数据结构中有 Currency 字段
  } catch (error) {
    console.error('获取单据数据错误:', error);
    ElMessage.error('获取数据失败，请稍后重试');
  }
}// 获取单据类型数据
const fetchFinaDocStatusData = async () => {
  try {
    const response = await axios.get('/find/finaDocStatus'); // 调用货币数据接口
    FinaDocStatusData.value = response.data.FinaDocStatus; // 假设返回的数据结构中有 Currency 字段
  } catch (error) {
    console.error('获取单据类型数据错误:', error);
    ElMessage.error('获取数据失败，请稍后重试');
  }
};

const paginatedSaleData = computed(() => {
  let filteredData = shouldInData.value;

  // 如果有搜索查询
  if (searchQuery.value) {
    console.log("s2s")
    if (isExactMatch.value === false) {
      if (onlyId.value === false) {
        // 模糊匹配多个字段
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value) ||
          item.BillReceNum.includes(searchQuery.value) ||
          item.DocDate.includes(searchQuery.value) ||
          item.ExpReceDate.includes(searchQuery.value) ||
          item.FinaDocType.includes(searchQuery.value) ||
          item.FinaDocStatus.includes(searchQuery.value) ||
          item.Merc.includes(searchQuery.value) ||
          item.AcctName.includes(searchQuery.value) ||
          item.BankAccName.includes(searchQuery.value) ||
          item.AccName.includes(searchQuery.value) ||
          item.TotAmt.toString().includes(searchQuery.value) ||
          item.Currency.includes(searchQuery.value) ||
          item.Notes.includes(searchQuery.value) ||
          item.FileName.includes(searchQuery.value)
        );
      } else {
        // 仅匹配 ID
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value)
        );
      }
    } else {
      if (onlyId.value === false) {
        // 精确匹配多个字段
        filteredData = filteredData.filter(item =>
          item.ID.toString() === searchQuery.value ||
          item.BillReceNum === searchQuery.value ||
          item.DocDate === searchQuery.value ||
          item.ExpReceDate === searchQuery.value ||
          item.FinaDocType === searchQuery.value ||
          item.FinaDocStatus === searchQuery.value ||
          item.Merc === searchQuery.value ||
          item.AcctName === searchQuery.value ||
          item.BankAccName === searchQuery.value ||
          item.AccName === searchQuery.value ||
          item.TotAmt.toString() === searchQuery.value ||
          item.Currency === searchQuery.value ||
          item.Notes === searchQuery.value ||
          item.FileName === searchQuery.value
        );
      } else {
        // 仅精确匹配 ID
        filteredData = filteredData.filter(item =>
          item.ID.toString() === searchQuery.value
        );
      }
    }
  }

  // 分页逻辑
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});
onMounted(() => {
  fetchShouldInData();
  fetchAcctData(); // 获取会计实体信息
  fetchMerchantData(); // 获取购买方信息
  fetchFinaDocStatusData(); // 获取单据状态数据
  fetchFinaDocTypeData(); // 获取单据类型数据
  fetchAcctBankData(); // 获取我方银行账户数据
  fetchBankAccountData(); // 获取对方银行账户数据
  fetchCurrencyData(); // 新增：获取货币数据

  searchQuery.value = route.query.searchQuery || '';
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
// 定义接口请求函数
const fetchAcctData = async () => {
  try {
    const response = await axios.get('/find/acct'); // 调用会计实体信息接口
    acctData.value = response.data.Acct; // 假设返回的数据结构中有 Acct 字段
  } catch (error) {
    console.error('获取会计实体信息失败:', error);
    ElMessage.error('获取会计实体信息失败');
  }
};

const fetchMerchantData = async () => {
  try {
    const response = await axios.get('/find/merchant'); // 调用购买方信息接口
    merchantData.value = response.data.Merchant; // 假设返回的数据结构中有 Merchant 字段
  } catch (error) {
    console.error('获取购买方信息失败:', error);
    ElMessage.error('获取购买方信息失败');
  }
};

const fetchAcctBankData = async () => {
  try {
    const response = await axios.get('/find/acctBank'); // 调用我方银行账户接口
    acctBankData.value = response.data.AcctBank; // 假设返回的数据结构中有 AcctBank 字段
  } catch (error) {
    console.error('获取我方银行账户失败:', error);
    ElMessage.error('获取我方银行账户失败');
  }
};

const fetchBankAccountData = async () => {
  try {
    const response = await axios.get('/find/bankAccount'); // 调用对方银行账户接口
    bankAccountData.value = response.data.BankAccount; // 假设返回的数据结构中有 BankAccount 字段
  } catch (error) {
    console.error('获取对方银行账户失败:', error);
    ElMessage.error('获取对方银行账户失败');
  }
};

const fetchSaleData = async (ID) => {
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/shouldIn/sale', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    SaleData.value = response.data.Sale; // 假设返回的数据结构中有 PrdtInfo 字段
    nowId.value = ID
    SaleVisible.value = true;
  } catch (error) {
    console.error('获取失败:', error);
    ElMessage.error('获取失败');
  }
};

const fetchInData = async (ID) => {
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/shouldIn/in', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    InData.value = response.data.In; // 假设返回的数据结构中有 PrdtInfo 字段
    nowId.value = ID
    InVisible.value = true;
  } catch (error) {
    console.error('获取失败:', error);
    ElMessage.error('获取失败');
  }
};

// 销售订单信息对话框显示状态
const showShouldInDialog = ref(false);
const showshowShouldInDialog = ref(false);

// 表格数据（初始值为空数组）
const acctData = ref([]); // 会计实体信息
const merchantData = ref([]); // 购买方信息
const acctBankData = ref([]); // 我方银行账户数据
const bankAccountData = ref([]); // 对方银行账户数据
const saleData = ref([]); // 销售订单信息

// 根据当前选中的菜单项动态更改标题和按钮文本
const headerTitle = computed(() => {
  return '应收账款单信息';
});

const addButtonText = computed(() => {
  return '添加应收账款单信息';
});

// 添加按钮点击事件
const handleAdd = () => {
  showShouldInDialog.value = true;
};

const handleEdit = (index, row) => {
  // 将当前行的数据赋值给 saleForm
  shouldInForm.value = { ...row };

  // 检查是否已上传文件
  if (row.FileId) {
    shouldInForm.value.FileId = row.FileId; // 保存 FileId
    shouldInForm.value.FileName = row.FileName; // 保存文件名
  }

  showShouldInDialog.value = true; // 打开销售订单信息对话框
};

// 查看按钮逻辑
const handleView = (index, row) => {
  shouldInForm.value = { ...row }; // 将当前行的数据赋值给 saleForm
  showshowShouldInDialog.value = true; // 打开查看销售订单信息对话框
};

// 删除按钮逻辑
const handleDelete = (index, ID) => {
  ElMessageBox.confirm('确定要删除该应收账款单信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    axios.post('/delete/shouldIn', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchShouldInData(); // 重新获取销售订单信息数据
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
// 重置销售订单信息表单
const resetShouldInForm = () => {

  shouldInForm.value = {
    BillReceNum: '', // 应收账款单号
    DocDate: '', // 单据日期
    ExpReceDate: '', // 预计收款日期
    FinaDocType: '', // 财务单据类型
    FinaDocStatus: '', // 财务单据状态
    MerchantId: '', // 付款方 ID
    Merc: '', // 付款方名称
    AcctId: '', // 收款方 ID
    AcctName: '', // 收款方名称
    BankAccountId: '', // 付款银行账户 ID
    BankAccName: '', // 付款银行账户名称
    AcctBankId: '', // 收款银行账户 ID
    AccName: '', // 收款银行账户名称
    TotAmt: '', // 总金额
    Currency: '', // 币种
    Notes: '', // 描述
    FileId: '', // 文件 ID
    FileName: '', // 文件名
  };

  file.value = null; // 重置文件对象
  if (uploadRef.value) {
    uploadRef.value.clearFiles(); // 清空文件列表
  }
};

const uploadRef = ref(null);

const submitShouldInForm = async () => {
  try {
    const formData = new FormData();

    // 添加其他字段
    Object.keys(shouldInForm.value).forEach((key) => {
      formData.append(key, shouldInForm.value[key]);
    });


    // 添加文件
    if (file.value) {
      formData.append('file', file.value);
    }

    // 提交表单

    const response = await axios.post('/save/shouldIn', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    if (response.status === 200) {
      ElMessage.success('销售订单信息保存成功');
      showShouldInDialog.value = false; // 关闭对话框
      fetchShouldInData(); // 重新获取销售订单信息数据
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存销售订单信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};
// 监听 change 事件并更新 AcctName
const onAcctChange = (value) => {
  const selectedAcct = acctData.value.find(acct => acct.ID === value);
  if (selectedAcct) {
    saleForm.value.AcctName = selectedAcct.AcctName;
  }
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
    link.setAttribute('download', fileName || `file_${fileId}`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error('文件下载失败:', error);
    ElMessage.error('文件下载失败');
  }
};
const handleFileChange = (uploadFile) => {
  file.value = uploadFile.raw; // 保存选择的文件
  saleForm.value.FileName = uploadFile.name; // 更新文件名
};

// 监听 change 事件并更新 Merc
const onMerchantChange = (value) => {
  const selectedMerchant = merchantData.value.find(merchant => merchant.ID === value);
  if (selectedMerchant) {
    saleForm.value.Merc = selectedMerchant.Merc;
  }
};

// 监听 change 事件并更新 AccName
const onAcctBankChange = (value) => {
  const selectedAcctBank = acctBankData.value.find(acctBank => acctBank.ID === value);
  if (selectedAcctBank) {
    saleForm.value.AccName = selectedAcctBank.AccName;
  }
};

// 监听 change 事件并更新 BankAccName
const onBankAccountChange = (value) => {
  const selectedBankAccount = bankAccountData.value.find(bankAccount => bankAccount.ID === value);
  if (selectedBankAccount) {
    saleForm.value.BankAccName = selectedBankAccount.BankAccName;
  }
};

// 分页逻辑
const handlePageChange = (page) => {
  currentPage.value = page;
};
</script>

<style src="../assets/styles/Bottom.css"></style>

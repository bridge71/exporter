<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-20'" />
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>

        <HeaderComponent :header-title="headerTitle" :add-button-text="addButtonText" v-model:search-query="searchQuery"
          @toggle-match-mode="toggleMatchMode" @toggle-id-mode="toggleIDMode" @add="handleAdd" />
        <el-header height="1px">
        </el-header>

        <el-main>

          <!-- 应付账款单信息表格 -->
          <el-table :data="paginatedBuyData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="BillReceNum" label="应付账款单号" width="220%"></el-table-column>

            <el-table-column label="付款方" width="220%">
              <template #default="scope">
                <span v-if="scope.row.Acct.AcctName">{{ scope.row.Acct.AcctName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column label="收款方" width="220%">
              <template #default="scope">
                <span v-if="scope.row.Merchant.Merc">{{ scope.row.Merchant.Merc }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column label="付款银行账户" width="220%">
              <template #default="scope">
                <span v-if="scope.row.AcctBank.AccName">{{ scope.row.AcctBank.AccName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column label="收款银行账户" width="220%">
              <template #default="scope">
                <span v-if="scope.row.BankAccount.BankAccName">{{ scope.row.BankAccount.BankAccName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>

            <el-table-column prop="DocDate" label="单据日期" width="220%"></el-table-column>
            <el-table-column prop="ExpReceDate" label="预计付款日期" width="220%"></el-table-column>
            <el-table-column prop="FinaDocType" label="单据类型" width="220%"></el-table-column>
            <el-table-column prop="FinaDocStatus" label="单据状态" width="220%"></el-table-column>
            <el-table-column prop="TotAmt" label="总金额" width="220%"></el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="描述" width="220%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>

            <el-table-column label="操作" fixed="right" width="420%">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="fetchCostInfoData(scope.row.ID)" type="text" size="small">产品明细</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.ID)" type="text" size="small">删除</el-button>
                  <el-button @click="fetchBuyData(scope.row.ID)" type="text" size="small">采购订单</el-button>
                  <el-button @click="fetchPurRecData(scope.row.ID)" type="text" size="small">采购收货单</el-button>
                  <el-button @click="fetchOutData(scope.row.ID)" type="text" size="small">收款单</el-button>
                  <el-button @click="fetchCostInfoData(scope.row.ID)" type="text" size="small">费用明细</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="buyData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <el-dialog v-model="costInfoVisible" title="费用明细" width="80%">
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="costInfoID" placeholder="请输入ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addCostInfo(nowID)">添加</el-button>
      </div>

      <el-table :data="costInfoData" height="400" style="width: 100%">

        <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
        <el-table-column prop="ExpType" label="费用类型" width="220%"></el-table-column>
        <el-table-column prop="Rates" label="收费标准" width="220%"></el-table-column>
        <el-table-column prop="UnitPrice" label="单价" width="220%"></el-table-column>
        <el-table-column prop="Number" label="数量" width="220%"></el-table-column>
        <el-table-column prop="Amount" label="金额" width="220%"></el-table-column>
        <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
        <el-table-column label="操作" fixed="right" width="150">
          <template #default="scope">
            <!-- <el-button type="text" size="small" @click="viewProduct(scope.row)">查看</el-button> -->

            <el-button type="text" size="small" @click="CheckCostInfo(scope.row.ID)">跳转</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="BuyVisible" title="采购订单" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="BuyID" placeholder="请输入ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addBuy(nowID)">添加</el-button>
      </div>

      <!-- 产品明细表格 -->
      <el-table :data="BuyData" height="400" style="width: 100%">
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

            <el-button type="text" size="small" @click="CheckBuy(scope.row.ID)">跳转</el-button>
            <el-button type="text" size="small" @click="DeleteBuy(scope.$index, nowID, scope.row.ID)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="purRecVisible" title="采购收货单" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="purRecID" placeholder="请输入采购收货单ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addPurRec(nowID)">添加</el-button>
      </div>

      <!-- 采购收货表格 -->
      <el-table :data="purRecData" height="400" style="width: 100%">
        <el-table-column prop="ID" label="ID" width="60%" />
        <el-table-column prop="SaleInvNum" label="销售发票号" width="150%" />
        <el-table-column prop="AcctName" label="发货人" width="150%" />
        <el-table-column prop="Note1" label="提单货物描述" width="360%" />
        <el-table-column prop="SaleInvDate" label="销售发票日期" width="420%" />
        <el-table-column label="操作" fixed="right" width="150">
          <template #default="scope">
            <!-- <el-button type="text" size="small" @click="viewProduct(scope.row)">查看</el-button> -->
            <el-button type="text" size="small" @click="DeletePurRec(scope.$index, nowID, scope.row.ID)">删除</el-button>

            <el-button type="text" size="small" @click="CheckPurRec(scope.row.ID)">跳转</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="OutVisible" title="收款单" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="OutID" placeholder="请输入收款单ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addOut(nowID)">添加</el-button>
      </div>
      <el-table :data="OutData" style="width: 100%" max-height="450">
        <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
        <el-table-column prop="ReceNum" label="账款单号" width="220%"></el-table-column>
        <el-table-column prop="RealReceDate" label="实际付款日期" width="220%"></el-table-column>
        <el-table-column prop="ExpReceDate" label="预计付款日期" width="220%"></el-table-column>
        <el-table-column prop="FinaDocType" label="单据类型" width="220%"></el-table-column>
        <el-table-column prop="FinaDocStatus" label="单据状态" width="220%"></el-table-column>
        <el-table-column prop="Merc" label="收款方" width="220%"></el-table-column>
        <el-table-column prop="AcctName" label="付款方" width="220%"></el-table-column>
        <el-table-column prop="BankAccName" label="收款银行账户" width="220%"></el-table-column>
        <el-table-column prop="AccName" label="付款银行账户" width="220%"></el-table-column>
        <el-table-column prop="TotAmt" label="付款金额" width="220%"></el-table-column>
        <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
        <el-table-column prop="Notes" label="描述" width="220%"></el-table-column>
        <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>

        <el-table-column label="操作" fixed="right" width="160%">
          <template #default="scope">
            <el-row type="flex" justify="space-between">
              <el-button type="text" size="small" @click="DeleteOut(scope.$index, nowID, scope.row.ID)">删除</el-button>
              <el-button type="text" size="small" @click="CheckOut(scope.row.ID)">跳转</el-button>
            </el-row>
          </template>
        </el-table-column>
      </el-table>
      <!-- 采购发货表格 -->
    </el-dialog>

    <el-dialog v-model="showShouldOutDialog" title="应付账款单" width="80%" @close="resetShouldOutForm">
      <el-form :model="shouldOutForm" label-width="150px" :rules="buyRules" ref="shouldOutFormRef">
        <!-- 第一行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="应付账款单号" prop="BillReceNum">
              <el-input v-model="shouldOutForm.BillReceNum" placeholder="请输入应付账款单号"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据日期" prop="DocDate">
              <el-date-picker v-model="shouldOutForm.DocDate" type="date" placeholder="请选择单据日期"></el-date-picker>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="预计收款日期" prop="ExpReceDate">
              <el-date-picker v-model="shouldOutForm.ExpReceDate" type="date" placeholder="请选择预计收款日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 第三行 -->

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据类型" prop="FinaDocType">
              <el-select v-model="shouldOutForm.FinaDocType" @change="onFinaDocTypeChange" placeholder="请选择财务单据类型">
                <el-option v-for="type in FinaDocTypeData" :key="type.FinaDocTypeID" :label="type.FinaDocType"
                  :value="type.FinaDocType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单据状态" prop="FinaDocStatus">
              <el-select v-model="shouldOutForm.FinaDocStatus" @change="onFinaDocStatusChange" placeholder="请选择财务单据状态">
                <el-option v-for="st in FinaDocStatusData" :key="st.FinaDocStatusID" :label="st.FinaDocStatus"
                  :value="st.FinaDocStatus"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="收款方" prop="MerchantID">
              <el-select v-model="shouldOutForm.MerchantID" @change="onMerchantChange" placeholder="请选择收款方">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="付款方" prop="AcctID">
              <el-select v-model="shouldOutForm.AcctID" @change="onAcctChange" placeholder="请选择付款方">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收款银行账户" prop="BankAccountID">
              <el-select v-model="shouldOutForm.BankAccountID" @change="onBankAccountChange" placeholder="请选择收款银行账户">
                <el-option v-for="bankAccount in bankAccountData" :key="bankAccount.ID" :label="bankAccount.BankAccName"
                  :value="bankAccount.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="付款银行账户" prop="AcctBankID">
              <el-select v-model="shouldOutForm.AcctBankID" @change="onAcctBankChange" placeholder="请选择付款银行账户">
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
              <el-input v-model="shouldOutForm.TotAmt"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="CurrencyID">
              <el-select v-model="shouldOutForm.Currency" placeholder="请选择币种">
                <el-option v-for="currency in currencyData" :key="currency.Currency" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="shouldOutForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">


              <el-upload v-if="!shouldOutForm.FileID" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(shouldOutForm.FileID, shouldOutForm.FileName)">
                下载文件
              </el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="shouldOutForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showShouldOutDialog = false">取消</el-button>
          <el-button type="primary" @click="submitShouldOutForm()">提交</el-button>
        </span>
      </template>
    </el-dialog>


    <el-dialog v-model="showshowShouldOutDialog" title="应付账款单" width="80%" @close="resetShouldOutForm">
      <el-form :model="shouldOutForm" label-width="150px" :rules="buyRules" ref="shouldOutFormRef">
        <!-- 第一行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="应付账款单号" prop="BillReceNum">
              <el-input v-model="shouldOutForm.BillReceNum" placeholder="请输入应付账款单号" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据日期" prop="DocDate">
              <el-date-picker v-model="shouldOutForm.DocDate" type="date" placeholder="请选择单据日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="预计收款日期" prop="ExpReceDate">
              <el-date-picker v-model="shouldOutForm.ExpReceDate" type="date" placeholder="请选择预计收款日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 第三行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据类型" prop="FinaDocType">
              <el-select v-model="shouldOutForm.FinaDocType" @change="onFinaDocTypeChange" placeholder="请选择财务单据类型"
                :disabled="true">
                <el-option v-for="type in FinaDocTypeData" :key="type.FinaDocTypeID" :label="type.FinaDocType"
                  :value="type.FinaDocType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单据状态" prop="FinaDocStatus">
              <el-select v-model="shouldOutForm.FinaDocStatus" @change="onFinaDocStatusChange" placeholder="请选择财务单据状态"
                :disabled="true">
                <el-option v-for="st in FinaDocStatusData" :key="st.FinaDocStatusID" :label="st.FinaDocStatus"
                  :value="st.FinaDocStatus"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收款方" prop="MerchantID">
              <el-select v-model="shouldOutForm.MerchantID" @change="onMerchantChange" placeholder="请选择收款方"
                :disabled="true">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="付款方" prop="AcctID">
              <el-select v-model="shouldOutForm.AcctID" @change="onAcctChange" placeholder="请选择付款方" :disabled="true">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收款银行账户" prop="BankAccountID">
              <el-select v-model="shouldOutForm.BankAccountID" @change="onBankAccountChange" placeholder="请选择收款银行账户"
                :disabled="true">
                <el-option v-for="bankAccount in bankAccountData" :key="bankAccount.ID" :label="bankAccount.BankAccName"
                  :value="bankAccount.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="付款银行账户" prop="AcctBankID">
              <el-select v-model="shouldOutForm.AcctBankID" @change="onAcctBankChange" placeholder="请选择付款银行账户"
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
              <el-input v-model="shouldOutForm.TotAmt" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="CurrencyID">
              <el-select v-model="shouldOutForm.Currency" placeholder="请选择币种" :disabled="true">
                <el-option v-for="currency in currencyData" :key="currency.Currency" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="shouldOutForm.Notes" type="textarea" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!shouldOutForm.FileID" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true" :disabled="true">
                <el-button type="primary" :disabled="true">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(shouldOutForm.FileID, shouldOutForm.FileName)">
                下载文件
              </el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="shouldOutForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showShouldOutDialog = false">取消</el-button>
          <!-- 隐藏提交按钮 -->
          <!-- <el-button type="primary" @click="submitShouldOutForm()">提交</el-button> -->
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
const onlyID = ref(true);
const toggleMatchMode = () => {
  isExactMatch.value = !isExactMatch.value;
};

const toggleIDMode = () => {
  onlyID.value = !onlyID.value;
};

const costInfoVisible = ref(false);

const costInfoID = ref(null);

const costInfoData = ref([]); // 存储产品明细数据
const fetchCostInfoData = async (ID) => {
  console.log("id ? ", ID)

  try {
    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/shouldOut/costInfo', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    });
    costInfoVisible.value = true;
    // prdtInfoData.value = response.data.PrdtInfo; // 假设返回的数据结构中有 PrdtInfo 字段

    costInfoData.value = Object.assign([], response.data.CostInfo); // 强制更新
    nowID.value = ID
    console.log("nowid", nowID.value)
  } catch (error) {
    console.error('获取产品明细失败:', error);
    ElMessage.error('获取产品明细失败');
  }
};
const addCostInfo = async (ID) => {

  console.log(nowID.value)
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("CostInfoID", costInfoID.value)

    const response = await axios.post('/add/shouldOut/costInfo', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    await fetchCostInfoData(nowID.value)
    prdtInfoID.value = ''
  } catch (error) {
    console.error('添加失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

const CheckCostInfo = (ID) => {
  try {
    // 确保 ID 是字符串
    const searchQuery = String(ID);

    // 使用路由的 resolve 方法生成完整路径
    const route = router.resolve({
      name: 'Cost', // 路由名称
      query: { searchQuery }, // 传递的查询参数（对象形式）
    });

    // 在新标签页打开
    window.open(route.href, '_blank');
  } catch (error) {
    ElMessage.error("查看失败");
  }
};
const OutVisible = ref(false)
const OutID = ref(null)
const OutData = ref([])

const DeleteOut = (index, ID, OutID) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该产品信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    console.log('Confirmed delete', ID); // 添加调试信息

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("OutID", OutID)

    axios.post('/delete/shouldOut/out', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })

      // axios.post('/delete/sale/prdtInfo', {
      //   "ID": ID,
      //   "PrdtInfoID": PrdtInfoID.value
      // })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchOutData(nowID.value); // 重新获取会计实体信息数据
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
const addOut = async (ID) => {

  console.log(nowID.value)
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("OutID", OutID.value)

    const response = await axios.post('/add/shouldOut/out', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchOutData(nowID.value)
    OutVisible.value = ''
  } catch (error) {
    console.error('添加失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};
const CheckOut = (ID) => {
  try {
    // 确保 ID 是字符串
    const searchQuery = String(ID);

    // 使用路由的 resolve 方法生成完整路径
    const route = router.resolve({
      name: 'Out', // 路由名称
      query: { searchQuery }, // 传递的查询参数（对象形式）
    });

    // 在新标签页打开
    window.open(route.href, '_blank');
  } catch (error) {
    ElMessage.error("查看失败");
  }
};
const BuyVisible = ref(false)
const BuyID = ref(null)
const BuyData = ref([])

const DeleteBuy = (index, ID, BuyID) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该产品信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    console.log('Confirmed delete', ID); // 添加调试信息

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("BuyID", BuyID)

    axios.post('/delete/shouldOut/buy', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })

      // axios.post('/delete/sale/prdtInfo', {
      //   "ID": ID,
      //   "PrdtInfoID": PrdtInfoID.value
      // })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchBuyData(nowID.value); // 重新获取会计实体信息数据
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
const addBuy = async (ID) => {

  console.log(nowID.value)
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("BuyID", BuyID.value)

    const response = await axios.post('/add/shouldOut/buy', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchBuyData(nowID.value)
    BuyVisible.value = ''
  } catch (error) {
    console.error('添加失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};
const CheckBuy = (ID) => {
  try {
    // 确保 ID 是字符串
    const searchQuery = String(ID);

    // 使用路由的 resolve 方法生成完整路径
    const route = router.resolve({
      name: 'Buy', // 路由名称
      query: { searchQuery }, // 传递的查询参数（对象形式）
    });

    // 在新标签页打开
    window.open(route.href, '_blank');
  } catch (error) {
    ElMessage.error("查看失败");
  }
};
const shouldOutForm = ref({
  BillReceNum: '',
  DocDate: '', // 单据日期
  ExpReceDate: '', // 预计收款日期
  FinaDocType: '', // 财务单据类型
  FinaDocStatus: '', // 财务单据状态
  MerchantID: '', // 收款方 ID
  Merc: '', // 收款方名称
  AcctID: '', // 付款方 ID
  AcctName: '', // 付款方名称
  BankAccountID: '', // 收款银行账户 ID
  BankAccName: '', // 收款银行账户名称
  AcctBankID: '', // 付款银行账户 ID
  AccName: '', // 付款银行账户名称
  TotAmt: '', // 总金额
  Currency: '', // 币种
  Notes: '', // 描述
  FileID: '', // 文件 ID
  FileName: '', // 文件名
});

// 表单验证规则
const shouldOutRules = {
  BillReceNum: [{ required: true, message: '请输入应付账款单号', trigger: 'blur' }],
  TotAmt: [{ type: 'number', message: '必须为正数', trigger: 'blur' }],
};
// 控制主弹窗显示
const purRecVisible = ref(false);

const purRecData = ref([]);
const purRecID = ref(null);

const CheckPurRec = (ID) => {
  try {
    // 确保 ID 是字符串
    const searchQuery = String(ID);

    // 使用路由的 resolve 方法生成完整路径
    const route = router.resolve({
      name: 'Purrec', // 路由名称
      query: { searchQuery }, // 传递的查询参数（对象形式）
    });

    // 在新标签页打开
    window.open(route.href, '_blank');
  } catch (error) {
    ElMessage.error("查看失败");
  }
};
// 删除按钮逻辑
const DeletePurRec = (index, ID, PurRecID) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonTitle: '取消',
    type: 'warning'
  }).then(() => {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("PurrecID", PurRecID)

    axios.post('/delete/shouldOut/purrec', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchPurRecData(nowID.value); // 重新获取会计实体信息数据
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
const addPurRec = async (ID) => {

  console.log(nowID.value)
  // acctForm.value.ID = parseInt(acctForm.value.ID, 10);

  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("PurRecID", purRecID.value)

    const response = await axios.post('/add/shouldOut/purrec', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchPurRecData(nowID.value)
    purRecID.value = ''
  } catch (error) {
    console.error('添加产品明细失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

const fetchPurRecData = async (ID) => {
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/shouldOut/purrec', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    purRecData.value = response.data.Purrec; // 假设返回的数据结构中有 PrdtInfo 字段
    nowID.value = ID
    purRecVisible.value = true;
    console.log("nowid", nowID.value)
  } catch (error) {
    console.error('获取产品明细失败:', error);
    // ElMessage.error('获取采购收货单失败');
    ElMessage.error(error.response.data.RetMessage);
  }
};
const nowID = ref(null);

const file = ref(null);
const searchQuery = ref(''); // 添加搜索查询字段
const currentPage = ref(1); // 当前页码
const pageSize = 8; // 每页显示的行数



const shouldOutData = ref([])

const fetchShouldOutData = async () => {
  try {
    const response = await axios.get('/find/shouldOut'); // 调用货币数据接口
    shouldOutData.value = response.data.ShouldOut; // 假设返回的数据结构中有 Currency 字段
    console.log(shouldOutData.value)
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

const paginatedBuyData = computed(() => {
  let filteredData = shouldOutData.value;


  if (searchQuery.value) {
    console.log(isExactMatch.value);
    console.log(onlyID.value);

    if (isExactMatch.value === false) {
      if (onlyID.value === false) {
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value) ||
          item.BillReceNum.includes(searchQuery.value) ||
          (item.Acct && item.Acct.AcctName && item.Acct.AcctName.includes(searchQuery.value)) ||
          (item.Merchant && item.Merchant.Merc && item.Merchant.Merc.includes(searchQuery.value)) ||
          (item.AcctBank && item.AcctBank.AccName && item.AcctBank.AccName.includes(searchQuery.value)) ||
          (item.BankAccount && item.BankAccount.BankAccName && item.BankAccount.BankAccName.includes(searchQuery.value)) ||
          item.DocDate.includes(searchQuery.value) ||
          item.ExpReceDate.includes(searchQuery.value) ||
          item.FinaDocType.includes(searchQuery.value) ||
          item.FinaDocStatus.includes(searchQuery.value) ||
          item.TotAmt.toString().includes(searchQuery.value) ||
          item.Currency.includes(searchQuery.value) ||
          item.Notes.includes(searchQuery.value) ||
          item.FileName.includes(searchQuery.value)
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
          item.BillReceNum === searchQuery.value ||
          (item.Acct && item.Acct.AcctName && item.Acct.AcctName === searchQuery.value) ||
          (item.Merchant && item.Merchant.Merc && item.Merchant.Merc === searchQuery.value) ||
          (item.AcctBank && item.AcctBank.AccName && item.AcctBank.AccName === searchQuery.value) ||
          (item.BankAccount && item.BankAccount.BankAccName && item.BankAccount.BankAccName === searchQuery.value) ||
          item.DocDate === searchQuery.value ||
          item.ExpReceDate === searchQuery.value ||
          item.FinaDocType === searchQuery.value ||
          item.FinaDocStatus === searchQuery.value ||
          item.TotAmt.toString() === searchQuery.value ||
          item.Currency === searchQuery.value ||
          item.Notes === searchQuery.value ||
          item.FileName === searchQuery.value
        );
      } else {
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
  fetchShouldOutData();
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

const fetchBuyData = async (ID) => {
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/shouldOut/buy', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    BuyData.value = response.data.Buy; // 假设返回的数据结构中有 PrdtInfo 字段
    nowID.value = ID
    BuyVisible.value = true;
  } catch (error) {
    console.error('获取失败:', error);
    ElMessage.error('获取失败');
  }
};

const fetchOutData = async (ID) => {
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/shouldOut/out', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    OutData.value = response.data.Out; // 假设返回的数据结构中有 PrdtInfo 字段
    nowID.value = ID
    OutVisible.value = true;
  } catch (error) {
    console.error('获取失败:', error);
    ElMessage.error('获取失败');
  }
};

// 应付账款单信息对话框显示状态
const showShouldOutDialog = ref(false);
const showshowShouldOutDialog = ref(false);

// 表格数据（初始值为空数组）
const acctData = ref([]); // 会计实体信息
const merchantData = ref([]); // 购买方信息
const acctBankData = ref([]); // 我方银行账户数据
const bankAccountData = ref([]); // 对方银行账户数据
const buyData = ref([]); // 采购订单信息

// 根据当前选中的菜单项动态更改标题和按钮文本
const headerTitle = computed(() => {
  return '应付账款单信息';
});

const addButtonText = computed(() => {
  return '添加应付账款单信息';
});

// 添加按钮点击事件
const handleAdd = () => {
  showShouldOutDialog.value = true;
};

const handleEdit = (index, row) => {
  // 将当前行的数据赋值给 saleForm
  shouldOutForm.value = { ...row };

  // 检查是否已上传文件
  if (row.FileID) {
    shouldOutForm.value.FileID = row.FileID; // 保存 FileID
    shouldOutForm.value.FileName = row.FileName; // 保存文件名
  }

  showShouldOutDialog.value = true; // 打开销售订单信息对话框
};

// 查看按钮逻辑
const handleView = (index, row) => {
  shouldOutForm.value = { ...row }; // 将当前行的数据赋值给 saleForm
  showshowShouldOutDialog.value = true; // 打开查看销售订单信息对话框
};

// 删除按钮逻辑
const handleDelete = (index, ID) => {
  ElMessageBox.confirm('确定要删除该应付账款单信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonTitle: '取消',
    type: 'warning'
  }).then(() => {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    axios.post('/delete/shouldOut', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchShouldOutData(); // 重新获取销售订单信息数据
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
const resetShouldOutForm = () => {

  shouldOutForm.value = {
    BillReceNum: '',
    DocDate: '', // 单据日期
    ExpReceDate: '', // 预计收款日期
    FinaDocType: '', // 财务单据类型
    FinaDocStatus: '', // 财务单据状态
    MerchantID: '', // 收款方 ID
    Merc: '', // 收款方名称
    AcctID: '', // 付款方 ID
    AcctName: '', // 付款方名称
    BankAccountID: '', // 收款银行账户 ID
    BankAccName: '', // 收款银行账户名称
    AcctBankID: '', // 付款银行账户 ID
    AccName: '', // 付款银行账户名称
    TotAmt: '', // 总金额
    Currency: '', // 币种
    Notes: '', // 描述
    FileID: '', // 文件 ID
    FileName: '', // 文件名
  };

  file.value = null; // 重置文件对象
  if (uploadRef.value) {
    uploadRef.value.clearFiles(); // 清空文件列表
  }
};

const uploadRef = ref(null);

const submitShouldOutForm = async () => {
  try {
    const formData = new FormData();

    // 添加其他字段
    Object.keys(shouldOutForm.value).forEach((key) => {
      formData.append(key, shouldOutForm.value[key]);
    });


    // 添加文件
    if (file.value) {
      formData.append('file', file.value);
    }

    // 提交表单

    const response = await axios.post('/save/shouldOut', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    if (response.status === 200) {
      ElMessage.success('采购订单信息保存成功');
      showShouldOutDialog.value = false; // 关闭对话框
      fetchShouldOutData(); // 重新获取销售订单信息数据
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存采购订单信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};
// 监听 change 事件并更新 AcctName
const onAcctChange = (value) => {
  const selectedAcct = acctData.value.find(acct => acct.ID === value);
  if (selectedAcct) {
    shouldOutForm.value.AcctName = selectedAcct.AcctName;
  }
};

const downloadFile = async (fileID, fileName) => {
  try {
    const response = await axios.post(
      '/file',
      { FileID: fileID },
      {
        responseType: 'blob',
      }
    );

    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', fileName || `file_${fileID}`);
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
  shouldOutForm.value.FileName = uploadFile.name; // 更新文件名
};

// 监听 change 事件并更新 Merc
const onMerchantChange = (value) => {
  const selectedMerchant = merchantData.value.find(merchant => merchant.ID === value);
  if (selectedMerchant) {
    shouldOutForm.value.Merc = selectedMerchant.Merc;
  }
};

// 监听 change 事件并更新 AccName
const onAcctBankChange = (value) => {
  const selectedAcctBank = acctBankData.value.find(acctBank => acctBank.ID === value);
  if (selectedAcctBank) {
    shouldOutForm.value.AccName = selectedAcctBank.AccName;
  }
};

// 监听 change 事件并更新 BankAccName
const onBankAccountChange = (value) => {
  const selectedBankAccount = bankAccountData.value.find(bankAccount => bankAccount.ID === value);
  if (selectedBankAccount) {
    shouldOutForm.value.BankAccName = selectedBankAccount.BankAccName;
  }
};

// 分页逻辑
const handlePageChange = (page) => {
  currentPage.value = page;
};
</script>
<style src="../assets/styles/Bottom.css"></style>

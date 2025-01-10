<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-22'" />
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>

        <HeaderComponent :header-title="headerTitle" :add-button-text="addButtonText" v-model:search-query="searchQuery"
          @toggle-match-mode="toggleMatchMode" @toggle-id-mode="toggleIDMode" @add="handleAdd" />
        <el-header height="1px">
        </el-header>

        <el-main>

          <!-- 采购订单信息表格 -->
          <el-table :data="paginatedBuyData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="ReceNum" label="付款单号" width="220%"></el-table-column>
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
            <el-table-column prop="RealReceDate" label="实际付款日期" width="220%"></el-table-column>
            <el-table-column prop="ExpReceDate" label="预计付款日期" width="220%"></el-table-column>
            <el-table-column prop="FinaDocType" label="单据类型" width="220%"></el-table-column>
            <el-table-column prop="FinaDocStatus" label="单据状态" width="220%"></el-table-column>
            <el-table-column prop="TotAmt" label="付款金额" width="220%"></el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="描述" width="220%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>

            <el-table-column label="操作" fixed="right" width="420%">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.ID)" type="text" size="small">删除</el-button>
                  <el-button @click="fetchBuyData(scope.row.ID)" type="text" size="small">采购订单</el-button>
                  <el-button @click="fetchPurrecData(scope.row.ID)" type="text" size="small">采购收货单</el-button>
                  <el-button @click="fetchShouldOutData(scope.row.ID)" type="text" size="small">应付账款单</el-button>
                  <!-- <el-button @click="fetchOutData(scope.row.ID)" type="text" size="small">付款单</el-button> -->
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="buyData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <el-dialog v-model="ShouldOutVisible" title="应付账款单" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="ShouldOutID" placeholder="请输入ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addShouldOut(nowID)">添加</el-button>
      </div>

      <!-- 产品明细表格 -->
      <el-table :data="ShouldOutData" height="400" style="width: 100%">

        <el-table-column prop="ID" label="ID" width="60%" />
        <el-table-column prop="BillReceNum" label="应付账款单号" width="220%"></el-table-column>
        <el-table-column prop="DocDate" label="单据日期" width="220%"></el-table-column>
        <el-table-column prop="ExpReceDate" label="预计付款日期" width="220%"></el-table-column>
        <el-table-column prop="FinaDocType" label="单据类型" width="220%"></el-table-column>
        <el-table-column prop="FinaDocStatus" label="单据状态" width="220%"></el-table-column>
        <!-- <el-table-column prop="Merc" label="付款方" width="220%"></el-table-column> -->
        <!-- <el-table-column prop="AcctName" label="收款方" width="220%"></el-table-column> -->
        <!-- <el-table-column prop="BankAccName" label="付款银行账户" width="220%"></el-table-column> -->
        <!-- <el-table-column prop="AccName" label="收款银行账户" width="220%"></el-table-column> -->
        <el-table-column prop="TotAmt" label="总金额" width="220%"></el-table-column>
        <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
        <el-table-column prop="Notes" label="描述" width="220%"></el-table-column>
        <el-table-column label="操作" fixed="right" width="150">
          <template #default="scope">
            <!-- <el-button type="text" size="small" @click="viewProduct(scope.row)">查看</el-button> -->

            <el-button type="text" size="small" @click="CheckShouldOut(scope.row.ID)">跳转</el-button>
            <el-button type="text" size="small"
              @click="DeleteShouldOut(scope.$index, nowID, scope.row.ID)">删除</el-button>
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

      <el-table :data="BuyData" height="400" style="width: 80%">
        <el-table-column prop="ID" label="ID" width="60%" />
        <el-table-column prop="OrderNum" label="订单编号" width="150%" />
        <el-table-column prop="Merc" label="购买方" width="150%" />
        <el-table-column prop="AcctName" label="销售方" width="150%" />
        <el-table-column prop="SpecName" label="包装规格" width="150%" />
        <el-table-column prop="TotAmt" label="总金额" width="150%" />
        <el-table-column prop="Currency" label="币种" width="150%" />
        <el-table-column label="操作" fixed="right" width="150%">
          <template #default="scope">
            <!-- <el-button type="text" size="small" @click="viewProduct(scope.row)">查看</el-button> -->

            <el-button type="text" size="small" @click="CheckBuy(scope.row.ID)">跳转</el-button>
            <el-button type="text" size="small" @click="DeleteBuy(scope.$index, nowID, scope.row.ID)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="purrecVisible" title="采购收货单" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="purrecID" placeholder="请输入采购收货单ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addPurrec(nowID)">添加</el-button>
      </div>

      <!-- 采购收货单表格 -->
      <el-table :data="purrecData" height="400" style="width: 80%">
        <el-table-column prop="ID" label="ID" width="60%" />
        <el-table-column prop="SaleInvNum" label="采购发票号" width="150%" />
        <!-- <el-table-column prop="AcctName" label="发货人" width="150%" /> -->
        <el-table-column prop="Note1" label="提单货物描述" width="360%" />
        <el-table-column prop="SaleInvDate" label="采购发票日期" width="420%" />
        <el-table-column label="操作" fixed="right" width="150%">
          <template #default="scope">
            <!-- <el-button type="text" size="small" @click="viewProduct(scope.row)">查看</el-button> -->
            <el-button type="text" size="small" @click="DeletePurrec(scope.$index, nowID, scope.row.ID)">删除</el-button>

            <el-button type="text" size="small" @click="CheckPurrec(scope.row.ID)">跳转</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog v-model="showOutDialog" title="付款单" width="80%" @close="resetOutForm">
      <el-form :model="OutForm" label-width="150px" :rules="Rules" ref="OutFormRef">
        <!-- 第一行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="付款单号" prop="ReceNum">
              <el-input v-model="OutForm.ReceNum" placeholder="请输入付款单号"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="实际付款日期" prop="RealReceDate">
              <el-date-picker v-model="OutForm.RealReceDate" type="date" placeholder="请选择单据日期"></el-date-picker>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="预计付款日期" prop="ExpReceDate">
              <el-date-picker v-model="OutForm.ExpReceDate" type="date" placeholder="请选择预计付款日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 第三行 -->

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据类型" prop="FinaDocType">
              <el-select v-model="OutForm.FinaDocType" @change="onFinaDocTypeChange" placeholder="请选择财务单据类型">
                <el-option v-for="type in FinaDocTypeData" :key="type.FinaDocTypeID" :label="type.FinaDocType"
                  :value="type.FinaDocType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单据状态" prop="FinaDocStatus">
              <el-select v-model="OutForm.FinaDocStatus" @change="onFinaDocStatusChange" placeholder="请选择财务单据状态">
                <el-option v-for="st in FinaDocStatusData" :key="st.FinaDocStatusID" :label="st.FinaDocStatus"
                  :value="st.FinaDocStatus"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="收款方" prop="MerchantID">
              <el-select v-model="OutForm.MerchantID" @change="onMerchantChange" placeholder="请选择收款方">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="付款方" prop="AcctID">
              <el-select v-model="OutForm.AcctID" @change="onAcctChange" placeholder="请选择付款方">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收款银行账户" prop="BankAccountID">
              <el-select v-model="OutForm.BankAccountID" @change="onBankAccountChange" placeholder="请选择收款银行账户">
                <el-option v-for="bankAccount in bankAccountData" :key="bankAccount.ID" :label="bankAccount.BankAccName"
                  :value="bankAccount.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="付款银行账户" prop="AcctBankID">
              <el-select v-model="OutForm.AcctBankID" @change="onAcctBankChange" placeholder="请选择付款银行账户">
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
              <el-input v-model="OutForm.TotAmt"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="CurrencyID">
              <el-select v-model="OutForm.Currency" placeholder="请选择币种">
                <el-option v-for="currency in currencyData" :key="currency.Currency" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="OutForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">


              <el-upload v-if="!OutForm.FileID" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(OutForm.FileID, OutForm.FileName)">
                下载文件
              </el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="OutForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showOutDialog = false">取消</el-button>
          <el-button type="primary" @click="submitOutForm()">提交</el-button>
        </span>
      </template>
    </el-dialog>


    <el-dialog v-model="showshowOutDialog" title="付款单" width="80%" @close="resetOutForm">
      <el-form :model="OutForm" label-width="150px" ref="OutFormRef">
        <!-- 第一行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="付款单号">
              <el-input v-model="OutForm.ReceNum" :readonly="true" placeholder="请输入付款单号"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="实际付款日期">
              <el-date-picker v-model="OutForm.RealReceDate" type="date" placeholder="请选择单据日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="预计付款日期">
              <el-date-picker v-model="OutForm.ExpReceDate" type="date" placeholder="请选择预计付款日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 第三行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单据类型">
              <el-select v-model="OutForm.FinaDocType" placeholder="请选择财务单据类型" :disabled="true">
                <el-option v-for="type in FinaDocTypeData" :key="type.FinaDocTypeID" :label="type.FinaDocType"
                  :value="type.FinaDocType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单据状态">
              <el-select v-model="OutForm.FinaDocStatus" placeholder="请选择财务单据状态" :disabled="true">
                <el-option v-for="st in FinaDocStatusData" :key="st.FinaDocStatusID" :label="st.FinaDocStatus"
                  :value="st.FinaDocStatus"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收款方">
              <el-select v-model="OutForm.MerchantID" placeholder="请选择收款方" :disabled="true">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="付款方">
              <el-select v-model="OutForm.AcctID" placeholder="请选择付款方" :disabled="true">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收款银行账户">
              <el-select v-model="OutForm.BankAccountID" placeholder="请选择收款银行账户" :disabled="true">
                <el-option v-for="bankAccount in bankAccountData" :key="bankAccount.ID" :label="bankAccount.BankAccName"
                  :value="bankAccount.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="付款银行账户">
              <el-select v-model="OutForm.AcctBankID" placeholder="请选择付款银行账户" :disabled="true">
                <el-option v-for="acctBank in acctBankData" :key="acctBank.ID" :label="acctBank.AccName"
                  :value="acctBank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 第四行 -->
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总金额">
              <el-input v-model="OutForm.TotAmt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种">
              <el-select v-model="OutForm.Currency" placeholder="请选择币种" :disabled="true">
                <el-option v-for="currency in currencyData" :key="currency.Currency" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注">
              <el-input v-model="OutForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-button v-if="OutForm.FileID" type="success" @click="downloadFile(OutForm.FileID, OutForm.FileName)">
                下载文件
              </el-button>
              <span v-else>无文件</span>
            </el-form-item>

            <el-form-item label="文件名">
              <el-input v-model="OutForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showOutDialog = false">关闭</el-button>
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

const CheckShouldOut = (ID) => {
  try {
    // 确保 ID 是字符串
    const searchQuery = String(ID);
    console.log(searchQuery)

    // 使用路由的 resolve 方法生成完整路径
    const route = router.resolve({
      name: 'ShouldOut', // 路由名称
      query: { searchQuery }, // 传递的查询参数（对象形式）
    });

    // 在新标签页打开
    window.open(route.href, '_blank');
  } catch (error) {
    ElMessage.error("查看失败");
  }
};
const ShouldOutVisible = ref(false)
const ShouldOutID = ref(null)
const ShouldOutData = ref([])

const fetchShouldOutData = async (ID) => {
  console.log("id ? ", ID)

  try {
    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/out/shouldOut', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    });
    ShouldOutVisible.value = true;

    ShouldOutData.value = Object.assign([], response.data.ShouldOut); // 强制更新
    nowID.value = ID

    console.log("nowid", nowID.value)
  } catch (error) {
    console.error('获取失败:', error);
    ElMessage.error('获取失败');
  }
};
const addShouldOut = async (ID) => {

  console.log(nowID.value)
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("ShouldOutID", ShouldOutID.value)

    const response = await axios.post('/add/out/shouldOut', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchShouldOutData(nowID.value)
    ShouldOutVisible.value = ''
  } catch (error) {
    console.error('添加失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

const DeleteShouldOut = (index, ID, id) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该产品信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    console.log('Confirmed delete', ID); // 添加调试信息

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("ShouldOutID", id)

    axios.post('/delete/out/shouldOut', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchShouldOutData(nowID.value); // 重新获取应付账款单信息数据
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

    axios.post('/delete/out/buy', params, {
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
          fetchBuyData(nowID.value); // 重新获取产品信息数据
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

    const response = await axios.post('/add/out/buy', params, {
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
const OutForm = ref({
  ReceNum: '', // 付款单号
  RealReceDate: '', // 单据日期
  ExpReceDate: '', // 预计付款日期
  FinaDocType: '', // 财务单据类型
  FinaDocStatus: '', // 财务单据状态
  MerchantID: '', // 付款方 ID
  Merc: '', // 付款方名称
  AcctID: '', // 收款方 ID
  AcctName: '', // 收款方名称
  BankAccountID: '', // 付款银行账户 ID
  BankAccName: '', // 付款银行账户名称
  AcctBankID: '', // 收款银行账户 ID
  AccName: '', // 收款银行账户名称
  TotAmt: '', // 总金额
  Currency: '', // 币种
  Notes: '', // 描述
  FileID: '', // 文件 ID
  FileName: '', // 文件名
});

// 控制主弹窗显示
const purrecVisible = ref(false);

const purrecData = ref([]);
const purrecID = ref(null);

const CheckPurrec = (ID) => {
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
const DeletePurrec = (index, ID, PurrecID) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该采购收货单信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("PurrecID", PurrecID)

    axios.post('/delete/out/purrec', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchPurrecData(nowID.value); // 重新获取采购收货单信息数据
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
const addPurrec = async (ID) => {

  console.log(nowID.value)
  // acctForm.value.ID = parseInt(acctForm.value.ID, 10);

  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("PurrecID", purrecID.value)

    const response = await axios.post('/add/out/purrec', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchPurrecData(nowID.value)
    purrecID.value = ''
  } catch (error) {
    console.error('添加产品明细失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

const fetchPurrecData = async (ID) => {
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/out/purrec', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    purrecData.value = response.data.Purrec; // 假设返回的数据结构中有 Purrec 字段
    nowID.value = ID
    purrecVisible.value = true;
    console.log("nowid", nowID.value)
  } catch (error) {
    console.error('获取产品明细失败:', error);
    ElMessage.error('获取产品明细失败');
  }
};
const nowID = ref(null);

const file = ref(null);
const searchQuery = ref(''); // 添加搜索查询字段
const currentPage = ref(1); // 当前页码
const pageSize = 8; // 每页显示的行数



const OutData = ref([])

const fetchOutData = async () => {
  try {
    const response = await axios.get('/find/out'); // 调用货币数据接口
    OutData.value = response.data.Out; // 假设返回的数据结构中有 Currency 字段
    console.log(OutData.value)
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
  let filteredData = OutData.value;


  if (searchQuery.value) {
    console.log(isExactMatch.value);
    console.log(onlyID.value);

    if (isExactMatch.value === false) {
      if (onlyID.value === false) {
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value) ||
          item.ReceNum.includes(searchQuery.value) ||
          (item.Acct && item.Acct.AcctName && item.Acct.AcctName.includes(searchQuery.value)) ||
          (item.Merchant && item.Merchant.Merc && item.Merchant.Merc.includes(searchQuery.value)) ||
          (item.AcctBank && item.AcctBank.AccName && item.AcctBank.AccName.includes(searchQuery.value)) ||
          (item.BankAccount && item.BankAccount.BankAccName && item.BankAccount.BankAccName.includes(searchQuery.value)) ||
          item.RealReceDate.includes(searchQuery.value) ||
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
          item.ReceNum === searchQuery.value ||
          (item.Acct && item.Acct.AcctName && item.Acct.AcctName === searchQuery.value) ||
          (item.Merchant && item.Merchant.Merc && item.Merchant.Merc === searchQuery.value) ||
          (item.AcctBank && item.AcctBank.AccName && item.AcctBank.AccName === searchQuery.value) ||
          (item.BankAccount && item.BankAccount.BankAccName && item.BankAccount.BankAccName === searchQuery.value) ||
          item.RealReceDate === searchQuery.value ||
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
  fetchOutData();
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

    const response = await axios.post('/find/out/buy', params, {
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

// 付款单信息对话框显示状态
const showOutDialog = ref(false);
const showshowOutDialog = ref(false);

// 表格数据（初始值为空数组）
const acctData = ref([]); // 会计实体信息
const merchantData = ref([]); // 购买方信息
const acctBankData = ref([]); // 我方银行账户数据
const bankAccountData = ref([]); // 对方银行账户数据
const buyData = ref([]); // 采购订单信息

// 根据当前选中的菜单项动态更改标题和按钮文本
const headerTitle = computed(() => {
  return '付款单信息';
});

const addButtonText = computed(() => {
  return '添加付款单信息';
});

// 添加按钮点击事件
const handleAdd = () => {
  showOutDialog.value = true;
};

const handleEdit = (index, row) => {
  // 将当前行的数据赋值给 OutForm
  OutForm.value = { ...row };

  // 检查是否已上传文件
  if (row.FileID) {
    OutForm.value.FileID = row.FileID; // 保存 FileID
    OutForm.value.FileName = row.FileName; // 保存文件名
  }

  showOutDialog.value = true; // 打开付款单信息对话框
};

// 查看按钮逻辑
const handleView = (index, row) => {
  OutForm.value = { ...row }; // 将当前行的数据赋值给 OutForm
  showshowOutDialog.value = true; // 打开查看付款单信息对话框
};

// 删除按钮逻辑
const handleDelete = (index, ID) => {
  ElMessageBox.confirm('确定要删除该付款单信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    axios.post('/delete/out', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchOutData(); // 重新获取付款单信息数据
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
// 重置付款单信息表单
const resetOutForm = () => {

  OutForm.value = {
    ReceNum: '', // 付款单号
    RealReceDate: '', // 单据日期
    ExpReceDate: '', // 预计付款日期
    FinaDocType: '', // 财务单据类型
    FinaDocStatus: '', // 财务单据状态
    MerchantID: '', // 付款方 ID
    Merc: '', // 付款方名称
    AcctID: '', // 收款方 ID
    AcctName: '', // 收款方名称
    BankAccountID: '', // 付款银行账户 ID
    BankAccName: '', // 付款银行账户名称
    AcctBankID: '', // 收款银行账户 ID
    AccName: '', // 收款银行账户名称
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

const submitOutForm = async () => {
  try {
    const formData = new FormData();

    // 添加其他字段
    Object.keys(OutForm.value).forEach((key) => {
      formData.append(key, OutForm.value[key]);
    });

    // 添加文件
    if (file.value) {
      formData.append('file', file.value);
    }

    // 提交表单
    const response = await axios.post('/save/out', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    if (response.status === 200) {
      ElMessage.success('付款单信息保存成功');
      showOutDialog.value = false; // 关闭对话框
      fetchOutData(); // 重新获取付款单信息数据
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存付款单信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};
// 监听 change 事件并更新 AcctName
const onAcctChange = (value) => {
  const selectedAcct = acctData.value.find(acct => acct.ID === value);
  if (selectedAcct) {
    OutForm.value.AcctName = selectedAcct.AcctName;
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
  OutForm.value.FileName = uploadFile.name; // 更新文件名
};

// 监听 change 事件并更新 Merc
const onMerchantChange = (value) => {
  const selectedMerchant = merchantData.value.find(merchant => merchant.ID === value);
  if (selectedMerchant) {
    OutForm.value.Merc = selectedMerchant.Merc;
  }
};

// 监听 change 事件并更新 AccName
const onAcctBankChange = (value) => {
  const selectedAcctBank = acctBankData.value.find(acctBank => acctBank.ID === value);
  if (selectedAcctBank) {
    OutForm.value.AccName = selectedAcctBank.AccName;
  }
};

// 监听 change 事件并更新 BankAccName
const onBankAccountChange = (value) => {
  const selectedBankAccount = bankAccountData.value.find(bankAccount => bankAccount.ID === value);
  if (selectedBankAccount) {
    OutForm.value.BankAccName = selectedBankAccount.BankAccName;
  }
};

// 分页逻辑
const handlePageChange = (page) => {
  currentPage.value = page;
};
</script>

<style src="../assets/styles/Bottom.css"></style>

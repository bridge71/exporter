<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-16'" />
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>

        <HeaderComponent :header-title="headerTitle" :add-button-text="addButtonText" v-model:search-query="searchQuery"
          @toggle-match-mode="toggleMatchMode" @toggle-id-mode="toggleIDMode" @add="handleAdd" />
        <el-header height="1px">
        </el-header>

        <el-main>

          <!-- 销售订单信息表格 -->
          <el-table :data="paginatedBuyData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="OrderNum" label="订单编号" width="220%"></el-table-column>
            <el-table-column prop="AcctName" label="销售方" width="220%"></el-table-column>
            <el-table-column prop="Merc" label="购买方" width="220%"></el-table-column>
            <el-table-column prop="QualStd" label="质量标准" width="220%"></el-table-column>
            <el-table-column prop="OrderDate" label="订单日期" width="420%"></el-table-column>
            <el-table-column prop="BillValidity" label="账单有效期" width="420%"></el-table-column>
            <el-table-column prop="BussOrderSta" label="单据状态" width="220%"></el-table-column>
            <el-table-column label="单据要求" width="700%">
              <template #default="scope">
                <div v-if="scope.row.DocReq && scope.row.DocReq.length > 0">
                  <el-tag v-for="(docReq, index) in scope.row.DocReq" :key="index" type="info" style="margin: 2px;">
                    {{ docReq.DocReq || '未知单据要求' }}
                  </el-tag>
                </div>
                <span v-else>无单据要求</span>
              </template>
            </el-table-column>
            <el-table-column prop="StartShip" label="发货开始日期" width="420%"></el-table-column>
            <el-table-column prop="EndShip" label="发货截止日期" width="420%"></el-table-column>
            <el-table-column prop="SrcPlace" label="起运地" width="220%"></el-table-column>
            <el-table-column prop="Des" label="目的地" width="220%"></el-table-column>
            <el-table-column prop="PayMtdName" label="付款方式" width="220%"></el-table-column>
            <el-table-column prop="TotAmt" label="总金额" width="220%"></el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="TotNum" label="总件数" width="220%"></el-table-column>
            <el-table-column prop="SpecName" label="包装规格" width="220%"></el-table-column>
            <el-table-column prop="TotalNetWeight" label="总净重" width="220%"></el-table-column>
            <el-table-column prop="UnitMeas" label="单位" width="220%"></el-table-column>
            <el-table-column prop="AccName" label="我方银行账户" width="220%"></el-table-column>
            <el-table-column prop="BankAccName" label="对方银行账户" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="220%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>

            <el-table-column label="操作" fixed="right" width="420%">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.ID)" type="text" size="small">删除</el-button>
                  <el-button @click="fetchPrdtInfoData(scope.row.ID)" type="text" size="small">产品明细</el-button>
                  <el-button @click="fetchPurrecData(scope.row.ID)" type="text" size="small">采购收货单</el-button>

                  <el-button @click="fetchShouldOutData(scope.row.ID)" type="text" size="small">应付账款单</el-button>
                  <el-button @click="fetchOutData(scope.row.ID)" type="text" size="small">付款单</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="buyData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <el-dialog v-model="OutVisible" title="付款单" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="OutID" placeholder="请输入付款单ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addOut(nowID)">添加</el-button>
      </div>
      <el-table :data="OutData" style="width: 100%" max-height="450">
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
              <el-button type="text" size="small" @click="DeleteOut(scope.$index, nowID, scope.row.ID)">删除</el-button>
              <el-button type="text" size="small" @click="CheckOut(scope.row.ID)">跳转</el-button>
            </el-row>
          </template>
        </el-table-column>
      </el-table>
      <!-- 销售发货表格 -->
    </el-dialog>

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
        <el-table-column prop="Merc" label="付款方" width="220%"></el-table-column>
        <el-table-column prop="AcctName" label="收款方" width="220%"></el-table-column>
        <el-table-column prop="BankAccName" label="付款银行账户" width="220%"></el-table-column>
        <el-table-column prop="AccName" label="收款银行账户" width="220%"></el-table-column>
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
    <!-- 弹出表单 -->
    <el-dialog v-model="prdtInfoVisible" title="产品明细" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="prdtInfoID" placeholder="请输入产品ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addPrdtInfo(nowID)">添加</el-button>
      </div>

      <!-- 产品明细表格 -->
      <el-table :data="prdtInfoData" height="400" style="width: 100%">
        <el-table-column prop="ID" label="ID" width="60%" />
        <el-table-column prop="CatEngName" label="产品" width="150%" />
        <el-table-column prop="BrandEngName" label="品牌" width="150%" />
        <el-table-column prop="PackSpec" label="包装规格" width="150%" />
        <el-table-column prop="Currency" label="币种" width="100%" />
        <el-table-column prop="UnitPrice" label="单价" width="70%" />
        <el-table-column prop="TradeTerm" label="贸易条款" width="100%" />
        <el-table-column prop="DeliveryLoc" label="交货地点" width="100%" />

        <el-table-column prop="Factory" label="生产工厂" width="150%" />
        <el-table-column prop="Unit" label="单位" width="150%" />
        <el-table-column prop="Amount" label="金额" width="100%" />
        <el-table-column prop="ItemNum" label="件数" width="70%" />
        <el-table-column prop="Weight" label="重量" width="100%" />
        <el-table-column prop="WeightUnit" label="重量单位" width="100%" />
        <el-table-column label="操作" fixed="right" width="150">
          <template #default="scope">
            <!-- <el-button type="text" size="small" @click="viewProduct(scope.row)">查看</el-button> -->

            <el-button type="text" size="small" @click="CheckPrdtInfo(scope.row.ID)">跳转</el-button>
            <!-- <el-button type="text" size="small" -->
            <!--   @click="DeletePrdtInfo(scope.$index, nowID, scope.row.ID)">删除</el-button> -->
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

      <!-- 销售发货表格 -->
      <el-table :data="purrecData" height="400" style="width: 100%">
        <el-table-column prop="ID" label="ID" width="60%" />
        <el-table-column prop="SaleInvNum" label="采购发票号" width="150%" />
        <el-table-column prop="AcctName" label="收货人" width="150%" />
        <el-table-column prop="Note1" label="提单货物描述" width="360%" />
        <el-table-column prop="SaleInvDate" label="采购发票日期" width="420%" />
        <el-table-column label="操作" fixed="right" width="150">
          <template #default="scope">
            <!-- <el-button type="text" size="small" @click="viewProduct(scope.row)">查看</el-button> -->
            <el-button type="text" size="small" @click="DeletePurrec(scope.$index, nowID, scope.row.ID)">删除</el-button>

            <el-button type="text" size="small" @click="CheckPurrec(scope.row.ID)">跳转</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <!-- 添加销售订单信息的对话框 -->
    <el-dialog v-model="showBuyDialog" title="采购订单信息" width="80%" @close="resetBuyForm">
      <el-form :model="buyForm" label-width="150px" :rules="buyRules" ref="buyFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="订单编号" prop="OrderNum">
              <el-input v-model="buyForm.OrderNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="订单日期" prop="OrderDate">
              <el-date-picker v-model="buyForm.OrderDate" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="销售方" prop="AcctID">
              <el-select v-model="buyForm.AcctID" @change="onAcctChange" placeholder="请选择销售方">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="购买方" prop="MerchantID">
              <el-select v-model="buyForm.MerchantID" @change="onMerchantChange" placeholder="请选择购买方">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="单据要求" prop="DocReq">
              <el-select v-model="buyForm.DocReq" multiple placeholder="请选择单据要求">
                <el-option v-for="docReq in docReqData" :key="docReq.DocReqID" :label="docReq.DocReq"
                  :value="docReq.DocReqID">
                  {{ docReq.DocReq }}
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="付款方式" prop="PayMentMethodID">
              <el-select v-model="buyForm.PayMentMethodID" placeholder="请选择付款方式">
                <el-option v-for="payMentMethod in payMentMethodData" :key="payMentMethod.ID"
                  :label="payMentMethod.PayMtdName" :value="payMentMethod.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账单有效期" prop="BillValidity">
              <el-date-picker v-model="buyForm.BillValidity" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单据状态" prop="BussOrderSta">
              <el-select v-model="buyForm.BussOrderSta" placeholder="请选择单据状态">
                <el-option v-for="bussOrderSta in bussOrderStaData" :key="bussOrderSta.BussOrderSta"
                  :label="bussOrderSta.BussOrderSta" :value="bussOrderSta.BussOrderSta"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="发货开始日期" prop="StartShip">
              <el-date-picker v-model="buyForm.StartShip" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="发货截止日期" prop="EndShip">
              <el-date-picker v-model="buyForm.EndShip" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="起运地" prop="SrcPlace">
              <el-select v-model="buyForm.SrcPlace" placeholder="请选择起运地">
                <el-option v-for="srcPlace in srcPlaceData" :key="srcPlace" :label="srcPlace"
                  :value="srcPlace"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="目的地" prop="Des">
              <el-select v-model="buyForm.Des" placeholder="请选择目的地">
                <el-option v-for="des in desData" :key="des" :label="des" :value="des"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="9">
            <el-form-item label="质量标准" prop="QualStd">
              <el-select v-model="buyForm.QualStd" placeholder="请选择质量标准">
                <el-option v-for="qualStd in qualStdData" :key="qualStd.QualStd" :label="qualStd.QualStd"
                  :value="qualStd.QualStd"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="总金额" prop="TotAmt">
              <el-input v-model="buyForm.TotAmt"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="7">
            <el-form-item label="币种" prop="CurrencyID">
              <el-select v-model="buyForm.Currency" @change="onCurrencyChange" placeholder="请选择币种">
                <el-option v-for="currency in currencyData" :key="currency.Currency" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总件数" prop="TotNum">
              <el-input v-model="buyForm.TotNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpecID">
              <el-select v-model="buyForm.PackSpecID" @change="onPackSpecChange" placeholder="请选择包装规格">
                <el-option v-for="packSpec in packSpecData" :key="packSpec.ID" :label="packSpec.SpecName"
                  :value="packSpec.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总净重" prop="TotalNetWeight">
              <el-input v-model="buyForm.TotalNetWeight"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas">
              <el-select v-model="buyForm.UnitMeas" placeholder="请选择单位">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeas" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="我方银行账户" prop="AcctBankID">
              <el-select v-model="buyForm.AcctBankID" @change="onAcctBankChange" placeholder="请选择我方银行账户">
                <el-option v-for="acctBank in acctBankData" :key="acctBank.ID" :label="acctBank.AccName"
                  :value="acctBank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="对方银行账户" prop="BankAccountID">
              <el-select v-model="buyForm.BankAccountID" @change="onBankAccountChange" placeholder="请选择对方银行账户">
                <el-option v-for="bankAccount in bankAccountData" :key="bankAccount.ID" :label="bankAccount.BankAccName"
                  :value="bankAccount.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="buyForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">


              <el-upload v-if="!buyForm.FileID" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(buyForm.FileID, buyForm.FileName)">
                下载文件
              </el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="buyForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitBuyForm">保存</el-button>
            <el-button @click="showBuyDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>


    <!-- 查看销售订单信息的对话框 -->

    <el-dialog v-model="showshowBuyDialog" title="采购订单信息" width="80%" @close="resetBuyForm">
      <el-form :model="buyForm" label-width="150px" :rules="buyRules" ref="buyFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="订单编号" prop="OrderNum">
              <el-input v-model="buyForm.OrderNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="订单日期" prop="OrderDate">
              <el-date-picker v-model="buyForm.OrderDate" type="date" placeholder="选择日期"
                :readonly="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="销售方" prop="AcctID">
              <el-select v-model="buyForm.AcctID" @change="onAcctChange" placeholder="请选择销售方" :disabled="true">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="购买方" prop="MerchantID">
              <el-select v-model="buyForm.MerchantID" @change="onMerchantChange" placeholder="请选择购买方" :disabled="true">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="单据要求" prop="DocReq">
              <div v-if="buyForm.DocReq && buyForm.DocReq.length > 0">
                <el-tag v-for="(docReq, index) in buyForm.DocReq" :key="index" type="info" style="margin: 2px;">
                  {{ docReq.DocReq || '未知单据要求' }}
                </el-tag>
              </div>
              <span v-else>无单据要求</span>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="付款方式" prop="PayMentMethodID">
              <el-select v-model="buyForm.PayMentMethodID" @change="onPayMentMethodChange" placeholder="请选择付款方式"
                :disabled="true">
                <el-option v-for="payMentMethod in payMentMethodData" :key="payMentMethod.ID"
                  :label="payMentMethod.PayMtdName" :value="payMentMethod.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账单有效期" prop="BillValidity">
              <el-date-picker v-model="buyForm.BillValidity" type="date" placeholder="选择日期"
                :readonly="true"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单据状态" prop="BussOrderSta">
              <el-select v-model="buyForm.BussOrderSta" placeholder="请选择单据状态" :disabled="true">
                <el-option v-for="bussOrderSta in bussOrderStaData" :key="bussOrderSta.BussOrderSta"
                  :label="bussOrderSta.BussOrderSta" :value="bussOrderSta.BussOrderSta"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="发货开始日期" prop="StartShip">
              <el-date-picker v-model="buyForm.StartShip" type="date" placeholder="选择日期"
                :readonly="true"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="发货截止日期" prop="EndShip">
              <el-date-picker v-model="buyForm.EndShip" type="date" placeholder="选择日期"
                :readonly="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="起运地" prop="SrcPlace">
              <el-select v-model="buyForm.SrcPlace" placeholder="请选择起运地" :disabled="true">
                <el-option v-for="srcPlace in srcPlaceData" :key="srcPlace" :label="srcPlace"
                  :value="srcPlace"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="目的地" prop="Des">
              <el-select v-model="buyForm.Des" placeholder="请选择目的地" :disabled="true">
                <el-option v-for="des in desData" :key="des" :label="des" :value="des"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="9">
            <el-form-item label="质量标准" prop="QualStd">
              <el-select v-model="buyForm.QualStd" placeholder="请选择质量标准" :disabled="true">
                <el-option v-for="qualStd in qualStdData" :key="qualStd.QualStd" :label="qualStd.QualStd"
                  :value="qualStd.QualStd"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="总金额" prop="TotAmt">
              <el-input v-model="buyForm.TotAmt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="7">
            <el-form-item label="币种" prop="CurrencyID">
              <el-select v-model="buyForm.Currency" @change="onCurrencyChange" placeholder="请选择币种" :disabled="true">
                <el-option v-for="currency in currencyData" :key="currency.Currency" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总件数" prop="TotNum">
              <el-input v-model="buyForm.TotNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpecID">
              <el-select v-model="buyForm.PackSpecID" @change="onPackSpecChange" placeholder="请选择包装规格" :disabled="true">
                <el-option v-for="packSpec in packSpecData" :key="packSpec.ID" :label="packSpec.SpecName"
                  :value="packSpec.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总净重" prop="TotalNetWeight">
              <el-input v-model="buyForm.TotalNetWeight" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas">
              <el-select v-model="buyForm.UnitMeas" placeholder="请选择单位" :disabled="true">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeas" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="我方银行账户" prop="AcctBankID">
              <el-select v-model="buyForm.AcctBankID" @change="onAcctBankChange" placeholder="请选择我方银行账户"
                :disabled="true">
                <el-option v-for="acctBank in acctBankData" :key="acctBank.ID" :label="acctBank.AccName"
                  :value="acctBank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="对方银行账户" prop="BankAccountID">
              <el-select v-model="buyForm.BankAccountID" @change="onBankAccountChange" placeholder="请选择对方银行账户"
                :disabled="true">
                <el-option v-for="bankAccount in bankAccountData" :key="bankAccount.ID" :label="bankAccount.BankAccName"
                  :value="bankAccount.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="buyForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-button v-if="buyForm.FileID" type="success" @click="downloadFile(buyForm.FileID, buyForm.FileName)">
                下载文件
              </el-button>
              <span v-else>无文件</span>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="buyForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showshowBuyDialog = false">关闭</el-button>
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
import axios from 'axios'; // 引入 axios
import SideMenu from '@/components/SideMenu.vue'; // 引入 SideMenu 组件
import { useRouter } from 'vue-router';
const router = useRouter();
import { useRoute } from 'vue-router';
const route = useRoute();
// 匹配模式（默认是模糊匹配）
const isExactMatch = ref(true);
const onlyID = ref(true);
const toggleMatchMode = () => {
  console.log("check onlyID", isExactMatch.value)
  isExactMatch.value = !isExactMatch.value;
};

const toggleIDMode = () => {

  console.log("check match", onlyID.value)
  onlyID.value = !onlyID.value;
};

const OutVisible = ref(false)
const OutID = ref(null)
const OutData = ref([])

const fetchOutData = async (ID) => {
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/buy/out', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    OutData.value = response.data.Out; // 假设返回的数据结构中有 Out 字段
    nowID.value = ID
    OutVisible.value = true;
  } catch (error) {
    console.error('获取失败:', error);
    ElMessage.error('获取失败');
  }
};
const DeleteOut = (index, ID, OutID) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该付款单信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    console.log('Confirmed delete', ID); // 添加调试信息

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("OutID", OutID)

    axios.post('/delete/buy/out', params, {
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
          fetchOutData(nowID.value); // 重新获取付款单信息数据
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

    const response = await axios.post('/add/buy/out', params, {
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

    const response = await axios.post('/find/buy/shouldOut', params, {
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

    const response = await axios.post('/add/buy/shouldOut', params, {
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
  ElMessageBox.confirm('确定要删除该应付账款单信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    console.log('Confirmed delete', ID); // 添加调试信息

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("ShouldOutID", id)

    axios.post('/delete/buy/shouldOut', params, {
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
const CheckPrdtInfo = (ID) => {
  try {
    // 确保 ID 是字符串
    const searchQuery = String(ID);

    // 使用路由的 resolve 方法生成完整路径
    const route = router.resolve({
      name: 'Prdt', // 路由名称
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

    axios.post('/delete/buy/purrec', params, {
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

    const response = await axios.post('/add/buy/purrec', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchPurrecData(nowID.value)
    purrecID.value = ''
  } catch (error) {
    console.error('添加采购收货单失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

const fetchPurrecData = async (ID) => {
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/buy/purrec', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    purrecData.value = response.data.Purrec; // 假设返回的数据结构中有 Purrec 字段
    nowID.value = ID
    purrecVisible.value = true;
    console.log("nowid", nowID.value)
  } catch (error) {
    console.error('获取采购收货单失败:', error);
    ElMessage.error('获取采购收货单失败');
  }
};
// 控制主弹窗显示
const prdtInfoVisible = ref(false);
const nowID = ref(null);

const prdtInfoID = ref(null);

const prdtInfoData = ref([]); // 存储产品明细数据

// 删除按钮逻辑
const DeletePrdtInfo = (index, ID, PrdtInfoID) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该产品信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    console.log('Confirmed delete', ID); // 添加调试信息

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("PrdtInfoID", PrdtInfoID)

    axios.post('/delete/buy/prdtInfo', params, {
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
          fetchPrdtInfoData(nowID.value); // 重新获取会计实体信息数据
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
const fetchPrdtInfoData = async (ID) => {
  console.log("id ? ", ID)

  try {
    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/buy/prdtInfo', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    });
    prdtInfoVisible.value = true;
    // prdtInfoData.value = response.data.PrdtInfo; // 假设返回的数据结构中有 PrdtInfo 字段

    prdtInfoData.value = Object.assign([], response.data.PrdtInfo); // 强制更新
    nowID.value = ID
    console.log("nowid", nowID.value)
  } catch (error) {
    console.error('获取产品明细失败:', error);
    ElMessage.error('获取产品明细失败');
  }
};
const addPrdtInfo = async (ID) => {

  console.log(nowID.value)
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("PrdtInfoID", prdtInfoID.value)

    const response = await axios.post('/add/buy/prdtInfo', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchPrdtInfoData(nowID.value)
    prdtInfoID.value = ''
  } catch (error) {
    console.error('添加产品明细失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};
const file = ref(null);
const searchQuery = ref(''); // 添加搜索查询字段
const currentPage = ref(1); // 当前页码
const pageSize = 8; // 每页显示的行数


const paginatedBuyData = computed(() => {
  let filteredData = buyData.value;
  if (searchQuery.value) {

    if (isExactMatch.value === false) {
      if (onlyID.value === false) {
        console.log("sss ", buyData.value)
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value) ||
          item.OrderNum.includes(searchQuery.value) ||
          item.OrderDate.includes(searchQuery.value) ||
          item.AcctName.includes(searchQuery.value) ||
          item.Merc.includes(searchQuery.value) ||
          item.QualStd.includes(searchQuery.value) ||
          item.BillValidity.includes(searchQuery.value) ||
          item.BussOrderSta.includes(searchQuery.value) ||
          item.StartShip.includes(searchQuery.value) ||
          item.EndShip.includes(searchQuery.value) ||
          item.SrcPlace.includes(searchQuery.value) ||
          item.Des.includes(searchQuery.value) ||
          item.PayMtdName.includes(searchQuery.value) ||
          item.TotAmt.toString().includes(searchQuery.value) ||
          item.Currency.includes(searchQuery.value) ||
          item.TotNum.toString().includes(searchQuery.value) ||
          item.SpecName.includes(searchQuery.value) ||
          item.TotalNetWeight.includes(searchQuery.value) ||
          item.UnitMeas.includes(searchQuery.value) ||
          item.AccName.includes(searchQuery.value) ||
          item.BankAccName.includes(searchQuery.value) ||
          item.Notes.includes(searchQuery.value)
        );
      } else {
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value)
        );

      }
    } else {
      if (onlyID === false) {
        filteredData = filteredData.filter(item =>
          item.ID.toString() === searchQuery.value ||
          item.OrderNum === searchQuery.value ||
          item.OrderDate === searchQuery.value ||
          item.AcctName === searchQuery.value ||
          item.Merc === searchQuery.value ||
          item.QualStd === searchQuery.value ||
          item.BillValidity === searchQuery.value ||
          item.BussOrderSta === searchQuery.value ||
          item.StartShip === searchQuery.value ||
          item.EndShip === searchQuery.value ||
          item.SrcPlace === searchQuery.value ||
          item.Des === searchQuery.value ||
          item.PayMtdName === searchQuery.value ||
          item.TotAmt.toString() === searchQuery.value ||
          item.Currency === searchQuery.value ||
          item.TotNum.toString() === searchQuery.value ||
          item.SpecName === searchQuery.value ||
          item.TotalNetWeight === searchQuery.value ||
          item.UnitMeas === searchQuery.value ||
          item.AccName === searchQuery.value ||
          item.BankAccName === searchQuery.value ||
          item.Notes === searchQuery.value
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
onMounted(() => {

  fetchDocReqData(); // 获取单据要求数据
  fetchAcctData(); // 获取会计实体信息
  fetchMerchantData(); // 获取购买方信息
  fetchQualStdData(); // 获取质量标准数据
  fetchBussOrderStaData(); // 获取单据状态数据
  fetchSrcPlaceData(); // 获取起运地数据
  fetchDesData(); // 获取目的地数据
  fetchPayMentMethodData(); // 获取付款方式数据
  fetchPackSpecData(); // 获取包装规格数据
  fetchUnitMeasData(); // 获取单位数据
  fetchAcctBankData(); // 获取我方银行账户数据
  fetchBankAccountData(); // 获取对方银行账户数据
  fetchBuyData(); // 获取采购订单信息
  fetchCurrencyData(); // 新增：获取货币数据

  searchQuery.value = route.query.searchQuery || '';
  // fetchPrdtInfoData(); // 新增：获取产品明细数据
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

const fetchQualStdData = async () => {
  try {
    const response = await axios.get('/find/qualStd'); // 调用质量标准接口
    qualStdData.value = response.data.QualStd; // 假设返回的数据结构中有 QualStd 字段
  } catch (error) {
    console.error('获取质量标准失败:', error);
    ElMessage.error('获取质量标准失败');
  }
};

const fetchBussOrderStaData = async () => {
  try {
    const response = await axios.get('/find/bussOrderSta'); // 调用单据状态接口
    bussOrderStaData.value = response.data.BussOrderSta; // 假设返回的数据结构中有 BussOrderSta 字段
  } catch (error) {
    console.error('获取单据状态失败:', error);
    ElMessage.error('获取单据状态失败');
  }
};

const fetchDesData = async () => {
  try {
    const response = await axios.get('/find/spot'); // 调用目的地接口
    desData.value = response.data.Spot.map(spot => spot.InvLocName); // 提取 Spot 数组中的 InvLocName 字段
  } catch (error) {
    console.error('获取目的地失败:', error);
    ElMessage.error('获取目的地失败');
  }
};


const fetchPayMentMethodData = async () => {
  try {
    const response = await axios.get('/find/payMentMethod'); // 调用付款方式接口
    payMentMethodData.value = response.data.PayMentMethod; // 假设返回的数据结构中有 PayMentMethod 字段
  } catch (error) {
    console.error('获取付款方式失败:', error);
    ElMessage.error('获取付款方式失败');
  }
};

const fetchPackSpecData = async () => {
  try {
    const response = await axios.get('/find/packSpec'); // 调用包装规格接口
    packSpecData.value = response.data.PackSpec; // 假设返回的数据结构中有 PackSpec 字段
  } catch (error) {
    console.error('获取包装规格失败:', error);
    ElMessage.error('获取包装规格失败');
  }
};

const fetchUnitMeasData = async () => {
  try {
    const response = await axios.get('/find/unitMeas'); // 调用单位接口
    unitMeasData.value = response.data.UnitMeas; // 假设返回的数据结构中有 UnitMeas 字段
  } catch (error) {
    console.error('获取单位失败:', error);
    ElMessage.error('获取单位失败');
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

const fetchBuyData = async () => {
  try {
    const response = await axios.get('/find/buy'); // 调用采购订单信息接口
    buyData.value = response.data.Buy; // 假设返回的数据结构中有 Buy 字段
  } catch (error) {
    console.error('获取采购订单信息失败:', error);
    ElMessage.error('获取采购订单信息失败');
  }
};

// 采购订单信息对话框显示状态
const showBuyDialog = ref(false);
const showshowBuyDialog = ref(false);

// 采购订单信息表单数据
const buyForm = ref({
  OrderNum: '',
  OrderDate: '',
  AcctID: '',
  AcctName: '',
  MerchantID: '',
  Merc: '',
  QualStd: '',
  BillValidity: '',
  BussOrderSta: '',
  StartShip: '',
  EndShip: '',
  SrcPlace: '',
  Des: '',
  PayMentMethodID: '',
  PayMtdName: '',
  TotAmt: 0,
  Currency: '',
  TotNum: 0,
  PackSpecID: '',
  SpecName: '',
  TotalNetWeight: '',
  UnitMeas: '',
  AcctBankID: '',
  AccName: '',
  BankAccountID: '',
  BankAccName: '',
  Notes: '',
  FileID: '', // 新增：文件 ID
  FileName: '', // 新增：文件名
  DocReq: [], // 新增：单据要求
  Display: '',
});

const docReqData = ref([]); // 存储单据要求数据

// 获取单据要求数据
const fetchDocReqData = async () => {
  try {
    const response = await axios.get('/find/docReq'); // 调用单据要求接口
    docReqData.value = response.data.DocReq; // 假设返回的数据结构中有 DocReq 字段
  } catch (error) {
    console.error('获取单据要求失败:', error);
    ElMessage.error('获取单据要求失败');
  }
};
// 采购订单信息表单验证规则
const buyRules = {
  OrderNum: [{ required: true, message: '请输入订单编号', trigger: 'blur' }],
  OrderDate: [{ required: true, message: '请选择订单日期', trigger: 'blur' }],
  AcctID: [{ required: true, message: '请选择销售方', trigger: 'blur' }],
  MerchantID: [{ required: true, message: '请选择购买方', trigger: 'blur' }],
  QualStd: [{ required: true, message: '请选择质量标准', trigger: 'blur' }],
  BillValidity: [{ required: true, message: '请选择账单有效期', trigger: 'blur' }],
  BussOrderSta: [{ required: true, message: '请选择单据状态', trigger: 'blur' }],
  StartShip: [{ required: true, message: '请选择发货开始日期', trigger: 'blur' }],
  EndShip: [{ required: true, message: '请选择发货截止日期', trigger: 'blur' }],
  SrcPlace: [{ required: true, message: '请选择起运地', trigger: 'blur' }],
  Des: [{ required: true, message: '请选择目的地', trigger: 'blur' }],
  PayMentMethodID: [{ required: true, message: '请选择付款方式', trigger: 'blur' }],
  TotAmt: [{ required: true, message: '请输入总金额', trigger: 'blur' }],
  Currency: [{ required: true, message: '请输入币种', trigger: 'blur' }],
  TotNum: [{ required: true, message: '请输入总件数', trigger: 'blur' }],
  PackSpecID: [{ required: true, message: '请选择包装规格', trigger: 'blur' }],
  TotalNetWeight: [{ required: true, message: '请输入总净重', trigger: 'blur' }],
  UnitMeas: [{ required: true, message: '请选择单位', trigger: 'blur' }],
  AcctBankID: [{ required: true, message: '请选择我方银行账户', trigger: 'blur' }],
  BankAccountID: [{ required: true, message: '请选择对方银行账户', trigger: 'blur' }],
};

// 表格数据（初始值为空数组）
const acctData = ref([]); // 会计实体信息
const merchantData = ref([]); // 购买方信息
const qualStdData = ref([]); // 质量标准数据
const bussOrderStaData = ref([]); // 单据状态数据
const srcPlaceData = ref([]); // 起运地数据
const desData = ref([]); // 目的地数据
const payMentMethodData = ref([]); // 付款方式数据
const packSpecData = ref([]); // 包装规格数据
const unitMeasData = ref([]); // 单位数据
const acctBankData = ref([]); // 我方银行账户数据
const bankAccountData = ref([]); // 对方银行账户数据
const buyData = ref([]); // 采购订单信息

// 根据当前选中的菜单项动态更改标题和按钮文本
const headerTitle = computed(() => {
  return '采购订单信息';
});

const addButtonText = computed(() => {
  return '添加采购订单信息';
});

// 添加按钮点击事件
const handleAdd = () => {
  showBuyDialog.value = true;
};

const handleEdit = (index, row) => {
  // 将当前行的数据赋值给 buyForm
  buyForm.value = { ...row };

  // 如果 PrdtInfos 是对象数组，转换为 ID 数组
  if (row.PrdtInfos && Array.isArray(row.PrdtInfos)) {
    buyForm.value.PrdtInfos = row.PrdtInfos.map(item => item.ID);
  }

  if (row.DocReq && Array.isArray(row.DocReq)) {
    buyForm.value.DocReq = row.DocReq.map(item => item.DocReqID);
  }
  // 检查是否已上传文件
  if (row.FileID) {
    buyForm.value.FileID = row.FileID; // 保存 FileID
    buyForm.value.FileName = row.FileName; // 保存文件名
  }

  showBuyDialog.value = true; // 打开采购订单信息对话框
};

// 查看按钮逻辑
const handleView = (index, row) => {
  buyForm.value = { ...row }; // 将当前行的数据赋值给 buyForm
  showshowBuyDialog.value = true; // 打开查看采购订单信息对话框
};

// 删除按钮逻辑
const handleDelete = (index, ID) => {
  ElMessageBox.confirm('确定要删除该采购订单信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/buy', {
      "ID": ID,
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchBuyData(); // 重新获取采购订单信息数据
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
// 重置采购订单信息表单
const resetBuyForm = () => {
  buyForm.value = {
    OrderNum: '',
    OrderDate: '',
    AcctID: '',
    AcctName: '',
    MerchantID: '',
    Merc: '',
    QualStd: '',
    BillValidity: '',
    BussOrderSta: '',
    StartShip: '',
    EndShip: '',
    SrcPlace: '',
    Des: '',
    PayMentMethodID: '',
    PayMtdName: '',
    TotAmt: 0,
    Currency: '',
    TotNum: 0,
    PackSpecID: '',
    SpecName: '',
    TotalNetWeight: '',
    UnitMeas: '',
    AcctBankID: '',
    AccName: '',
    BankAccountID: '',
    BankAccName: '',
    Notes: '',
    FileID: '', // 重置文件 ID
    FileName: '', // 重置文件名
    // PrdtInfos: [], // 重置产品明细
  };
  file.value = null; // 重置文件对象
  if (uploadRef.value) {
    uploadRef.value.clearFiles(); // 清空文件列表
  }
};

const uploadRef = ref(null);

const submitBuyForm = async () => {
  try {
    const formData = new FormData();

    // 添加其他字段
    Object.keys(buyForm.value).forEach((key) => {
      if (key != `DocReq`) {
        formData.append(key, buyForm.value[key]);
      }
    });

    console.log("aa", buyForm.value.DocReq)
    // 处理 DocReq
    buyForm.value.DocReq.forEach((docReqID, index) => {
      formData.append(`DocReq[${index}][DocReqID]`, docReqID);
      console.log("hh", docReqID)
    });
    // 添加文件
    if (file.value) {
      formData.append('file', file.value);
    }

    // 提交表单
    const response = await axios.post('/save/buy', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    if (response.status === 200) {
      ElMessage.success('采购订单信息保存成功');
      showBuyDialog.value = false; // 关闭对话框
      fetchBuyData(); // 重新获取采购订单信息数据
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
    buyForm.value.AcctName = selectedAcct.AcctName;
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
  buyForm.value.FileName = uploadFile.name; // 更新文件名
};
const fetchSrcPlaceData = async () => {
  try {
    const response = await axios.get('/find/spot'); // 调用起运地接口
    srcPlaceData.value = response.data.Spot.map(spot => spot.InvLocName); // 提取 Spot 数组中的 InvLocName 字段
  } catch (error) {
    console.error('获取起运地失败:', error);
    ElMessage.error('获取起运地失败');
  }
};

// 监听 change 事件并更新 Merc
const onMerchantChange = (value) => {
  const selectedMerchant = merchantData.value.find(merchant => merchant.ID === value);
  if (selectedMerchant) {
    buyForm.value.Merc = selectedMerchant.Merc;
  }
};

// 监听 change 事件并更新 PayMtdName
// const onPayMentMethodChange = (value) => {
//   const selectedPayMentMethod = payMentMethodData.value.find(payMentMethod => payMentMethod.ID === value);
//   if (selectedPayMentMethod) {
//     buyForm.value.PayMtdName = selectedPayMentMethod.PayMtdName;
//   }
// };

// 监听 change 事件并更新 SpecName
const onPackSpecChange = (value) => {
  const selectedPackSpec = packSpecData.value.find(packSpec => packSpec.ID === value);
  if (selectedPackSpec) {
    buyForm.value.SpecName = selectedPackSpec.SpecName;
  }
};

// 监听 change 事件并更新 AccName
const onAcctBankChange = (value) => {
  const selectedAcctBank = acctBankData.value.find(acctBank => acctBank.ID === value);
  if (selectedAcctBank) {
    buyForm.value.AccName = selectedAcctBank.AccName;
  }
};

// 监听 change 事件并更新 BankAccName
const onBankAccountChange = (value) => {
  const selectedBankAccount = bankAccountData.value.find(bankAccount => bankAccount.ID === value);
  if (selectedBankAccount) {
    buyForm.value.BankAccName = selectedBankAccount.BankAccName;
  }
};

// 分页逻辑
const handlePageChange = (page) => {
  currentPage.value = page;
};
</script>

<style src="../assets/styles/Bottom.css"></style>

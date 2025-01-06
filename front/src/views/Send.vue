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

        <el-header style="display: flex; justify-content: space-between; align-items: center;">
          <h2>{{ headerTitle }}</h2>
          <div>
            搜索：
            <el-input v-model="searchQuery" placeholder="输入要搜索的关键字" style="width: 200px;" />
            <el-button type="primary" @click="handleAdd">{{ addButtonText }}</el-button>
          </div>
        </el-header>
        <el-main>

          <!-- 销售订单信息表格 -->
          <el-table :data="paginatedSaleData" style="width: 100%" max-height="450">
            <el-table-column prop="OrderNum" label="订单编号" width="220%"></el-table-column>
            <el-table-column prop="OrderDate" label="订单日期" width="220%"></el-table-column>
            <el-table-column prop="AcctName" label="销售方" width="220%"></el-table-column>
            <el-table-column prop="Merc" label="购买方" width="220%"></el-table-column>
            <el-table-column prop="QualStd" label="质量标准" width="220%"></el-table-column>
            <el-table-column prop="BillValidity" label="账单有效期" width="220%"></el-table-column>
            <el-table-column prop="BussOrderSta" label="单据状态" width="220%"></el-table-column>


            <el-table-column label="单据要求" width="200">
              <template #default="scope">
                <div v-if="scope.row.DocReq && scope.row.DocReq.length > 0">
                  <el-tag v-for="(docReq, index) in scope.row.DocReq" :key="index" type="info" style="margin: 2px;">
                    {{ docReq.DocReq || '未知单据要求' }}
                  </el-tag>
                </div>
                <span v-else>无单据要求</span>
              </template>
            </el-table-column>
            <el-table-column prop="StartShip" label="发货开始日期" width="220%"></el-table-column>
            <el-table-column prop="EndShip" label="发货截止日期" width="220%"></el-table-column>
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

            <el-table-column label="产品明细" width="300">
              <template #default="scope">
                <div v-for="prdtInfo in scope.row.PrdtInfos" :key="prdtInfo.ID">
                  {{ prdtInfo.ID }} - {{ prdtInfo.CatEngName }} - {{ prdtInfo.BrandEngName }} - {{ prdtInfo.PackSpec }}
                  -
                  {{
                    prdtInfo.Currency }} - {{ prdtInfo.UnitPrice }} - {{ prdtInfo.TradeTerm }} - {{ prdtInfo.DeliveryLoc
                  }}
                </div>
              </template>
            </el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.ID)" type="text" size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="saleData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>


    <!-- 添加销售订单信息的对话框 -->
    <el-dialog v-model="showSaleDialog" title="销售订单信息" width="80%" @close="resetSaleForm">
      <el-form :model="saleForm" label-width="150px" :rules="saleRules" ref="saleFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="订单编号" prop="OrderNum">
              <el-input v-model="saleForm.OrderNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="订单日期" prop="OrderDate">
              <el-date-picker v-model="saleForm.OrderDate" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="销售方" prop="AcctId">
              <el-select v-model="saleForm.AcctId" @change="onAcctChange" placeholder="请选择销售方">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="购买方" prop="MerchantId">
              <el-select v-model="saleForm.MerchantId" @change="onMerchantChange" placeholder="请选择购买方">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">

            <el-form-item label="单据要求" prop="DocReq">
              <el-select v-model="saleForm.DocReq" multiple placeholder="请选择单据要求" @click="onDocReqChange">
                <el-option v-for="docReq in docReqData" :key="docReq.DocReqId" :label="docReq.DocReq"
                  :value="docReq.DocReqId">
                  {{ docReq.DocReq }}
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="付款方式" prop="PayMentMethodId">
              <el-select v-model="saleForm.PayMentMethodId" @change="onPayMentMethodChange" placeholder="请选择付款方式">
                <el-option v-for="payMentMethod in payMentMethodData" :key="payMentMethod.ID"
                  :label="payMentMethod.PayMtdName" :value="payMentMethod.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账单有效期" prop="BillValidity">
              <el-date-picker v-model="saleForm.BillValidity" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单据状态" prop="BussOrderSta">
              <el-select v-model="saleForm.BussOrderSta" placeholder="请选择单据状态">
                <el-option v-for="bussOrderSta in bussOrderStaData" :key="bussOrderSta.BussOrderSta"
                  :label="bussOrderSta.BussOrderSta" :value="bussOrderSta.BussOrderSta"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="发货开始日期" prop="StartShip">
              <el-date-picker v-model="saleForm.StartShip" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="发货截止日期" prop="EndShip">
              <el-date-picker v-model="saleForm.EndShip" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="起运地" prop="SrcPlace">
              <el-select v-model="saleForm.SrcPlace" placeholder="请选择起运地">
                <el-option v-for="srcPlace in srcPlaceData" :key="srcPlace" :label="srcPlace"
                  :value="srcPlace"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="目的地" prop="Des">
              <el-select v-model="saleForm.Des" placeholder="请选择目的地">
                <el-option v-for="des in desData" :key="des" :label="des" :value="des"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="9">
            <el-form-item label="质量标准" prop="QualStd">
              <el-select v-model="saleForm.QualStd" placeholder="请选择质量标准">
                <el-option v-for="qualStd in qualStdData" :key="qualStd.QualStd" :label="qualStd.QualStd"
                  :value="qualStd.QualStd"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="总金额" prop="TotAmt">
              <el-input v-model="saleForm.TotAmt"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="7">
            <el-form-item label="币种" prop="CurrencyId">
              <el-select v-model="saleForm.Currency" @change="onCurrencyChange" placeholder="请选择币种">
                <el-option v-for="currency in currencyData" :key="currency.Currency" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总件数" prop="TotNum">
              <el-input v-model="saleForm.TotNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpecId">
              <el-select v-model="saleForm.PackSpecId" @change="onPackSpecChange" placeholder="请选择包装规格">
                <el-option v-for="packSpec in packSpecData" :key="packSpec.ID" :label="packSpec.SpecName"
                  :value="packSpec.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总净重" prop="TotalNetWeight">
              <el-input v-model="saleForm.TotalNetWeight"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas">
              <el-select v-model="saleForm.UnitMeas" placeholder="请选择单位">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeas" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="我方银行账户" prop="AcctBankId">
              <el-select v-model="saleForm.AcctBankId" @change="onAcctBankChange" placeholder="请选择我方银行账户">
                <el-option v-for="acctBank in acctBankData" :key="acctBank.ID" :label="acctBank.AccName"
                  :value="acctBank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="对方银行账户" prop="BankAccountId">
              <el-select v-model="saleForm.BankAccountId" @change="onBankAccountChange" placeholder="请选择对方银行账户">
                <el-option v-for="bankAccount in bankAccountData" :key="bankAccount.ID" :label="bankAccount.BankAccName"
                  :value="bankAccount.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="saleForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">


              <el-upload v-if="!saleForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(saleForm.FileId, saleForm.FileName)">
                下载文件
              </el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="saleForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>



        <el-select v-model="saleForm.PrdtInfos" multiple placeholder="请选择产品明细" @click="onPrdtInfoChange">
          <el-option v-for="prdtInfo in prdtInfoData" :key="prdtInfo.ID"
            :label="`${prdtInfo.ID} - ${prdtInfo.CatEngName} - ${prdtInfo.BrandEngName} - ${prdtInfo.PackSpec} - ${prdtInfo.Currency} - ${prdtInfo.UnitPrice} - ${prdtInfo.TradeTerm} - ${prdtInfo.DeliveryLoc}`"
            :value="prdtInfo.ID">
            <span>{{ prdtInfo.ID }}</span> -
            <span>{{ prdtInfo.CatEngName }}</span> -
            <span>{{ prdtInfo.BrandEngName }}</span> -
            <span>{{ prdtInfo.PackSpec }}</span> -
            <span>{{ prdtInfo.Currency }}</span> -
            <span>{{ prdtInfo.UnitPrice }}</span> -
            <span>{{ prdtInfo.TradeTerm }}</span> -
            <span>{{ prdtInfo.DeliveryLoc }}</span>
          </el-option>
        </el-select>
        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitSaleForm">保存</el-button>
            <el-button @click="showSaleDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>


    <!-- 查看销售订单信息的对话框 -->

    <el-dialog v-model="showshowSaleDialog" title="销售订单信息" width="80%" @close="resetSaleForm">
      <el-form :model="saleForm" label-width="150px" :rules="saleRules" ref="saleFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="订单编号" prop="OrderNum">
              <el-input v-model="saleForm.OrderNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="订单日期" prop="OrderDate">
              <el-date-picker v-model="saleForm.OrderDate" type="date" placeholder="选择日期"
                :readonly="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="销售方" prop="AcctId">
              <el-select v-model="saleForm.AcctId" @change="onAcctChange" placeholder="请选择销售方" :disabled="true">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="购买方" prop="MerchantId">
              <el-select v-model="saleForm.MerchantId" @change="onMerchantChange" placeholder="请选择购买方" :disabled="true">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="单据要求" prop="DocReq">
              <div v-if="saleForm.DocReq && saleForm.DocReq.length > 0">
                <el-tag v-for="(docReq, index) in saleForm.DocReq" :key="index" type="info" style="margin: 2px;">
                  {{ docReq.DocReq || '未知单据要求' }}
                </el-tag>
              </div>
              <span v-else>无单据要求</span>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="付款方式" prop="PayMentMethodId">
              <el-select v-model="saleForm.PayMentMethodId" @change="onPayMentMethodChange" placeholder="请选择付款方式"
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
              <el-date-picker v-model="saleForm.BillValidity" type="date" placeholder="选择日期"
                :readonly="true"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单据状态" prop="BussOrderSta">
              <el-select v-model="saleForm.BussOrderSta" placeholder="请选择单据状态" :disabled="true">
                <el-option v-for="bussOrderSta in bussOrderStaData" :key="bussOrderSta.BussOrderSta"
                  :label="bussOrderSta.BussOrderSta" :value="bussOrderSta.BussOrderSta"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="发货开始日期" prop="StartShip">
              <el-date-picker v-model="saleForm.StartShip" type="date" placeholder="选择日期"
                :readonly="true"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="发货截止日期" prop="EndShip">
              <el-date-picker v-model="saleForm.EndShip" type="date" placeholder="选择日期"
                :readonly="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="起运地" prop="SrcPlace">
              <el-select v-model="saleForm.SrcPlace" placeholder="请选择起运地" :disabled="true">
                <el-option v-for="srcPlace in srcPlaceData" :key="srcPlace" :label="srcPlace"
                  :value="srcPlace"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="目的地" prop="Des">
              <el-select v-model="saleForm.Des" placeholder="请选择目的地" :disabled="true">
                <el-option v-for="des in desData" :key="des" :label="des" :value="des"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="9">
            <el-form-item label="质量标准" prop="QualStd">
              <el-select v-model="saleForm.QualStd" placeholder="请选择质量标准" :disabled="true">
                <el-option v-for="qualStd in qualStdData" :key="qualStd.QualStd" :label="qualStd.QualStd"
                  :value="qualStd.QualStd"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="总金额" prop="TotAmt">
              <el-input v-model="saleForm.TotAmt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="7">
            <el-form-item label="币种" prop="CurrencyId">
              <el-select v-model="saleForm.Currency" @change="onCurrencyChange" placeholder="请选择币种" :disabled="true">
                <el-option v-for="currency in currencyData" :key="currency.Currency" :label="currency.Currency"
                  :value="currency.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总件数" prop="TotNum">
              <el-input v-model="saleForm.TotNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpecId">
              <el-select v-model="saleForm.PackSpecId" @change="onPackSpecChange" placeholder="请选择包装规格"
                :disabled="true">
                <el-option v-for="packSpec in packSpecData" :key="packSpec.ID" :label="packSpec.SpecName"
                  :value="packSpec.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总净重" prop="TotalNetWeight">
              <el-input v-model="saleForm.TotalNetWeight" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas">
              <el-select v-model="saleForm.UnitMeas" placeholder="请选择单位" :disabled="true">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeas" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="我方银行账户" prop="AcctBankId">
              <el-select v-model="saleForm.AcctBankId" @change="onAcctBankChange" placeholder="请选择我方银行账户"
                :disabled="true">
                <el-option v-for="acctBank in acctBankData" :key="acctBank.ID" :label="acctBank.AccName"
                  :value="acctBank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="对方银行账户" prop="BankAccountId">
              <el-select v-model="saleForm.BankAccountId" @change="onBankAccountChange" placeholder="请选择对方银行账户"
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
              <el-input v-model="saleForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-button v-if="saleForm.FileId" type="success"
                @click="downloadFile(saleForm.FileId, saleForm.FileName)">
                下载文件
              </el-button>
              <span v-else>无文件</span>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="saleForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="产品明细">
              <el-select v-model="saleForm.PrdtInfos" multiple placeholder="请选择产品明细" @click="onPrdtInfoChange"
                :disabled="true">
                <el-option v-for="prdtInfo in prdtInfoData" :key="prdtInfo.ID"
                  :label="`${prdtInfo.ID} - ${prdtInfo.CatEngName} - ${prdtInfo.BrandEngName} - ${prdtInfo.PackSpec} - ${prdtInfo.Currency} - ${prdtInfo.UnitPrice} - ${prdtInfo.TradeTerm} - ${prdtInfo.DeliveryLoc}`"
                  :value="prdtInfo.ID">
                  <span>{{ prdtInfo.ID }}</span> -
                  <span>{{ prdtInfo.CatEngName }}</span> -
                  <span>{{ prdtInfo.BrandEngName }}</span> -
                  <span>{{ prdtInfo.PackSpec }}</span> -
                  <span>{{ prdtInfo.Currency }}</span> -
                  <span>{{ prdtInfo.UnitPrice }}</span> -
                  <span>{{ prdtInfo.TradeTerm }}</span> -
                  <span>{{ prdtInfo.DeliveryLoc }}</span>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showshowSaleDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import axios from 'axios'; // 引入 axios
import SideMenu from '@/components/SideMenu.vue'; // 引入 SideMenu 组件

const file = ref(null);
const searchQuery = ref(''); // 添加搜索查询字段
const currentPage = ref(1); // 当前页码
const pageSize = 8; // 每页显示的行数

const onPrdtInfoChange = (selectedPrdtInfos) => {
  // 直接更新 saleForm.PrdtInfos，不需要额外处理
  saleForm.value.PrdtInfos = selectedPrdtInfos;
};

const onDocReqChange = (selectedDocReq) => {
  saleForm.value.DocReq = selectedDocReq;
};
const paginatedSaleData = computed(() => {
  let filteredData = saleData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
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
  fetchSaleData(); // 获取销售订单信息
  fetchCurrencyData(); // 新增：获取货币数据
  fetchPrdtInfoData(); // 新增：获取产品明细数据
});

const prdtInfoData = ref([]); // 存储产品明细数据

const fetchPrdtInfoData = async () => {
  try {
    const response = await axios.get('/find/prdtInfo'); // 调用产品明细接口
    prdtInfoData.value = response.data.PrdtInfo; // 假设返回的数据结构中有 PrdtInfo 字段
  } catch (error) {
    console.error('获取产品明细失败:', error);
    ElMessage.error('获取产品明细失败，请稍后重试');
  }
};
const currencyData = ref([]); // 存储货币数据

const fetchCurrencyData = async () => {
  try {
    const response = await axios.get('/find/currency'); // 调用货币数据接口
    currencyData.value = response.data.Currency; // 假设返回的数据结构中有 Currency 字段
  } catch (error) {
    console.error('获取货币数据失败:', error);
    ElMessage.error('获取货币数据失败，请稍后重试');
  }
};
// 定义接口请求函数
const fetchAcctData = async () => {
  try {
    const response = await axios.get('/find/acct'); // 调用会计实体信息接口
    acctData.value = response.data.Acct; // 假设返回的数据结构中有 Acct 字段
  } catch (error) {
    console.error('获取会计实体信息失败:', error);
    ElMessage.error('获取会计实体信息失败，请稍后重试');
  }
};

const fetchMerchantData = async () => {
  try {
    const response = await axios.get('/find/merchant'); // 调用购买方信息接口
    merchantData.value = response.data.Merchant; // 假设返回的数据结构中有 Merchant 字段
  } catch (error) {
    console.error('获取购买方信息失败:', error);
    ElMessage.error('获取购买方信息失败，请稍后重试');
  }
};

const fetchQualStdData = async () => {
  try {
    const response = await axios.get('/find/qualStd'); // 调用质量标准接口
    qualStdData.value = response.data.QualStd; // 假设返回的数据结构中有 QualStd 字段
  } catch (error) {
    console.error('获取质量标准失败:', error);
    ElMessage.error('获取质量标准失败，请稍后重试');
  }
};

const fetchBussOrderStaData = async () => {
  try {
    const response = await axios.get('/find/bussOrderSta'); // 调用单据状态接口
    bussOrderStaData.value = response.data.BussOrderSta; // 假设返回的数据结构中有 BussOrderSta 字段
  } catch (error) {
    console.error('获取单据状态失败:', error);
    ElMessage.error('获取单据状态失败，请稍后重试');
  }
};

const fetchDesData = async () => {
  try {
    const response = await axios.get('/find/spot'); // 调用目的地接口
    desData.value = response.data.Spot.map(spot => spot.InvLocName); // 提取 Spot 数组中的 InvLocName 字段
  } catch (error) {
    console.error('获取目的地失败:', error);
    ElMessage.error('获取目的地失败，请稍后重试');
  }
};


const fetchPayMentMethodData = async () => {
  try {
    const response = await axios.get('/find/payMentMethod'); // 调用付款方式接口
    payMentMethodData.value = response.data.PayMentMethod; // 假设返回的数据结构中有 PayMentMethod 字段
  } catch (error) {
    console.error('获取付款方式失败:', error);
    ElMessage.error('获取付款方式失败，请稍后重试');
  }
};

const fetchPackSpecData = async () => {
  try {
    const response = await axios.get('/find/packSpec'); // 调用包装规格接口
    packSpecData.value = response.data.PackSpec; // 假设返回的数据结构中有 PackSpec 字段
  } catch (error) {
    console.error('获取包装规格失败:', error);
    ElMessage.error('获取包装规格失败，请稍后重试');
  }
};

const fetchUnitMeasData = async () => {
  try {
    const response = await axios.get('/find/unitMeas'); // 调用单位接口
    unitMeasData.value = response.data.UnitMeas; // 假设返回的数据结构中有 UnitMeas 字段
  } catch (error) {
    console.error('获取单位失败:', error);
    ElMessage.error('获取单位失败，请稍后重试');
  }
};

const fetchAcctBankData = async () => {
  try {
    const response = await axios.get('/find/acctBank'); // 调用我方银行账户接口
    acctBankData.value = response.data.AcctBank; // 假设返回的数据结构中有 AcctBank 字段
  } catch (error) {
    console.error('获取我方银行账户失败:', error);
    ElMessage.error('获取我方银行账户失败，请稍后重试');
  }
};

const fetchBankAccountData = async () => {
  try {
    const response = await axios.get('/find/bankAccount'); // 调用对方银行账户接口
    bankAccountData.value = response.data.BankAccount; // 假设返回的数据结构中有 BankAccount 字段
  } catch (error) {
    console.error('获取对方银行账户失败:', error);
    ElMessage.error('获取对方银行账户失败，请稍后重试');
  }
};

const fetchSaleData = async () => {
  try {
    const response = await axios.get('/find/sale'); // 调用销售订单信息接口
    saleData.value = response.data.Sale; // 假设返回的数据结构中有 Sale 字段
  } catch (error) {
    console.error('获取销售订单信息失败:', error);
    ElMessage.error('获取销售订单信息失败，请稍后重试');
  }
};

// 销售订单信息对话框显示状态
const showSaleDialog = ref(false);
const showshowSaleDialog = ref(false);

// 销售订单信息表单数据
const saleForm = ref({
  OrderNum: '',
  OrderDate: '',
  AcctId: '',
  AcctName: '',
  MerchantId: '',
  Merc: '',
  QualStd: '',
  BillValidity: '',
  BussOrderSta: '',
  StartShip: '',
  EndShip: '',
  SrcPlace: '',
  Des: '',
  PayMentMethodId: '',
  PayMtdName: '',
  TotAmt: 0,
  Currency: '',
  TotNum: 0,
  PackSpecId: '',
  SpecName: '',
  TotalNetWeight: '',
  UnitMeas: '',
  AcctBankId: '',
  AccName: '',
  BankAccountId: '',
  BankAccName: '',
  Notes: '',
  FileId: '', // 新增：文件 ID
  FileName: '', // 新增：文件名
  PrdtInfos: [], // 新增：产品明细
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
    ElMessage.error('获取单据要求失败，请稍后重试');
  }
};
// 销售订单信息表单验证规则
const saleRules = {
  OrderNum: [{ required: true, message: '请输入订单编号', trigger: 'blur' }],
  OrderDate: [{ required: true, message: '请选择订单日期', trigger: 'blur' }],
  AcctId: [{ required: true, message: '请选择销售方', trigger: 'blur' }],
  MerchantId: [{ required: true, message: '请选择购买方', trigger: 'blur' }],
  QualStd: [{ required: true, message: '请选择质量标准', trigger: 'blur' }],
  BillValidity: [{ required: true, message: '请选择账单有效期', trigger: 'blur' }],
  BussOrderSta: [{ required: true, message: '请选择单据状态', trigger: 'blur' }],
  StartShip: [{ required: true, message: '请选择发货开始日期', trigger: 'blur' }],
  EndShip: [{ required: true, message: '请选择发货截止日期', trigger: 'blur' }],
  SrcPlace: [{ required: true, message: '请选择起运地', trigger: 'blur' }],
  Des: [{ required: true, message: '请选择目的地', trigger: 'blur' }],
  PayMentMethodId: [{ required: true, message: '请选择付款方式', trigger: 'blur' }],
  TotAmt: [{ required: true, message: '请输入总金额', trigger: 'blur' }],
  Currency: [{ required: true, message: '请输入币种', trigger: 'blur' }],
  TotNum: [{ required: true, message: '请输入总件数', trigger: 'blur' }],
  PackSpecId: [{ required: true, message: '请选择包装规格', trigger: 'blur' }],
  TotalNetWeight: [{ required: true, message: '请输入总净重', trigger: 'blur' }],
  UnitMeas: [{ required: true, message: '请选择单位', trigger: 'blur' }],
  AcctBankId: [{ required: true, message: '请选择我方银行账户', trigger: 'blur' }],
  BankAccountId: [{ required: true, message: '请选择对方银行账户', trigger: 'blur' }],
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
const saleData = ref([]); // 销售订单信息

// 根据当前选中的菜单项动态更改标题和按钮文本
const headerTitle = computed(() => {
  return '销售订单信息';
});

const addButtonText = computed(() => {
  return '添加销售订单信息';
});

// 添加按钮点击事件
const handleAdd = () => {
  showSaleDialog.value = true;
};


const handleEdit = (index, row) => {
  saleForm.value = { ...row }; // 将当前行的数据赋值给 saleForm
  showSaleDialog.value = true; // 打开销售订单信息对话框

  // 检查是否已上传文件
  if (row.FileId) {
    saleForm.value.FileId = row.FileId; // 保存 FileId
    saleForm.value.FileName = row.FileName; // 保存文件名
  }
};
// 查看按钮逻辑
const handleView = (index, row) => {
  saleForm.value = { ...row }; // 将当前行的数据赋值给 saleForm
  showshowSaleDialog.value = true; // 打开查看销售订单信息对话框
};

// 删除按钮逻辑
const handleDelete = (index, ID) => {
  ElMessageBox.confirm('确定要删除该销售订单信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/sale', {
      "ID": ID,
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchSaleData(); // 重新获取销售订单信息数据
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
const resetSaleForm = () => {
  saleForm.value = {
    OrderNum: '',
    OrderDate: '',
    AcctId: '',
    AcctName: '',
    MerchantId: '',
    Merc: '',
    QualStd: '',
    BillValidity: '',
    BussOrderSta: '',
    StartShip: '',
    EndShip: '',
    SrcPlace: '',
    Des: '',
    PayMentMethodId: '',
    PayMtdName: '',
    TotAmt: 0,
    Currency: '',
    TotNum: 0,
    PackSpecId: '',
    SpecName: '',
    TotalNetWeight: '',
    UnitMeas: '',
    AcctBankId: '',
    AccName: '',
    BankAccountId: '',
    BankAccName: '',
    Notes: '',
    FileId: '', // 重置文件 ID
    FileName: '', // 重置文件名
    PrdtInfos: [], // 重置产品明细
  };
  file.value = null; // 重置文件对象
  if (uploadRef.value) {
    uploadRef.value.clearFiles(); // 清空文件列表
  }
};

const uploadRef = ref(null);

const submitSaleForm = async () => {
  try {
    const formData = new FormData();

    // 添加其他字段
    Object.keys(saleForm.value).forEach((key) => {
      if (key !== 'PrdtInfos' && key != `DocReq`) {
        formData.append(key, saleForm.value[key]);
      }
    });

    // 处理 PrdtInfos
    saleForm.value.PrdtInfos.forEach((prdtInfo, index) => {
      formData.append(`PrdtInfos[${index}][ID]`, prdtInfo);
      console.log("sss", prdtInfo)
      // 如果需要提交其他字段，可以继续添加
      // formData.append(`PrdtInfos[${index}][CatEngName]`, prdtInfo.CatEngName);
      // formData.append(`PrdtInfos[${index}][BrandEngName]`, prdtInfo.BrandEngName);
      // ...
    });

    console.log("aa1", saleForm.value.PrdtInfos)
    console.log("aa", saleForm.value.DocReq)
    // 处理 DocReq
    saleForm.value.DocReq.forEach((docReqId, index) => {
      formData.append(`DocReq[${index}][DocReqId]`, docReqId);
      console.log("hh", docReqId)
    });
    // 添加文件
    if (file.value) {
      formData.append('file', file.value);
    }

    // 提交表单
    const response = await axios.post('/save/sale', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    if (response.status === 200) {
      ElMessage.success('销售订单信息保存成功');
      showSaleDialog.value = false; // 关闭对话框
      fetchSaleData(); // 重新获取销售订单信息数据
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
    ElMessage.error('文件下载失败，请稍后重试');
  }
};
const handleFileChange = (uploadFile) => {
  file.value = uploadFile.raw; // 保存选择的文件
  saleForm.value.FileName = uploadFile.name; // 更新文件名
};
const fetchSrcPlaceData = async () => {
  try {
    const response = await axios.get('/find/spot'); // 调用起运地接口
    srcPlaceData.value = response.data.Spot.map(spot => spot.InvLocName); // 提取 Spot 数组中的 InvLocName 字段
  } catch (error) {
    console.error('获取起运地失败:', error);
    ElMessage.error('获取起运地失败，请稍后重试');
  }
};

// 监听 change 事件并更新 Merc
const onMerchantChange = (value) => {
  const selectedMerchant = merchantData.value.find(merchant => merchant.ID === value);
  if (selectedMerchant) {
    saleForm.value.Merc = selectedMerchant.Merc;
  }
};

// 监听 change 事件并更新 PayMtdName
const onPayMentMethodChange = (value) => {
  const selectedPayMentMethod = payMentMethodData.value.find(payMentMethod => payMentMethod.ID === value);
  if (selectedPayMentMethod) {
    saleForm.value.PayMtdName = selectedPayMentMethod.PayMtdName;
  }
};

// 监听 change 事件并更新 SpecName
const onPackSpecChange = (value) => {
  const selectedPackSpec = packSpecData.value.find(packSpec => packSpec.ID === value);
  if (selectedPackSpec) {
    saleForm.value.SpecName = selectedPackSpec.SpecName;
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

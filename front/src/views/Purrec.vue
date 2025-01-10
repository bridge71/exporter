<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-18'" />
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
          <el-table :data="paginatedPurrecData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="SaleInvNum" label="采购发票号" width="220%"></el-table-column>
            <el-table-column prop="Acct1Name" label="客户" width="220%"></el-table-column>
            <el-table-column prop="Acct2Name" label="收货人" width="220%"></el-table-column>
            <el-table-column prop="Acct3Name" label="通知人" width="220%"></el-table-column>

            <el-table-column label="发货人" width="220%">
              <template #default="scope">
                <span v-if="scope.row.Merchant.Merc">{{ scope.row.Merchant.Merc }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column label="付款银行" width="220%">
              <template #default="scope">
                <span v-if="scope.row.AcctBank.AccName">{{ scope.row.AcctBank.AccName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column label="付款方式" width="220%">
              <template #default="scope">
                <span v-if="scope.row.PayMentMethod.PayMtdName">{{ scope.row.PayMentMethod.PayMtdName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column label="包装规格" width="220%">
              <template #default="scope">
                <span v-if="scope.row.PackSpec.SpecName">{{ scope.row.PackSpec.SpecName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column prop="SrcPlace" label="起运地" width="220%"></el-table-column>
            <el-table-column prop="Des" label="目的地" width="220%"></el-table-column>
            <el-table-column prop="SaleInvDate" label="采购发票日期" width="420%"></el-table-column>
            <el-table-column prop="ShipName" label="船名" width="220%"></el-table-column>
            <el-table-column prop="Voyage" label="航次" width="220%"></el-table-column>
            <el-table-column prop="TotNum" label="总件数" width="220%"></el-table-column>
            <el-table-column prop="TotalNetWeight" label="总净重" width="220%"></el-table-column>
            <el-table-column prop="UnitMeas1" label="单位" width="220%"></el-table-column>
            <el-table-column prop="GrossWt" label="总毛重" width="220%"></el-table-column>
            <el-table-column prop="UnitMeas2" label="单位" width="220%"></el-table-column>
            <el-table-column prop="TotVol" label="总体积" width="220%"></el-table-column>
            <el-table-column prop="UnitMeas3" label="单位" width="220%"></el-table-column>
            <el-table-column prop="BillLadNum" label="提单号" width="220%"></el-table-column>
            <el-table-column prop="DateOfShip" label="发货(开船)日期" width="420%"></el-table-column>
            <el-table-column prop="Note1" label="提单货物描述" width="220%"></el-table-column>
            <el-table-column prop="Note2" label="箱单货物描述" width="220%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="420%">
              <template #default="scope">
                <el-row type="flex" justify="space-evenly">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.ID)" type="text" size="small">删除</el-button>
                  <el-button @click="fetchPrdtInfoData(scope.row.ID)" type="text" size="small">产品明细</el-button>
                  <el-button @click="fetchLoadingInfoData(scope.row.ID)" type="text" size="small">装货明细</el-button>
                  <el-button @click="fetchBuyData(scope.row.ID)" type="text" size="small">采购订单</el-button>

                  <el-button @click="fetchShouldOutData(scope.row.ID)" type="text" size="small">应付账款单</el-button>
                  <el-button @click="fetchOutData(scope.row.ID)" type="text" size="small">付款单</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="purrecData.length"
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
        <!-- <el-table-column prop="Merc" label="付款方" width="220%"></el-table-column> -->
        <!-- <el-table-column prop="AcctName" label="收款方" width="220%"></el-table-column> -->
        <!-- <el-table-column prop="BankAccName" label="付款银行账户" width="220%"></el-table-column> -->
        <!-- <el-table-column prop="AccName" label="收款银行账户" width="220%"></el-table-column> -->
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

    <el-dialog v-model="BuyVisible" title="采购订单" width="55%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="BuyID" placeholder="请输入ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addBuy(nowID)">添加</el-button>
      </div>

      <!-- 产品明细表格 -->
      <el-table :data="BuyData" height="400" style="width: 100%">
        <el-table-column prop="ID" label="ID" width="60%" />
        <el-table-column prop="OrderNum" label="订单编号" width="150%" />
        <!-- <el-table-column prop="Merc" label="购买方" width="150%" /> -->
        <!-- <el-table-column prop="AcctName" label="销售方" width="150%" /> -->
        <!-- <el-table-column prop="SpecName" label="包装规格" width="150%" /> -->
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
    <el-dialog v-model="prdtInfoVisible" title="产品明细" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="prdtInfoID" placeholder="请输入产品ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addPrdtInfo(nowID)">添加</el-button>
      </div>

      <!-- 产品明细表格 -->
      <el-table :data="prdtInfoData" height="400" style="width: 100%">
        <el-table-column prop="ID" label="ID" width="60%" />
        <!-- <el-table-column prop="CatEngName" label="产品" width="150%" /> -->
        <!-- <el-table-column prop="BrandEngName" label="品牌" width="150%" /> -->
        <!-- <el-table-column prop="PackSpec" label="包装规格" width="150%" /> -->
        <el-table-column prop="Currency" label="币种" width="100%" />
        <el-table-column prop="UnitPrice" label="单价" width="70%" />
        <el-table-column prop="TradeTerm" label="贸易条款" width="100%" />
        <!-- <el-table-column prop="DeliveryLoc" label="交货地点" width="100%" /> -->


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
    <!-- 添加发货单信息的对话框 -->

    <el-dialog v-model="LoadingInfoVisible" title="装货明细" width="52%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="LoadingInfoID" placeholder="请输入装货明细ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addLoadingInfo(nowID)">添加</el-button>
      </div>

      <el-table :data="LoadingInfoData" height="400" style="width: 100%">
        <el-table-column prop="ID" label="ID" width="60%" />
        <!-- <el-table-column prop="Product" label="产品" width="150%" /> -->
        <!-- <el-table-column prop="Brand" label="品牌" width="150%" /> -->
        <el-table-column prop="PrdtPlant" label="生产工厂" width="150%" />
        <el-table-column prop="Currency" label="币种" width="100%" />
        <el-table-column prop="UnitPrice" label="单价" width="70%" />
        <el-table-column prop="TradeTerm" label="贸易条款" width="100%" />
        <!-- <el-table-column prop="DeliveryLoc" label="交货地点" width="100%" /> -->
        <el-table-column label="操作" fixed="right" width="150">
          <template #default="scope">

            <el-button type="text" size="small" @click="CheckLoadingInfo(scope.row.ID)">跳转</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
    <el-dialog v-model="showPurrecDialog" title="采购收货信息" width="80%" @close="resetPurrecForm">
      <el-form :model="purrecForm" label-width="150px" :rules="purrecRules" ref="purrecFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="采购发票号" prop="SaleInvNum">
              <el-input v-model="purrecForm.SaleInvNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="采购发票日期" prop="SaleInvDate">
              <el-date-picker v-model="purrecForm.SaleInvDate" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客户" prop="Acct1ID">
              <el-select v-model="purrecForm.Acct1ID" @change="onAcct1Change" placeholder="请选择客户">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="发货人" prop="MerchantID">
              <el-select v-model="purrecForm.MerchantID" @change="onMerchantChange" placeholder="请选择通知人">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

        </el-row>

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="收货人" prop="Acct2ID">
              <el-select v-model="purrecForm.Acct2ID" @change="onAcct2Change" placeholder="请选择收货人">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="通知人" prop="Acct3ID">
              <el-select v-model="purrecForm.Acct3ID" @change="onAcct3Change" placeholder="请选择通知人">
                <el-option v-for="acct in merchantData" :key="acct.ID" :label="acct.AcctName"
                  :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="起运地" prop="SrcPlace">
              <el-select v-model="purrecForm.SrcPlace" placeholder="请选择起运地">
                <el-option v-for="srcPlace in srcPlaceData" :key="srcPlace" :label="srcPlace"
                  :value="srcPlace"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="目的地" prop="Des">
              <el-select v-model="purrecForm.Des" placeholder="请选择目的地">
                <el-option v-for="des in desData" :key="des" :label="des" :value="des"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="船名" prop="ShipName">
              <el-input v-model="purrecForm.ShipName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="航次" prop="Voyage">
              <el-input v-model="purrecForm.Voyage"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总件数" prop="TotNum">
              <el-input v-model="purrecForm.TotNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpecID">
              <el-select v-model="purrecForm.PackSpecID" @change="onPackSpecChange" placeholder="请选择包装规格">
                <el-option v-for="packSpec in packSpecData" :key="packSpec.ID" :label="packSpec.SpecName"
                  :value="packSpec.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总净重" prop="TotalNetWeight">
              <el-input v-model="purrecForm.TotalNetWeight"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas1">
              <el-select v-model="purrecForm.UnitMeas1" placeholder="请选择单位">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasID" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总毛重" prop="GrossWt">
              <el-input v-model="purrecForm.GrossWt"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas2">
              <el-select v-model="purrecForm.UnitMeas2" placeholder="请选择单位">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasID" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总体积" prop="TotVol">
              <el-input v-model="purrecForm.TotVol"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas3">
              <el-select v-model="purrecForm.UnitMeas3" placeholder="请选择单位">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasID" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="提单号" prop="BillLadNum">
              <el-input v-model="purrecForm.BillLadNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="发货（开船）日期" prop="DateOfShip">
              <el-date-picker v-model="purrecForm.DateOfShip" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="提单货物描述" prop="Note1">
              <el-input v-model="purrecForm.Note1" type="textarea" :rows="4"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="箱单货物描述" prop="Note2">
              <el-input v-model="purrecForm.Note2" type="textarea" :rows="4"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="付款方式" prop="PayMentMethodID">
              <el-select v-model="purrecForm.PayMentMethodID" @change="onPayMentMethodChange" placeholder="请选择付款方式">
                <el-option v-for="payMentMethod in payMentMethodData" :key="payMentMethod.ID"
                  :label="payMentMethod.PayMtdName" :value="payMentMethod.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="付款银行" prop="AcctBankID">
              <el-select v-model="purrecForm.AcctBankID" placeholder="请选择付款银行">
                <el-option v-for="bank in acctBankData" :key="bank.ID" :label="bank.AccName"
                  :value="bank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="提单">
              <el-upload v-if="!purrecForm.File1ID" ref="upload1Ref" action="" :limit="1" :on-change="handleFile1Change"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(purrecForm.File1ID, purrecForm.File1Name)">
                下载文件
              </el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="File1Name">
              <el-input v-model="purrecForm.File1Name" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="其他单据">
              <el-upload v-if="!purrecForm.File2ID" ref="upload2Ref" action="" :limit="1" :on-change="handleFile2Change"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(purrecForm.File2ID, purrecForm.File2Name)">
                下载文件
              </el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="File2Name">
              <el-input v-model="purrecForm.File2Name" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitPurrecForm">保存</el-button>
            <el-button @click="showPurrecDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <el-dialog v-model="showshowPurrecDialog" title="采购收货信息" width="80%" @close="resetPurrecForm">
      <el-form :model="purrecForm" label-width="150px" ref="purrecFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="采购发票号">
              <el-input v-model="purrecForm.SaleInvNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="采购发票日期">
              <el-date-picker v-model="purrecForm.SaleInvDate" type="date" placeholder="选择日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客户">
              <el-select v-model="purrecForm.Acct1ID" placeholder="请选择客户" :disabled="true">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="发货人">
              <el-select v-model="purrecForm.MerchantID" placeholder="请选择通知人" :disabled="true">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收货人">
              <el-select v-model="purrecForm.Acct2ID" placeholder="请选择收货人" :disabled="true">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="通知人">
              <el-select v-model="purrecForm.Acct3ID" placeholder="请选择通知人" :disabled="true">
                <el-option v-for="acct in merchantData" :key="acct.ID" :label="acct.AcctName"
                  :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="起运地">
              <el-select v-model="purrecForm.SrcPlace" placeholder="请选择起运地" :disabled="true">
                <el-option v-for="srcPlace in srcPlaceData" :key="srcPlace" :label="srcPlace"
                  :value="srcPlace"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="目的地">
              <el-select v-model="purrecForm.Des" placeholder="请选择目的地" :disabled="true">
                <el-option v-for="des in desData" :key="des" :label="des" :value="des"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="船名">
              <el-input v-model="purrecForm.ShipName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="航次">
              <el-input v-model="purrecForm.Voyage" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总件数">
              <el-input v-model="purrecForm.TotNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装规格">
              <el-select v-model="purrecForm.PackSpecID" placeholder="请选择包装规格" :disabled="true">
                <el-option v-for="packSpec in packSpecData" :key="packSpec.ID" :label="packSpec.SpecName"
                  :value="packSpec.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总净重">
              <el-input v-model="purrecForm.TotalNetWeight" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位">
              <el-select v-model="purrecForm.UnitMeas1" placeholder="请选择单位" :disabled="true">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasID" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总毛重">
              <el-input v-model="purrecForm.GrossWt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位">
              <el-select v-model="purrecForm.UnitMeas2" placeholder="请选择单位" :disabled="true">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasID" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总体积">
              <el-input v-model="purrecForm.TotVol" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位">
              <el-select v-model="purrecForm.UnitMeas3" placeholder="请选择单位" :disabled="true">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasID" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="提单号">
              <el-input v-model="purrecForm.BillLadNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="发货（开船）日期">
              <el-date-picker v-model="purrecForm.DateOfShip" type="date" placeholder="选择日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="提单货物描述">
              <el-input v-model="purrecForm.Note1" type="textarea" :rows="4" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="箱单货物描述">
              <el-input v-model="purrecForm.Note2" type="textarea" :rows="4" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="付款方式">
              <el-select v-model="purrecForm.PayMentMethodID" placeholder="请选择付款方式" :disabled="true">
                <el-option v-for="payMentMethod in payMentMethodData" :key="payMentMethod.ID"
                  :label="payMentMethod.PayMtdName" :value="payMentMethod.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="收款银行">
              <el-select v-model="purrecForm.BankAccountID" placeholder="请选择收款银行" :disabled="true">
                <el-option v-for="bank in bankAccountData" :key="bank.ID" :label="bank.BankAccName"
                  :value="bank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="提单">
              <el-button v-if="purrecForm.File1ID" type="success"
                @click="downloadFile(purrecForm.File1ID, purrecForm.File1Name)">
                下载文件
              </el-button>
              <span v-else>无文件</span>
            </el-form-item>
            <el-form-item label="文件名">
              <el-input v-model="purrecForm.File1Name" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="其他单据">
              <el-button v-if="purrecForm.File2ID" type="success"
                @click="downloadFile(purrecForm.File2ID, purrecForm.File2Name)">
                下载文件
              </el-button>
              <span v-else>无文件</span>
            </el-form-item>
            <el-form-item label="文件名">
              <el-input v-model="purrecForm.File2Name" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showPurrecDialog = false">关闭</el-button>
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

const OutVisible = ref(false)
const OutID = ref(null)
const OutData = ref([])

const fetchOutData = async (ID) => {
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/purrec/out', params, {
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

    axios.post('/delete/purrec/out', params, {
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

    const response = await axios.post('/add/purrec/out', params, {
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

const CheckLoadingInfo = (ID) => {
  try {
    // 确保 ID 是字符串
    const searchQuery = String(ID);
    console.log(searchQuery)

    // 使用路由的 resolve 方法生成完整路径
    const route = router.resolve({
      name: 'LoadingInfo', // 路由名称
      query: { searchQuery }, // 传递的查询参数（对象形式）
    });

    // 在新标签页打开
    window.open(route.href, '_blank');
  } catch (error) {
    ElMessage.error("查看失败");
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
const isExactMatch = ref(true);
const onlyID = ref(true);
const toggleMatchMode = () => {
  isExactMatch.value = !isExactMatch.value;
  console.log("isExactMatch", isExactMatch.value)
};

const toggleIDMode = () => {
  onlyID.value = !onlyID.value;
};
const LoadingInfoVisible = ref(false)
const LoadingInfoID = ref(null)
const LoadingInfoData = ref([])

const fetchLoadingInfoData = async (ID) => {
  console.log("id ? ", ID)

  try {
    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/purrec/loadingInfo', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    });
    LoadingInfoVisible.value = true;
    // prdtInfoData.value = response.data.PrdtInfo; // 假设返回的数据结构中有 PrdtInfo 字段

    LoadingInfoData.value = Object.assign([], response.data.LoadingInfo); // 强制更新
    nowID.value = ID
    console.log("nowid", nowID.value)
  } catch (error) {
    console.error('获取产品明细失败:', error);
    ElMessage.error('获取失败');
  }
};
const addLoadingInfo = async (ID) => {

  console.log(nowID.value)
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("LoadingInfoID", LoadingInfoID.value)

    const response = await axios.post('/add/purrec/loadingInfo', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
    ElMessage.success("添加成功");
    fetchLoadingInfoData(nowID.value)
    LoadingInfoVisible.value = ''
  } catch (error) {
    console.error('添加失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

const BuyVisible = ref(false)
const BuyID = ref(null)
const BuyData = ref([])

const fetchBuyData = async (ID) => {
  console.log("id ? ", ID)

  try {
    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段

    const response = await axios.post('/find/purrec/buy', params, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    });
    BuyVisible.value = true;
    // prdtInfoData.value = response.data.PrdtInfo; // 假设返回的数据结构中有 PrdtInfo 字段

    BuyData.value = Object.assign([], response.data.Buy); // 强制更新
    nowID.value = ID
    console.log("nowid", nowID.value)
  } catch (error) {
    console.error('获取失败:', error);
    ElMessage.error('获取失败');
  }
};
const addBuy = async (ID) => {

  console.log(nowID.value)
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("BuyID", BuyID.value)

    const response = await axios.post('/add/purrec/buy', params, {
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

    const response = await axios.post('/find/purrec/shouldOut', params, {
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

    const response = await axios.post('/add/purrec/shouldOut', params, {
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

    axios.post('/delete/purrec/shouldOut', params, {
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
const prdtInfoVisible = ref(false);
const nowID = ref(null);

const prdtInfoID = ref(null);

const prdtInfoData = ref([]); // 存储产品明细数据

// 删除按钮逻辑
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

    axios.post('/delete/purrec/buy', params, {
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
          fetchBuyData(nowID.value); // 重新获取应付账款单信息数据
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

    const response = await axios.post('/find/purrec/prdtInfo', params, {
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
    ElMessage.error('获取失败');
  }
};
const addPrdtInfo = async (ID) => {

  console.log(nowID.value)
  try {

    const params = new URLSearchParams();
    params.append('ID', ID); // 添加表单字段
    params.append("PrdtInfoID", prdtInfoID.value)

    const response = await axios.post('/add/purrec/prdtInfo', params, {
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
const file1 = ref(null);
const file2 = ref(null);
const searchQuery = ref(''); // 添加搜索查询字段
const currentPage = ref(1); // 当前页码
const pageSize = 8; // 每页显示的行数


const paginatedPurrecData = computed(() => {
  let filteredData = purrecData.value;
  console.log(purrecData.value);
  if (searchQuery.value) {
    if (isExactMatch.value === false) {
      if (onlyID.value === false) {
        console.log("打印了");
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value) ||
          item.SaleInvNum.includes(searchQuery.value) ||
          item.Merchant1Name.includes(searchQuery.value) ||
          item.Merchant2Name.includes(searchQuery.value) ||
          item.Merchant3Name.includes(searchQuery.value) ||
          item.AcctName.includes(searchQuery.value) ||
          item.SrcPlace.includes(searchQuery.value) ||
          item.Des.includes(searchQuery.value) ||
          item.SaleInvDate.includes(searchQuery.value) ||
          item.ShipName.includes(searchQuery.value) ||
          item.Voyage.includes(searchQuery.value) ||
          item.TotNum.toString().includes(searchQuery.value) ||
          item.SpecName.includes(searchQuery.value) ||
          item.TotalNetWeight.toString().includes(searchQuery.value) ||
          item.UnitMeas1.includes(searchQuery.value) ||
          item.GrossWt.toString().includes(searchQuery.value) ||
          item.UnitMeas2.includes(searchQuery.value) ||
          item.TotVol.includes(searchQuery.value) ||
          item.UnitMeas3.includes(searchQuery.value) ||
          item.BillLadNum.includes(searchQuery.value) ||
          item.DateOfShip.includes(searchQuery.value) ||
          item.Note1.includes(searchQuery.value) ||
          item.Note2.includes(searchQuery.value) ||
          item.PayMtdName.includes(searchQuery.value) ||
          item.AccName.includes(searchQuery.value)
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
          item.SaleInvNum === searchQuery.value ||
          item.Merchant1Name === searchQuery.value ||
          item.Merchant2Name === searchQuery.value ||
          item.Merchant3Name === searchQuery.value ||
          item.AcctName === searchQuery.value ||
          item.SrcPlace === searchQuery.value ||
          item.Des === searchQuery.value ||
          item.SaleInvDate === searchQuery.value ||
          item.ShipName === searchQuery.value ||
          item.Voyage === searchQuery.value ||
          item.TotNum.toString() === searchQuery.value ||
          item.SpecName === searchQuery.value ||
          item.TotalNetWeight === searchQuery.value ||
          item.UnitMeas1 === searchQuery.value ||
          item.GrossWt === searchQuery.value ||
          item.UnitMeas2 === searchQuery.value ||
          item.TotVol === searchQuery.value ||
          item.UnitMeas3 === searchQuery.value ||
          item.BillLadNum === searchQuery.value ||
          item.DateOfShip === searchQuery.value ||
          item.Note1 === searchQuery.value ||
          item.Note2 === searchQuery.value ||
          item.PayMtdName === searchQuery.value ||
          item.AccName === searchQuery.value
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
  fetchPurrecData(); // 获取采购收货单信息
  fetchCurrencyData(); // 新增：获取货币数据

  searchQuery.value = route.query.searchQuery || '';
  // searchQuery.value = ''
});

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

const fetchPurrecData = async () => {
  try {
    const response = await axios.get('/find/purrec'); // 调用采购收货单信息接口
    purrecData.value = response.data.Purrec; // 假设返回的数据结构中有 Purrec 字段
  } catch (error) {
    console.error('获取采购收货单信息失败:', error);
    ElMessage.error('获取采购收货单信息失败，请稍后重试');
  }
};

// 采购收货单信息对话框显示状态
const showPurrecDialog = ref(false);
const showshowPurrecDialog = ref(false);
const purrecForm = ref({
  SaleInvNum: '', // 采购发票号
  SaleInvDate: '', // 采购发票日期
  // Sales: [], // 采购订单，多表关联
  Acct1ID: '',
  Acct1Name: '',
  Acct2ID: '',
  Acct2Name: '',
  Acct3ID: '',
  Acct3Name: '',
  SrcPlace: '', // 起运地
  Des: '', // 目的地
  ShipName: '', // 船名
  Voyage: '', // 航次
  TotNum: '', // 总件数
  PackSpecID: '', // 包装规格 ID
  SpecName: '', // 包装规格名称
  TotalNetWeight: '', // 总净重
  UnitMeas1: '', // 单位 1
  GrossWt: '', // 总毛重
  UnitMeas2: '', // 单位 2
  TotVol: '', // 总体积
  UnitMeas3: '', // 尺码
  BillLadNum: '', // 提单号
  DateOfShip: '', // 发货（开船）日期
  Note1: '', // 提单货物描述
  Note2: '', // 箱单货物描述
  PayMentMethodID: '', // 付款方式 ID
  PayMtdName: '', // 付款方式名称
  AcctBankID: '', // 收款银行 ID
  AccName: '', // 收款银行名称
  // PrdtInfos: [], // 产品明细，多表关联
  // LoadingInfos: [], // 装货明细，多表关联
  File1Name: '',
  File1ID: '',
  File2Name: '',
  File2ID: '',
  MerchantID: '',
  MerchantName: '',

  // ShouldIns: [], // 应付账款单
  // Ins: [], // 付款单
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
// 采购收货单信息表单验证规则
const purrecRules = {
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
const purrecData = ref([]); // 采购收货单信息

// 根据当前选中的菜单项动态更改标题和按钮文本
const headerTitle = computed(() => {
  return '采购收货单信息';
});

const addButtonText = computed(() => {
  return '添加采购收货单信息';
});

// 添加按钮点击事件
const handleAdd = () => {
  showPurrecDialog.value = true;
};

const handleEdit = (index, row) => {
  // 将当前行的数据赋值给 purrecForm
  purrecForm.value = { ...row };

  // 如果 PrdtInfos 是对象数组，转换为 ID 数组
  if (row.PrdtInfos && Array.isArray(row.PrdtInfos)) {
    purrecForm.value.PrdtInfos = row.PrdtInfos.map(item => item.ID);
  }

  if (row.DocReq && Array.isArray(row.DocReq)) {
    purrecForm.value.DocReq = row.DocReq.map(item => item.DocReqID);
  }
  // // 检查是否已上传文件
  // if (row.FileID) {
  //   purrecForm.value.FileID = row.FileID; // 保存 FileID
  //   purrecForm.value.FileName = row.FileName; // 保存文件名
  // }

  showPurrecDialog.value = true; // 打开采购收货单信息对话框
};

// 查看按钮逻辑
const handleView = (index, row) => {
  purrecForm.value = { ...row }; // 将当前行的数据赋值给 purrecForm
  showshowPurrecDialog.value = true; // 打开查看采购收货单信息对话框
};

// 删除按钮逻辑
const handleDelete = (index, ID) => {
  ElMessageBox.confirm('确定要删除该采购收货单信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {

    const formData = new FormData();
    formData.append("ID", ID)
    axios.post('/delete/purrec', formData,
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
        }
      },
    )
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchPurrecData(); // 重新获取采购收货单信息数据
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

// 重置采购收货单信息表单
const resetPurrecForm = () => {
  purrecForm.value = {
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
    File1Name: '',
    File1ID: '',
    File2Name: '',
    File2ID: '',
    Merchant1ID: '',
    Merchant2ID: '',
    Merchant3ID: '',
    Merchant1Name: '',
    Merchant2Name: '',
    Merchant3Name: '',
  };
  if (upload1Ref.value) {
    upload1Ref.value.clearFiles(); // 清空文件列表
  }
  if (upload2Ref.value) {
    upload2Ref.value.clearFiles(); // 清空文件列表
  }
};

const upload1Ref = ref(null);
const upload2Ref = ref(null);

const submitPurrecForm = async () => {
  try {
    const formData = new FormData();

    // 添加其他字段
    Object.keys(purrecForm.value).forEach((key) => {
      formData.append(key, purrecForm.value[key]);
    });

    // 添加文件
    if (file1.value) {
      formData.append('file1', file1.value);
    }

    if (file2.value) {
      formData.append('file2', file2.value);
    }
    // 提交表单
    console.log(formData)
    const response = await axios.post('/save/purrec', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    if (response.status === 200) {
      ElMessage.success('采购收货单信息保存成功');
      showPurrecDialog.value = false; // 关闭对话框
      fetchPurrecData(); // 重新获取采购收货单信息数据
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存采购收货单信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};
// 监听 change 事件并更新 AcctName
const onAcctChange = (value) => {
  const selectedAcct = acctData.value.find(acct => acct.ID === value);
  if (selectedAcct) {
    purrecForm.value.AcctName = selectedAcct.AcctName;
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
    ElMessage.error('文件下载失败，请稍后重试');
  }
};

const handleFile1Change = (up) => {
  file1.value = up.raw; // 保存选择的文件
  purrecForm.value.File1Name = up.name; // 更新文件名
};
const handleFile2Change = (up) => {
  file2.value = up.raw; // 保存选择的文件
  purrecForm.value.File2Name = up.name; // 更新文件名
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
    purrecForm.value.Merc = selectedMerchant.Merc;
  }
};

const onAcct1Change = (value) => {
  const selectedAcct = acctData.value.find(acct => acct.ID === value);
  if (selectedAcct) {
    purrecForm.value.Acct1Name = selectedAcct.AcctName;
  }
};
const onAcct2Change = (value) => {

  const selectedAcct = acctData.value.find(acct => acct.ID === value);
  if (selectedAcct) {
    purrecForm.value.Acct2Name = selectedAcct.AcctName;
  }
};
const onAcct3Change = (value) => {

  const selectedAcct = acctData.value.find(acct => acct.ID === value);
  if (selectedAcct) {
    purrecForm.value.Acct3Name = selectedAcct.AcctName;
  }
};
// 监听 change 事件并更新 PayMtdName
const onPayMentMethodChange = (value) => {
  const selectedPayMentMethod = payMentMethodData.value.find(payMentMethod => payMentMethod.ID === value);
  if (selectedPayMentMethod) {
    purrecForm.value.PayMtdName = selectedPayMentMethod.PayMtdName;
  }
};

// 监听 change 事件并更新 SpecName
const onPackSpecChange = (value) => {
  const selectedPackSpec = packSpecData.value.find(packSpec => packSpec.ID === value);
  if (selectedPackSpec) {
    purrecForm.value.SpecName = selectedPackSpec.SpecName;
  }
};

// 监听 change 事件并更新 AccName
const onAcctBankChange = (value) => {
  const selectedAcctBank = acctBankData.value.find(acctBank => acctBank.ID === value);
  if (selectedAcctBank) {
    purrecForm.value.AccName = selectedAcctBank.AccName;
  }
};

// 监听 change 事件并更新 BankAccName
const onBankAccountChange = (value) => {
  const selectedBankAccount = bankAccountData.value.find(bankAccount => bankAccount.ID === value);
  if (selectedBankAccount) {
    purrecForm.value.BankAccName = selectedBankAccount.BankAccName;
  }
};

// 分页逻辑
const handlePageChange = (page) => {
  currentPage.value = page;
};
</script>

<style src="../assets/styles/Bottom.css"></style>

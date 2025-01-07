<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-17'" />
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
          <el-table :data="paginatedSendData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="SaleInvNum" label="销售发票号" width="220%"></el-table-column>
            <el-table-column prop="Merchant1Name" label="客户" width="220%"></el-table-column>
            <el-table-column prop="Merchant2Name" label="收货人" width="220%"></el-table-column>
            <el-table-column prop="Merchant3Name" label="通知人" width="220%"></el-table-column>
            <el-table-column prop="AcctName" label="发货人" width="420%"></el-table-column>
            <el-table-column prop="SrcPlace" label="起运地" width="220%"></el-table-column>
            <el-table-column prop="Des" label="目的地" width="220%"></el-table-column>
            <el-table-column prop="SaleInvDate" label="销售发票日期" width="420%"></el-table-column>
            <el-table-column prop="ShipName" label="船名" width="220%"></el-table-column>
            <el-table-column prop="Voyage" label="航次" width="220%"></el-table-column>
            <el-table-column prop="TotNum" label="总件数" width="220%"></el-table-column>
            <el-table-column prop="SpecName" label="包装规格" width="220%"></el-table-column>
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
            <el-table-column prop="PayMtdName" label="付款方式" width="220%"></el-table-column>
            <el-table-column prop="AccName" label="收款银行" width="220%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="420%">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.ID)" type="text" size="small">删除</el-button>
                  <el-button @click="fetchPrdtInfoData(scope.row.ID)" type="text" size="small">产品明细</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="sendData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>


    <el-dialog v-model="prdtInfoVisible" title="产品明细" width="80%">
      <!-- 添加按钮和输入框 -->
      <div style="text-align: right; margin-bottom: 20px;">
        <el-input v-model="prdtInfoId" placeholder="请输入产品ID" style="width: 200px; margin-right: 10px;" />
        <el-button type="primary" @click="addPrdtInfo(nowId)">添加</el-button>
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
        <el-table-column label="操作" width="150">
          <template #default="scope">
            <!-- <el-button type="text" size="small" @click="viewProduct(scope.row)">查看</el-button> -->
            <el-button type="text" size="small"
              @click="DeletePrdtInfo(scope.$index, nowId, scope.row.ID)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
    <!-- 添加发货单信息的对话框 -->

    <el-dialog v-model="showSendDialog" title="销售发货信息" width="80%" @close="resetSendForm">
      <el-form :model="sendForm" label-width="150px" :rules="sendRules" ref="sendFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="销售发票号" prop="SaleInvNum">
              <el-input v-model="sendForm.SaleInvNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="销售发票日期" prop="SaleInvDate">
              <el-date-picker v-model="sendForm.SaleInvDate" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>


        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客户" prop="Merchants">
              <el-select v-model="sendForm.Merchant1Id" @change="onMerchantChange" placeholder="请选择客户">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="发货人" prop="AcctId">
              <el-select v-model="sendForm.AcctId" @change="onAcctChange" placeholder="请选择发货人">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="收货人" prop="Merchants">
              <el-select v-model="sendForm.Merchant2Id" @change="onMerchantChange" placeholder="请选择收货人">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="通知人" prop="Merchants">
              <el-select v-model="sendForm.Merchant3Id" @change="onMerchantChange" placeholder="请选择通知人">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="起运地" prop="SrcPlace">
              <el-select v-model="sendForm.SrcPlace" placeholder="请选择起运地">
                <el-option v-for="srcPlace in srcPlaceData" :key="srcPlace" :label="srcPlace"
                  :value="srcPlace"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="目的地" prop="Des">
              <el-select v-model="sendForm.Des" placeholder="请选择目的地">
                <el-option v-for="des in desData" :key="des" :label="des" :value="des"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="船名" prop="ShipName">
              <el-input v-model="sendForm.ShipName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="航次" prop="Voyage">
              <el-input v-model="sendForm.Voyage"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总件数" prop="TotNum">
              <el-input v-model="sendForm.TotNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpecId">
              <el-select v-model="sendForm.PackSpecId" @change="onPackSpecChange" placeholder="请选择包装规格">
                <el-option v-for="packSpec in packSpecData" :key="packSpec.ID" :label="packSpec.SpecName"
                  :value="packSpec.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总净重" prop="TotalNetWeight">
              <el-input v-model="sendForm.TotalNetWeight"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas1">
              <el-select v-model="sendForm.UnitMeas1" placeholder="请选择单位">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasId" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总毛重" prop="GrossWt">
              <el-input v-model="sendForm.GrossWt"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas2">
              <el-select v-model="sendForm.UnitMeas2" placeholder="请选择单位">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasId" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总体积" prop="TotVol">
              <el-input v-model="sendForm.TotVol"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas3">
              <el-select v-model="sendForm.UnitMeas3" placeholder="请选择单位">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasId" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="提单号" prop="BillLadNum">
              <el-input v-model="sendForm.BillLadNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="发货（开船）日期" prop="DateOfShip">
              <el-date-picker v-model="sendForm.DateOfShip" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="提单货物描述" prop="Note1">
              <el-input v-model="sendForm.Note1" type="textarea" :rows="4"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="箱单货物描述" prop="Note2">
              <el-input v-model="sendForm.Note2" type="textarea" :rows="4"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="付款方式" prop="PayMentMethodId">
              <el-select v-model="sendForm.PayMentMethodId" @change="onPayMentMethodChange" placeholder="请选择付款方式">
                <el-option v-for="payMentMethod in payMentMethodData" :key="payMentMethod.ID"
                  :label="payMentMethod.PayMtdName" :value="payMentMethod.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="收款银行" prop="AcctBankId">
              <el-select v-model="sendForm.AcctBankId" @change="onAcctBankChange" placeholder="请选择收款银行">
                <el-option v-for="acctBank in acctBankData" :key="acctBank.ID" :label="acctBank.AccName"
                  :value="acctBank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="提单">
              <el-upload v-if="!sendForm.File1Id" ref="upload1Ref" action="" :limit="1" :on-change="handleFile1Change"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(sendForm.File1Id, sendForm.File1Name)">
                下载文件
              </el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="File1Name">
              <el-input v-model="sendForm.File1Name" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="其他单据">
              <el-upload v-if="!sendForm.File2Id" ref="upload2Ref" action="" :limit="1" :on-change="handleFile2Change"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success" @click="downloadFile(sendForm.File2Id, sendForm.File2Name)">
                下载文件
              </el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="File2Name">
              <el-input v-model="sendForm.File2Name" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitSendForm">保存</el-button>
            <el-button @click="showSendDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <el-dialog v-model="showshowSendDialog" title="销售发货信息" width="80%" @close="resetSendForm">
      <el-form :model="sendForm" label-width="150px" :rules="sendRules" ref="sendFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="销售发票号" prop="SaleInvNum">
              <el-input v-model="sendForm.SaleInvNum" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="销售发票日期" prop="SaleInvDate">
              <el-date-picker v-model="sendForm.SaleInvDate" type="date" placeholder="选择日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客户" prop="Merchants">
              <el-select v-model="sendForm.Merchant1Id" @change="onMerchantChange" placeholder="请选择客户" :disabled="true">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="发货人" prop="AcctId">
              <el-select v-model="sendForm.AcctId" @change="onAcctChange" placeholder="请选择发货人" :disabled="true">
                <el-option v-for="acct in acctData" :key="acct.ID" :label="acct.AcctName" :value="acct.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="收货人" prop="Merchants">
              <el-select v-model="sendForm.Merchant2Id" @change="onMerchantChange" placeholder="请选择收货人"
                :disabled="true">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="通知人" prop="Merchants">
              <el-select v-model="sendForm.Merchant3Id" @change="onMerchantChange" placeholder="请选择通知人"
                :disabled="true">
                <el-option v-for="merchant in merchantData" :key="merchant.ID" :label="merchant.Merc"
                  :value="merchant.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="起运地" prop="SrcPlace">
              <el-select v-model="sendForm.SrcPlace" placeholder="请选择起运地" :disabled="true">
                <el-option v-for="srcPlace in srcPlaceData" :key="srcPlace" :label="srcPlace"
                  :value="srcPlace"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="目的地" prop="Des">
              <el-select v-model="sendForm.Des" placeholder="请选择目的地" :disabled="true">
                <el-option v-for="des in desData" :key="des" :label="des" :value="des"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="船名" prop="ShipName">
              <el-input v-model="sendForm.ShipName" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="航次" prop="Voyage">
              <el-input v-model="sendForm.Voyage" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总件数" prop="TotNum">
              <el-input v-model="sendForm.TotNum" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpecId">
              <el-select v-model="sendForm.PackSpecId" @change="onPackSpecChange" placeholder="请选择包装规格"
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
              <el-input v-model="sendForm.TotalNetWeight" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas1">
              <el-select v-model="sendForm.UnitMeas1" placeholder="请选择单位" :disabled="true">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasId" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总毛重" prop="GrossWt">
              <el-input v-model="sendForm.GrossWt" :disabled="true"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas2">
              <el-select v-model="sendForm.UnitMeas2" placeholder="请选择单位" :disabled="true">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasId" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="总体积" prop="TotVol">
              <el-input v-model="sendForm.TotVol" :disabled="true"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="单位" prop="UnitMeas3">
              <el-select v-model="sendForm.UnitMeas3" placeholder="请选择单位" :disabled="true">
                <el-option v-for="unitMeas in unitMeasData" :key="unitMeas.UnitMeasId" :label="unitMeas.UnitMeas"
                  :value="unitMeas.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="提单号" prop="BillLadNum">
              <el-input v-model="sendForm.BillLadNum" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="发货（开船）日期" prop="DateOfShip">
              <el-date-picker v-model="sendForm.DateOfShip" type="date" placeholder="选择日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="提单货物描述" prop="Note1">
              <el-input v-model="sendForm.Note1" type="textarea" :rows="4" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="箱单货物描述" prop="Note2">
              <el-input v-model="sendForm.Note2" type="textarea" :rows="4" :disabled="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="付款方式" prop="PayMentMethodId">
              <el-select v-model="sendForm.PayMentMethodId" @change="onPayMentMethodChange" placeholder="请选择付款方式"
                :disabled="true">
                <el-option v-for="payMentMethod in payMentMethodData" :key="payMentMethod.ID"
                  :label="payMentMethod.PayMtdName" :value="payMentMethod.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="收款银行" prop="AcctBankId">
              <el-select v-model="sendForm.AcctBankId" @change="onAcctBankChange" placeholder="请选择收款银行"
                :disabled="true">
                <el-option v-for="acctBank in acctBankData" :key="acctBank.ID" :label="acctBank.AccName"
                  :value="acctBank.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="提单">
              <el-button v-if="sendForm.File1Id" type="success"
                @click="downloadFile(sendForm.File1Id, sendForm.File1Name)">
                下载文件
              </el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="File1Name">
              <el-input v-model="sendForm.File1Name" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="其他单据">
              <el-button v-if="sendForm.File2Id" type="success"
                @click="downloadFile(sendForm.File2Id, sendForm.File2Name)">
                下载文件
              </el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="File2Name">
              <el-input v-model="sendForm.File2Name" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showshowSendDialog = false">关闭</el-button>
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

const prdtInfoData = ref([]); // 存储产品明细数据

const fetchPrdtInfoData = async (ID) => {
  try {
    const response = await axios.get('/find/send/prdtInfo', {
      "ID": ID,
    }); // 调用产品明细接口
    prdtInfoData.value = response.data.PrdtInfo; // 假设返回的数据结构中有 PrdtInfo 字段
    prdtInfoVisible.value = true;
    nowId.value = ID
    console.log("nowid", nowId.value)
  } catch (error) {
    console.error('获取产品明细失败:', error);
    ElMessage.error('获取产品明细失败');
  }
};
const prdtInfoVisible = ref(false);
const nowId = ref(null);

const prdtInfoId = ref(null);
const addPrdtInfo = async (ID) => {

  console.log(nowId.value)
  // acctForm.value.ID = parseInt(acctForm.value.ID, 10);
  try {
    const response = await axios.post('/add/send/prdtInfo', {
      "ID": ID,
      "PrdtInfoId": parseInt(prdtInfoId.value, 10)
    }); // 调用产品明细接口

    ElMessage.success("添加成功");
    fetchPrdtInfoData(nowId.value)
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

const paginatedSendData = computed(() => {
  let filteredData = sendData.value;

  // 如果有搜索条件，过滤数据
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.SaleInvNum.includes(searchQuery.value) ||
      item.SaleInvDate.includes(searchQuery.value) ||
      // item.Merchant1.includes(searchQuery.value) ||
      // item.Merchant2.includes(searchQuery.value) ||
      // item.Merchant3.includes(searchQuery.value) ||
      item.AcctName.includes(searchQuery.value) ||
      item.SrcPlace.includes(searchQuery.value) ||
      item.Des.includes(searchQuery.value) ||
      item.ShipName.includes(searchQuery.value) ||
      item.Voyage.includes(searchQuery.value) ||
      item.TotNum.toString().includes(searchQuery.value) ||
      item.SpecName.includes(searchQuery.value) ||
      item.TotalNetWeight.includes(searchQuery.value) ||
      item.UnitMeas1.includes(searchQuery.value) ||
      item.GrossWt.includes(searchQuery.value) ||
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
  }

  // 计算分页的起始和结束位置
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;

  // 返回分页后的数据
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
  fetchSendData(); // 获取销售订单信息
  fetchCurrencyData(); // 新增：获取货币数据
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

const fetchSendData = async () => {
  try {
    const response = await axios.get('/find/send'); // 调用销售订单信息接口
    sendData.value = response.data.Send; // 假设返回的数据结构中有 Send 字段
  } catch (error) {
    console.error('获取销售订单信息失败:', error);
    ElMessage.error('获取销售订单信息失败，请稍后重试');
  }
};

// 销售订单信息对话框显示状态
const showSendDialog = ref(false);
const showshowSendDialog = ref(false);
const sendForm = ref({
  SaleInvNum: '', // 销售发票号
  SaleInvDate: '', // 销售发票日期
  // Sales: [], // 销售订单，多表关联
  AcctId: '', // 发货人 ID
  AcctName: '', // 发货人名称
  SrcPlace: '', // 起运地
  Des: '', // 目的地
  ShipName: '', // 船名
  Voyage: '', // 航次
  TotNum: '', // 总件数
  PackSpecId: '', // 包装规格 ID
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
  PayMentMethodId: '', // 付款方式 ID
  PayMtdName: '', // 付款方式名称
  AcctBankId: '', // 收款银行 ID
  AccName: '', // 收款银行名称
  // PrdtInfos: [], // 产品明细，多表关联
  // LoadingInfos: [], // 装货明细，多表关联
  File1Name: '',
  File1Id: '',
  File2Name: '',
  File2Id: '',
  Merchant1Id: '',
  Merchant2Id: '',
  Merchant3Id: '',
  Merchant1Name: '',
  Merchant2Name: '',
  Merchant3Name: '',

  // ShouldIns: [], // 应收账款单
  // Ins: [], // 收款单
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
const sendRules = {
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
const sendData = ref([]); // 销售订单信息

// 根据当前选中的菜单项动态更改标题和按钮文本
const headerTitle = computed(() => {
  return '销售发货单信息';
});

const addButtonText = computed(() => {
  return '添加销售发货单信息';
});

// 添加按钮点击事件
const handleAdd = () => {
  showSendDialog.value = true;
};

const handleEdit = (index, row) => {
  // 将当前行的数据赋值给 sendForm
  sendForm.value = { ...row };

  // 如果 PrdtInfos 是对象数组，转换为 ID 数组
  if (row.PrdtInfos && Array.isArray(row.PrdtInfos)) {
    sendForm.value.PrdtInfos = row.PrdtInfos.map(item => item.ID);
  }

  if (row.DocReq && Array.isArray(row.DocReq)) {
    sendForm.value.DocReq = row.DocReq.map(item => item.DocReqId);
  }
  // // 检查是否已上传文件
  // if (row.FileId) {
  //   sendForm.value.FileId = row.FileId; // 保存 FileId
  //   sendForm.value.FileName = row.FileName; // 保存文件名
  // }

  showSendDialog.value = true; // 打开销售订单信息对话框
};

// 查看按钮逻辑
const handleView = (index, row) => {
  sendForm.value = { ...row }; // 将当前行的数据赋值给 sendForm
  showshowSendDialog.value = true; // 打开查看销售订单信息对话框
};

// 删除按钮逻辑
const handleDelete = (index, ID) => {
  ElMessageBox.confirm('确定要删除该销售订单信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/send', {
      "ID": ID,
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchSendData(); // 重新获取销售订单信息数据
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
const resetSendForm = () => {
  sendForm.value = {
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
    File1Name: '',
    File1Id: '',
    File2Name: '',
    File2Id: '',
    Merchant1Id: '',
    Merchant2Id: '',
    Merchant3Id: '',
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

const submitSendForm = async () => {
  try {
    const formData = new FormData();

    // 添加其他字段
    Object.keys(sendForm.value).forEach((key) => {
      formData.append(key, sendForm.value[key]);
    });

    // 添加文件
    if (file1.value) {
      formData.append('file1', file1.value);
    }

    if (file2.value) {
      formData.append('file2', file2.value);
    }
    // 提交表单
    const response = await axios.post('/save/send', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    if (response.status === 200) {
      ElMessage.success('销售订单信息保存成功');
      showSendDialog.value = false; // 关闭对话框
      fetchSendData(); // 重新获取销售订单信息数据
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
    sendForm.value.AcctName = selectedAcct.AcctName;
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

const handleFile1Change = (up) => {
  file1.value = up.raw; // 保存选择的文件
  sendForm.value.File1Name = up.name; // 更新文件名
};
const handleFile2Change = (up) => {
  file2.value = up.raw; // 保存选择的文件
  sendForm.value.File2Name = up.name; // 更新文件名
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
    sendForm.value.Merc = selectedMerchant.Merc;
  }
};

// 监听 change 事件并更新 PayMtdName
const onPayMentMethodChange = (value) => {
  const selectedPayMentMethod = payMentMethodData.value.find(payMentMethod => payMentMethod.ID === value);
  if (selectedPayMentMethod) {
    sendForm.value.PayMtdName = selectedPayMentMethod.PayMtdName;
  }
};

// 监听 change 事件并更新 SpecName
const onPackSpecChange = (value) => {
  const selectedPackSpec = packSpecData.value.find(packSpec => packSpec.ID === value);
  if (selectedPackSpec) {
    sendForm.value.SpecName = selectedPackSpec.SpecName;
  }
};

// 监听 change 事件并更新 AccName
const onAcctBankChange = (value) => {
  const selectedAcctBank = acctBankData.value.find(acctBank => acctBank.ID === value);
  if (selectedAcctBank) {
    sendForm.value.AccName = selectedAcctBank.AccName;
  }
};

// 监听 change 事件并更新 BankAccName
const onBankAccountChange = (value) => {
  const selectedBankAccount = bankAccountData.value.find(bankAccount => bankAccount.ID === value);
  if (selectedBankAccount) {
    sendForm.value.BankAccName = selectedBankAccount.BankAccName;
  }
};

// 分页逻辑
const handlePageChange = (page) => {
  currentPage.value = page;
};
</script>

<style src="../assets/styles/Bottom.css"></style>

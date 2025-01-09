<template>
  <el-container>
    <!-- 左侧导航栏 -->
    <el-aside width="200px">
      <el-menu default-active="customer" @select="handleMenuSelect">
        <el-menu-item index="customer">客户类型</el-menu-item>
        <el-menu-item index="supplier">供应商类型</el-menu-item>
        <el-menu-item index="PrdtType">产品类型</el-menu-item>
        <el-menu-item index="FoodAddType">食品添加剂类型</el-menu-item>
        <el-menu-item index="FeedAddType">饲料添加剂类型</el-menu-item>
        <el-menu-item index="UnitMeas">计量单位</el-menu-item>
        <el-menu-item index="PackType">包装类型</el-menu-item>
        <el-menu-item index="ConType">集装箱类型</el-menu-item>
        <el-menu-item index="Currency">币种</el-menu-item>
        <el-menu-item index="TradeTerm">贸易条款</el-menu-item>
        <el-menu-item index="Nation">国家</el-menu-item>
        <el-menu-item index="Port">港口</el-menu-item>
        <el-menu-item index="TaxType">税号类型</el-menu-item>
        <el-menu-item index="BrandType">品牌类型</el-menu-item>
        <el-menu-item index="EduLevel">学历</el-menu-item>
        <el-menu-item index="Dept">公司部门</el-menu-item>
        <el-menu-item index="Position">公司岗位</el-menu-item>
        <el-menu-item index="QualStd">质量标准</el-menu-item>
        <el-menu-item index="InvLoc">库存地点位置</el-menu-item>
        <el-menu-item index="DocReq">单据要求</el-menu-item>
        <el-menu-item index="PayMth">转账方式</el-menu-item>
        <el-menu-item index="PayLimit">后付款转账期限</el-menu-item>
        <el-menu-item index="FinaDocStatus">财务单据状态</el-menu-item>
        <el-menu-item index="FinaDocType">财务单据类型</el-menu-item>
        <el-menu-item index="ExpType">费用类型</el-menu-item>
        <el-menu-item index="Rates">收费标准</el-menu-item>
        <el-menu-item index="BussOrderSta">业务单据状态</el-menu-item>
      </el-menu>
    </el-aside>

    <!-- 右侧内容 -->
    <el-main>
      <!-- 添加按钮 -->
      <div class="header-buttons">
        <el-button @click="handleBack" type="default" icon="el-icon-arrow-left">返回</el-button>
        <el-button type="primary" @click="handleAdd">添加</el-button>
      </div>

      <!-- 表格 -->
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="type" label="类型" />
        <el-table-column label="操作">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 添加/编辑对话框 -->
      <el-dialog v-model="dialogVisible" :title="dialogTitle">
        <el-form :model="form">
          <el-form-item label="类型">
            <el-input v-model="form.type" />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { ElMessageBox, ElMessage } from 'element-plus';
export default {
  setup() {
    const activeMenu = ref('customer'); // 当前选中的菜单项
    const tableData = ref([]); // 表格数据
    const dialogVisible = ref(false); // 对话框是否可见
    const dialogTitle = ref(''); // 对话框标题
    const form = ref({ type: '', id: '' }); // 表单数据
    const isEdit = ref(false); // 是否是编辑模式


    const handleBack = () => {
      window.history.back();
    };
    // 获取数据
    const fetchData = async () => {
      let url;
      if (activeMenu.value === 'customer') {
        url = '/find/mercType';
      }
      else if (activeMenu.value === 'supplier') {
        url = '/find/suprType';
      }
      else if (activeMenu.value === 'PrdtType') {
        url = '/find/prdtType';
      }
      else if (activeMenu.value === 'FoodAddType') {
        url = '/find/foodAddType';
      }
      else if (activeMenu.value === 'FeedAddType') {
        url = '/find/feedAddType';
      }
      else if (activeMenu.value === 'UnitMeas') {
        url = '/find/unitMeas';
      }
      else if (activeMenu.value === 'PackType') {
        url = '/find/packType';
      }
      else if (activeMenu.value === 'ConType') {
        url = '/find/conType';
      }
      else if (activeMenu.value === 'Currency') {
        url = '/find/currency';
      }
      else if (activeMenu.value === 'TradeTerm') {
        url = '/find/tradeTerm';
      }
      else if (activeMenu.value === 'Nation') {
        url = '/find/nation';
      }
      else if (activeMenu.value === 'Port') {
        url = '/find/port';
      }
      else if (activeMenu.value === 'TaxType') {
        url = '/find/taxType';
      }
      else if (activeMenu.value === 'BrandType') {
        url = '/find/brandType';
      }
      else if (activeMenu.value === 'EduLevel') {
        url = '/find/eduLevel';
      }
      else if (activeMenu.value === 'Dept') {
        url = '/find/dept';
      }
      else if (activeMenu.value === 'Position') {
        url = '/find/position';
      }
      else if (activeMenu.value === 'QualStd') {
        url = '/find/qualStd';
      }
      else if (activeMenu.value === 'InvLoc') {
        url = '/find/invLoc';
      }
      else if (activeMenu.value === 'DocReq') {
        url = '/find/docReq';
      }
      else if (activeMenu.value === 'PayMth') {
        url = '/find/payMth';
      }
      else if (activeMenu.value === 'PayLimit') {
        url = '/find/payLimit';
      }
      else if (activeMenu.value === 'FinaDocStatus') {
        url = '/find/finaDocStatus';
      }
      else if (activeMenu.value === 'FinaDocType') {
        url = '/find/finaDocType';
      }
      else if (activeMenu.value === 'ExpType') {
        url = '/find/expType';
      }
      else if (activeMenu.value === 'Rates') {
        url = '/find/rates';
      }
      else if (activeMenu.value === 'BussOrderSta') {
        url = '/find/bussOrderSta';
      }
      const response = await axios.get(url);
      if (activeMenu.value === 'customer') {
        tableData.value = response.data.MercType.map(item => ({
          type: item.MercType,
          id: item.MercTypeID,
        }));
      }
      else if (activeMenu.value === 'supplier') {
        tableData.value = response.data.SuprType.map(item => ({
          type: item.SuprType,
          id: item.SuprTypeID,
        }));
      }

      else if (activeMenu.value === 'PrdtType') {
        tableData.value = response.data.PrdtType.map(item => ({
          type: item.PrdtType,
          id: item.PrdtTypeID,
        }));
      }
      else if (activeMenu.value === 'FoodAddType') {
        tableData.value = response.data.FoodAddType.map(item => ({
          type: item.FoodAddType,
          id: item.FoodAddTypeID,
        }));
      }
      else if (activeMenu.value === 'FeedAddType') {
        tableData.value = response.data.FeedAddType.map(item => ({
          type: item.FeedAddType,
          id: item.FeedAddTypeID,
        }));
      }
      else if (activeMenu.value === 'UnitMeas') {
        tableData.value = response.data.UnitMeas.map(item => ({
          type: item.UnitMeas,
          id: item.UnitMeasID,
        }));
      }
      else if (activeMenu.value === 'PackType') {
        tableData.value = response.data.PackType.map(item => ({
          type: item.PackType,
          id: item.PackTypeID,
        }));
      }
      else if (activeMenu.value === 'ConType') {
        tableData.value = response.data.ConType.map(item => ({
          type: item.ConType,
          id: item.ConTypeID,
        }));
      }
      else if (activeMenu.value === 'Currency') {
        tableData.value = response.data.Currency.map(item => ({
          type: item.Currency,
          id: item.CurrencyID,
        }));
      }
      else if (activeMenu.value === 'TradeTerm') {
        tableData.value = response.data.TradeTerm.map(item => ({
          type: item.TradeTerm,
          id: item.TradeTermID,
        }));
      }
      else if (activeMenu.value === 'Nation') {
        tableData.value = response.data.Nation.map(item => ({
          type: item.Nation,
          id: item.NationID,
        }));
      }
      else if (activeMenu.value === 'Port') {
        tableData.value = response.data.Port.map(item => ({
          type: item.Port,
          id: item.PortID,
        }));
      }
      else if (activeMenu.value === 'TaxType') {
        tableData.value = response.data.TaxType.map(item => ({
          type: item.TaxType,
          id: item.TaxTypeID,
        }));
      }
      else if (activeMenu.value === 'BrandType') {
        tableData.value = response.data.BrandType.map(item => ({
          type: item.BrandType,
          id: item.BrandTypeID,
        }));
      }
      else if (activeMenu.value === 'EduLevel') {
        tableData.value = response.data.EduLevel.map(item => ({
          type: item.EduLevel,
          id: item.EduLevelID,
        }));
      }
      else if (activeMenu.value === 'Dept') {
        tableData.value = response.data.Dept.map(item => ({
          type: item.Dept,
          id: item.DeptID,
        }));
      }
      else if (activeMenu.value === 'Position') {
        tableData.value = response.data.Position.map(item => ({
          type: item.Position,
          id: item.PositionID,
        }));
      }
      else if (activeMenu.value === 'QualStd') {
        tableData.value = response.data.QualStd.map(item => ({
          type: item.QualStd,
          id: item.QualStdID,
        }));
      }
      else if (activeMenu.value === 'InvLoc') {
        tableData.value = response.data.InvLoc.map(item => ({
          type: item.InvLoc,
          id: item.InvLocID,
        }));
      }
      else if (activeMenu.value === 'DocReq') {
        tableData.value = response.data.DocReq.map(item => ({
          type: item.DocReq,
          id: item.DocReqID,
        }));
      }
      else if (activeMenu.value === 'PayMth') {
        tableData.value = response.data.PayMth.map(item => ({
          type: item.PayMth,
          id: item.PayMthID,
        }));
      }
      else if (activeMenu.value === 'PayLimit') {
        tableData.value = response.data.PayLimit.map(item => ({
          type: item.PayLimit,
          id: item.PayLimitID,
        }));
      }
      else if (activeMenu.value === 'FinaDocStatus') {
        tableData.value = response.data.FinaDocStatus.map(item => ({
          type: item.FinaDocStatus,
          id: item.FinaDocStatusID,
        }));
      }
      else if (activeMenu.value === 'FinaDocType') {
        tableData.value = response.data.FinaDocType.map(item => ({
          type: item.FinaDocType,
          id: item.FinaDocTypeID,
        }));
      }
      else if (activeMenu.value === 'ExpType') {
        tableData.value = response.data.ExpType.map(item => ({
          type: item.ExpType,
          id: item.ExpTypeID,
        }));
      }
      else if (activeMenu.value === 'Rates') {
        tableData.value = response.data.Rates.map(item => ({
          type: item.Rates,
          id: item.RatesID,
        }));
      }
      else if (activeMenu.value === 'BussOrderSta') {
        tableData.value = response.data.BussOrderSta.map(item => ({
          type: item.BussOrderSta,
          id: item.BussOrderStaID,
        }));
      }
    };

    // 处理菜单选择
    const handleMenuSelect = (index) => {
      activeMenu.value = index;
      fetchData();
    };

    // 处理添加按钮点击
    const handleAdd = () => {
      dialogTitle.value = '添加';
      form.value = { type: '', id: '' };
      isEdit.value = false;
      dialogVisible.value = true;
    };

    // 处理编辑按钮点击
    const handleEdit = (row) => {
      dialogTitle.value = '编辑';
      form.value = { type: row.type, id: row.id };
      isEdit.value = true;
      dialogVisible.value = true;
    };

    // 处理删除按钮点击
    const handleDelete = async (row) => {
      let url;
      if (activeMenu.value === 'customer') {
        url = '/delete/mercType';
      }
      else if (activeMenu.value === 'supplier') {
        url = '/delete/suprType';
      }
      else if (activeMenu.value === 'PrdtType') {
        url = '/delete/prdtType';
      }
      else if (activeMenu.value === 'FoodAddType') {
        url = '/delete/foodAddType';
      }
      else if (activeMenu.value === 'FeedAddType') {
        url = '/delete/feedAddType';
      }
      else if (activeMenu.value === 'UnitMeas') {
        url = '/delete/unitMeas';
      }
      else if (activeMenu.value === 'PackType') {
        url = '/delete/packType';
      }
      else if (activeMenu.value === 'ConType') {
        url = '/delete/conType';
      }
      else if (activeMenu.value === 'Currency') {
        url = '/delete/currency';
      }
      else if (activeMenu.value === 'TradeTerm') {
        url = '/delete/tradeTerm';
      }
      else if (activeMenu.value === 'Nation') {
        url = '/delete/nation';
      }
      else if (activeMenu.value === 'Port') {
        url = '/delete/port';
      }
      else if (activeMenu.value === 'TaxType') {
        url = '/delete/taxType';
      }
      else if (activeMenu.value === 'BrandType') {
        url = '/delete/brandType';
      }
      else if (activeMenu.value === 'EduLevel') {
        url = '/delete/eduLevel';
      }
      else if (activeMenu.value === 'Dept') {
        url = '/delete/dept';
      }
      else if (activeMenu.value === 'Position') {
        url = '/delete/position';
      }
      else if (activeMenu.value === 'QualStd') {
        url = '/delete/qualStd';
      }
      else if (activeMenu.value === 'InvLoc') {
        url = '/delete/invLoc';
      }
      else if (activeMenu.value === 'DocReq') {
        url = '/delete/docReq';
      }
      else if (activeMenu.value === 'PayMth') {
        url = '/delete/payMth';
      }
      else if (activeMenu.value === 'PayLimit') {
        url = '/delete/payLimit';
      }
      else if (activeMenu.value === 'FinaDocStatus') {
        url = '/delete/finaDocStatus';
      }
      else if (activeMenu.value === 'FinaDocType') {
        url = '/delete/finaDocType';
      }
      else if (activeMenu.value === 'ExpType') {
        url = '/delete/expType';
      }
      else if (activeMenu.value === 'Rates') {
        url = '/delete/rates';
      }
      else if (activeMenu.value === 'BussOrderSta') {
        url = '/delete/bussOrderSta';
      }

      let data = {};
      if (activeMenu.value === 'customer') {
        data = { MercTypeID: row.id }; // 如果是客户类型，传递 MercTypeID
      }
      else if (activeMenu.value === 'supplier') {
        data = { SuprTypeID: row.id }; // 如果是供应商类型，传递 SuprTypeID
      }
      else if (activeMenu.value === 'PrdtType') {
        data = { PrdtTypeID: row.id }; // 如果是产品类型，传递 PrdtTypeID
      }
      else if (activeMenu.value === 'FoodAddType') {
        data = { FoodAddTypeID: row.id }; // 如果是食品添加剂类型，传递 FoodAddTypeID
      }
      else if (activeMenu.value === 'FeedAddType') {
        data = { FeedAddTypeID: row.id }; // 如果是饲料添加剂类型，传递 FeedAddTypeID
      }
      else if (activeMenu.value === 'UnitMeas') {
        data = { UnitMeasID: row.id }; // 如果是计量单位，传递 UnitMeasID
      }
      else if (activeMenu.value === 'PackType') {
        data = { PackTypeID: row.id }; // 如果是包装类型，传递 PackTypeID
      }
      else if (activeMenu.value === 'ConType') {
        data = { ConTypeID: row.id }; // 如果是集装箱类型，传递 ConTypeID
      }
      else if (activeMenu.value === 'Currency') {
        data = { CurrencyID: row.id }; // 如果是币种，传递 CurrencyID
      }
      else if (activeMenu.value === 'TradeTerm') {
        data = { TradeTermID: row.id }; // 如果是贸易条款，传递 TradeTermID
      }
      else if (activeMenu.value === 'Nation') {
        data = { NationID: row.id }; // 如果是国家，传递 NationID
      }
      else if (activeMenu.value === 'Port') {
        data = { PortID: row.id }; // 如果是港口，传递 PortID
      }
      else if (activeMenu.value === 'TaxType') {
        data = { TaxTypeID: row.id }; // 如果是税号类型，传递 TaxTypeID
      }
      else if (activeMenu.value === 'BrandType') {
        data = { BrandTypeID: row.id }; // 如果是品牌类型，传递 BrandTypeID
      }
      else if (activeMenu.value === 'EduLevel') {
        data = { EduLevelID: row.id }; // 如果是学历，传递 EduLevelID
      }
      else if (activeMenu.value === 'Dept') {
        data = { DeptID: row.id }; // 如果是公司部门，传递 DeptID
      }
      else if (activeMenu.value === 'Position') {
        data = { PositionID: row.id }; // 如果是公司岗位，传递 PositionID
      }
      else if (activeMenu.value === 'QualStd') {
        data = { QualStdID: row.id }; // 如果是质量标准，传递 QualStdID
      }
      else if (activeMenu.value === 'InvLoc') {
        data = { InvLocID: row.id }; // 如果是库存地点位置，传递 InvLocID
      }
      else if (activeMenu.value === 'DocReq') {
        data = { DocReqID: row.id }; // 如果是单据要求，传递 DocReqID
      }
      else if (activeMenu.value === 'PayMth') {
        data = { PayMthID: row.id }; // 如果是转账方式，传递 PayMthID
      }
      else if (activeMenu.value === 'PayLimit') {
        data = { PayLimitID: row.id }; // 如果是后付款转账期限，传递 PayLimitID
      }
      else if (activeMenu.value === 'FinaDocStatus') {
        data = { FinaDocStatusID: row.id }; // 如果是财务单据状态，传递 FinaDocStatusID
      }
      else if (activeMenu.value === 'FinaDocType') {
        data = { FinaDocTypeID: row.id }; // 如果是财务单据类型，传递 FinaDocTypeID
      }
      else if (activeMenu.value === 'ExpType') {
        data = { ExpTypeID: row.id }; // 如果是费用类型，传递 ExpTypeID
      }
      else if (activeMenu.value === 'Rates') {
        data = { RatesID: row.id }; // 如果是收费标准，传递 RatesID
      }
      else if (activeMenu.value === 'BussOrderSta') {
        data = { BussOrderStaID: row.id }; // 如果是业务单据状态，传递 BussOrderStaID
      }
      await axios.post(url, data);
      // await axios.post(url, { MercTypeID: row.id, SuprTypeID: row.id });
      fetchData();
    };


    const handleSubmit = async () => {
      try {
        let url;
        let data;
        if (activeMenu.value === 'customer') {
          url = '/save/mercType';
          if (isEdit.value) {
            data = { MercType: form.value.type, MercTypeID: form.value.id }; // 编辑时传递 MercType 和 MercTypeID
          } else {
            data = { MercType: form.value.type }; // 添加时只传递 MercType
          }
        }
        else if (activeMenu.value === 'supplier') {
          url = '/save/suprType';
          if (isEdit.value) {
            data = { SuprType: form.value.type, SuprTypeID: form.value.id }; // 编辑时传递 SuprType 和 SuprTypeID
          } else {
            data = { SuprType: form.value.type }; // 添加时只传递 SuprType
          }
        }

        else if (activeMenu.value === 'PrdtType') {
          url = '/save/prdtType';
          if (isEdit.value) {
            data = { PrdtType: form.value.type, PrdtTypeID: form.value.id }; // 编辑时传递 PrdtType 和 PrdtTypeID
          } else {
            data = { PrdtType: form.value.type }; // 添加时只传递 PrdtType
          }
        }
        else if (activeMenu.value === 'FoodAddType') {
          url = '/save/foodAddType';
          if (isEdit.value) {
            data = { FoodAddType: form.value.type, FoodAddTypeID: form.value.id }; // 编辑时传递 FoodAddType 和 FoodAddTypeID
          } else {
            data = { FoodAddType: form.value.type }; // 添加时只传递 FoodAddType
          }
        }
        else if (activeMenu.value === 'FeedAddType') {
          url = '/save/feedAddType';
          if (isEdit.value) {
            data = { FeedAddType: form.value.type, FeedAddTypeID: form.value.id }; // 编辑时传递 FeedAddType 和 FeedAddTypeID
          } else {
            data = { FeedAddType: form.value.type }; // 添加时只传递 FeedAddType
          }
        }
        else if (activeMenu.value === 'UnitMeas') {
          url = '/save/unitMeas';
          if (isEdit.value) {
            data = { UnitMeas: form.value.type, UnitMeasID: form.value.id }; // 编辑时传递 UnitMeas 和 UnitMeasID
          } else {
            data = { UnitMeas: form.value.type }; // 添加时只传递 UnitMeas
          }
        }
        else if (activeMenu.value === 'PackType') {
          url = '/save/packType';
          if (isEdit.value) {
            data = { PackType: form.value.type, PackTypeID: form.value.id }; // 编辑时传递 PackType 和 PackTypeID
          } else {
            data = { PackType: form.value.type }; // 添加时只传递 PackType
          }
        }
        else if (activeMenu.value === 'ConType') {
          url = '/save/conType';
          if (isEdit.value) {
            data = { ConType: form.value.type, ConTypeID: form.value.id }; // 编辑时传递 ConType 和 ConTypeID
          } else {
            data = { ConType: form.value.type }; // 添加时只传递 ConType
          }
        }
        else if (activeMenu.value === 'Currency') {
          url = '/save/currency';
          if (isEdit.value) {
            data = { Currency: form.value.type, CurrencyID: form.value.id }; // 编辑时传递 Currency 和 CurrencyID
          } else {
            data = { Currency: form.value.type }; // 添加时只传递 Currency
          }
        }
        else if (activeMenu.value === 'TradeTerm') {
          url = '/save/tradeTerm';
          if (isEdit.value) {
            data = { TradeTerm: form.value.type, TradeTermID: form.value.id }; // 编辑时传递 TradeTerm 和 TradeTermID
          } else {
            data = { TradeTerm: form.value.type }; // 添加时只传递 TradeTerm
          }
        }
        else if (activeMenu.value === 'Nation') {
          url = '/save/nation';
          if (isEdit.value) {
            data = { Nation: form.value.type, NationID: form.value.id }; // 编辑时传递 Nation 和 NationID
          } else {
            data = { Nation: form.value.type }; // 添加时只传递 Nation
          }
        }
        else if (activeMenu.value === 'Port') {
          url = '/save/port';
          if (isEdit.value) {
            data = { Port: form.value.type, PortID: form.value.id }; // 编辑时传递 Port 和 PortID
          } else {
            data = { Port: form.value.type }; // 添加时只传递 Port
          }
        }
        else if (activeMenu.value === 'TaxType') {
          url = '/save/taxType';
          if (isEdit.value) {
            data = { TaxType: form.value.type, TaxTypeID: form.value.id }; // 编辑时传递 TaxType 和 TaxTypeID
          } else {
            data = { TaxType: form.value.type }; // 添加时只传递 TaxType
          }
        }
        else if (activeMenu.value === 'BrandType') {
          url = '/save/brandType';
          if (isEdit.value) {
            data = { BrandType: form.value.type, BrandTypeID: form.value.id }; // 编辑时传递 BrandType 和 BrandTypeID
          } else {
            data = { BrandType: form.value.type }; // 添加时只传递 BrandType
          }
        }
        else if (activeMenu.value === 'EduLevel') {
          url = '/save/eduLevel';
          if (isEdit.value) {
            data = { EduLevel: form.value.type, EduLevelID: form.value.id }; // 编辑时传递 EduLevel 和 EduLevelID
          } else {
            data = { EduLevel: form.value.type }; // 添加时只传递 EduLevel
          }
        }
        else if (activeMenu.value === 'Dept') {
          url = '/save/dept';
          if (isEdit.value) {
            data = { Dept: form.value.type, DeptID: form.value.id }; // 编辑时传递 Dept 和 DeptID
          } else {
            data = { Dept: form.value.type }; // 添加时只传递 Dept
          }
        }
        else if (activeMenu.value === 'Position') {
          url = '/save/position';
          if (isEdit.value) {
            data = { Position: form.value.type, PositionID: form.value.id }; // 编辑时传递 Position 和 PositionID
          } else {
            data = { Position: form.value.type }; // 添加时只传递 Position
          }
        }
        else if (activeMenu.value === 'QualStd') {
          url = '/save/qualStd';
          if (isEdit.value) {
            data = { QualStd: form.value.type, QualStdID: form.value.id }; // 编辑时传递 QualStd 和 QualStdID
          } else {
            data = { QualStd: form.value.type }; // 添加时只传递 QualStd
          }
        }
        else if (activeMenu.value === 'InvLoc') {
          url = '/save/invLoc';
          if (isEdit.value) {
            data = { InvLoc: form.value.type, InvLocID: form.value.id }; // 编辑时传递 InvLoc 和 InvLocID
          } else {
            data = { InvLoc: form.value.type }; // 添加时只传递 InvLoc
          }
        }
        else if (activeMenu.value === 'DocReq') {
          url = '/save/docReq';
          if (isEdit.value) {
            data = { DocReq: form.value.type, DocReqID: form.value.id }; // 编辑时传递 DocReq 和 DocReqID
          } else {
            data = { DocReq: form.value.type }; // 添加时只传递 DocReq
          }
        }
        else if (activeMenu.value === 'PayMth') {
          url = '/save/payMth';
          if (isEdit.value) {
            data = { PayMth: form.value.type, PayMthID: form.value.id }; // 编辑时传递 PayMth 和 PayMthID
          } else {
            data = { PayMth: form.value.type }; // 添加时只传递 PayMth
          }
        }
        else if (activeMenu.value === 'PayLimit') {
          url = '/save/payLimit';
          if (isEdit.value) {
            data = { PayLimit: form.value.type, PayLimitID: form.value.id }; // 编辑时传递 PayLimit 和 PayLimitID
          } else {
            data = { PayLimit: form.value.type }; // 添加时只传递 PayLimit
          }
        }
        else if (activeMenu.value === 'FinaDocStatus') {
          url = '/save/finaDocStatus';
          if (isEdit.value) {
            data = { FinaDocStatus: form.value.type, FinaDocStatusID: form.value.id }; // 编辑时传递 FinaDocStatus 和 FinaDocStatusID
          } else {
            data = { FinaDocStatus: form.value.type }; // 添加时只传递 FinaDocStatus
          }
        }
        else if (activeMenu.value === 'FinaDocType') {
          url = '/save/finaDocType';
          if (isEdit.value) {
            data = { FinaDocType: form.value.type, FinaDocTypeID: form.value.id }; // 编辑时传递 FinaDocType 和 FinaDocTypeID
          } else {
            data = { FinaDocType: form.value.type }; // 添加时只传递 FinaDocType
          }
        }
        else if (activeMenu.value === 'ExpType') {
          url = '/save/expType';
          if (isEdit.value) {
            data = { ExpType: form.value.type, ExpTypeID: form.value.id }; // 编辑时传递 ExpType 和 ExpTypeID
          } else {
            data = { ExpType: form.value.type }; // 添加时只传递 ExpType
          }
        }
        else if (activeMenu.value === 'Rates') {
          url = '/save/rates';
          if (isEdit.value) {
            data = { Rates: form.value.type, RatesID: form.value.id }; // 编辑时传递 Rates 和 RatesID
          } else {
            data = { Rates: form.value.type }; // 添加时只传递 Rates
          }
        }
        else if (activeMenu.value === 'BussOrderSta') {
          url = '/save/bussOrderSta';
          if (isEdit.value) {
            data = { BussOrderSta: form.value.type, BussOrderStaID: form.value.id }; // 编辑时传递 BussOrderSta 和 BussOrderStaID
          } else {
            data = { BussOrderSta: form.value.type }; // 添加时只传递 BussOrderSta
          }
        }
        // 调用保存接口
        const response = await axios.post(url, data);

        // 检查响应状态码，确保请求成功
        if (response.status === 200) {
          dialogVisible.value = false; // 关闭对话框
          fetchData(); // 重新获取数据
        } else {
          // 处理其他非200状态码的情况
          throw new Error(`请求失败，状态码：${response.status}`);
        }
      } catch (error) {
        // 捕获并处理异常
        console.error('提交数据时发生错误:', error);

        ElMessage.error(response.data.RetMessage || '删除失败');
        // 可以根据不同的错误类型进行不同的处理
        if (error.response) {
          // 请求已发出，但服务器响应状态码不在 2xx 范围内
          console.error('服务器返回错误:', error.response.data);
        } else if (error.request) {
          // 请求已发出，但没有收到响应
          console.error('未收到服务器响应:', error.request);
        } else {
          // 其他错误
          console.error('发生错误:', error.message);
        }

        // 可以在这里显示错误提示给用户
        alert('提交数据失败，请稍后重试。');
      }
    };
    // 组件挂载时获取数据
    onMounted(() => {
      fetchData();
    });

    return {
      activeMenu,
      tableData,
      dialogVisible,
      dialogTitle,
      form,
      isEdit,
      handleMenuSelect,
      handleAdd,
      handleEdit,
      handleDelete,
      handleSubmit,
      handleBack,
    };
  },
};
</script>


<style src="../assets/styles/DictionaryManager.css" scoped></style>

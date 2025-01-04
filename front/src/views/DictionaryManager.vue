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
        <el-menu-item index="Postion">公司岗位</el-menu-item>
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
      else if (activeMenu.value === 'Postion') {
        url = '/find/postion';
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
          id: item.MercTypeId,
        }));
      }
      else if (activeMenu.value === 'supplier') {
        tableData.value = response.data.SuprType.map(item => ({
          type: item.SuprType,
          id: item.SuprTypeId,
        }));
      }

      else if (activeMenu.value === 'PrdtType') {
        tableData.value = response.data.PrdtType.map(item => ({
          type: item.PrdtType,
          id: item.PrdtTypeId,
        }));
      }
      else if (activeMenu.value === 'FoodAddType') {
        tableData.value = response.data.FoodAddType.map(item => ({
          type: item.FoodAddType,
          id: item.FoodAddTypeId,
        }));
      }
      else if (activeMenu.value === 'FeedAddType') {
        tableData.value = response.data.FeedAddType.map(item => ({
          type: item.FeedAddType,
          id: item.FeedAddTypeId,
        }));
      }
      else if (activeMenu.value === 'UnitMeas') {
        tableData.value = response.data.UnitMeas.map(item => ({
          type: item.UnitMeas,
          id: item.UnitMeasId,
        }));
      }
      else if (activeMenu.value === 'PackType') {
        tableData.value = response.data.PackType.map(item => ({
          type: item.PackType,
          id: item.PackTypeId,
        }));
      }
      else if (activeMenu.value === 'ConType') {
        tableData.value = response.data.ConType.map(item => ({
          type: item.ConType,
          id: item.ConTypeId,
        }));
      }
      else if (activeMenu.value === 'Currency') {
        tableData.value = response.data.Currency.map(item => ({
          type: item.Currency,
          id: item.CurrencyId,
        }));
      }
      else if (activeMenu.value === 'TradeTerm') {
        tableData.value = response.data.TradeTerm.map(item => ({
          type: item.TradeTerm,
          id: item.TradeTermId,
        }));
      }
      else if (activeMenu.value === 'Nation') {
        tableData.value = response.data.Nation.map(item => ({
          type: item.Nation,
          id: item.NationId,
        }));
      }
      else if (activeMenu.value === 'Port') {
        tableData.value = response.data.Port.map(item => ({
          type: item.Port,
          id: item.PortId,
        }));
      }
      else if (activeMenu.value === 'TaxType') {
        tableData.value = response.data.TaxType.map(item => ({
          type: item.TaxType,
          id: item.TaxTypeId,
        }));
      }
      else if (activeMenu.value === 'BrandType') {
        tableData.value = response.data.BrandType.map(item => ({
          type: item.BrandType,
          id: item.BrandTypeId,
        }));
      }
      else if (activeMenu.value === 'EduLevel') {
        tableData.value = response.data.EduLevel.map(item => ({
          type: item.EduLevel,
          id: item.EduLevelId,
        }));
      }
      else if (activeMenu.value === 'Dept') {
        tableData.value = response.data.Dept.map(item => ({
          type: item.Dept,
          id: item.DeptId,
        }));
      }
      else if (activeMenu.value === 'Postion') {
        tableData.value = response.data.Postion.map(item => ({
          type: item.Postion,
          id: item.PostionId,
        }));
      }
      else if (activeMenu.value === 'QualStd') {
        tableData.value = response.data.QualStd.map(item => ({
          type: item.QualStd,
          id: item.QualStdId,
        }));
      }
      else if (activeMenu.value === 'InvLoc') {
        tableData.value = response.data.InvLoc.map(item => ({
          type: item.InvLoc,
          id: item.InvLocId,
        }));
      }
      else if (activeMenu.value === 'DocReq') {
        tableData.value = response.data.DocReq.map(item => ({
          type: item.DocReq,
          id: item.DocReqId,
        }));
      }
      else if (activeMenu.value === 'PayMth') {
        tableData.value = response.data.PayMth.map(item => ({
          type: item.PayMth,
          id: item.PayMthId,
        }));
      }
      else if (activeMenu.value === 'PayLimit') {
        tableData.value = response.data.PayLimit.map(item => ({
          type: item.PayLimit,
          id: item.PayLimitId,
        }));
      }
      else if (activeMenu.value === 'FinaDocStatus') {
        tableData.value = response.data.FinaDocStatus.map(item => ({
          type: item.FinaDocStatus,
          id: item.FinaDocStatusId,
        }));
      }
      else if (activeMenu.value === 'FinaDocType') {
        tableData.value = response.data.FinaDocType.map(item => ({
          type: item.FinaDocType,
          id: item.FinaDocTypeId,
        }));
      }
      else if (activeMenu.value === 'ExpType') {
        tableData.value = response.data.ExpType.map(item => ({
          type: item.ExpType,
          id: item.ExpTypeId,
        }));
      }
      else if (activeMenu.value === 'Rates') {
        tableData.value = response.data.Rates.map(item => ({
          type: item.Rates,
          id: item.RatesId,
        }));
      }
      else if (activeMenu.value === 'BussOrderSta') {
        tableData.value = response.data.BussOrderSta.map(item => ({
          type: item.BussOrderSta,
          id: item.BussOrderStaId,
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
      else if (activeMenu.value === 'Postion') {
        url = '/delete/postion';
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
        data = { MercTypeId: row.id }; // 如果是客户类型，传递 MercTypeId
      }
      else if (activeMenu.value === 'supplier') {
        data = { SuprTypeId: row.id }; // 如果是供应商类型，传递 SuprTypeId
      }
      else if (activeMenu.value === 'PrdtType') {
        data = { PrdtTypeId: row.id }; // 如果是产品类型，传递 PrdtTypeId
      }
      else if (activeMenu.value === 'FoodAddType') {
        data = { FoodAddTypeId: row.id }; // 如果是食品添加剂类型，传递 FoodAddTypeId
      }
      else if (activeMenu.value === 'FeedAddType') {
        data = { FeedAddTypeId: row.id }; // 如果是饲料添加剂类型，传递 FeedAddTypeId
      }
      else if (activeMenu.value === 'UnitMeas') {
        data = { UnitMeasId: row.id }; // 如果是计量单位，传递 UnitMeasId
      }
      else if (activeMenu.value === 'PackType') {
        data = { PackTypeId: row.id }; // 如果是包装类型，传递 PackTypeId
      }
      else if (activeMenu.value === 'ConType') {
        data = { ConTypeId: row.id }; // 如果是集装箱类型，传递 ConTypeId
      }
      else if (activeMenu.value === 'Currency') {
        data = { CurrencyId: row.id }; // 如果是币种，传递 CurrencyId
      }
      else if (activeMenu.value === 'TradeTerm') {
        data = { TradeTermId: row.id }; // 如果是贸易条款，传递 TradeTermId
      }
      else if (activeMenu.value === 'Nation') {
        data = { NationId: row.id }; // 如果是国家，传递 NationId
      }
      else if (activeMenu.value === 'Port') {
        data = { PortId: row.id }; // 如果是港口，传递 PortId
      }
      else if (activeMenu.value === 'TaxType') {
        data = { TaxTypeId: row.id }; // 如果是税号类型，传递 TaxTypeId
      }
      else if (activeMenu.value === 'BrandType') {
        data = { BrandTypeId: row.id }; // 如果是品牌类型，传递 BrandTypeId
      }
      else if (activeMenu.value === 'EduLevel') {
        data = { EduLevelId: row.id }; // 如果是学历，传递 EduLevelId
      }
      else if (activeMenu.value === 'Dept') {
        data = { DeptId: row.id }; // 如果是公司部门，传递 DeptId
      }
      else if (activeMenu.value === 'Postion') {
        data = { PostionId: row.id }; // 如果是公司岗位，传递 PostionId
      }
      else if (activeMenu.value === 'QualStd') {
        data = { QualStdId: row.id }; // 如果是质量标准，传递 QualStdId
      }
      else if (activeMenu.value === 'InvLoc') {
        data = { InvLocId: row.id }; // 如果是库存地点位置，传递 InvLocId
      }
      else if (activeMenu.value === 'DocReq') {
        data = { DocReqId: row.id }; // 如果是单据要求，传递 DocReqId
      }
      else if (activeMenu.value === 'PayMth') {
        data = { PayMthId: row.id }; // 如果是转账方式，传递 PayMthId
      }
      else if (activeMenu.value === 'PayLimit') {
        data = { PayLimitId: row.id }; // 如果是后付款转账期限，传递 PayLimitId
      }
      else if (activeMenu.value === 'FinaDocStatus') {
        data = { FinaDocStatusId: row.id }; // 如果是财务单据状态，传递 FinaDocStatusId
      }
      else if (activeMenu.value === 'FinaDocType') {
        data = { FinaDocTypeId: row.id }; // 如果是财务单据类型，传递 FinaDocTypeId
      }
      else if (activeMenu.value === 'ExpType') {
        data = { ExpTypeId: row.id }; // 如果是费用类型，传递 ExpTypeId
      }
      else if (activeMenu.value === 'Rates') {
        data = { RatesId: row.id }; // 如果是收费标准，传递 RatesId
      }
      else if (activeMenu.value === 'BussOrderSta') {
        data = { BussOrderStaId: row.id }; // 如果是业务单据状态，传递 BussOrderStaId
      }
      await axios.post(url, data);
      // await axios.post(url, { MercTypeId: row.id, SuprTypeId: row.id });
      fetchData();
    };


    const handleSubmit = async () => {
      try {
        let url;
        let data;
        if (activeMenu.value === 'customer') {
          url = '/save/mercType';
          if (isEdit.value) {
            data = { MercType: form.value.type, MercTypeId: form.value.id }; // 编辑时传递 MercType 和 MercTypeId
          } else {
            data = { MercType: form.value.type }; // 添加时只传递 MercType
          }
        }
        else if (activeMenu.value === 'supplier') {
          url = '/save/suprType';
          if (isEdit.value) {
            data = { SuprType: form.value.type, SuprTypeId: form.value.id }; // 编辑时传递 SuprType 和 SuprTypeId
          } else {
            data = { SuprType: form.value.type }; // 添加时只传递 SuprType
          }
        }

        else if (activeMenu.value === 'PrdtType') {
          url = '/save/prdtType';
          if (isEdit.value) {
            data = { PrdtType: form.value.type, PrdtTypeId: form.value.id }; // 编辑时传递 PrdtType 和 PrdtTypeId
          } else {
            data = { PrdtType: form.value.type }; // 添加时只传递 PrdtType
          }
        }
        else if (activeMenu.value === 'FoodAddType') {
          url = '/save/foodAddType';
          if (isEdit.value) {
            data = { FoodAddType: form.value.type, FoodAddTypeId: form.value.id }; // 编辑时传递 FoodAddType 和 FoodAddTypeId
          } else {
            data = { FoodAddType: form.value.type }; // 添加时只传递 FoodAddType
          }
        }
        else if (activeMenu.value === 'FeedAddType') {
          url = '/save/feedAddType';
          if (isEdit.value) {
            data = { FeedAddType: form.value.type, FeedAddTypeId: form.value.id }; // 编辑时传递 FeedAddType 和 FeedAddTypeId
          } else {
            data = { FeedAddType: form.value.type }; // 添加时只传递 FeedAddType
          }
        }
        else if (activeMenu.value === 'UnitMeas') {
          url = '/save/unitMeas';
          if (isEdit.value) {
            data = { UnitMeas: form.value.type, UnitMeasId: form.value.id }; // 编辑时传递 UnitMeas 和 UnitMeasId
          } else {
            data = { UnitMeas: form.value.type }; // 添加时只传递 UnitMeas
          }
        }
        else if (activeMenu.value === 'PackType') {
          url = '/save/packType';
          if (isEdit.value) {
            data = { PackType: form.value.type, PackTypeId: form.value.id }; // 编辑时传递 PackType 和 PackTypeId
          } else {
            data = { PackType: form.value.type }; // 添加时只传递 PackType
          }
        }
        else if (activeMenu.value === 'ConType') {
          url = '/save/conType';
          if (isEdit.value) {
            data = { ConType: form.value.type, ConTypeId: form.value.id }; // 编辑时传递 ConType 和 ConTypeId
          } else {
            data = { ConType: form.value.type }; // 添加时只传递 ConType
          }
        }
        else if (activeMenu.value === 'Currency') {
          url = '/save/currency';
          if (isEdit.value) {
            data = { Currency: form.value.type, CurrencyId: form.value.id }; // 编辑时传递 Currency 和 CurrencyId
          } else {
            data = { Currency: form.value.type }; // 添加时只传递 Currency
          }
        }
        else if (activeMenu.value === 'TradeTerm') {
          url = '/save/tradeTerm';
          if (isEdit.value) {
            data = { TradeTerm: form.value.type, TradeTermId: form.value.id }; // 编辑时传递 TradeTerm 和 TradeTermId
          } else {
            data = { TradeTerm: form.value.type }; // 添加时只传递 TradeTerm
          }
        }
        else if (activeMenu.value === 'Nation') {
          url = '/save/nation';
          if (isEdit.value) {
            data = { Nation: form.value.type, NationId: form.value.id }; // 编辑时传递 Nation 和 NationId
          } else {
            data = { Nation: form.value.type }; // 添加时只传递 Nation
          }
        }
        else if (activeMenu.value === 'Port') {
          url = '/save/port';
          if (isEdit.value) {
            data = { Port: form.value.type, PortId: form.value.id }; // 编辑时传递 Port 和 PortId
          } else {
            data = { Port: form.value.type }; // 添加时只传递 Port
          }
        }
        else if (activeMenu.value === 'TaxType') {
          url = '/save/taxType';
          if (isEdit.value) {
            data = { TaxType: form.value.type, TaxTypeId: form.value.id }; // 编辑时传递 TaxType 和 TaxTypeId
          } else {
            data = { TaxType: form.value.type }; // 添加时只传递 TaxType
          }
        }
        else if (activeMenu.value === 'BrandType') {
          url = '/save/brandType';
          if (isEdit.value) {
            data = { BrandType: form.value.type, BrandTypeId: form.value.id }; // 编辑时传递 BrandType 和 BrandTypeId
          } else {
            data = { BrandType: form.value.type }; // 添加时只传递 BrandType
          }
        }
        else if (activeMenu.value === 'EduLevel') {
          url = '/save/eduLevel';
          if (isEdit.value) {
            data = { EduLevel: form.value.type, EduLevelId: form.value.id }; // 编辑时传递 EduLevel 和 EduLevelId
          } else {
            data = { EduLevel: form.value.type }; // 添加时只传递 EduLevel
          }
        }
        else if (activeMenu.value === 'Dept') {
          url = '/save/dept';
          if (isEdit.value) {
            data = { Dept: form.value.type, DeptId: form.value.id }; // 编辑时传递 Dept 和 DeptId
          } else {
            data = { Dept: form.value.type }; // 添加时只传递 Dept
          }
        }
        else if (activeMenu.value === 'Postion') {
          url = '/save/postion';
          if (isEdit.value) {
            data = { Postion: form.value.type, PostionId: form.value.id }; // 编辑时传递 Postion 和 PostionId
          } else {
            data = { Postion: form.value.type }; // 添加时只传递 Postion
          }
        }
        else if (activeMenu.value === 'QualStd') {
          url = '/save/qualStd';
          if (isEdit.value) {
            data = { QualStd: form.value.type, QualStdId: form.value.id }; // 编辑时传递 QualStd 和 QualStdId
          } else {
            data = { QualStd: form.value.type }; // 添加时只传递 QualStd
          }
        }
        else if (activeMenu.value === 'InvLoc') {
          url = '/save/invLoc';
          if (isEdit.value) {
            data = { InvLoc: form.value.type, InvLocId: form.value.id }; // 编辑时传递 InvLoc 和 InvLocId
          } else {
            data = { InvLoc: form.value.type }; // 添加时只传递 InvLoc
          }
        }
        else if (activeMenu.value === 'DocReq') {
          url = '/save/docReq';
          if (isEdit.value) {
            data = { DocReq: form.value.type, DocReqId: form.value.id }; // 编辑时传递 DocReq 和 DocReqId
          } else {
            data = { DocReq: form.value.type }; // 添加时只传递 DocReq
          }
        }
        else if (activeMenu.value === 'PayMth') {
          url = '/save/payMth';
          if (isEdit.value) {
            data = { PayMth: form.value.type, PayMthId: form.value.id }; // 编辑时传递 PayMth 和 PayMthId
          } else {
            data = { PayMth: form.value.type }; // 添加时只传递 PayMth
          }
        }
        else if (activeMenu.value === 'PayLimit') {
          url = '/save/payLimit';
          if (isEdit.value) {
            data = { PayLimit: form.value.type, PayLimitId: form.value.id }; // 编辑时传递 PayLimit 和 PayLimitId
          } else {
            data = { PayLimit: form.value.type }; // 添加时只传递 PayLimit
          }
        }
        else if (activeMenu.value === 'FinaDocStatus') {
          url = '/save/finaDocStatus';
          if (isEdit.value) {
            data = { FinaDocStatus: form.value.type, FinaDocStatusId: form.value.id }; // 编辑时传递 FinaDocStatus 和 FinaDocStatusId
          } else {
            data = { FinaDocStatus: form.value.type }; // 添加时只传递 FinaDocStatus
          }
        }
        else if (activeMenu.value === 'FinaDocType') {
          url = '/save/finaDocType';
          if (isEdit.value) {
            data = { FinaDocType: form.value.type, FinaDocTypeId: form.value.id }; // 编辑时传递 FinaDocType 和 FinaDocTypeId
          } else {
            data = { FinaDocType: form.value.type }; // 添加时只传递 FinaDocType
          }
        }
        else if (activeMenu.value === 'ExpType') {
          url = '/save/expType';
          if (isEdit.value) {
            data = { ExpType: form.value.type, ExpTypeId: form.value.id }; // 编辑时传递 ExpType 和 ExpTypeId
          } else {
            data = { ExpType: form.value.type }; // 添加时只传递 ExpType
          }
        }
        else if (activeMenu.value === 'Rates') {
          url = '/save/rates';
          if (isEdit.value) {
            data = { Rates: form.value.type, RatesId: form.value.id }; // 编辑时传递 Rates 和 RatesId
          } else {
            data = { Rates: form.value.type }; // 添加时只传递 Rates
          }
        }
        else if (activeMenu.value === 'BussOrderSta') {
          url = '/save/bussOrderSta';
          if (isEdit.value) {
            data = { BussOrderSta: form.value.type, BussOrderStaId: form.value.id }; // 编辑时传递 BussOrderSta 和 BussOrderStaId
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
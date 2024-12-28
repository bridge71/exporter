<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <el-menu default-active="1" class="el-menu-vertical-demo">
          <el-submenu index="1">
            <template #title>会计实体信息</template>
            <el-menu-item index="1-1" @click="showAcctDialog = true">会计实体信息</el-menu-item>
            <el-menu-item index="1-2" @click="showBankDialog = true">会计实体银行账户信息</el-menu-item>
            <el-menu-item index="1-3">客商信息</el-menu-item>
            <!-- 其他菜单项 -->
          </el-submenu>
          <!-- 其他菜单 -->
        </el-menu>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>
        <el-header style="display: flex; justify-content: space-between; align-items: center;">
          <h2>会计实体信息</h2>
          <el-button type="primary" @click="showAcctDialog = true">添加会计实体信息</el-button>
        </el-header>
        <el-main>
          <!-- 会计实体信息表格 -->
          <el-table :data="acctData" style="width: 100%">
            <el-table-column prop="AcctCode" label="会计实体编码"></el-table-column>
            <el-table-column prop="AcctAbbr" label="会计实体缩写"></el-table-column>
            <el-table-column prop="EtyAbbr" label="实体简称"></el-table-column>
            <el-table-column prop="AcctName" label="会计实体名称"></el-table-column>
            <el-table-column prop="AcctEngName" label="会计实体英文名称"></el-table-column>
            <el-table-column prop="AcctAddr" label="会计实体地址"></el-table-column>
            <el-table-column prop="Nation" label="国家"></el-table-column>
            <el-table-column prop="TaxType" label="税号类型"></el-table-column>
            <el-table-column prop="TaxCode" label="税号"></el-table-column>
            <el-table-column prop="PhoneNum" label="联系电话"></el-table-column>
            <el-table-column prop="Email" label="邮箱"></el-table-column>
            <el-table-column prop="Website" label="网站"></el-table-column>
            <el-table-column prop="RegDate" label="注册日期"></el-table-column>
            <el-table-column prop="Notes" label="备注"></el-table-column>
            <el-table-column label="操作">
              <template #default="scope">
                <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加会计实体信息的对话框 -->
    <el-dialog v-model="showAcctDialog" title="会计实体信息" width="50%">
      <el-form :model="acctForm" label-width="150px" :rules="acctRules" ref="acctFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="会计实体编码" prop="AcctCode" :required="true">
              <el-input v-model="acctForm.AcctCode" placeholder="国家代码+流水码"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体缩写" prop="AcctAbbr" :required="true">
              <el-input v-model="acctForm.AcctAbbr" placeholder="名称缩写，一般2-3个字母"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="实体简称" prop="EtyAbbr" :required="true">
              <el-input v-model="acctForm.EtyAbbr" placeholder="客户名称简称，一般1-2个单词"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体名称" prop="AcctName">
              <el-input v-model="acctForm.AcctName" placeholder="全称"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="会计实体英文名称" prop="AcctEngName">
              <el-input v-model="acctForm.AcctEngName" placeholder="英文全称（仅限注册在中国大陆的会计实体）"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体地址" prop="AcctAddr">
              <el-input v-model="acctForm.AcctAddr"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="acctForm.Nation"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-select v-model="acctForm.TaxType" placeholder="请选择">
                <el-option label="类型1" value="type1"></el-option>
                <el-option label="类型2" value="type2"></el-option>
                <!-- 添加更多选项 -->
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="税号" prop="TaxCode">
              <el-input v-model="acctForm.TaxCode"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="acctForm.PhoneNum"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="acctForm.Email"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="网站" prop="Website">
              <el-input v-model="acctForm.Website"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册日期" prop="RegDate">
              <el-date-picker v-model="acctForm.RegDate" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="acctForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="营业执照" prop="License">
              <el-upload action="https://jsonplaceholder.typicode.com/posts/" :limit="1">
                <el-button type="primary">上传文件</el-button>
              </el-upload>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 会计实体银行账户信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联银行账户信息">
              <el-select v-model="acctForm.BankAccounts" multiple placeholder="请选择银行账户信息">
                <el-option v-for="bank in bankData" :key="bank.AccNum" :label="`${bank.AccName} (${bank.AccNum})`"
                  :value="bank.AccNum"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitAcctForm">保存</el-button>
            <el-button @click="showAcctDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 添加会计实体银行账户信息的对话框 -->
    <el-dialog v-model="showBankDialog" title="会计实体银行账户信息" width="50%">
      <el-form :model="bankForm" label-width="150px" :rules="bankRules" ref="bankFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账户名称" prop="AccName">
              <el-input v-model="bankForm.AccName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="账号" prop="AccNum">
              <el-input v-model="bankForm.AccNum"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-input v-model="bankForm.Currency"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="开户行名称" prop="BankName">
              <el-input v-model="bankForm.BankName"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行号" prop="BankNum">
              <el-input v-model="bankForm.BankNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="SWIFT CODE" prop="SwiftCode">
              <el-input v-model="bankForm.SwiftCode"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行地址" prop="BankAddr">
              <el-input v-model="bankForm.BankAddr"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="bankForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="附件" prop="File">
              <el-upload action="https://jsonplaceholder.typicode.com/posts/" :limit="1">
                <el-button type="primary">上传文件</el-button>
              </el-upload>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 会计实体信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联会计实体信息">
              <el-select v-model="bankForm.AcctCode" placeholder="请选择会计实体信息">
                <el-option v-for="acct in acctData" :key="acct.AcctCode" :label="`${acct.AcctName} (${acct.AcctCode})`"
                  :value="acct.AcctCode"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitBankForm">保存</el-button>
            <el-button @click="showBankDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue';

// 会计实体信息对话框显示状态
const showAcctDialog = ref(false);

// 会计实体银行账户信息对话框显示状态
const showBankDialog = ref(false);

// 会计实体信息表单数据
const acctForm = ref({
  AcctCode: '',
  AcctAbbr: '',
  EtyAbbr: '',
  AcctName: '',
  AcctEngName: '',
  AcctAddr: '',
  Nation: '',
  TaxType: '',
  TaxCode: '',
  PhoneNum: '',
  Email: '',
  Website: '',
  RegDate: '',
  Notes: '',
  License: '',
  BankAccounts: [] // 关联的银行账户信息
});

// 会计实体银行账户信息表单数据
const bankForm = ref({
  AccName: '',
  AccNum: '',
  Currency: '',
  BankName: '',
  BankNum: '',
  SwiftCode: '',
  BankAddr: '',
  Notes: '',
  File: '',
  AcctCode: '' // 关联的会计实体信息
});

// 会计实体信息表单验证规则
const acctRules = {
  AcctCode: [{ required: true, message: '请输入会计实体编码', trigger: 'blur' }],
  AcctAbbr: [{ required: true, message: '请输入会计实体缩写', trigger: 'blur' }],
  EtyAbbr: [{ required: true, message: '请输入实体简称', trigger: 'blur' }]
};

// 会计实体银行账户信息表单验证规则
const bankRules = {
  AccName: [{ required: true, message: '请输入账户名称', trigger: 'blur' }],
  AccNum: [{ required: true, message: '请输入账号', trigger: 'blur' }]
};

// 会计实体信息表单提交逻辑
const submitAcctForm = () => {
  console.log('会计实体信息:', acctForm.value);
  showAcctDialog.value = false;
};

// 会计实体银行账户信息表单提交逻辑
const submitBankForm = () => {
  console.log('会计实体银行账户信息:', bankForm.value);
  showBankDialog.value = false;
};

// 表格数据（示例数据）
const acctData = ref([
  {
    AcctCode: 'HKG0001',
    AcctAbbr: 'LFHK',
    EtyAbbr: 'Lonford',
    AcctName: 'Lonford International',
    AcctEngName: 'Lonford International',
    AcctAddr: 'Hong Kong',
    Nation: 'China',
    TaxType: 'VAT',
    TaxCode: '123456',
    PhoneNum: '123-456-7890',
    Email: 'info@lonford.com',
    Website: 'www.lonford.com',
    RegDate: '2023-01-01',
    Notes: '示例数据',
    BankAccounts: [] // 关联的银行账户信息
  }
]);

// 银行账户数据（示例数据）
const bankData = ref([
  {
    AccName: '账户1',
    AccNum: '1234567890',
    Currency: 'CNY',
    BankName: '中国银行',
    BankNum: '123456',
    SwiftCode: 'BKCHCNBJ',
    BankAddr: '北京市',
    Notes: '示例数据',
    AcctCode: 'HKG0001' // 关联的会计实体信息
  }
]);

// 操作按钮逻辑
const handleView = (index, row) => {
  console.log('查看', index, row);
};

const handleEdit = (index, row) => {
  console.log('编辑', index, row);
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  margin-top: 60px;
}

.el-header {
  background-color: #f5f5f5;
  padding: 0 20px;
}

.el-aside {
  background-color: #304156;
}

.el-menu {
  border-right: none;
}

.el-menu-item,
.el-submenu__title {
  color: #fff;
}

.el-menu-item:hover,
.el-submenu__title:hover {
  background-color: #263445;
}

.el-main {
  padding: 20px;
}
</style>

<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <el-menu :default-active="activeMenu" class="el-menu-vertical-demo" @select="handleMenuSelect">
          <el-submenu index="1">
            <template #title>会计实体信息</template>
            <el-menu-item index="1-1">会计实体信息</el-menu-item>
            <el-menu-item index="1-2">会计实体银行账户信息</el-menu-item>
            <el-menu-item index="1-3">客商信息</el-menu-item>
            <!-- 其他菜单项 -->
          </el-submenu>
          <!-- 其他菜单 -->
        </el-menu>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>
        <el-header style="display: flex; justify-content: space-between; align-items: center;">
          <h2>{{ headerTitle }}</h2>
          <el-button type="primary" @click="handleAdd">{{ addButtonText }}</el-button>
        </el-header>
        <el-main>
          <!-- 会计实体信息表格 -->
          <el-table :data="paginatedAcctData" style="width: 100%" max-height="450" v-if="activeMenu === '1-1'">
            <el-table-column prop="AcctCode" label="会计实体编码" width="160%"></el-table-column>
            <el-table-column prop="AcctAbbr" label="会计实体缩写" width="160%"></el-table-column>
            <el-table-column prop="EtyAbbr" label="实体简称" width="160%"></el-table-column>
            <el-table-column prop="AcctName" label="会计实体名称" width="220%"></el-table-column>


            <el-table-column label="绑定的银行账户信息" width="220%">
              <!-- <el-input v-model="acctForm.BankAccounts" :readonly="true"></el-input> -->
              <template #default="scope">
                <span v-if="scope.row.AcctBanks && scope.row.AcctBanks.length > 0">
                  {{ scope.row.AcctBanks.map(bank => bank.AccNum).join(', ') }}
                </span>
                <span v-else>无</span>
              </template>
            </el-table-column>

            <el-table-column prop="AcctEngName" label="会计实体英文名称" width="220%"></el-table-column>
            <el-table-column prop="AcctAddr" label="会计实体地址" width="220%"></el-table-column>
            <el-table-column prop="Nation" label="国家"></el-table-column>
            <el-table-column prop="TaxType" label="税号类型" width="220%"></el-table-column>
            <el-table-column prop="TaxCode" label="税号" width="220%"></el-table-column>
            <el-table-column prop="PhoneNum" label="联系电话" width="220%"></el-table-column>
            <el-table-column prop="Email" label="邮箱" width="220%"></el-table-column>
            <el-table-column prop="Website" label="网站" width="220%"></el-table-column>
            <el-table-column prop="RegDate" label="注册日期" width="450%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="220%"></el-table-column>

            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="150">
              <template #default="scope">
                <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
              </template>
            </el-table-column>

          </el-table>

          <!-- 会计实体银行账户信息表格 -->
          <el-table :data="paginatedBankData" style="width: 100%" max-height="450" v-if="activeMenu === '1-2'">
            <el-table-column prop="AccName" label="账户名称" width="220%"></el-table-column>
            <el-table-column prop="AccNum" label="账号" width="220%"></el-table-column>

            <el-table-column label="对应的会计实体信息" width="320%">
              <template #default="scope">
                <span v-if="scope.row.AcctName">{{ scope.row.AcctName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="BankName" label="开户行名称" width="220%"></el-table-column>
            <el-table-column prop="BankNum" label="开户行号" width="320%"></el-table-column>
            <el-table-column prop="SwiftCode" label="SWIFT CODE" width="220%"></el-table-column>
            <el-table-column prop="BankAddr" label="开户行地址" width="280%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="320%"></el-table-column>

            <el-table-column prop="FileName" label="文件名" width="320%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="150">
              <template #default="scope">
                <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize"
            :total="activeMenu === '1-1' ? acctData.length : bankData.length" layout="prev, pager, next"
            @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加会计实体信息的对话框 -->
    <el-dialog v-model="showAcctDialog" title="会计实体信息" width="80%" @close="resetAcctForm">
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
            <el-form-item label="会计实体名称" prop="AcctName" :required="true">
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

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!acctForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(acctForm.FileId, acctForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="acctForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 会计实体银行账户信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联银行账户信息">

              <!-- <el-input v-model="acctForm.BankAccounts" :readonly="true"></el-input> -->
              <el-input v-model="acctForm.AcctBanksDisplay" :readonly="true"></el-input>
              <!-- <el-select v-model="acctForm.BankAccounts" multiple placeholder="请选择银行账户信息"> -->
              <!--   <el-option v-for="bank in bankData" :key="bank.AccNum" :label="`${bank.AccName} (${bank.AccNum})`" -->
              <!--     :value="bank.AccNum"></el-option> -->
              <!-- </el-select> -->
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
    <el-dialog v-model="showBankDialog" title="会计实体银行账户信息" width="80%" @close="resetBankForm">
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

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!bankForm.FileId" ref="bankUploadRef" action="" :limit="1"
                :on-change="handleBankFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(bankForm.FileId, bankForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="bankForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>


        <!-- 会计实体信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联会计实体信息">
              <el-select v-model="bankForm.AcctId" @change=onAcctChange placeholder="请选择会计实体信息">
                <el-option v-for="acct in acctData" :key="acct.AcctCode" :label="`${acct.AcctName} (${acct.AcctCode})`"
                  :value="acct.AcctId"></el-option>
              </el-select>
              <!-- <el-select v-model="selectedAcct" placeholder="请选择会计实体信息"> -->
              <!--   <el-option v-for="acct in acctData" :key="acct.AcctCode" :label="`${acct.AcctName} `" -->
              <!--     :value="{ AcctId: acct.AcctId, AcctName: acct.AcctName }"></el-option> -->
              <!-- </el-select> -->

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


    <!-- 查看会计实体信息的对话框 -->
    <el-dialog v-model="showshowAcctDialog" title="会计实体信息" width="80%" @close="resetAcctForm">
      <el-form :model="acctForm" label-width="150px" :rules="acctRules" ref="acctFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="会计实体编码" prop="AcctCode" :required="true">
              <el-input v-model="acctForm.AcctCode" placeholder="国家代码+流水码" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体缩写" prop="AcctAbbr" :required="true">
              <el-input v-model="acctForm.AcctAbbr" placeholder="名称缩写，一般2-3个字母" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="实体简称" prop="EtyAbbr" :required="true">
              <el-input v-model="acctForm.EtyAbbr" placeholder="客户名称简称，一般1-2个单词" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体名称" prop="AcctName" :required="true">
              <el-input v-model="acctForm.AcctName" placeholder="全称" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="会计实体英文名称" prop="AcctEngName">
              <el-input v-model="acctForm.AcctEngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体地址" prop="AcctAddr">
              <el-input v-model="acctForm.AcctAddr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="acctForm.Nation" :readonly="true"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-input v-model="acctForm.TaxType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="税号" prop="TaxCode">
              <el-input v-model="acctForm.TaxCode" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="acctForm.PhoneNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="acctForm.Email" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="网站" prop="Website">
              <el-input v-model="acctForm.Website" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册日期" prop="RegDate">
              <el-date-picker v-model="acctForm.RegDate" type="date" placeholder="选择日期"
                :readonly="true"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="acctForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!acctForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(acctForm.FileId, acctForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="acctForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 会计实体银行账户信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联银行账户信息">

              <!-- <el-input v-model="acctForm.BankAccounts" :readonly="true"></el-input> -->
              <el-input v-model="acctForm.AcctBanksDisplay" :readonly="true"></el-input>
              <!-- <el-select v-model="acctForm.BankAccounts" multiple placeholder="请选择银行账户信息"> -->
              <!--   <el-option v-for="bank in bankData" :key="bank.AccNum" :label="`${bank.AccName} (${bank.AccNum})`" -->
              <!--     :value="bank.AccNum"></el-option> -->
              <!-- </el-select> -->
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showshowAcctDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 添加会计实体银行账户信息的对话框 -->
    <el-dialog v-model="showshowBankDialog" title="会计实体银行账户信息" width="80%" @close="resetBankForm">
      <el-form :model="bankForm" label-width="150px" :rules="bankRules" ref="bankFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="账户名称" prop="AccName">
              <el-input v-model="bankForm.AccName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="账号" prop="AccNum">
              <el-input v-model="bankForm.AccNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-input v-model="bankForm.Currency" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="开户行名称" prop="BankName">
              <el-input v-model="bankForm.BankName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行号" prop="BankNum">
              <el-input v-model="bankForm.BankNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="SWIFT CODE" prop="SwiftCode">
              <el-input v-model="bankForm.SwiftCode" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开户行地址" prop="BankAddr">
              <el-input v-model="bankForm.BankAddr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="bankForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!bankForm.FileId" ref="bankUploadRef" action="" :limit="1"
                :on-change="handleBankFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(bankForm.FileId, bankForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="bankForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 会计实体信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联会计实体信息">
              <el-select v-model="bankForm.AcctId" @change=onAcctChange placeholder="请选择会计实体信息">
                <el-option v-for="acct in acctData" :key="acct.AcctCode" :label="`${acct.AcctName} (${acct.AcctCode})`"
                  :value="acct.AcctId"></el-option>
              </el-select>
              <!-- <el-select v-model="selectedAcct" placeholder="请选择会计实体信息"> -->
              <!--   <el-option v-for="acct in acctData" :key="acct.AcctCode" :label="`${acct.AcctName} `" -->
              <!--     :value="{ AcctId: acct.AcctId, AcctName: acct.AcctName }"></el-option> -->
              <!-- </el-select> -->

            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showshowBankDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>

import { ref, onMounted, computed } from 'vue';

import { ElMessage } from 'element-plus'; // 引入 ElMessage
import axios from 'axios'; // 引入 axios
const downloadFile = async (fileId, fileName) => {
  try {
    const response = await axios.post(
      '/file',
      { FileId: fileId }, // 提供 FileId 表单
      {
        responseType: 'blob', // 指定响应类型为二进制数据
      }
    );

    // 创建下载链接
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    // link.setAttribute('download', `file_${fileId}`); // 设置下载文件名
    link.setAttribute('download', `${fileName}` || `file_${fileId}`); // 使用 fileName 作为文件名，如果不存在则使用默认文件名
    document.body.appendChild(link);
    link.click(); // 触发下载
    document.body.removeChild(link); // 移除链接
    window.URL.revokeObjectURL(url); // 释放 URL 对象
  } catch (error) {
    console.error('文件下载失败:', error);
    ElMessage.error('文件下载失败，请稍后重试');
  }
};
// 文件数据（会计实体银行账户信息）
const bankFile = ref(null);

// 文件选择事件（会计实体银行账户信息）
const handleBankFileChange = (uploadFile) => {
  bankFile.value = uploadFile.raw; // 保存选择的文件
};
// 获取 el-upload 组件的引用
const uploadRef = ref(null);
// 文件数据
const file = ref(null);

// 文件选择事件
const handleFileChange = (uploadFile) => {
  file.value = uploadFile.raw; // 保存选择的文件
};
// 查看按钮逻辑
const handleView = (index, row) => {
  if (activeMenu.value === '1-1') {
    // 填充会计实体信息表单
    acctForm.value = { ...row }; // 将当前行的数据赋值给 acctForm
    // 将 AcctBanks 转换为可显示的字符串
    acctForm.value.AcctBanksDisplay = row.AcctBanks && row.AcctBanks.length > 0
      ? row.AcctBanks.map(bank => bank.AccNum).join(', ')
      : '无';
    showshowAcctDialog.value = true; // 打开会计实体信息对话框
    // 检查是否已上传文件
    if (row.FileId) {
      acctForm.value.FileId = row.FileId; // 保存 FileId
      acctForm.value.FileName = row.FileName; // 保存 FileId
    }
  } else if (activeMenu.value === '1-2') {
    // 填充会计实体银行账户信息表单
    bankForm.value = { ...row }; // 将当前行的数据赋值给 bankForm
    selectedAcct.value = { AcctId: row.AcctId, AcctName: row.AcctName }; // 填充关联的会计实体信息
    showshowBankDialog.value = true; // 打开会计实体银行账户信息对话框
    // 检查是否已上传文件
    if (row.FileId) {
      bankForm.value.FileId = row.FileId; // 保存 FileId
      bankForm.value.FileName = row.FileName; // 保存 FileId
    }
  }
};
// 重置会计实体信息表单
const resetAcctForm = () => {
  acctForm.value = {
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
  };
  file.value = null; // 重置文件数据
  if (uploadRef.value) {
    uploadRef.value.clearFiles(); // 清空文件列表
  }
};
// 获取 el-upload 组件的引用（会计实体银行账户信息）
const bankUploadRef = ref(null);
// 重置会计实体银行账户信息表单
const resetBankForm = () => {
  bankForm.value = {
    AccName: '',
    AccNum: '',
    Currency: '',
    BankName: '',
    BankNum: '',
    SwiftCode: '',
    BankAddr: '',
    Notes: '',
    File: '',
    AcctId: '', // 关联的会计实体 ID
    AcctName: '' // 关联的会计实体名称
  };
  selectedAcct.value = null; // 重置选择的会计实体信息
  bankFile.value = null; // 重置文件数据
  if (bankUploadRef.value) {
    bankUploadRef.value.clearFiles(); // 清空文件列表
  }
};
// 监听 change 事件并更新 AcctName
function onAcctChange(value) {
  const selectedAcct = acctData.value.find(acct => acct.AcctId === value);
  if (selectedAcct) {
    bankForm.value.AcctName = selectedAcct.AcctName;
  }
}
const selectedAcctName = computed(() => {
  const acct = acctData.value.find(acct => acct.AcctId === bankForm.value.AcctId);
  return acct ? acct.AcctName : '';
});
const selectedAcct = ref({ AcctId: '', AcctName: '' }); // 默认值为空对象
// 会计实体信息表单提交逻辑
const submitAcctForm = async () => {
  try {
    // acctForm.value.AcctId = parseInt(acctForm.value.AcctId, 10);

    // const response = await axios.post('/save/acct', acctForm.value); // 调用保存会计实体信息接口
    const formData = new FormData(); // 创建 FormData 对象

    // 添加表单数据
    Object.keys(acctForm.value).forEach((key) => {
      formData.append(key, acctForm.value[key]);
    });

    // 添加文件
    if (file.value) {
      formData.append('file', file.value);
    }

    const response = await axios.post('/save/acct', formData, {
      headers: {
        'Content-Type': 'multipart/form-data', // 设置请求头
      },
    });
    if (response.status === 200) {
      ElMessage.success('会计实体信息保存成功');
      showAcctDialog.value = false; // 关闭对话框
      fetchAcctData(); // 重新获取会计实体信息数据
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 会计实体银行账户信息表单提交逻辑
const submitBankForm = async () => {
  try {
    const formData = new FormData(); // 创建 FormData 对象

    // 添加表单数据
    Object.keys(bankForm.value).forEach((key) => {
      formData.append(key, bankForm.value[key]);
    });

    // 添加文件
    if (bankFile.value) {
      formData.append('file', bankFile.value);
    }

    const response = await axios.post('/save/acctBank', formData, {
      headers: {
        'Content-Type': 'multipart/form-data', // 设置请求头
      },
    });
    // 将选择的会计实体信息赋值给 bankForm
    // const response = await axios.post('/save/acctBank', bankForm.value); // 调用保存会计实体银行账户信息接口
    if (response.status === 200) {
      ElMessage.success('会计实体银行账户信息保存成功');
      showBankDialog.value = false; // 关闭对话框
      fetchAcctBankData(); // 重新获取会计实体银行账户信息数据
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存会计实体银行账户信息失败:', error);
    ElMessage.error('保存会计实体银行账户信息失败，请稍后重试');
  }
};
const handlePageChange = (page) => {
  currentPage.value = page;
};
const currentPage = ref(1); // 当前页码
const pageSize = 8; // 每页显示的行数
// 计算当前页显示的会计实体信息数据
const paginatedAcctData = computed(() => {
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return acctData.value.slice(start, end);
});

// 计算当前页显示的会计实体银行账户信息数据
const paginatedBankData = computed(() => {
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return bankData.value.slice(start, end);
});
onMounted(() => {
  fetchAcctData(); // 获取会计实体信息
  fetchAcctBankData(); // 获取会计实体银行账户信息
});
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

const fetchAcctBankData = async () => {
  try {
    const response = await axios.get('/find/acctBank'); // 调用会计实体银行账户信息接口
    bankData.value = response.data.AcctBank; // 假设返回的数据结构中有 AcctBank 字段
  } catch (error) {
    ElMessage.error('获取会计实体银行账户信息失败，请稍后重试');
    console.error('获取会计实体银行账户信息失败:', error);
  }
};
// 当前选中的菜单项
const activeMenu = ref('1-1');

// 会计实体信息对话框显示状态
const showAcctDialog = ref(false);

// 会计实体银行账户信息对话框显示状态
const showBankDialog = ref(false);

// 会计实体信息对话框显示状态
const showshowAcctDialog = ref(false);

// 会计实体银行账户信息对话框显示状态
const showshowBankDialog = ref(false);
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
  AcctId: '',
  FileName: '', // 添加 FileName 字段
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
  AcctId: '',
  FileId: '', // 添加 FileName 字段
  FileName: '', // 添加 FileName 字段
  AcctName: ''
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



// 表格数据（初始值为空数组）
const acctData = ref([]); // 会计实体信息
const bankData = ref([]); // 会计实体银行账户信息
// 根据当前选中的菜单项动态更改标题和按钮文本
const headerTitle = computed(() => {
  return activeMenu.value === '1-1' ? '会计实体信息' : '会计实体银行账户信息';
});

const addButtonText = computed(() => {
  return activeMenu.value === '1-1' ? '添加会计实体信息' : '添加会计实体银行账户信息';
});

// 菜单项选择事件
const handleMenuSelect = (index) => {
  activeMenu.value = index;
  if (index === '1-1') {
    fetchAcctData(); // 重新获取会计实体信息数据
  } else if (index === '1-2') {
    fetchAcctBankData()
  }
};

// 添加按钮点击事件
const handleAdd = () => {
  if (activeMenu.value === '1-1') {
    showAcctDialog.value = true;
  } else if (activeMenu.value === '1-2') {
    showBankDialog.value = true;
  }
};


// 编辑按钮逻辑
const handleEdit = (index, row) => {
  if (activeMenu.value === '1-1') {
    // 填充会计实体信息表单

    acctForm.value = { ...row }; // 将当前行的数据赋值给 acctForm
    acctForm.value.AcctBanksDisplay = row.AcctBanks && row.AcctBanks.length > 0
      ? row.AcctBanks.map(bank => bank.AccNum).join(', ')
      : '无';
    showAcctDialog.value = true; // 打开会计实体信息对话框
  } else if (activeMenu.value === '1-2') {
    // 填充会计实体银行账户信息表单
    bankForm.value = { ...row }; // 将当前行的数据赋值给 bankForm
    selectedAcct.value = { AcctId: row.AcctId, AcctName: row.AcctName }; // 填充关联的会计实体信息
    showBankDialog.value = true; // 打开会计实体银行账户信息对话框
  }
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

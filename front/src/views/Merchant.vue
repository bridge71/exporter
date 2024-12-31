<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <el-menu :default-active="activeMenu" class="el-menu-vertical-demo" @select="handleMenuSelect">
            <!-- 菜单项内容保持不变 -->
            <el-submenu index="1">
              <template #title>会计实体信息</template>
              <el-menu-item index="1-1" @click="pushAcct">会计实体信息</el-menu-item>
              <el-menu-item index="1-2" @click="pushAcctBank">会计实体银行账户信息</el-menu-item>
              <el-menu-item index="1-3" @click="pushMerchant">客商信息</el-menu-item>
              <el-menu-item index="1-4">联系人信息</el-menu-item>
              <el-menu-item index="1-5">银行账户信息</el-menu-item>
              <el-menu-item index="1-6">库存地点信息</el-menu-item>
              <el-menu-item index="1-7">付款方式信息</el-menu-item>
              <el-menu-item index="1-8">品类信息</el-menu-item>
              <el-menu-item index="1-9">品牌信息</el-menu-item>
              <el-menu-item index="1-10">包装规格信息</el-menu-item>
              <el-menu-item index="1-11">员工信息</el-menu-item>
              <el-menu-item index="1-12">产品明细</el-menu-item>
              <el-menu-item index="1-13">装货明细</el-menu-item>
              <el-menu-item index="1-14">费用明细</el-menu-item>
              <el-menu-item index="1-15">销售订单</el-menu-item>
              <el-menu-item index="1-16">采购收单</el-menu-item>
              <el-menu-item index="1-17">应付账款单</el-menu-item>
              <el-menu-item index="1-18">收款单</el-menu-item>
              <el-menu-item index="1-19">付款单</el-menu-item>
              <!-- 其他菜单项 -->
            </el-submenu>
            <!-- 其他菜单 -->
          </el-menu>
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>
        <el-header style="display: flex; justify-content: space-between; align-items: center;">
          <h2>{{ headerTitle }}</h2>
          <el-button type="primary" @click="handleAdd">{{ addButtonText }}</el-button>
        </el-header>
        <el-main>
          <!-- 商客信息表格 -->
          <el-table :data="paginatedMerchantData" style="width: 100%" max-height="450" v-if="activeMenu === '1-3'">
            <el-table-column prop="MercCode" label="客商编码" width="160%"></el-table-column>
            <el-table-column prop="MercAbbr" label="客商缩写" width="160%"></el-table-column>
            <el-table-column prop="ShortMerc" label="客商简称" width="160%"></el-table-column>
            <el-table-column prop="Merc" label="客商名称" width="220%"></el-table-column>
            <el-table-column prop="EngName" label="客商英文名称" width="220%"></el-table-column>
            <el-table-column prop="Address" label="地址" width="220%"></el-table-column>
            <el-table-column prop="Nation" label="国家"></el-table-column>
            <el-table-column prop="PhoneNum" label="联系电话" width="220%"></el-table-column>
            <el-table-column prop="Email" label="邮箱" width="220%"></el-table-column>
            <el-table-column prop="Fax" label="传真" width="220%"></el-table-column>
            <el-table-column prop="Website" label="网址" width="220%"></el-table-column>
            <el-table-column prop="TaxType" label="税号类型" width="220%"></el-table-column>
            <el-table-column prop="TaxCode" label="税号" width="220%"></el-table-column>
            <el-table-column prop="MerchantType" label="客商类型" width="220%"></el-table-column>
            <el-table-column prop="RegCap" label="注册资本" width="220%"></el-table-column>
            <el-table-column prop="OperCap" label="营业资本" width="220%"></el-table-column>
            <el-table-column label="联系人信息" width="220%">
              <template #default="scope">
                <span v-if="scope.row.Contacts && scope.row.Contacts.length > 0">
                  {{ scope.row.Contacts.map(contact => contact.Name).join(', ') }}
                </span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column label="银行账户信息" width="220%">
              <template #default="scope">
                <span v-if="scope.row.BankAccounts && scope.row.BankAccounts.length > 0">
                  {{ scope.row.BankAccounts.map(bank => bank.AccNum).join(', ') }}
                </span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.MercId)" type="text"
                    size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize"
            :total="activeMenu === '1-1' ? merchantData.length : 0" layout="prev, pager, next"
            @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加客商信息的对话框 -->
    <el-dialog v-model="showMerchantDialog" title="客商信息" width="80%" @close="resetMerchantForm">
      <el-form :model="merchantForm" label-width="150px" :rules="merchantRules" ref="merchantFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商编码" prop="MercCode" :required="true">
              <el-input v-model="merchantForm.MercCode" placeholder="国家代码+流水码"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商缩写" prop="MercAbbr" :required="true">
              <el-input v-model="merchantForm.MercAbbr" placeholder="名称缩写，一般2-3个字母"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商简称" prop="ShortMerc" :required="true">
              <el-input v-model="merchantForm.ShortMerc" placeholder="客户名称简称，一般1-2个单词"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商名称" prop="Merc" :required="true">
              <el-input v-model="merchantForm.Merc" placeholder="全称"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商英文名称" prop="EngName">
              <el-input v-model="merchantForm.EngName" placeholder="英文全称"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地址" prop="Address">
              <el-input v-model="merchantForm.Address"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="merchantForm.Nation"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="merchantForm.PhoneNum"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="merchantForm.Email"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="传真" prop="Fax">
              <el-input v-model="merchantForm.Fax"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="网址" prop="Website">
              <el-input v-model="merchantForm.Website"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-select v-model="merchantForm.TaxType" placeholder="请选择">
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
              <el-input v-model="merchantForm.TaxCode"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商类型" prop="MerchantType">
              <el-input v-model="merchantForm.MerchantType"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册资本" prop="RegCap">
              <el-input v-model="merchantForm.RegCap"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="营业资本" prop="OperCap">
              <el-input v-model="merchantForm.OperCap"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 联系人信息 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="联系人信息">
              <el-input v-model="merchantForm.ContactsDisplay" :readonly="true"></el-input>
              <!-- <el-select v-model="merchantForm.Contacts" multiple placeholder="请选择联系人信息">
                <el-option v-for="contact in contactData" :key="contact.ContactId" :label="`${contact.Name} (${contact.PhoneNum})`"
                  :value="contact.ContactId"></el-option>
              </el-select> -->
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 银行账户信息 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="银行账户信息">
              <el-input v-model="merchantForm.BankAccountsDisplay" :readonly="true"></el-input>
              <!-- <el-select v-model="merchantForm.BankAccounts" multiple placeholder="请选择银行账户信息">
                <el-option v-for="bank in bankData" :key="bank.AccNum" :label="`${bank.AccName} (${bank.AccNum})`"
                  :value="bank.AccNum"></el-option>
              </el-select> -->
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!merchantForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(merchantForm.FileId, merchantForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="merchantForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitMerchantForm">保存</el-button>
            <el-button @click="showMerchantDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看客商信息的对话框 -->
    <el-dialog v-model="showshowMerchantDialog" title="客商信息" width="80%" @close="resetMerchantForm">
      <el-form :model="merchantForm" label-width="150px" :rules="merchantRules" ref="merchantFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商编码" prop="MercCode" :required="true">
              <el-input v-model="merchantForm.MercCode" placeholder="国家代码+流水码" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商缩写" prop="MercAbbr" :required="true">
              <el-input v-model="merchantForm.MercAbbr" placeholder="名称缩写，一般2-3个字母" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商简称" prop="ShortMerc" :required="true">
              <el-input v-model="merchantForm.ShortMerc" placeholder="客户名称简称，一般1-2个单词" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商名称" prop="Merc" :required="true">
              <el-input v-model="merchantForm.Merc" placeholder="全称" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商英文名称" prop="EngName">
              <el-input v-model="merchantForm.EngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地址" prop="Address">
              <el-input v-model="merchantForm.Address" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="merchantForm.Nation" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="merchantForm.PhoneNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="merchantForm.Email" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="传真" prop="Fax">
              <el-input v-model="merchantForm.Fax" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="网址" prop="Website">
              <el-input v-model="merchantForm.Website" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-input v-model="merchantForm.TaxType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="税号" prop="TaxCode">
              <el-input v-model="merchantForm.TaxCode" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商类型" prop="MerchantType">
              <el-input v-model="merchantForm.MerchantType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册资本" prop="RegCap">
              <el-input v-model="merchantForm.RegCap" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="营业资本" prop="OperCap">
              <el-input v-model="merchantForm.OperCap" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看客商信息的对话框 -->
    <el-dialog v-model="showshowMerchantDialog" title="客商信息" width="80%" @close="resetMerchantForm">
      <el-form :model="merchantForm" label-width="150px" :rules="merchantRules" ref="merchantFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商编码" prop="MercCode" :required="true">
              <el-input v-model="merchantForm.MercCode" placeholder="国家代码+流水码" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商缩写" prop="MercAbbr" :required="true">
              <el-input v-model="merchantForm.MercAbbr" placeholder="名称缩写，一般2-3个字母" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商简称" prop="ShortMerc" :required="true">
              <el-input v-model="merchantForm.ShortMerc" placeholder="客户名称简称，一般1-2个单词" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商名称" prop="Merc" :required="true">
              <el-input v-model="merchantForm.Merc" placeholder="全称" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="客商英文名称" prop="EngName">
              <el-input v-model="merchantForm.EngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地址" prop="Address">
              <el-input v-model="merchantForm.Address" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="merchantForm.Nation" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="merchantForm.PhoneNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="merchantForm.Email" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="传真" prop="Fax">
              <el-input v-model="merchantForm.Fax" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="网址" prop="Website">
              <el-input v-model="merchantForm.Website" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-input v-model="merchantForm.TaxType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="税号" prop="TaxCode">
              <el-input v-model="merchantForm.TaxCode" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="客商类型" prop="MerchantType">
              <el-input v-model="merchantForm.MerchantType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册资本" prop="RegCap">
              <el-input v-model="merchantForm.RegCap" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="营业资本" prop="OperCap">
              <el-input v-model="merchantForm.OperCap" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 联系人信息 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="联系人信息">
              <el-input v-model="merchantForm.ContactsDisplay" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 银行账户信息 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="银行账户信息">
              <el-input v-model="merchantForm.BankAccountsDisplay" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!merchantForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(merchantForm.FileId, merchantForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="merchantForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showshowMerchantDialog = false">关闭</el-button>
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
import { useRouter } from 'vue-router';

const router = useRouter();

const pushAcct = () => {
  router.push('/acct');
};
const pushAcctBank = () => {
  router.push('/acctBank');
};

const pushMerchant = () => {
  router.push('/merchant');
};
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

const handleDelete = (index, MercId) => {
  ElMessageBox.confirm('确定要删除该客商信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/merchant', {
      "MercCode": "ss",
      "MercAbbr": "ss",
      "ShortMerc": "ss",
      "MercId": MercId,
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchMerchantData(); // 重新获取客商信息数据
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

// 文件数据（客商信息）
const file = ref(null);

// 文件选择事件（客商信息）
const handleFileChange = (uploadFile) => {
  file.value = uploadFile.raw; // 保存选择的文件
};

// 查看按钮逻辑
const handleView = (index, row) => {
  // 填充客商信息表单
  merchantForm.value = { ...row }; // 将当前行的数据赋值给 merchantForm
  // 将 Contacts 和 BankAccounts 转换为可显示的字符串
  merchantForm.value.ContactsDisplay = row.Contacts && row.Contacts.length > 0
    ? row.Contacts.map(contact => contact.Name).join(', ')
    : '无';
  merchantForm.value.BankAccountsDisplay = row.BankAccounts && row.BankAccounts.length > 0
    ? row.BankAccounts.map(bank => bank.AccNum).join(', ')
    : '无';
  showshowMerchantDialog.value = true; // 打开客商信息对话框
  // 检查是否已上传文件
  if (row.FileId) {
    merchantForm.value.FileId = row.FileId; // 保存 FileId
    merchantForm.value.FileName = row.FileName; // 保存 FileId
  }
};

// 获取 el-upload 组件的引用（客商信息）
const uploadRef = ref(null);

// 重置客商信息表单
const resetMerchantForm = () => {
  merchantForm.value = {
    MercCode: '',
    MercAbbr: '',
    ShortMerc: '',
    Merc: '',
    EngName: '',
    Address: '',
    Nation: '',
    PhoneNum: '',
    Email: '',
    Fax: '',
    Website: '',
    TaxType: '',
    TaxCode: '',
    MerchantType: '',
    RegCap: '',
    OperCap: '',
    Contacts: [],
    BankAccounts: [],
    File: '',
    FileId: '',
    FileName: '',
    MerchantId: ''
  };
  file.value = null; // 重置文件数据
  if (uploadRef.value) {
    uploadRef.value.clearFiles(); // 清空文件列表
  }
};

// 会计实体信息选择
const onAcctChange = (value) => {
  const selectedAcct = acctData.value.find(acct => acct.AcctId === value);
  if (selectedAcct) {
    merchantForm.value.AcctName = selectedAcct.AcctName;
  }
};

// 商客信息表单提交逻辑
const submitMerchantForm = async () => {
  try {
    const formData = new FormData(); // 创建 FormData 对象

    console.log(merchantForm.value.MercCode)
    console.log(merchantForm.value.ShortMerc)
    console.log(merchantForm.value.MercAbbr)
    // 添加表单数据
    Object.keys(merchantForm.value).forEach((key) => {
      formData.append(key, merchantForm.value[key]);
    });

    // 添加文件
    if (file.value) {
      formData.append('file', file.value);
    }

    const response = await axios.post('/save/merchant', formData, {
      headers: {
        'Content-Type': 'multipart/form-data', // 设置请求头
      },
    });

    if (response.status === 200) {
      ElMessage.success('客商信息保存成功');
      showMerchantDialog.value = false; // 关闭对话框
      fetchMerchantData(); // 重新获取客商信息数据
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    ElMessage.error(error.response.data.RetMessage || '保存失败');
    console.error('保存客商信息失败:', error);
    ElMessage.error('保存客商信息失败，请稍后重试');
  }
};

const handlePageChange = (page) => {
  currentPage.value = page;
};

const currentPage = ref(1); // 当前页码
const pageSize = 8; // 每页显示的行数

// 计算当前页显示的客商信息数据
const paginatedMerchantData = computed(() => {
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return merchantData.value.slice(start, end);
});

onMounted(() => {
  fetchMerchantData(); // 获取客商信息
});

// 定义接口请求函数
const fetchMerchantData = async () => {
  try {
    const response = await axios.get('/find/merchant'); // 调用客商信息接口
    merchantData.value = response.data.Merchant; // 假设返回的数据结构中有 Merchant 字段
  } catch (error) {
    console.error('获取客商信息失败:', error);
    ElMessage.error('获取客商信息失败，请稍后重试');
  }
};

// 当前选中的菜单项
const activeMenu = ref('1-3');

// 商客信息对话框显示状态
const showMerchantDialog = ref(false);

// 查看客商信息对话框显示状态
const showshowMerchantDialog = ref(false);

// 商客信息表单数据
const merchantForm = ref({
  MercCode: '',
  MercAbbr: '',
  ShortMerc: '',
  Merc: '',
  EngName: '',
  Address: '',
  Nation: '',
  PhoneNum: '',
  Email: '',
  Fax: '',
  Website: '',
  TaxType: '',
  TaxCode: '',
  MerchantType: '',
  RegCap: '',
  OperCap: '',
  Contacts: [],
  BankAccounts: [],
  // File: '',
  FileId: '',
  FileName: '',
  MercId: ''
});

// 商客信息表单验证规则
const merchantRules = {
  MercCode: [{ required: true, message: '请输入客商编码', trigger: 'blur' }],
  MercAbbr: [{ required: true, message: '请输入客商缩写', trigger: 'blur' }],
  ShortMerc: [{ required: true, message: '请输入客商简称', trigger: 'blur' }],
  Merc: [{ required: true, message: '请输入客商名称', trigger: 'blur' }]
};

// 表格数据（初始值为空数组）
const merchantData = ref([]); // 商客信息

// 根据当前选中的菜单项动态更改标题和按钮文本
const headerTitle = computed(() => {
  return activeMenu.value === '1-3' ? '客商信息' : '';
});

const addButtonText = computed(() => {
  return activeMenu.value === '1-3' ? '添加客商信息' : '';
});

// 菜单项选择事件
const handleMenuSelect = (index) => {
  activeMenu.value = index;
  if (index === '1-3') {
    fetchMerchantData(); // 重新获取客商信息数据
  }
};

// 添加按钮点击事件
const handleAdd = () => {
  showMerchantDialog.value = true;
};

// 编辑按钮逻辑
const handleEdit = (index, row) => {
  merchantForm.value = { ...row }; // 将当前行的数据赋值给 merchantForm
  // 将 Contacts 和 BankAccounts 转换为可显示的字符串
  merchantForm.value.ContactsDisplay = row.Contacts && row.Contacts.length > 0
    ? row.Contacts.map(contact => contact.Name).join(', ')
    : '无';
  merchantForm.value.BankAccountsDisplay = row.BankAccounts && row.BankAccounts.length > 0
    ? row.BankAccounts.map(bank => bank.AccNum).join(', ')
    : '无';
  showMerchantDialog.value = true; // 打开客商信息对话框
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

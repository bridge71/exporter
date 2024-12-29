<template>
  <div class="table-container">
    <el-button @click="dialogVisible = true" type="primary" style="margin-bottom: 10px;">添加条目</el-button>
    <el-table
      :data="tableData"
      border
      style="width: 100%"
      @selection-change="handleSelectionChange"
      highlight-current-row>
      <el-table-column type="selection" width="50"></el-table-column>
      <el-table-column prop="id" label="#" sortable width="60"></el-table-column>
      <el-table-column prop="code" label="编码" sortable></el-table-column>
      <el-table-column prop="abbreviation" label="缩写"></el-table-column>
      <el-table-column prop="shortName" label="简称"></el-table-column>
      <el-table-column prop="companyName" label="名称"></el-table-column>
      <el-table-column prop="englishName" label="英文名称"></el-table-column>
      <el-table-column prop="address" label="地址"></el-table-column>
      <el-table-column prop="country" label="国家"></el-table-column>
      <el-table-column prop="phone" label="联系电话"></el-table-column>
      <el-table-column prop="fax" label="传真"></el-table-column>
      <el-table-column prop="taxType" label="税号类型"></el-table-column>
      <el-table-column prop="taxNumber" label="税号"></el-table-column>
      <el-table-column prop="merchantType" label="客商类型"></el-table-column>
      <el-table-column prop="registeredCapital" label="注册资本"></el-table-column>
      <el-table-column prop="businessLicense" label="营业执照"></el-table-column>
      <el-table-column prop="contactPerson" label="联系人"></el-table-column>
      <el-table-column prop="bankAccount" label="银行账户信息"></el-table-column>
      <el-table-column prop="email" label="邮箱"></el-table-column>
      <el-table-column prop="website" label="网站"></el-table-column>
      <el-table-column label="操作" fixed="right" width="120">
        <template #default="{ row }">
          <el-button @click="editRow(row)" type="text">编辑</el-button>
          <el-button @click="deleteRow(row)" type="text">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 添加 Dialog -->
    <el-dialog title="添加条目" v-model="dialogVisible" width="50%">
      <el-form :model="form" label-width="120px">
        <el-form-item label="编码">
          <el-input v-model="form.code"></el-input>
        </el-form-item>
        <el-form-item label="缩写">
          <el-input v-model="form.abbreviation"></el-input>
        </el-form-item>
        <el-form-item label="简称">
          <el-input v-model="form.shortName"></el-input>
        </el-form-item>
        <el-form-item label="名称">
          <el-input v-model="form.companyName"></el-input>
        </el-form-item>
        <el-form-item label="英文名称">
          <el-input v-model="form.englishName"></el-input>
        </el-form-item>
        <el-form-item label="地址">
          <el-input v-model="form.address"></el-input>
        </el-form-item>
        <el-form-item label="国家">
          <el-input v-model="form.country"></el-input>
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="form.phone"></el-input>
        </el-form-item>
        <el-form-item label="传真">
          <el-input v-model="form.fax"></el-input>
        </el-form-item>
        <el-form-item label="税号类型">
          <el-input v-model="form.taxType"></el-input>
        </el-form-item>
        <el-form-item label="税号">
          <el-input v-model="form.taxNumber"></el-input>
        </el-form-item>
        <el-form-item label="客商类型">
          <el-input v-model="form.merchantType"></el-input>
        </el-form-item>
        <el-form-item label="注册资本">
          <el-input v-model="form.registeredCapital"></el-input>
        </el-form-item>
        <el-form-item label="营业执照">
          <el-input v-model="form.businessLicense"></el-input>
        </el-form-item>
        <el-form-item label="联系人">
          <el-input v-model="form.contactPerson"></el-input>
        </el-form-item>
        <el-form-item label="银行账户信息">
          <el-input v-model="form.bankAccount"></el-input>
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="form.email"></el-input>
        </el-form-item>
        <el-form-item label="网站">
          <el-input v-model="form.website"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="addRow">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 编辑 Dialog -->
    <el-dialog title="编辑条目" v-model="editDialogVisible" width="50%">
      <el-form :model="editForm" label-width="120px">
        <el-form-item label="编码">
          <el-input v-model="editForm.code"></el-input>
        </el-form-item>
        <el-form-item label="缩写">
          <el-input v-model="editForm.abbreviation"></el-input>
        </el-form-item>
        <el-form-item label="简称">
          <el-input v-model="editForm.shortName"></el-input>
        </el-form-item>
        <el-form-item label="名称">
          <el-input v-model="editForm.companyName"></el-input>
        </el-form-item>
        <el-form-item label="英文名称">
          <el-input v-model="editForm.englishName"></el-input>
        </el-form-item>
        <el-form-item label="地址">
          <el-input v-model="editForm.address"></el-input>
        </el-form-item>
        <el-form-item label="国家">
          <el-input v-model="editForm.country"></el-input>
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="editForm.phone"></el-input>
        </el-form-item>
        <el-form-item label="传真">
          <el-input v-model="editForm.fax"></el-input>
        </el-form-item>
        <el-form-item label="税号类型">
          <el-input v-model="editForm.taxType"></el-input>
        </el-form-item>
        <el-form-item label="税号">
          <el-input v-model="editForm.taxNumber"></el-input>
        </el-form-item>
        <el-form-item label="客商类型">
          <el-input v-model="editForm.merchantType"></el-input>
        </el-form-item>
        <el-form-item label="注册资本">
          <el-input v-model="editForm.registeredCapital"></el-input>
        </el-form-item>
        <el-form-item label="营业执照">
          <el-input v-model="editForm.businessLicense"></el-input>
        </el-form-item>
        <el-form-item label="联系人">
          <el-input v-model="editForm.contactPerson"></el-input>
        </el-form-item>
        <el-form-item label="银行账户信息">
          <el-input v-model="editForm.bankAccount"></el-input>
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="editForm.email"></el-input>
        </el-form-item>
        <el-form-item label="网站">
          <el-input v-model="editForm.website"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="editDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveEdit">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue';

export default {
  setup() {
    const tableData = ref([
      // 示例数据
      { id: 1, code: 'CHN0002', abbreviation: '大成', shortName: '大成品牌', companyName: '大成工业集团有限公司', englishName: 'Changchun Dacheng Industrial Group', address: '长春市大成路', country: '中国', phone: '1234567890', fax: '0987654321', taxType: '企业所得税', taxNumber: '123456789', merchantType: '制造商', registeredCapital: '5000万', businessLicense: '123456789XYZ', contactPerson: '张三', bankAccount: '建设银行 12345678901234', email: 'info@dacheng.com', website: 'www.dacheng.com' },
      // 更多数据
    ]);

    const dialogVisible = ref(false);
    const editDialogVisible = ref(false);
    const form = ref({
      id: tableData.value.length + 1,
      code: '',
      abbreviation: '',
      shortName: '',
      companyName: '',
      englishName: '',
      address: '',
      country: '',
      phone: '',
      fax: '',
      taxType: '',
      taxNumber: '',
      merchantType: '',
      registeredCapital: '',
      businessLicense: '',
      contactPerson: '',
      bankAccount: '',
      email: '',
      website: ''
    });
    const editForm = ref({});

    const handleSelectionChange = (selection) => {
      console.log('Selected items:', selection);
    };

    const editRow = (row) => {
      editForm.value = { ...row };
      editDialogVisible.value = true;
    };

    const deleteRow = (row) => {
      tableData.value = tableData.value.filter(item => item.id !== row.id);
    };

    const addRow = () => {
      tableData.value.push({ ...form.value });
      dialogVisible.value = false;
      form.value = {
        id: tableData.value.length + 1,
        code: '',
        abbreviation: '',
        shortName: '',
        companyName: '',
        englishName: '',
        address: '',
        country: '',
        phone: '',
        fax: '',
        taxType: '',
        taxNumber: '',
        merchantType: '',
        registeredCapital: '',
        businessLicense: '',
        contactPerson: '',
        bankAccount: '',
        email: '',
        website: ''
      };
    };

    const saveEdit = () => {
      const index = tableData.value.findIndex(item => item.id === editForm.value.id);
      if (index !== -1) {
        tableData.value[index] = { ...editForm.value };
      }
      editDialogVisible.value = false;
    };

    return { tableData, dialogVisible, editDialogVisible, form, editForm, handleSelectionChange, editRow, deleteRow, addRow, saveEdit };
  }
};
</script>

<style scoped>
.table-container {
  width: 100%; /* 确保表格容器宽度占满父容器 */
  overflow-x: auto; /* 启用水平滚动 */
  position: relative; /* 确保表格固定位置 */
  height: calc(100vh - 60px); /* 根据实际情况调整高度 */
}

.el-table {
  min-width: 1500px; /* 设置表格最小宽度，根据实际列数调整 */
  width: 100%; /* 确保表格宽度占满容器 */
}
</style>
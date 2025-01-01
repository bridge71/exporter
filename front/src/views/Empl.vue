<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-11'" />
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
          <!-- 员工信息表格 -->
          <el-table :data="paginatedEmplData" style="width: 100%" max-height="450">
            <el-table-column prop="EmpName" label="员工姓名" width="220%"></el-table-column>
            <el-table-column prop="EmpEngName" label="员工英文名" width="220%"></el-table-column>
            <el-table-column prop="Dept" label="部门" width="220%"></el-table-column>
            <el-table-column prop="Position" label="岗位" width="220%"></el-table-column>
            <el-table-column prop="JoinDate" label="入职日期" width="220%"></el-table-column>
            <el-table-column prop="Gender" label="性别" width="220%"></el-table-column>
            <el-table-column prop="Age" label="年龄" width="220%"></el-table-column>
            <el-table-column prop="EduLevel" label="学历" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="320%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="320%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.EmpID)" type="text"
                    size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="emplData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加员工信息的对话框 -->
    <el-dialog v-model="showEmplDialog" title="员工信息" width="80%" @close="resetEmplForm">
      <el-form :model="emplForm" label-width="150px" :rules="emplRules" ref="emplFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="员工姓名" prop="EmpName">
              <el-input v-model="emplForm.EmpName"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="员工英文名" prop="EmpEngName">
              <el-input v-model="emplForm.EmpEngName"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="部门" prop="Dept">
              <el-select v-model="emplForm.Dept" @change="onDeptChange" placeholder="请选择部门">
                <el-option v-for="dept in deptData" :key="dept.Dept" :label="dept.Dept" :value="dept.Dept"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="岗位" prop="Position">
              <el-select v-model="emplForm.Position" @change="onPositionChange" placeholder="请选择岗位">
                <el-option v-for="position in positionData" :key="position.Position" :label="position.Position"
                  :value="position.Position"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="入职日期" prop="JoinDate">
              <el-date-picker v-model="emplForm.JoinDate" type="date" placeholder="选择入职日期"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="性别" prop="Gender">
              <el-select v-model="emplForm.Gender" placeholder="请选择性别">
                <el-option label="男" value="男"></el-option>
                <el-option label="女" value="女"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="年龄" prop="Age">
              <el-input v-model.number="emplForm.Age" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="学历" prop="EduLevel">
              <el-select v-model="emplForm.EduLevel" @change="onEduLevelChange" placeholder="请选择学历">
                <el-option v-for="eduLevel in eduLevelData" :key="eduLevel.EduLevel" :label="eduLevel.EduLevel"
                  :value="eduLevel.EduLevel"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="身份证号" prop="IDCardNum">
              <el-input v-model="emplForm.IDCardNum"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="毕业学校" prop="GradSchool">
              <el-input v-model="emplForm.GradSchool"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="专业" prop="Major">
              <el-input v-model="emplForm.Major"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="居住地区" prop="ResArea">
              <el-input v-model="emplForm.ResArea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="详细地址" prop="AddrDetail">
              <el-input v-model="emplForm.AddrDetail"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="手机号码" prop="MobileNum">
              <el-input v-model="emplForm.MobileNum"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="电子邮箱" prop="EmailAddr">
              <el-input v-model="emplForm.EmailAddr"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="紧急联系人" prop="EmergContact">
              <el-input v-model="emplForm.EmergContact"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="紧急联系方式" prop="EmergContactInfo">
              <el-input v-model="emplForm.EmergContactInfo"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="emplForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!emplForm.FileId" ref="emplUploadRef" action="" :limit="1"
                :on-change="handleEmplFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(emplForm.FileId, emplForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="emplForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitEmplForm">保存</el-button>
            <el-button @click="showEmplDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看员工信息的对话框 -->
    <el-dialog v-model="showViewEmplDialog" title="员工信息" width="80%" @close="resetEmplForm">
      <el-form :model="emplForm" label-width="150px" :rules="emplRules" ref="emplFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="员工姓名" prop="EmpName">
              <el-input v-model="emplForm.EmpName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="员工英文名" prop="EmpEngName">
              <el-input v-model="emplForm.EmpEngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="部门" prop="Dept">
              <el-select v-model="emplForm.Dept" @change="onDeptChange" placeholder="请选择部门" :disabled="true">
                <el-option v-for="dept in deptData" :key="dept.Dept" :label="dept.Dept" :value="dept.Dept"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="岗位" prop="Position">
              <el-select v-model="emplForm.Position" @change="onPositionChange" placeholder="请选择岗位" :disabled="true">
                <el-option v-for="position in positionData" :key="position.Position" :label="position.Position"
                  :value="position.Position"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="入职日期" prop="JoinDate">
              <el-date-picker v-model="emplForm.JoinDate" type="date" placeholder="选择入职日期"
                :disabled="true"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="性别" prop="Gender">
              <el-select v-model="emplForm.Gender" placeholder="请选择性别" :disabled="true">
                <el-option label="男" value="男"></el-option>
                <el-option label="女" value="女"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="年龄" prop="Age">
              <el-input v-model.number="emplForm.Age" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="学历" prop="EduLevel">
              <el-select v-model="emplForm.EduLevel" @change="onEduLevelChange" placeholder="请选择学历" :disabled="true">
                <el-option v-for="eduLevel in eduLevelData" :key="eduLevel.EduLevel" :label="eduLevel.EduLevel"
                  :value="eduLevel.EduLevel"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="身份证号" prop="IDCardNum">
              <el-input v-model="emplForm.IDCardNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="毕业学校" prop="GradSchool">
              <el-input v-model="emplForm.GradSchool" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="专业" prop="Major">
              <el-input v-model="emplForm.Major" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="居住地区" prop="ResArea">
              <el-input v-model="emplForm.ResArea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="详细地址" prop="AddrDetail">
              <el-input v-model="emplForm.AddrDetail" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="手机号码" prop="MobileNum">
              <el-input v-model="emplForm.MobileNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="电子邮箱" prop="EmailAddr">
              <el-input v-model="emplForm.EmailAddr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="紧急联系人" prop="EmergContact">
              <el-input v-model="emplForm.EmergContact" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="紧急联系方式" prop="EmergContactInfo">
              <el-input v-model="emplForm.EmergContactInfo" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="emplForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-button v-if="emplForm.FileId" type="success"
                @click="downloadFile(emplForm.FileId, emplForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="emplForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewEmplDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import axios from 'axios';
import SideMenu from '@/components/SideMenu.vue';

const searchQuery = ref('');
const currentPage = ref(1);
const pageSize = 8;

const emplData = ref([]);
const deptData = ref([]); // 部门数据
const positionData = ref([]); // 岗位数据
const eduLevelData = ref([]); // 学历数据
const showEmplDialog = ref(false);
const showViewEmplDialog = ref(false);
const emplForm = ref({
  EmpID: '',
  EmpName: '',
  EmpEngName: '',
  Dept: '',
  Position: '',
  JoinDate: '',
  Gender: '',
  Age: 0,
  IDCardNum: '',
  EduLevel: '',
  GradSchool: '',
  Major: '',
  ResArea: '',
  AddrDetail: '',
  MobileNum: '',
  EmailAddr: '',
  EmergContact: '',
  EmergContactInfo: '',
  Notes: '',
  FileId: '',
  FileName: '',
});
const emplUploadRef = ref(null);
const emplFile = ref(null);

const handlePageChange = (page) => {
  currentPage.value = page;
};

const paginatedEmplData = computed(() => {
  let filteredData = emplData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.EmpName.includes(searchQuery.value) ||
      item.EmpEngName.includes(searchQuery.value) ||
      item.Dept.includes(searchQuery.value) ||
      item.Position.includes(searchQuery.value) ||
      item.JoinDate.includes(searchQuery.value) ||
      item.Gender.includes(searchQuery.value) ||
      item.Age.toString().includes(searchQuery.value) ||
      item.EduLevel.includes(searchQuery.value) ||
      item.Notes.includes(searchQuery.value) ||
      item.FileName.includes(searchQuery.value)
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

const handleAdd = () => {
  showEmplDialog.value = true;
};

const handleEdit = (index, row) => {
  emplForm.value = { ...row };
  showEmplDialog.value = true;
};

const handleView = (index, row) => {
  emplForm.value = { ...row };
  showViewEmplDialog.value = true;
};

const handleDelete = (index, EmpID) => {
  ElMessageBox.confirm('确定要删除该员工信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    axios.post('/delete/empl', { EmpID })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchEmplData();
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

const resetEmplForm = () => {
  emplForm.value = {
    EmpID: '',
    EmpName: '',
    EmpEngName: '',
    Dept: '',
    Position: '',
    JoinDate: '',
    Gender: '',
    Age: 0,
    IDCardNum: '',
    EduLevel: '',
    GradSchool: '',
    Major: '',
    ResArea: '',
    AddrDetail: '',
    MobileNum: '',
    EmailAddr: '',
    EmergContact: '',
    EmergContactInfo: '',
    Notes: '',
    FileId: '',
    FileName: '',
  };
  emplFile.value = null;
  if (emplUploadRef.value) {
    emplUploadRef.value.clearFiles();
  }
};

const onDeptChange = (value) => {
  const selectedDept = deptData.value.find(dept => dept.Dept === value);
  if (selectedDept) {
    emplForm.value.Dept = selectedDept.Dept;
  }
};

const onPositionChange = (value) => {
  const selectedPosition = positionData.value.find(position => position.Position === value);
  if (selectedPosition) {
    emplForm.value.Position = selectedPosition.Position;
  }
};

const onEduLevelChange = (value) => {
  const selectedEduLevel = eduLevelData.value.find(eduLevel => eduLevel.EduLevel === value);
  if (selectedEduLevel) {
    emplForm.value.EduLevel = selectedEduLevel.EduLevel;
  }
};

const submitEmplForm = async () => {
  try {
    const formData = new FormData();
    Object.keys(emplForm.value).forEach((key) => {
      formData.append(key, emplForm.value[key]);
    });
    if (emplFile.value) {
      formData.append('file', emplFile.value);
    }
    const response = await axios.post('/save/empl', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    if (response.status === 200) {
      ElMessage.success('员工信息保存成功');
      showEmplDialog.value = false;
      fetchEmplData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存员工信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

const handleEmplFileChange = (uploadFile) => {
  emplFile.value = uploadFile.raw;
};

const downloadFile = async (fileId, fileName) => {
  try {
    const response = await axios.post('/file', { FileId: fileId }, { responseType: 'blob' });
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

onMounted(() => {
  fetchEmplData();
  fetchDeptData();
  fetchPositionData();
  fetchEduLevelData();
});

const fetchEmplData = async () => {
  try {
    const response = await axios.get('/find/empl');
    emplData.value = response.data.Empl;
  } catch (error) {
    console.error('获取员工信息失败:', error);
    ElMessage.error('获取员工信息失败，请稍后重试');
  }
};

const fetchDeptData = async () => {
  try {
    const response = await axios.get('/find/dept');
    deptData.value = response.data.Dept;
  } catch (error) {
    console.error('获取部门信息失败:', error);
    ElMessage.error('获取部门信息失败，请稍后重试');
  }
};

const fetchPositionData = async () => {
  try {
    const response = await axios.get('/find/position');
    positionData.value = response.data.Position;
  } catch (error) {
    console.error('获取岗位信息失败:', error);
    ElMessage.error('获取岗位信息失败，请稍后重试');
  }
};

const fetchEduLevelData = async () => {
  try {
    const response = await axios.get('/find/eduLevel');
    eduLevelData.value = response.data.EduLevel;
  } catch (error) {
    console.error('获取学历信息失败:', error);
    ElMessage.error('获取学历信息失败，请稍后重试');
  }
};

const headerTitle = computed(() => '员工信息');
const addButtonText = computed(() => '添加员工信息');

const emplRules = {
  EmpName: [{ required: true, message: '请输入员工姓名', trigger: 'blur' }],
  EmpEngName: [{ required: true, message: '请输入员工英文名', trigger: 'blur' }],
  Dept: [{ required: true, message: '请选择部门', trigger: 'blur' }],
  Position: [{ required: true, message: '请选择岗位', trigger: 'blur' }],
  JoinDate: [{ required: true, message: '请选择入职日期', trigger: 'blur' }],
  Gender: [{ required: true, message: '请选择性别', trigger: 'blur' }],
  Age: [{ required: true, message: '请输入年龄', trigger: 'blur' }],
  EduLevel: [{ required: true, message: '请选择学历', trigger: 'blur' }],
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

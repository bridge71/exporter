<template>
  <div style="display: flex; height: 100vh;">
    <el-menu default-active="2" class="el-menu-vertical-demo" :router="true">
      <el-menu-item @click="handleAcctClick">
        <el-icon>
          <location />
        </el-icon>
        <span>会计实体信息</span>
      </el-menu-item>
      <el-menu-item index="/acctBank">
        <el-icon>
          <document />
        </el-icon>
        <span>会计实体信息银行账户信息</span>
      </el-menu-item>
      <el-menu-item index="/testHome/CL">
        <el-icon>
          <document />
        </el-icon>
        <span>客商信息</span>
      </el-menu-item>
      <el-menu-item index="/item-three">
        <el-icon>
          <document />
        </el-icon>
        <span>联系人信息</span>
      </el-menu-item>
      <el-menu-item index="/item-four">
        <el-icon>
          <document />
        </el-icon>
        <span>银行账户信息</span>
      </el-menu-item>
      <el-menu-item index="/item-one">
        <el-icon>
          <document />
        </el-icon>
        <span>库存地点信息</span>
      </el-menu-item>
      <el-menu-item index="/navigator-two">
        <el-icon><icon-menu /></el-icon>
        <span>付款方式信息</span>
      </el-menu-item>
      <el-menu-item index="/navigator-four">
        <el-icon>
          <setting />
        </el-icon>
        <span>品类信息</span>
      </el-menu-item>
      <!-- 增加更多菜单项 -->
      <el-menu-item index="/item-five">
        <el-icon>
          <document />
        </el-icon>
        <span>品牌信息</span>
      </el-menu-item>
      <el-menu-item index="/item-six">
        <el-icon>
          <document />
        </el-icon>
        <span>包装规格信息</span>
      </el-menu-item>
      <el-menu-item index="/item-seven">
        <el-icon>
          <document />
        </el-icon>
        <span>员工信息</span>
      </el-menu-item>
      <el-menu-item index="/item-eight">
        <el-icon>
          <document />
        </el-icon>
        <span>产品明细</span>
      </el-menu-item>
      <el-menu-item index="/item-nine">
        <el-icon>
          <document />
        </el-icon>
        <span>装货明细</span>
      </el-menu-item>
      <el-menu-item index="/item-ten">
        <el-icon>
          <document />
        </el-icon>
        <span>费用明细</span>
      </el-menu-item>
      <el-menu-item index="/item-eleven">
        <el-icon>
          <document />
        </el-icon>
        <span>销售订单</span>
      </el-menu-item>
      <el-menu-item index="/item-twelve">
        <el-icon>
          <document />
        </el-icon>
        <span>采购收单</span>
      </el-menu-item>
      <el-menu-item index="/item-thirteen">
        <el-icon>
          <document />
        </el-icon>
        <span>应付账款单</span>
      </el-menu-item>
      <el-menu-item index="/item-fourteen">
        <el-icon>
          <document />
        </el-icon>
        <span>收款单</span>
      </el-menu-item>
      <el-menu-item index="/item-fifteen">
        <el-icon>
          <document />
        </el-icon>
        <span>付款单</span>
      </el-menu-item>
    </el-menu>
    <div style="flex: 1; overflow: auto;">
      <el-main>
        <router-view />
      </el-main>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router';
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue';

const router = useRouter();

// 处理会计实体信息点击事件
const handleAcctClick = () => {
  // 假设 acct 是你要传递的数据
  const acct = {
    id: 1,
    name: '会计实体1',
    code: 'ACCT001',
  };

  // 跳转到 /acct 路由，并传递 acct 数据
  router.push({
    path: '/acct',
    query: { acct: JSON.stringify(acct) }, // 将 acct 对象转为字符串传递
  });
};
</script>

<style src="../assets/styles/Bottom.css"></style>

<template>
  <el-menu :default-active="activeMenu" class="el-menu-vertical-demo" @select="handleMenuSelect">
    <el-menu-item index="1-1" @click="pushAcct">会计实体信息</el-menu-item>
    <el-menu-item index="1-2" @click="pushAcctBank">会计实体银行账户信息</el-menu-item>
    <el-menu-item index="1-3" @click="pushMerchant">客商信息</el-menu-item>
    <el-menu-item index="1-4" @click="pushCust">联系人信息</el-menu-item>
    <el-menu-item index="1-5" @click="pushBankAccount">银行账户信息</el-menu-item>
    <el-menu-item index="1-6" @click="pushInvLoc">库存地点信息</el-menu-item>
    <el-menu-item index="1-7" @click="pushPaymentMethod">付款方式信息</el-menu-item>
    <el-menu-item index="1-8" @click="pushCat">品类信息</el-menu-item>
    <el-menu-item index="1-9" @click="pushBrand">品牌信息</el-menu-item>
    <el-menu-item index="1-10" @click="pushSpec">包装规格信息</el-menu-item>
    <el-menu-item index="1-11" @click="pushEmpl">员工信息</el-menu-item>
    <el-menu-item index="1-12" @click="pushPrdt">产品明细</el-menu-item>
    <el-menu-item index="1-13" @click="pushUnload">装货明细</el-menu-item>
    <el-menu-item index="1-14" @click="pushCost">费用明细</el-menu-item>
    <el-menu-item index="1-15" @click="pushSale">销售订单</el-menu-item>
    <el-menu-item index="1-16" @click="pushBuy">采购收单</el-menu-item>
    <el-menu-item index="1-17" @click="pushSend">销售发货单</el-menu-item>
    <el-menu-item index="1-18" @click="pushPurrec">采购收货单</el-menu-item>
    <el-menu-item index="1-19" @click="pushShouldIn">应收账款单</el-menu-item>
    <el-menu-item index="1-20" @click="pushShouldOut">应付账款单</el-menu-item>
    <el-menu-item index="1-21" @click="pushIn">收款单</el-menu-item>
    <el-menu-item index="1-22" @click="pushOut">付款单</el-menu-item>
  </el-menu>
</template>

<script setup>
import { ref, watch } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
// 通过 props 接收默认选中项
const props = defineProps({
  defaultActive: {
    type: String,
    default: '1-1', // 默认值
  },
});

// 动态设置 activeMenu
const activeMenu = ref(props.defaultActive);

// 监听 props 的变化，更新 activeMenu
watch(
  () => props.defaultActive,
  (newVal) => {
    activeMenu.value = newVal;
  }
);

// 处理菜单项选择事件
const handleMenuSelect = (index) => {
  console.log('Selected menu index:', index);
};


// 路由跳转函数
const pushAcct = () => {
  router.push('/acct');
};

const pushAcctBank = () => {
  router.push('/acctBank');
};

const pushMerchant = () => {
  router.push('/merchant');
};

const pushCust = () => {
  router.push('/cust');
};
const pushContact = () => {
  router.push('/contact');
};

const pushBankAccount = () => {
  router.push('/bankAccount');
};

const pushInvLoc = () => {
  router.push('/invLoc');
};

const pushPaymentMethod = () => {
  router.push('/paymentMethod');
};

const pushCat = () => {
  router.push('/cat');
};

const pushBrand = () => {
  router.push('/brand');
};

const pushSpec = () => {
  router.push('/spec');
};

const pushEmpl = () => {
  router.push('/empl');
};

const pushPrdt = () => {
  router.push('/prdt');
};

const pushUnload = () => {
  router.push('/unload');
};

const pushCost = () => {
  router.push('/cost');
};

const pushSale = () => {
  router.push('/sale');
};

const pushBuy = () => {
  router.push('/buy');
};

const pushSend = () => {
  router.push('/send');
};

const pushPurrec = () => {
  router.push('/purrec');
};

const pushShouldIn = () => {
  router.push('/shouldIn');
};

const pushShouldOut = () => {
  router.push('/shouldOut');
};

const pushIn = () => {
  router.push('/in');
};

const pushOut = () => {
  router.push('/out');
};
</script>

<style scoped>
.el-menu-vertical-demo {
  width: 200px;
  /* 设置菜单宽度 */
  background-color: white;
  /* 背景设置为白色 */
}

.el-menu-item {
  color: black;
  /* 文字设置为黑色 */
}

.el-menu-item:hover {
  background-color: #f5f5f5;
  /* 鼠标悬停时的背景色 */
}

.el-menu-item.is-active {
  color: black;
  /* 激活项的文字颜色 */
  background-color: #e0e0e0;
  /* 激活项的背景色 */
}
</style>

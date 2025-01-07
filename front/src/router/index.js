// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
// 导入views中的页面组件
import Login from '@/views/Login.vue';
import DictionaryManager from '@/views/DictionaryManager.vue';

// 创建路由实例
const router = createRouter({
  history: createWebHistory(),  // 使用 HTML5 历史模式

  routes: [

    {
      path: "/dictionaryManager",
      name: "dictionaryManager",
      component: () => import('@/views/DictionaryManager.vue'),
    },
    {
      path: '/acct', // 账户页面
      name: 'Acct',
      component: () => import('@/views/Acct.vue'),
    },
    {
      path: '/acctBank', // 账户银行页面
      name: 'AcctBank',
      component: () => import('@/views/AcctBank.vue'),
    },
    {
      path: '/merchant', // 商户页面
      name: 'Merchant',
      component: () => import('@/views/Merchant.vue'),
    },
    {
      path: '/cust', // 客户页面
      name: 'Cust',
      component: () => import('@/views/Cust.vue'),
    },
    {
      path: '/bankAccount', // 银行账户页面
      name: 'BankAccount',
      component: () => import('@/views/BankAccount.vue'),
    },
    {
      path: '/invLoc', // 库存位置页面
      name: 'InvLoc',
      component: () => import('@/views/InvLoc.vue'),
    },
    {
      path: '/paymentMethod', // 支付方式页面
      name: 'PaymentMethod',
      component: () => import('@/views/PaymentMethod.vue'),
    },
    {
      path: '/cat', // 分类页面
      name: 'Cat',
      component: () => import('@/views/Cat.vue'),
    },
    {
      path: '/brand', // 品牌页面
      name: 'Brand',
      component: () => import('@/views/Brand.vue'),
    },
    {
      path: '/spec', // 规格页面
      name: 'Spec',
      component: () => import('@/views/Spec.vue'),
    },
    {
      path: '/empl', // 员工页面
      name: 'Empl',
      component: () => import('@/views/Empl.vue'),
    },
    {
      path: '/prdt', // 产品页面
      name: 'Prdt',
      component: () => import('@/views/Prdt.vue'),
    },
    {
      path: '/loadingInfo', // 卸载页面
      name: 'LoadingInfo',
      component: () => import('@/views/LoadingInfo.vue'),
    },
    {
      path: '/cost', // 成本页面
      name: 'Cost',
      component: () => import('@/views/Cost.vue'),
    },
    {
      path: '/sale', // 销售页面
      name: 'Sale',
      component: () => import('@/views/Sale.vue'),
    },
    {
      path: '/buy', // 购买页面
      name: 'Buy',
      component: () => import('@/views/Buy.vue'),
    },
    {
      path: '/send', // 发送页面
      name: 'Send',
      component: () => import('@/views/Send.vue'),
    },
    {
      path: '/purrec', // 采购记录页面
      name: 'Purrec',
      component: () => import('@/views/Purrec.vue'),
    },
    {
      path: '/shouldIn', // 应收页面
      name: 'ShouldIn',
      component: () => import('@/views/ShouldIn.vue'),
    },
    {
      path: '/shouldOut', // 应付页面
      name: 'ShouldOut',
      component: () => import('@/views/ShouldOut.vue'),
    },
    {
      path: '/in', // 入库页面
      name: 'In',
      component: () => import('@/views/In.vue'),
    },
    {
      path: '/out', // 出库页面
      name: 'Out',
      component: () => import('@/views/Out.vue'),
    },
    {
      path: '/',  // 显示登录页面
      name: 'Login',
      component: Login,
    },
    {
      path: '/dictionaryManager',  // 显示登录页面
      name: 'dictionaryManager',
      component: DictionaryManager,
    }

  ]
});

export default router;

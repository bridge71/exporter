// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
// 导入views中的页面组件
import Login from '@/views/Login.vue';
<<<<<<< HEAD
import Home from '@/views/TestHome.vue';
import Acct from '@/views/Acct.vue';
import AcctBank from '@/views/AcctBank.vue';
import BankAccount from '@/views/BankAccount.vue';
import Brand from '@/views/Brand.vue';
import Buy from '@/views/Buy.vue';
import Cat from '@/views/Cat.vue';
import Cost from '@/views/Cost.vue';
import Cust from '@/views/Cust.vue';
import DictionaryManager from '@/views/DictionaryManager.vue';  // 新增导入
import Empl from '@/views/Empl.vue';
import In from '@/views/In.vue';
import InvLoc from '@/views/InvLoc.vue';
import Merchant from '@/views/Merchant.vue';
import Out from '@/views/Out.vue';
import PackSpec from '@/views/PackSpec.vue';
import Password from '@/views/Password.vue';
import PaymentMethod from '@/views/PaymentMethod.vue';
import Prdt from '@/views/Prdt.vue';
import Purrec from '@/views/Purrec.vue';
import Register from '@/views/Register.vue';
import Sale from '@/views/Sale.vue';
import Send from '@/views/Send.vue';
import ShouldIn from '@/views/ShouldIn.vue';
import ShouldOut from '@/views/ShouldOut.vue';
import Spec from '@/views/Spec.vue';
import Unload from '@/views/Unload.vue';
import Username from '@/views/Username.vue';
=======
>>>>>>> 707122bc314398eb5eb8f3097fdd1b4c50d72907

// 创建路由实例
const router = createRouter({
  history: createWebHistory(),  // 使用 HTML5 历史模式

  routes: [
<<<<<<< HEAD
    {
      path: '/',
      name: 'login',
      component: Login,
    },
    {
      path: '/acct',
      name: 'acct',
      component: Acct,
    },
    {
      path: '/acctBank',
      name: 'acctBank',
      component: AcctBank,
    },
    {
      path: '/bankAccount',
      name: 'bankAccount',
      component: BankAccount,
    },
    {
      path: '/brand',
      name: 'brand',
      component: Brand,
    },
    {
      path: '/buy',
      name: 'buy',
      component: Buy,
    },
    {
      path: '/cat',
      name: 'cat',
      component: Cat,
    },
    {
      path: '/cost',
      name: 'cost',
      component: Cost,
    },
    {
      path: '/cust',
      name: 'cust',
      component: Cust,
    },
    {
      path: '/dictionaryManager',  // 新增路由配置
      name: 'dictionaryManager',
      component: DictionaryManager,
    },
    {
      path: '/empl',
      name: 'empl',
      component: Empl,
    },
    {
      path: '/in',
      name: 'in',
      component: In,
    },
    {
      path: '/invLoc',
      name: 'invLoc',
      component: InvLoc,
    },
    {
      path: '/merchant',
      name: 'merchant',
      component: Merchant,
    },
    {
      path: '/out',
      name: 'out',
      component: Out,
    },
    {
      path: '/packSpec',
      name: 'packSpec',
      component: PackSpec,
    },
    {
      path: '/password',
      name: 'password',
      component: Password,
    },
    {
      path: '/paymentMethod',
      name: 'paymentMethod',
      component: PaymentMethod,
    },
    {
      path: '/prdt',
      name: 'prdt',
      component: Prdt,
    },
    {
      path: '/purrec',
      name: 'purrec',
      component: Purrec,
    },
    {
      path: '/register',
      name: 'register',
      component: Register,
    },
    {
      path: '/sale',
      name: 'sale',
      component: Sale,
    },
    {
      path: '/send',
      name: 'send',
      component: Send,
    },
    {
      path: '/shouldIn',
      name: 'shouldIn',
      component: ShouldIn,
    },
    {
      path: '/shouldOut',
      name: 'shouldOut',
      component: ShouldOut,
    },
    {
      path: '/spec',
      name: 'spec',
      component: Spec,
    },
    {
      path: '/unload',
      name: 'unload',
      component: Unload,
    },
    {
      path: '/username',
      name: 'username',
      component: Username,
    },
=======

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
      path: '/unload', // 卸载页面
      name: 'Unload',
      component: () => import('@/views/Unload.vue'),
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
      path: '/login',  // 显示登录页面
      name: 'Login',
      component: Login,
    },

>>>>>>> 707122bc314398eb5eb8f3097fdd1b4c50d72907
  ]
});

export default router;
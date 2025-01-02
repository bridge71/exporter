// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
// 导入views中的页面组件
import Login from '@/views/Login.vue';
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

// 创建路由实例
const router = createRouter({
  history: createWebHistory(),  // 使用 HTML5 历史模式

  routes: [
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
  ]
});

export default router;
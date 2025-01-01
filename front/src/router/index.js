// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
// 导入views中的页面组件
import Login from '@/views/Login.vue';
import Home from '@/views/TestHome.vue';
import Acct from '@/views/Acct.vue';
import AcctBank from '@/views/AcctBank.vue';
import TestHome from '@/views/TestHome.vue';
import Client from '@/views/Client.vue';
import Contact from '@/views/Contact.vue';
import BankAccount from '@/views/BankAccount.vue';
import InventoryLocation from '@/views/InventoryLocation.vue';
import PaymentMethod from '@/views/PaymentMethod.vue';
import Category from '@/views/Category.vue';
import Brand from '@/views/Brand.vue';
import PackagingSpec from '@/views/PackagingSpec.vue';
import Employee from '@/views/Employee.vue';
import ProductDetail from '@/views/ProductDetail.vue';
import LoadingDetail from '@/views/LoadingDetail.vue';
import CostDetail from '@/views/CostDetail.vue';
import SalesOrder from '@/views/SalesOrder.vue';
import PurchaseOrder from '@/views/PurchaseOrder.vue';
import PayableInvoice from '@/views/PayableInvoice.vue';
import Receipt from '@/views/Receipt.vue';
import Payment from '@/views/Payment.vue';
import Register from '@/views/Register.vue';

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
      path: '/clinet',
      name: 'Clinet',
      component: Client,

    },
    {
      path: '/login',  // 显示登录页面
      name: 'Login',
      component: Login,
    },
    {
      path: '/contact',
      name: 'contact',
      component: Contact,
    },
    {
      path: '/bankAccount',
      name: 'bankAccount',
      component: BankAccount,
    },
    {
      path: '/inventoryLocation',
      name: 'inventoryLocation',
      component: InventoryLocation,
    },
    {
      path: '/paymentMethod',
      name: 'paymentMethod',
      component: PaymentMethod,
    },
    {
      path: '/category',
      name: 'category',
      component: Category,
    },
    {
      path: '/brand',
      name: 'brand',
      component: Brand,
    },
    {
      path: '/packagingSpec',
      name: 'packagingSpec',
      component: PackagingSpec,
    },
    {
      path: '/employee',
      name: 'employee',
      component: Employee,
    },
    {
      path: '/productDetail',
      name: 'productDetail',
      component: ProductDetail,
    },
    {
      path: '/loadingDetail',
      name: 'loadingDetail',
      component: LoadingDetail,
    },
    {
      path: '/costDetail',
      name: 'costDetail',
      component: CostDetail,
    },
    {
      path: '/salesOrder',
      name: 'salesOrder',
      component: SalesOrder,
    },
    {
      path: '/purchaseOrder',
      name: 'purchaseOrder',
      component: PurchaseOrder,
    },
    {
      path: '/payableInvoice',
      name: 'payableInvoice',
      component: PayableInvoice,
    },
    {
      path: '/receipt',
      name: 'receipt',
      component: Receipt,
    },
    {
      path: '/payment',
      name: 'payment',
      component: Payment,
    },
    {
      path: '/register',
      name: 'register',
      component: Register,
    },
    
  ]
});

export default router;

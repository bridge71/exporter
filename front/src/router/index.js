// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
// 导入views中的页面组件
import Login from '@/views/Login.vue';
<<<<<<< HEAD
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
=======
>>>>>>> a5db68559becfe39342b7aa8b2004fa37d254cb6

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
      path: '/clinet',
      name: 'Clinet',
      component: Client,

    },
=======

    { path: '/', component: () => import('@/views/Acct.vue') },
    { path: '/acct', component: () => import('@/views/Acct.vue') },
    { path: '/acctBank', component: () => import('@/views/AcctBank.vue') },
    { path: '/merchant', component: () => import('@/views/Merchant.vue') },
    { path: '/cust', component: () => import('@/views/Cust.vue') },
    { path: '/bankAccount', component: () => import('@/views/BankAccount.vue') },
    { path: '/invLoc', component: () => import('@/views/InvLoc.vue') },
    { path: '/paymentMethod', component: () => import('@/views/PaymentMethod.vue') },
    { path: '/cat', component: () => import('@/views/Cat.vue') },
    { path: '/brand', component: () => import('@/views/Brand.vue') },
    { path: '/spec', component: () => import('@/views/Spec.vue') },
    { path: '/empl', component: () => import('@/views/Empl.vue') },
    { path: '/prdt', component: () => import('@/views/Prdt.vue') },
    { path: '/unload', component: () => import('@/views/Unload.vue') },
    { path: '/cost', component: () => import('@/views/Cost.vue') },
    { path: '/sale', component: () => import('@/views/Sale.vue') },
    { path: '/buy', component: () => import('@/views/Buy.vue') },
    { path: '/send', component: () => import('@/views/Send.vue') },
    { path: '/purrec', component: () => import('@/views/Purrec.vue') },
    { path: '/shouldIn', component: () => import('@/views/ShouldIn.vue') },
    { path: '/shouldOut', component: () => import('@/views/ShouldOut.vue') },
    { path: '/in', component: () => import('@/views/In.vue') },
    { path: '/out', component: () => import('@/views/Out.vue') },
>>>>>>> a5db68559becfe39342b7aa8b2004fa37d254cb6
    {
      path: '/login',  // 显示登录页面
      name: 'Login',
      component: Login,
    },
<<<<<<< HEAD
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
    
=======

>>>>>>> a5db68559becfe39342b7aa8b2004fa37d254cb6
  ]
});

export default router;

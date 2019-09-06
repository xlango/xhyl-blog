// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import iView from 'iview';
import axios from 'axios'
import {post,fetch,patch,put} from './request/http'
//定义全局变量
Vue.prototype.$post=post;
Vue.prototype.$fetch=fetch;
Vue.prototype.$patch=patch;
Vue.prototype.$put=put;
import 'iview/dist/styles/iview.css'
import 'iview/dist/iview.min.js'
import VueRouter  from 'vue-router'
import AddBlog from '@/components/AddBlog'
import Home from '@/components/Home'
import Login from '@/components/Login'
import Me from '@/components/Me'
import QuestionList from '@/components/QuestionList'
import Registe from '@/components/Registe'
import Archive from '@/components/Archive'
import Test from '@/components/Test'


//验证码
import SlideVerify from 'vue-monoplasty-slide-verify';


Vue.config.productionTip = false
Vue.use(iView)
Vue.use(VueRouter)
Vue.use(SlideVerify);


//配置路由
const router=new VueRouter({
  linkActiveClass:'router-link-active',
  routes:[
    {path:"/",component:Home},
    // {path:"/home",component:Home},
    {path:"/article/add",component:AddBlog,meta:{requireAuth:true}},
    // {path:"/article/list",component:BlogList},
    {path:"/login",component:Login},
    {path:"/registe",component:Registe},
    {path:"/me",component:Me,meta:{requireAuth:true}},
    {path:"/question/list",component:QuestionList},
    {path:"/test",component:Test},
    {path:"/archive",component:Archive},
  ],
  mode:"history"
})


router.beforeEach((to, from, next) => {
  console.log("跳转前-----------------",to.matched.some(res => res.meta.requireAuth))
  if (to.matched.some(res => res.meta.requireAuth)) { // 验证是否需要登陆
    if (localStorage.getItem('userId')) { // 查询本地存储信息是否已经登陆
      next();
    } else {
      next({
        path: '/login', // 未登录则跳转至login页面
        query: {redirect: to.fullPath} // 登陆成功后回到当前页面，这里传值给login页面，to.fullPath为当前点击的页面
        });
    }
  } else {
    next();
  }

});
/* eslint-disable no-new */
new Vue({
  router:router,
  el: '#app',
  components: { App },
  template: '<App/>'
})

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
    {path:"/article/add",component:AddBlog},
    // {path:"/article/list",component:BlogList},
    {path:"/login",component:Login},
    {path:"/registe",component:Registe},
    {path:"/me",component:Me},
    {path:"/question/list",component:QuestionList},
  ],
  mode:"history"
})

/* eslint-disable no-new */
new Vue({
  router:router,
  el: '#app',
  components: { App },
  template: '<App/>'
})

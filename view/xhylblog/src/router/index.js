import Vue from 'vue'
import Router from 'vue-router'
import { routers, base } from './config'
 
Vue.use(Router)
 
let route = []
routers.map(v => {
  v.child.map(k => {
    if (k.url) {
      return route.push({
        path: `${base}admin${k.url}`,
        component: () => import(`../pages${k.url}.vue`)
      })
    }
    if (k.child) {
      k.child.map(j => {
        route.push({
          path: `${base}admin${j.url}`,
          component: () => import(`../pages${j.url}.vue`)
        })
      })
    }
  })
})
 
let adminRoute = [
  {
    path: `${base}login`,
    component: () => import('../pages/login/index.vue')
  },
  {
    path: `${base}admin`,
    component: () => import('../pages/admin.vue'),
    children: route
  }
]
 
export default new Router({
  mode: 'history',
  routes: adminRoute
})
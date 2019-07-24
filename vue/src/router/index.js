import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Card from '@/components/Card'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: '主页',
      component: Home
    },
    {
      path: '/card',
      name: '卡片',
      component: Card
    }
  ]
})

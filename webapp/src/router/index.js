import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from '@/views/Home.vue'
import Database from '@/views/Database.vue'

import Actors from '@/components/Actors.vue'
import Actor from '@/components/Actor.vue'
import Classes from '@/views/Classes.vue'
import Skills from '@/views/Skills.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/database',
    name: '',
    component: Database,
    children: [
      { path: '', redirect: 'actors' },
      {
        path: 'actors',
        component: Actors,
        children: [{
          path: ':id',
          component: Actor
        }]
      },
      { path: 'classes', component: Classes },
      { path: 'skills', component: Skills }
    ]
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = new VueRouter({
  routes,
  linkActiveClass: 'is-active'
})

export default router

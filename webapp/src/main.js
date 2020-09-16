import Vue from 'vue'

import { sync } from 'vuex-router-sync'
import app from './App.vue'
import router from './router'
import store from './store'

import '@/assets/scss/main.scss'

sync(store, router)

Vue.config.productionTip = false

const vm = new Vue({
  router,
  store,
  render: h => h(app)
}).$mount('#app')

export { vm }

import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import { routes } from "./routes"

import Buefy from 'buefy'
Vue.use(Buefy)
Vue.use(VueRouter)

window.axios = require('axios')

const router = new VueRouter({
    mode: 'history',
    routes
})

new Vue({
    el: '#app',
    router,
    render: h => h(App)
})
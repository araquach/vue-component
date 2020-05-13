import Vue from 'vue'
import Vuex from 'vuex'

import auth from "./modules/auth"
import team from "./modules/team"

Vue.use(Vuex)

export const store = new Vuex.Store({
    modules: {
        auth,
        team
    }
})
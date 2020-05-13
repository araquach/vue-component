import Home from './components/home'
import About from './components/About'
import Login from './components/auth/Login'
import Register from  './components/auth/Register'
import Protected from './components/Protected'
import TeamAll from './components/team/TeamAll'
import TeamDetail from './components/team/TeamDetail'
import {store} from "./store/store";

export const routes = [
    {
        path: '',
        component: Home,
        name: 'home'
    },
    {
        path: '/login',
        component: Login,
        name: 'login'
    },
    {
        path: '/about',
        component: About,
        name: 'about'
    },
    {
        path: '/register',
        component: Register,
        name: 'register'
    },
    {
        path: '/team',
        component: TeamAll,
        name: 'team'
    },
    {
        path: '/team/:slug',
        component: TeamDetail,
        name: 'team-detail',
        params: true
    },
    {
        path: '/protected',
        component: Protected,
        name: 'protected',
        beforeEnter: (to, from, next) => {
            if (store.getters.isLoggedIn) {
                next()
                return
            }
            next('/')
        }}
]


import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import(/* webpackChunkName: "Home" */ '@/views/Home.vue')
    },
    {
        path: '/buyers',
        name: 'Buyers',
        component: () => import(/* webpackChunkName: "Buyer" */ '@/views/Buyers.vue')
    },
    {
        path: '/buyers/:buyerId/transactions',
        name: 'Transactions',
        component: () => import(/* webpackChunkName: "Transactions" */ '@/views/Transactions.vue')
    },
    {
        path: '*',
        name: 'NotFound',
        component: () => import(/* webpackChunkName: "NotFound" */ '@/views/NotFound.vue')
    },
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router

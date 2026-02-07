import {createRouter, createWebHistory} from "vue-router";
import {authAfterDenyRouter, authBeforeAllowRouter, baseRouters} from "./base.js";
import {localStore} from "../store/local.js";

const router = createRouter({
    history: createWebHistory(), routes: baseRouters, scrollBehavior: () => ({left: 0, top: 0}),
})

router.beforeEach(async (to, from, next) => {
    let token = localStore().auth.token
    if (token) {
        if (authAfterDenyRouter.indexOf(to.name) >= 0) {
            next({name: baseRouters[2].name})
        } else {
            next()
        }
    } else {
        // 未登录
        if (authBeforeAllowRouter.indexOf(to.name) >= 0) {
            next()
        } else {
            next({name: 'Login'})
        }
    }
    localStore().menuKey = to.name
})

router.afterEach((to) => {
})

export default router


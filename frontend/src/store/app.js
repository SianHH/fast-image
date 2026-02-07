import {defineStore} from "pinia";

export const appStore = defineStore('app', {
    state: () => ({
        width: window.innerWidth,
        height: window.innerHeight,
        auth: {
            token: '',
            expAt: 0,
        },
        contentStyle: {
            maxWidth: '980px',
            margin: 'auto',
        },
    }),
    getters: {
        drawerWidthAdapter: (state) => {
            if (state?.width < 500) {
                return state?.width
            }
            return 400
        }
    },
    actions: {
        async refreshUserInfo() {

        }
    },
    persist: false
})
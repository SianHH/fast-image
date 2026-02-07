import {defineStore} from "pinia";

export const localStore = defineStore('local', {
    state: () => ({
        darkTheme: false,
        layout: 'simple',
        copyType: 'md',
        isCollapsed: true,
        menuKey: '',
        auth: {
            token: '',
            expAt: ''
        },
    }),
    getters: {},
    actions: {
        setLayout(name) {
            this.layout = name
        }
    },
    persist: true
})
export const baseRouters = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('../views/auth/login/index.vue'),
        meta: {
            title: '登录',
            layout: 'empty',
            hidden: 1,
        }
    },
    {
        path: '/image/upload',
        name: 'ImageUpload',
        component: () => import('../views/image/index.vue'),
        meta: {
            title: '上传图片',
            layout: 'simple',
            hidden: 1,
            icon: '<svg t="1769082833717" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5685"><path d="M433.1 480 174.9 480c-25.9 0-46.9-21-46.9-46.9L128 174.9c0-25.9 21-46.9 46.9-46.9l258.2 0c25.9 0 46.9 21 46.9 46.9l0 258.2C480 459 459 480 433.1 480z" p-id="5686"></path><path d="M433.1 896 174.9 896c-25.9 0-46.9-21-46.9-46.9L128 590.9c0-25.9 21-46.9 46.9-46.9l258.2 0c25.9 0 46.9 21 46.9 46.9l0 258.2C480 875 459 896 433.1 896z" p-id="5687"></path><path d="M849.1 480 590.9 480c-25.9 0-46.9-21-46.9-46.9L544 174.9c0-25.9 21-46.9 46.9-46.9l258.2 0c25.9 0 46.9 21 46.9 46.9l0 258.2C896 459 875 480 849.1 480z" p-id="5688"></path><path d="M849.1 896 590.9 896c-25.9 0-46.9-21-46.9-46.9L544 590.9c0-25.9 21-46.9 46.9-46.9l258.2 0c25.9 0 46.9 21 46.9 46.9l0 258.2C896 875 875 896 849.1 896z" p-id="5689"></path></svg>',
        }
    },
    {
        path: '/403',
        name: '403',
        component: () => import('../views/public/403.vue'),
        meta: {
            title: '403',
            layout: 'empty',
            hidden: 1,
        }
    },
    {
        path: '/:pathMatch(.*)',
        name: '404',
        component: () => import('../views/public/404.vue'),
        meta: {
            title: '404',
            layout: 'empty',
        }
    }
]

const findRouterName = (router) => {
    if (!router) {
        return []
    }
    let result = []
    for (let i = 0; i < router?.length; i++) {
        let r = router[i]
        // if (r?.meta?.hidden !== 2) {
        //     continue
        // }
        let children = findRouterName(r.children)
        if (children?.length > 0) {
            result.push(...children)
        } else {
            result.push(r.name)
        }
    }
    return result
}

// 授权前，可以访问的页面
export const authBeforeAllowRouter = ['Login', '403', '404']
// 授权后，无法访问的页面
export const authAfterDenyRouter = ['Login']

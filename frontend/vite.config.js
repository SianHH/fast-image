import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
    plugins: [
        vue(),
    ],
    server: {
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:8083',
                changeOrigin: true,
                ws: true
            }
        },
        host: '0.0.0.0'
    },
    build: {
        // 代码分割配置
        rollupOptions: {
            output: {
                // 手动指定块分割策略
                manualChunks(id) {
                    // 将 node_modules 中的包分割成单独的 chunk
                    if (id.includes('node_modules')) {
                        // 按包名分割
                        if (id.includes('lodash')) {
                            return 'vendor-lodash'
                        }
                        if (id.includes('axios')) {
                            return 'vendor-axios'
                        }

                        if (id.includes('naive-ui')) {
                            return 'vendor-naive-ui'
                        }

                        if (id.includes('echarts')) {
                            return 'vendor-echarts'
                        }

                        if (id.includes('chart.js')) {
                            return 'vendor-chartjs'
                        }

                        if (id.includes('md-editor-v3')) {
                            return 'vendor-md-editor'
                        }

                        if (id.includes('/vue/') || id.includes('@vue/')) {
                            return 'vendor-vue'
                        }

                        if (id.includes('pinia')) {
                            return 'vendor-pinia'
                        }

                        if (id.includes('moment') || id.includes('dayjs')) {
                            return 'vendor-date'
                        }

                        if (id.includes('qrcode')) {
                            return 'vendor-qrcode'
                        }

                        if (id.includes('devicon') || id.includes('vicons') || id.includes('flag-icons')) {
                            return 'vendor-icons'
                        }

                        // 其他第三方依赖
                        return 'vendor'
                    }

                    // // 按业务模块分割
                    // if (id.includes('src/components')) {
                    //     return 'components'
                    // }

                    // if (id.includes('src/utils')) {
                    //     return 'utils'
                    // }
                },

                // 块文件名格式
                chunkFileNames: 'assets/js/[name]-[hash].js',
                entryFileNames: 'assets/js/[name]-[hash].js',
                assetFileNames: 'assets/[ext]/[name]-[hash].[ext]'
            }
        },

        // 启用/禁用代码分割
        cssCodeSplit: true,

        // 目标浏览器
        target: 'esnext',

        // 模块转换目标
        modulePreload: {
            polyfill: true
        },

        // 限制 chunk 大小警告
        chunkSizeWarningLimit: 200
    }
})

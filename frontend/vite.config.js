import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            '~': `${path.resolve(__dirname, 'src')}`,
        }
    },
    css: {
        preprocessorOptions:
            {
                scss: {
                    charset: false,
                }
            }
    },
    plugins: [vue()]
})

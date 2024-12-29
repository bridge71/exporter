import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
  server: {
    port: 7777,
    host: '0.0.0.0',
    open: true,
    strictPort: false,
    https: false,
    proxy: {
      '/find': {
        target: 'http://10.111.3.44:6767',
        changeOrigin: true,
        secure: false
      },

      '/save': {
        target: 'http://10.111.3.44:6767',
        changeOrigin: true,
        secure: false
      },
      '/auth': {
        target: 'http://10.111.3.44:6767',
        changeOrigin: true,
        secure: false
      },
    }
  },
})

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'



// https://vite.dev/config/
export default defineConfig(({mode})=>{
    const root = process.cwd()
    const env = loadEnv(mode,root)
    console.log(env)
    return {
        plugins: [
      vue(),
      AutoImport({
        resolvers: [ElementPlusResolver()],
      }),
      Components({
        resolvers: [ElementPlusResolver()],
      }),
  ],
  base: env.VITE_NODE_ENV === 'production' ? env.VITE_APP_PATH : '/',
    }

})

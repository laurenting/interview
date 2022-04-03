import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import * as ElIcons from '@element-plus/icons-vue'
import '~/styles/index.scss'

const app = createApp(App)
app
  .use(router)
  .use(ElementPlus, {
    size: 'mini',
    zIndex: 3000
  })
  .mount('#app')

// 统一注册Icon图标
for (const iconName in ElIcons) {
  if (Reflect.has(ElIcons, iconName)) {
    const item = ElIcons[iconName]
    app.component(iconName, item)
  }
}

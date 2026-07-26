import './assets/main.css'
import "tailwindcss";
// Note: base.css is still imported via main.css @import chain but overridden

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')

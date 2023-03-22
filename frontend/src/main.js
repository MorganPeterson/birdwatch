import { createApp } from 'vue'
import { plugin, defaultConfig } from '@formkit/vue'

import App from './App.vue'

import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap'

import './assets/main.css'

const app = createApp(App)

app.provide('axios', app.config.globalProperties.axios)

app.use(plugin, defaultConfig)

app.mount('#app')

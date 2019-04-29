import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './registerServiceWorker'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import {leancloud} from "./assets/js/config";
import AV from 'leancloud-storage'
import  {i18n} from "./i18n";


AV.init({
  appId: leancloud.APP_ID,
  appKey: leancloud.APP_KEY
})

Vue.config.productionTip = false

Vue.use(ElementUI)
Vue.use(mavonEditor)


new Vue({
  router,
  i18n,
  store,
  render: h => h(App)
}).$mount('#app')

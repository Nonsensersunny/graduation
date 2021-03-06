import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './registerServiceWorker'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import {i18n} from "./i18n"
import AV from './assets/js/av'
import {leancloud} from "@/assets/js/config"
import VueSocketIO from 'vue-socket.io'

AV.init({
  appId: leancloud.APP_ID,
  appKey: leancloud.APP_KEY
})

router.beforeEach((to, from, next) => {
  if (!store.state.isLogin && to.path != '/signin' && to.path != '/') {
    ElementUI.Message.error(i18n.t("message.common.login_first"))
    next()
  }
  next()
})

Vue.config.productionTip = false

Vue.use(ElementUI)
Vue.use(mavonEditor)
// Vue.use(new VueSocketIO({
//   debug: true,
//   connection: 'http://127.0.0.1:8100',
//   vuex: {
//     store,
//     actionPrefix: 'SOCKET_',
//     mutaionPrefix: 'SOCKET_'
//   },
//   options: {
//     // path: '/u/ws'
//   }
// }))

new Vue({
  router,
  i18n,
  store,
  render: h => h(App)
}).$mount('#app')

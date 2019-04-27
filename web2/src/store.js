import Vue from 'vue'
import Vuex from 'vuex'
import {GuestHttp, UserHttp} from "@/assets/js/api";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isLogin: false,
    contents: [],
    rankeduser: [],
    profile: {}
  },
  getters: {
    rankedUsers(state) { return state.rankeduser },
    profile(state) { return state.profile }
  },
  mutations: {

  },
  actions: {
    async checkLoginStatus(context) {
      let resp = await UserHttp.checkLoginStatus()
      if (resp) {
        context.state.profile = resp
        context.state.isLogin = true
      } else {
        context.state.profile = {}
        context.state.isLogin = false
      }
    },
    async userLogin(context, playload) {
      context.state.profile = await GuestHttp.signin(playload);
    },
    async userRegister(context, playload) {
      await GuestHttp.signup(playload);
    },
    async getRankedUsers(context) {
      let resp = await UserHttp.getRankedUsers()
      context.state.rankeduser = resp
    },
    async getContentById(context, id) {
      let resp = await GuestHttp.getContentById(id);
      return resp;
    },
    async createContent(context, {title, author, category, content}) {
      await UserHttp.createContent(title, author, category, content)
    },
    async getRankedContent(context) {
      return await GuestHttp.getRankedContent()
    },
    async getContentByCat(context, cat) {
      return await GuestHttp.getContentByCat(cat)
    },
    async getTopContent(context) {
      return await UserHttp.getTopContent()
    }
  }
})

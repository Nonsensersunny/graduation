import Vue from 'vue'
import Vuex from 'vuex'
import {GuestHttp, UserHttp} from "@/assets/js/api";
import {langs, i18n} from "./i18n";

Vue.use(Vuex)


export default new Vuex.Store({
  state: {
    isLogin: false,
    contents: [],
    rankeduser: [],
    profile: {},
    lang: "en",
    langs: langs,
    avatar: '',
    cats: []
  },
  getters: {
    rankedUsers(state) { return state.rankeduser },
    profile(state) { return state.profile },
    selectedLang(state) { return state.lang },
    langs(state) { return state.langs },
    categories(state) { return state.cats }
  },
  mutations: {
    selectLang(state, lang) {
      state.lang = lang
      i18n.locale = lang
    }
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
      context.state.isLogin = true
    },
    async userRegister(context, playload) {
      await GuestHttp.signup(playload);
    },
    async userLogout(context) {
      await UserHttp.userLogout(context.state.profile.username)
      context.state.isLogin = false
      context.state.profile = {}
    },
    async getProfileByName(context, username) {
      return await GuestHttp.getProfileByName(username);
    },
    async getProfileById(context, id) {
      return await GuestHttp.getProfileById(id);
    },
    async getRankedUsers(context) {
      context.state.rankeduser = await UserHttp.getRankedUsers()
    },
    async getContentById(context, id) {
      return await GuestHttp.getContentById(id);
    },
    async createContent(context, {title, author, category, content}) {
      return await UserHttp.createContent(title, author, category, content);
    },
    async getRankedContent(context) {
      return await GuestHttp.getRankedContent()
    },
    async getContentByCat(context, cat) {
      return await GuestHttp.getContentByCat(cat)
    },
    async getTopContent(context) {
      return await UserHttp.getTopContent()
    },
    async updateUserProfile(context, user) {
      context.state.profile = await UserHttp.updateUserProfile(user)
    },
    async getCategories(context) {
      context.state.cats = await GuestHttp.getCategories()
    }
  }
})

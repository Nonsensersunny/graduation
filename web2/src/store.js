import Vue from 'vue'
import Vuex from 'vuex'
import {GuestHttp} from "@/assets/api";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isLogin: false,
    username: '',
  },
  mutations: {

  },
  actions: {
    async userLogin(context, playload) {
      await GuestHttp.signin(playload);
      context.state.username = playload.username;
      context.state.isLogin = true;

    },
    async userRegister(context, playload) {
      await GuestHttp.signup(playload);
    }
  }
})

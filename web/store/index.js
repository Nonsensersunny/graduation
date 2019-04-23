import {GuestHttp, UserHttp} from "../assets/api";

export const state = () => ({

})

export const mutations = {

}

export const actions = {
    async userLogin(context, playload) {
        await GuestHttp.signin(playload);
        context.state.username = playload.username;
        context.state.isLogin = true;

    }
}


import axios from "../.nuxt/axios";

export const GuestHttp = {
  client: axios.create({
    baseURL: '/g'
  }),
  async signin(user) {
    let data = new FormData();
    data.set('username', user.username);
    data.set('password', user.password);
    let res = this.client.post()
  },
  async signup(user) {
    let data = new FormData();
    data.set('username', user.username);
    data.set('password', user.password);
    let res = this.client.post()
  }
}

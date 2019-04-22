import axios from "../.nuxt/axios";

import {ErrorCode, RespError} from "type"

export const GuestHttp = {
  client: axios.create({
    baseURL: '/g'
  }),
  async signin(user) {
    let data = new FormData();
    data.set('username', user.username);
    data.set('password', user.password);
    let res = this.client.post('/register', user)
  },
  async signup(user) {
    let data = new FormData();
    data.set('username', user.username);
    data.set('password', user.password);
    let res = this.client.post('/login', user)
  }
}

export const UserHttp = {
  client: axios.create({
    baseURL: '/u',
  }),

}

function errorHandler(e) {
  if (!e.response) {
    console.error(e);
    let error = new RespError(0, ErrorCode.Undefined, '');
    return Promise.reject(error);
  }
  let resp = e.response;
  if (resp.status >= 500) {
    let error = new RespError(resp.status, ErrorCode.ServerError, '');
    return Promise.reject(error);
  }
  let errorBody = resp.data.error;
  if (!errorBody || !errorBody.code) {
    let error = new RespError(resp.status, ErrorCode.Undefined, resp.data);
    return Promise.reject(error);
  }
  let error = new RespError(resp.status, errorBody.code, resp.msg);
  return Promise.reject(error);
}

GuestHttp.client.interceptors.response.use(resp => resp, errorHandler);
UserHttp.client.interceptor.response.use(resp => resp, errorHandler);

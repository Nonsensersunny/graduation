import axios from "axios";

import {ErrorCode, RespError, Content} from "@/assets/type"

// axios.defaults.withCredentials = true;

export const GuestHttp = {
  client: axios.create({
    baseURL: 'http://127.0.0.1:8888/g',
    withCredentials:true
  }),
  async signin(user) {
    let data = new FormData();
    data.set('username', user.username);
    data.set('password', user.password);
    console.log(data)
    let res = await this.client.post('/login', user);
    return res;
  },
  async signup(user) {
    let data = new FormData();
    data.set('username', user.username);
    data.set('password', user.password);
    let res = await this.client.post('/register', user);
    return res;
  }
}

export const UserHttp = {
  async createContent(title, author, category, content) {
    let data = new Content(title, author, category, content);
    let resp = this.client.post('/content', data)
    return resp['message']
  },
  client: axios.create({
    baseURL: 'http://127.0.0.1:8888/u',
    withCredentials:true
  }),
  async getContentById(id) {
    let resp = await this.client.get(`/content/${id}`);
    return resp.data["content"]
  }
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
UserHttp.client.interceptors.response.use(resp => resp, errorHandler);
// GuestHttp.client.request({withCredentials: true})
// UserHttp.client.request({withCredentials: true})
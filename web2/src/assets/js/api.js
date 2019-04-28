import axios from "axios";

import {ErrorCode, RespError, Content} from "@/assets/js/type"
import {server} from "@/assets/js/config";

// axios.defaults.withCredentials = true;

export const GuestHttp = {
  client: axios.create({
    baseURL: `http://${server.host}:${server.port}/g`,
    withCredentials:true
  }),
  async signin(user) {
    let data = new FormData();
    data.set('username', user.username);
    data.set('password', user.password);
    let res = await this.client.post('/login', user);
    return res.data.data["data"];
  },
  async signup(user) {
    let data = new FormData();
    data.set('username', user.username);
    data.set('password', user.password);
    let res = await this.client.post('/register', user);
    return res;
  },
  async getContentById(id) {
    let resp = await this.client.get(`/content/${id}`);
    return resp.data.data["data"]
  },
  async getContentByCat(cat) {
    let resp = await this.client.get(`/contents/${cat}`)
    return resp.data.data['data']
  },
  async getRankedContent() {
    let resp = await this.client.get('/contents');
    return resp.data.data['data']
  },
  async getProfileByName(username) {
    let resp = await this.client.get(`/profile/username/${username}`);
    return resp.data.data['data']
  },
  async getProfileById(id) {
    let resp = await this.client.get(`/profile/id/${id}`);
    return resp.data.data['data']
  },
}

export const UserHttp = {
  client: axios.create({
    baseURL: `http://${server.host}:${server.port}/u`,
    withCredentials:true
  }),
  async checkLoginStatus() {
    let resp = await this.client.get("/status");
    return resp.data.data["data"]
  },
  async createContent(title, author, category, content) {
    let data = new Content(title, author, category, content);
    let resp = await this.client.post('/content', data)
    return resp.data.data['status']
  },
  async getRankedUsers() {
    let resp = await this.client.get("/rank");
    return resp.data.data["users"]
  },
  async getTopContent() {
    let resp = await this.client.get("/contents");
    return resp.data.data["data"]
  },
  async updateUserProfile(user) {
    let resp = await this.client.post(user);
    return resp.data.data["data"]
  }
}

function errorHandler(e) {
  console.log(e)
  if (!e.response) {
    console.error(e);
    let error = new RespError(-1, ErrorCode.Undefined, '');
    return Promise.reject(error);
  }
  let response = e.response;
  if (response.code != 0) {
    let error = new RespError(response.code, response.error, '')
    return Promise.reject(error)
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
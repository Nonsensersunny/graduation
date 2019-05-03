import axios from "axios";

import {ErrorCode, RespError, Content} from "@/assets/js/type"
import {server} from "@/assets/js/config";

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
    return await this.client.post('/register', user);
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
  async getCategories() {
    let resp = await this.client.get('/cats');
    return resp.data.data['data']
  }
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
    let resp = await this.client.post("/profile/update", user);
    return resp.data.data["data"]
  },
  async userLogout(username) {
    await this.client.get(`/logout/${username}`)
  },
  async createComment(comment) {
    let resp = await this.client.post("/comment", comment);
    return resp.data.data["data"]
  },
  async userGetContentById(cid, uid) {
    let resp = await this.client.get(`/content/${cid}/${uid}`);
    return resp.data.data["data"]
  },
  async createFavorite(uid, cid) {
    let resp = await this.client.get(`/fav/${uid}/${cid}`);
    return resp.data.data["status"]
  },
  async deleteFavorite(uid, cid) {
    let resp = await this.client.get(`/vaf/${uid}/${cid}`);
    return resp.data.data["status"]
  },
  async createLink(link) {
    let resp = await this.client.post("/link", link);
    return resp.data.data["status"]
  },
  async deleteLink(id) {
    let resp = await this.client.get(`/link/${id}`);
    return resp.data.data["status"]
  },
  async getLinksByUserId(id) {
    let resp = await this.client.get(`/links/${id}`);
    return resp.data.data["data"]
  }
}

function errorHandler(e) {
  console.dir(e)
  if (!e.response) {
    console.error("response",e);
    let error = new RespError(-1, ErrorCode.Undefined, '');
    return Promise.reject(error);
  }
  let response = e.response;
  if (response.data.code != 0) {
    let error = new RespError(response.code, response.error, '')
    return Promise.reject(error)
  }
}

GuestHttp.client.interceptors.response.use(resp => resp, errorHandler);
UserHttp.client.interceptors.response.use(resp => resp, errorHandler);
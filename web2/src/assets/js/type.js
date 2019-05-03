export class RespError {
  constructor(error, code, data) {
    this.error = error;
    this.code = code;
    this.data = data;
  }
}

export const ErrorCode = Object.freeze({
  Undefined: 0,
  UnknownError: -1,
  ServerError: 1,
  BadParams: 2,
  RegisterFail: 1000,
  LoginFail: 1001,

  InvalidContent: 2000,
  CreateContentFail: 2001,
  FetchContentFail: 2002,

  InvalidPathPara: 3000,
  InvalidFormPara: 3001
});

export class Content {
  constructor(title, author, category, content) {
    this.title = title;
    this.author = author;
    this.category = category;
    this.content = content;
  }
}

export class User {
  constructor() {
    this.id = 0
    this.username = ''
    this.password = ''
    this.avatar = ''
    this.description = ''
    this.mail = ''
    this.sex = 0
  }
}

export class Comment {
  constructor(from, to, cid, content) {
    this.from = from
    this.to = to
    this.content = content
    this.cid = cid
  }
}

export class Link {
  constructor(uid, name, href) {
    this.uid = uid
    this.name = name
    this.href = href
  }
}
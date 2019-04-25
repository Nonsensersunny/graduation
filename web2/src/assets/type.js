export class RespError {
  constructor(status, code, msg) {
    this.status = status;
    this.code = code;
    this.msg = msg;
  }
}

export const ErrorCode = Object.freeze({
  Undefined: 0,
  ServerError: 1,
  BadParams: 2,
});

export class Content {
  constructor(title, author, category, content) {
    this.title = title;
    this.author = author;
    this.category = category;
    this.content = content;
  }
}
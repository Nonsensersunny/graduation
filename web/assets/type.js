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

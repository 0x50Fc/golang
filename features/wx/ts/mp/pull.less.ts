
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Scope } from '../Scope';
import { User } from '../User';
import { MessageType } from '../MessageType';

/**
 * 拉去关注微信公众号的用户
 * @method GET
 */
export interface Request {

  /**
   * appid
   */
  appid: string

}

export interface WXMPPullData {
  /**
   * 拉取的数量
   */
  count: int32
}

export interface Response extends BaseResponse {
  data?: WXMPPullData
}

export function handle(req: Request): Response {
  return {
    errno: ErrCode.OK
  }
}

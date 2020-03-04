
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Scope } from '../Scope';
import { User } from '../User';
import { MessageType } from '../MessageType';

/**
 * 上传媒体文件
 * @method POST
 */
export interface Request {

  /**
   * appid
   */
  appid: string

  /**
   * 文件类型, 默认 image
   */
  type?: string

  /**
   * 文件名
   */
  name: string

  /**
   * 文件内容 base64
   */
  content: string

}

export interface AppUpData {
  /**
   * 文件类型
   */
  type: string
  /**
   * 媒体标示
   */
  media_id: string
  /**
   * 创建时间
   */
  created_at: int64
}

export interface Response extends BaseResponse {
  data?: AppUpData
}

export function handle(req: Request): Response {
  return {
    errno: ErrCode.OK
  }
}

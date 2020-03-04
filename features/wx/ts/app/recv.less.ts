
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Scope } from '../Scope';
import { User } from '../User';
import { MessageType } from '../MessageType';

/**
 * 接收消息
 * @method POST
 */
export interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * 内容类型 json/xml 默认 json
     */
    type?: string

    /**
     * encodingKey
     */
    encodingKey: string

    /**
     * echostr
     */
    echostr: string

    /**
     * nonce
     */
    nonce: string

    /**
     * timestamp
     */
    timestamp: string

    /**
     * signature
     */
    signature: string

    /**
     * token
     */
    token: string

    /**
     * 内容
     */
    content: string

}

export interface Response extends BaseResponse {
    data?: any
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}


import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Scope } from '../Scope';
import { User } from '../User';
import { MessageType } from '../MessageType';

/**
 * 开放平台编码消息
 * @method POST
 */
export interface Request {

    /**
     * Token
     */
    token: string

    /**
     * encodingKey
     */
    encodingKey: string

    /**
     * nonce
     */
    nonce?: string

    /**
     * timestamp
     */
    timestamp?: string

    /**
     * 内容 JSON
     */
    content: string

}

export interface OpenEncryptData {

    /**
     * nonce
     */
    nonce: string

    /**
     * timestamp
     */
    timestamp: string

    /**
     * 签名
     */
    signature: string

    /**
     * 编码后内容
     */
    content: string
}

export interface Response extends BaseResponse {
    data?: OpenEncryptData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

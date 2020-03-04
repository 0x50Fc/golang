
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Scope } from '../Scope';
import { Ticket, TicketType } from '../Ticket';

/**
 * 开发平台获取JS签名配置信息
 * @method GET
 */
export interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * noncestr 不存在是自动生成
     */
    noncestr?: string

    /**
     * noncestr 不存时是自动生成
     */
    timestamp?: int64

    /**
     * 签名URL
     */
    url: string
}

export interface OpenConfigData {
    /**
     * appid
     */
    appid: string

    /**
     * timestamp
     */
    timestamp: int64

    /**
     * nonceStr
     */
    nonceStr: string

    /**
     * signature
     */
    signature: string

}


export interface Response extends BaseResponse {
    ticket?: Ticket
    data?: OpenConfigData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

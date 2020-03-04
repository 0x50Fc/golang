
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Scope } from '../Scope';

/**
 * 开发平台获取授权URL
 * @method GET
 */
export interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * state
     */
    state?: string

    /**
     * scope
     */
    scope?: Scope

    /**
     * redirect_uri
     */
    redirect_uri: string

}


export interface OpenAuthorizeData {
    /**
     * 授权URL
     */
    url: string
}


export interface Response extends BaseResponse {
    data?: OpenAuthorizeData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

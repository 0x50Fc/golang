
import { BaseResponse, ErrCode } from "../../lib/BaseResponse"
import { int64, int32 } from "../../lib/less";
import { Scope } from '../../Scope';

/**
 * 开发平台 公众号授权 获取授权URL
 * @method GET
 */
export interface Request {

    /**
     * 授权方式
     */
    openType?: OpenMPOpenType

    /**
     * 授权类型
     */
    authType?: OpenMPAuthType

    /**
     * 公众号/小程序 Appid
     */
    appid?: string

    /**
     * redirect_uri
     */
    redirect_uri: string

}

export enum OpenMPAuthType {
    MP = 1,
    App = 2,
    MP_APP = 3
}

export enum OpenMPOpenType {
    WX = 1,
    Web = 2
}

export interface OpenMPAuthorizeData {
    /**
     * 授权URL
     */
    url: string
}


export interface Response extends BaseResponse {
    data?: OpenMPAuthorizeData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

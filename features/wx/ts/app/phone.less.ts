
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { User } from '../User';

/**
 * 小程序获取手机号
 * @method POST
 */
export interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * openid
     */
    openid?: string

    /**
     * session_key
     */
    session_key?: string

    /**
     * 用户信息编码数据
     */
    encryptedData: string

    /**
     * 加密算法的初始向量
     */
    iv: string

}

export interface AppPhoneData {

    /**
     * 手机号
     */
    phone: string

    /**
     * 国家区号
     */
    country: string

    /**
     * 用户
     */
    user?: User
}

export interface Response extends BaseResponse {
    data?: AppPhoneData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

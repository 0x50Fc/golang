
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { User } from '../User';

/**
 * 小程序登录
 * @method POST
 */
export interface Request {

    /**
     * appid
     */
    appid: string

    /**
     * js_code
     */
    code: string

    /**
     * 用户信息编码数据
     */
    encryptedData?: string

    /**
     * 加密算法的初始向量
     */
    iv?: string
}

export interface AppLoginData {

    /**
     * session_key
     */
    session_key: string

    /**
     * 用户
     */
    user: User
}


export interface Response extends BaseResponse {
    data?: AppLoginData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}


import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { User } from "./User";
import { int64, int32 } from "./lib/less";

/**
 * 用户数量
 * @method GET
 */
export interface Request {

    /**
     * 用户ID,逗号分割
     */
    ids?: string

    /**
     * 用户名
     */
    name?: string

    /**
     * 昵称
     */
    nick?: string

    /**
     * 模糊匹配关键字
     */
    q?: string

    /**
     * 用户名前缀
     */
    prefix?: string

    /**
     * 用户名后缀
     */
    suffix?: string

}


export interface CountData {
    /**
     * 总记录数
     */
    total: int32
}


export interface Response extends BaseResponse {
    data?: CountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

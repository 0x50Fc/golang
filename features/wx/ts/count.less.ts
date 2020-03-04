
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { UserType } from './UserType';

/**
 * 数量
 * @method GET
 */
export interface Request {

    /**
     * 用户ID
     */
    uid?: int64

    /**
     * 类型,多个逗号分割
     */
    type?: string

    /**
     * appid
     */
    appid?: string

    /**
     * openid
     */
    openid?: string

    /**
     * uniqueid
     */
    unionid?: string

    /**
     * 模糊匹配关键字
     */
    q?: string

    /**
     * 状态 多个逗号分割
     */
    state?: string

    /**
     * 是否绑定
     */
    bind?: boolean

    /**
     * 是否有用户信息
     */
    info?: boolean

    /**
     * 绑定开始时间
     */
    startTime?: int32

    /**
     * 绑定结束时间
     */
    endTime?: int32

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

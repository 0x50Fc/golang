
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { User } from './User';
import { Page } from './Page';

/**
 * 查询
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
     * unionid
     */
    unionid?: string

    /**
     * 状态 多个逗号分割
     */
    state?: string

    /**
     * 是否绑定
     */
    bind?: boolean

    /**
     * 绑定开始时间
     */
    startTime?: int32

    /**
     * 绑定结束时间
     */
    endTime?: int32

    /**
     * 是否有用户信息
     */
    info?: boolean

    /**
     * 模糊匹配关键字
     */
    q?: string

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface QueryData {
    /**
     * 用户
     */
    items: User[]

    /**
     * 分页
     */
    page?: Page
}


export interface Response extends BaseResponse {
    data?: QueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

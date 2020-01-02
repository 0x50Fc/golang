
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { User } from "./User";
import { int64, int32 } from "./lib/less";

/**
 * 查询用户
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

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface QueryDataPage {
    /**
     * 分页位置
     */
    p: int32
    /**
    * 单页记录数
    */
    n: int32
    /**
     * 总页数
     */
    count: int32
    /**
     * 总记录数
     */
    total: int32
}

export interface QueryData {
    /**
     * 用户
     */
    items: User[]

    /**
     * 分页
     */
    page?: QueryDataPage
}


export interface Response extends BaseResponse {
    data?: QueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

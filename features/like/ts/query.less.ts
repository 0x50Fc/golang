
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { Like } from "./Like";

/**
 * 查询
 * @method GET
 */
export interface Request {

    /**
     * 目标
     */
    tid: int64

    /**
     * 项ID 默认 0
     */
    iid?: int64

    /**
     * 用户ID
     */
    uid?: int64

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
     * 赞
     */
    items: Like[]

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

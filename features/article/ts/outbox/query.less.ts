
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Page } from "../Query";
import { Outbox } from "../Outbox";

/**
 * 查询发件箱
 * @method GET
 */
export interface Request {

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 是否发布
     */
    isPublished?: boolean

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


export interface OutboxQueryData {
    
    /**
     * 发件
     */
    items: Outbox[]

    /**
     * 分页
     */
    page?: Page
}


export interface Response extends BaseResponse {
    data?: OutboxQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

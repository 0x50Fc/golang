
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";

/**
 * 查询发件箱数量
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

}


export interface OutboxCountData {

    /**
     * 总记录数
     */
    total: int32
}


export interface Response extends BaseResponse {
    data?: OutboxCountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

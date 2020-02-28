
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Page } from '../Query';
import { User } from '../User';

/**
 * 查询评论的用户
 * @method GET
 */
export interface Request {

    /**
     * 评论目标ID
     */
    eid: int64

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

    /**
     * 最大时间
     */
    maxtime?: string

    /**
     * 最小时间
     */
    mintime?: string

}

export interface UserQueryData {

    items: User[]

    page ?: Page

}

export interface Response extends BaseResponse {
    data?: UserQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

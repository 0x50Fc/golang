
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { QueryData } from '../Query';
import { Comment } from "../Comment";

/**
 * 查询回收站中的评论
 * @method GET
 */
export interface Request {

    /**
     * 评论ID
     */
    id?: int64

    /**
     * 父级别
     */
    pid?: int64

    /**
     * 评论目标ID
     */
    eid: int64

    /**
     * 用户ID,0不验证
     */
    uid?: int64

    /**
     * 内容模糊查询
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

export interface Response extends BaseResponse {
    data?: QueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

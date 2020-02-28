
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { CountData } from '../Query';
import { Comment } from "../Comment";

/**
 * 获取回收站中评论数量
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


}

export interface Response extends BaseResponse {
    data?: CountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

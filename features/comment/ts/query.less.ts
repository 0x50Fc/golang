
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { QueryData } from './Query';
import { Comment } from "./Comment";

/**
 * 查询评论
 * @method GET
 */
export interface Request {

    /**
     * 评论ID
     */
    id?: int64

    /**
     * 评论ID 多个逗号分割
     */
    ids?: string

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
     * 用户ID,0不验证 不获取博主的评论
     */
    bloguid?: int64

    /**
     * 内容模糊查询
     */
    q?: string


    /**
     * path模糊查询,查询一个评论下的所有回复
     */
    path?: string

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

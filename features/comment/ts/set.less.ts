
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Comment } from "./Comment";

/**
 * 修改评论
 * @method POST
 */
interface Request {

    /**
     * 评论ID
     */
    id: int64

    /**
     * 评论目标ID
     */
    eid: int64

    /**
     * 用户ID,0不验证
     */
    uid?: int64

    /**
     * 内容
     */
    body?: string

    /**
     * 其他选项 JSON 叠加
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Comment
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

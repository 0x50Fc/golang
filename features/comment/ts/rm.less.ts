
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Comment } from "./Comment";

/**
 * 删除评论
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

}

interface Response extends BaseResponse {
    data?: Comment
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

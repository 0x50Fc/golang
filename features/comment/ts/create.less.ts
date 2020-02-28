
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Comment } from "./Comment";

/**
 * 创建评论
 * @method POST
 */
interface Request {

    /**
     * 父级ID
     */
    pid: int64

    /**
     * 评论目标ID
     * @required true
     */
    eid: int64

    /**
     * 用户ID
     * @required true
     */
    uid: int64

    /**
     * 内容
     * @required true
     */
    body: string

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

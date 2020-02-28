
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { FeedStatus, Feed } from "./Feed";

/**
 * 修改动态
 * @method POST
 */
interface Request {

    /**
     * 动态ID
     */
    id: int64

    /**
     * 状态
     */
    status?: FeedStatus

    /**
     * 内容
     */
    body?: string

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

    /**
     * 发布时间
     */
    ctime?: int64
}

interface Response extends BaseResponse {
    data?: Feed
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

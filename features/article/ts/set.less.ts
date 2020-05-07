
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Article } from "./Article";

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
     * 状态 多个逗号分割
     */
    state?: string

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
    data?: Article
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

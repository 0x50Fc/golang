
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { Outbox } from "../Outbox";

/**
 * 发布
 * @method POST
 */
interface Request {

    /**
     * 草稿ID
     */
    id: int64

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 内容
     */
    body?: string

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Outbox
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

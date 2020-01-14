
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Inbox } from "./Inbox";

/**
 * 修改
 * @method POST
 */
interface Request {

    /**
     * ID
     */
    id: int64

    /**
     * 用户ID
     */
    uid: int64

    /**
     * 发布者ID
     */
    fuid?: int64

    /**
     * 其他数据 JSON 叠加数据
     */
    options?: string

}

interface Response extends BaseResponse {
    data?: Inbox
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}


import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Inbox } from "./Inbox";

/**
 * 删除
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

}

interface Response extends BaseResponse {
    data?: Inbox
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

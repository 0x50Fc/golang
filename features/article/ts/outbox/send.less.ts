
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

}

interface Response extends BaseResponse {
    data?: Outbox
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

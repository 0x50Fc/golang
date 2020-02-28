
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { Feed } from "./Feed";

/**
 * 获取动态
 * @method GET
 */
interface Request {

    /**
     * 动态ID
     */
    id: int64

}

interface Response extends BaseResponse {
    data?: Feed
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

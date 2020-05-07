
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { Article } from "../Article";

/**
 * 恢复动态
 * @method POST
 */
interface Request {

    /**
     * 动态ID
     */
    id: int64

}

interface Response extends BaseResponse {
    data?: Article
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

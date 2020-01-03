
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";

/**
 * 获取ID
 * @method GET
 */
interface Request {

}

interface Response extends BaseResponse {
    data?: int64
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

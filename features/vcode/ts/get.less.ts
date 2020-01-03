
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { VCode } from "./VCode";

/**
 * 获取验证码
 * @method GET
 */
interface Request {

    /**
     * Key
     */
    key: string

}

interface Response extends BaseResponse {
    data?: VCode
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

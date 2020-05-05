
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { VCode } from "./VCode";

/**
 * 校验验证码
 * @method POST
 */
interface Request {

    /**
     * Key
     */
    key: string

    /**
     * 数字验证码
     * @length 12
     */
    code?: string

    /**
     * 32位 HASH
     * @length 32
     */
    hash?: string

}

interface Response extends BaseResponse {
    data?: VCode
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

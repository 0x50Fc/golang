
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { VCode } from "./VCode";

/**
 * 创建验证码
 * @method POST
 */
interface Request {

    /**
     * Key
     */
    key: string

    /**
     * 超时时间(秒)
     */
    expires: number

    /**
     * 验证码长度 默认 4
     */
    length?: number
}

interface Response extends BaseResponse {
    data?: VCode
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

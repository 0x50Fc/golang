
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";

/**
 * 验证初始化
 * @method POST
 */
interface Request {

    /**
     * 极验ID
     */
    captchaId: string

    /**
     * 验证唯一键
     */
    key: string

    /**
     * 超时时间
     */
    expires: int32
}

interface Response extends BaseResponse {
    data?: any
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

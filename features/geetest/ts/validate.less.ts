
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 ,int32} from "./lib/less";

/**
 * 验证
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

    challenge: string

    validate: string

    seccode: string

}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

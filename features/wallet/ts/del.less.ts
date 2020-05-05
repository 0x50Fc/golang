
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { VCode } from "./VCode";

/**
 * 删除验证码
 * @method POST
 */
interface Request {

    /**
     * Key
     */
    key: string

}

interface Response extends BaseResponse {

}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

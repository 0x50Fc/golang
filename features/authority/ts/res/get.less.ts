
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64 } from "../lib/less";
import { Res } from "../Res";

/**
 * 获取资源
 * @method GET
 */
interface Request {

    /**
     * 资源ID
     */
    id?: int64

    /**
     * 资源路径
     */
    path?: string
}

interface Response extends BaseResponse {
    data?: Res
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

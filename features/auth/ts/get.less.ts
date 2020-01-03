
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64 } from "./lib/less";
import { AuthType, Auth } from "./Auth";

/**
 * 获取
 * @method GET
 */
interface Request {

    /**
     * 键值
     */
    key: string

    /**
     * 类型
     */
    type?: AuthType

}

interface Response extends BaseResponse {
    auth?: Auth
    data?: any
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}

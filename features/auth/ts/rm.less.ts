
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { AuthType, Auth } from "./Auth";

/**
 * 删除
 * @method POST
 */
interface Request {

    /**
     * 键值
     */
    key: string

}

interface Response extends BaseResponse {
    auth?: Auth
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
